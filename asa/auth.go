/**
Copyright (C) 2021 Mehmet Gungoren.
This file is part of apple-search-ads-go, a package for working with Apple's
Search Ads API.
apple-search-ads-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
apple-search-ads-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with apple-search-ads-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asa

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"io/ioutil"
	"net/http"
	"time"
)

// ErrMissingPEM happens when the bytes cannot be decoded as a PEM block.
var ErrMissingPEM = errors.New("no PEM blob found")

// ErrInvalidPrivateKey happens when a key cannot be parsed as a ECDSA PKCS8 private key.
var ErrInvalidPrivateKey = errors.New("key could not be parsed as a valid ecdsa.PrivateKey")

// AuthTransport is an http.RoundTripper implementation that stores the JWT created.
// If the token expires, the Rotate function should be called to update the stored token.
type AuthTransport struct {
	Transport    http.RoundTripper
	jwtGenerator jwtGenerator
	orgID       string
}

type jwtGenerator interface {
	Token() (string, error)
	IsValid() bool
}

type standardJWTGenerator struct {
	keyID          string
	issuerID       string
	clientID       string
	expireDuration time.Duration
	privateKey     *ecdsa.PrivateKey

	accessToken *accessToken
}

type accessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`

	expiresAfter time.Time
}

// NewTokenConfig returns a new AuthTransport instance that customizes the Authentication header of the request during transport.
// It can be customized further by supplying a custom http.RoundTripper instance to the Transport field.
func NewTokenConfig(orgID string, keyID string, teamID string, clientID string, expireDuration time.Duration, privateKey []byte) (*AuthTransport, error) {
	key, err := parsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	gen := &standardJWTGenerator{
		keyID:          keyID,
		issuerID:       teamID,
		clientID:       clientID,
		privateKey:     key,
		expireDuration: expireDuration,
	}
	_, err = gen.Token()

	return &AuthTransport{
		Transport:    newTransport(),
		jwtGenerator: gen,
		orgID: orgID,
	}, err
}

func parsePrivateKey(blob []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(blob)
	if block == nil {
		return nil, ErrMissingPEM
	}

	parsedKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return parsedKey, nil
}

// RoundTrip implements the http.RoundTripper interface to set the Authorization header.
func (t AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.jwtGenerator.Token()
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("X-AP-Context", fmt.Sprintf("orgId=%s", t.orgID))

	return t.transport().RoundTrip(req)
}

// Client returns a new http.Client instance for use with apple_search_ads.Client.
func (t *AuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *AuthTransport) transport() http.RoundTripper {
	if t.Transport == nil {
		t.Transport = newTransport()
	}

	return t.Transport
}

func (g *standardJWTGenerator) Token() (string, error) {
	if g.IsValid() {
		return g.accessToken.AccessToken, nil
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, g.claims())
	t.Header["kid"] = g.keyID

	token, err := t.SignedString(g.privateKey)
	if err != nil {
		return "", err
	}

	accessTkn, err := g.generateAccessToken(token)
	if err != nil {
		return "", err
	}

	g.accessToken = accessTkn

	return token, nil
}

func (g *standardJWTGenerator) generateAccessToken(token string) (*accessToken, error) {
	url := fmt.Sprintf("https://appleid.apple.com/auth/oauth2/token?grant_type=client_credentials&client_id=%s&client_secret=%s&scope=searchadsorg", g.clientID, token)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Host", "appleid.apple.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, errors.New(string(b))
	}

	accessToken := &accessToken{}
	if err := json.Unmarshal(b, accessToken); err != nil {
		return nil, err
	}
	expiresTime := time.Now().Add(time.Second * time.Duration(accessToken.ExpiresIn))
	accessToken.expiresAfter = expiresTime
	return accessToken, nil
}

func (g *standardJWTGenerator) IsValid() bool {
	if g.accessToken == nil || g.accessToken.AccessToken == "" {
		return false
	}

	if g.accessToken.expiresAfter.Before(time.Now()) {
		return false
	}

	return true
}

func (g *standardJWTGenerator) claims() jwt.Claims {
	expiry := time.Now().Add(g.expireDuration)

	return jwt.StandardClaims{
		Audience:  jwt.ClaimStrings{"https://appleid.apple.com"},
		Subject:   g.clientID,
		Issuer:    g.issuerID,
		ExpiresAt: jwt.At(expiry),
	}
}

func newTransport() http.RoundTripper {
	return &http.Transport{
		IdleConnTimeout: defaultTimeout,
	}
}
