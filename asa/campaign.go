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

type CampaignService service

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

type CampaignSupplySource string

const (
	CampaignSupplySourceAppstoreSearchResults CampaignSupplySource = "APPSTORE_SEARCH_RESULTS"
	CampaignSupplySourceNews                  CampaignSupplySource = "NEWS"
	CampaignSupplySourceStocks                CampaignSupplySource = "STOCKS"
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

type Campaign struct {
	AdamId                             int64                                      `json:"adamId"`
	AdChannelType                      CampaignAdChannelType                      `json:"adChannelType"`
	BillingEvent                       string                                     `json:"billingEvent"`
	BudgetAmount                       Money                                      `json:"budgetAmount"`
	BudgetOrders                       []int64                                    `json:"budgetOrders"`
	CountriesOrRegions                 []Region                                   `json:"countriesOrRegions"`
	CountryOrRegionServingStateReasons CampaignCountryOrRegionServingStateReasons `json:"countryOrRegionServingStateReasons"`
	DailyBudgetAmount                  *Money                                     `json:"dailyBudgetAmount"`
	Deleted                            bool                                       `json:"deleted"`
	DisplayStatus                      CampaignDisplayStatus                      `json:"displayStatus"`
	EndTime                            *DateTime                                  `json:"endTime"`
	Id                                 int64                                      `json:"id"`
	LocInvoiceDetails                  LOCInvoiceDetails                          `json:"locInvoiceDetails"`
	ModificationTime                   DateTime                                   `json:"modificationTime"`
	Name                               string                                     `json:"name"`
	OrgId                              int64                                      `json:"orgId"`
	PaymentModel                       PaymentModel                               `json:"paymentModel"`
	ServingStateReasons                []CampaignServingStateReason               `json:"servingStateReasons"`
	ServingStatus                      CampaignServingStatus                      `json:"servingStatus"`
	StartTime                          DateTime                                   `json:"startTime"`
	Status                             CampaignStatus                             `json:"status"`
	SupplySources                      []CampaignSupplySource                     `json:"supplySources"`
}

type LOCInvoiceDetails struct {
	BillingContactEmail string `json:"billingContactEmail"`
	BuyerEmail          string `json:"buyerEmail"`
	BuyerName           string `json:"buyerName"`
	ClientName          string `json:"clientName"`
	OrderNumber         string `json:"orderNumber"`
}

type CampaignResponse struct {
	Error      ErrorResponseBody `json:"error"`
	Pagination PageDetail        `json:"pagination"`
	Campaign   Campaign          `json:"data"`
}

func (s *CampaignService) GetAllCampaigns(ctx context.Context, params *GetAllCampaignQuery) (*CampaignListResponse, *Response, error) {
	res := new(CampaignListResponse)
	resp, err := s.client.get(ctx, "campaigns", &params, res)
	return res, resp, err
}

func (s *CampaignService) GetCampaign(ctx context.Context, campaignId int64) (*CampaignResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d", campaignId)
	res := new(CampaignResponse)
	resp, err := s.client.get(ctx, url, nil, res)
	return res, resp, err
}

func (s *CampaignService) FindCampaigns(ctx context.Context, selector *Selector) (*CampaignListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/find")
	res := new(CampaignListResponse)
	resp, err := s.client.post(ctx, url, selector, res)
	return res, resp, err
}

func (s *CampaignService) DeleteCampaign(ctx context.Context, campaignId int64) (*Response, error) {
	url := fmt.Sprintf("campaigns/%d", campaignId)
	resp, err := s.client.delete(ctx, url, nil)
	return resp, err
}

func (s *CampaignService) CreateCampaign(ctx context.Context, campaign *Campaign) (*CampaignResponse, *Response, error) {
	url := fmt.Sprintf("campaigns")
	res := new(CampaignResponse)
	resp, err := s.client.post(ctx, url, campaign, res)
	return res, resp, err
}

func (s *CampaignService) UpdateCampaign(ctx context.Context, campaignId int64, req *UpdateCampaignRequest) (*CampaignResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d", campaignId)
	res := new(CampaignResponse)
	resp, err := s.client.put(ctx, url, req, res)
	return res, resp, err
}
