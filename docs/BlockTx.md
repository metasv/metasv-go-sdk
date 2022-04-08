# BlockTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Txid** | Pointer to **string** | Transaction hash. | [optional] 
**Height** | Pointer to **int32** | Block height of this tx. | [optional] 
**BlockHash** | Pointer to **string** | Hash of the block | [optional] 
**Size** | Pointer to **int64** | transaction size | [optional] 
**InputCount** | Pointer to **int32** | Input count in this transaction | [optional] 
**OutputCount** | Pointer to **int32** | Output count in this transaction. | [optional] 
**LockTime** | Pointer to **int64** | Lock time of this transaction | [optional] 
**Fee** | Pointer to **int64** | Trasaction fee. | [optional] 
**Confirmations** | Pointer to **int32** | Confirmations of this transaction, -1 if unconfirmed. | [optional] 
**Timestamp** | Pointer to **int64** | Block timestamp for the transaction, confirmed tx only. | [optional] 

## Methods

### NewBlockTx

`func NewBlockTx() *BlockTx`

NewBlockTx instantiates a new BlockTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockTxWithDefaults

`func NewBlockTxWithDefaults() *BlockTx`

NewBlockTxWithDefaults instantiates a new BlockTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxid

`func (o *BlockTx) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *BlockTx) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *BlockTx) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *BlockTx) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetHeight

`func (o *BlockTx) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BlockTx) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BlockTx) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *BlockTx) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetBlockHash

`func (o *BlockTx) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *BlockTx) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *BlockTx) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *BlockTx) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetSize

`func (o *BlockTx) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BlockTx) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BlockTx) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BlockTx) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetInputCount

`func (o *BlockTx) GetInputCount() int32`

GetInputCount returns the InputCount field if non-nil, zero value otherwise.

### GetInputCountOk

`func (o *BlockTx) GetInputCountOk() (*int32, bool)`

GetInputCountOk returns a tuple with the InputCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCount

`func (o *BlockTx) SetInputCount(v int32)`

SetInputCount sets InputCount field to given value.

### HasInputCount

`func (o *BlockTx) HasInputCount() bool`

HasInputCount returns a boolean if a field has been set.

### GetOutputCount

`func (o *BlockTx) GetOutputCount() int32`

GetOutputCount returns the OutputCount field if non-nil, zero value otherwise.

### GetOutputCountOk

`func (o *BlockTx) GetOutputCountOk() (*int32, bool)`

GetOutputCountOk returns a tuple with the OutputCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCount

`func (o *BlockTx) SetOutputCount(v int32)`

SetOutputCount sets OutputCount field to given value.

### HasOutputCount

`func (o *BlockTx) HasOutputCount() bool`

HasOutputCount returns a boolean if a field has been set.

### GetLockTime

`func (o *BlockTx) GetLockTime() int64`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *BlockTx) GetLockTimeOk() (*int64, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *BlockTx) SetLockTime(v int64)`

SetLockTime sets LockTime field to given value.

### HasLockTime

`func (o *BlockTx) HasLockTime() bool`

HasLockTime returns a boolean if a field has been set.

### GetFee

`func (o *BlockTx) GetFee() int64`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *BlockTx) GetFeeOk() (*int64, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *BlockTx) SetFee(v int64)`

SetFee sets Fee field to given value.

### HasFee

`func (o *BlockTx) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetConfirmations

`func (o *BlockTx) GetConfirmations() int32`

GetConfirmations returns the Confirmations field if non-nil, zero value otherwise.

### GetConfirmationsOk

`func (o *BlockTx) GetConfirmationsOk() (*int32, bool)`

GetConfirmationsOk returns a tuple with the Confirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmations

`func (o *BlockTx) SetConfirmations(v int32)`

SetConfirmations sets Confirmations field to given value.

### HasConfirmations

`func (o *BlockTx) HasConfirmations() bool`

HasConfirmations returns a boolean if a field has been set.

### GetTimestamp

`func (o *BlockTx) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlockTx) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlockTx) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BlockTx) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


