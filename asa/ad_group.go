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

// AdGroupService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/ad_groups
type AdGroupService service

type AdGroupResponse struct {
	Error      *GeneralErrorResponse `json:"error,omitempty"`
	PageDetail *PageDetail           `json:"pageDetail,omitempty"`
	AdGroup    *AdGroup              `json:"data,omitempty"`
}

// AdGroupDisplayStatus defines model for AdGroupDisplayStatus.
//
// https://developer.apple.com/documentation/apple_search_ads/adgroup
type AdGroupDisplayStatus string

const (
	// AdGroupDisplayStatusDelete is for an ad group display status on Deleted.
	AdGroupDisplayStatusDelete AdGroupDisplayStatus = "DELETED"
	// AdGroupDisplayStatusOnHold is for an ad group display status on On Hold.
	AdGroupDisplayStatusOnHold AdGroupDisplayStatus = "ON_HOLD"
	// AdGroupDisplayStatusPaused is for an ad group display status on Paused.
	AdGroupDisplayStatusPaused AdGroupDisplayStatus = "PAUSED"
	// AdGroupDisplayStatusRunning is for an ad group display status on Running.
	AdGroupDisplayStatusRunning AdGroupDisplayStatus = "RUNNING"
)

// AdGroupPricingModel defines model for AdGroupPricingModel.
//
// https://developer.apple.com/documentation/apple_search_ads/adgroup
type AdGroupPricingModel string

const (
	//AdGroupPricingModelCPC is for an ad group pricing model CPC
	AdGroupPricingModelCPC AdGroupPricingModel = "CPC"
	//AdGroupPricingModelCPM is for an ad group pricing model CPM
	AdGroupPricingModelCPM AdGroupPricingModel = "CPM"
)

type ServingStateReason string

const (
	ServingStateReasonAdGroupPausedByUser         ServingStateReason = "AD_GROUP_PAUSED_BY_USER"
	ServingStateReasonAdGroupEndDateReached       ServingStateReason = "ADGROUP_END_DATE_REACHED"
	ServingStateReasonAppNotSupport               ServingStateReason = "APP_NOT_SUPPORT"
	ServingStateReasonAudienceBelowThreshold      ServingStateReason = "AUDIENCE_BELOW_THRESHOLD"
	ServingStateReasonCampaignNotRunning          ServingStateReason = "CAMPAIGN_NOT_RUNNING"
	ServingStateReasonDeletedByUser               ServingStateReason = "DELETED_BY_USER"
	ServingStateReasonPendingAudienceVerification ServingStateReason = "PENDING_AUDIENCE_VERIFICATION"
	ServingStateReasonStartDateInTheFuture        ServingStateReason = "START_DATE_IN_THE_FUTURE"
)

type AdGroupServingStatus string

const (
	AdGroupServingStatusNotRunning AdGroupServingStatus = "NOT_RUNNING"
	AdGroupServingStatusRunning    AdGroupServingStatus = "RUNNING"
)

type AdGroupStatus string

const (
	AdGroupStatusEnabled AdGroupStatus = "ENABLED"
	AdGroupStatusPaused  AdGroupStatus = "PAUSED"
)

type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type AdGroup struct {
	AutomatedKeywordsOptIn bool                 `json:"automatedKeywordsOptIn,omitempty"`
	CampaignID             int64                `json:"campaignID,omitempty"`
	CpaGoal                *Money               `json:"cpaGoal,omitempty"`
	DefaultBidAmount       *Money               `json:"defaultBidAmount"`
	Deleted                bool                 `json:"deleted"`
	DisplayStatus          AdGroupDisplayStatus `json:"displayStatus"`
	EndTime                DateTime             `json:"endTime,omitempty"`
	Id                     int64                `json:"id,omitempty"`
	ModificationTime       DateTime             `json:"modificationTime,omitempty"`
	Name                   string               `json:"name,omitempty"`
	OrgId                  int64                `json:"orgId,omitempty"`
	PricingModel           AdGroupPricingModel  `json:"pricingModel"`
	ServingStateReasons    []ServingStateReason `json:"servingStateReasons,omitempty"`
	ServingStatus          AdGroupServingStatus `json:"servingStatus"`
	StartTime              DateTime             `json:"startTime,omitempty"`
	Status                 AdGroupStatus        `json:"status,omitempty"`
	TargetDimensions       *TargetDimensions    `json:"targetDimensions,omitempty"`
}

type TargetDimensions struct {
	AdminArea      *AdminAreaCriteria     `json:"adminArea,omitempty"`
	Age            *AgeCriteria           `json:"age,omitempty"`
	AppDownloaders *AppDownloaderCriteria `json:"appDownloaders"`
	Country        *CountryCriteria       `json:"country,omitempty"`
	DayPart        *DayPartCriteria       `json:"daypart,omitempty"`
	DeviceClass    *DeviceClassCriteria   `json:"deviceClass,omitempty"`
	Gender         *GenderCriteria        `json:"gender,omitempty"`
	Locality       *LocalityCriteria      `json:"locality,omitempty"`
}

type AdminAreaCriteria struct {
	Included []string `json:"included,omitempty"`
}

type AgeCriteria struct {
	Included []*AgeRange `json:"included,omitempty"`
}

type AgeRange struct {
	MaxAge int32 `json:"maxAge,omitempty"`
	MinAge int32 `json:"minAge,omitempty"`
}

type AppDownloaderCriteria struct {
	Included []string `json:"included,omitempty"`
	Excluded []string `json:"excluded,omitempty"`
}

type CountryCriteria struct {
	Included []string `json:"included,omitempty"`
}

type DayPartCriteria struct {
	UserTime *DaypartDetail `json:"userTime,omitempty"`
}

type DaypartDetail struct {
	Included []int32 `json:"included,omitempty"`
}

type AdGroupDeviceClass string

const (
	AdGroupDeviceClassIpad   AdGroupDeviceClass = "IPAD"
	AdGroupDeviceClassIphone AdGroupDeviceClass = "IPHONE"
)

type DeviceClassCriteria struct {
	Included []AdGroupDeviceClass `json:"included,omitempty"`
}

type AdGroupGender string

const (
	AdGroupGenderFemale AdGroupGender = "F"
	AdGroupGenderMale   AdGroupGender = "M"
)

type GenderCriteria struct {
	Included []AdGroupGender `json:"included,omitempty"`
}

type LocalityCriteria struct {
	Included []string `json:"included,omitempty"`
}

type AdGroupListResponse struct {
	AdGroup    []AdGroup          `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	PageDetail *PageDetail        `json:"pageDetail,omitempty"`
}

func (s *AdGroupService) CreateAdGroup(ctx context.Context, campaignId int64, adGroup *AdGroup) (*AdGroupResponse, *Response, error) {
	url := fmt.Sprintf("/campaigns/%d/adgroups", campaignId)
	res := new(AdGroupResponse)
	resp, err := s.client.post(ctx, url, adGroup, res)
	return res, resp, err
}

type GetAllAdGroupsQuery struct {
	Limit  int32 `url:"limit,omitempty"`
	Offset int32 `url:"offset,omitempty"`
}

type AdGroupUpdateRequest struct {
	AutomatedKeywordsOptIn bool              `json:"automatedKeywordsOptIn,omitempty"`
	CpaGoal                *Money            `json:"cpaGoal,omitempty"`
	DefaultBidAmount       *Money            `json:"defaultBidAmount,omitempty"`
	EndTime                DateTime          `json:"endTime,omitempty"`
	Name                   string            `json:"name,omitempty"`
	StartTime              DateTime          `json:"startTime,omitempty"`
	Status                 AdGroupStatus     `json:"status,omitempty"`
	TargetingDimensions    *TargetDimensions `json:"targetingDimensions"`
}

type ConditionOperator string

const (
	ConditionOperatorEquals      ConditionOperator = "EQUALS"
	ConditionOperatorGreaterThan ConditionOperator = "GREATER_THAN"
	ConditionOperatorLessThan    ConditionOperator = "LESS_THAN"
	ConditionOperatorIn          ConditionOperator = "IN"
	ConditionOperatorLike        ConditionOperator = "LIKE"
	ConditionOperatorStartsWith  ConditionOperator = "STARTSWITH"
	ConditionOperatorContains    ConditionOperator = "CONTAINS"
	ConditionOperatorEndsWith    ConditionOperator = "ENDSWITH"
	ConditionOperatorNotEqual    ConditionOperator = "NOT_EQUALS"
	ConditionOperatorIs          ConditionOperator = "IS"
	ConditionOperatorContainsAny ConditionOperator = "CONTAINS_ANY"
	ConditionOperatorContainsAll ConditionOperator = "CONTAINS_ALL"
)

type Selector struct {
	Conditions []*Condition `json:"conditions,omitempty"`
	Fields     []string     `json:"fields,omitempty"`
	OrderBy    []*Sorting   `json:"orderBy,omitempty"`
	Pagination *Pagination  `json:"pagination,omitempty"`
}

type Condition struct {
	Field    string            `json:"field"`
	Operator ConditionOperator `json:"operator"`
	Values   []string          `json:"values"`
}

type SortOrder string

const (
	SortingOrderAscending  SortOrder = "ASCENDING"
	SortingOrderDescending SortOrder = "DESCENDING"
)

type Sorting struct {
	Field     string    `json:"field"`
	SortOrder SortOrder `json:"sortOrder"`
}

type Pagination struct {
	Limit  uint32 `json:"limit"`
	Offset uint32 `json:"offset"`
}

func (s *AdGroupService) FindAdGroups(ctx context.Context, campaignId int64, selector *Selector) (*AdGroupListResponse, *Response, error) {
	url := fmt.Sprintf("/campaigns/%d/adgroups/find", campaignId)
	res := new(AdGroupListResponse)
	resp, err := s.client.post(ctx, url, selector, res)
	return res, resp, err
}

func (s *AdGroupService) GetAdGroup(ctx context.Context, campaignId int64, adGroupId int64) (*AdGroupResponse, *Response, error) {
	url := fmt.Sprintf("/campaigns/%d/adgroups/%d", campaignId, adGroupId)
	res := new(AdGroupResponse)
	resp, err := s.client.get(ctx, url, nil, res)
	return res, resp, err
}

func (s *AdGroupService) GetAllAdGroups(ctx context.Context, campaignId int64, params *GetAllAdGroupsQuery) (*AdGroupListResponse, *Response, error) {
	url := fmt.Sprintf("/campaigns/%d/adgroups", campaignId)
	res := new(AdGroupListResponse)
	resp, err := s.client.get(ctx, url, &params, res)
	return res, resp, err
}

func (s *AdGroupService) UpdateAdGroup(ctx context.Context, campaignId int64, adGroupId int64, req *AdGroupUpdateRequest) (*AdGroupResponse, *Response, error) {
	url := fmt.Sprintf("/campaigns/%d/adgroups/%d", campaignId, adGroupId)
	res := new(AdGroupResponse)
	resp, err := s.client.put(ctx, url, req, res)
	return res, resp, err
}

func (s *AdGroupService) DeleteAdGroup(ctx context.Context, campaignId int64, adGroupId int64) (*Response, error) {
	url := fmt.Sprintf("/campaigns/%d/adgroups/%d", campaignId, adGroupId)
	resp, err := s.client.delete(ctx, url, nil)
	return resp, err
}
