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

// AddressUtxo Utxos by address(or addresses) order by height asc.
type AddressUtxo struct {
	// paging flag
	Flag *string `json:"flag,omitempty"`
	// Address of the utxo
	Address *string `json:"address,omitempty"`
	// txid of the utxo
	Txid *string `json:"txid,omitempty"`
	// output index in the tx
	OutIndex *int32 `json:"outIndex,omitempty"`
	// Value of the utxo
	Value *int64 `json:"value,omitempty"`
	// Height of the utxo, -1 if not confirmed
	Height *int32 `json:"height,omitempty"`
}

// NewAddressUtxo instantiates a new AddressUtxo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressUtxo() *AddressUtxo {
	this := AddressUtxo{}
	return &this
}

// NewAddressUtxoWithDefaults instantiates a new AddressUtxo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressUtxoWithDefaults() *AddressUtxo {
	this := AddressUtxo{}
	return &this
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *AddressUtxo) GetFlag() string {
	if o == nil || o.Flag == nil {
		var ret string
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUtxo) GetFlagOk() (*string, bool) {
	if o == nil || o.Flag == nil {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *AddressUtxo) HasFlag() bool {
	if o != nil && o.Flag != nil {
		return true
	}

	return false
}

// SetFlag gets a reference to the given string and assigns it to the Flag field.
func (o *AddressUtxo) SetFlag(v string) {
	o.Flag = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AddressUtxo) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUtxo) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AddressUtxo) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AddressUtxo) SetAddress(v string) {
	o.Address = &v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *AddressUtxo) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUtxo) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *AddressUtxo) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *AddressUtxo) SetTxid(v string) {
	o.Txid = &v
}

// GetOutIndex returns the OutIndex field value if set, zero value otherwise.
func (o *AddressUtxo) GetOutIndex() int32 {
	if o == nil || o.OutIndex == nil {
		var ret int32
		return ret
	}
	return *o.OutIndex
}

// GetOutIndexOk returns a tuple with the OutIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUtxo) GetOutIndexOk() (*int32, bool) {
	if o == nil || o.OutIndex == nil {
		return nil, false
	}
	return o.OutIndex, true
}

// HasOutIndex returns a boolean if a field has been set.
func (o *AddressUtxo) HasOutIndex() bool {
	if o != nil && o.OutIndex != nil {
		return true
	}

	return false
}

// SetOutIndex gets a reference to the given int32 and assigns it to the OutIndex field.
func (o *AddressUtxo) SetOutIndex(v int32) {
	o.OutIndex = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AddressUtxo) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUtxo) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AddressUtxo) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *AddressUtxo) SetValue(v int64) {
	o.Value = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *AddressUtxo) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUtxo) GetHeightOk() (*int32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *AddressUtxo) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *AddressUtxo) SetHeight(v int32) {
	o.Height = &v
}

func (o AddressUtxo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Flag != nil {
		toSerialize["flag"] = o.Flag
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.OutIndex != nil {
		toSerialize["outIndex"] = o.OutIndex
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	return json.Marshal(toSerialize)
}

type NullableAddressUtxo struct {
	value *AddressUtxo
	isSet bool
}

func (v NullableAddressUtxo) Get() *AddressUtxo {
	return v.value
}

func (v *NullableAddressUtxo) Set(val *AddressUtxo) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressUtxo) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressUtxo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressUtxo(val *AddressUtxo) *NullableAddressUtxo {
	return &NullableAddressUtxo{value: val, isSet: true}
}

func (v NullableAddressUtxo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressUtxo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
