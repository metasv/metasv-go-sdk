# XpubUtxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Xpub** | Pointer to **string** | xpub of the utxo | [optional] 
**Address** | Pointer to **string** | Address string of this utxo | [optional] 
**AddressType** | Pointer to **int32** | Address type, 0 for receive address, 1 for change address. path is {{addressType}}/{{addressIndex}} | [optional] 
**AddressIndex** | Pointer to **int32** | Address index. Address path in the xpub is {{addressType}}/{{addressIndex}} | [optional] 
**Txid** | Pointer to **string** | Txid for this utxo. | [optional] 
**TxIndex** | Pointer to **int32** | Output index for the Utxo. | [optional] 
**Value** | Pointer to **int64** | Satoshi value of the utxo. | [optional] 
**Height** | Pointer to **int32** | The height of this utxo, -1 for unconfirmed utxo. | [optional] 
**Flag** | Pointer to **int64** | The paging flag of utxo | [optional] 

## Methods

### NewXpubUtxo

`func NewXpubUtxo() *XpubUtxo`

NewXpubUtxo instantiates a new XpubUtxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXpubUtxoWithDefaults

`func NewXpubUtxoWithDefaults() *XpubUtxo`

NewXpubUtxoWithDefaults instantiates a new XpubUtxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXpub

`func (o *XpubUtxo) GetXpub() string`

GetXpub returns the Xpub field if non-nil, zero value otherwise.

### GetXpubOk

`func (o *XpubUtxo) GetXpubOk() (*string, bool)`

GetXpubOk returns a tuple with the Xpub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpub

`func (o *XpubUtxo) SetXpub(v string)`

SetXpub sets Xpub field to given value.

### HasXpub

`func (o *XpubUtxo) HasXpub() bool`

HasXpub returns a boolean if a field has been set.

### GetAddress

`func (o *XpubUtxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *XpubUtxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *XpubUtxo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *XpubUtxo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressType

`func (o *XpubUtxo) GetAddressType() int32`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *XpubUtxo) GetAddressTypeOk() (*int32, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *XpubUtxo) SetAddressType(v int32)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *XpubUtxo) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetAddressIndex

`func (o *XpubUtxo) GetAddressIndex() int32`

GetAddressIndex returns the AddressIndex field if non-nil, zero value otherwise.

### GetAddressIndexOk

`func (o *XpubUtxo) GetAddressIndexOk() (*int32, bool)`

GetAddressIndexOk returns a tuple with the AddressIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressIndex

`func (o *XpubUtxo) SetAddressIndex(v int32)`

SetAddressIndex sets AddressIndex field to given value.

### HasAddressIndex

`func (o *XpubUtxo) HasAddressIndex() bool`

HasAddressIndex returns a boolean if a field has been set.

### GetTxid

`func (o *XpubUtxo) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *XpubUtxo) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *XpubUtxo) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *XpubUtxo) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxIndex

`func (o *XpubUtxo) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *XpubUtxo) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *XpubUtxo) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.

### HasTxIndex

`func (o *XpubUtxo) HasTxIndex() bool`

HasTxIndex returns a boolean if a field has been set.

### GetValue

`func (o *XpubUtxo) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *XpubUtxo) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *XpubUtxo) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *XpubUtxo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetHeight

`func (o *XpubUtxo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *XpubUtxo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *XpubUtxo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *XpubUtxo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetFlag

`func (o *XpubUtxo) GetFlag() int64`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *XpubUtxo) GetFlagOk() (*int64, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *XpubUtxo) SetFlag(v int64)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *XpubUtxo) HasFlag() bool`

HasFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


