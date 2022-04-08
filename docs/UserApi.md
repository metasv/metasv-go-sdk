# \UserApi

All URIs are relative to *https://apiv2.metasv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserClientPubkeyGet**](UserApi.md#UserClientPubkeyGet) | **Get** /user/clientPubkey | See https://github.com/metasv/metasv-client-signature for details. Get all trusted public keys.
[**UserClientPubkeyPost**](UserApi.md#UserClientPubkeyPost) | **Post** /user/clientPubkey | See https://github.com/metasv/metasv-client-signature for details. Register new client public key.
[**UserClientPubkeyPubkeyDelete**](UserApi.md#UserClientPubkeyPubkeyDelete) | **Delete** /user/clientPubkey/{pubkey} | Delete a registered client public key.
[**UserTrafficGet**](UserApi.md#UserTrafficGet) | **Get** /user/traffic | Get traffic details for every hour(valid in six months).
[**UserTrafficSumGet**](UserApi.md#UserTrafficSumGet) | **Get** /user/traffic/sum | Sum all traffic in the current month.



## UserClientPubkeyGet

> []string UserClientPubkeyGet(ctx).Execute()

See https://github.com/metasv/metasv-client-signature for details. Get all trusted public keys.

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
    resp, r, err := api_client.UserApi.UserClientPubkeyGet(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserClientPubkeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserClientPubkeyGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserClientPubkeyGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserClientPubkeyGetRequest struct via the builder pattern


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


## UserClientPubkeyPost

> ClientPubkeyResult UserClientPubkeyPost(ctx).ClientPubkeyRequest(clientPubkeyRequest).Execute()

See https://github.com/metasv/metasv-client-signature for details. Register new client public key.

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
    clientPubkeyRequest := *openapiclient.NewClientPubkeyRequest() // ClientPubkeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.UserClientPubkeyPost(context.Background()).ClientPubkeyRequest(clientPubkeyRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserClientPubkeyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserClientPubkeyPost`: ClientPubkeyResult
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserClientPubkeyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserClientPubkeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientPubkeyRequest** | [**ClientPubkeyRequest**](ClientPubkeyRequest.md) |  | 

### Return type

[**ClientPubkeyResult**](ClientPubkeyResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserClientPubkeyPubkeyDelete

> ClientPubkeyResult UserClientPubkeyPubkeyDelete(ctx, pubkey).Execute()

Delete a registered client public key.



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
    pubkey := "pubkey_example" // string | the pubkey to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.UserClientPubkeyPubkeyDelete(context.Background(), pubkey).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserClientPubkeyPubkeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserClientPubkeyPubkeyDelete`: ClientPubkeyResult
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserClientPubkeyPubkeyDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | the pubkey to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserClientPubkeyPubkeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientPubkeyResult**](ClientPubkeyResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTrafficGet

> []UserHourlyTraffic UserTrafficGet(ctx).StartTime(startTime).EndTime(endTime).Execute()

Get traffic details for every hour(valid in six months).



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
    startTime := int64(789) // int64 | define start query time(unix timestamp) (optional)
    endTime := int64(789) // int64 | define end query time(unix timestamp) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.UserTrafficGet(context.Background()).StartTime(startTime).EndTime(endTime).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserTrafficGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserTrafficGet`: []UserHourlyTraffic
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserTrafficGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserTrafficGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | define start query time(unix timestamp) | 
 **endTime** | **int64** | define end query time(unix timestamp) | 

### Return type

[**[]UserHourlyTraffic**](UserHourlyTraffic.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTrafficSumGet

> UserTrafficSum UserTrafficSumGet(ctx).Execute()

Sum all traffic in the current month.



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
    resp, r, err := api_client.UserApi.UserTrafficSumGet(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserTrafficSumGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserTrafficSumGet`: UserTrafficSum
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserTrafficSumGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserTrafficSumGetRequest struct via the builder pattern


### Return type

[**UserTrafficSum**](UserTrafficSum.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

