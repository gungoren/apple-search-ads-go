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

func TestGetAllAdGroups(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AdGroups.GetAllAdGroups(ctx, 1, &GetAllAdGroupsQuery{})
	})
}

func TestGetAdGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AdGroups.GetAdGroup(ctx, 1, 99)
	})
}

func TestDeleteAdGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.AdGroups.DeleteAdGroup(ctx, 1, 99)
	})
}

func TestUpdateAdGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AdGroups.UpdateAdGroup(ctx, 1, 99, &AdGroupUpdateRequest{})
	})
}

func TestCreateAdGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AdGroups.CreateAdGroup(ctx, 1, &AdGroup{})
	})
}

func TestFindAdGroups(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AdGroups.FindAdGroups(ctx, 1, &Selector{})
	})
}
