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
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCampaignLevelReports(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ReportingResponseBody{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetCampaignLevelReports(ctx, &ReportingRequest{})
	})
}

func TestGetAdGroupLevelReports(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ReportingResponseBody{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetAdGroupLevelReports(ctx, 1, &ReportingRequest{})
	})
}

func TestGetKeywordLevelReports(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ReportingResponseBody{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetKeywordLevelReports(ctx, 1, &ReportingRequest{})
	})
}

func TestGetSearchTermLevelReports(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ReportingResponseBody{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetSearchTermLevelReports(ctx, 1, &ReportingRequest{})
	})
}

func TestGetCreativeSetLevelReports(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &ReportingResponseBody{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Reporting.GetCreativeSetLevelReports(ctx, 1, &ReportingRequest{})
	})
}

func deserializeFileToReportingResponse(t *testing.T, sampleJSONResponse string) *ReportingResponseBody {
	t.Helper()

	content, err := ioutil.ReadFile(filepath.Clean(sampleJSONResponse))
	assert.NoError(t, err)

	model := &ReportingResponseBody{}

	err = json.Unmarshal(content, model)
	assert.NoError(t, err)

	return model
}

func TestGetCampaignLevelReportsResponseDeserialization(t *testing.T) {
	t.Parallel()

	sampleJSONResponse := "../test/response_body_json_files/get_campaign_level_reports.json"

	deserializeFileToReportingResponse(t, sampleJSONResponse)
}

func TestGetCampaignLevelReportsWithGranularityResponseDeserialization(t *testing.T) {
	t.Parallel()

	sampleJSONResponse := "../test/response_body_json_files/get_campaign_level_reports_with_granularity.json"

	deserializeFileToReportingResponse(t, sampleJSONResponse)
}

func TestGetAdGroupLevelReportsResponseDeserialization(t *testing.T) {
	t.Parallel()

	sampleJSONResponse := "../test/response_body_json_files/get_ad_group_level_reports.json"

	deserializeFileToReportingResponse(t, sampleJSONResponse)
}

func TestGetKeywordLevelReportsResponseDeserialization(t *testing.T) {
	t.Parallel()

	sampleJSONResponse := "../test/response_body_json_files/get_keyword_level_reports.json"

	deserializeFileToReportingResponse(t, sampleJSONResponse)
}

func TestGetSearchTermLevelReportsResponseDeserialization(t *testing.T) {
	t.Parallel()

	sampleJSONResponse := "../test/response_body_json_files/get_search_term_level_reports.json"

	deserializeFileToReportingResponse(t, sampleJSONResponse)
}

func TestEmptyResponseDeserialization(t *testing.T) {
	t.Parallel()

	sampleJSONResponse := "../test/response_body_json_files/reporting_response_body_with_empty.json"

	report := deserializeFileToReportingResponse(t, sampleJSONResponse)

	assert.Equal(t, 0, len(report.ReportingCampaign.ReportingDataResponse.Rows))
	assert.Equal(t, 0, report.Pagination.TotalResults)
}

func TestGetCampaignLevelReportsRequestSerialization(t *testing.T) {
	t.Parallel()

	pagination := &Pagination{
		Limit:  1000,
		Offset: 0,
	}
	selector := &Selector{
		Conditions: []*Condition{
			{
				Field:    "countriesOrRegions",
				Operator: ConditionOperatorContainsAny,
				Values:   []string{"US", "GB"},
			},
			{
				Field:    "countryOrRegion",
				Operator: ConditionOperatorIn,
				Values:   []string{"US"},
			},
		},
		Pagination: pagination,
		OrderBy: []*Sorting{
			{
				Field:     "countryOrRegion",
				SortOrder: SortingOrderAscending,
			},
		},
	}

	startTime, _ := time.Parse("2006-01-02", "2020-08-04")
	endTimeTime, _ := time.Parse("2006-01-02", "2020-08-14")

	reportingRequest := &ReportingRequest{
		StartTime:                  Date{startTime},
		EndTime:                    Date{endTimeTime},
		TimeZone:                   ReportingRequestTimeZoneUTC,
		GroupBy:                    []ReportingRequestGroupBy{ReportingRequestGroupByTypeCountryOrRegion},
		ReturnGrandTotals:          true,
		ReturnRecordsWithNoMetrics: true,
		ReturnRowTotals:            true,
		Selector:                   selector,
	}

	val1, err := json.Marshal(reportingRequest)
	assert.NoError(t, err)

	val2, err := ioutil.ReadFile("../test/request_body_json_files/get_campaign_level_reports.json")
	assert.NoError(t, err)

	require.JSONEq(t, string(val1), string(val2))
}

func TestGetCampaignLevelReportsWithGranularityRequestSerialization(t *testing.T) {
	t.Parallel()

	pagination := &Pagination{
		Limit:  1000,
		Offset: 0,
	}
	selector := &Selector{
		Conditions: []*Condition{
			{
				Field:    "countriesOrRegions",
				Operator: ConditionOperatorContainsAny,
				Values:   []string{"US", "GB"},
			},
			{
				Field:    "countryOrRegion",
				Operator: ConditionOperatorIn,
				Values:   []string{"US"},
			},
		},
		Pagination: pagination,
		OrderBy: []*Sorting{
			{
				Field:     "countryOrRegion",
				SortOrder: SortingOrderAscending,
			},
		},
	}

	startTime, _ := time.Parse("2006-01-02", "2020-08-04")
	endTimeTime, _ := time.Parse("2006-01-02", "2020-08-14")

	reportingRequest := &ReportingRequest{
		StartTime:                  Date{startTime},
		EndTime:                    Date{endTimeTime},
		TimeZone:                   ReportingRequestTimeZoneUTC,
		GroupBy:                    []ReportingRequestGroupBy{ReportingRequestGroupByTypeCountryOrRegion},
		ReturnGrandTotals:          false,
		ReturnRecordsWithNoMetrics: true,
		ReturnRowTotals:            false,
		Granularity:                ReportingRequestGranularityTypeDaily,
		Selector:                   selector,
	}

	val1, err := json.Marshal(reportingRequest)
	assert.NoError(t, err)

	val2, err := ioutil.ReadFile("../test/request_body_json_files/get_campaign_level_reports_with_granularity.json")
	assert.NoError(t, err)

	require.JSONEq(t, string(val1), string(val2))
}

func TestKeywordLevelReportsRequestSerialization(t *testing.T) {
	t.Parallel()

	pagination := &Pagination{
		Limit:  1000,
		Offset: 0,
	}
	selector := &Selector{
		Conditions: []*Condition{
			{
				Field:    "deleted",
				Operator: ConditionOperatorIn,
				Values:   []string{"false", "true"},
			},
		},
		Pagination: pagination,
		OrderBy: []*Sorting{
			{
				Field:     "localSpend",
				SortOrder: SortingOrderAscending,
			},
		},
	}

	startTime, _ := time.Parse("2006-01-02", "2020-07-01")
	endTimeTime, _ := time.Parse("2006-01-02", "2020-07-02")

	reportingRequest := &ReportingRequest{
		StartTime:                  Date{startTime},
		EndTime:                    Date{endTimeTime},
		TimeZone:                   ReportingRequestTimeZoneUTC,
		ReturnGrandTotals:          true,
		ReturnRecordsWithNoMetrics: true,
		ReturnRowTotals:            true,
		Granularity:                ReportingRequestGranularityTypeDaily,
		Selector:                   selector,
	}

	val1, err := json.Marshal(reportingRequest)
	assert.NoError(t, err)

	val2, err := ioutil.ReadFile("../test/request_body_json_files/get_keyword_level_reports.json")
	assert.NoError(t, err)

	require.JSONEq(t, string(val1), string(val2))
}

func TestSearchTermLevelReportsRequestSerialization(t *testing.T) {
	t.Parallel()

	pagination := &Pagination{
		Limit:  1000,
		Offset: 0,
	}
	selector := &Selector{
		Pagination: pagination,
		OrderBy: []*Sorting{
			{
				Field:     "impressions",
				SortOrder: SortingOrderDescending,
			},
		},
	}

	startTime, _ := time.Parse("2006-01-02", "2020-06-03")
	endTimeTime, _ := time.Parse("2006-01-02", "2020-06-11")

	reportingRequest := &ReportingRequest{
		StartTime:                  Date{startTime},
		EndTime:                    Date{endTimeTime},
		TimeZone:                   ReportingRequestTimeZoneORTZ,
		ReturnGrandTotals:          true,
		ReturnRecordsWithNoMetrics: false,
		ReturnRowTotals:            true,
		GroupBy:                    []ReportingRequestGroupBy{ReportingRequestGroupByTypeCountryOrRegion},
		Selector:                   selector,
	}

	val1, err := json.Marshal(reportingRequest)
	assert.NoError(t, err)

	val2, err := ioutil.ReadFile("../test/request_body_json_files/get_search_term_level_reports.json")
	assert.NoError(t, err)

	require.JSONEq(t, string(val1), string(val2))
}
