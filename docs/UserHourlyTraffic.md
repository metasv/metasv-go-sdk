# UserHourlyTraffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HourTime** | Pointer to **int64** | The unix timestamp of the beginning of the hour. | [optional] 
**HourTimeString** | Pointer to **string** | Parsed utc time string for readability. | [optional] 
**TotalBytes** | Pointer to **int64** | Total outbound bytes within this hour | [optional] 

## Methods

### NewUserHourlyTraffic

`func NewUserHourlyTraffic() *UserHourlyTraffic`

NewUserHourlyTraffic instantiates a new UserHourlyTraffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserHourlyTrafficWithDefaults

`func NewUserHourlyTrafficWithDefaults() *UserHourlyTraffic`

NewUserHourlyTrafficWithDefaults instantiates a new UserHourlyTraffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHourTime

`func (o *UserHourlyTraffic) GetHourTime() int64`

GetHourTime returns the HourTime field if non-nil, zero value otherwise.

### GetHourTimeOk

`func (o *UserHourlyTraffic) GetHourTimeOk() (*int64, bool)`

GetHourTimeOk returns a tuple with the HourTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourTime

`func (o *UserHourlyTraffic) SetHourTime(v int64)`

SetHourTime sets HourTime field to given value.

### HasHourTime

`func (o *UserHourlyTraffic) HasHourTime() bool`

HasHourTime returns a boolean if a field has been set.

### GetHourTimeString

`func (o *UserHourlyTraffic) GetHourTimeString() string`

GetHourTimeString returns the HourTimeString field if non-nil, zero value otherwise.

### GetHourTimeStringOk

`func (o *UserHourlyTraffic) GetHourTimeStringOk() (*string, bool)`

GetHourTimeStringOk returns a tuple with the HourTimeString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourTimeString

`func (o *UserHourlyTraffic) SetHourTimeString(v string)`

SetHourTimeString sets HourTimeString field to given value.

### HasHourTimeString

`func (o *UserHourlyTraffic) HasHourTimeString() bool`

HasHourTimeString returns a boolean if a field has been set.

### GetTotalBytes

`func (o *UserHourlyTraffic) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *UserHourlyTraffic) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *UserHourlyTraffic) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *UserHourlyTraffic) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


