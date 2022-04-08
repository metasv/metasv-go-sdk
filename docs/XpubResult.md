# XpubResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Xpub** | Pointer to **string** | return xpub if broadcast success | [optional] 
**Message** | Pointer to **string** | return messages if broadcast failed | [optional] 

## Methods

### NewXpubResult

`func NewXpubResult() *XpubResult`

NewXpubResult instantiates a new XpubResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXpubResultWithDefaults

`func NewXpubResultWithDefaults() *XpubResult`

NewXpubResultWithDefaults instantiates a new XpubResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXpub

`func (o *XpubResult) GetXpub() string`

GetXpub returns the Xpub field if non-nil, zero value otherwise.

### GetXpubOk

`func (o *XpubResult) GetXpubOk() (*string, bool)`

GetXpubOk returns a tuple with the Xpub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpub

`func (o *XpubResult) SetXpub(v string)`

SetXpub sets Xpub field to given value.

### HasXpub

`func (o *XpubResult) HasXpub() bool`

HasXpub returns a boolean if a field has been set.

### GetMessage

`func (o *XpubResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *XpubResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *XpubResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *XpubResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


