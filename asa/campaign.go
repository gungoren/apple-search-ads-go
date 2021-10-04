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

// CampaignService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/campaigns
type CampaignService service

// GetAllCampaignQuery defines query parameter for GetAllCampaigns endpoint.
type GetAllCampaignQuery struct {
	Limit  int32 `url:"limit,omitempty"`
	Offset int32 `url:"offset,omitempty"`
}

type CampaignAdChannelType string

const (
	CampaignAdChannelTypeSearch  CampaignAdChannelType = "SEARCH"
	CampaignAdChannelTypeDisplay CampaignAdChannelType = "DISPLAY"
)

type Region string

const (
	RegionUS Region = "US"
)

type CampaignDisplayStatus string

const (
	CampaignDisplayStatusRunning CampaignDisplayStatus = "RUNNING"
	CampaignDisplayStatusOnHold  CampaignDisplayStatus = "ON_HOLD"
	CampaignDisplayStatusPaused  CampaignDisplayStatus = "PAUSED"
	CampaignDisplayStatusDeleted CampaignDisplayStatus = "DELETED"
)

type PaymentModel string

const (
	PaymentModelPayG   PaymentModel = "PAYG"
	PaymentModelLoc    PaymentModel = "LOC"
	PaymentModelNotSet PaymentModel = ""
)

type CampaignServingStateReason string

const (
	CampaignServingStateReasonNoPaymentMethodOnFile       CampaignServingStateReason = "NO_PAYMENT_METHOD_ON_FILE"
	CampaignServingStateReasonMissingBoOrInvoicingFields  CampaignServingStateReason = "MISSING_BO_OR_INVOICING_FIELDS"
	CampaignServingStateReasonPausedByUser                CampaignServingStateReason = "PAUSED_BY_USER"
	CampaignServingStateReasonDeletedByUser               CampaignServingStateReason = "DELETED_BY_USER"
	CampaignServingStateReasonCampaignEndDateReached      CampaignServingStateReason = "CAMPAIGN_END_DATE_REACHED"
	CampaignServingStateReasonCampaignStartDateInFuture   CampaignServingStateReason = "CAMPAIGN_START_DATE_IN_FUTURE"
	CampaignServingStateReasonDailyCapExhausted           CampaignServingStateReason = "DAILY_CAP_EXHAUSTED"
	CampaignServingStateReasonTotalBudgetExhausted        CampaignServingStateReason = "TOTAL_BUDGET_EXHAUSTED"
	CampaignServingStateReasonCreditCardDeclined          CampaignServingStateReason = "CREDIT_CARD_DECLINED"
	CampaignServingStateReasonAppNotEligible              CampaignServingStateReason = "APP_NOT_ELIGIBLE"
	CampaignServingStateReasonAppNotEligibleSearchads     CampaignServingStateReason = "APP_NOT_ELIGIBLE_SEARCHADS"
	CampaignServingStateReasonAppNotPublishedYet          CampaignServingStateReason = "APP_NOT_PUBLISHED_YET"
	CampaignServingStateReasonBoStartDateInFuture         CampaignServingStateReason = "BO_START_DATE_IN_FUTURE"
	CampaignServingStateReasonBoEndDateReached            CampaignServingStateReason = "BO_END_DATE_REACHED"
	CampaignServingStateReasonBoExhausted                 CampaignServingStateReason = "BO_EXHAUSTED"
	CampaignServingStateReasonOrgPaymentTypeChanged       CampaignServingStateReason = "ORG_PAYMENT_TYPE_CHANGED"
	CampaignServingStateReasonOrgSuspendedPolicyViolation CampaignServingStateReason = "ORG_SUSPENDED_POLICY_VIOLATION"
	CampaignServingStateReasonOrgSuspendedFraud           CampaignServingStateReason = "ORG_SUSPENDED_FRAUD"
	CampaignServingStateReasonOrgChargeBackDisputed       CampaignServingStateReason = "ORG_CHARGE_BACK_DISPUTED"
	CampaignServingStateReasonPausedBySystem              CampaignServingStateReason = "PAUSED_BY_SYSTEM"
	CampaignServingStateReasonLocExhausted                CampaignServingStateReason = "LOC_EXHAUSTED"
	CampaignServingStateReasonTaxVerificationPending      CampaignServingStateReason = "TAX_VERIFICATION_PENDING"
	CampaignServingStateReasonSapinLawAgentUnknown        CampaignServingStateReason = "SAPIN_LAW_AGENT_UNKNOWN"
	CampaignServingStateReasonSapinLawFrenchBizUnknown    CampaignServingStateReason = "SAPIN_LAW_FRENCH_BIZ_UNKNOWN"
	CampaignServingStateReasonSapinLawFrenchBiz           CampaignServingStateReason = "SAPIN_LAW_FRENCH_BIZ"
	CampaignServingStateReasonNoEligibleCountries         CampaignServingStateReason = "NO_ELIGIBLE_COUNTRIES"
	CampaignServingStateReasonAdGroupMissing              CampaignServingStateReason = "AD_GROUP_MISSING"
)

// CampaignSupplySource is the supply source of ads to use in a campaign.
type CampaignSupplySource string

const (
	// CampaignSupplySourceAppstoreSearchResults is for a campaign supply source on APPSTORE_SEARCH_RESULTS.
	CampaignSupplySourceAppstoreSearchResults CampaignSupplySource = "APPSTORE_SEARCH_RESULTS"
	// CampaignSupplySourceNews is for a campaign supply source on NEWS.
	CampaignSupplySourceNews CampaignSupplySource = "NEWS"
	// CampaignSupplySourceStocks is for a campaign supply source on STOCKS.
	CampaignSupplySourceStocks CampaignSupplySource = "STOCKS"
)

type CampaignServingStatus string

const (
	CampaignServingStatusRunning    CampaignServingStatus = "RUNNING"
	CampaignServingStatusNotRunning CampaignServingStatus = "NOT_RUNNING"
)

type CampaignStatus string

const (
	CampaignStatusEnabled CampaignStatus = "ENABLED"
	CampaignStatusPaused  CampaignStatus = "PAUSED"
)

// CampaignCountryOrRegionServingStateReasons is the reasons why a campaign can’t run
//
// https://developer.apple.com/documentation/apple_search_ads/campaign/countryorregionservingstatereasons
type CampaignCountryOrRegionServingStateReasons map[Region]CampaignCountryOrRegionServingStateReason

type CampaignCountryOrRegionServingStateReason string

const (
	CampaignCountryOrRegionServingStateReasonAppNotEligible           CampaignCountryOrRegionServingStateReason = "APP_NOT_ELIGIBLE"
	CampaignCountryOrRegionServingStateReasonAppNotEligibleSearchAds  CampaignCountryOrRegionServingStateReason = "APP_NOT_ELIGIBLE_SEARCHADS"
	CampaignCountryOrRegionServingStateReasonAppNotPublishedYet       CampaignCountryOrRegionServingStateReason = "APP_NOT_PUBLISHED_YET"
	CampaignCountryOrRegionServingStateReasonSapinLawAgentUnknown     CampaignCountryOrRegionServingStateReason = "SAPIN_LAW_AGENT_UNKNOWN"
	CampaignCountryOrRegionServingStateReasonSapinLawFrenchBizUnknown CampaignCountryOrRegionServingStateReason = "SAPIN_LAW_FRENCH_BIZ_UNKNOWN"
	CampaignCountryOrRegionServingStateReasonSapinLawFrenchBiz        CampaignCountryOrRegionServingStateReason = "SAPIN_LAW_FRENCH_BIZ"
)

// Campaign is the response to a request to create and fetch campaigns
//
// https://developer.apple.com/documentation/apple_search_ads/campaign
type Campaign struct {
	AdamID                             int64                                      `json:"adamId,omitempty"`
	AdChannelType                      CampaignAdChannelType                      `json:"adChannelType,omitempty"`
	BillingEvent                       string                                     `json:"billingEvent,omitempty"`
	BudgetAmount                       *Money                                     `json:"budgetAmount,omitempty"`
	BudgetOrders                       []int64                                    `json:"budgetOrders,omitempty"`
	CountriesOrRegions                 []Region                                   `json:"countriesOrRegions,omitempty"`
	CountryOrRegionServingStateReasons CampaignCountryOrRegionServingStateReasons `json:"countryOrRegionServingStateReasons,omitempty"`
	DailyBudgetAmount                  *Money                                     `json:"dailyBudgetAmount,omitempty"`
	Deleted                            bool                                       `json:"deleted,omitempty"`
	DisplayStatus                      CampaignDisplayStatus                      `json:"displayStatus,omitempty"`
	EndTime                            *DateTime                                  `json:"endTime,omitempty"`
	ID                                 int64                                      `json:"id,omitempty"`
	LocInvoiceDetails                  *LOCInvoiceDetails                         `json:"locInvoiceDetails,omitempty"`
	ModificationTime                   DateTime                                   `json:"modificationTime,omitempty"`
	Name                               string                                     `json:"name,omitempty"`
	OrgID                              int64                                      `json:"orgId,omitempty"`
	PaymentModel                       PaymentModel                               `json:"paymentModel,omitempty"`
	ServingStateReasons                []CampaignServingStateReason               `json:"servingStateReasons,omitempty"`
	ServingStatus                      CampaignServingStatus                      `json:"servingStatus,omitempty"`
	StartTime                          DateTime                                   `json:"startTime,omitempty"`
	Status                             CampaignStatus                             `json:"status,omitempty"`
	SupplySources                      []CampaignSupplySource                     `json:"supplySources,omitempty"`
}

// LOCInvoiceDetails is the response to a request to fetch campaign details for a standard invoicing payment model
//
// https://developer.apple.com/documentation/apple_search_ads/locinvoicedetails
type LOCInvoiceDetails struct {
	BillingContactEmail string `json:"billingContactEmail,omitempty"`
	BuyerEmail          string `json:"buyerEmail,omitempty"`
	BuyerName           string `json:"buyerName,omitempty"`
	ClientName          string `json:"clientName,omitempty"`
	OrderNumber         string `json:"orderNumber,omitempty"`
}

// CampaignResponse is a container for the campaign response body
//
// https://developer.apple.com/documentation/apple_search_ads/campaignresponse
type CampaignResponse struct {
	Campaign   *Campaign          `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// CampaignListResponse is the response details of campaign requests
//
// https://developer.apple.com/documentation/apple_search_ads/campaignlistresponse
type CampaignListResponse struct {
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
	Campaigns  []Campaign         `json:"data,omitempty"`
}

// ErrorResponseBody is a container for the error response body
//
// https://developer.apple.com/documentation/apple_search_ads/errorresponsebody
type ErrorResponseBody struct {
	Errors []ErrorResponseItem `json:"errors,omitempty"`
}

// APIErrorResponse A container for the error response body
//
// https://developer.apple.com/documentation/apple_search_ads/apierrorresponse
type APIErrorResponse struct {
	Error ErrorResponseBody `json:"error,omitempty"`
}

// ErrorResponseItemMessageCode is a system-assigned error code.
type ErrorResponseItemMessageCode string

const (
	ErrorResponseItemMessageCodeUnauthorized      ErrorResponseItemMessageCode = "UNAUTHORIZED"
	ErrorResponseItemMessageCodeInvalidDateFormat ErrorResponseItemMessageCode = "INVALID_DATE_FORMAT"
)

// ErrorResponseItem is the error response details in the response body
//
// https://developer.apple.com/documentation/apple_search_ads/errorresponseitem
type ErrorResponseItem struct {
	Field       string                       `json:"field"`
	Message     string                       `json:"message"`
	MessageCode ErrorResponseItemMessageCode `json:"messageCode"`
}

// GetAllCampaigns Fetches all of an organization’s assigned campaigns
//
// https://developer.apple.com/documentation/apple_search_ads/get_all_campaigns
func (s *CampaignService) GetAllCampaigns(ctx context.Context, params *GetAllCampaignQuery) (*CampaignListResponse, *Response, error) {
	res := new(CampaignListResponse)
	resp, err := s.client.get(ctx, "campaigns", &params, res)

	return res, resp, err
}

// GetCampaign Fetches a specific campaign by campaign identifier
//
// https://developer.apple.com/documentation/apple_search_ads/get_a_campaign
func (s *CampaignService) GetCampaign(ctx context.Context, campaignID int64) (*CampaignResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d", campaignID)
	res := new(CampaignResponse)
	resp, err := s.client.get(ctx, url, nil, res)

	return res, resp, err
}

// FindCampaigns Fetches campaigns with selector operators
//
// https://developer.apple.com/documentation/apple_search_ads/find_campaigns
func (s *CampaignService) FindCampaigns(ctx context.Context, selector *Selector) (*CampaignListResponse, *Response, error) {
	url := "campaigns/find"
	res := new(CampaignListResponse)
	resp, err := s.client.post(ctx, url, selector, res)

	return res, resp, err
}

// DeleteCampaign Deletes a specific campaign by campaign identifier
//
// https://developer.apple.com/documentation/apple_search_ads/delete_a_campaign
func (s *CampaignService) DeleteCampaign(ctx context.Context, campaignID int64) (*Response, error) {
	url := fmt.Sprintf("campaigns/%d", campaignID)
	resp, err := s.client.delete(ctx, url, nil)

	return resp, err
}

// CreateCampaign Creates a campaign to promote an app
//
// https://developer.apple.com/documentation/apple_search_ads/create_a_campaign
func (s *CampaignService) CreateCampaign(ctx context.Context, campaign *Campaign) (*CampaignResponse, *Response, error) {
	url := "campaigns"
	res := new(CampaignResponse)
	resp, err := s.client.post(ctx, url, campaign, res)

	return res, resp, err
}

// CampaignUpdate is the list of campaign fields that are updatable
//
// https://developer.apple.com/documentation/apple_search_ads/campaignupdate
type CampaignUpdate struct {
	BudgetAmount       *Money            `json:"budgetAmount,omitempty"`
	BudgetOrders       int64             `json:"budgetOrders,omitempty"`
	CountriesOrRegions []string          `json:"countriesOrRegions,omitempty"`
	DailyBudgetAmount  *Money            `json:"dailyBudgetAmount,omitempty"`
	LOCInvoiceDetails  LOCInvoiceDetails `json:"locInvoiceDetails,omitempty"`
	Name               string            `json:"name,omitempty"`
	Status             *CampaignStatus   `json:"status,omitempty"`
}

// UpdateCampaignRequest is the payload properties to clear Geo Targeting from a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/updatecampaignrequest
type UpdateCampaignRequest struct {
	Campaign                                 *CampaignUpdate `json:"campaign"`
	ClearGeoTargetingOnCountryOrRegionChange bool            `json:"clearGeoTargetingOnCountryOrRegionChange"`
}

// UpdateCampaign Updates a campaign with a campaign identifier
//
// https://developer.apple.com/documentation/apple_search_ads/update_a_campaign
func (s *CampaignService) UpdateCampaign(ctx context.Context, campaignID int64, req *UpdateCampaignRequest) (*CampaignResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d", campaignID)
	res := new(CampaignResponse)
	resp, err := s.client.put(ctx, url, req, res)

	return res, resp, err
}
