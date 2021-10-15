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

// CampaignAdChannelType is the channel type of ad in a campaign.
type CampaignAdChannelType string

const (
	// CampaignAdChannelTypeSearch When supplySources is APPSTORE_SEARCH_RESULTS, the adChannelType must be SEARCH.
	CampaignAdChannelTypeSearch CampaignAdChannelType = "SEARCH"
	// CampaignAdChannelTypeDisplay When supplySources is APPSTORE_SEARCH_TAB, the adChannelType must be DISPLAY.
	CampaignAdChannelTypeDisplay CampaignAdChannelType = "DISPLAY"
)

// CampaignDisplayStatus is the status of the campaign.
type CampaignDisplayStatus string

const (
	// CampaignDisplayStatusRunning is for a campaign status on RUNNING.
	CampaignDisplayStatusRunning CampaignDisplayStatus = "RUNNING"
	// CampaignDisplayStatusOnHold is for a campaign status on ON_HOLD.
	CampaignDisplayStatusOnHold CampaignDisplayStatus = "ON_HOLD"
	// CampaignDisplayStatusPaused is for a campaign status on PAUSED.
	CampaignDisplayStatusPaused CampaignDisplayStatus = "PAUSED"
	// CampaignDisplayStatusDeleted is for a campaign status on DELETED.
	CampaignDisplayStatusDeleted CampaignDisplayStatus = "DELETED"
)

// PaymentModel is the payment model that you set through the Search Ads UI.
type PaymentModel string

const (
	// PaymentModelPayG is a pay-as-you-go payment mode.
	PaymentModelPayG PaymentModel = "PAYG"
	// PaymentModelLoc is a line-of-credit payment model.
	PaymentModelLoc PaymentModel = "LOC"
	// PaymentModelNotSet is represent there is no set payment method.
	PaymentModelNotSet PaymentModel = ""
)

// CampaignServingStateReason is a reason that displays when a campaign can’t run.
type CampaignServingStateReason string

const (
	// CampaignServingStateReasonNoPaymentMethodOnFile is for a campaign serving state reason for NO_PAYMENT_METHOD_ON_FILE.
	CampaignServingStateReasonNoPaymentMethodOnFile CampaignServingStateReason = "NO_PAYMENT_METHOD_ON_FILE"
	// CampaignServingStateReasonMissingBoOrInvoicingFields is for a campaign serving state reason for MISSING_BO_OR_INVOICING_FIELDS.
	CampaignServingStateReasonMissingBoOrInvoicingFields CampaignServingStateReason = "MISSING_BO_OR_INVOICING_FIELDS"
	// CampaignServingStateReasonPausedByUser is for a campaign serving state reason for PAUSED_BY_USER.
	CampaignServingStateReasonPausedByUser CampaignServingStateReason = "PAUSED_BY_USER"
	// CampaignServingStateReasonDeletedByUser is for a campaign serving state reason for DELETED_BY_USER.
	CampaignServingStateReasonDeletedByUser CampaignServingStateReason = "DELETED_BY_USER"
	// CampaignServingStateReasonCampaignEndDateReached is for a campaign serving state reason for CAMPAIGN_END_DATE_REACHED.
	CampaignServingStateReasonCampaignEndDateReached CampaignServingStateReason = "CAMPAIGN_END_DATE_REACHED"
	// CampaignServingStateReasonCampaignStartDateInFuture is for a campaign serving state reason for CAMPAIGN_START_DATE_IN_FUTURE.
	CampaignServingStateReasonCampaignStartDateInFuture CampaignServingStateReason = "CAMPAIGN_START_DATE_IN_FUTURE"
	// CampaignServingStateReasonDailyCapExhausted is for a campaign serving state reason for DAILY_CAP_EXHAUSTED.
	CampaignServingStateReasonDailyCapExhausted CampaignServingStateReason = "DAILY_CAP_EXHAUSTED"
	// CampaignServingStateReasonTotalBudgetExhausted is for a campaign serving state reason for TOTAL_BUDGET_EXHAUSTED.
	CampaignServingStateReasonTotalBudgetExhausted CampaignServingStateReason = "TOTAL_BUDGET_EXHAUSTED"
	// CampaignServingStateReasonCreditCardDeclined is for a campaign serving state reason for CREDIT_CARD_DECLINED.
	CampaignServingStateReasonCreditCardDeclined CampaignServingStateReason = "CREDIT_CARD_DECLINED"
	// CampaignServingStateReasonAppNotEligible is for a campaign serving state reason for APP_NOT_ELIGIBLE.
	CampaignServingStateReasonAppNotEligible CampaignServingStateReason = "APP_NOT_ELIGIBLE"
	// CampaignServingStateReasonAppNotEligibleSearchads is for a campaign serving state reason for APP_NOT_ELIGIBLE_SEARCHADS.
	CampaignServingStateReasonAppNotEligibleSearchads CampaignServingStateReason = "APP_NOT_ELIGIBLE_SEARCHADS"
	// CampaignServingStateReasonAppNotPublishedYet is for a campaign serving state reason for APP_NOT_PUBLISHED_YET.
	CampaignServingStateReasonAppNotPublishedYet CampaignServingStateReason = "APP_NOT_PUBLISHED_YET"
	// CampaignServingStateReasonBoStartDateInFuture is for a campaign serving state reason for BO_START_DATE_IN_FUTURE.
	CampaignServingStateReasonBoStartDateInFuture CampaignServingStateReason = "BO_START_DATE_IN_FUTURE"
	// CampaignServingStateReasonBoEndDateReached is for a campaign serving state reason for BO_END_DATE_REACHED.
	CampaignServingStateReasonBoEndDateReached CampaignServingStateReason = "BO_END_DATE_REACHED"
	// CampaignServingStateReasonBoExhausted is for a campaign serving state reason for BO_EXHAUSTED.
	CampaignServingStateReasonBoExhausted CampaignServingStateReason = "BO_EXHAUSTED"
	// CampaignServingStateReasonOrgPaymentTypeChanged is for a campaign serving state reason for ORG_PAYMENT_TYPE_CHANGED.
	CampaignServingStateReasonOrgPaymentTypeChanged CampaignServingStateReason = "ORG_PAYMENT_TYPE_CHANGED"
	// CampaignServingStateReasonOrgSuspendedPolicyViolation is for a campaign serving state reason for ORG_SUSPENDED_POLICY_VIOLATION.
	CampaignServingStateReasonOrgSuspendedPolicyViolation CampaignServingStateReason = "ORG_SUSPENDED_POLICY_VIOLATION"
	// CampaignServingStateReasonOrgSuspendedFraud is for a campaign serving state reason for ORG_SUSPENDED_FRAUD.
	CampaignServingStateReasonOrgSuspendedFraud CampaignServingStateReason = "ORG_SUSPENDED_FRAUD"
	// CampaignServingStateReasonOrgChargeBackDisputed is for a campaign serving state reason for ORG_CHARGE_BACK_DISPUTED.
	CampaignServingStateReasonOrgChargeBackDisputed CampaignServingStateReason = "ORG_CHARGE_BACK_DISPUTED"
	// CampaignServingStateReasonPausedBySystem is for a campaign serving state reason for PAUSED_BY_SYSTEM.
	CampaignServingStateReasonPausedBySystem CampaignServingStateReason = "PAUSED_BY_SYSTEM"
	// CampaignServingStateReasonLocExhausted is for a campaign serving state reason for LOC_EXHAUSTED.
	CampaignServingStateReasonLocExhausted CampaignServingStateReason = "LOC_EXHAUSTED"
	// CampaignServingStateReasonTaxVerificationPending is for a campaign serving state reason for TAX_VERIFICATION_PENDING.
	CampaignServingStateReasonTaxVerificationPending CampaignServingStateReason = "TAX_VERIFICATION_PENDING"
	// CampaignServingStateReasonSapinLawAgentUnknown is for a campaign serving state reason for SAPIN_LAW_AGENT_UNKNOWN.
	CampaignServingStateReasonSapinLawAgentUnknown CampaignServingStateReason = "SAPIN_LAW_AGENT_UNKNOWN"
	// CampaignServingStateReasonSapinLawFrenchBizUnknown is for a campaign serving state reason for SAPIN_LAW_FRENCH_BIZ_UNKNOWN.
	CampaignServingStateReasonSapinLawFrenchBizUnknown CampaignServingStateReason = "SAPIN_LAW_FRENCH_BIZ_UNKNOWN"
	// CampaignServingStateReasonSapinLawFrenchBiz is for a campaign serving state reason for SAPIN_LAW_FRENCH_BIZ.
	CampaignServingStateReasonSapinLawFrenchBiz CampaignServingStateReason = "SAPIN_LAW_FRENCH_BIZ"
	// CampaignServingStateReasonNoEligibleCountries is for a campaign serving state reason for NO_ELIGIBLE_COUNTRIES.
	CampaignServingStateReasonNoEligibleCountries CampaignServingStateReason = "NO_ELIGIBLE_COUNTRIES"
	// CampaignServingStateReasonAdGroupMissing is for a campaign serving state reason for AD_GROUP_MISSING.
	CampaignServingStateReasonAdGroupMissing CampaignServingStateReason = "AD_GROUP_MISSING"
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

// CampaignServingStatus is the status of the campaign.
type CampaignServingStatus string

const (
	// CampaignServingStatusRunning is for a campaign serving status source on RUNNING.
	CampaignServingStatusRunning CampaignServingStatus = "RUNNING"
	// CampaignServingStatusNotRunning is for a campaign supply source on NOT_RUNNING.
	CampaignServingStatusNotRunning CampaignServingStatus = "NOT_RUNNING"
)

// CampaignStatus is the user-controlled status to enable or pause the campaign.
type CampaignStatus string

const (
	// CampaignStatusEnabled is for a campaign status on ENABLED.
	CampaignStatusEnabled CampaignStatus = "ENABLED"
	// CampaignStatusPaused is for a campaign status source on PAUSED.
	CampaignStatusPaused CampaignStatus = "PAUSED"
)

// CampaignCountryOrRegionServingStateReasons is the reasons why a campaign can’t run
//
// https://developer.apple.com/documentation/apple_search_ads/campaign/countryorregionservingstatereasons
type CampaignCountryOrRegionServingStateReasons map[string]CampaignCountryOrRegionServingStateReason

// CampaignCountryOrRegionServingStateReason is a reason that returns when a campaign can’t run for a specified country or region.
type CampaignCountryOrRegionServingStateReason string

const (
	// CampaignCountryOrRegionServingStateReasonAppNotEligible is for a campaign country or region serving state reason on APP_NOT_ELIGIBLE.
	CampaignCountryOrRegionServingStateReasonAppNotEligible CampaignCountryOrRegionServingStateReason = "APP_NOT_ELIGIBLE"
	// CampaignCountryOrRegionServingStateReasonAppNotEligibleSearchAds is for a campaign country or region serving state reason on APP_NOT_ELIGIBLE_SEARCHADS.
	CampaignCountryOrRegionServingStateReasonAppNotEligibleSearchAds CampaignCountryOrRegionServingStateReason = "APP_NOT_ELIGIBLE_SEARCHADS"
	// CampaignCountryOrRegionServingStateReasonAppNotPublishedYet is for a campaign country or region serving state reason on APP_NOT_PUBLISHED_YET.
	CampaignCountryOrRegionServingStateReasonAppNotPublishedYet CampaignCountryOrRegionServingStateReason = "APP_NOT_PUBLISHED_YET"
	// CampaignCountryOrRegionServingStateReasonSapinLawAgentUnknown is for a campaign country or region serving state reason on SAPIN_LAW_AGENT_UNKNOWN.
	CampaignCountryOrRegionServingStateReasonSapinLawAgentUnknown CampaignCountryOrRegionServingStateReason = "SAPIN_LAW_AGENT_UNKNOWN"
	// CampaignCountryOrRegionServingStateReasonSapinLawFrenchBizUnknown is for a campaign country or region serving state reason on SAPIN_LAW_FRENCH_BIZ_UNKNOWN.
	CampaignCountryOrRegionServingStateReasonSapinLawFrenchBizUnknown CampaignCountryOrRegionServingStateReason = "SAPIN_LAW_FRENCH_BIZ_UNKNOWN"
	// CampaignCountryOrRegionServingStateReasonSapinLawFrenchBiz is for a campaign country or region serving state reason on SAPIN_LAW_FRENCH_BIZ.
	CampaignCountryOrRegionServingStateReasonSapinLawFrenchBiz CampaignCountryOrRegionServingStateReason = "SAPIN_LAW_FRENCH_BIZ"
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
	CountriesOrRegions                 []string                                   `json:"countriesOrRegions,omitempty"`
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
	Campaigns  []*Campaign        `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
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
	// ErrorResponseItemMessageCodeUnauthorized is for an error response item message code on UNAUTHORIZED.
	ErrorResponseItemMessageCodeUnauthorized ErrorResponseItemMessageCode = "UNAUTHORIZED"
	// ErrorResponseItemMessageCodeInvalidDateFormat is for an error response item message code on INVALID_DATE_FORMAT.
	ErrorResponseItemMessageCodeInvalidDateFormat ErrorResponseItemMessageCode = "INVALID_DATE_FORMAT"
)

// ErrorResponseItem is the error response details in the response body
//
// https://developer.apple.com/documentation/apple_search_ads/errorresponseitem
type ErrorResponseItem struct {
	Field       string                       `json:"field,omitempty"`
	Message     string                       `json:"message,omitempty"`
	MessageCode ErrorResponseItemMessageCode `json:"messageCode,omitempty"`
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
