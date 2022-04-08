# XpubLiteBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **int64** | confirmed balance plus unconfirmed balance | [optional] 

## Methods

### NewXpubLiteBalance

`func NewXpubLiteBalance() *XpubLiteBalance`

NewXpubLiteBalance instantiates a new XpubLiteBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXpubLiteBalanceWithDefaults

`func NewXpubLiteBalanceWithDefaults() *XpubLiteBalance`

NewXpubLiteBalanceWithDefaults instantiates a new XpubLiteBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *XpubLiteBalance) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *XpubLiteBalance) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *XpubLiteBalance) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *XpubLiteBalance) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


