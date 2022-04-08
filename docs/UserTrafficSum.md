# UserTrafficSum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThisMonth** | Pointer to **int64** | Summed traffic for currentMonth. | [optional] 
**LastMonth** | Pointer to **int64** | Summed traffic for lastMonth. | [optional] 

## Methods

### NewUserTrafficSum

`func NewUserTrafficSum() *UserTrafficSum`

NewUserTrafficSum instantiates a new UserTrafficSum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTrafficSumWithDefaults

`func NewUserTrafficSumWithDefaults() *UserTrafficSum`

NewUserTrafficSumWithDefaults instantiates a new UserTrafficSum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThisMonth

`func (o *UserTrafficSum) GetThisMonth() int64`

GetThisMonth returns the ThisMonth field if non-nil, zero value otherwise.

### GetThisMonthOk

`func (o *UserTrafficSum) GetThisMonthOk() (*int64, bool)`

GetThisMonthOk returns a tuple with the ThisMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThisMonth

`func (o *UserTrafficSum) SetThisMonth(v int64)`

SetThisMonth sets ThisMonth field to given value.

### HasThisMonth

`func (o *UserTrafficSum) HasThisMonth() bool`

HasThisMonth returns a boolean if a field has been set.

### GetLastMonth

`func (o *UserTrafficSum) GetLastMonth() int64`

GetLastMonth returns the LastMonth field if non-nil, zero value otherwise.

### GetLastMonthOk

`func (o *UserTrafficSum) GetLastMonthOk() (*int64, bool)`

GetLastMonthOk returns a tuple with the LastMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMonth

`func (o *UserTrafficSum) SetLastMonth(v int64)`

SetLastMonth sets LastMonth field to given value.

### HasLastMonth

`func (o *UserTrafficSum) HasLastMonth() bool`

HasLastMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


