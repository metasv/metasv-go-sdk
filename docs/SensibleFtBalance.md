# SensibleFtBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeHash** | Pointer to **string** | Codehash of the token. | [optional] 
**Genesis** | Pointer to **string** | Genesis of the token. | [optional] 
**Name** | Pointer to **string** | Name of the token. | [optional] 
**Symbol** | Pointer to **string** | Symbol of the token. | [optional] 
**Decimal** | Pointer to **int32** | The decimal position. | [optional] 
**SensibleId** | Pointer to **string** | SensibleId of the token | [optional] 
**UtxoCount** | Pointer to **int32** | Number of utxos for this token. | [optional] 
**Confirmed** | Pointer to **int64** | Confirmed balance of the token. | [optional] 
**ConfirmedString** | Pointer to **string** | Confirmed balance of the token(In string format). | [optional] 
**Unconfirmed** | Pointer to **int64** | Unconfirmed balance of the token. | [optional] 
**UnconfirmedString** | Pointer to **string** | Unconfirmed balance of the token(In string format). | [optional] 

## Methods

### NewSensibleFtBalance

`func NewSensibleFtBalance() *SensibleFtBalance`

NewSensibleFtBalance instantiates a new SensibleFtBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensibleFtBalanceWithDefaults

`func NewSensibleFtBalanceWithDefaults() *SensibleFtBalance`

NewSensibleFtBalanceWithDefaults instantiates a new SensibleFtBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeHash

`func (o *SensibleFtBalance) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *SensibleFtBalance) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *SensibleFtBalance) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.

### HasCodeHash

`func (o *SensibleFtBalance) HasCodeHash() bool`

HasCodeHash returns a boolean if a field has been set.

### GetGenesis

`func (o *SensibleFtBalance) GetGenesis() string`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *SensibleFtBalance) GetGenesisOk() (*string, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *SensibleFtBalance) SetGenesis(v string)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *SensibleFtBalance) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetName

`func (o *SensibleFtBalance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SensibleFtBalance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SensibleFtBalance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SensibleFtBalance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSymbol

`func (o *SensibleFtBalance) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SensibleFtBalance) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SensibleFtBalance) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SensibleFtBalance) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDecimal

`func (o *SensibleFtBalance) GetDecimal() int32`

GetDecimal returns the Decimal field if non-nil, zero value otherwise.

### GetDecimalOk

`func (o *SensibleFtBalance) GetDecimalOk() (*int32, bool)`

GetDecimalOk returns a tuple with the Decimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimal

`func (o *SensibleFtBalance) SetDecimal(v int32)`

SetDecimal sets Decimal field to given value.

### HasDecimal

`func (o *SensibleFtBalance) HasDecimal() bool`

HasDecimal returns a boolean if a field has been set.

### GetSensibleId

`func (o *SensibleFtBalance) GetSensibleId() string`

GetSensibleId returns the SensibleId field if non-nil, zero value otherwise.

### GetSensibleIdOk

`func (o *SensibleFtBalance) GetSensibleIdOk() (*string, bool)`

GetSensibleIdOk returns a tuple with the SensibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensibleId

`func (o *SensibleFtBalance) SetSensibleId(v string)`

SetSensibleId sets SensibleId field to given value.

### HasSensibleId

`func (o *SensibleFtBalance) HasSensibleId() bool`

HasSensibleId returns a boolean if a field has been set.

### GetUtxoCount

`func (o *SensibleFtBalance) GetUtxoCount() int32`

GetUtxoCount returns the UtxoCount field if non-nil, zero value otherwise.

### GetUtxoCountOk

`func (o *SensibleFtBalance) GetUtxoCountOk() (*int32, bool)`

GetUtxoCountOk returns a tuple with the UtxoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoCount

`func (o *SensibleFtBalance) SetUtxoCount(v int32)`

SetUtxoCount sets UtxoCount field to given value.

### HasUtxoCount

`func (o *SensibleFtBalance) HasUtxoCount() bool`

HasUtxoCount returns a boolean if a field has been set.

### GetConfirmed

`func (o *SensibleFtBalance) GetConfirmed() int64`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *SensibleFtBalance) GetConfirmedOk() (*int64, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *SensibleFtBalance) SetConfirmed(v int64)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *SensibleFtBalance) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetConfirmedString

`func (o *SensibleFtBalance) GetConfirmedString() string`

GetConfirmedString returns the ConfirmedString field if non-nil, zero value otherwise.

### GetConfirmedStringOk

`func (o *SensibleFtBalance) GetConfirmedStringOk() (*string, bool)`

GetConfirmedStringOk returns a tuple with the ConfirmedString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedString

`func (o *SensibleFtBalance) SetConfirmedString(v string)`

SetConfirmedString sets ConfirmedString field to given value.

### HasConfirmedString

`func (o *SensibleFtBalance) HasConfirmedString() bool`

HasConfirmedString returns a boolean if a field has been set.

### GetUnconfirmed

`func (o *SensibleFtBalance) GetUnconfirmed() int64`

GetUnconfirmed returns the Unconfirmed field if non-nil, zero value otherwise.

### GetUnconfirmedOk

`func (o *SensibleFtBalance) GetUnconfirmedOk() (*int64, bool)`

GetUnconfirmedOk returns a tuple with the Unconfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfirmed

`func (o *SensibleFtBalance) SetUnconfirmed(v int64)`

SetUnconfirmed sets Unconfirmed field to given value.

### HasUnconfirmed

`func (o *SensibleFtBalance) HasUnconfirmed() bool`

HasUnconfirmed returns a boolean if a field has been set.

### GetUnconfirmedString

`func (o *SensibleFtBalance) GetUnconfirmedString() string`

GetUnconfirmedString returns the UnconfirmedString field if non-nil, zero value otherwise.

### GetUnconfirmedStringOk

`func (o *SensibleFtBalance) GetUnconfirmedStringOk() (*string, bool)`

GetUnconfirmedStringOk returns a tuple with the UnconfirmedString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfirmedString

`func (o *SensibleFtBalance) SetUnconfirmedString(v string)`

SetUnconfirmedString sets UnconfirmedString field to given value.

### HasUnconfirmedString

`func (o *SensibleFtBalance) HasUnconfirmedString() bool`

HasUnconfirmedString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


