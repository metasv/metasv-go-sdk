/*
 * MetaSV API Spec
 *
 * API definition for MetaSV provided apis
 *
 * API version: 2.1.4
 * Contact: heqiming@metasv.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metasv

import (
	"encoding/json"
)

// SensibleFtUtxo Sensible fungible token Utxo belongs to the specified address
type SensibleFtUtxo struct {
	// Address string of this utxo
	Address *string `json:"address,omitempty"`
	// Codehash of this utxo.
	CodeHash *string `json:"codeHash,omitempty"`
	// Genesis of this utxo.
	Genesis *string `json:"genesis,omitempty"`
	// Name of the token.
	Name *string `json:"name,omitempty"`
	// Symbol of the token.
	Symbol *string `json:"symbol,omitempty"`
	// SensibleId of the token
	SensibleId *string `json:"sensibleId,omitempty"`
	// The decimal position.
	Decimal *int32 `json:"decimal,omitempty"`
	// Txid for this utxo.
	Txid *string `json:"txid,omitempty"`
	// Output index for the Utxo.
	TxIndex *int32 `json:"txIndex,omitempty"`
	// Token value of the utxo(Irrelavant to satoshi value).
	Value *int64 `json:"value,omitempty"`
	// Token value of the utxo(In string format)
	ValueString *string `json:"valueString,omitempty"`
	// Bsv value of the utxo(Irrelavant to token value)
	Satoshi *int64 `json:"satoshi,omitempty"`
	// Bsv value of the utxo(In string format)
	SatoshiString *string `json:"satoshiString,omitempty"`
	// The height of this utxo, -1 for unconfirmed utxo.
	Height *int32 `json:"height,omitempty"`
	// Flag used for paging
	Flag *string `json:"flag,omitempty"`
}

// NewSensibleFtUtxo instantiates a new SensibleFtUtxo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensibleFtUtxo() *SensibleFtUtxo {
	this := SensibleFtUtxo{}
	return &this
}

// NewSensibleFtUtxoWithDefaults instantiates a new SensibleFtUtxo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensibleFtUtxoWithDefaults() *SensibleFtUtxo {
	this := SensibleFtUtxo{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *SensibleFtUtxo) SetAddress(v string) {
	o.Address = &v
}

// GetCodeHash returns the CodeHash field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetCodeHash() string {
	if o == nil || o.CodeHash == nil {
		var ret string
		return ret
	}
	return *o.CodeHash
}

// GetCodeHashOk returns a tuple with the CodeHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetCodeHashOk() (*string, bool) {
	if o == nil || o.CodeHash == nil {
		return nil, false
	}
	return o.CodeHash, true
}

// HasCodeHash returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasCodeHash() bool {
	if o != nil && o.CodeHash != nil {
		return true
	}

	return false
}

// SetCodeHash gets a reference to the given string and assigns it to the CodeHash field.
func (o *SensibleFtUtxo) SetCodeHash(v string) {
	o.CodeHash = &v
}

// GetGenesis returns the Genesis field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetGenesis() string {
	if o == nil || o.Genesis == nil {
		var ret string
		return ret
	}
	return *o.Genesis
}

// GetGenesisOk returns a tuple with the Genesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetGenesisOk() (*string, bool) {
	if o == nil || o.Genesis == nil {
		return nil, false
	}
	return o.Genesis, true
}

// HasGenesis returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasGenesis() bool {
	if o != nil && o.Genesis != nil {
		return true
	}

	return false
}

// SetGenesis gets a reference to the given string and assigns it to the Genesis field.
func (o *SensibleFtUtxo) SetGenesis(v string) {
	o.Genesis = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SensibleFtUtxo) SetName(v string) {
	o.Name = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SensibleFtUtxo) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSensibleId returns the SensibleId field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetSensibleId() string {
	if o == nil || o.SensibleId == nil {
		var ret string
		return ret
	}
	return *o.SensibleId
}

// GetSensibleIdOk returns a tuple with the SensibleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetSensibleIdOk() (*string, bool) {
	if o == nil || o.SensibleId == nil {
		return nil, false
	}
	return o.SensibleId, true
}

// HasSensibleId returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasSensibleId() bool {
	if o != nil && o.SensibleId != nil {
		return true
	}

	return false
}

// SetSensibleId gets a reference to the given string and assigns it to the SensibleId field.
func (o *SensibleFtUtxo) SetSensibleId(v string) {
	o.SensibleId = &v
}

// GetDecimal returns the Decimal field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetDecimal() int32 {
	if o == nil || o.Decimal == nil {
		var ret int32
		return ret
	}
	return *o.Decimal
}

// GetDecimalOk returns a tuple with the Decimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetDecimalOk() (*int32, bool) {
	if o == nil || o.Decimal == nil {
		return nil, false
	}
	return o.Decimal, true
}

// HasDecimal returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasDecimal() bool {
	if o != nil && o.Decimal != nil {
		return true
	}

	return false
}

// SetDecimal gets a reference to the given int32 and assigns it to the Decimal field.
func (o *SensibleFtUtxo) SetDecimal(v int32) {
	o.Decimal = &v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *SensibleFtUtxo) SetTxid(v string) {
	o.Txid = &v
}

// GetTxIndex returns the TxIndex field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetTxIndex() int32 {
	if o == nil || o.TxIndex == nil {
		var ret int32
		return ret
	}
	return *o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetTxIndexOk() (*int32, bool) {
	if o == nil || o.TxIndex == nil {
		return nil, false
	}
	return o.TxIndex, true
}

// HasTxIndex returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasTxIndex() bool {
	if o != nil && o.TxIndex != nil {
		return true
	}

	return false
}

// SetTxIndex gets a reference to the given int32 and assigns it to the TxIndex field.
func (o *SensibleFtUtxo) SetTxIndex(v int32) {
	o.TxIndex = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *SensibleFtUtxo) SetValue(v int64) {
	o.Value = &v
}

// GetValueString returns the ValueString field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetValueString() string {
	if o == nil || o.ValueString == nil {
		var ret string
		return ret
	}
	return *o.ValueString
}

// GetValueStringOk returns a tuple with the ValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetValueStringOk() (*string, bool) {
	if o == nil || o.ValueString == nil {
		return nil, false
	}
	return o.ValueString, true
}

// HasValueString returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasValueString() bool {
	if o != nil && o.ValueString != nil {
		return true
	}

	return false
}

// SetValueString gets a reference to the given string and assigns it to the ValueString field.
func (o *SensibleFtUtxo) SetValueString(v string) {
	o.ValueString = &v
}

// GetSatoshi returns the Satoshi field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetSatoshi() int64 {
	if o == nil || o.Satoshi == nil {
		var ret int64
		return ret
	}
	return *o.Satoshi
}

// GetSatoshiOk returns a tuple with the Satoshi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetSatoshiOk() (*int64, bool) {
	if o == nil || o.Satoshi == nil {
		return nil, false
	}
	return o.Satoshi, true
}

// HasSatoshi returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasSatoshi() bool {
	if o != nil && o.Satoshi != nil {
		return true
	}

	return false
}

// SetSatoshi gets a reference to the given int64 and assigns it to the Satoshi field.
func (o *SensibleFtUtxo) SetSatoshi(v int64) {
	o.Satoshi = &v
}

// GetSatoshiString returns the SatoshiString field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetSatoshiString() string {
	if o == nil || o.SatoshiString == nil {
		var ret string
		return ret
	}
	return *o.SatoshiString
}

// GetSatoshiStringOk returns a tuple with the SatoshiString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetSatoshiStringOk() (*string, bool) {
	if o == nil || o.SatoshiString == nil {
		return nil, false
	}
	return o.SatoshiString, true
}

// HasSatoshiString returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasSatoshiString() bool {
	if o != nil && o.SatoshiString != nil {
		return true
	}

	return false
}

// SetSatoshiString gets a reference to the given string and assigns it to the SatoshiString field.
func (o *SensibleFtUtxo) SetSatoshiString(v string) {
	o.SatoshiString = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetHeightOk() (*int32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *SensibleFtUtxo) SetHeight(v int32) {
	o.Height = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *SensibleFtUtxo) GetFlag() string {
	if o == nil || o.Flag == nil {
		var ret string
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleFtUtxo) GetFlagOk() (*string, bool) {
	if o == nil || o.Flag == nil {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *SensibleFtUtxo) HasFlag() bool {
	if o != nil && o.Flag != nil {
		return true
	}

	return false
}

// SetFlag gets a reference to the given string and assigns it to the Flag field.
func (o *SensibleFtUtxo) SetFlag(v string) {
	o.Flag = &v
}

func (o SensibleFtUtxo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.CodeHash != nil {
		toSerialize["codeHash"] = o.CodeHash
	}
	if o.Genesis != nil {
		toSerialize["genesis"] = o.Genesis
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.SensibleId != nil {
		toSerialize["sensibleId"] = o.SensibleId
	}
	if o.Decimal != nil {
		toSerialize["decimal"] = o.Decimal
	}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.TxIndex != nil {
		toSerialize["txIndex"] = o.TxIndex
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueString != nil {
		toSerialize["valueString"] = o.ValueString
	}
	if o.Satoshi != nil {
		toSerialize["satoshi"] = o.Satoshi
	}
	if o.SatoshiString != nil {
		toSerialize["satoshiString"] = o.SatoshiString
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Flag != nil {
		toSerialize["flag"] = o.Flag
	}
	return json.Marshal(toSerialize)
}

type NullableSensibleFtUtxo struct {
	value *SensibleFtUtxo
	isSet bool
}

func (v NullableSensibleFtUtxo) Get() *SensibleFtUtxo {
	return v.value
}

func (v *NullableSensibleFtUtxo) Set(val *SensibleFtUtxo) {
	v.value = val
	v.isSet = true
}

func (v NullableSensibleFtUtxo) IsSet() bool {
	return v.isSet
}

func (v *NullableSensibleFtUtxo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensibleFtUtxo(val *SensibleFtUtxo) *NullableSensibleFtUtxo {
	return &NullableSensibleFtUtxo{value: val, isSet: true}
}

func (v NullableSensibleFtUtxo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensibleFtUtxo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
