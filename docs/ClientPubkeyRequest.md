# ClientPubkeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to **string** | The hex format public key to register. (i.e. 02fd17dd0c52e54e5eed4ebe1e75df5e48df422f81c26520d44380bef1691fdd98) | [optional] 

## Methods

### NewClientPubkeyRequest

`func NewClientPubkeyRequest() *ClientPubkeyRequest`

NewClientPubkeyRequest instantiates a new ClientPubkeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientPubkeyRequestWithDefaults

`func NewClientPubkeyRequestWithDefaults() *ClientPubkeyRequest`

NewClientPubkeyRequestWithDefaults instantiates a new ClientPubkeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *ClientPubkeyRequest) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *ClientPubkeyRequest) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *ClientPubkeyRequest) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *ClientPubkeyRequest) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


