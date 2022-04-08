# AsyncBroadcastResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Txid** | Pointer to **string** | the txid of this tx | [optional] 
**State** | Pointer to **string** | the state of this tx, possible values PENDING, BROADCASTED, INVALID, UNKNOWN | [optional] 
**Message** | Pointer to **string** | return messages if broadcast failed | [optional] 

## Methods

### NewAsyncBroadcastResult

`func NewAsyncBroadcastResult() *AsyncBroadcastResult`

NewAsyncBroadcastResult instantiates a new AsyncBroadcastResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncBroadcastResultWithDefaults

`func NewAsyncBroadcastResultWithDefaults() *AsyncBroadcastResult`

NewAsyncBroadcastResultWithDefaults instantiates a new AsyncBroadcastResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxid

`func (o *AsyncBroadcastResult) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *AsyncBroadcastResult) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *AsyncBroadcastResult) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *AsyncBroadcastResult) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetState

`func (o *AsyncBroadcastResult) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AsyncBroadcastResult) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AsyncBroadcastResult) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AsyncBroadcastResult) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessage

`func (o *AsyncBroadcastResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AsyncBroadcastResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AsyncBroadcastResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AsyncBroadcastResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


