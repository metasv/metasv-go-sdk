# XpubAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Xpub** | Pointer to **string** | xpub in the query | [optional] 
**Address** | Pointer to **string** | Address in the query | [optional] 
**AddressType** | Pointer to **int32** | Address type, 0 for receive address, 1 for change address. path is {{addressType}}/{{addressIndex}} | [optional] 
**AddressIndex** | Pointer to **int32** | Address index. Address path in the xpub is {{addressType}}/{{addressIndex}} | [optional] 

## Methods

### NewXpubAddress

`func NewXpubAddress() *XpubAddress`

NewXpubAddress instantiates a new XpubAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXpubAddressWithDefaults

`func NewXpubAddressWithDefaults() *XpubAddress`

NewXpubAddressWithDefaults instantiates a new XpubAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXpub

`func (o *XpubAddress) GetXpub() string`

GetXpub returns the Xpub field if non-nil, zero value otherwise.

### GetXpubOk

`func (o *XpubAddress) GetXpubOk() (*string, bool)`

GetXpubOk returns a tuple with the Xpub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpub

`func (o *XpubAddress) SetXpub(v string)`

SetXpub sets Xpub field to given value.

### HasXpub

`func (o *XpubAddress) HasXpub() bool`

HasXpub returns a boolean if a field has been set.

### GetAddress

`func (o *XpubAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *XpubAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *XpubAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *XpubAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressType

`func (o *XpubAddress) GetAddressType() int32`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *XpubAddress) GetAddressTypeOk() (*int32, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *XpubAddress) SetAddressType(v int32)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *XpubAddress) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetAddressIndex

`func (o *XpubAddress) GetAddressIndex() int32`

GetAddressIndex returns the AddressIndex field if non-nil, zero value otherwise.

### GetAddressIndexOk

`func (o *XpubAddress) GetAddressIndexOk() (*int32, bool)`

GetAddressIndexOk returns a tuple with the AddressIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressIndex

`func (o *XpubAddress) SetAddressIndex(v int32)`

SetAddressIndex sets AddressIndex field to given value.

### HasAddressIndex

`func (o *XpubAddress) HasAddressIndex() bool`

HasAddressIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


