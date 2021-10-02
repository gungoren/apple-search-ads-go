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
