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

type CampaignListResponse struct {
	Error      ErrorResponseBody `json:"error"`
	Pagination PageDetail        `json:"pagination"`
	Campaigns  []Campaign        `json:"data"`
}

type ErrorResponseBody struct {
	Errors []ErrorResponseItem `json:"errors"`
}

type GeneralErrorResponse struct {
	Error ErrorResponseBody `json:"error"`
}

type ErrorResponseItemMessageCode string

const (
	ErrorResponseItemMessageCodeUnauthorized      ErrorResponseItemMessageCode = "UNAUTHORIZED"
	ErrorResponseItemMessageCodeInvalidDateFormat ErrorResponseItemMessageCode = "INVALID_DATE_FORMAT"
)

type ErrorResponseItem struct {
	Field       string                       `json:"field"`
	Message     string                       `json:"message"`
	MessageCode ErrorResponseItemMessageCode `json:"messageCode"`
}
