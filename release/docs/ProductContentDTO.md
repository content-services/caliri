# ProductContentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**ContentDTO**](ContentDTO.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewProductContentDTO

`func NewProductContentDTO(content ContentDTO, ) *ProductContentDTO`

NewProductContentDTO instantiates a new ProductContentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductContentDTOWithDefaults

`func NewProductContentDTOWithDefaults() *ProductContentDTO`

NewProductContentDTOWithDefaults instantiates a new ProductContentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ProductContentDTO) GetContent() ContentDTO`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ProductContentDTO) GetContentOk() (*ContentDTO, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ProductContentDTO) SetContent(v ContentDTO)`

SetContent sets Content field to given value.


### GetEnabled

`func (o *ProductContentDTO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProductContentDTO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProductContentDTO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ProductContentDTO) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


