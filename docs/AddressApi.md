# \AddressApi

All URIs are relative to *https://apiv2.metasv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressAddressBalanceGet**](AddressApi.md#AddressAddressBalanceGet) | **Get** /address/{address}/balance | Get address balance by specific address.
[**AddressAddressUtxoGet**](AddressApi.md#AddressAddressUtxoGet) | **Get** /address/{address}/utxo | Get address utxos by specific address(100 per page).



## AddressAddressBalanceGet

> AddressBalance AddressAddressBalanceGet(ctx, address).Execute()

Get address balance by specific address.

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
    address := "address_example" // string | the requested address

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.AddressAddressBalanceGet(context.Background(), address).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.AddressAddressBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressAddressBalanceGet`: AddressBalance
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.AddressAddressBalanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressAddressBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressBalance**](AddressBalance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressAddressUtxoGet

> []AddressUtxo AddressAddressUtxoGet(ctx, address).Flag(flag).Execute()

Get address utxos by specific address(100 per page).

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
    address := "address_example" // string | the requested address
    flag := "flag_example" // string | The last id of the last query(Paging flag) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.AddressAddressUtxoGet(context.Background(), address).Flag(flag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.AddressAddressUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddressAddressUtxoGet`: []AddressUtxo
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.AddressAddressUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressAddressUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flag** | **string** | The last id of the last query(Paging flag) | 

### Return type

[**[]AddressUtxo**](AddressUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

