# AddressTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flag** | Pointer to **string** | Paging flag | [optional] 
**Address** | Pointer to **string** | The address of the record | [optional] 
**Time** | Pointer to **int64** | timestamp of the tx | [optional] 
**Height** | Pointer to **int32** | Block Height of the tx, -1 if not confirmed | [optional] 
**Income** | Pointer to **int64** | total income of the address from this tx | [optional] 
**Outcome** | Pointer to **int64** | total outcome of the address from this tx | [optional] 
**Txid** | Pointer to **string** | txid of this record | [optional] 

## Methods

### NewAddressTx

`func NewAddressTx() *AddressTx`

NewAddressTx instantiates a new AddressTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTxWithDefaults

`func NewAddressTxWithDefaults() *AddressTx`

NewAddressTxWithDefaults instantiates a new AddressTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlag

`func (o *AddressTx) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *AddressTx) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *AddressTx) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *AddressTx) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetAddress

`func (o *AddressTx) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressTx) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressTx) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddressTx) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTime

`func (o *AddressTx) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AddressTx) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AddressTx) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *AddressTx) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetHeight

`func (o *AddressTx) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *AddressTx) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *AddressTx) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *AddressTx) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetIncome

`func (o *AddressTx) GetIncome() int64`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *AddressTx) GetIncomeOk() (*int64, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *AddressTx) SetIncome(v int64)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *AddressTx) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetOutcome

`func (o *AddressTx) GetOutcome() int64`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AddressTx) GetOutcomeOk() (*int64, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AddressTx) SetOutcome(v int64)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *AddressTx) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetTxid

`func (o *AddressTx) GetTxid() string`

GetTxid returns the Txid field if non-nil, zero value otherwise.

### GetTxidOk

`func (o *AddressTx) GetTxidOk() (*string, bool)`

GetTxidOk returns a tuple with the Txid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxid

`func (o *AddressTx) SetTxid(v string)`

SetTxid sets Txid field to given value.

### HasTxid

`func (o *AddressTx) HasTxid() bool`

HasTxid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


