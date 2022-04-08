# SensibleNftSellUtxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address string of this utxo | [optional] 
**ContractAddress** | Pointer to **string** | Address calculated from contract hash(p2ch). | [optional] 
**Txid** | Pointer to **string** | Txid for this utxo. | [optional] 
**TxIndex** | Pointer to **int32** | Output index for the Utxo. | [optional] 
**CodeHash** | Pointer to **string** | Codehash of this utxo. | [optional] 
**Genesis** | Pointer to **string** | Genesis of this utxo. | [optional] 
**TokenIndex** | Pointer to **int64** | The index of this NFT. | [optional] 
**Price** | Pointer to **int64** | the price of nft. | [optional] 
**Satoshi** | Pointer to **int64** | Bsv value of the utxo(Irrelavant to token value) | [optional] 
**SatoshiString** | Pointer to **string** | Bsv value of the utxo(In string format) | [optional] 
**Height** | Pointer to **int32** | The height of this utxo, -1 for unconfirmed utxo. | [optional] 
**IsReady** | Pointer to **bool** | Is current nft transfered into sell contract, If not ready, the following fields will be null | [optional] 
**SensibleId** | Pointer to **string** | SensibleId of the token | [optional] 
**MetaTxid** | Pointer to **string** | The metanet tx describing the nft. | [optional] 
**MetaOutputIndex** | Pointer to **int32** | Symbol of the token. | [optional] 
**TokenSupply** | Pointer to **int64** | The total supply of this NFT. | [optional] 
**Flag** | Pointer to **string** | Flag used for paging | [optional] 

## Methods

### NewSensibleNftSellUtxo

`func NewSensibleNftSellUtxo() *SensibleNftSellUtxo`

NewSensibleNftSellUtxo instantiates a new SensibleNftSellUtxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensibleNftSellUtxoWithDefaults

`func NewSensibleNftSellUtxoWithDefaults() *SensibleNftSellUtxo`

NewSensibleNftSellUtxoWithDefaults instantiates a new SensibleNftSellUtxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SensibleNftSellUtxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SensibleNftSellUtxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SensibleNftSellUtxo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SensibleNftSellUtxo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetContractAddress

`func (o *SensibleNftSellUtxo) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *SensibleNftSellUtxo) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *SensibleNftSellUtxo) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *SensibleNftSellUtxo) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetTxid

`func (o *SensibleNftSellUtxo) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *SensibleNftSellUtxo) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *SensibleNftSellUtxo) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *SensibleNftSellUtxo) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxIndex

`func (o *SensibleNftSellUtxo) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *SensibleNftSellUtxo) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *SensibleNftSellUtxo) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.

### HasTxIndex

`func (o *SensibleNftSellUtxo) HasTxIndex() bool`

HasTxIndex returns a boolean if a field has been set.

### GetCodeHash

`func (o *SensibleNftSellUtxo) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *SensibleNftSellUtxo) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *SensibleNftSellUtxo) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.

### HasCodeHash

`func (o *SensibleNftSellUtxo) HasCodeHash() bool`

HasCodeHash returns a boolean if a field has been set.

### GetGenesis

`func (o *SensibleNftSellUtxo) GetGenesis() string`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *SensibleNftSellUtxo) GetGenesisOk() (*string, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *SensibleNftSellUtxo) SetGenesis(v string)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *SensibleNftSellUtxo) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetTokenIndex

`func (o *SensibleNftSellUtxo) GetTokenIndex() int64`

GetTokenIndex returns the TokenIndex field if non-nil, zero value otherwise.

### GetTokenIndexOk

`func (o *SensibleNftSellUtxo) GetTokenIndexOk() (*int64, bool)`

GetTokenIndexOk returns a tuple with the TokenIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIndex

`func (o *SensibleNftSellUtxo) SetTokenIndex(v int64)`

SetTokenIndex sets TokenIndex field to given value.

### HasTokenIndex

`func (o *SensibleNftSellUtxo) HasTokenIndex() bool`

HasTokenIndex returns a boolean if a field has been set.

### GetPrice

`func (o *SensibleNftSellUtxo) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SensibleNftSellUtxo) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SensibleNftSellUtxo) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SensibleNftSellUtxo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSatoshi

`func (o *SensibleNftSellUtxo) GetSatoshi() int64`

GetSatoshi returns the Satoshi field if non-nil, zero value otherwise.

### GetSatoshiOk

`func (o *SensibleNftSellUtxo) GetSatoshiOk() (*int64, bool)`

GetSatoshiOk returns a tuple with the Satoshi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshi

`func (o *SensibleNftSellUtxo) SetSatoshi(v int64)`

SetSatoshi sets Satoshi field to given value.

### HasSatoshi

`func (o *SensibleNftSellUtxo) HasSatoshi() bool`

HasSatoshi returns a boolean if a field has been set.

### GetSatoshiString

`func (o *SensibleNftSellUtxo) GetSatoshiString() string`

GetSatoshiString returns the SatoshiString field if non-nil, zero value otherwise.

### GetSatoshiStringOk

`func (o *SensibleNftSellUtxo) GetSatoshiStringOk() (*string, bool)`

GetSatoshiStringOk returns a tuple with the SatoshiString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshiString

`func (o *SensibleNftSellUtxo) SetSatoshiString(v string)`

SetSatoshiString sets SatoshiString field to given value.

### HasSatoshiString

`func (o *SensibleNftSellUtxo) HasSatoshiString() bool`

HasSatoshiString returns a boolean if a field has been set.

### GetHeight

`func (o *SensibleNftSellUtxo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SensibleNftSellUtxo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SensibleNftSellUtxo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *SensibleNftSellUtxo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetIsReady

`func (o *SensibleNftSellUtxo) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *SensibleNftSellUtxo) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *SensibleNftSellUtxo) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.

### HasIsReady

`func (o *SensibleNftSellUtxo) HasIsReady() bool`

HasIsReady returns a boolean if a field has been set.

### GetSensibleId

`func (o *SensibleNftSellUtxo) GetSensibleId() string`

GetSensibleId returns the SensibleId field if non-nil, zero value otherwise.

### GetSensibleIdOk

`func (o *SensibleNftSellUtxo) GetSensibleIdOk() (*string, bool)`

GetSensibleIdOk returns a tuple with the SensibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensibleId

`func (o *SensibleNftSellUtxo) SetSensibleId(v string)`

SetSensibleId sets SensibleId field to given value.

### HasSensibleId

`func (o *SensibleNftSellUtxo) HasSensibleId() bool`

HasSensibleId returns a boolean if a field has been set.

### GetMetaTxid

`func (o *SensibleNftSellUtxo) GetMetaTxid() string`

GetMetaTxid returns the MetaTxid field if non-nil, zero value otherwise.

### GetMetaTxidOk

`func (o *SensibleNftSellUtxo) GetMetaTxidOk() (*string, bool)`

GetMetaTxidOk returns a tuple with the MetaTxid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTxid

`func (o *SensibleNftSellUtxo) SetMetaTxid(v string)`

SetMetaTxid sets MetaTxid field to given value.

### HasMetaTxid

`func (o *SensibleNftSellUtxo) HasMetaTxid() bool`

HasMetaTxid returns a boolean if a field has been set.

### GetMetaOutputIndex

`func (o *SensibleNftSellUtxo) GetMetaOutputIndex() int32`

GetMetaOutputIndex returns the MetaOutputIndex field if non-nil, zero value otherwise.

### GetMetaOutputIndexOk

`func (o *SensibleNftSellUtxo) GetMetaOutputIndexOk() (*int32, bool)`

GetMetaOutputIndexOk returns a tuple with the MetaOutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaOutputIndex

`func (o *SensibleNftSellUtxo) SetMetaOutputIndex(v int32)`

SetMetaOutputIndex sets MetaOutputIndex field to given value.

### HasMetaOutputIndex

`func (o *SensibleNftSellUtxo) HasMetaOutputIndex() bool`

HasMetaOutputIndex returns a boolean if a field has been set.

### GetTokenSupply

`func (o *SensibleNftSellUtxo) GetTokenSupply() int64`

GetTokenSupply returns the TokenSupply field if non-nil, zero value otherwise.

### GetTokenSupplyOk

`func (o *SensibleNftSellUtxo) GetTokenSupplyOk() (*int64, bool)`

GetTokenSupplyOk returns a tuple with the TokenSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSupply

`func (o *SensibleNftSellUtxo) SetTokenSupply(v int64)`

SetTokenSupply sets TokenSupply field to given value.

### HasTokenSupply

`func (o *SensibleNftSellUtxo) HasTokenSupply() bool`

HasTokenSupply returns a boolean if a field has been set.

### GetFlag

`func (o *SensibleNftSellUtxo) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *SensibleNftSellUtxo) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *SensibleNftSellUtxo) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *SensibleNftSellUtxo) HasFlag() bool`

HasFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


