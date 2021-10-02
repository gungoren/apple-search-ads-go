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

import "time"

type ReportingResponseBody struct {
	ReportingCampaign ReportingCampaign `json:"data"`
	Pagination        PageDetail        `json:"pagination"`
}

type ReportingCampaign struct {
	ReportingDataResponse ReportingDataResponse `json:"reportingDataResponse"`
}

type ReportingDataResponse struct {
	Rows        []Row           `json:"row"`
	GrandTotals *GrandTotalsRow `json:"grandTotals,omitempty"`
}

type Row struct {
	Other       bool                `json:"other"`
	Granularity []*ExtendedSpendRow `json:"granularity,omitempty"`
	Total       *SpendRow           `json:"total,omitempty"`
	Metadata    *MetadataObject     `json:"metadata,omitempty"`
	Insights    *KeywordInsights    `json:"insights"`
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

type MetadataObject struct {
	AdGroupID                          int64                                      `json:"adGroupID"`
	AdGroupName                        string                                     `json:"adGroupName"`
	CampaignID                         int64                                      `json:"campaignId"`
	CampaignName                       string                                     `json:"campaignName"`
	Deleted                            bool                                       `json:"deleted"`
	CampaignStatus                     CampaignStatus                             `json:"campaignStatus"`
	App                                *Application                               `json:"app,omitempty"`
	ServingStatus                      CampaignServingStatus                      `json:"servingStatus"`
	ServingStateReasons                []CampaignServingStateReason               `json:"servingStateReasons"`
	CountriesOrRegions                 []Region                                   `json:"countriesOrRegions"`
	ModificationTime                   time.Time                                  `json:"modificationTime"`
	TotalBudget                        Money                                      `json:"totalBudget"`
	DailyBudget                        Money                                      `json:"dailyBudget"`
	DisplayStatus                      CampaignDisplayStatus                      `json:"displayStatus"`
	SupplySources                      []CampaignSupplySource                     `json:"supplySources"`
	AdChannelType                      CampaignAdChannelType                      `json:"adChannelType"`
	OrgID                              int                                        `json:"orgId"`
	CountryOrRegionServingStateReasons CampaignCountryOrRegionServingStateReasons `json:"countryOrRegionServingStateReasons"`
	BillingEvent                       string                                     `json:"billingEvent"`
	KeywordID                          int64                                      `json:"keywordID"`
	MatchType                          *ReportingKeywordMatchType                 `json:"matchType"`
	CountryOrRegion                    Region                                     `json:"countryOrRegion"`
	SearchTermText                     []string                                   `json:"SearchTermText"`
	SearchTermSource                   *SearchTermSource                          `json:"searchTermSource"`
}

type GrandTotalsRow struct {
	Other bool     `json:"other"`
	Total SpendRow `json:"total"`
}

type SpendRow struct {
	AvgCPA         Money   `json:"avgCPA"`
	AvgCPT         Money   `json:"avgCPT"`
	AvgCPM         Money   `json:"avgCPM"`
	ConversionRate float64 `json:"conversionRate"`
	Impressions    int64   `json:"impressions"`
	Installs       int64   `json:"installs"`
	LatOffInstalls int64   `json:"latOffInstalls"`
	LatOnInstalls  int64   `json:"latOnInstalls"`
	LocalSpend     Money   `json:"localSpend"`
	NewDownloads   int64   `json:"newDownloads"`
	ReDownloads    int64   `json:"redownloads"`
	Taps           int64   `json:"taps"`
	Ttr            float64 `json:"ttr"`
}

type ExtendedSpendRow struct {
	AvgCPA         *Money    `json:"avgCPA,omitempty"`
	AvgCPT         *Money    `json:"avgCPT,omitempty"`
	AvgCPM         *Money    `json:"avgCPM,omitempty"`
	ConversionRate float64   `json:"conversionRate,omitempty"`
	Impressions    int64     `json:"impressions,omitempty"`
	Installs       int64     `json:"installs,omitempty"`
	LatOffInstalls int64     `json:"latOffInstalls,omitempty"`
	LatOnInstalls  int64     `json:"latOnInstalls,omitempty"`
	LocalSpend     *Money    `json:"localSpend,omitempty"`
	NewDownloads   int64     `json:"newDownloads,omitempty"`
	ReDownloads    int64     `json:"redownloads,omitempty"`
	Taps           int64     `json:"taps,omitempty"`
	Ttr            float64   `json:"ttr,omitempty"`
	Date           time.Time `json:"date,omitempty"`
}

type KeywordInsights struct {
	BidRecommendation KeywordBidRecommendation `json:"bidRecommendation"`
}

type KeywordBidRecommendation struct {
	BidMax *Money `json:"bidMax"`
	BidMin *Money `json:"bidMin"`
}
