# BroadcastResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Txid** | Pointer to **string** | return txid if broadcast success | [optional] 
**Message** | Pointer to **string** | return messages if broadcast failed | [optional] 

## Methods

### NewBroadcastResult

`func NewBroadcastResult() *BroadcastResult`

NewBroadcastResult instantiates a new BroadcastResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastResultWithDefaults

`func NewBroadcastResultWithDefaults() *BroadcastResult`

NewBroadcastResultWithDefaults instantiates a new BroadcastResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxid

`func (o *BroadcastResult) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *BroadcastResult) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *BroadcastResult) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *BroadcastResult) HasTxid() bool`

HasTxid returns a boolean if a field has been set.

### GetMessage

`func (o *BroadcastResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BroadcastResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BroadcastResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BroadcastResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


