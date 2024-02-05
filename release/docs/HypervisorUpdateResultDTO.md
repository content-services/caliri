# HypervisorUpdateResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**[]HypervisorConsumerDTO**](HypervisorConsumerDTO.md) |  | 
**Updated** | [**[]HypervisorConsumerDTO**](HypervisorConsumerDTO.md) |  | 
**Unchanged** | [**[]HypervisorConsumerDTO**](HypervisorConsumerDTO.md) |  | 
**FailedUpdate** | **[]string** |  | 

## Methods

### NewHypervisorUpdateResultDTO

`func NewHypervisorUpdateResultDTO(created []HypervisorConsumerDTO, updated []HypervisorConsumerDTO, unchanged []HypervisorConsumerDTO, failedUpdate []string, ) *HypervisorUpdateResultDTO`

NewHypervisorUpdateResultDTO instantiates a new HypervisorUpdateResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorUpdateResultDTOWithDefaults

`func NewHypervisorUpdateResultDTOWithDefaults() *HypervisorUpdateResultDTO`

NewHypervisorUpdateResultDTOWithDefaults instantiates a new HypervisorUpdateResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *HypervisorUpdateResultDTO) GetCreated() []HypervisorConsumerDTO`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *HypervisorUpdateResultDTO) GetCreatedOk() (*[]HypervisorConsumerDTO, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *HypervisorUpdateResultDTO) SetCreated(v []HypervisorConsumerDTO)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *HypervisorUpdateResultDTO) GetUpdated() []HypervisorConsumerDTO`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *HypervisorUpdateResultDTO) GetUpdatedOk() (*[]HypervisorConsumerDTO, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *HypervisorUpdateResultDTO) SetUpdated(v []HypervisorConsumerDTO)`

SetUpdated sets Updated field to given value.


### GetUnchanged

`func (o *HypervisorUpdateResultDTO) GetUnchanged() []HypervisorConsumerDTO`

GetUnchanged returns the Unchanged field if non-nil, zero value otherwise.

### GetUnchangedOk

`func (o *HypervisorUpdateResultDTO) GetUnchangedOk() (*[]HypervisorConsumerDTO, bool)`

GetUnchangedOk returns a tuple with the Unchanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnchanged

`func (o *HypervisorUpdateResultDTO) SetUnchanged(v []HypervisorConsumerDTO)`

SetUnchanged sets Unchanged field to given value.


### GetFailedUpdate

`func (o *HypervisorUpdateResultDTO) GetFailedUpdate() []string`

GetFailedUpdate returns the FailedUpdate field if non-nil, zero value otherwise.

### GetFailedUpdateOk

`func (o *HypervisorUpdateResultDTO) GetFailedUpdateOk() (*[]string, bool)`

GetFailedUpdateOk returns a tuple with the FailedUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedUpdate

`func (o *HypervisorUpdateResultDTO) SetFailedUpdate(v []string)`

SetFailedUpdate sets FailedUpdate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


