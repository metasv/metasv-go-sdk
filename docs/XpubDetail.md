# XpubDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Xpub** | Pointer to **string** | String encoded extended pubkey (xpub) | [optional] 
**ReceiveIndex** | Pointer to **int32** | Next unused receive index, path /0/index | [optional] 
**MaxReceiveIndex** | Pointer to **int32** | Max lookahead receive index. | [optional] 
**ChangeIndex** | Pointer to **int32** | Next unused change index, path /1/index | [optional] 
**MaxChangeIndex** | Pointer to **int32** | Max lookahead change index. | [optional] 
**Mode** | Pointer to **int32** | Current xpub process mode, 0 means preparing(not ready), 1 means synchronizing(ready) | [optional] 
**SkipHeight** | Pointer to **int32** | Skip blocks before skipHeight while searching transactions. This will speed up sync time. | [optional] 
**ProcessHeight** | Pointer to **int32** | Xpub current processed height. | [optional] 

## Methods

### NewXpubDetail

`func NewXpubDetail() *XpubDetail`

NewXpubDetail instantiates a new XpubDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXpubDetailWithDefaults

`func NewXpubDetailWithDefaults() *XpubDetail`

NewXpubDetailWithDefaults instantiates a new XpubDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXpub

`func (o *XpubDetail) GetXpub() string`

GetXpub returns the Xpub field if non-nil, zero value otherwise.

### GetXpubOk

`func (o *XpubDetail) GetXpubOk() (*string, bool)`

GetXpubOk returns a tuple with the Xpub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpub

`func (o *XpubDetail) SetXpub(v string)`

SetXpub sets Xpub field to given value.

### HasXpub

`func (o *XpubDetail) HasXpub() bool`

HasXpub returns a boolean if a field has been set.

### GetReceiveIndex

`func (o *XpubDetail) GetReceiveIndex() int32`

GetReceiveIndex returns the ReceiveIndex field if non-nil, zero value otherwise.

### GetReceiveIndexOk

`func (o *XpubDetail) GetReceiveIndexOk() (*int32, bool)`

GetReceiveIndexOk returns a tuple with the ReceiveIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveIndex

`func (o *XpubDetail) SetReceiveIndex(v int32)`

SetReceiveIndex sets ReceiveIndex field to given value.

### HasReceiveIndex

`func (o *XpubDetail) HasReceiveIndex() bool`

HasReceiveIndex returns a boolean if a field has been set.

### GetMaxReceiveIndex

`func (o *XpubDetail) GetMaxReceiveIndex() int32`

GetMaxReceiveIndex returns the MaxReceiveIndex field if non-nil, zero value otherwise.

### GetMaxReceiveIndexOk

`func (o *XpubDetail) GetMaxReceiveIndexOk() (*int32, bool)`

GetMaxReceiveIndexOk returns a tuple with the MaxReceiveIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReceiveIndex

`func (o *XpubDetail) SetMaxReceiveIndex(v int32)`

SetMaxReceiveIndex sets MaxReceiveIndex field to given value.

### HasMaxReceiveIndex

`func (o *XpubDetail) HasMaxReceiveIndex() bool`

HasMaxReceiveIndex returns a boolean if a field has been set.

### GetChangeIndex

`func (o *XpubDetail) GetChangeIndex() int32`

GetChangeIndex returns the ChangeIndex field if non-nil, zero value otherwise.

### GetChangeIndexOk

`func (o *XpubDetail) GetChangeIndexOk() (*int32, bool)`

GetChangeIndexOk returns a tuple with the ChangeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeIndex

`func (o *XpubDetail) SetChangeIndex(v int32)`

SetChangeIndex sets ChangeIndex field to given value.

### HasChangeIndex

`func (o *XpubDetail) HasChangeIndex() bool`

HasChangeIndex returns a boolean if a field has been set.

### GetMaxChangeIndex

`func (o *XpubDetail) GetMaxChangeIndex() int32`

GetMaxChangeIndex returns the MaxChangeIndex field if non-nil, zero value otherwise.

### GetMaxChangeIndexOk

`func (o *XpubDetail) GetMaxChangeIndexOk() (*int32, bool)`

GetMaxChangeIndexOk returns a tuple with the MaxChangeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChangeIndex

`func (o *XpubDetail) SetMaxChangeIndex(v int32)`

SetMaxChangeIndex sets MaxChangeIndex field to given value.

### HasMaxChangeIndex

`func (o *XpubDetail) HasMaxChangeIndex() bool`

HasMaxChangeIndex returns a boolean if a field has been set.

### GetMode

`func (o *XpubDetail) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *XpubDetail) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *XpubDetail) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *XpubDetail) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSkipHeight

`func (o *XpubDetail) GetSkipHeight() int32`

GetSkipHeight returns the SkipHeight field if non-nil, zero value otherwise.

### GetSkipHeightOk

`func (o *XpubDetail) GetSkipHeightOk() (*int32, bool)`

GetSkipHeightOk returns a tuple with the SkipHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipHeight

`func (o *XpubDetail) SetSkipHeight(v int32)`

SetSkipHeight sets SkipHeight field to given value.

### HasSkipHeight

`func (o *XpubDetail) HasSkipHeight() bool`

HasSkipHeight returns a boolean if a field has been set.

### GetProcessHeight

`func (o *XpubDetail) GetProcessHeight() int32`

GetProcessHeight returns the ProcessHeight field if non-nil, zero value otherwise.

### GetProcessHeightOk

`func (o *XpubDetail) GetProcessHeightOk() (*int32, bool)`

GetProcessHeightOk returns a tuple with the ProcessHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessHeight

`func (o *XpubDetail) SetProcessHeight(v int32)`

SetProcessHeight sets ProcessHeight field to given value.

### HasProcessHeight

`func (o *XpubDetail) HasProcessHeight() bool`

HasProcessHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


