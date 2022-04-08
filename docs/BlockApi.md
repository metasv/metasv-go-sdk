# \BlockApi

All URIs are relative to *https://apiv2.metasv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockBlockIdGet**](BlockApi.md#BlockBlockIdGet) | **Get** /block/{blockId} | Get block request by height or hash
[**BlockGet**](BlockApi.md#BlockGet) | **Get** /block | Get recent block list by paging. 30 items per page.
[**BlockInfoGet**](BlockApi.md#BlockInfoGet) | **Get** /block/info | Get current blockchain info from full node.



## BlockBlockIdGet

> BlockHeader BlockBlockIdGet(ctx, blockId).Execute()

Get block request by height or hash

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
    blockId := "blockId_example" // string | The block id, height or hash acceptable.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.BlockBlockIdGet(context.Background(), blockId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BlockBlockIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlockBlockIdGet`: BlockHeader
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.BlockBlockIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** | The block id, height or hash acceptable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlockBlockIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockHeader**](BlockHeader.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockGet

> []BlockHeader BlockGet(ctx).Last(last).Execute()

Get recent block list by paging. 30 items per page.

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
    last := int32(56) // int32 | paging flag, height of last item in last page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.BlockGet(context.Background()).Last(last).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BlockGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlockGet`: []BlockHeader
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.BlockGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **last** | **int32** | paging flag, height of last item in last page | 

### Return type

[**[]BlockHeader**](BlockHeader.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockInfoGet

> BlockchainInfo BlockInfoGet(ctx).Execute()

Get current blockchain info from full node.

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
    resp, r, err := api_client.BlockApi.BlockInfoGet(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BlockInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlockInfoGet`: BlockchainInfo
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.BlockInfoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBlockInfoGetRequest struct via the builder pattern


### Return type

[**BlockchainInfo**](BlockchainInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

