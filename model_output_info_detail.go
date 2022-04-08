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

// OutputInfoDetail spent status and value info of the output with detailed output script.
type OutputInfoDetail struct {
	// txid that this output is in.
	Txid *string `json:"txid,omitempty"`
	// index of this output in the tx.
	Index *int32 `json:"index,omitempty"`
	// output scrypt in hex format
	Script *string `json:"script,omitempty"`
	// parsed address of this output, empty for non standard.
	Address *string `json:"address,omitempty"`
	// value of this output
	Value *int64 `json:"value,omitempty"`
	// this output is spent or not, true if spent
	Spent *bool `json:"spent,omitempty"`
	// txid that spent this output
	SpentTxid *string `json:"spentTxid,omitempty"`
	// vin index of the spent tx
	SpentIndex *int32 `json:"spentIndex,omitempty"`
	// height of the spent tx(-1 if unconfirmed, 0 if unspent)
	SpentHeight *int32 `json:"spentHeight,omitempty"`
}

// NewOutputInfoDetail instantiates a new OutputInfoDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputInfoDetail() *OutputInfoDetail {
	this := OutputInfoDetail{}
	return &this
}

// NewOutputInfoDetailWithDefaults instantiates a new OutputInfoDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputInfoDetailWithDefaults() *OutputInfoDetail {
	this := OutputInfoDetail{}
	return &this
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *OutputInfoDetail) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfoDetail) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *OutputInfoDetail) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *OutputInfoDetail) SetTxid(v string) {
	o.Txid = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *OutputInfoDetail) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfoDetail) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *OutputInfoDetail) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *OutputInfoDetail) SetIndex(v int32) {
	o.Index = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *OutputInfoDetail) GetScript() string {
	if o == nil || o.Script == nil {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfoDetail) GetScriptOk() (*string, bool) {
	if o == nil || o.Script == nil {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *OutputInfoDetail) HasScript() bool {
	if o != nil && o.Script != nil {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *OutputInfoDetail) SetScript(v string) {
	o.Script = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OutputInfoDetail) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfoDetail) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OutputInfoDetail) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *OutputInfoDetail) SetAddress(v string) {
	o.Address = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OutputInfoDetail) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfoDetail) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OutputInfoDetail) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *OutputInfoDetail) SetValue(v int64) {
	o.Value = &v
}

// GetSpent returns the Spent field value if set, zero value otherwise.
func (o *OutputInfoDetail) GetSpent() bool {
	if o == nil || o.Spent == nil {
		var ret bool
		return ret
	}
	return *o.Spent
}

// GetSpentOk returns a tuple with the Spent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfoDetail) GetSpentOk() (*bool, bool) {
	if o == nil || o.Spent == nil {
		return nil, false
	}
	return o.Spent, true
}

// HasSpent returns a boolean if a field has been set.
func (o *OutputInfoDetail) HasSpent() bool {
	if o != nil && o.Spent != nil {
		return true
	}

	return false
}

// SetSpent gets a reference to the given bool and assigns it to the Spent field.
func (o *OutputInfoDetail) SetSpent(v bool) {
	o.Spent = &v
}

// GetSpentTxid returns the SpentTxid field value if set, zero value otherwise.
func (o *OutputInfoDetail) GetSpentTxid() string {
	if o == nil || o.SpentTxid == nil {
		var ret string
		return ret
	}
	return *o.SpentTxid
}

// GetSpentTxidOk returns a tuple with the SpentTxid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfoDetail) GetSpentTxidOk() (*string, bool) {
	if o == nil || o.SpentTxid == nil {
		return nil, false
	}
	return o.SpentTxid, true
}

// HasSpentTxid returns a boolean if a field has been set.
func (o *OutputInfoDetail) HasSpentTxid() bool {
	if o != nil && o.SpentTxid != nil {
		return true
	}

	return false
}

// SetSpentTxid gets a reference to the given string and assigns it to the SpentTxid field.
func (o *OutputInfoDetail) SetSpentTxid(v string) {
	o.SpentTxid = &v
}

// GetSpentIndex returns the SpentIndex field value if set, zero value otherwise.
func (o *OutputInfoDetail) GetSpentIndex() int32 {
	if o == nil || o.SpentIndex == nil {
		var ret int32
		return ret
	}
	return *o.SpentIndex
}

// GetSpentIndexOk returns a tuple with the SpentIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfoDetail) GetSpentIndexOk() (*int32, bool) {
	if o == nil || o.SpentIndex == nil {
		return nil, false
	}
	return o.SpentIndex, true
}

// HasSpentIndex returns a boolean if a field has been set.
func (o *OutputInfoDetail) HasSpentIndex() bool {
	if o != nil && o.SpentIndex != nil {
		return true
	}

	return false
}

// SetSpentIndex gets a reference to the given int32 and assigns it to the SpentIndex field.
func (o *OutputInfoDetail) SetSpentIndex(v int32) {
	o.SpentIndex = &v
}

// GetSpentHeight returns the SpentHeight field value if set, zero value otherwise.
func (o *OutputInfoDetail) GetSpentHeight() int32 {
	if o == nil || o.SpentHeight == nil {
		var ret int32
		return ret
	}
	return *o.SpentHeight
}

// GetSpentHeightOk returns a tuple with the SpentHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfoDetail) GetSpentHeightOk() (*int32, bool) {
	if o == nil || o.SpentHeight == nil {
		return nil, false
	}
	return o.SpentHeight, true
}

// HasSpentHeight returns a boolean if a field has been set.
func (o *OutputInfoDetail) HasSpentHeight() bool {
	if o != nil && o.SpentHeight != nil {
		return true
	}

	return false
}

// SetSpentHeight gets a reference to the given int32 and assigns it to the SpentHeight field.
func (o *OutputInfoDetail) SetSpentHeight(v int32) {
	o.SpentHeight = &v
}

func (o OutputInfoDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Script != nil {
		toSerialize["script"] = o.Script
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Spent != nil {
		toSerialize["spent"] = o.Spent
	}
	if o.SpentTxid != nil {
		toSerialize["spentTxid"] = o.SpentTxid
	}
	if o.SpentIndex != nil {
		toSerialize["spentIndex"] = o.SpentIndex
	}
	if o.SpentHeight != nil {
		toSerialize["spentHeight"] = o.SpentHeight
	}
	return json.Marshal(toSerialize)
}

type NullableOutputInfoDetail struct {
	value *OutputInfoDetail
	isSet bool
}

func (v NullableOutputInfoDetail) Get() *OutputInfoDetail {
	return v.value
}

func (v *NullableOutputInfoDetail) Set(val *OutputInfoDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputInfoDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputInfoDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputInfoDetail(val *OutputInfoDetail) *NullableOutputInfoDetail {
	return &NullableOutputInfoDetail{value: val, isSet: true}
}

func (v NullableOutputInfoDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputInfoDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
