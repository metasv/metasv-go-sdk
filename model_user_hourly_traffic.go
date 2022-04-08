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

// UserHourlyTraffic Summed outbound traffic within particular hour.
type UserHourlyTraffic struct {
	// The unix timestamp of the beginning of the hour.
	HourTime *int64 `json:"hourTime,omitempty"`
	// Parsed utc time string for readability.
	HourTimeString *string `json:"hourTimeString,omitempty"`
	// Total outbound bytes within this hour
	TotalBytes *int64 `json:"totalBytes,omitempty"`
}

// NewUserHourlyTraffic instantiates a new UserHourlyTraffic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserHourlyTraffic() *UserHourlyTraffic {
	this := UserHourlyTraffic{}
	return &this
}

// NewUserHourlyTrafficWithDefaults instantiates a new UserHourlyTraffic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserHourlyTrafficWithDefaults() *UserHourlyTraffic {
	this := UserHourlyTraffic{}
	return &this
}

// GetHourTime returns the HourTime field value if set, zero value otherwise.
func (o *UserHourlyTraffic) GetHourTime() int64 {
	if o == nil || o.HourTime == nil {
		var ret int64
		return ret
	}
	return *o.HourTime
}

// GetHourTimeOk returns a tuple with the HourTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserHourlyTraffic) GetHourTimeOk() (*int64, bool) {
	if o == nil || o.HourTime == nil {
		return nil, false
	}
	return o.HourTime, true
}

// HasHourTime returns a boolean if a field has been set.
func (o *UserHourlyTraffic) HasHourTime() bool {
	if o != nil && o.HourTime != nil {
		return true
	}

	return false
}

// SetHourTime gets a reference to the given int64 and assigns it to the HourTime field.
func (o *UserHourlyTraffic) SetHourTime(v int64) {
	o.HourTime = &v
}

// GetHourTimeString returns the HourTimeString field value if set, zero value otherwise.
func (o *UserHourlyTraffic) GetHourTimeString() string {
	if o == nil || o.HourTimeString == nil {
		var ret string
		return ret
	}
	return *o.HourTimeString
}

// GetHourTimeStringOk returns a tuple with the HourTimeString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserHourlyTraffic) GetHourTimeStringOk() (*string, bool) {
	if o == nil || o.HourTimeString == nil {
		return nil, false
	}
	return o.HourTimeString, true
}

// HasHourTimeString returns a boolean if a field has been set.
func (o *UserHourlyTraffic) HasHourTimeString() bool {
	if o != nil && o.HourTimeString != nil {
		return true
	}

	return false
}

// SetHourTimeString gets a reference to the given string and assigns it to the HourTimeString field.
func (o *UserHourlyTraffic) SetHourTimeString(v string) {
	o.HourTimeString = &v
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise.
func (o *UserHourlyTraffic) GetTotalBytes() int64 {
	if o == nil || o.TotalBytes == nil {
		var ret int64
		return ret
	}
	return *o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserHourlyTraffic) GetTotalBytesOk() (*int64, bool) {
	if o == nil || o.TotalBytes == nil {
		return nil, false
	}
	return o.TotalBytes, true
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *UserHourlyTraffic) HasTotalBytes() bool {
	if o != nil && o.TotalBytes != nil {
		return true
	}

	return false
}

// SetTotalBytes gets a reference to the given int64 and assigns it to the TotalBytes field.
func (o *UserHourlyTraffic) SetTotalBytes(v int64) {
	o.TotalBytes = &v
}

func (o UserHourlyTraffic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HourTime != nil {
		toSerialize["hourTime"] = o.HourTime
	}
	if o.HourTimeString != nil {
		toSerialize["hourTimeString"] = o.HourTimeString
	}
	if o.TotalBytes != nil {
		toSerialize["totalBytes"] = o.TotalBytes
	}
	return json.Marshal(toSerialize)
}

type NullableUserHourlyTraffic struct {
	value *UserHourlyTraffic
	isSet bool
}

func (v NullableUserHourlyTraffic) Get() *UserHourlyTraffic {
	return v.value
}

func (v *NullableUserHourlyTraffic) Set(val *UserHourlyTraffic) {
	v.value = val
	v.isSet = true
}

func (v NullableUserHourlyTraffic) IsSet() bool {
	return v.isSet
}

func (v *NullableUserHourlyTraffic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserHourlyTraffic(val *UserHourlyTraffic) *NullableUserHourlyTraffic {
	return &NullableUserHourlyTraffic{value: val, isSet: true}
}

func (v NullableUserHourlyTraffic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserHourlyTraffic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}