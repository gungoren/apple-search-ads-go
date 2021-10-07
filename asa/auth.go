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
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// ErrMissingPEM happens when the bytes cannot be decoded as a PEM block.
var ErrMissingPEM = errors.New("no PEM blob found")

// ErrInvalidPrivateKey happens when a key cannot be parsed as a ECDSA PKCS8 private key.
var ErrInvalidPrivateKey = errors.New("key could not be parsed as a valid ecdsa.PrivateKey")

// ErrHTTPTokenBadRequest happens when apple generate token http request failed.
var ErrHTTPTokenBadRequest = errors.New("generate auth token failed with")

// AuthTransport is an http.RoundTripper implementation that stores the JWT created.
// If the token expires, the Rotate function should be called to update the stored token.
type AuthTransport struct {
	Transport    http.RoundTripper
	jwtGenerator jwtGenerator
	orgID        string
}

type jwtGenerator interface {
	Token() (string, error)
	AccessToken() (string, error)
	IsTokenValid() bool
	IsAccessTokenValid() bool
	Client() *authClient
}

type standardJWTGenerator struct {
	keyID          string
	issuerID       string
	clientID       string
	expireDuration time.Duration
	privateKey     *ecdsa.PrivateKey

	accessToken *accessToken
	token       string
	client      *authClient
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

	return &AuthTransport{
		Transport:    newTransport(),
		jwtGenerator: gen,
		orgID:        orgID,
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
	token, err := t.jwtGenerator.AccessToken()
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
	if g.IsTokenValid() {
		return g.token, nil
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, g.claims())
	t.Header["kid"] = g.keyID

	token, err := t.SignedString(g.privateKey)
	if err != nil {
		return "", err
	}

	g.token = token

	return token, nil
}

func (g *standardJWTGenerator) AccessToken() (string, error) {
	if g.IsAccessTokenValid() {
		return g.accessToken.AccessToken, nil
	}

	token, err := g.Token()
	if err != nil {
		return "", err
	}

	accessTkn, err := g.generateAccessToken(token)
	if err != nil {
		return "", err
	}

	g.accessToken = accessTkn

	return accessTkn.AccessToken, nil
}

func (g *standardJWTGenerator) Client() *authClient {
	if g.client == nil {
		g.client = &authClient{
			client:  &http.Client{},
			baseURL: defaultAuthURL,
		}
	}

	return g.client
}

type authClient struct {
	client  *http.Client
	baseURL string
}

func (g *standardJWTGenerator) generateAccessToken(token string) (*accessToken, error) {
	authClient := g.Client()
	url := fmt.Sprintf("%s/oauth2/token?grant_type=client_credentials&client_id=%s&client_secret=%s&scope=searchadsorg", authClient.baseURL, g.clientID, token)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Host", "appleid.apple.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := authClient.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("http error %w: %s", ErrHTTPTokenBadRequest, string(b))
	}

	accessToken := &accessToken{}
	if err := json.Unmarshal(b, accessToken); err != nil {
		return nil, err
	}

	accessToken.expiresAfter = time.Now().Add(time.Second * time.Duration(accessToken.ExpiresIn))

	return accessToken, nil
}

func (g *standardJWTGenerator) IsTokenValid() bool {
	if g.token == "" {
		return false
	}

	parsed, err := jwt.Parse(
		g.token,
		jwt.KnownKeyfunc(jwt.SigningMethodES256, g.privateKey),
		jwt.WithAudience("https://appleid.apple.com"),
		jwt.WithIssuer(g.issuerID),
	)
	if err != nil {
		return false
	}

	return parsed.Valid
}

func (g *standardJWTGenerator) IsAccessTokenValid() bool {
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
