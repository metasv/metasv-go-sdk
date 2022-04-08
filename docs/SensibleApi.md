# \SensibleApi

All URIs are relative to *https://apiv2.metasv.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SensibleFtAddressAddressBalanceConfirmedGet**](SensibleApi.md#SensibleFtAddressAddressBalanceConfirmedGet) | **Get** /sensible/ft/address/{address}/balance/confirmed | Get all sensible token balances for specific address ignoring all unconfirmed txs.
[**SensibleFtAddressAddressBalanceGet**](SensibleApi.md#SensibleFtAddressAddressBalanceGet) | **Get** /sensible/ft/address/{address}/balance | Get all sensible token balances for specific address.
[**SensibleFtAddressAddressUtxoGet**](SensibleApi.md#SensibleFtAddressAddressUtxoGet) | **Get** /sensible/ft/address/{address}/utxo | Get all sensible token utxos for specific address.
[**SensibleNftAddressAddressCountConfirmedGet**](SensibleApi.md#SensibleNftAddressAddressCountConfirmedGet) | **Get** /sensible/nft/address/{address}/count/confirmed | Get confirmed utxo count for specific nft(ignore all unconfirmed txs).
[**SensibleNftAddressAddressSummaryGet**](SensibleApi.md#SensibleNftAddressAddressSummaryGet) | **Get** /sensible/nft/address/{address}/summary | Get nft summary(NFT count group by genesis) for address.
[**SensibleNftAddressAddressUtxoGet**](SensibleApi.md#SensibleNftAddressAddressUtxoGet) | **Get** /sensible/nft/address/{address}/utxo | Get all sensible nft token utxos for specific address.
[**SensibleNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet**](SensibleApi.md#SensibleNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet) | **Get** /sensible/nft/auction/codeHash/{codeHash}/nftId/{nftId}/utxo | Get all sensible nft token utxos by codeHash and genesisId.
[**SensibleNftGenesisCodeHashGenesisSummaryGet**](SensibleApi.md#SensibleNftGenesisCodeHashGenesisSummaryGet) | **Get** /sensible/nft/genesis/{codeHash}/{genesis}/summary | Get nft summary(count group by address) for specific codeHash and genesisId(result cached for 60s).
[**SensibleNftGenesisCodeHashGenesisUtxoGet**](SensibleApi.md#SensibleNftGenesisCodeHashGenesisUtxoGet) | **Get** /sensible/nft/genesis/{codeHash}/{genesis}/utxo | Get all sensible nft token utxos by codeHash and genesisId.
[**SensibleNftSellAddressAddressUtxoGet**](SensibleApi.md#SensibleNftSellAddressAddressUtxoGet) | **Get** /sensible/nft/sell/address/{address}/utxo | Get all sensible sell sell utxos for specific address.
[**SensibleNftSellGenesisCodeHashGenesisUtxoGet**](SensibleApi.md#SensibleNftSellGenesisCodeHashGenesisUtxoGet) | **Get** /sensible/nft/sell/genesis/{codeHash}/{genesis}/utxo | Get all sensible nft token utxos by codeHash and genesisId.
[**SensibleNftSellV2AddressAddressUtxoGet**](SensibleApi.md#SensibleNftSellV2AddressAddressUtxoGet) | **Get** /sensible/nft/sellV2/address/{address}/utxo | Get all sensible sell sell utxos for specific address.
[**SensibleNftSellV2GenesisCodeHashGenesisUtxoGet**](SensibleApi.md#SensibleNftSellV2GenesisCodeHashGenesisUtxoGet) | **Get** /sensible/nft/sellV2/genesis/{codeHash}/{genesis}/utxo | Get all sensible nft token utxos by codeHash and genesisId.



## SensibleFtAddressAddressBalanceConfirmedGet

> int64 SensibleFtAddressAddressBalanceConfirmedGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Execute()

Get all sensible token balances for specific address ignoring all unconfirmed txs.

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
    codeHash := "codeHash_example" // string | Filter by contract code hash
    genesis := "genesis_example" // string | Filter by contract genesis

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleFtAddressAddressBalanceConfirmedGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleFtAddressAddressBalanceConfirmedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleFtAddressAddressBalanceConfirmedGet`: int64
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleFtAddressAddressBalanceConfirmedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleFtAddressAddressBalanceConfirmedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 

### Return type

**int64**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensibleFtAddressAddressBalanceGet

> []SensibleFtBalance SensibleFtAddressAddressBalanceGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Execute()

Get all sensible token balances for specific address.

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
    codeHash := "codeHash_example" // string | Filter by contract code hash (optional)
    genesis := "genesis_example" // string | Filter by contract genesis (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleFtAddressAddressBalanceGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleFtAddressAddressBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleFtAddressAddressBalanceGet`: []SensibleFtBalance
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleFtAddressAddressBalanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleFtAddressAddressBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 

### Return type

[**[]SensibleFtBalance**](SensibleFtBalance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensibleFtAddressAddressUtxoGet

> []SensibleFtUtxo SensibleFtAddressAddressUtxoGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()

Get all sensible token utxos for specific address.

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
    codeHash := "codeHash_example" // string | Filter by contract code hash (optional)
    genesis := "genesis_example" // string | Filter by contract genesis (optional)
    flag := "flag_example" // string | The flag of the last query Item(Paging flag) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleFtAddressAddressUtxoGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleFtAddressAddressUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleFtAddressAddressUtxoGet`: []SensibleFtUtxo
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleFtAddressAddressUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleFtAddressAddressUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 
 **flag** | **string** | The flag of the last query Item(Paging flag) | 

### Return type

[**[]SensibleFtUtxo**](SensibleFtUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensibleNftAddressAddressCountConfirmedGet

> int32 SensibleNftAddressAddressCountConfirmedGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Execute()

Get confirmed utxo count for specific nft(ignore all unconfirmed txs).

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
    codeHash := "codeHash_example" // string | Filter by contract code hash
    genesis := "genesis_example" // string | Filter by contract genesis

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleNftAddressAddressCountConfirmedGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleNftAddressAddressCountConfirmedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleNftAddressAddressCountConfirmedGet`: int32
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleNftAddressAddressCountConfirmedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleNftAddressAddressCountConfirmedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 

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


## SensibleNftAddressAddressSummaryGet

> []SensibleNftAddressSummary SensibleNftAddressAddressSummaryGet(ctx, address).Execute()

Get nft summary(NFT count group by genesis) for address.

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
    resp, r, err := api_client.SensibleApi.SensibleNftAddressAddressSummaryGet(context.Background(), address).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleNftAddressAddressSummaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleNftAddressAddressSummaryGet`: []SensibleNftAddressSummary
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleNftAddressAddressSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleNftAddressAddressSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SensibleNftAddressSummary**](SensibleNftAddressSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensibleNftAddressAddressUtxoGet

> []SensibleNftUtxo SensibleNftAddressAddressUtxoGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()

Get all sensible nft token utxos for specific address.

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
    codeHash := "codeHash_example" // string | Filter by contract code hash (optional)
    genesis := "genesis_example" // string | Filter by contract genesis (optional)
    flag := "flag_example" // string | The flag of the last query Item(Paging flag) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleNftAddressAddressUtxoGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleNftAddressAddressUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleNftAddressAddressUtxoGet`: []SensibleNftUtxo
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleNftAddressAddressUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | the requested address | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleNftAddressAddressUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 
 **flag** | **string** | The flag of the last query Item(Paging flag) | 

### Return type

[**[]SensibleNftUtxo**](SensibleNftUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensibleNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet

> []SensibleNftAuctionUtxo SensibleNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet(ctx, codeHash, nftId).Execute()

Get all sensible nft token utxos by codeHash and genesisId.

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
    codeHash := "codeHash_example" // string | Code hash of the token.
    nftId := "nftId_example" // string | Nft id of this auction.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet(context.Background(), codeHash, nftId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet`: []SensibleNftAuctionUtxo
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleNftAuctionCodeHashCodeHashNftIdNftIdUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** | Code hash of the token. | 
**nftId** | **string** | Nft id of this auction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleNftAuctionCodeHashCodeHashNftIdNftIdUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SensibleNftAuctionUtxo**](SensibleNftAuctionUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensibleNftGenesisCodeHashGenesisSummaryGet

> []SensibleNftGenesisSummary SensibleNftGenesisCodeHashGenesisSummaryGet(ctx, codeHash, genesis).Execute()

Get nft summary(count group by address) for specific codeHash and genesisId(result cached for 60s).

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
    codeHash := "codeHash_example" // string | Code hash of the token.
    genesis := "genesis_example" // string | Contract genesis

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleNftGenesisCodeHashGenesisSummaryGet(context.Background(), codeHash, genesis).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleNftGenesisCodeHashGenesisSummaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleNftGenesisCodeHashGenesisSummaryGet`: []SensibleNftGenesisSummary
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleNftGenesisCodeHashGenesisSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** | Code hash of the token. | 
**genesis** | **string** | Contract genesis | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleNftGenesisCodeHashGenesisSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SensibleNftGenesisSummary**](SensibleNftGenesisSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensibleNftGenesisCodeHashGenesisUtxoGet

> []SensibleNftUtxo SensibleNftGenesisCodeHashGenesisUtxoGet(ctx, codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()

Get all sensible nft token utxos by codeHash and genesisId.

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
    codeHash := "codeHash_example" // string | Code hash of the token.
    genesis := "genesis_example" // string | Contract genesis
    tokenIndex := int64(789) // int64 | Find exact token Index. (optional)
    max := int64(789) // int64 | Token index not bigger than this(include this). (optional)
    min := int64(789) // int64 | Token index not less than this(include this). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleNftGenesisCodeHashGenesisUtxoGet(context.Background(), codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleNftGenesisCodeHashGenesisUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleNftGenesisCodeHashGenesisUtxoGet`: []SensibleNftUtxo
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleNftGenesisCodeHashGenesisUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** | Code hash of the token. | 
**genesis** | **string** | Contract genesis | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleNftGenesisCodeHashGenesisUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenIndex** | **int64** | Find exact token Index. | 
 **max** | **int64** | Token index not bigger than this(include this). | 
 **min** | **int64** | Token index not less than this(include this). | 

### Return type

[**[]SensibleNftUtxo**](SensibleNftUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensibleNftSellAddressAddressUtxoGet

> []SensibleNftSellUtxo SensibleNftSellAddressAddressUtxoGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()

Get all sensible sell sell utxos for specific address.

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
    address := "address_example" // string | Owner address.
    codeHash := "codeHash_example" // string | Filter by contract code hash (optional)
    genesis := "genesis_example" // string | Filter by contract genesis (optional)
    flag := "flag_example" // string | The flag of the last query Item(Paging flag) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleNftSellAddressAddressUtxoGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleNftSellAddressAddressUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleNftSellAddressAddressUtxoGet`: []SensibleNftSellUtxo
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleNftSellAddressAddressUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Owner address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleNftSellAddressAddressUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 
 **flag** | **string** | The flag of the last query Item(Paging flag) | 

### Return type

[**[]SensibleNftSellUtxo**](SensibleNftSellUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensibleNftSellGenesisCodeHashGenesisUtxoGet

> []SensibleNftSellUtxo SensibleNftSellGenesisCodeHashGenesisUtxoGet(ctx, codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()

Get all sensible nft token utxos by codeHash and genesisId.

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
    codeHash := "codeHash_example" // string | Code hash of the token.
    genesis := "genesis_example" // string | Contract genesis
    tokenIndex := int64(789) // int64 | Find exact token Index. (optional)
    max := int64(789) // int64 | Token index not bigger than this(include this). (optional)
    min := int64(789) // int64 | Token index not less than this(include this). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleNftSellGenesisCodeHashGenesisUtxoGet(context.Background(), codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleNftSellGenesisCodeHashGenesisUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleNftSellGenesisCodeHashGenesisUtxoGet`: []SensibleNftSellUtxo
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleNftSellGenesisCodeHashGenesisUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** | Code hash of the token. | 
**genesis** | **string** | Contract genesis | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleNftSellGenesisCodeHashGenesisUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenIndex** | **int64** | Find exact token Index. | 
 **max** | **int64** | Token index not bigger than this(include this). | 
 **min** | **int64** | Token index not less than this(include this). | 

### Return type

[**[]SensibleNftSellUtxo**](SensibleNftSellUtxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensibleNftSellV2AddressAddressUtxoGet

> []SensibleNftSellV2Utxo SensibleNftSellV2AddressAddressUtxoGet(ctx, address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()

Get all sensible sell sell utxos for specific address.

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
    address := "address_example" // string | Owner address.
    codeHash := "codeHash_example" // string | Filter by contract code hash (optional)
    genesis := "genesis_example" // string | Filter by contract genesis (optional)
    flag := "flag_example" // string | The flag of the last query Item(Paging flag) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleNftSellV2AddressAddressUtxoGet(context.Background(), address).CodeHash(codeHash).Genesis(genesis).Flag(flag).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleNftSellV2AddressAddressUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleNftSellV2AddressAddressUtxoGet`: []SensibleNftSellV2Utxo
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleNftSellV2AddressAddressUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Owner address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleNftSellV2AddressAddressUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeHash** | **string** | Filter by contract code hash | 
 **genesis** | **string** | Filter by contract genesis | 
 **flag** | **string** | The flag of the last query Item(Paging flag) | 

### Return type

[**[]SensibleNftSellV2Utxo**](SensibleNftSellV2Utxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensibleNftSellV2GenesisCodeHashGenesisUtxoGet

> []SensibleNftSellV2Utxo SensibleNftSellV2GenesisCodeHashGenesisUtxoGet(ctx, codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()

Get all sensible nft token utxos by codeHash and genesisId.

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
    codeHash := "codeHash_example" // string | Code hash of the token.
    genesis := "genesis_example" // string | Contract genesis
    tokenIndex := int64(789) // int64 | Find exact token Index. (optional)
    max := int64(789) // int64 | Token index not bigger than this(include this). (optional)
    min := int64(789) // int64 | Token index not less than this(include this). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensibleApi.SensibleNftSellV2GenesisCodeHashGenesisUtxoGet(context.Background(), codeHash, genesis).TokenIndex(tokenIndex).Max(max).Min(min).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SensibleApi.SensibleNftSellV2GenesisCodeHashGenesisUtxoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensibleNftSellV2GenesisCodeHashGenesisUtxoGet`: []SensibleNftSellV2Utxo
    fmt.Fprintf(os.Stdout, "Response from `SensibleApi.SensibleNftSellV2GenesisCodeHashGenesisUtxoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**codeHash** | **string** | Code hash of the token. | 
**genesis** | **string** | Contract genesis | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensibleNftSellV2GenesisCodeHashGenesisUtxoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tokenIndex** | **int64** | Find exact token Index. | 
 **max** | **int64** | Token index not bigger than this(include this). | 
 **min** | **int64** | Token index not less than this(include this). | 

### Return type

[**[]SensibleNftSellV2Utxo**](SensibleNftSellV2Utxo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

