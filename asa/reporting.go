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

// ReportingService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/reports
type ReportingService service

type ReportingRequestGranularity string

const (
	ReportingRequestGranularityTypeHourly  ReportingRequestGranularity = "HOURLY"
	ReportingRequestGranularityTypeDaily   ReportingRequestGranularity = "DAILY"
	ReportingRequestGranularityTypeWeekly  ReportingRequestGranularity = "WEEKLY"
	ReportingRequestGranularityTypeMonthly ReportingRequestGranularity = "MONTHLY"
)

type ReportingRequestTimeZone string

const (
	ReportingRequestTimeZoneUTC  ReportingRequestTimeZone = "UTC"
	ReportingRequestTimeZoneORTZ ReportingRequestTimeZone = "ORTZ"
)

type ReportingRequestGroupBy string

const (
	ReportingRequestGroupByTypeAdminArea       ReportingRequestGroupBy = "adminArea"
	ReportingRequestGroupByTypeAgeRange        ReportingRequestGroupBy = "ageRange"
	ReportingRequestGroupByTypeCountryCode     ReportingRequestGroupBy = "countryCode"
	ReportingRequestGroupByTypeCountryOrRegion ReportingRequestGroupBy = "countryOrRegion"
	ReportingRequestGroupByTypeDeviceClass     ReportingRequestGroupBy = "deviceClass"
	ReportingRequestGroupByTypeGender          ReportingRequestGroupBy = "gender"
	ReportingRequestGroupByTypeLocality        ReportingRequestGroupBy = "locality"
)

// ReportingRequest is the report request body
//
// https://developer.apple.com/documentation/apple_search_ads/reportingrequest
type ReportingRequest struct {
	StartTime                  DateTime                     `json:"startTime"`
	EndTime                    DateTime                     `json:"endTime"`
	Granularity                *ReportingRequestGranularity `json:"granularity,omitempty"`
	TimeZone                   *ReportingRequestTimeZone    `json:"timeZone,omitempty"`
	GroupBy                    []ReportingRequestGroupBy    `json:"groupBy,omitempty"`
	ReturnGrandTotals          bool                         `json:"returnGrandTotals"`
	ReturnRecordsWithNoMetrics bool                         `json:"returnRecordsWithNoMetrics"`
	ReturnRowTotals            bool                         `json:"returnRowTotals"`
	Selector                   *Selector                    `json:"selector,omitempty"`
}

type ReportingResponseBody struct {
	ReportingCampaign *ReportingResponse `json:"data,omitempty"`
	Pagination        *PageDetail        `json:"pagination,omitempty"`
}

// ReportingResponse is a container for report metrics
//
// https://developer.apple.com/documentation/apple_search_ads/reportingresponse
type ReportingResponse struct {
	ReportingDataResponse *ReportingDataResponse `json:"reportingDataResponse,omitempty"`
}

// ReportingDataResponse is the total metrics for a report
//
// https://developer.apple.com/documentation/apple_search_ads/reportingdataresponse
type ReportingDataResponse struct {
	Rows        []Row           `json:"row,omitempty"`
	GrandTotals *GrandTotalsRow `json:"grandTotals,omitempty"`
}

// Row is the report metrics organized by time granularity.
//
// https://developer.apple.com/documentation/apple_search_ads/row
type Row struct {
	Other       bool                `json:"other,omitempty"`
	Granularity []*ExtendedSpendRow `json:"granularity,omitempty"`
	Total       *SpendRow           `json:"total,omitempty"`
	Metadata    *MetaDataObject     `json:"metadata,omitempty"`
	Insights    *InsightsObject     `json:"insights,omitempty"`
}

type ReportingKeywordMatchType string

const (
	ReportingKeywordMatchTypeAuto  ReportingKeywordMatchType = "AUTO"
	ReportingKeywordMatchTypeExact ReportingKeywordMatchType = "EXACT"
	ReportingKeywordMatchTypeBroad ReportingKeywordMatchType = "BROAD"
)

type SearchTermSource string

const (
	SearchTermSourceAuto     SearchTermSource = "AUTO"
	SearchTermSourceTargeted SearchTermSource = "TARGETED"
)

// CampaignAppDetail is the app data to fetch from campaign-level reports
//
// https://developer.apple.com/documentation/apple_search_ads/campaignappdetail
type CampaignAppDetail struct {
	AppName string `json:"appName"`
	AdamID  int64  `json:"adamId"`
}

// MetaDataObject is the report response objects
//
// https://developer.apple.com/documentation/apple_search_ads/metadataobject
type MetaDataObject struct {
	AdGroupID                          int64                                       `json:"adGroupID,omitempty"`
	AdGroupName                        string                                      `json:"adGroupName,omitempty"`
	CampaignID                         int64                                       `json:"campaignId,omitempty"`
	CampaignName                       string                                      `json:"campaignName,omitempty"`
	Deleted                            bool                                        `json:"deleted,omitempty"`
	CampaignStatus                     CampaignStatus                              `json:"campaignStatus,omitempty"`
	App                                *CampaignAppDetail                          `json:"app,omitempty"`
	ServingStatus                      CampaignServingStatus                       `json:"servingStatus,omitempty"`
	ServingStateReasons                []CampaignServingStateReason                `json:"servingStateReasons,omitempty"`
	CountriesOrRegions                 []Region                                    `json:"countriesOrRegions,omitempty"`
	ModificationTime                   DateTime                                    `json:"modificationTime,omitempty"`
	TotalBudget                        *Money                                      `json:"totalBudget,omitempty"`
	DailyBudget                        *Money                                      `json:"dailyBudget,omitempty"`
	DisplayStatus                      CampaignDisplayStatus                       `json:"displayStatus,omitempty"`
	SupplySources                      []CampaignSupplySource                      `json:"supplySources,omitempty"`
	AdChannelType                      CampaignAdChannelType                       `json:"adChannelType,omitempty"`
	OrgID                              int                                         `json:"orgId,omitempty"`
	CountryOrRegionServingStateReasons *CampaignCountryOrRegionServingStateReasons `json:"countryOrRegionServingStateReasons,omitempty"`
	BillingEvent                       string                                      `json:"billingEvent,omitempty"`
	KeywordID                          int64                                       `json:"keywordID,omitempty"`
	MatchType                          *ReportingKeywordMatchType                  `json:"matchType,omitempty"`
	CountryOrRegion                    *Region                                     `json:"countryOrRegion,omitempty"`
	SearchTermText                     []string                                    `json:"SearchTermText,omitempty"`
	SearchTermSource                   *SearchTermSource                           `json:"searchTermSource,omitempty"`
}

// GrandTotalsRow is the summary of cumulative metrics
//
// https://developer.apple.com/documentation/apple_search_ads/grandtotalsrow
type GrandTotalsRow struct {
	Other bool      `json:"other,omitempty"`
	Total *SpendRow `json:"total,omitempty"`
}

// SpendRow is the reporting response metrics
//
// https://developer.apple.com/documentation/apple_search_ads/spendrow
type SpendRow struct {
	AvgCPA         *Money  `json:"avgCPA,omitempty"`
	AvgCPT         *Money  `json:"avgCPT,omitempty"`
	AvgCPM         *Money  `json:"avgCPM,omitempty"`
	ConversionRate float64 `json:"conversionRate,omitempty"`
	Impressions    int64   `json:"impressions,omitempty"`
	Installs       int64   `json:"installs,omitempty"`
	LatOffInstalls int64   `json:"latOffInstalls,omitempty"`
	LatOnInstalls  int64   `json:"latOnInstalls,omitempty"`
	LocalSpend     *Money  `json:"localSpend,omitempty"`
	NewDownloads   int64   `json:"newDownloads,omitempty"`
	ReDownloads    int64   `json:"redownloads,omitempty"`
	Taps           int64   `json:"taps,omitempty"`
	Ttr            float64 `json:"ttr,omitempty"`
}

// ExtendedSpendRow is the descriptions of metrics with dates
//
// https://developer.apple.com/documentation/apple_search_ads/extendedspendrow
type ExtendedSpendRow struct {
	AvgCPA         *Money  `json:"avgCPA,omitempty"`
	AvgCPT         *Money  `json:"avgCPT,omitempty"`
	AvgCPM         *Money  `json:"avgCPM,omitempty"`
	ConversionRate float64 `json:"conversionRate,omitempty"`
	Impressions    int64   `json:"impressions,omitempty"`
	Installs       int64   `json:"installs,omitempty"`
	LatOffInstalls int64   `json:"latOffInstalls,omitempty"`
	LatOnInstalls  int64   `json:"latOnInstalls,omitempty"`
	LocalSpend     *Money  `json:"localSpend,omitempty"`
	NewDownloads   int64   `json:"newDownloads,omitempty"`
	ReDownloads    int64   `json:"redownloads,omitempty"`
	Taps           int64   `json:"taps,omitempty"`
	Ttr            float64 `json:"ttr,omitempty"`
	Date           Date    `json:"date,omitempty"`
}

// InsightsObject is a parent object for bid recommendations
//
// https://developer.apple.com/documentation/apple_search_ads/insightsobject
type InsightsObject struct {
	BidRecommendation *KeywordBidRecommendation `json:"bidRecommendation,omitempty"`
}

// KeywordBidRecommendation is the bid recommendation range for a keyword
//
// https://developer.apple.com/documentation/apple_search_ads/keywordbidrecommendation
type KeywordBidRecommendation struct {
	BidMax *Money `json:"bidMax,omitempty"`
	BidMin *Money `json:"bidMin,omitempty"`
}

// GetCampaignLevelReports fetches reports for campaigns
//
// https://developer.apple.com/documentation/apple_search_ads/get_campaign-level_reports
func (s *ReportingService) GetCampaignLevelReports(ctx context.Context, params *ReportingRequest) (*ReportingResponseBody, *Response, error) {
	url := "reports/campaigns"
	res := new(ReportingResponseBody)
	resp, err := s.client.post(ctx, url, &params, res)

	return res, resp, err
}

// GetAdGroupLevelReports fetches reports for ad groups within a campaig
//
// https://developer.apple.com/documentation/apple_search_ads/get_ad_group-level_reports
func (s *ReportingService) GetAdGroupLevelReports(ctx context.Context, campaignID int64, params *ReportingRequest) (*ReportingResponseBody, *Response, error) {
	url := fmt.Sprintf("reports/campaigns/%d/adgroups", campaignID)
	res := new(ReportingResponseBody)
	resp, err := s.client.post(ctx, url, &params, res)

	return res, resp, err
}

// GetKeywordLevelReports fetches reports for targeting keywords within a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/get_keyword-level_reports
func (s *ReportingService) GetKeywordLevelReports(ctx context.Context, campaignID int64, params *ReportingRequest) (*ReportingResponseBody, *Response, error) {
	url := fmt.Sprintf("reports/campaigns/%d/keywords", campaignID)
	res := new(ReportingResponseBody)
	resp, err := s.client.post(ctx, url, &params, res)

	return res, resp, err
}

// GetSearchTermLevelReports fetches reports for search terms within a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/get_search_term-level_reports
func (s *ReportingService) GetSearchTermLevelReports(ctx context.Context, campaignID int64, params *ReportingRequest) (*ReportingResponseBody, *Response, error) {
	url := fmt.Sprintf("reports/campaigns/%d/searchterms", campaignID)
	res := new(ReportingResponseBody)
	resp, err := s.client.post(ctx, url, &params, res)

	return res, resp, err
}

// GetCreativeSetLevelReports fetches reports for Creative Sets within a campaign
//
// https://developer.apple.com/documentation/apple_search_ads/get_creative_set-level_reports
func (s *ReportingService) GetCreativeSetLevelReports(ctx context.Context, campaignID int64, params *ReportingRequest) (*ReportingResponseBody, *Response, error) {
	url := fmt.Sprintf("reports/campaigns/%d/creativesets", campaignID)
	res := new(ReportingResponseBody)
	resp, err := s.client.post(ctx, url, &params, res)

	return res, resp, err
}
