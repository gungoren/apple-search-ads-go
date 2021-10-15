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

// CreativeSetsService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/creative_sets
type CreativeSetsService service

// MediaCreativeSetRequest is the request body for getting Creative Set assets
//
// https://developer.apple.com/documentation/apple_search_ads/mediacreativesetrequest
type MediaCreativeSetRequest struct {
	AssetsGenIds     []string `json:"assetsGenIds,omitempty"`
	CountryOrRegions []string `json:"countryOrRegions"`
}

// MediaCreativeSetDetailResponse is the response data to Creative Set asset requests
//
// https://developer.apple.com/documentation/apple_search_ads/mediacreativesetdetailresponse
type MediaCreativeSetDetailResponse struct {
	CreativeSetAssetsDetail *CreativeSetAssetsDetail `json:"data,omitempty"`
}

// CreativeSetAssetsDetail is the asset details to create a Creative Set
//
// https://developer.apple.com/documentation/apple_search_ads/creativesetassetsdetail
type CreativeSetAssetsDetail struct {
	CreativeSetDetails map[string]CreativeSetLocaleDetail `json:"creativeSetDetails,omitempty"`
}

// CreativeSetLocaleDetail is the localized information about a Creative Set
//
// https://developer.apple.com/documentation/apple_search_ads/creativesetlocaledetail
type CreativeSetLocaleDetail struct {
	AppPreviewDeviceWithAssets map[string]MediaAppPreviewOrScreenshotsDetail `json:"appPreviewDeviceWithAssets,omitempty"`
	IsPrimaryLocale            bool                                          `json:"isPrimaryLocale,omitempty"`
	LanguageCode               string                                        `json:"languageCode,omitempty"`
	LanguageDisplayName        string                                        `json:"languageDisplayName"`
}

// MediaAppPreviewOrScreenshotsDetail is the app asset details of a device
//
// https://developer.apple.com/documentation/apple_search_ads/mediaappprevieworscreenshotsdetail
type MediaAppPreviewOrScreenshotsDetail struct {
	DeviceDisplayName           string                          `json:"deviceDisplayName,omitempty"`
	FallBackDevicesDisplayNames map[string]string               `json:"fallBackDevicesDisplayNames,omitempty"`
	Screenshots                 []*MediaAppPreviewOrScreenshots `json:"screenshots,omitempty"`
	AppPreviews                 []*MediaAppPreviewOrScreenshots `json:"appPreviews,omitempty"`
}

// MediaAppPreviewOrScreenshotsAssetType The type of asset.
type MediaAppPreviewOrScreenshotsAssetType string

const (
	// MediaAppPreviewOrScreenshotsAssetTypeAppPreview is for a media app preview screenshot asset type on App Preview.
	MediaAppPreviewOrScreenshotsAssetTypeAppPreview MediaAppPreviewOrScreenshotsAssetType = "APP_PREVIEW"
	// MediaAppPreviewOrScreenshotsAssetTypeScreenshot is for a media app preview screenshot asset type on App Screenshot.
	MediaAppPreviewOrScreenshotsAssetTypeScreenshot MediaAppPreviewOrScreenshotsAssetType = "SCREENSHOT"
)

// MediaAppPreviewOrScreenshotsOrientation is the orientation of the asset that you upload to App Store Connect.
type MediaAppPreviewOrScreenshotsOrientation string

const (
	// MediaAppPreviewOrScreenshotsOrientationPortrait is for a media app preview or screenshots orientation on Portrait.
	MediaAppPreviewOrScreenshotsOrientationPortrait MediaAppPreviewOrScreenshotsOrientation = "PORTRAIT"
	// MediaAppPreviewOrScreenshotsOrientationLandscape  is for a media app preview or screenshots orientation on Landscape.
	MediaAppPreviewOrScreenshotsOrientationLandscape MediaAppPreviewOrScreenshotsOrientation = "LANDSCAPE"
	// MediaAppPreviewOrScreenshotsOrientationUnknown  is for a media app preview or screenshots orientation on Unknown.
	MediaAppPreviewOrScreenshotsOrientationUnknown MediaAppPreviewOrScreenshotsOrientation = "UNKNOWN"
)

// MediaAppPreviewOrScreenshots is the asset details of the app preview or app screenshots
//
// https://developer.apple.com/documentation/apple_search_ads/mediaappprevieworscreenshots
type MediaAppPreviewOrScreenshots struct {
	AssetGenID   string                                  `json:"assetGenId,omitempty"`
	AssetType    MediaAppPreviewOrScreenshotsAssetType   `json:"assetType"`
	AssetURL     string                                  `json:"assetURL,omitempty"`
	Orientation  MediaAppPreviewOrScreenshotsOrientation `json:"orientation"`
	SortPosition int64                                   `json:"sortPosition,omitempty"`
	SourceHeight int32                                   `json:"sourceHeight,omitempty"`
	SourceWidth  int32                                   `json:"sourceWidth,omitempty"`
}

// GetCreativeAppAssets Fetches assets to use with Creative Sets
//
// https://developer.apple.com/documentation/apple_search_ads/get_app_language_device_sizes_and_asset_details
func (s *CreativeSetsService) GetCreativeAppAssets(ctx context.Context, adamID int64, params *MediaCreativeSetRequest) (*MediaCreativeSetDetailResponse, *Response, error) {
	url := fmt.Sprintf("creativeappassets/%d", adamID)
	res := new(MediaCreativeSetDetailResponse)
	resp, err := s.client.post(ctx, url, *params, res)

	return res, resp, err
}

// AppPreviewDevicesMappingResponse is the app preview device mapping response to display name and size mapping requests
//
// https://developer.apple.com/documentation/apple_search_ads/apppreviewdevicesmappingresponse
type AppPreviewDevicesMappingResponse struct {
	AppPreviewDevices map[string]string `json:"data"`
}

// GetAppPreviewDeviceSizes Fetches supported app preview device size mappings
//
// https://developer.apple.com/documentation/apple_search_ads/get_app_language_device_sizes_and_asset_details
func (s *CreativeSetsService) GetAppPreviewDeviceSizes(ctx context.Context) (*AppPreviewDevicesMappingResponse, *Response, error) {
	url := "creativeappassets/devices"
	res := new(AppPreviewDevicesMappingResponse)
	resp, err := s.client.get(ctx, url, nil, res)

	return res, resp, err
}

// CreateAdGroupCreativeSetRequest is the response to a request to create an ad group Creative Set
//
// https://developer.apple.com/documentation/apple_search_ads/createadgroupcreativesetrequest
type CreateAdGroupCreativeSetRequest struct {
	CreativeSet *CreativeSetCreate `json:"creativeSet,omitempty"`
}

// CreativeSetCreate is the response to creating a Creative Set
//
// https://developer.apple.com/documentation/apple_search_ads/creativesetcreate
type CreativeSetCreate struct {
	AdamID       int64    `json:"adamId,omitempty"`
	Name         string   `json:"name,omitempty"`
	LanguageCode string   `json:"languageCode,omitempty"`
	AssetsGenIds []string `json:"assetsGenIds,omitempty"`
}

// AdGroupCreativeSetResponse is a container for the ad group Creative Set response body
//
// https://developer.apple.com/documentation/apple_search_ads/adgroupcreativesetresponse
type AdGroupCreativeSetResponse struct {
	AdGroupCreativeSet *AdGroupCreativeSet `json:"data,omitempty"`
	Error              *ErrorResponseBody  `json:"error,omitempty"`
	Pagination         *PageDetail         `json:"pagination,omitempty"`
}

// AdGroupCreativeSet is the assignment relationship between an ad group and a Creative Set
//
// https://developer.apple.com/documentation/apple_search_ads/adgroupcreativeset
type AdGroupCreativeSet struct {
	AdGroupID            int64                            `json:"adGroupId,omitempty"`
	CampaignID           int64                            `json:"campaignId,omitempty"`
	CreativeSetID        int64                            `json:"creativeSetId,omitempty"`
	Deleted              bool                             `json:"deleted"`
	ID                   int64                            `json:"id"`
	ModificationTime     DateTime                         `json:"modificationTime"`
	ServingStatus        AdGroupServingStatus             `json:"servingStatus"`
	ServingStatusReasons []CreativeSetsServingStateReason `json:"servingStatusReasons"`
	Status               AdGroupStatus                    `json:"status"`
}

// CreativeSetsServingStateReason is a reason when a adgroupcreativeset is not running.
type CreativeSetsServingStateReason string

const (
	// CreativeSetsServingStateReasonPausedBySystem is for a adgroup creative set serving state reason for PAUSED_BY_SYSTEM.
	CreativeSetsServingStateReasonPausedBySystem CreativeSetsServingStateReason = "PAUSED_BY_SYSTEM"
	// CreativeSetsServingStateReasonPausedByUser is for a adgroup creative set serving state reason for PAUSED_BY_USER.
	CreativeSetsServingStateReasonPausedByUser CreativeSetsServingStateReason = "PAUSED_BY_USER"
	// CreativeSetsServingStateReasonDeletedByUser is for a adgroup creative set serving state reason for DELETED_BY_USER.
	CreativeSetsServingStateReasonDeletedByUser CreativeSetsServingStateReason = "DELETED_BY_USER"
	// CreativeSetsServingStateReasonCreativeSetInvalid is for a adgroup creative set serving state reason for CREATIVE_SET_INVALID.
	CreativeSetsServingStateReasonCreativeSetInvalid CreativeSetsServingStateReason = "CREATIVE_SET_INVALID"
)

// CreateAdGroupCreativeSets Creates a Creative Set and assigns it to an ad group
//
// https://developer.apple.com/documentation/apple_search_ads/create_ad_group_creative_sets
func (s *CreativeSetsService) CreateAdGroupCreativeSets(ctx context.Context, campaignID int64, adgroupID int64, body *CreateAdGroupCreativeSetRequest) (*AdGroupCreativeSetResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/adgroupcreativesets/creativesets", campaignID, adgroupID)
	res := new(AdGroupCreativeSetResponse)
	resp, err := s.client.post(ctx, url, body, res)

	return res, resp, err
}

// FindAdGroupCreativeSetRequest Selector objects available to filter returned data.
//
// https://developer.apple.com/documentation/apple_search_ads/findadgroupcreativesetrequest
type FindAdGroupCreativeSetRequest struct {
	Selector *Selector `json:"selector"`
}

// AdGroupCreativeSetListResponse is the response details of ad group Creative Set requests
//
// https://developer.apple.com/documentation/apple_search_ads/adgroupcreativesetlistresponse
type AdGroupCreativeSetListResponse struct {
	AdGroupCreativeSets []*AdGroupCreativeSet `json:"data,omitempty"`
	Error               *ErrorResponseBody    `json:"error,omitempty"`
	Pagination          *PageDetail           `json:"pagination,omitempty"`
}

// FindAdGroupCreativeSets Fetches all assigned Creative Sets for ad groups
//
// https://developer.apple.com/documentation/apple_search_ads/find_ad_group_creative_sets
func (s *CreativeSetsService) FindAdGroupCreativeSets(ctx context.Context, campaignID int64, body *FindAdGroupCreativeSetRequest) (*AdGroupCreativeSetListResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroupcreativesets/find", campaignID)
	res := new(AdGroupCreativeSetListResponse)
	resp, err := s.client.post(ctx, url, body, res)

	return res, resp, err
}

// AdGroupCreativeSetUpdate is the response to ad group Creative Set update requests
//
// https://developer.apple.com/documentation/apple_search_ads/adgroupcreativesetupdate
type AdGroupCreativeSetUpdate struct {
	Status AdGroupStatus `json:"status"`
}

// UpdateAdGroupCreativeSets Updates an ad group Creative Set using an identifier
//
// https://developer.apple.com/documentation/apple_search_ads/update_ad_group_creative_sets
func (s *CreativeSetsService) UpdateAdGroupCreativeSets(ctx context.Context, campaignID int64, adgroupID int64, adGroupCreativeSetID int64, body *AdGroupCreativeSetUpdate) (*AdGroupCreativeSetResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/adgroupcreativesets/%d", campaignID, adgroupID, adGroupCreativeSetID)
	res := new(AdGroupCreativeSetResponse)
	resp, err := s.client.put(ctx, url, body, res)

	return res, resp, err
}

// DeleteAdGroupCreativeSets Deletes Creative Sets from a specified ad group
//
// https://developer.apple.com/documentation/apple_search_ads/delete_ad_group_creative_sets
func (s *CreativeSetsService) DeleteAdGroupCreativeSets(ctx context.Context, campaignID int64, adgroupID int64, adGroupCreativeSetIDs []int64) (*IntegerResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/adgroupcreativesets/delete/bulk", campaignID, adgroupID)
	res := new(IntegerResponse)
	resp, err := s.client.post(ctx, url, adGroupCreativeSetIDs, res)

	return res, resp, err
}

// CreativeSetResponse is the response to update a Creative Set request
//
// https://developer.apple.com/documentation/apple_search_ads/creativesetresponse
type CreativeSetResponse struct {
	CreativeSet *CreativeSet       `json:"data,omitempty"`
	Error       *ErrorResponseBody `json:"error,omitempty"`
}

// CreativeSetStatus is the user-controlled status to enable or pause the Creative Set.
type CreativeSetStatus string

const (
	// CreativeSetStatusValid is for a creative set status on Valid.
	CreativeSetStatusValid CreativeSetStatus = "VALID"
	// CreativeSetStatusInvalid is for a creative set status on InValid.
	CreativeSetStatusInvalid CreativeSetStatus = "INVALID"
)

// CreativeSetStatusReason is the reason for the Creative Set status.
type CreativeSetStatusReason string

const (
	// CreativeSetStatusReasonAssetDeleted is for a creative set status reason on Asset Deleted.
	CreativeSetStatusReasonAssetDeleted CreativeSetStatusReason = "ASSET_DELETED"
)

// CreativeSet is the basic details of a Creative Set
//
// https://developer.apple.com/documentation/apple_search_ads/creativeset
type CreativeSet struct {
	ID                int64                     `json:"id,omitempty"`
	Name              string                    `json:"name,omitempty"`
	AdamID            int64                     `json:"adamID,omitempty"`
	CreativeSetAssets []*CreativeSetAsset       `json:"creativeSetAssets,omitempty"`
	LanguageCode      string                    `json:"languageCode,omitempty"`
	OrgID             int64                     `json:"orgID,omitempty"`
	Status            CreativeSetStatus         `json:"status,omitempty"`
	StatusReasons     []CreativeSetStatusReason `json:"statusReasons,omitempty"`
}

// CreativeSetAsset is the assets of a Creative Set
//
// https://developer.apple.com/documentation/apple_search_ads/creativesetasset
type CreativeSetAsset struct {
	Asset *Asset `json:"asset,omitempty"`
	ID    int64  `json:"id,omitempty"`
}

// Asset is the assets for creating Creative Sets
//
// https://developer.apple.com/documentation/apple_search_ads/asset
type Asset struct {
	AppPreviewDevice string                                  `json:"appPreviewDevice,omitempty"`
	AssetGenID       string                                  `json:"assetGenId,omitempty"`
	Deleted          bool                                    `json:"deleted"`
	Orientation      MediaAppPreviewOrScreenshotsOrientation `json:"orientation"`
	Type             MediaAppPreviewOrScreenshotsAssetType   `json:"type"`
}

// GetCreativeSetVariationQuery defines query parameter for GetCreativeSetVariation endpoint.
type GetCreativeSetVariationQuery struct {
	IncludeDeletedCreativeSetAssets bool `url:"includeDeletedCreativeSetAssets,omitempty"`
}

// GetCreativeSetVariation Get a Creative Set Ad Variation
//
// https://developer.apple.com/documentation/apple_search_ads/get_a_creative_set_ad_variation
func (s *CreativeSetsService) GetCreativeSetVariation(ctx context.Context, creativeSetID int64, params *GetCreativeSetVariationQuery) (*CreativeSetResponse, *Response, error) {
	url := fmt.Sprintf("creativesets/%d", creativeSetID)
	res := new(CreativeSetResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// FindCreativeSetRequest is the request to find Creative Sets
//
// https://developer.apple.com/documentation/apple_search_ads/findcreativesetrequest
type FindCreativeSetRequest struct {
	Selector                        *Selector `json:"selector,omitempty"`
	IncludeDeletedCreativeSetAssets bool      `json:"includeDeletedCreativeSetAssets,omitempty"`
}

// CreativeSetListResponse is the response to the request to find Creative Sets
//
// https://developer.apple.com/documentation/apple_search_ads/creativesetlistresponse
type CreativeSetListResponse struct {
	CreativeSets []*CreativeSet     `json:"data,omitempty"`
	Error        *ErrorResponseBody `json:"error,omitempty"`
	Pagination   *PageDetail        `json:"pagination,omitempty"`
}

// FindCreativeSets Fetches all assigned Creative Sets for an organization
//
// https://developer.apple.com/documentation/apple_search_ads/find_creative_sets
func (s *CreativeSetsService) FindCreativeSets(ctx context.Context, params *FindCreativeSetRequest) (*CreativeSetListResponse, *Response, error) {
	url := "creativesets/find"
	res := new(CreativeSetListResponse)
	resp, err := s.client.post(ctx, url, params, res)

	return res, resp, err
}

// AssignAdGroupCreativeSetRequest is the request to assign a Creative Set to an ad group
//
// https://developer.apple.com/documentation/apple_search_ads/assignadgroupcreativesetrequest
type AssignAdGroupCreativeSetRequest struct {
	CreativeSetID int64 `json:"creativeSetID"`
}

// AssignCreativeSetsToAdGroup Creates a Creative Set assignment to an ad group
//
// https://developer.apple.com/documentation/apple_search_ads/assign_creative_sets_to_an_ad_group
func (s *CreativeSetsService) AssignCreativeSetsToAdGroup(ctx context.Context, campaignID int64, adgroupID int64, request *AssignAdGroupCreativeSetRequest) (*AdGroupCreativeSetResponse, *Response, error) {
	url := fmt.Sprintf("campaigns/%d/adgroups/%d/adgroupcreativesets", campaignID, adgroupID)
	res := new(AdGroupCreativeSetResponse)
	resp, err := s.client.post(ctx, url, request, res)

	return res, resp, err
}

// CreativeSetUpdate is the details of an update to a Creative Set request
//
// https://developer.apple.com/documentation/apple_search_ads/creativesetupdate
type CreativeSetUpdate struct {
	Name string `json:"name"`
}

// UpdateCreativeSets Updates a Creative Set name using an identifier
//
// https://developer.apple.com/documentation/apple_search_ads/update_creative_sets
func (s *CreativeSetsService) UpdateCreativeSets(ctx context.Context, creativeSetID int64, request *CreativeSetUpdate) (*CreativeSetResponse, *Response, error) {
	url := fmt.Sprintf("creativesets/%d", creativeSetID)
	res := new(CreativeSetResponse)
	resp, err := s.client.put(ctx, url, request, res)

	return res, resp, err
}
