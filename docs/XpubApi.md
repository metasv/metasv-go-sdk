# \XpubApi

All URIs are relative to *https://apiv2.metasv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**XpubLiteXpubAddressAddressGet**](XpubApi.md#XpubLiteXpubAddressAddressGet) | **Get** /xpubLite/{xpub}/address/{address} | Get xpub address type and index. Only index under /0/70 /1/30 is valid. Otherwise not found.
[**XpubLiteXpubBalanceGet**](XpubApi.md#XpubLiteXpubBalanceGet) | **Get** /xpubLite/{xpub}/balance | Get xpub balances including confirmed and unconfirmed.
[**XpubLiteXpubUtxoGet**](XpubApi.md#XpubLiteXpubUtxoGet) | **Get** /xpubLite/{xpub}/utxo | Get xpub utxos by specific xpub(300 per page).
[**XpubPost**](XpubApi.md#XpubPost) | **Post** /xpub | Register a new xpub.
[**XpubXpubAddressAddressGet**](XpubApi.md#XpubXpubAddressAddressGet) | **Get** /xpub/{xpub}/address/{address} | Get xpub address type and index. Only index under max index is valid. Otherwise not found.
[**XpubXpubBalanceGet**](XpubApi.md#XpubXpubBalanceGet) | **Get** /xpub/{xpub}/balance | Get xpub balances including confirmed and unconfirmed.
[**XpubXpubDelete**](XpubApi.md#XpubXpubDelete) | **Delete** /xpub/{xpub} | Delete a registered xpub.
[**XpubXpubGet**](XpubApi.md#XpubXpubGet) | **Get** /xpub/{xpub} | Get xpub detail and status. Only registered xpub available.
[**XpubXpubTxsGet**](XpubApi.md#XpubXpubTxsGet) | **Get** /xpub/{xpub}/txs | Get xpub transaction history by specific xpub(100 per page).
[**XpubXpubUtxoCountGet**](XpubApi.md#XpubXpubUtxoCountGet) | **Get** /xpub/{xpub}/utxo/count | Get valid utxo count(including confirmed and unconfirmed).
[**XpubXpubUtxoGet**](XpubApi.md#XpubXpubUtxoGet) | **Get** /xpub/{xpub}/utxo | Get xpub utxos by specific xpub.
[**XpubsCountGet**](XpubApi.md#XpubsCountGet) | **Get** /xpubs/count | Get the total count of registered xpubs.
[**XpubsGet**](XpubApi.md#XpubsGet) | **Get** /xpubs | Get all registered xpub strings. 300 per page.



## XpubLiteXpubAddressAddressGet

> XpubAddress XpubLiteXpubAddressAddressGet(ctx, xpub, address).Execute()

Get xpub address type and index. Only index under /0/70 /1/30 is valid. Otherwise not found.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xpub := "xpub_example" // string | the requested xpub
    address := "address_example" // string | the requested address

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubLiteXpubAddressAddressGet(context.Background(), xpub, address).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubLiteXpubAddressAddressGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubLiteXpubAddressAddressGet`: XpubAddress
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubLiteXpubAddressAddressGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the requested xpub | 
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubLiteXpubAddressAddressGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**XpubAddress**](XpubAddress.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubLiteXpubBalanceGet

> XpubLiteBalance XpubLiteXpubBalanceGet(ctx, xpub).Execute()

Get xpub balances including confirmed and unconfirmed.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xpub := "xpub_example" // string | the xpub to search

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubLiteXpubBalanceGet(context.Background(), xpub).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubLiteXpubBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubLiteXpubBalanceGet`: XpubLiteBalance
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubLiteXpubBalanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the xpub to search | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubLiteXpubBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**XpubLiteBalance**](XpubLiteBalance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubLiteXpubUtxoGet

> []XpubUtxo XpubLiteXpubUtxoGet(ctx, xpub).Limit(limit).Execute()

Get xpub utxos by specific xpub(300 per page).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xpub := "xpub_example" // string | the requested xpub
    limit := int32(56) // int32 | The max items returned in this query(default 300), not bigger than 5000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubLiteXpubUtxoGet(context.Background(), xpub).Limit(limit).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubLiteXpubUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubLiteXpubUtxoGet`: []XpubUtxo
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubLiteXpubUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the requested xpub | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubLiteXpubUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max items returned in this query(default 300), not bigger than 5000. | 

### Return type

[**[]XpubUtxo**](XpubUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubPost

> XpubResult XpubPost(ctx).XpubRequest(xpubRequest).Execute()

Register a new xpub.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xpubRequest := *openapiclient.NewXpubRequest() // XpubRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubPost(context.Background()).XpubRequest(xpubRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubPost`: XpubResult
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXpubPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xpubRequest** | [**XpubRequest**](XpubRequest.md) |  | 

### Return type

[**XpubResult**](XpubResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubXpubAddressAddressGet

> XpubAddress XpubXpubAddressAddressGet(ctx, xpub, address).Execute()

Get xpub address type and index. Only index under max index is valid. Otherwise not found.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xpub := "xpub_example" // string | the requested xpub
    address := "address_example" // string | the requested address

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubXpubAddressAddressGet(context.Background(), xpub, address).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubXpubAddressAddressGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubXpubAddressAddressGet`: XpubAddress
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubXpubAddressAddressGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the requested xpub | 
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubXpubAddressAddressGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**XpubAddress**](XpubAddress.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubXpubBalanceGet

> XpubBalance XpubXpubBalanceGet(ctx, xpub).Execute()

Get xpub balances including confirmed and unconfirmed.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xpub := "xpub_example" // string | the xpub to search

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubXpubBalanceGet(context.Background(), xpub).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubXpubBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubXpubBalanceGet`: XpubBalance
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubXpubBalanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the xpub to search | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubXpubBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**XpubBalance**](XpubBalance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubXpubDelete

> XpubResult XpubXpubDelete(ctx, xpub).Execute()

Delete a registered xpub.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xpub := "xpub_example" // string | the xpub to search

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubXpubDelete(context.Background(), xpub).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubXpubDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubXpubDelete`: XpubResult
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubXpubDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the xpub to search | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubXpubDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**XpubResult**](XpubResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubXpubGet

> XpubDetail XpubXpubGet(ctx, xpub).Execute()

Get xpub detail and status. Only registered xpub available.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xpub := "xpub_example" // string | the xpub to search

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubXpubGet(context.Background(), xpub).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubXpubGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubXpubGet`: XpubDetail
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubXpubGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the xpub to search | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubXpubGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**XpubDetail**](XpubDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubXpubTxsGet

> []XPubTransaction XpubXpubTxsGet(ctx, xpub).Flag(flag).Execute()

Get xpub transaction history by specific xpub(100 per page).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xpub := "xpub_example" // string | the requested xpub
    flag := "flag_example" // string | The last id of the last query(Paging flag) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubXpubTxsGet(context.Background(), xpub).Flag(flag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubXpubTxsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubXpubTxsGet`: []XPubTransaction
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubXpubTxsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the requested xpub | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubXpubTxsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flag** | **string** | The last id of the last query(Paging flag) | 

### Return type

[**[]XPubTransaction**](XPubTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubXpubUtxoCountGet

> int32 XpubXpubUtxoCountGet(ctx, xpub).Execute()

Get valid utxo count(including confirmed and unconfirmed).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xpub := "xpub_example" // string | the requested xpub

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubXpubUtxoCountGet(context.Background(), xpub).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubXpubUtxoCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubXpubUtxoCountGet`: int32
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubXpubUtxoCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the requested xpub | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubXpubUtxoCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**int32**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubXpubUtxoGet

> []XpubUtxo XpubXpubUtxoGet(ctx, xpub).Limit(limit).Execute()

Get xpub utxos by specific xpub.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xpub := "xpub_example" // string | the requested xpub
    limit := int64(789) // int64 | The max items returned in this query(default 300), not bigger than 5000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubXpubUtxoGet(context.Background(), xpub).Limit(limit).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubXpubUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubXpubUtxoGet`: []XpubUtxo
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubXpubUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xpub** | **string** | the requested xpub | 

### Other Parameters

Other parameters are passed through a pointer to a apiXpubXpubUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | The max items returned in this query(default 300), not bigger than 5000. | 

### Return type

[**[]XpubUtxo**](XpubUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubsCountGet

> int32 XpubsCountGet(ctx).Execute()

Get the total count of registered xpubs.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubsCountGet(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubsCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubsCountGet`: int32
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubsCountGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiXpubsCountGetRequest struct via the builder pattern


### Return type

**int32**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XpubsGet

> []string XpubsGet(ctx).Flag(flag).Execute()

Get all registered xpub strings. 300 per page.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    flag := "flag_example" // string | The last xpub of the last query(Paging flag) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.XpubApi.XpubsGet(context.Background()).Flag(flag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `XpubApi.XpubsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XpubsGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `XpubApi.XpubsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXpubsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flag** | **string** | The last xpub of the last query(Paging flag) | 

### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

