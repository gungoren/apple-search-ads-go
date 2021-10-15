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
)

// AppService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/search_apps_and_geolocations
type AppService service

// SearchAppsQuery defines query parameter for SearchApps endpoint.
type SearchAppsQuery struct {
	Limit           int32  `url:"limit,omitempty"`
	Offset          int32  `url:"offset,omitempty"`
	Query           string `url:"query,omitempty"`
	ReturnOwnedApps bool   `url:"returnOwnedApps,omitempty"`
}

// AppInfoListResponse is the response details of app search requests
//
// https://developer.apple.com/documentation/apple_search_ads/appinfolistresponse
type AppInfoListResponse struct {
	AppInfos   []*AppInfo         `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// AppInfo is the response to an app search request
//
// https://developer.apple.com/documentation/apple_search_ads/appinfo
type AppInfo struct {
	AdamID               int64    `json:"adamId,omitempty"`
	AppName              string   `json:"appName,omitempty"`
	CountryOrRegionCodes []string `json:"countryOrRegionCodes,omitempty"`
	DeveloperName        string   `json:"developerName"`
}

// SearchApps Searches for iOS apps to promote in a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/search_for_ios_apps
func (s *AppService) SearchApps(ctx context.Context, params *SearchAppsQuery) (*AppInfoListResponse, *Response, error) {
	url := "search/apps"
	res := new(AppInfoListResponse)
	resp, err := s.client.get(ctx, url, &params, res)

	return res, resp, err
}
