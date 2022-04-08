# XpubBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confirmed** | Pointer to **int64** | confirmed balance | [optional] 
**Unconfirmed** | Pointer to **int64** | unconfirmed balance | [optional] 

## Methods

### NewXpubBalance

`func NewXpubBalance() *XpubBalance`

NewXpubBalance instantiates a new XpubBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXpubBalanceWithDefaults

`func NewXpubBalanceWithDefaults() *XpubBalance`

NewXpubBalanceWithDefaults instantiates a new XpubBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmed

`func (o *XpubBalance) GetConfirmed() int64`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *XpubBalance) GetConfirmedOk() (*int64, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *XpubBalance) SetConfirmed(v int64)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *XpubBalance) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetUnconfirmed

`func (o *XpubBalance) GetUnconfirmed() int64`

GetUnconfirmed returns the Unconfirmed field if non-nil, zero value otherwise.

### GetUnconfirmedOk

`func (o *XpubBalance) GetUnconfirmedOk() (*int64, bool)`

GetUnconfirmedOk returns a tuple with the Unconfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfirmed

`func (o *XpubBalance) SetUnconfirmed(v int64)`

SetUnconfirmed sets Unconfirmed field to given value.

### HasUnconfirmed

`func (o *XpubBalance) HasUnconfirmed() bool`

HasUnconfirmed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


