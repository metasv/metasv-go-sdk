# SensibleNftSellV2Utxo

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
**FeeAddressPkh** | Pointer to **string** | The address to receive fees. | [optional] 
**FeeRate** | Pointer to **int32** | feeRate | [optional] 
**NftId** | Pointer to **string** | nftId | [optional] 
**SellerAddressPkh** | Pointer to **string** | The address pkh of seller | [optional] 
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

### NewSensibleNftSellV2Utxo

`func NewSensibleNftSellV2Utxo() *SensibleNftSellV2Utxo`

NewSensibleNftSellV2Utxo instantiates a new SensibleNftSellV2Utxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensibleNftSellV2UtxoWithDefaults

`func NewSensibleNftSellV2UtxoWithDefaults() *SensibleNftSellV2Utxo`

NewSensibleNftSellV2UtxoWithDefaults instantiates a new SensibleNftSellV2Utxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SensibleNftSellV2Utxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SensibleNftSellV2Utxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SensibleNftSellV2Utxo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SensibleNftSellV2Utxo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetContractAddress

`func (o *SensibleNftSellV2Utxo) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *SensibleNftSellV2Utxo) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *SensibleNftSellV2Utxo) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *SensibleNftSellV2Utxo) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetTxid

`func (o *SensibleNftSellV2Utxo) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *SensibleNftSellV2Utxo) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *SensibleNftSellV2Utxo) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *SensibleNftSellV2Utxo) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxIndex

`func (o *SensibleNftSellV2Utxo) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *SensibleNftSellV2Utxo) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *SensibleNftSellV2Utxo) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.

### HasTxIndex

`func (o *SensibleNftSellV2Utxo) HasTxIndex() bool`

HasTxIndex returns a boolean if a field has been set.

### GetCodeHash

`func (o *SensibleNftSellV2Utxo) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *SensibleNftSellV2Utxo) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *SensibleNftSellV2Utxo) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.

### HasCodeHash

`func (o *SensibleNftSellV2Utxo) HasCodeHash() bool`

HasCodeHash returns a boolean if a field has been set.

### GetGenesis

`func (o *SensibleNftSellV2Utxo) GetGenesis() string`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *SensibleNftSellV2Utxo) GetGenesisOk() (*string, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *SensibleNftSellV2Utxo) SetGenesis(v string)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *SensibleNftSellV2Utxo) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetTokenIndex

`func (o *SensibleNftSellV2Utxo) GetTokenIndex() int64`

GetTokenIndex returns the TokenIndex field if non-nil, zero value otherwise.

### GetTokenIndexOk

`func (o *SensibleNftSellV2Utxo) GetTokenIndexOk() (*int64, bool)`

GetTokenIndexOk returns a tuple with the TokenIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIndex

`func (o *SensibleNftSellV2Utxo) SetTokenIndex(v int64)`

SetTokenIndex sets TokenIndex field to given value.

### HasTokenIndex

`func (o *SensibleNftSellV2Utxo) HasTokenIndex() bool`

HasTokenIndex returns a boolean if a field has been set.

### GetPrice

`func (o *SensibleNftSellV2Utxo) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SensibleNftSellV2Utxo) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SensibleNftSellV2Utxo) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SensibleNftSellV2Utxo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetFeeAddressPkh

`func (o *SensibleNftSellV2Utxo) GetFeeAddressPkh() string`

GetFeeAddressPkh returns the FeeAddressPkh field if non-nil, zero value otherwise.

### GetFeeAddressPkhOk

`func (o *SensibleNftSellV2Utxo) GetFeeAddressPkhOk() (*string, bool)`

GetFeeAddressPkhOk returns a tuple with the FeeAddressPkh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAddressPkh

`func (o *SensibleNftSellV2Utxo) SetFeeAddressPkh(v string)`

SetFeeAddressPkh sets FeeAddressPkh field to given value.

### HasFeeAddressPkh

`func (o *SensibleNftSellV2Utxo) HasFeeAddressPkh() bool`

HasFeeAddressPkh returns a boolean if a field has been set.

### GetFeeRate

`func (o *SensibleNftSellV2Utxo) GetFeeRate() int32`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *SensibleNftSellV2Utxo) GetFeeRateOk() (*int32, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *SensibleNftSellV2Utxo) SetFeeRate(v int32)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *SensibleNftSellV2Utxo) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetNftId

`func (o *SensibleNftSellV2Utxo) GetNftId() string`

GetNftId returns the NftId field if non-nil, zero value otherwise.

### GetNftIdOk

`func (o *SensibleNftSellV2Utxo) GetNftIdOk() (*string, bool)`

GetNftIdOk returns a tuple with the NftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftId

`func (o *SensibleNftSellV2Utxo) SetNftId(v string)`

SetNftId sets NftId field to given value.

### HasNftId

`func (o *SensibleNftSellV2Utxo) HasNftId() bool`

HasNftId returns a boolean if a field has been set.

### GetSellerAddressPkh

`func (o *SensibleNftSellV2Utxo) GetSellerAddressPkh() string`

GetSellerAddressPkh returns the SellerAddressPkh field if non-nil, zero value otherwise.

### GetSellerAddressPkhOk

`func (o *SensibleNftSellV2Utxo) GetSellerAddressPkhOk() (*string, bool)`

GetSellerAddressPkhOk returns a tuple with the SellerAddressPkh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerAddressPkh

`func (o *SensibleNftSellV2Utxo) SetSellerAddressPkh(v string)`

SetSellerAddressPkh sets SellerAddressPkh field to given value.

### HasSellerAddressPkh

`func (o *SensibleNftSellV2Utxo) HasSellerAddressPkh() bool`

HasSellerAddressPkh returns a boolean if a field has been set.

### GetSatoshi

`func (o *SensibleNftSellV2Utxo) GetSatoshi() int64`

GetSatoshi returns the Satoshi field if non-nil, zero value otherwise.

### GetSatoshiOk

`func (o *SensibleNftSellV2Utxo) GetSatoshiOk() (*int64, bool)`

GetSatoshiOk returns a tuple with the Satoshi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshi

`func (o *SensibleNftSellV2Utxo) SetSatoshi(v int64)`

SetSatoshi sets Satoshi field to given value.

### HasSatoshi

`func (o *SensibleNftSellV2Utxo) HasSatoshi() bool`

HasSatoshi returns a boolean if a field has been set.

### GetSatoshiString

`func (o *SensibleNftSellV2Utxo) GetSatoshiString() string`

GetSatoshiString returns the SatoshiString field if non-nil, zero value otherwise.

### GetSatoshiStringOk

`func (o *SensibleNftSellV2Utxo) GetSatoshiStringOk() (*string, bool)`

GetSatoshiStringOk returns a tuple with the SatoshiString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshiString

`func (o *SensibleNftSellV2Utxo) SetSatoshiString(v string)`

SetSatoshiString sets SatoshiString field to given value.

### HasSatoshiString

`func (o *SensibleNftSellV2Utxo) HasSatoshiString() bool`

HasSatoshiString returns a boolean if a field has been set.

### GetHeight

`func (o *SensibleNftSellV2Utxo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SensibleNftSellV2Utxo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SensibleNftSellV2Utxo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *SensibleNftSellV2Utxo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetIsReady

`func (o *SensibleNftSellV2Utxo) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *SensibleNftSellV2Utxo) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *SensibleNftSellV2Utxo) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.

### HasIsReady

`func (o *SensibleNftSellV2Utxo) HasIsReady() bool`

HasIsReady returns a boolean if a field has been set.

### GetSensibleId

`func (o *SensibleNftSellV2Utxo) GetSensibleId() string`

GetSensibleId returns the SensibleId field if non-nil, zero value otherwise.

### GetSensibleIdOk

`func (o *SensibleNftSellV2Utxo) GetSensibleIdOk() (*string, bool)`

GetSensibleIdOk returns a tuple with the SensibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensibleId

`func (o *SensibleNftSellV2Utxo) SetSensibleId(v string)`

SetSensibleId sets SensibleId field to given value.

### HasSensibleId

`func (o *SensibleNftSellV2Utxo) HasSensibleId() bool`

HasSensibleId returns a boolean if a field has been set.

### GetMetaTxid

`func (o *SensibleNftSellV2Utxo) GetMetaTxid() string`

GetMetaTxid returns the MetaTxid field if non-nil, zero value otherwise.

### GetMetaTxidOk

`func (o *SensibleNftSellV2Utxo) GetMetaTxidOk() (*string, bool)`

GetMetaTxidOk returns a tuple with the MetaTxid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTxid

`func (o *SensibleNftSellV2Utxo) SetMetaTxid(v string)`

SetMetaTxid sets MetaTxid field to given value.

### HasMetaTxid

`func (o *SensibleNftSellV2Utxo) HasMetaTxid() bool`

HasMetaTxid returns a boolean if a field has been set.

### GetMetaOutputIndex

`func (o *SensibleNftSellV2Utxo) GetMetaOutputIndex() int32`

GetMetaOutputIndex returns the MetaOutputIndex field if non-nil, zero value otherwise.

### GetMetaOutputIndexOk

`func (o *SensibleNftSellV2Utxo) GetMetaOutputIndexOk() (*int32, bool)`

GetMetaOutputIndexOk returns a tuple with the MetaOutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaOutputIndex

`func (o *SensibleNftSellV2Utxo) SetMetaOutputIndex(v int32)`

SetMetaOutputIndex sets MetaOutputIndex field to given value.

### HasMetaOutputIndex

`func (o *SensibleNftSellV2Utxo) HasMetaOutputIndex() bool`

HasMetaOutputIndex returns a boolean if a field has been set.

### GetTokenSupply

`func (o *SensibleNftSellV2Utxo) GetTokenSupply() int64`

GetTokenSupply returns the TokenSupply field if non-nil, zero value otherwise.

### GetTokenSupplyOk

`func (o *SensibleNftSellV2Utxo) GetTokenSupplyOk() (*int64, bool)`

GetTokenSupplyOk returns a tuple with the TokenSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSupply

`func (o *SensibleNftSellV2Utxo) SetTokenSupply(v int64)`

SetTokenSupply sets TokenSupply field to given value.

### HasTokenSupply

`func (o *SensibleNftSellV2Utxo) HasTokenSupply() bool`

HasTokenSupply returns a boolean if a field has been set.

### GetFlag

`func (o *SensibleNftSellV2Utxo) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *SensibleNftSellV2Utxo) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *SensibleNftSellV2Utxo) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *SensibleNftSellV2Utxo) HasFlag() bool`

HasFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


