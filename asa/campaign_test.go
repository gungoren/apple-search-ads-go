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
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCampaigns(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &CampaignListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Campaigns.GetAllCampaigns(ctx, &GetAllCampaignQuery{})
	})
}

func TestGetCampaign(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &CampaignResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Campaigns.GetCampaign(ctx, 1)
	})
}

func TestDeleteCampaign(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Campaigns.DeleteCampaign(ctx, 1)
	})
}

func TestUpdateCampaign(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &CampaignResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Campaigns.UpdateCampaign(ctx, 1, &UpdateCampaignRequest{})
	})
}

func TestCreateCampaign(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &CampaignResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Campaigns.CreateCampaign(ctx, &Campaign{})
	})
}

func TestFindCampaigns(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &CampaignListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Campaigns.FindCampaigns(ctx, &Selector{})
	})
}

func TestDeserializationWithUnAuthorized(t *testing.T) {
	t.Parallel()

	content, err := ioutil.ReadFile("../test/response_body_json_files/error_with_unauthorized.json")
	assert.NoError(t, err)

	model := &APIErrorResponse{}

	err = json.Unmarshal(content, model)
	assert.NoError(t, err)
}

func TestCampaignListResponseDeserialization(t *testing.T) {
	t.Parallel()

	content, err := ioutil.ReadFile("../test/response_body_json_files/campaign_list_response.json")
	assert.NoError(t, err)

	model := &CampaignListResponse{}

	err = json.Unmarshal(content, model)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(model.Campaigns))
	assert.Equal(t, 1, model.Pagination.TotalResults)
	campaign := model.Campaigns[0]
	assert.Equal(t, PaymentModelLoc, campaign.PaymentModel)
}

func TestCampaignListResponseWithEmptyDeserialization(t *testing.T) {
	t.Parallel()

	content, err := ioutil.ReadFile("../test/response_body_json_files/campaign_list_response_with_empty.json")
	assert.NoError(t, err)

	model := &CampaignListResponse{}

	err = json.Unmarshal(content, model)
	assert.NoError(t, err)

	assert.Equal(t, 0, len(model.Campaigns))
	assert.Equal(t, 0, model.Pagination.TotalResults)
}
