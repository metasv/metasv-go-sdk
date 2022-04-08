# SensibleNftAuctionUtxo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address string of this utxo | [optional] 
**Txid** | Pointer to **string** | Txid for this utxo. | [optional] 
**TxIndex** | Pointer to **int32** | Output index for the Utxo. | [optional] 
**BidBsvPrice** | Pointer to **int64** | bidBsvPrice | [optional] 
**BidTimestamp** | Pointer to **int64** | bidTimestamp | [optional] 
**BidderAddressPkh** | Pointer to **string** | bidderAddressPkh | [optional] 
**CodeHash** | Pointer to **string** | Codehash of this utxo. | [optional] 
**Genesis** | Pointer to **string** | Genesis of this utxo. | [optional] 
**EndTimestamp** | Pointer to **int64** | endTimestamp | [optional] 
**FeeAddressPkh** | Pointer to **string** | feeAddressPkh | [optional] 
**FeeAmount** | Pointer to **int64** | feeAmount | [optional] 
**FeeRate** | Pointer to **int32** | feeRate | [optional] 
**Height** | Pointer to **int32** | The height of this utxo, -1 for unconfirmed utxo. | [optional] 
**NftCodeHash** | Pointer to **string** | nftCodeHash | [optional] 
**NftId** | Pointer to **string** | nftId of the auctioning fnt | [optional] 
**SenderAddressPkh** | Pointer to **string** | senderAddressPkh | [optional] 
**SensibleId** | Pointer to **string** | SensibleId of the utxo | [optional] 
**StartBsvPrice** | Pointer to **int64** | startBsvPrice | [optional] 
**Satoshi** | Pointer to **int64** | Bsv value of the utxo(Irrelavant to token value) | [optional] 
**SatoshiString** | Pointer to **string** | Bsv value of the utxo(In string format) | [optional] 
**ContractAddress** | Pointer to **string** | The hash160 of the script(p2ch address) | [optional] 
**IsReady** | Pointer to **bool** | this the nft send to this contract address | [optional] 
**Flag** | Pointer to **string** | Flag used for paging | [optional] 

## Methods

### NewSensibleNftAuctionUtxo

`func NewSensibleNftAuctionUtxo() *SensibleNftAuctionUtxo`

NewSensibleNftAuctionUtxo instantiates a new SensibleNftAuctionUtxo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensibleNftAuctionUtxoWithDefaults

`func NewSensibleNftAuctionUtxoWithDefaults() *SensibleNftAuctionUtxo`

NewSensibleNftAuctionUtxoWithDefaults instantiates a new SensibleNftAuctionUtxo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SensibleNftAuctionUtxo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SensibleNftAuctionUtxo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SensibleNftAuctionUtxo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SensibleNftAuctionUtxo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTxid

`func (o *SensibleNftAuctionUtxo) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *SensibleNftAuctionUtxo) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *SensibleNftAuctionUtxo) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *SensibleNftAuctionUtxo) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetTxIndex

`func (o *SensibleNftAuctionUtxo) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *SensibleNftAuctionUtxo) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *SensibleNftAuctionUtxo) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.

### HasTxIndex

`func (o *SensibleNftAuctionUtxo) HasTxIndex() bool`

HasTxIndex returns a boolean if a field has been set.

### GetBidBsvPrice

`func (o *SensibleNftAuctionUtxo) GetBidBsvPrice() int64`

GetBidBsvPrice returns the BidBsvPrice field if non-nil, zero value otherwise.

### GetBidBsvPriceOk

`func (o *SensibleNftAuctionUtxo) GetBidBsvPriceOk() (*int64, bool)`

GetBidBsvPriceOk returns a tuple with the BidBsvPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidBsvPrice

`func (o *SensibleNftAuctionUtxo) SetBidBsvPrice(v int64)`

SetBidBsvPrice sets BidBsvPrice field to given value.

### HasBidBsvPrice

`func (o *SensibleNftAuctionUtxo) HasBidBsvPrice() bool`

HasBidBsvPrice returns a boolean if a field has been set.

### GetBidTimestamp

`func (o *SensibleNftAuctionUtxo) GetBidTimestamp() int64`

GetBidTimestamp returns the BidTimestamp field if non-nil, zero value otherwise.

### GetBidTimestampOk

`func (o *SensibleNftAuctionUtxo) GetBidTimestampOk() (*int64, bool)`

GetBidTimestampOk returns a tuple with the BidTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidTimestamp

`func (o *SensibleNftAuctionUtxo) SetBidTimestamp(v int64)`

SetBidTimestamp sets BidTimestamp field to given value.

### HasBidTimestamp

`func (o *SensibleNftAuctionUtxo) HasBidTimestamp() bool`

HasBidTimestamp returns a boolean if a field has been set.

### GetBidderAddressPkh

`func (o *SensibleNftAuctionUtxo) GetBidderAddressPkh() string`

GetBidderAddressPkh returns the BidderAddressPkh field if non-nil, zero value otherwise.

### GetBidderAddressPkhOk

`func (o *SensibleNftAuctionUtxo) GetBidderAddressPkhOk() (*string, bool)`

GetBidderAddressPkhOk returns a tuple with the BidderAddressPkh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidderAddressPkh

`func (o *SensibleNftAuctionUtxo) SetBidderAddressPkh(v string)`

SetBidderAddressPkh sets BidderAddressPkh field to given value.

### HasBidderAddressPkh

`func (o *SensibleNftAuctionUtxo) HasBidderAddressPkh() bool`

HasBidderAddressPkh returns a boolean if a field has been set.

### GetCodeHash

`func (o *SensibleNftAuctionUtxo) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *SensibleNftAuctionUtxo) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *SensibleNftAuctionUtxo) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.

### HasCodeHash

`func (o *SensibleNftAuctionUtxo) HasCodeHash() bool`

HasCodeHash returns a boolean if a field has been set.

### GetGenesis

`func (o *SensibleNftAuctionUtxo) GetGenesis() string`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *SensibleNftAuctionUtxo) GetGenesisOk() (*string, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *SensibleNftAuctionUtxo) SetGenesis(v string)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *SensibleNftAuctionUtxo) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetEndTimestamp

`func (o *SensibleNftAuctionUtxo) GetEndTimestamp() int64`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *SensibleNftAuctionUtxo) GetEndTimestampOk() (*int64, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *SensibleNftAuctionUtxo) SetEndTimestamp(v int64)`

SetEndTimestamp sets EndTimestamp field to given value.

### HasEndTimestamp

`func (o *SensibleNftAuctionUtxo) HasEndTimestamp() bool`

HasEndTimestamp returns a boolean if a field has been set.

### GetFeeAddressPkh

`func (o *SensibleNftAuctionUtxo) GetFeeAddressPkh() string`

GetFeeAddressPkh returns the FeeAddressPkh field if non-nil, zero value otherwise.

### GetFeeAddressPkhOk

`func (o *SensibleNftAuctionUtxo) GetFeeAddressPkhOk() (*string, bool)`

GetFeeAddressPkhOk returns a tuple with the FeeAddressPkh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAddressPkh

`func (o *SensibleNftAuctionUtxo) SetFeeAddressPkh(v string)`

SetFeeAddressPkh sets FeeAddressPkh field to given value.

### HasFeeAddressPkh

`func (o *SensibleNftAuctionUtxo) HasFeeAddressPkh() bool`

HasFeeAddressPkh returns a boolean if a field has been set.

### GetFeeAmount

`func (o *SensibleNftAuctionUtxo) GetFeeAmount() int64`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *SensibleNftAuctionUtxo) GetFeeAmountOk() (*int64, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *SensibleNftAuctionUtxo) SetFeeAmount(v int64)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *SensibleNftAuctionUtxo) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetFeeRate

`func (o *SensibleNftAuctionUtxo) GetFeeRate() int32`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *SensibleNftAuctionUtxo) GetFeeRateOk() (*int32, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *SensibleNftAuctionUtxo) SetFeeRate(v int32)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *SensibleNftAuctionUtxo) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetHeight

`func (o *SensibleNftAuctionUtxo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SensibleNftAuctionUtxo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SensibleNftAuctionUtxo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *SensibleNftAuctionUtxo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetNftCodeHash

`func (o *SensibleNftAuctionUtxo) GetNftCodeHash() string`

GetNftCodeHash returns the NftCodeHash field if non-nil, zero value otherwise.

### GetNftCodeHashOk

`func (o *SensibleNftAuctionUtxo) GetNftCodeHashOk() (*string, bool)`

GetNftCodeHashOk returns a tuple with the NftCodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftCodeHash

`func (o *SensibleNftAuctionUtxo) SetNftCodeHash(v string)`

SetNftCodeHash sets NftCodeHash field to given value.

### HasNftCodeHash

`func (o *SensibleNftAuctionUtxo) HasNftCodeHash() bool`

HasNftCodeHash returns a boolean if a field has been set.

### GetNftId

`func (o *SensibleNftAuctionUtxo) GetNftId() string`

GetNftId returns the NftId field if non-nil, zero value otherwise.

### GetNftIdOk

`func (o *SensibleNftAuctionUtxo) GetNftIdOk() (*string, bool)`

GetNftIdOk returns a tuple with the NftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftId

`func (o *SensibleNftAuctionUtxo) SetNftId(v string)`

SetNftId sets NftId field to given value.

### HasNftId

`func (o *SensibleNftAuctionUtxo) HasNftId() bool`

HasNftId returns a boolean if a field has been set.

### GetSenderAddressPkh

`func (o *SensibleNftAuctionUtxo) GetSenderAddressPkh() string`

GetSenderAddressPkh returns the SenderAddressPkh field if non-nil, zero value otherwise.

### GetSenderAddressPkhOk

`func (o *SensibleNftAuctionUtxo) GetSenderAddressPkhOk() (*string, bool)`

GetSenderAddressPkhOk returns a tuple with the SenderAddressPkh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddressPkh

`func (o *SensibleNftAuctionUtxo) SetSenderAddressPkh(v string)`

SetSenderAddressPkh sets SenderAddressPkh field to given value.

### HasSenderAddressPkh

`func (o *SensibleNftAuctionUtxo) HasSenderAddressPkh() bool`

HasSenderAddressPkh returns a boolean if a field has been set.

### GetSensibleId

`func (o *SensibleNftAuctionUtxo) GetSensibleId() string`

GetSensibleId returns the SensibleId field if non-nil, zero value otherwise.

### GetSensibleIdOk

`func (o *SensibleNftAuctionUtxo) GetSensibleIdOk() (*string, bool)`

GetSensibleIdOk returns a tuple with the SensibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensibleId

`func (o *SensibleNftAuctionUtxo) SetSensibleId(v string)`

SetSensibleId sets SensibleId field to given value.

### HasSensibleId

`func (o *SensibleNftAuctionUtxo) HasSensibleId() bool`

HasSensibleId returns a boolean if a field has been set.

### GetStartBsvPrice

`func (o *SensibleNftAuctionUtxo) GetStartBsvPrice() int64`

GetStartBsvPrice returns the StartBsvPrice field if non-nil, zero value otherwise.

### GetStartBsvPriceOk

`func (o *SensibleNftAuctionUtxo) GetStartBsvPriceOk() (*int64, bool)`

GetStartBsvPriceOk returns a tuple with the StartBsvPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartBsvPrice

`func (o *SensibleNftAuctionUtxo) SetStartBsvPrice(v int64)`

SetStartBsvPrice sets StartBsvPrice field to given value.

### HasStartBsvPrice

`func (o *SensibleNftAuctionUtxo) HasStartBsvPrice() bool`

HasStartBsvPrice returns a boolean if a field has been set.

### GetSatoshi

`func (o *SensibleNftAuctionUtxo) GetSatoshi() int64`

GetSatoshi returns the Satoshi field if non-nil, zero value otherwise.

### GetSatoshiOk

`func (o *SensibleNftAuctionUtxo) GetSatoshiOk() (*int64, bool)`

GetSatoshiOk returns a tuple with the Satoshi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshi

`func (o *SensibleNftAuctionUtxo) SetSatoshi(v int64)`

SetSatoshi sets Satoshi field to given value.

### HasSatoshi

`func (o *SensibleNftAuctionUtxo) HasSatoshi() bool`

HasSatoshi returns a boolean if a field has been set.

### GetSatoshiString

`func (o *SensibleNftAuctionUtxo) GetSatoshiString() string`

GetSatoshiString returns the SatoshiString field if non-nil, zero value otherwise.

### GetSatoshiStringOk

`func (o *SensibleNftAuctionUtxo) GetSatoshiStringOk() (*string, bool)`

GetSatoshiStringOk returns a tuple with the SatoshiString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatoshiString

`func (o *SensibleNftAuctionUtxo) SetSatoshiString(v string)`

SetSatoshiString sets SatoshiString field to given value.

### HasSatoshiString

`func (o *SensibleNftAuctionUtxo) HasSatoshiString() bool`

HasSatoshiString returns a boolean if a field has been set.

### GetContractAddress

`func (o *SensibleNftAuctionUtxo) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *SensibleNftAuctionUtxo) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *SensibleNftAuctionUtxo) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *SensibleNftAuctionUtxo) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetIsReady

`func (o *SensibleNftAuctionUtxo) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *SensibleNftAuctionUtxo) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *SensibleNftAuctionUtxo) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.

### HasIsReady

`func (o *SensibleNftAuctionUtxo) HasIsReady() bool`

HasIsReady returns a boolean if a field has been set.

### GetFlag

`func (o *SensibleNftAuctionUtxo) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *SensibleNftAuctionUtxo) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *SensibleNftAuctionUtxo) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *SensibleNftAuctionUtxo) HasFlag() bool`

HasFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


