# AddressBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | current address | [optional] 
**Confirmed** | Pointer to **int64** | confirmed balance | [optional] 
**Unconfirmed** | Pointer to **int64** | unconfirmed balance | [optional] 

## Methods

### NewAddressBalance

`func NewAddressBalance() *AddressBalance`

NewAddressBalance instantiates a new AddressBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressBalanceWithDefaults

`func NewAddressBalanceWithDefaults() *AddressBalance`

NewAddressBalanceWithDefaults instantiates a new AddressBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressBalance) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressBalance) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressBalance) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddressBalance) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetConfirmed

`func (o *AddressBalance) GetConfirmed() int64`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *AddressBalance) GetConfirmedOk() (*int64, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *AddressBalance) SetConfirmed(v int64)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *AddressBalance) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetUnconfirmed

`func (o *AddressBalance) GetUnconfirmed() int64`

GetUnconfirmed returns the Unconfirmed field if non-nil, zero value otherwise.

### GetUnconfirmedOk

`func (o *AddressBalance) GetUnconfirmedOk() (*int64, bool)`

GetUnconfirmedOk returns a tuple with the Unconfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfirmed

`func (o *AddressBalance) SetUnconfirmed(v int64)`

SetUnconfirmed sets Unconfirmed field to given value.

### HasUnconfirmed

`func (o *AddressBalance) HasUnconfirmed() bool`

HasUnconfirmed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


