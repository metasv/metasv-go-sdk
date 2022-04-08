# UtxoPickRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** | The address list to pick utxo from | [optional] 
**Amount** | Pointer to **int64** | The total amount you want to pick | [optional] 

## Methods

### NewUtxoPickRequest

`func NewUtxoPickRequest() *UtxoPickRequest`

NewUtxoPickRequest instantiates a new UtxoPickRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtxoPickRequestWithDefaults

`func NewUtxoPickRequestWithDefaults() *UtxoPickRequest`

NewUtxoPickRequestWithDefaults instantiates a new UtxoPickRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *UtxoPickRequest) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *UtxoPickRequest) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *UtxoPickRequest) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *UtxoPickRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAmount

`func (o *UtxoPickRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UtxoPickRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UtxoPickRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UtxoPickRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


