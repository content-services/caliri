# EnvironmentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ContentPrefix** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NestedOwnerDTO**](NestedOwnerDTO.md) |  | [optional] 
**EnvironmentContent** | Pointer to [**[]EnvironmentContentDTO**](EnvironmentContentDTO.md) |  | [optional] 

## Methods

### NewEnvironmentDTO

`func NewEnvironmentDTO() *EnvironmentDTO`

NewEnvironmentDTO instantiates a new EnvironmentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDTOWithDefaults

`func NewEnvironmentDTOWithDefaults() *EnvironmentDTO`

NewEnvironmentDTOWithDefaults instantiates a new EnvironmentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *EnvironmentDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EnvironmentDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EnvironmentDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EnvironmentDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *EnvironmentDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *EnvironmentDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *EnvironmentDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *EnvironmentDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContentPrefix

`func (o *EnvironmentDTO) GetContentPrefix() string`

GetContentPrefix returns the ContentPrefix field if non-nil, zero value otherwise.

### GetContentPrefixOk

`func (o *EnvironmentDTO) GetContentPrefixOk() (*string, bool)`

GetContentPrefixOk returns a tuple with the ContentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPrefix

`func (o *EnvironmentDTO) SetContentPrefix(v string)`

SetContentPrefix sets ContentPrefix field to given value.

### HasContentPrefix

`func (o *EnvironmentDTO) HasContentPrefix() bool`

HasContentPrefix returns a boolean if a field has been set.

### GetOwner

`func (o *EnvironmentDTO) GetOwner() NestedOwnerDTO`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *EnvironmentDTO) GetOwnerOk() (*NestedOwnerDTO, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *EnvironmentDTO) SetOwner(v NestedOwnerDTO)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *EnvironmentDTO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetEnvironmentContent

`func (o *EnvironmentDTO) GetEnvironmentContent() []EnvironmentContentDTO`

GetEnvironmentContent returns the EnvironmentContent field if non-nil, zero value otherwise.

### GetEnvironmentContentOk

`func (o *EnvironmentDTO) GetEnvironmentContentOk() (*[]EnvironmentContentDTO, bool)`

GetEnvironmentContentOk returns a tuple with the EnvironmentContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentContent

`func (o *EnvironmentDTO) SetEnvironmentContent(v []EnvironmentContentDTO)`

SetEnvironmentContent sets EnvironmentContent field to given value.

### HasEnvironmentContent

`func (o *EnvironmentDTO) HasEnvironmentContent() bool`

HasEnvironmentContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


