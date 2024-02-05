# ActivationKeyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NestedOwnerDTO**](NestedOwnerDTO.md) |  | [optional] 
**ReleaseVer** | Pointer to [**ReleaseVerDTO**](ReleaseVerDTO.md) |  | [optional] 
**ServiceLevel** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**AddOns** | Pointer to **[]string** |  | [optional] 
**AutoAttach** | Pointer to **bool** |  | [optional] 
**Pools** | Pointer to [**[]ActivationKeyPoolDTO**](ActivationKeyPoolDTO.md) |  | [optional] 
**Products** | Pointer to [**[]ActivationKeyProductDTO**](ActivationKeyProductDTO.md) |  | [optional] 
**ContentOverrides** | Pointer to [**[]ContentOverrideDTO**](ContentOverrideDTO.md) |  | [optional] 

## Methods

### NewActivationKeyDTO

`func NewActivationKeyDTO() *ActivationKeyDTO`

NewActivationKeyDTO instantiates a new ActivationKeyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivationKeyDTOWithDefaults

`func NewActivationKeyDTOWithDefaults() *ActivationKeyDTO`

NewActivationKeyDTOWithDefaults instantiates a new ActivationKeyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ActivationKeyDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ActivationKeyDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ActivationKeyDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ActivationKeyDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ActivationKeyDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ActivationKeyDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ActivationKeyDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ActivationKeyDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *ActivationKeyDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivationKeyDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivationKeyDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivationKeyDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActivationKeyDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivationKeyDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivationKeyDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActivationKeyDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ActivationKeyDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActivationKeyDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActivationKeyDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActivationKeyDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *ActivationKeyDTO) GetOwner() NestedOwnerDTO`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ActivationKeyDTO) GetOwnerOk() (*NestedOwnerDTO, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ActivationKeyDTO) SetOwner(v NestedOwnerDTO)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ActivationKeyDTO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetReleaseVer

`func (o *ActivationKeyDTO) GetReleaseVer() ReleaseVerDTO`

GetReleaseVer returns the ReleaseVer field if non-nil, zero value otherwise.

### GetReleaseVerOk

`func (o *ActivationKeyDTO) GetReleaseVerOk() (*ReleaseVerDTO, bool)`

GetReleaseVerOk returns a tuple with the ReleaseVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVer

`func (o *ActivationKeyDTO) SetReleaseVer(v ReleaseVerDTO)`

SetReleaseVer sets ReleaseVer field to given value.

### HasReleaseVer

`func (o *ActivationKeyDTO) HasReleaseVer() bool`

HasReleaseVer returns a boolean if a field has been set.

### GetServiceLevel

`func (o *ActivationKeyDTO) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *ActivationKeyDTO) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *ActivationKeyDTO) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *ActivationKeyDTO) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetUsage

`func (o *ActivationKeyDTO) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ActivationKeyDTO) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ActivationKeyDTO) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ActivationKeyDTO) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetRole

`func (o *ActivationKeyDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ActivationKeyDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ActivationKeyDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ActivationKeyDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAddOns

`func (o *ActivationKeyDTO) GetAddOns() []string`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *ActivationKeyDTO) GetAddOnsOk() (*[]string, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *ActivationKeyDTO) SetAddOns(v []string)`

SetAddOns sets AddOns field to given value.

### HasAddOns

`func (o *ActivationKeyDTO) HasAddOns() bool`

HasAddOns returns a boolean if a field has been set.

### GetAutoAttach

`func (o *ActivationKeyDTO) GetAutoAttach() bool`

GetAutoAttach returns the AutoAttach field if non-nil, zero value otherwise.

### GetAutoAttachOk

`func (o *ActivationKeyDTO) GetAutoAttachOk() (*bool, bool)`

GetAutoAttachOk returns a tuple with the AutoAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAttach

`func (o *ActivationKeyDTO) SetAutoAttach(v bool)`

SetAutoAttach sets AutoAttach field to given value.

### HasAutoAttach

`func (o *ActivationKeyDTO) HasAutoAttach() bool`

HasAutoAttach returns a boolean if a field has been set.

### GetPools

`func (o *ActivationKeyDTO) GetPools() []ActivationKeyPoolDTO`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *ActivationKeyDTO) GetPoolsOk() (*[]ActivationKeyPoolDTO, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *ActivationKeyDTO) SetPools(v []ActivationKeyPoolDTO)`

SetPools sets Pools field to given value.

### HasPools

`func (o *ActivationKeyDTO) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetProducts

`func (o *ActivationKeyDTO) GetProducts() []ActivationKeyProductDTO`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ActivationKeyDTO) GetProductsOk() (*[]ActivationKeyProductDTO, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ActivationKeyDTO) SetProducts(v []ActivationKeyProductDTO)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ActivationKeyDTO) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetContentOverrides

`func (o *ActivationKeyDTO) GetContentOverrides() []ContentOverrideDTO`

GetContentOverrides returns the ContentOverrides field if non-nil, zero value otherwise.

### GetContentOverridesOk

`func (o *ActivationKeyDTO) GetContentOverridesOk() (*[]ContentOverrideDTO, bool)`

GetContentOverridesOk returns a tuple with the ContentOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentOverrides

`func (o *ActivationKeyDTO) SetContentOverrides(v []ContentOverrideDTO)`

SetContentOverrides sets ContentOverrides field to given value.

### HasContentOverrides

`func (o *ActivationKeyDTO) HasContentOverrides() bool`

HasContentOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


