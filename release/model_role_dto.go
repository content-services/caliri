/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the RoleDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleDTO{}

// RoleDTO Represent the relationship between users and the permissions they have
type RoleDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Users []UserDTO `json:"users,omitempty"`
	Permissions []PermissionBlueprintDTO `json:"permissions,omitempty"`
}

// NewRoleDTO instantiates a new RoleDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleDTO() *RoleDTO {
	this := RoleDTO{}
	return &this
}

// NewRoleDTOWithDefaults instantiates a new RoleDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleDTOWithDefaults() *RoleDTO {
	this := RoleDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RoleDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RoleDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *RoleDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *RoleDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *RoleDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *RoleDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleDTO) SetName(v string) {
	o.Name = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *RoleDTO) GetUsers() []UserDTO {
	if o == nil || IsNil(o.Users) {
		var ret []UserDTO
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleDTO) GetUsersOk() ([]UserDTO, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *RoleDTO) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserDTO and assigns it to the Users field.
func (o *RoleDTO) SetUsers(v []UserDTO) {
	o.Users = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RoleDTO) GetPermissions() []PermissionBlueprintDTO {
	if o == nil || IsNil(o.Permissions) {
		var ret []PermissionBlueprintDTO
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleDTO) GetPermissionsOk() ([]PermissionBlueprintDTO, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RoleDTO) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []PermissionBlueprintDTO and assigns it to the Permissions field.
func (o *RoleDTO) SetPermissions(v []PermissionBlueprintDTO) {
	o.Permissions = v
}

func (o RoleDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableRoleDTO struct {
	value *RoleDTO
	isSet bool
}

func (v NullableRoleDTO) Get() *RoleDTO {
	return v.value
}

func (v *NullableRoleDTO) Set(val *RoleDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleDTO(val *RoleDTO) *NullableRoleDTO {
	return &NullableRoleDTO{value: val, isSet: true}
}

func (v NullableRoleDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


