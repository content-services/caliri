# HypervisorConsumerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NestedOwnerDTO**](NestedOwnerDTO.md) |  | [optional] 

## Methods

### NewHypervisorConsumerDTO

`func NewHypervisorConsumerDTO() *HypervisorConsumerDTO`

NewHypervisorConsumerDTO instantiates a new HypervisorConsumerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorConsumerDTOWithDefaults

`func NewHypervisorConsumerDTOWithDefaults() *HypervisorConsumerDTO`

NewHypervisorConsumerDTOWithDefaults instantiates a new HypervisorConsumerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *HypervisorConsumerDTO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HypervisorConsumerDTO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HypervisorConsumerDTO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HypervisorConsumerDTO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *HypervisorConsumerDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HypervisorConsumerDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HypervisorConsumerDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HypervisorConsumerDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *HypervisorConsumerDTO) GetOwner() NestedOwnerDTO`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *HypervisorConsumerDTO) GetOwnerOk() (*NestedOwnerDTO, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *HypervisorConsumerDTO) SetOwner(v NestedOwnerDTO)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *HypervisorConsumerDTO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


