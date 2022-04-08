# XpubRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Xpub** | Pointer to **string** | The xpub to register. | [optional] 
**SkipHeight** | Pointer to **int32** | Skip transactions before this height. Default is 0. Ignore this while deleting xpub. | [optional] 
**InitReceiveIndex** | Pointer to **int32** | Set the init maxReceiveIndex to {initReceiveIndex}(less than 5000) before the initial crawl , default is 200. This could increase cost. | [optional] 
**InitChangeIndex** | Pointer to **int32** | Set the init maxChangeIndex(less than 5000) before the initial crawl , default is 200. This could increase cost. | [optional] 

## Methods

### NewXpubRequest

`func NewXpubRequest() *XpubRequest`

NewXpubRequest instantiates a new XpubRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXpubRequestWithDefaults

`func NewXpubRequestWithDefaults() *XpubRequest`

NewXpubRequestWithDefaults instantiates a new XpubRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXpub

`func (o *XpubRequest) GetXpub() string`

GetXpub returns the Xpub field if non-nil, zero value otherwise.

### GetXpubOk

`func (o *XpubRequest) GetXpubOk() (*string, bool)`

GetXpubOk returns a tuple with the Xpub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpub

`func (o *XpubRequest) SetXpub(v string)`

SetXpub sets Xpub field to given value.

### HasXpub

`func (o *XpubRequest) HasXpub() bool`

HasXpub returns a boolean if a field has been set.

### GetSkipHeight

`func (o *XpubRequest) GetSkipHeight() int32`

GetSkipHeight returns the SkipHeight field if non-nil, zero value otherwise.

### GetSkipHeightOk

`func (o *XpubRequest) GetSkipHeightOk() (*int32, bool)`

GetSkipHeightOk returns a tuple with the SkipHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipHeight

`func (o *XpubRequest) SetSkipHeight(v int32)`

SetSkipHeight sets SkipHeight field to given value.

### HasSkipHeight

`func (o *XpubRequest) HasSkipHeight() bool`

HasSkipHeight returns a boolean if a field has been set.

### GetInitReceiveIndex

`func (o *XpubRequest) GetInitReceiveIndex() int32`

GetInitReceiveIndex returns the InitReceiveIndex field if non-nil, zero value otherwise.

### GetInitReceiveIndexOk

`func (o *XpubRequest) GetInitReceiveIndexOk() (*int32, bool)`

GetInitReceiveIndexOk returns a tuple with the InitReceiveIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitReceiveIndex

`func (o *XpubRequest) SetInitReceiveIndex(v int32)`

SetInitReceiveIndex sets InitReceiveIndex field to given value.

### HasInitReceiveIndex

`func (o *XpubRequest) HasInitReceiveIndex() bool`

HasInitReceiveIndex returns a boolean if a field has been set.

### GetInitChangeIndex

`func (o *XpubRequest) GetInitChangeIndex() int32`

GetInitChangeIndex returns the InitChangeIndex field if non-nil, zero value otherwise.

### GetInitChangeIndexOk

`func (o *XpubRequest) GetInitChangeIndexOk() (*int32, bool)`

GetInitChangeIndexOk returns a tuple with the InitChangeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitChangeIndex

`func (o *XpubRequest) SetInitChangeIndex(v int32)`

SetInitChangeIndex sets InitChangeIndex field to given value.

### HasInitChangeIndex

`func (o *XpubRequest) HasInitChangeIndex() bool`

HasInitChangeIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


