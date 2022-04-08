# TxOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Output index of the tx. | [optional] 
**Value** | Pointer to **int64** | Bitcoin Value in this output. | [optional] 
**Address** | Pointer to **string** | Parsed address from output | [optional] 
**LockScript** | Pointer to **string** | Hex formatted lockScript | [optional] 

## Methods

### NewTxOutput

`func NewTxOutput() *TxOutput`

NewTxOutput instantiates a new TxOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxOutputWithDefaults

`func NewTxOutputWithDefaults() *TxOutput`

NewTxOutputWithDefaults instantiates a new TxOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *TxOutput) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TxOutput) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TxOutput) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *TxOutput) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetValue

`func (o *TxOutput) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TxOutput) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TxOutput) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *TxOutput) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAddress

`func (o *TxOutput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TxOutput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TxOutput) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TxOutput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLockScript

`func (o *TxOutput) GetLockScript() string`

GetLockScript returns the LockScript field if non-nil, zero value otherwise.

### GetLockScriptOk

`func (o *TxOutput) GetLockScriptOk() (*string, bool)`

GetLockScriptOk returns a tuple with the LockScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockScript

`func (o *TxOutput) SetLockScript(v string)`

SetLockScript sets LockScript field to given value.

### HasLockScript

`func (o *TxOutput) HasLockScript() bool`

HasLockScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


