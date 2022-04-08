# TxDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxDetail** | Pointer to [**BlockTx**](BlockTx.md) |  | [optional] 
**Inputs** | Pointer to [**[]TxInput**](TxInput.md) |  | [optional] 
**Outputs** | Pointer to [**[]TxOutput**](TxOutput.md) |  | [optional] 

## Methods

### NewTxDetail

`func NewTxDetail() *TxDetail`

NewTxDetail instantiates a new TxDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxDetailWithDefaults

`func NewTxDetailWithDefaults() *TxDetail`

NewTxDetailWithDefaults instantiates a new TxDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxDetail

`func (o *TxDetail) GetTxDetail() BlockTx`

GetTxDetail returns the TxDetail field if non-nil, zero value otherwise.

### GetTxDetailOk

`func (o *TxDetail) GetTxDetailOk() (*BlockTx, bool)`

GetTxDetailOk returns a tuple with the TxDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDetail

`func (o *TxDetail) SetTxDetail(v BlockTx)`

SetTxDetail sets TxDetail field to given value.

### HasTxDetail

`func (o *TxDetail) HasTxDetail() bool`

HasTxDetail returns a boolean if a field has been set.

### GetInputs

`func (o *TxDetail) GetInputs() []TxInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *TxDetail) GetInputsOk() (*[]TxInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *TxDetail) SetInputs(v []TxInput)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *TxDetail) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetOutputs

`func (o *TxDetail) GetOutputs() []TxOutput`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *TxDetail) GetOutputsOk() (*[]TxOutput, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *TxDetail) SetOutputs(v []TxOutput)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *TxDetail) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


