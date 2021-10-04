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
	"fmt"
)

// KeywordService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/targeting_keywords_and_negative_keywords
type KeywordService service

// KeywordMatchType defines model for Keyword Match Type.
//
// https://developer.apple.com/documentation/apple_search_ads/keyword
type KeywordMatchType string

const (
	KeywordMatchTypeBroad KeywordMatchType = "Broad"
	KeywordMatchTypeExact KeywordMatchType = "Exact"
)

// KeywordStatus defines model for Keyword Status.
//
// https://developer.apple.com/documentation/apple_search_ads/keyword
type KeywordStatus string

const (
	KeywordStatusActive KeywordStatus = "ACTIVE"
	KeywordStatusPaused KeywordStatus = "PAUSED"
)

// Keyword defines model for Keyword.
//
// https://developer.apple.com/documentation/apple_search_ads/keyword
type Keyword struct {
	AdGroupId        int64            `json:"adGroupId,omitempty"`
	BidAmount        Money            `json:"bidAmount,omitempty"`
	Deleted          bool             `json:"deleted,omitempty"`
	Id               int64            `json:"id,omitempty"`
	MatchType        KeywordMatchType `json:"matchType,omitempty"`
	ModificationTime DateTime         `json:"modificationTime,omitempty"`
	Status           KeywordStatus    `json:"status,omitempty"`
	Text             string           `json:"text,omitempty"`
}

// KeywordListResponse defines model for Keyword List Response.
//
//https://developer.apple.com/documentation/apple_search_ads/keywordlistresponse
type KeywordListResponse struct {
	Keywords   []*Keyword         `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	PageDetail *PageDetail        `json:"pageDetail,omitempty"`
}

// CreateTargetingKeywords Creates targeting keywords in ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/create_targeting_keywords
func (s *KeywordService) CreateTargetingKeywords(ctx context.Context, campaignId int64, adGroupId int64, keyword []*Keyword) (*KeywordListResponse, *Response, error) {
	url := fmt.Sprintf("/campaigns/%d/adgroups/%d/targetingkeywords/bulk", campaignId, adGroupId)
	res := new(KeywordListResponse)
	resp, err := s.client.post(ctx, url, keyword, res)
	return res, resp, err
}

// FindTargetingKeywords Fetches targeting keywords in a campaign’s ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/create_targeting_keywords
func (s *KeywordService) FindTargetingKeywords(ctx context.Context, campaignId int64, selector *Selector) (*KeywordListResponse, *Response, error) {
	url := fmt.Sprintf("/campaigns/%d/adgroups/targetingkeywords/find", campaignId)
	res := new(KeywordListResponse)
	resp, err := s.client.post(ctx, url, selector, res)
	return res, resp, err
}

// KeywordResponse is a container for the targeting keywords response body.
//
// https://developer.apple.com/documentation/apple_search_ads/keywordresponse
type KeywordResponse struct {
	Data       *Keyword           `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	PageDetail *PageDetail        `json:"pageDetail,omitempty"`
}

// GetTargetingKeyword Fetches a specific targeting keyword in an ad group
//
// https://developer.apple.com/documentation/apple_search_ads/get_a_targeting_keyword_in_an_ad_group
func (s *KeywordService) GetTargetingKeyword(ctx context.Context, campaignId int64, adGroupId int64, keywordId int64) (*KeywordResponse, *Response, error) {
	url := fmt.Sprintf("/campaigns/%d/adgroups/%d/targetingkeywords/%d", campaignId, adGroupId, keywordId)
	res := new(KeywordResponse)
	resp, err := s.client.get(ctx, url, nil, res)
	return res, resp, err
}

// GetAllTargetingKeywordsQuery defines query parameter for GetAllTargetingKeywords endpoint.
type GetAllTargetingKeywordsQuery struct {
	Limit  int32 `url:"limit,omitempty"`
	Offset int32 `url:"offset,omitempty"`
}

// GetAllTargetingKeywords Fetches all targeting keywords in ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/get_all_targeting_keywords_in_an_ad_group
func (s *KeywordService) GetAllTargetingKeywords(ctx context.Context, campaignId int64, adGroupId int64, params *GetAllTargetingKeywordsQuery) (*KeywordListResponse, *Response, error) {
	url := fmt.Sprintf("/campaigns/%d/adgroups/%d/targetingkeywords/", campaignId, adGroupId)
	res := new(KeywordListResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// KeywordUpdateRequest Targeting keyword parameters to use in requests and responses
//
// https://developer.apple.com/documentation/apple_search_ads/keywordupdaterequest
type KeywordUpdateRequest struct {
	AdGroupId        int64            `json:"adGroupId,omitempty"`
	BidAmount        *Money           `json:"bidAmount,omitempty"`
	Deleted          bool             `json:"deleted,omitempty"`
	Id               int64            `json:"id,omitempty"`
	MatchType        KeywordMatchType `json:"matchType"`
	ModificationTime DateTime         `json:"modificationTime"`
}

// UpdateTargetingKeywords Updates targeting keywords in ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/update_targeting_keywords
func (s *KeywordService) UpdateTargetingKeywords(ctx context.Context, campaignId int64, adGroupId int64, updateRequests []*KeywordUpdateRequest) (*KeywordListResponse, *Response, error) {
	url := fmt.Sprintf("/campaigns/%d/adgroups/%d/targetingkeywords/bulk", campaignId, adGroupId)
	res := new(KeywordListResponse)
	resp, err := s.client.put(ctx, url, updateRequests, res)
	return res, resp, err
}

// NegativeKeyword Negative keyword parameters to use in requests and responses
//
// https://developer.apple.com/documentation/apple_search_ads/negativekeyword
type NegativeKeyword struct {
	AdGroupId        int64            `json:"adGroupId,omitempty"`
	CampaignId       int64            `json:"campaignId,omitempty"`
	Deleted          bool             `json:"deleted,omitempty"`
	Id               int64            `json:"id,omitempty"`
	MatchType        KeywordMatchType `json:"matchType,omitempty"`
	ModificationTime DateTime         `json:"modificationTime,omitempty"`
	Status           KeywordStatus    `json:"status,omitempty"`
	Text             string           `json:"text,omitempty"`
}

// NegativeKeywordListResponse The response details of negative keyword requests
//
// https://developer.apple.com/documentation/apple_search_ads/negativekeywordlistresponse
type NegativeKeywordListResponse struct {
	Keywords   []*NegativeKeyword `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	PageDetail *PageDetail        `json:"pageDetail,omitempty"`
}

// CreateNegativeKeywords Creates negative keywords for a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/create_campaign_negative_keywords
func (s *KeywordService) CreateNegativeKeywords(ctx context.Context, campaignId int64, keyword []*NegativeKeyword) (*NegativeKeywordListResponse, *Response, error) {
	url := fmt.Sprintf("/campaigns/%d/negativekeywords/bulk", campaignId)
	res := new(NegativeKeywordListResponse)
	resp, err := s.client.post(ctx, url, keyword, res)
	return res, resp, err
}