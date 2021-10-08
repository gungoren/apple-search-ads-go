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

package util

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gungoren/apple-search-ads-go/asa"
	"io"
	"log"
	"os"
	"time"
)

var (
	orgID          = flag.String("oid", "", "org ID")
	keyID          = flag.String("kid", "", "key ID")
	teamID         = flag.String("tid", "", "team ID")
	clientID       = flag.String("cid", "", "client ID")
	privateKey     = flag.String("privatekey", "", "private key used to sign authorization token")
	privateKeyPath = flag.String("privatekeypath", "", "path to a private key used to sign authorization token")
)

// TokenConfig creates the auth transport using the required information
func TokenConfig() (auth *asa.AuthTransport, err error) {
	var secret []byte
	if *privateKey != "" {
		secret = []byte(*privateKey)
	} else if *privateKeyPath != "" {
		// Read private key file as []byte
		secret, err = os.ReadFile(*privateKeyPath)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("no private key provided to either the -privatekey or -privatekeypath flags")
	}

	// Create the token using the required information
	auth, err = asa.NewTokenConfig(*orgID, *keyID, *teamID, *clientID, 20*time.Minute, secret)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// GetCampaign returns a single asa.Campaign by filtering by its selector on Apple Search Ads
func GetCampaign(ctx context.Context, client *asa.Client, params *asa.Selector) (*asa.Campaign, error) {
	apps, _, err := client.Campaigns.FindCampaigns(ctx, params)
	if err != nil {
		return nil, err
	} else if len(apps.Campaigns) == 0 {
		return nil, fmt.Errorf("query for campaign returned no campaign")
	}
	return &apps.Campaigns[0], nil
}

// Close closes an open descriptor.
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
