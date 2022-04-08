# \OutpointApi

All URIs are relative to *https://apiv2.metasv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OutpointTxidIndexGet**](OutpointApi.md#OutpointTxidIndexGet) | **Get** /outpoint/{txid}/{index} | Get tx output(outpoint for vin) spent status.
[**VinTxidDetailGet**](OutpointApi.md#VinTxidDetailGet) | **Get** /vin/{txid}/detail | Get all output point of vins in the tx with detailed output script.



## OutpointTxidIndexGet

> OutputInfo OutpointTxidIndexGet(ctx, txid, index).Execute()

Get tx output(outpoint for vin) spent status.



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
    txid := "txid_example" // string | The txid of the output
    index := int32(56) // int32 | The index of the output in the tx.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OutpointApi.OutpointTxidIndexGet(context.Background(), txid, index).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpointApi.OutpointTxidIndexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutpointTxidIndexGet`: OutputInfo
    fmt.Fprintf(os.Stdout, "Response from `OutpointApi.OutpointTxidIndexGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txid** | **string** | The txid of the output | 
**index** | **int32** | The index of the output in the tx. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutpointTxidIndexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OutputInfo**](OutputInfo.md)

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
    resp, r, err := api_client.OutpointApi.VinTxidDetailGet(context.Background(), txid).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OutpointApi.VinTxidDetailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VinTxidDetailGet`: []OutputInfoDetail
    fmt.Fprintf(os.Stdout, "Response from `OutpointApi.VinTxidDetailGet`: %v\n", resp)
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

