# BlockHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **int32** | Block height. | [optional] 
**BlockHash** | Pointer to **string** | Block hash. | [optional] 
**Version** | Pointer to **int32** | Block version. | [optional] 
**PrevBlockHash** | Pointer to **string** | The block hash of the previous block. | [optional] 
**MerkleRoot** | Pointer to **string** | Hex formatted block merkle root. | [optional] 
**Timestamp** | Pointer to **int64** | Block timestamp. | [optional] 
**MedianTime** | Pointer to **int64** | Block median timestamp. | [optional] 
**Reward** | Pointer to **int64** | Miner total rewards, including miner fee. | [optional] 
**Miner** | Pointer to **string** | Guessed miner name. | [optional] 
**MinerAddress** | Pointer to **string** | Miner address that received rewards. | [optional] 
**TxCount** | Pointer to **int32** | Total txs count included in the block. | [optional] 
**InputCount** | Pointer to **int32** | Total input count in the block. | [optional] 
**OutputCount** | Pointer to **int32** | Total output count in the block | [optional] 
**Size** | Pointer to **int64** | Size of the block | [optional] 
**Bits** | Pointer to **int64** | Target bits. | [optional] 
**Nonce** | Pointer to **int64** | Block nonce | [optional] 
**Coinbase** | Pointer to **string** | Hex formated coinbase info. | [optional] 

## Methods

### NewBlockHeader

`func NewBlockHeader() *BlockHeader`

NewBlockHeader instantiates a new BlockHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockHeaderWithDefaults

`func NewBlockHeaderWithDefaults() *BlockHeader`

NewBlockHeaderWithDefaults instantiates a new BlockHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *BlockHeader) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BlockHeader) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BlockHeader) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *BlockHeader) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetBlockHash

`func (o *BlockHeader) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *BlockHeader) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *BlockHeader) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *BlockHeader) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetVersion

`func (o *BlockHeader) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlockHeader) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlockHeader) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BlockHeader) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPrevBlockHash

`func (o *BlockHeader) GetPrevBlockHash() string`

GetPrevBlockHash returns the PrevBlockHash field if non-nil, zero value otherwise.

### GetPrevBlockHashOk

`func (o *BlockHeader) GetPrevBlockHashOk() (*string, bool)`

GetPrevBlockHashOk returns a tuple with the PrevBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevBlockHash

`func (o *BlockHeader) SetPrevBlockHash(v string)`

SetPrevBlockHash sets PrevBlockHash field to given value.

### HasPrevBlockHash

`func (o *BlockHeader) HasPrevBlockHash() bool`

HasPrevBlockHash returns a boolean if a field has been set.

### GetMerkleRoot

`func (o *BlockHeader) GetMerkleRoot() string`

GetMerkleRoot returns the MerkleRoot field if non-nil, zero value otherwise.

### GetMerkleRootOk

`func (o *BlockHeader) GetMerkleRootOk() (*string, bool)`

GetMerkleRootOk returns a tuple with the MerkleRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleRoot

`func (o *BlockHeader) SetMerkleRoot(v string)`

SetMerkleRoot sets MerkleRoot field to given value.

### HasMerkleRoot

`func (o *BlockHeader) HasMerkleRoot() bool`

HasMerkleRoot returns a boolean if a field has been set.

### GetTimestamp

`func (o *BlockHeader) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlockHeader) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlockHeader) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BlockHeader) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMedianTime

`func (o *BlockHeader) GetMedianTime() int64`

GetMedianTime returns the MedianTime field if non-nil, zero value otherwise.

### GetMedianTimeOk

`func (o *BlockHeader) GetMedianTimeOk() (*int64, bool)`

GetMedianTimeOk returns a tuple with the MedianTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianTime

`func (o *BlockHeader) SetMedianTime(v int64)`

SetMedianTime sets MedianTime field to given value.

### HasMedianTime

`func (o *BlockHeader) HasMedianTime() bool`

HasMedianTime returns a boolean if a field has been set.

### GetReward

`func (o *BlockHeader) GetReward() int64`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *BlockHeader) GetRewardOk() (*int64, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *BlockHeader) SetReward(v int64)`

SetReward sets Reward field to given value.

### HasReward

`func (o *BlockHeader) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetMiner

`func (o *BlockHeader) GetMiner() string`

GetMiner returns the Miner field if non-nil, zero value otherwise.

### GetMinerOk

`func (o *BlockHeader) GetMinerOk() (*string, bool)`

GetMinerOk returns a tuple with the Miner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiner

`func (o *BlockHeader) SetMiner(v string)`

SetMiner sets Miner field to given value.

### HasMiner

`func (o *BlockHeader) HasMiner() bool`

HasMiner returns a boolean if a field has been set.

### GetMinerAddress

`func (o *BlockHeader) GetMinerAddress() string`

GetMinerAddress returns the MinerAddress field if non-nil, zero value otherwise.

### GetMinerAddressOk

`func (o *BlockHeader) GetMinerAddressOk() (*string, bool)`

GetMinerAddressOk returns a tuple with the MinerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinerAddress

`func (o *BlockHeader) SetMinerAddress(v string)`

SetMinerAddress sets MinerAddress field to given value.

### HasMinerAddress

`func (o *BlockHeader) HasMinerAddress() bool`

HasMinerAddress returns a boolean if a field has been set.

### GetTxCount

`func (o *BlockHeader) GetTxCount() int32`

GetTxCount returns the TxCount field if non-nil, zero value otherwise.

### GetTxCountOk

`func (o *BlockHeader) GetTxCountOk() (*int32, bool)`

GetTxCountOk returns a tuple with the TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxCount

`func (o *BlockHeader) SetTxCount(v int32)`

SetTxCount sets TxCount field to given value.

### HasTxCount

`func (o *BlockHeader) HasTxCount() bool`

HasTxCount returns a boolean if a field has been set.

### GetInputCount

`func (o *BlockHeader) GetInputCount() int32`

GetInputCount returns the InputCount field if non-nil, zero value otherwise.

### GetInputCountOk

`func (o *BlockHeader) GetInputCountOk() (*int32, bool)`

GetInputCountOk returns a tuple with the InputCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCount

`func (o *BlockHeader) SetInputCount(v int32)`

SetInputCount sets InputCount field to given value.

### HasInputCount

`func (o *BlockHeader) HasInputCount() bool`

HasInputCount returns a boolean if a field has been set.

### GetOutputCount

`func (o *BlockHeader) GetOutputCount() int32`

GetOutputCount returns the OutputCount field if non-nil, zero value otherwise.

### GetOutputCountOk

`func (o *BlockHeader) GetOutputCountOk() (*int32, bool)`

GetOutputCountOk returns a tuple with the OutputCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCount

`func (o *BlockHeader) SetOutputCount(v int32)`

SetOutputCount sets OutputCount field to given value.

### HasOutputCount

`func (o *BlockHeader) HasOutputCount() bool`

HasOutputCount returns a boolean if a field has been set.

### GetSize

`func (o *BlockHeader) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BlockHeader) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BlockHeader) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BlockHeader) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetBits

`func (o *BlockHeader) GetBits() int64`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *BlockHeader) GetBitsOk() (*int64, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *BlockHeader) SetBits(v int64)`

SetBits sets Bits field to given value.

### HasBits

`func (o *BlockHeader) HasBits() bool`

HasBits returns a boolean if a field has been set.

### GetNonce

`func (o *BlockHeader) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *BlockHeader) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *BlockHeader) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *BlockHeader) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetCoinbase

`func (o *BlockHeader) GetCoinbase() string`

GetCoinbase returns the Coinbase field if non-nil, zero value otherwise.

### GetCoinbaseOk

`func (o *BlockHeader) GetCoinbaseOk() (*string, bool)`

GetCoinbaseOk returns a tuple with the Coinbase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbase

`func (o *BlockHeader) SetCoinbase(v string)`

SetCoinbase sets Coinbase field to given value.

### HasCoinbase

`func (o *BlockHeader) HasCoinbase() bool`

HasCoinbase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


