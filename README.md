# apple-search-ads-go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/gungoren/apple-search-ads-go/asa)](https://pkg.go.dev/github.com/gungoren/apple-search-ads-go/asa)
[![Test Status](https://github.com/gungoren/apple-search-ads-go/workflows/Run%20Tests/badge.svg)](https://github.com/gungoren/apple-search-ads-go/actions?query=workflow%253A%2522Run+Tests%2522)
[![codecov](https://codecov.io/gh/gungoren/apple-search-ads-go/branch/master/graph/badge.svg?token=NGXNS17SV2)](https://codecov.io/gh/gungoren/apple-search-ads-go)

apple-search-ads-go is a Go client library for accessing Apple's [Apple Search Ads API](https://developer.apple.com/documentation/apple_search_ads).

## Usage

This project uses Go Modules. It requires **Go 1.16 or higher**.

```go
import "github.com/gungoren/apple-search-ads-go/asa"
```

Construct a new Apple Search Ads client, then use the various services on the client to access different parts of the Apple Search Ads API. For example:

```go
client := asa.NewClient(nil)

// list all apps with the bundle ID "com.sky.MyApp"
apps, _, err := client.Apps.ListApps(&asa.ListAppsQuery{
    FilterBundleID: []string{"com.sky.MyApp"},
})
```

The client is divided into logical chunks closely corresponding to the layout and structure of Apple's own documentation at <https://developer.apple.com/documentation/apple_search_ads>.

For more sample code snippets, head over to the [examples](https://github.com/gungoren/apple-search-ads-go/tree/master/examples) directory.

### Authentication

You may find that the code snippet above will always fail due to a lack of authorization. The Apple Search Ads API has no methods that allow for unauthorized requests. To make it easy to authenticate with Apple Search Ads, the apple-search-ads-go library offers a solution for signing and rotating JSON Web Tokens automatically. For example, the above snippet could be made to look a little more like this:

```go
import (
    "os"
    "time"

    "github.com/gungoren/apple-search-ads-go/asa"
)

func main() {
    // Key ID for the given private key, described in Apple Search Ads
    keyID := "...."
    // Issuer ID for the Apple Search Ads team
    issuerID := "...."
    // A duration value for the lifetime of a token. Apple Search Ads does not accept a token with a lifetime of longer than 20 minutes
    expiryDuration = 20*time.Minute
    // The bytes of the PKCS#8 private key created on Apple Search Ads. Keep this key safe as you can only download it once.
    privateKey = os.ReadFile("path/to/key")

    auth, err = asa.NewTokenConfig(keyID, issuerID, expiryDuration, privateKey)
    if err != nil {
        return nil, err
    }
    client := asa.NewClient(auth.Client())

    // list all apps with the bundle ID "com.sky.MyApp" in the authenticated user's team
    apps, _, err := client.Apps.ListApps(&asa.ListAppsQuery{
        FilterBundleID: []string{"com.sky.MyApp"},
    })
}
```

The authenticated client created here will automatically regenerate the token if it expires. Also note that all Apple Search Ads APIs are scoped to the credentials of the pre-configured key, so you can't use this API to make queries against the entire Apple Search Ads. For more information on creating the necessary credentials for the Apple Search Ads API, see the documentation at <https://developer.apple.com/documentation/apple_search_ads/implementing_oauth_for_the_apple_search_ads_api>.

### Rate Limiting

Apple imposes a rate limit on all API clients. The returned `Response.Rate` value contains the rate limit information from the most recent API call. If the API produces a rate limit error, it will be identifiable as an `ErrorResponse` with an error code of `429`.

Learn more about rate limiting at <https://developer.apple.com/documentation/appstoreconnectapi/identifying_rate_limits>.

### Pagination

All requests for resource collections (apps, builds, beta groups, etc.) support pagination. Responses for paginated resources will contain a `Links` property of type `PagedDocumentLinks`, with `Reference` URLs for first, next, and self. A `Reference` can have its cursor extracted with the `Cursor()` method, and that can be passed to a query param using its `Cursor` field. You can also find more information about the per-page limit and total count of resources in the response's `Meta` field of type `PagingInformation`.

```go
auth, _ = asa.NewTokenConfig(keyID, issuerID, expiryDuration, privateKey)
client := asa.NewClient(auth.Client())

opt := &asa.ListAppsQuery{
    FilterBundleID: []string{"com.sky.MyApp"},
}

var allApps []asa.App
for {
    apps, _, err := apps, _, err := client.Apps.ListApps(opt)
	if err != nil {
		return err
	}
	allApps = append(allApps, apps.Data...)
    if apps.Links.Next == nil {
        break
    }
    cursor := apps.Links.Next.Cursor()
    if cursor == "" {
        break
    }
    opt.Cursor = cursor
}
```

For complete usage of apple-search-ads-go, see the full [package docs](https://pkg.go.dev/github.com/gungoren/apple-search-ads-go/asa).

## Contributing

This project's primary goal is to cover the entire API surface exposed by the official Apple Search Ads API. Otherwise, it's being developed to aid in internal application development by the authors. Therefore, until the package's version stabilizes with v1, there isn't a strong roadmap beyond those stated goals. However, contributions are always welcome. If you want to get involved or you just want to offer feedback, please see [`CONTRIBUTING.md`](https://github.com/gungoren/.github/blob/master/CONTRIBUTING.md) for details.

## License

This library is licensed under the GNU General Public License v3.0 or later

See [LICENSE](./LICENSE) to see the full text.
