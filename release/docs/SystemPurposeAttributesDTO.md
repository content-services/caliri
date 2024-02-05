# SystemPurposeAttributesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to [**NestedOwnerDTO**](NestedOwnerDTO.md) |  | [optional] 
**SystemPurposeAttributes** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewSystemPurposeAttributesDTO

`func NewSystemPurposeAttributesDTO() *SystemPurposeAttributesDTO`

NewSystemPurposeAttributesDTO instantiates a new SystemPurposeAttributesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemPurposeAttributesDTOWithDefaults

`func NewSystemPurposeAttributesDTOWithDefaults() *SystemPurposeAttributesDTO`

NewSystemPurposeAttributesDTOWithDefaults instantiates a new SystemPurposeAttributesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *SystemPurposeAttributesDTO) GetOwner() NestedOwnerDTO`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SystemPurposeAttributesDTO) GetOwnerOk() (*NestedOwnerDTO, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SystemPurposeAttributesDTO) SetOwner(v NestedOwnerDTO)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SystemPurposeAttributesDTO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSystemPurposeAttributes

`func (o *SystemPurposeAttributesDTO) GetSystemPurposeAttributes() map[string][]string`

GetSystemPurposeAttributes returns the SystemPurposeAttributes field if non-nil, zero value otherwise.

### GetSystemPurposeAttributesOk

`func (o *SystemPurposeAttributesDTO) GetSystemPurposeAttributesOk() (*map[string][]string, bool)`

GetSystemPurposeAttributesOk returns a tuple with the SystemPurposeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPurposeAttributes

`func (o *SystemPurposeAttributesDTO) SetSystemPurposeAttributes(v map[string][]string)`

SetSystemPurposeAttributes sets SystemPurposeAttributes field to given value.

### HasSystemPurposeAttributes

`func (o *SystemPurposeAttributesDTO) HasSystemPurposeAttributes() bool`

HasSystemPurposeAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


