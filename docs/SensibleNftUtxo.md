# SensibleNftUtxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address string of this utxo | [optional] 
**Txid** | Pointer to **string** | Txid for this utxo. | [optional] 
**TxIndex** | Pointer to **int32** | Output index for the Utxo. | [optional] 
**CodeHash** | Pointer to **string** | Codehash of this utxo. | [optional] 
**Genesis** | Pointer to **string** | Genesis of this utxo. | [optional] 
**SensibleId** | Pointer to **string** | SensibleId of the token | [optional] 
**Height** | Pointer to **int32** | The height of this utxo, -1 for unconfirmed utxo. | [optional] 
**MetaTxid** | Pointer to **string** | The metanet tx describing the nft. | [optional] 
**MetaOutputIndex** | Pointer to **int32** | Symbol of the token. | [optional] 
**TokenSupply** | Pointer to **int64** | The total supply of this NFT. | [optional] 
**TokenIndex** | Pointer to **int64** | The index of this NFT. | [optional] 
**Satoshi** | Pointer to **int64** | Bsv value of the utxo(Irrelavant to token value) | [optional] 
**SatoshiString** | Pointer to **string** | Bsv value of the utxo(In string format) | [optional] 
**Flag** | Pointer to **string** | Flag used for paging | [optional] 

## Methods

### NewSensibleNftUtxo

`func NewSensibleNftUtxo() *SensibleNftUtxo`

NewSensibleNftUtxo instantiates a new SensibleNftUtxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensibleNftUtxoWithDefaults

`func NewSensibleNftUtxoWithDefaults() *SensibleNftUtxo`

NewSensibleNftUtxoWithDefaults instantiates a new SensibleNftUtxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SensibleNftUtxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SensibleNftUtxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SensibleNftUtxo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SensibleNftUtxo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTxid

`func (o *SensibleNftUtxo) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *SensibleNftUtxo) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *SensibleNftUtxo) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *SensibleNftUtxo) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxIndex

`func (o *SensibleNftUtxo) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *SensibleNftUtxo) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *SensibleNftUtxo) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.

### HasTxIndex

`func (o *SensibleNftUtxo) HasTxIndex() bool`

HasTxIndex returns a boolean if a field has been set.

### GetCodeHash

`func (o *SensibleNftUtxo) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *SensibleNftUtxo) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *SensibleNftUtxo) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.

### HasCodeHash

`func (o *SensibleNftUtxo) HasCodeHash() bool`

HasCodeHash returns a boolean if a field has been set.

### GetGenesis

`func (o *SensibleNftUtxo) GetGenesis() string`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *SensibleNftUtxo) GetGenesisOk() (*string, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *SensibleNftUtxo) SetGenesis(v string)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *SensibleNftUtxo) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetSensibleId

`func (o *SensibleNftUtxo) GetSensibleId() string`

GetSensibleId returns the SensibleId field if non-nil, zero value otherwise.

### GetSensibleIdOk

`func (o *SensibleNftUtxo) GetSensibleIdOk() (*string, bool)`

GetSensibleIdOk returns a tuple with the SensibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensibleId

`func (o *SensibleNftUtxo) SetSensibleId(v string)`

SetSensibleId sets SensibleId field to given value.

### HasSensibleId

`func (o *SensibleNftUtxo) HasSensibleId() bool`

HasSensibleId returns a boolean if a field has been set.

### GetHeight

`func (o *SensibleNftUtxo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SensibleNftUtxo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SensibleNftUtxo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *SensibleNftUtxo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMetaTxid

`func (o *SensibleNftUtxo) GetMetaTxid() string`

GetMetaTxid returns the MetaTxid field if non-nil, zero value otherwise.

### GetMetaTxidOk

`func (o *SensibleNftUtxo) GetMetaTxidOk() (*string, bool)`

GetMetaTxidOk returns a tuple with the MetaTxid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTxid

`func (o *SensibleNftUtxo) SetMetaTxid(v string)`

SetMetaTxid sets MetaTxid field to given value.

### HasMetaTxid

`func (o *SensibleNftUtxo) HasMetaTxid() bool`

HasMetaTxid returns a boolean if a field has been set.

### GetMetaOutputIndex

`func (o *SensibleNftUtxo) GetMetaOutputIndex() int32`

GetMetaOutputIndex returns the MetaOutputIndex field if non-nil, zero value otherwise.

### GetMetaOutputIndexOk

`func (o *SensibleNftUtxo) GetMetaOutputIndexOk() (*int32, bool)`

GetMetaOutputIndexOk returns a tuple with the MetaOutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaOutputIndex

`func (o *SensibleNftUtxo) SetMetaOutputIndex(v int32)`

SetMetaOutputIndex sets MetaOutputIndex field to given value.

### HasMetaOutputIndex

`func (o *SensibleNftUtxo) HasMetaOutputIndex() bool`

HasMetaOutputIndex returns a boolean if a field has been set.

### GetTokenSupply

`func (o *SensibleNftUtxo) GetTokenSupply() int64`

GetTokenSupply returns the TokenSupply field if non-nil, zero value otherwise.

### GetTokenSupplyOk

`func (o *SensibleNftUtxo) GetTokenSupplyOk() (*int64, bool)`

GetTokenSupplyOk returns a tuple with the TokenSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSupply

`func (o *SensibleNftUtxo) SetTokenSupply(v int64)`

SetTokenSupply sets TokenSupply field to given value.

### HasTokenSupply

`func (o *SensibleNftUtxo) HasTokenSupply() bool`

HasTokenSupply returns a boolean if a field has been set.

### GetTokenIndex

`func (o *SensibleNftUtxo) GetTokenIndex() int64`

GetTokenIndex returns the TokenIndex field if non-nil, zero value otherwise.

### GetTokenIndexOk

`func (o *SensibleNftUtxo) GetTokenIndexOk() (*int64, bool)`

GetTokenIndexOk returns a tuple with the TokenIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIndex

`func (o *SensibleNftUtxo) SetTokenIndex(v int64)`

SetTokenIndex sets TokenIndex field to given value.

### HasTokenIndex

`func (o *SensibleNftUtxo) HasTokenIndex() bool`

HasTokenIndex returns a boolean if a field has been set.

### GetSatoshi

`func (o *SensibleNftUtxo) GetSatoshi() int64`

GetSatoshi returns the Satoshi field if non-nil, zero value otherwise.

### GetSatoshiOk

`func (o *SensibleNftUtxo) GetSatoshiOk() (*int64, bool)`

GetSatoshiOk returns a tuple with the Satoshi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshi

`func (o *SensibleNftUtxo) SetSatoshi(v int64)`

SetSatoshi sets Satoshi field to given value.

### HasSatoshi

`func (o *SensibleNftUtxo) HasSatoshi() bool`

HasSatoshi returns a boolean if a field has been set.

### GetSatoshiString

`func (o *SensibleNftUtxo) GetSatoshiString() string`

GetSatoshiString returns the SatoshiString field if non-nil, zero value otherwise.

### GetSatoshiStringOk

`func (o *SensibleNftUtxo) GetSatoshiStringOk() (*string, bool)`

GetSatoshiStringOk returns a tuple with the SatoshiString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshiString

`func (o *SensibleNftUtxo) SetSatoshiString(v string)`

SetSatoshiString sets SatoshiString field to given value.

### HasSatoshiString

`func (o *SensibleNftUtxo) HasSatoshiString() bool`

HasSatoshiString returns a boolean if a field has been set.

### GetFlag

`func (o *SensibleNftUtxo) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *SensibleNftUtxo) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *SensibleNftUtxo) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *SensibleNftUtxo) HasFlag() bool`

HasFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


