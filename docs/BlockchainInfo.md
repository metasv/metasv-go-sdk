# BlockchainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** | mainnet or testnet | [optional] 
**Blocks** | Pointer to **int32** | current block count | [optional] 
**Headers** | Pointer to **int32** | current block header count | [optional] 
**BestBlockHash** | Pointer to **string** | current highest blockhash | [optional] 
**Difficulty** | Pointer to **string** | decimal string for current difficulty | [optional] 
**MedianTime** | Pointer to **int64** | current median time | [optional] 
**Chainwork** | Pointer to **string** | current pow | [optional] 

## Methods

### NewBlockchainInfo

`func NewBlockchainInfo() *BlockchainInfo`

NewBlockchainInfo instantiates a new BlockchainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainInfoWithDefaults

`func NewBlockchainInfoWithDefaults() *BlockchainInfo`

NewBlockchainInfoWithDefaults instantiates a new BlockchainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *BlockchainInfo) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *BlockchainInfo) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *BlockchainInfo) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *BlockchainInfo) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetBlocks

`func (o *BlockchainInfo) GetBlocks() int32`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *BlockchainInfo) GetBlocksOk() (*int32, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *BlockchainInfo) SetBlocks(v int32)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *BlockchainInfo) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetHeaders

`func (o *BlockchainInfo) GetHeaders() int32`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BlockchainInfo) GetHeadersOk() (*int32, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BlockchainInfo) SetHeaders(v int32)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BlockchainInfo) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBestBlockHash

`func (o *BlockchainInfo) GetBestBlockHash() string`

GetBestBlockHash returns the BestBlockHash field if non-nil, zero value otherwise.

### GetBestBlockHashOk

`func (o *BlockchainInfo) GetBestBlockHashOk() (*string, bool)`

GetBestBlockHashOk returns a tuple with the BestBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestBlockHash

`func (o *BlockchainInfo) SetBestBlockHash(v string)`

SetBestBlockHash sets BestBlockHash field to given value.

### HasBestBlockHash

`func (o *BlockchainInfo) HasBestBlockHash() bool`

HasBestBlockHash returns a boolean if a field has been set.

### GetDifficulty

`func (o *BlockchainInfo) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *BlockchainInfo) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *BlockchainInfo) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *BlockchainInfo) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetMedianTime

`func (o *BlockchainInfo) GetMedianTime() int64`

GetMedianTime returns the MedianTime field if non-nil, zero value otherwise.

### GetMedianTimeOk

`func (o *BlockchainInfo) GetMedianTimeOk() (*int64, bool)`

GetMedianTimeOk returns a tuple with the MedianTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianTime

`func (o *BlockchainInfo) SetMedianTime(v int64)`

SetMedianTime sets MedianTime field to given value.

### HasMedianTime

`func (o *BlockchainInfo) HasMedianTime() bool`

HasMedianTime returns a boolean if a field has been set.

### GetChainwork

`func (o *BlockchainInfo) GetChainwork() string`

GetChainwork returns the Chainwork field if non-nil, zero value otherwise.

### GetChainworkOk

`func (o *BlockchainInfo) GetChainworkOk() (*string, bool)`

GetChainworkOk returns a tuple with the Chainwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainwork

`func (o *BlockchainInfo) SetChainwork(v string)`

SetChainwork sets Chainwork field to given value.

### HasChainwork

`func (o *BlockchainInfo) HasChainwork() bool`

HasChainwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


