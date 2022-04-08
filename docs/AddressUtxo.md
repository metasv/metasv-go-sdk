# AddressUtxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flag** | Pointer to **string** | paging flag | [optional] 
**Address** | Pointer to **string** | Address of the utxo | [optional] 
**Txid** | Pointer to **string** | txid of the utxo | [optional] 
**OutIndex** | Pointer to **int32** | output index in the tx | [optional] 
**Value** | Pointer to **int64** | Value of the utxo | [optional] 
**Height** | Pointer to **int32** | Height of the utxo, -1 if not confirmed | [optional] 

## Methods

### NewAddressUtxo

`func NewAddressUtxo() *AddressUtxo`

NewAddressUtxo instantiates a new AddressUtxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressUtxoWithDefaults

`func NewAddressUtxoWithDefaults() *AddressUtxo`

NewAddressUtxoWithDefaults instantiates a new AddressUtxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlag

`func (o *AddressUtxo) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *AddressUtxo) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *AddressUtxo) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *AddressUtxo) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetAddress

`func (o *AddressUtxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressUtxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressUtxo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddressUtxo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTxid

`func (o *AddressUtxo) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *AddressUtxo) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *AddressUtxo) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *AddressUtxo) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetOutIndex

`func (o *AddressUtxo) GetOutIndex() int32`

GetOutIndex returns the OutIndex field if non-nil, zero value otherwise.

### GetOutIndexOk

`func (o *AddressUtxo) GetOutIndexOk() (*int32, bool)`

GetOutIndexOk returns a tuple with the OutIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutIndex

`func (o *AddressUtxo) SetOutIndex(v int32)`

SetOutIndex sets OutIndex field to given value.

### HasOutIndex

`func (o *AddressUtxo) HasOutIndex() bool`

HasOutIndex returns a boolean if a field has been set.

### GetValue

`func (o *AddressUtxo) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddressUtxo) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddressUtxo) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *AddressUtxo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetHeight

`func (o *AddressUtxo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *AddressUtxo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *AddressUtxo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *AddressUtxo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


