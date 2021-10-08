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

// KeywordMatchType is an automated keyword and bidding strategy.
type KeywordMatchType string

const (
	// KeywordMatchTypeBroad is used this value to ensure your ads don’t run on relevant, close variants of a keyword, such as singulars, plurals, misspellings, synonyms, related searches, and phrases that include that term (fully or partially).
	KeywordMatchTypeBroad KeywordMatchType = "Broad"
	// KeywordMatchTypeExact is used this value for the most control over searches your ad may appear in. You can target a specific term and its close variants, such as common misspellings and plurals. Your ad may receive fewer impressions as a result, but your tap-through rates (TTRs) and conversions on those impressions may be higher because you’re reaching users most interested in your app.
	KeywordMatchTypeExact KeywordMatchType = "Exact"
)

// KeywordStatus defines model for Keyword Status.
type KeywordStatus string

const (
	// KeywordStatusActive is for a keyword status on Active state.
	KeywordStatusActive KeywordStatus = "ACTIVE"
	// KeywordStatusPaused is for a keyword status on Paused state.
	KeywordStatusPaused KeywordStatus = "PAUSED"
)

// Keyword defines model for Keyword.
//
// https://developer.apple.com/documentation/apple_search_ads/keyword
type Keyword struct {
	AdGroupID        int64            `json:"adGroupId,omitempty"`
	BidAmount        Money            `json:"bidAmount,omitempty"`
	Deleted          bool             `json:"deleted,omitempty"`
	ID               int64            `json:"id,omitempty"`
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
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// CreateTargetingKeywords Creates targeting keywords in ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/create_targeting_keywords
func (s *KeywordService) CreateTargetingKeywords(ctx context.Context, campaignID int64, adGroupID int64, keyword []*Keyword) (*KeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/bulk", campaignID, adGroupID)
	res := new(KeywordListResponse)
	resp, err := s.client.post(ctx, url, keyword, res)

	return res, resp, err
}

// FindTargetingKeywords Fetches targeting keywords in a campaign’s ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/create_targeting_keywords
func (s *KeywordService) FindTargetingKeywords(ctx context.Context, campaignID int64, selector *Selector) (*KeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/targetingkeywords/find", campaignID)
	res := new(KeywordListResponse)
	resp, err := s.client.post(ctx, url, selector, res)

	return res, resp, err
}

// KeywordResponse is a container for the targeting keywords response body.
//
// https://developer.apple.com/documentation/apple_search_ads/keywordresponse
type KeywordResponse struct {
	Keyword    *Keyword           `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// GetTargetingKeyword Fetches a specific targeting keyword in an ad group
//
// https://developer.apple.com/documentation/apple_search_ads/get_a_targeting_keyword_in_an_ad_group
func (s *KeywordService) GetTargetingKeyword(ctx context.Context, campaignID int64, adGroupID int64, keywordID int64) (*KeywordResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/%d", campaignID, adGroupID, keywordID)
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
func (s *KeywordService) GetAllTargetingKeywords(ctx context.Context, campaignID int64, adGroupID int64, params *GetAllTargetingKeywordsQuery) (*KeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/", campaignID, adGroupID)
	res := new(KeywordListResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// KeywordUpdateRequest Targeting keyword parameters to use in requests and responses
//
// https://developer.apple.com/documentation/apple_search_ads/keywordupdaterequest
type KeywordUpdateRequest struct {
	AdGroupID        int64            `json:"adGroupId,omitempty"`
	BidAmount        *Money           `json:"bidAmount,omitempty"`
	Deleted          bool             `json:"deleted,omitempty"`
	ID               int64            `json:"id,omitempty"`
	MatchType        KeywordMatchType `json:"matchType"`
	ModificationTime DateTime         `json:"modificationTime"`
}

// UpdateTargetingKeywords Updates targeting keywords in ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/update_targeting_keywords
func (s *KeywordService) UpdateTargetingKeywords(ctx context.Context, campaignID int64, adGroupID int64, updateRequests []*KeywordUpdateRequest) (*KeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/bulk", campaignID, adGroupID)
	res := new(KeywordListResponse)
	resp, err := s.client.put(ctx, url, updateRequests, res)

	return res, resp, err
}

// NegativeKeyword Negative keyword parameters to use in requests and responses
//
// https://developer.apple.com/documentation/apple_search_ads/negativekeyword
type NegativeKeyword struct {
	AdGroupID        int64            `json:"adGroupId,omitempty"`
	CampaignID       int64            `json:"campaignId,omitempty"`
	Deleted          bool             `json:"deleted,omitempty"`
	ID               int64            `json:"id,omitempty"`
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
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// CreateNegativeKeywords Creates negative keywords for a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/create_campaign_negative_keywords
func (s *KeywordService) CreateNegativeKeywords(ctx context.Context, campaignID int64, keyword []*NegativeKeyword) (*NegativeKeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/bulk", campaignID)
	res := new(NegativeKeywordListResponse)
	resp, err := s.client.post(ctx, url, keyword, res)

	return res, resp, err
}

// CreateAdGroupNegativeKeywords Creates negative keywords in a specific ad group
//
// https://developer.apple.com/documentation/apple_search_ads/create_ad_group_negative_keywords
func (s *KeywordService) CreateAdGroupNegativeKeywords(ctx context.Context, campaignID int64, adGroupID int64, keyword []*NegativeKeyword) (*NegativeKeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/bulk", campaignID, adGroupID)
	res := new(NegativeKeywordListResponse)
	resp, err := s.client.post(ctx, url, keyword, res)

	return res, resp, err
}

// FindNegativeKeywords Fetches negative keywords for campaigns
//
// https://developer.apple.com/documentation/apple_search_ads/find_campaign_negative_keywords
func (s *KeywordService) FindNegativeKeywords(ctx context.Context, campaignID int64, selector *Selector) (*NegativeKeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/find", campaignID)
	res := new(NegativeKeywordListResponse)
	resp, err := s.client.post(ctx, url, selector, res)

	return res, resp, err
}

// FindAdGroupNegativeKeywords Fetches negative keywords in a campaign’s ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/find_ad_group_negative_keywords
func (s *KeywordService) FindAdGroupNegativeKeywords(ctx context.Context, campaignID int64, selector *Selector) (*NegativeKeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/negativekeywords/find", campaignID)
	res := new(NegativeKeywordListResponse)
	resp, err := s.client.post(ctx, url, selector, res)

	return res, resp, err
}

// NegativeKeywordResponse is a container for the negative keyword response body
//
// https://developer.apple.com/documentation/apple_search_ads/negativekeywordresponse
type NegativeKeywordResponse struct {
	NegativeKeyword *NegativeKeyword   `json:"data,omitempty"`
	Error           *ErrorResponseBody `json:"error,omitempty"`
	Pagination      *PageDetail        `json:"pagination,omitempty"`
}

// GetNegativeKeyword Fetches a specific negative keyword in a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/get_a_campaign_negative_keyword
func (s *KeywordService) GetNegativeKeyword(ctx context.Context, campaignID int64, keywordID int64) (*NegativeKeywordResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/%d", campaignID, keywordID)
	res := new(NegativeKeywordResponse)
	resp, err := s.client.get(ctx, url, nil, res)

	return res, resp, err
}

// GetAdGroupNegativeKeyword Fetches a specific negative keyword in an ad group
//
// https://developer.apple.com/documentation/apple_search_ads/get_an_ad_group_negative_keyword
func (s *KeywordService) GetAdGroupNegativeKeyword(ctx context.Context, campaignID int64, adGroupID int64, keywordID int64) (*NegativeKeywordResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/%d", campaignID, adGroupID, keywordID)
	res := new(NegativeKeywordResponse)
	resp, err := s.client.get(ctx, url, nil, res)

	return res, resp, err
}

// GetAllNegativeKeywordsQuery defines query parameter for GetAllNegativeKeywords endpoint.
type GetAllNegativeKeywordsQuery struct {
	Limit  int32 `url:"limit,omitempty"`
	Offset int32 `url:"offset,omitempty"`
}

// GetAllNegativeKeywords Fetches all negative keywords in a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/get_all_campaign_negative_keywords
func (s *KeywordService) GetAllNegativeKeywords(ctx context.Context, campaignID int64, params *GetAllNegativeKeywordsQuery) (*NegativeKeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/", campaignID)
	res := new(NegativeKeywordListResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetAllAdGroupNegativeKeywords Fetches all negative keywords in ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/get_all_ad_group_negative_keywords
func (s *KeywordService) GetAllAdGroupNegativeKeywords(ctx context.Context, campaignID int64, adGroupID int64, params *GetAllNegativeKeywordsQuery) (*NegativeKeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/", campaignID, adGroupID)
	res := new(NegativeKeywordListResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UpdateNegativeKeywords Updates negative keywords in a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/update_campaign_negative_keywords
func (s *KeywordService) UpdateNegativeKeywords(ctx context.Context, campaignID int64, updateRequests []*NegativeKeyword) (*NegativeKeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/bulk", campaignID)
	res := new(NegativeKeywordListResponse)
	resp, err := s.client.put(ctx, url, updateRequests, res)

	return res, resp, err
}

// UpdateAdGroupNegativeKeywords Updates negative keywords in an ad group
//
// https://developer.apple.com/documentation/apple_search_ads/update_ad_group_negative_keywords
func (s *KeywordService) UpdateAdGroupNegativeKeywords(ctx context.Context, campaignID int64, adGroupID int64, updateRequests []*NegativeKeyword) (*NegativeKeywordListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/bulk", campaignID, adGroupID)
	res := new(NegativeKeywordListResponse)
	resp, err := s.client.put(ctx, url, updateRequests, res)

	return res, resp, err
}

// IntegerResponse is a common integer type response
//
// https://developer.apple.com/documentation/apple_search_ads/integerresponse
type IntegerResponse struct {
	Data       int32              `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// DeleteNegativeKeywords Deletes negative keywords from a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/delete_campaign_negative_keywords
func (s *KeywordService) DeleteNegativeKeywords(ctx context.Context, campaignID int64, keywordIds []int64) (*IntegerResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/negativekeywords/delete/bulk", campaignID)
	res := new(IntegerResponse)
	resp, err := s.client.post(ctx, url, keywordIds, res)

	return res, resp, err
}

// DeleteAdGroupNegativeKeywords Deletes negative keywords from an ad group
//
// https://developer.apple.com/documentation/apple_search_ads/delete_ad_group_negative_keywords
func (s *KeywordService) DeleteAdGroupNegativeKeywords(ctx context.Context, campaignID int64, adGroupID int64, keywordIds []int64) (*IntegerResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/delete/bulk", campaignID, adGroupID)
	res := new(IntegerResponse)
	resp, err := s.client.post(ctx, url, keywordIds, res)

	return res, resp, err
}
