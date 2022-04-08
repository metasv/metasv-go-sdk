# \TxApi

All URIs are relative to *https://apiv2.metasv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TxBroadcastAsyncPost**](TxApi.md#TxBroadcastAsyncPost) | **Post** /tx/broadcast/async | Broadcast tx to metasv asyn(tx size no bigger than 250K).
[**TxBroadcastAsyncTxidStateGet**](TxApi.md#TxBroadcastAsyncTxidStateGet) | **Get** /tx/broadcast/async/{txid}/state | Query aync broadcast result.
[**TxBroadcastPost**](TxApi.md#TxBroadcastPost) | **Post** /tx/broadcast | Broadcast tx to metasv fullnode.
[**TxTxidGet**](TxApi.md#TxTxidGet) | **Get** /tx/{txid} | Get transaction detail by specific txid.
[**TxTxidRawGet**](TxApi.md#TxTxidRawGet) | **Get** /tx/{txid}/raw | Get transaction raw hex by specific txid.
[**TxTxidSeenGet**](TxApi.md#TxTxidSeenGet) | **Get** /tx/{txid}/seen | Whether MetaSV have seen this tx before. This is a fast approach to know if the tx exist or not.
[**VinTxidDetailGet**](TxApi.md#VinTxidDetailGet) | **Get** /vin/{txid}/detail | Get all output point of vins in the tx with detailed output script.



## TxBroadcastAsyncPost

> string TxBroadcastAsyncPost(ctx).TxRaw(txRaw).Execute()

Broadcast tx to metasv asyn(tx size no bigger than 250K).



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
    txRaw := *openapiclient.NewTxRaw() // TxRaw |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TxApi.TxBroadcastAsyncPost(context.Background()).TxRaw(txRaw).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TxApi.TxBroadcastAsyncPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TxBroadcastAsyncPost`: string
    fmt.Fprintf(os.Stdout, "Response from `TxApi.TxBroadcastAsyncPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTxBroadcastAsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **txRaw** | [**TxRaw**](TxRaw.md) |  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxBroadcastAsyncTxidStateGet

> AsyncBroadcastResult TxBroadcastAsyncTxidStateGet(ctx, txid).Execute()

Query aync broadcast result.



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
    txid := "txid_example" // string | the request ID to search, txhash

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TxApi.TxBroadcastAsyncTxidStateGet(context.Background(), txid).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TxApi.TxBroadcastAsyncTxidStateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TxBroadcastAsyncTxidStateGet`: AsyncBroadcastResult
    fmt.Fprintf(os.Stdout, "Response from `TxApi.TxBroadcastAsyncTxidStateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txid** | **string** | the request ID to search, txhash | 

### Other Parameters

Other parameters are passed through a pointer to a apiTxBroadcastAsyncTxidStateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncBroadcastResult**](AsyncBroadcastResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxBroadcastPost

> BroadcastResult TxBroadcastPost(ctx).TxRaw(txRaw).Execute()

Broadcast tx to metasv fullnode.



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
    txRaw := *openapiclient.NewTxRaw() // TxRaw |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TxApi.TxBroadcastPost(context.Background()).TxRaw(txRaw).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TxApi.TxBroadcastPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TxBroadcastPost`: BroadcastResult
    fmt.Fprintf(os.Stdout, "Response from `TxApi.TxBroadcastPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTxBroadcastPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **txRaw** | [**TxRaw**](TxRaw.md) |  | 

### Return type

[**BroadcastResult**](BroadcastResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxTxidGet

> TxDetail TxTxidGet(ctx, txid).ShowScript(showScript).Execute()

Get transaction detail by specific txid.



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
    txid := "txid_example" // string | the request ID to search, txhash
    showScript := true // bool | Return source script code or not (default false). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TxApi.TxTxidGet(context.Background(), txid).ShowScript(showScript).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TxApi.TxTxidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TxTxidGet`: TxDetail
    fmt.Fprintf(os.Stdout, "Response from `TxApi.TxTxidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txid** | **string** | the request ID to search, txhash | 

### Other Parameters

Other parameters are passed through a pointer to a apiTxTxidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showScript** | **bool** | Return source script code or not (default false). | 

### Return type

[**TxDetail**](TxDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxTxidRawGet

> TxRaw TxTxidRawGet(ctx, txid).Execute()

Get transaction raw hex by specific txid.

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
    txid := "txid_example" // string | the request ID to search, txhash

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TxApi.TxTxidRawGet(context.Background(), txid).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TxApi.TxTxidRawGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TxTxidRawGet`: TxRaw
    fmt.Fprintf(os.Stdout, "Response from `TxApi.TxTxidRawGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txid** | **string** | the request ID to search, txhash | 

### Other Parameters

Other parameters are passed through a pointer to a apiTxTxidRawGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TxRaw**](TxRaw.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TxTxidSeenGet

> bool TxTxidSeenGet(ctx, txid).Execute()

Whether MetaSV have seen this tx before. This is a fast approach to know if the tx exist or not.

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
    txid := "txid_example" // string | the request ID to search, txhash

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TxApi.TxTxidSeenGet(context.Background(), txid).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TxApi.TxTxidSeenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TxTxidSeenGet`: bool
    fmt.Fprintf(os.Stdout, "Response from `TxApi.TxTxidSeenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txid** | **string** | the request ID to search, txhash | 

### Other Parameters

Other parameters are passed through a pointer to a apiTxTxidSeenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VinTxidDetailGet

> []OutputInfoDetail VinTxidDetailGet(ctx, txid).Execute()

Get all output point of vins in the tx with detailed output script.



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
    txid := "txid_example" // string | The txid of the vins

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TxApi.VinTxidDetailGet(context.Background(), txid).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TxApi.VinTxidDetailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VinTxidDetailGet`: []OutputInfoDetail
    fmt.Fprintf(os.Stdout, "Response from `TxApi.VinTxidDetailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txid** | **string** | The txid of the vins | 

### Other Parameters

Other parameters are passed through a pointer to a apiVinTxidDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OutputInfoDetail**](OutputInfoDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

