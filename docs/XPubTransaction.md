# XPubTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Xpub** | Pointer to **string** | query xpub | [optional] 
**Txid** | Pointer to **string** | Txid for this transaction. | [optional] 
**MaxReceiveIndex** | Pointer to **int32** | Max lookahead receive index when processing this transaction. | [optional] 
**MaxChangeIndex** | Pointer to **int32** | Max lookahead change index when processing this transaction. | [optional] 
**Income** | Pointer to **int64** | Total received satoshis(Including all address) | [optional] 
**Outcome** | Pointer to **int64** | Total spent satoshis(Including all address) | [optional] 
**Height** | Pointer to **int32** | Height for this transaction. -1 for unconfirmed | [optional] 
**BlockIndex** | Pointer to **int32** | Block index for this transaction, -1 for unconfirmed | [optional] 
**BlockTime** | Pointer to **int64** | Block timestamp for this transaction, if unconfirmed, the time is first seen time. | [optional] 
**Flag** | Pointer to **string** | Paging flag, format blockTimestamp_blockIndex | [optional] 

## Methods

### NewXPubTransaction

`func NewXPubTransaction() *XPubTransaction`

NewXPubTransaction instantiates a new XPubTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXPubTransactionWithDefaults

`func NewXPubTransactionWithDefaults() *XPubTransaction`

NewXPubTransactionWithDefaults instantiates a new XPubTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXpub

`func (o *XPubTransaction) GetXpub() string`

GetXpub returns the Xpub field if non-nil, zero value otherwise.

### GetXpubOk

`func (o *XPubTransaction) GetXpubOk() (*string, bool)`

GetXpubOk returns a tuple with the Xpub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpub

`func (o *XPubTransaction) SetXpub(v string)`

SetXpub sets Xpub field to given value.

### HasXpub

`func (o *XPubTransaction) HasXpub() bool`

HasXpub returns a boolean if a field has been set.

### GetTxid

`func (o *XPubTransaction) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *XPubTransaction) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *XPubTransaction) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *XPubTransaction) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetMaxReceiveIndex

`func (o *XPubTransaction) GetMaxReceiveIndex() int32`

GetMaxReceiveIndex returns the MaxReceiveIndex field if non-nil, zero value otherwise.

### GetMaxReceiveIndexOk

`func (o *XPubTransaction) GetMaxReceiveIndexOk() (*int32, bool)`

GetMaxReceiveIndexOk returns a tuple with the MaxReceiveIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReceiveIndex

`func (o *XPubTransaction) SetMaxReceiveIndex(v int32)`

SetMaxReceiveIndex sets MaxReceiveIndex field to given value.

### HasMaxReceiveIndex

`func (o *XPubTransaction) HasMaxReceiveIndex() bool`

HasMaxReceiveIndex returns a boolean if a field has been set.

### GetMaxChangeIndex

`func (o *XPubTransaction) GetMaxChangeIndex() int32`

GetMaxChangeIndex returns the MaxChangeIndex field if non-nil, zero value otherwise.

### GetMaxChangeIndexOk

`func (o *XPubTransaction) GetMaxChangeIndexOk() (*int32, bool)`

GetMaxChangeIndexOk returns a tuple with the MaxChangeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChangeIndex

`func (o *XPubTransaction) SetMaxChangeIndex(v int32)`

SetMaxChangeIndex sets MaxChangeIndex field to given value.

### HasMaxChangeIndex

`func (o *XPubTransaction) HasMaxChangeIndex() bool`

HasMaxChangeIndex returns a boolean if a field has been set.

### GetIncome

`func (o *XPubTransaction) GetIncome() int64`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *XPubTransaction) GetIncomeOk() (*int64, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *XPubTransaction) SetIncome(v int64)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *XPubTransaction) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetOutcome

`func (o *XPubTransaction) GetOutcome() int64`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *XPubTransaction) GetOutcomeOk() (*int64, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *XPubTransaction) SetOutcome(v int64)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *XPubTransaction) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetHeight

`func (o *XPubTransaction) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *XPubTransaction) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *XPubTransaction) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *XPubTransaction) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetBlockIndex

`func (o *XPubTransaction) GetBlockIndex() int32`

GetBlockIndex returns the BlockIndex field if non-nil, zero value otherwise.

### GetBlockIndexOk

`func (o *XPubTransaction) GetBlockIndexOk() (*int32, bool)`

GetBlockIndexOk returns a tuple with the BlockIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIndex

`func (o *XPubTransaction) SetBlockIndex(v int32)`

SetBlockIndex sets BlockIndex field to given value.

### HasBlockIndex

`func (o *XPubTransaction) HasBlockIndex() bool`

HasBlockIndex returns a boolean if a field has been set.

### GetBlockTime

`func (o *XPubTransaction) GetBlockTime() int64`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *XPubTransaction) GetBlockTimeOk() (*int64, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *XPubTransaction) SetBlockTime(v int64)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *XPubTransaction) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetFlag

`func (o *XPubTransaction) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *XPubTransaction) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *XPubTransaction) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *XPubTransaction) HasFlag() bool`

HasFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


