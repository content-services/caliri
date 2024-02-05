# PermissionBlueprintDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NestedOwnerDTO**](NestedOwnerDTO.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Access** | Pointer to **string** |  | [optional] 

## Methods

### NewPermissionBlueprintDTO

`func NewPermissionBlueprintDTO() *PermissionBlueprintDTO`

NewPermissionBlueprintDTO instantiates a new PermissionBlueprintDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionBlueprintDTOWithDefaults

`func NewPermissionBlueprintDTOWithDefaults() *PermissionBlueprintDTO`

NewPermissionBlueprintDTOWithDefaults instantiates a new PermissionBlueprintDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PermissionBlueprintDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PermissionBlueprintDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PermissionBlueprintDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PermissionBlueprintDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *PermissionBlueprintDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PermissionBlueprintDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PermissionBlueprintDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PermissionBlueprintDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *PermissionBlueprintDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionBlueprintDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionBlueprintDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PermissionBlueprintDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwner

`func (o *PermissionBlueprintDTO) GetOwner() NestedOwnerDTO`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PermissionBlueprintDTO) GetOwnerOk() (*NestedOwnerDTO, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PermissionBlueprintDTO) SetOwner(v NestedOwnerDTO)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PermissionBlueprintDTO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetType

`func (o *PermissionBlueprintDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PermissionBlueprintDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PermissionBlueprintDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PermissionBlueprintDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccess

`func (o *PermissionBlueprintDTO) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *PermissionBlueprintDTO) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *PermissionBlueprintDTO) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *PermissionBlueprintDTO) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


