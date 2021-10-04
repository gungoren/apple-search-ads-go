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
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
//TODO: handle this error
func TestNewTokenConfig(t *testing.T) {
	t.Parallel()

	// This is a key that I generated solely for mocking purposes. This is not a
	// real secret, so don't get any funny ideas. If you need to regenerate it,
	// run this openssl command in a shell and copy the contents of key.pem to the string:
	//
	//   openssl ecparam -name prime256v1 -genkey -noout -out private-key.pem
	//
	// This will generate the ASN.1 PKCS#8 representation of the private key needed
	// to create a valid token. If you are looking at this test to see how to make a key,
	// reference Apple's documentation on this subject instead.
	//
	// https://developer.apple.com/documentation/apple_search_ads/implementing_oauth_for_the_apple_search_ads_api
	var privPEMData = []byte(`
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEID03iksZiO87BzijZdC2dz4GF+LxK6hOCVVsjxunKqcQoAoGCCqGSM49
AwEHoUQDQgAEFqu+MZYeAD3Zx9PfuAFG+fNmuy1IKXVtslFwgFgnz5BtuuZuycUP
LWfXSp67hz35UIbgO6NANWf4tzZ6fhTThA==
-----END EC PRIVATE KEY-----
`)

	token, err := NewTokenConfig("TEST_ORG_ID", "TEST_KEY_ID", "TEST_TEAM_ID", "TEST_CLIENT_ID", 20*time.Minute, privPEMData)
	assert.NoError(t, err)

	tok, err := token.jwtGenerator.Token()
	assert.NoError(t, err)

	components := strings.Split(tok, ".")
	assert.Equal(t, 3, len(components))

	tokCached, err := token.jwtGenerator.Token()
	assert.NoError(t, err)
	assert.Equal(t, tok, tokCached)
}*/

func TestNewTokenConfigBadPEM(t *testing.T) {
	t.Parallel()

	_, err := NewTokenConfig("TEST_ORG_ID", "TEST_KEY_ID", "TEST_TEAM_ID", "TEST_CLIENT_ID", 20*time.Minute, []byte("TEST"))
	assert.Error(t, err, "Expected error for invalid PEM, got nil")
}

func TestNewTokenConfigPrivateKeyNotPKCS8(t *testing.T) {
	t.Parallel()

	var badKey = []byte(`
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIXpcxwADKgwJSwxz24ypAMDFUHPrirqhcx0vimrl9L2oAoGCCqGSM49
AwEHoUQDQgAE7Ee8TlNaDqWa6O/Yw/nqHVEiJwYS+wt5cd7DC85nhsDxaU8M2asd
oH1YGuY57H3BQ3zLPVPsN+A8xnInGDa8yQ
-----END EC PRIVATE KEY-----
`)

	_, err := NewTokenConfig("TEST_ORG_ID", "TEST_KEY_ID", "TEST_TEAM_ID", "TEST_CLIENT_ID", 20*time.Minute, badKey)
	assert.Error(t, err, "Expected error for non-PKCS8 PEM, got nil")
}

func TestAuthTransport(t *testing.T) {
	t.Parallel()

	token := "TEST.TEST.TEST"
	transport := AuthTransport{
		jwtGenerator: &mockJWTGenerator{token: token},
	}
	client := transport.Client()

	req, _ := http.NewRequest("GET", "", nil)
	_, _ = client.Do(req) // nolint: bodyclose

	got, want := req.Header.Get("Authorization"), fmt.Sprintf("Bearer %s", token)
	assert.Equal(t, want, got)
}

func TestAuthTransportCustomTransport(t *testing.T) {
	t.Parallel()

	token := "TEST.TEST.TEST"
	transport := AuthTransport{
		jwtGenerator: &mockJWTGenerator{token: token},
	}
	client := transport.Client()

	req, _ := http.NewRequest("GET", "", nil) // nolint: noctx
	_, _ = client.Do(req)                     // nolint: bodyclose

	got, want := req.Header.Get("Authorization"), fmt.Sprintf("Bearer %s", token)
	assert.Equal(t, want, got)
}

type mockJWTGenerator struct {
	token string
}

func (g *mockJWTGenerator) Token() (string, error) {
	return g.token, nil
}

func (g *mockJWTGenerator) IsValid() bool {
	return true
}
