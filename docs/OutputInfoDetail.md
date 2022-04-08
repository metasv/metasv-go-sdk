# OutputInfoDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Txid** | Pointer to **string** | txid that this output is in. | [optional] 
**Index** | Pointer to **int32** | index of this output in the tx. | [optional] 
**Script** | Pointer to **string** | output scrypt in hex format | [optional] 
**Address** | Pointer to **string** | parsed address of this output, empty for non standard. | [optional] 
**Value** | Pointer to **int64** | value of this output | [optional] 
**Spent** | Pointer to **bool** | this output is spent or not, true if spent | [optional] 
**SpentTxid** | Pointer to **string** | txid that spent this output | [optional] 
**SpentIndex** | Pointer to **int32** | vin index of the spent tx | [optional] 
**SpentHeight** | Pointer to **int32** | height of the spent tx(-1 if unconfirmed, 0 if unspent) | [optional] 

## Methods

### NewOutputInfoDetail

`func NewOutputInfoDetail() *OutputInfoDetail`

NewOutputInfoDetail instantiates a new OutputInfoDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputInfoDetailWithDefaults

`func NewOutputInfoDetailWithDefaults() *OutputInfoDetail`

NewOutputInfoDetailWithDefaults instantiates a new OutputInfoDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxid

`func (o *OutputInfoDetail) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *OutputInfoDetail) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *OutputInfoDetail) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *OutputInfoDetail) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetIndex

`func (o *OutputInfoDetail) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *OutputInfoDetail) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *OutputInfoDetail) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *OutputInfoDetail) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetScript

`func (o *OutputInfoDetail) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *OutputInfoDetail) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *OutputInfoDetail) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *OutputInfoDetail) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetAddress

`func (o *OutputInfoDetail) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OutputInfoDetail) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OutputInfoDetail) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OutputInfoDetail) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetValue

`func (o *OutputInfoDetail) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OutputInfoDetail) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OutputInfoDetail) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *OutputInfoDetail) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSpent

`func (o *OutputInfoDetail) GetSpent() bool`

GetSpent returns the Spent field if non-nil, zero value otherwise.

### GetSpentOk

`func (o *OutputInfoDetail) GetSpentOk() (*bool, bool)`

GetSpentOk returns a tuple with the Spent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpent

`func (o *OutputInfoDetail) SetSpent(v bool)`

SetSpent sets Spent field to given value.

### HasSpent

`func (o *OutputInfoDetail) HasSpent() bool`

HasSpent returns a boolean if a field has been set.

### GetSpentTxid

`func (o *OutputInfoDetail) GetSpentTxid() string`

GetSpentTxid returns the SpentTxid field if non-nil, zero value otherwise.

### GetSpentTxidOk

`func (o *OutputInfoDetail) GetSpentTxidOk() (*string, bool)`

GetSpentTxidOk returns a tuple with the SpentTxid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentTxid

`func (o *OutputInfoDetail) SetSpentTxid(v string)`

SetSpentTxid sets SpentTxid field to given value.

### HasSpentTxid

`func (o *OutputInfoDetail) HasSpentTxid() bool`

HasSpentTxid returns a boolean if a field has been set.

### GetSpentIndex

`func (o *OutputInfoDetail) GetSpentIndex() int32`

GetSpentIndex returns the SpentIndex field if non-nil, zero value otherwise.

### GetSpentIndexOk

`func (o *OutputInfoDetail) GetSpentIndexOk() (*int32, bool)`

GetSpentIndexOk returns a tuple with the SpentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentIndex

`func (o *OutputInfoDetail) SetSpentIndex(v int32)`

SetSpentIndex sets SpentIndex field to given value.

### HasSpentIndex

`func (o *OutputInfoDetail) HasSpentIndex() bool`

HasSpentIndex returns a boolean if a field has been set.

### GetSpentHeight

`func (o *OutputInfoDetail) GetSpentHeight() int32`

GetSpentHeight returns the SpentHeight field if non-nil, zero value otherwise.

### GetSpentHeightOk

`func (o *OutputInfoDetail) GetSpentHeightOk() (*int32, bool)`

GetSpentHeightOk returns a tuple with the SpentHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentHeight

`func (o *OutputInfoDetail) SetSpentHeight(v int32)`

SetSpentHeight sets SpentHeight field to given value.

### HasSpentHeight

`func (o *OutputInfoDetail) HasSpentHeight() bool`

HasSpentHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


