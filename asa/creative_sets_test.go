package asa

import (
	"context"
	"testing"
)

func TestGetCreativeAppAssets(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &MediaCreativeSetDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.CreativeSets.GetCreativeAppAssets(ctx, 1, &MediaCreativeSetRequest{})
	})
}

func TestGetAppPreviewDeviceSizes(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppPreviewDevicesMappingResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.CreativeSets.GetAppPreviewDeviceSizes(ctx)
	})
}
