# ClientPubkeyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to **string** | return pubkey if created successfully | [optional] 
**Message** | Pointer to **string** | return messages if broadcast failed | [optional] 

## Methods

### NewClientPubkeyResult

`func NewClientPubkeyResult() *ClientPubkeyResult`

NewClientPubkeyResult instantiates a new ClientPubkeyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientPubkeyResultWithDefaults

`func NewClientPubkeyResultWithDefaults() *ClientPubkeyResult`

NewClientPubkeyResultWithDefaults instantiates a new ClientPubkeyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *ClientPubkeyResult) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *ClientPubkeyResult) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *ClientPubkeyResult) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *ClientPubkeyResult) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetMessage

`func (o *ClientPubkeyResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClientPubkeyResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClientPubkeyResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClientPubkeyResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


