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

func (s *ReportingService) GetCampaignLevelReports(ctx context.Context, params *ReportingRequest) (*ReportingResponseBody, *Response, error) {
	url := "reports/campaigns"
	res := new(ReportingResponseBody)
	resp, err := s.client.post(ctx, url, &params, res)
	return res, resp, err
}

func (s *ReportingService) GetAdGroupLevelReports(ctx context.Context, campaignId int64, params *ReportingRequest) (*ReportingResponseBody, *Response, error) {
	url := fmt.Sprintf("reports/campaigns/%d/adgroups", campaignId)
	res := new(ReportingResponseBody)
	resp, err := s.client.post(ctx, url, &params, res)
	return res, resp, err
}

func (s *ReportingService) GetKeywordLevelReports(ctx context.Context, campaignId int64, params *ReportingRequest) (*ReportingResponseBody, *Response, error) {
	url := fmt.Sprintf("reports/campaigns/%d/keywords", campaignId)
	res := new(ReportingResponseBody)
	resp, err := s.client.post(ctx, url, &params, res)
	return res, resp, err
}

func (s *ReportingService) GetSearchTermLevelReports(ctx context.Context, campaignId int64, params *ReportingRequest) (*ReportingResponseBody, *Response, error) {
	url := fmt.Sprintf("reports/campaigns/%d/searchterms", campaignId)
	res := new(ReportingResponseBody)
	resp, err := s.client.post(ctx, url, &params, res)
	return res, resp, err
}

func (s *ReportingService) GetCreativeSetLevelReports(ctx context.Context, campaignId int64, params *ReportingRequest) (*ReportingResponseBody, *Response, error) {
	url := fmt.Sprintf("reports/campaigns/%d/creativesets", campaignId)
	res := new(ReportingResponseBody)
	resp, err := s.client.post(ctx, url, &params, res)
	return res, resp, err
}
