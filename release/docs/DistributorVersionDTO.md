# DistributorVersionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Capabilities** | Pointer to [**[]DistributorVersionCapabilityDTO**](DistributorVersionCapabilityDTO.md) |  | [optional] 

## Methods

### NewDistributorVersionDTO

`func NewDistributorVersionDTO() *DistributorVersionDTO`

NewDistributorVersionDTO instantiates a new DistributorVersionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributorVersionDTOWithDefaults

`func NewDistributorVersionDTOWithDefaults() *DistributorVersionDTO`

NewDistributorVersionDTOWithDefaults instantiates a new DistributorVersionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DistributorVersionDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DistributorVersionDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DistributorVersionDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DistributorVersionDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *DistributorVersionDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *DistributorVersionDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *DistributorVersionDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *DistributorVersionDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *DistributorVersionDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DistributorVersionDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DistributorVersionDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DistributorVersionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DistributorVersionDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DistributorVersionDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DistributorVersionDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DistributorVersionDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *DistributorVersionDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DistributorVersionDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DistributorVersionDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DistributorVersionDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCapabilities

`func (o *DistributorVersionDTO) GetCapabilities() []DistributorVersionCapabilityDTO`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DistributorVersionDTO) GetCapabilitiesOk() (*[]DistributorVersionCapabilityDTO, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DistributorVersionDTO) SetCapabilities(v []DistributorVersionCapabilityDTO)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *DistributorVersionDTO) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


