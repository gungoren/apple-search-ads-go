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
	"testing"
)

func TestGetCampaignLevelReports(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ReportingResponseBody{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetCampaignLevelReports(ctx, &ReportingRequest{})
	})
}

func TestGetAdGroupLevelReports(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ReportingResponseBody{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetAdGroupLevelReports(ctx, 1, &ReportingRequest{})
	})
}

func TestGetKeywordLevelReports(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ReportingResponseBody{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetKeywordLevelReports(ctx, 1, &ReportingRequest{})
	})
}

func TestGetSearchTermLevelReports(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ReportingResponseBody{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetSearchTermLevelReports(ctx, 1, &ReportingRequest{})
	})
}

func TestGetCreativeSetLevelReports(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ReportingResponseBody{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetCreativeSetLevelReports(ctx, 1, &ReportingRequest{})
	})
}