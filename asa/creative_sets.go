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
	AssetGenId   string                                  `json:"assetGenId,omitempty"`
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
	url := fmt.Sprintf("/creativeappassets/%d", adamID)
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
	url := "/creativeappassets/devices"
	res := new(AppPreviewDevicesMappingResponse)
	resp, err := s.client.get(ctx, url, nil, res)

	return res, resp, err
}
