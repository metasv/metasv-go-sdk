# TxInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Input index of the tx. | [optional] 
**UtxoTxid** | Pointer to **string** | The outpoint utxo txid that this input spent | [optional] 
**UtxoIndex** | Pointer to **int32** | The outpoint utxo index that this input spent | [optional] 
**Address** | Pointer to **string** | Parsed address from pubkey(P2PKH input only). | [optional] 
**Value** | Pointer to **int64** | satoshi value of this input. | [optional] 
**UnlockScript** | Pointer to **string** | The hex of the input script. | [optional] 

## Methods

### NewTxInput

`func NewTxInput() *TxInput`

NewTxInput instantiates a new TxInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxInputWithDefaults

`func NewTxInputWithDefaults() *TxInput`

NewTxInputWithDefaults instantiates a new TxInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *TxInput) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TxInput) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TxInput) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *TxInput) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetUtxoTxid

`func (o *TxInput) GetUtxoTxid() string`

GetUtxoTxid returns the UtxoTxid field if non-nil, zero value otherwise.

### GetUtxoTxidOk

`func (o *TxInput) GetUtxoTxidOk() (*string, bool)`

GetUtxoTxidOk returns a tuple with the UtxoTxid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoTxid

`func (o *TxInput) SetUtxoTxid(v string)`

SetUtxoTxid sets UtxoTxid field to given value.

### HasUtxoTxid

`func (o *TxInput) HasUtxoTxid() bool`

HasUtxoTxid returns a boolean if a field has been set.

### GetUtxoIndex

`func (o *TxInput) GetUtxoIndex() int32`

GetUtxoIndex returns the UtxoIndex field if non-nil, zero value otherwise.

### GetUtxoIndexOk

`func (o *TxInput) GetUtxoIndexOk() (*int32, bool)`

GetUtxoIndexOk returns a tuple with the UtxoIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoIndex

`func (o *TxInput) SetUtxoIndex(v int32)`

SetUtxoIndex sets UtxoIndex field to given value.

### HasUtxoIndex

`func (o *TxInput) HasUtxoIndex() bool`

HasUtxoIndex returns a boolean if a field has been set.

### GetAddress

`func (o *TxInput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TxInput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TxInput) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TxInput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetValue

`func (o *TxInput) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TxInput) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TxInput) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *TxInput) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUnlockScript

`func (o *TxInput) GetUnlockScript() string`

GetUnlockScript returns the UnlockScript field if non-nil, zero value otherwise.

### GetUnlockScriptOk

`func (o *TxInput) GetUnlockScriptOk() (*string, bool)`

GetUnlockScriptOk returns a tuple with the UnlockScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockScript

`func (o *TxInput) SetUnlockScript(v string)`

SetUnlockScript sets UnlockScript field to given value.

### HasUnlockScript

`func (o *TxInput) HasUnlockScript() bool`

HasUnlockScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


