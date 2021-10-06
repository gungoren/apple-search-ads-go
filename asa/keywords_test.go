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

func TestCreateTargetingKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &KeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.CreateTargetingKeywords(ctx, 1, 99, []*Keyword{})
	})
}

func TestFindTargetingKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &KeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.FindTargetingKeywords(ctx, 1, &Selector{})
	})
}

func TestGetTargetingKeyword(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &KeywordResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.GetTargetingKeyword(ctx, 1, 99, 10001)
	})
}

func TestGetAllTargetingKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &KeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.GetAllTargetingKeywords(ctx, 1, 99, &GetAllTargetingKeywordsQuery{})
	})
}

func TestUpdateTargetingKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &KeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.UpdateTargetingKeywords(ctx, 1, 99, []*KeywordUpdateRequest{})
	})
}

func TestCreateNegativeKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &NegativeKeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.CreateNegativeKeywords(ctx, 1, []*NegativeKeyword{})
	})
}

func TestCreateAdGroupNegativeKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &NegativeKeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.CreateAdGroupNegativeKeywords(ctx, 1, 99, []*NegativeKeyword{})
	})
}

func TestFindNegativeKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &NegativeKeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.FindNegativeKeywords(ctx, 1, &Selector{})
	})
}

func TestFindAdGroupNegativeKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &NegativeKeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.FindAdGroupNegativeKeywords(ctx, 1, &Selector{})
	})
}

func TestGetNegativeKeyword(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &NegativeKeywordResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.GetNegativeKeyword(ctx, 1, 10001)
	})
}

func TestGetAdGroupNegativeKeyword(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &NegativeKeywordResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.GetAdGroupNegativeKeyword(ctx, 1, 99, 10001)
	})
}

func TestGetAllNegativeKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &NegativeKeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.GetAllNegativeKeywords(ctx, 1, &GetAllNegativeKeywordsQuery{})
	})
}

func TestGetAllAdGroupNegativeKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &NegativeKeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.GetAllAdGroupNegativeKeywords(ctx, 1, 99, &GetAllNegativeKeywordsQuery{})
	})
}

func TestUpdateNegativeKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &NegativeKeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.UpdateNegativeKeywords(ctx, 1, []*NegativeKeyword{})
	})
}

func TestUpdateAdGroupNegativeKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &NegativeKeywordListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.UpdateAdGroupNegativeKeywords(ctx, 1, 99, []*NegativeKeyword{})
	})
}

func TestDeleteNegativeKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &IntegerResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.DeleteNegativeKeywords(ctx, 1, []int64{})
	})
}

func TestDeleteAdGroupNegativeKeywords(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &IntegerResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Keywords.DeleteAdGroupNegativeKeywords(ctx, 1, 99, []int64{})
	})
}
