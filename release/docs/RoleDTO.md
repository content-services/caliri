# RoleDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]UserDTO**](UserDTO.md) |  | [optional] 
**Permissions** | Pointer to [**[]PermissionBlueprintDTO**](PermissionBlueprintDTO.md) |  | [optional] 

## Methods

### NewRoleDTO

`func NewRoleDTO() *RoleDTO`

NewRoleDTO instantiates a new RoleDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDTOWithDefaults

`func NewRoleDTOWithDefaults() *RoleDTO`

NewRoleDTOWithDefaults instantiates a new RoleDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *RoleDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RoleDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RoleDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RoleDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *RoleDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RoleDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RoleDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *RoleDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *RoleDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RoleDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsers

`func (o *RoleDTO) GetUsers() []UserDTO`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RoleDTO) GetUsersOk() (*[]UserDTO, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RoleDTO) SetUsers(v []UserDTO)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *RoleDTO) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetPermissions

`func (o *RoleDTO) GetPermissions() []PermissionBlueprintDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleDTO) GetPermissionsOk() (*[]PermissionBlueprintDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleDTO) SetPermissions(v []PermissionBlueprintDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RoleDTO) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


