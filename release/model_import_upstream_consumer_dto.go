/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the ImportUpstreamConsumerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportUpstreamConsumerDTO{}

// ImportUpstreamConsumerDTO Represents an import upstream consumer
type ImportUpstreamConsumerDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Name *string `json:"name,omitempty"`
	ApiUrl *string `json:"apiUrl,omitempty"`
	WebUrl *string `json:"webUrl,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	ContentAccessMode *string `json:"contentAccessMode,omitempty"`
	Type *ConsumerTypeDTO `json:"type,omitempty"`
}

// NewImportUpstreamConsumerDTO instantiates a new ImportUpstreamConsumerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportUpstreamConsumerDTO() *ImportUpstreamConsumerDTO {
	this := ImportUpstreamConsumerDTO{}
	return &this
}

// NewImportUpstreamConsumerDTOWithDefaults instantiates a new ImportUpstreamConsumerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportUpstreamConsumerDTOWithDefaults() *ImportUpstreamConsumerDTO {
	this := ImportUpstreamConsumerDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ImportUpstreamConsumerDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUpstreamConsumerDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ImportUpstreamConsumerDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ImportUpstreamConsumerDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ImportUpstreamConsumerDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUpstreamConsumerDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ImportUpstreamConsumerDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *ImportUpstreamConsumerDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportUpstreamConsumerDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUpstreamConsumerDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportUpstreamConsumerDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImportUpstreamConsumerDTO) SetId(v string) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ImportUpstreamConsumerDTO) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUpstreamConsumerDTO) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ImportUpstreamConsumerDTO) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ImportUpstreamConsumerDTO) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImportUpstreamConsumerDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUpstreamConsumerDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImportUpstreamConsumerDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImportUpstreamConsumerDTO) SetName(v string) {
	o.Name = &v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *ImportUpstreamConsumerDTO) GetApiUrl() string {
	if o == nil || IsNil(o.ApiUrl) {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUpstreamConsumerDTO) GetApiUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ApiUrl) {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *ImportUpstreamConsumerDTO) HasApiUrl() bool {
	if o != nil && !IsNil(o.ApiUrl) {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *ImportUpstreamConsumerDTO) SetApiUrl(v string) {
	o.ApiUrl = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *ImportUpstreamConsumerDTO) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl) {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUpstreamConsumerDTO) GetWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebUrl) {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *ImportUpstreamConsumerDTO) HasWebUrl() bool {
	if o != nil && !IsNil(o.WebUrl) {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *ImportUpstreamConsumerDTO) SetWebUrl(v string) {
	o.WebUrl = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ImportUpstreamConsumerDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUpstreamConsumerDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ImportUpstreamConsumerDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ImportUpstreamConsumerDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetContentAccessMode returns the ContentAccessMode field value if set, zero value otherwise.
func (o *ImportUpstreamConsumerDTO) GetContentAccessMode() string {
	if o == nil || IsNil(o.ContentAccessMode) {
		var ret string
		return ret
	}
	return *o.ContentAccessMode
}

// GetContentAccessModeOk returns a tuple with the ContentAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUpstreamConsumerDTO) GetContentAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentAccessMode) {
		return nil, false
	}
	return o.ContentAccessMode, true
}

// HasContentAccessMode returns a boolean if a field has been set.
func (o *ImportUpstreamConsumerDTO) HasContentAccessMode() bool {
	if o != nil && !IsNil(o.ContentAccessMode) {
		return true
	}

	return false
}

// SetContentAccessMode gets a reference to the given string and assigns it to the ContentAccessMode field.
func (o *ImportUpstreamConsumerDTO) SetContentAccessMode(v string) {
	o.ContentAccessMode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ImportUpstreamConsumerDTO) GetType() ConsumerTypeDTO {
	if o == nil || IsNil(o.Type) {
		var ret ConsumerTypeDTO
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUpstreamConsumerDTO) GetTypeOk() (*ConsumerTypeDTO, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ImportUpstreamConsumerDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ConsumerTypeDTO and assigns it to the Type field.
func (o *ImportUpstreamConsumerDTO) SetType(v ConsumerTypeDTO) {
	o.Type = &v
}

func (o ImportUpstreamConsumerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportUpstreamConsumerDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ApiUrl) {
		toSerialize["apiUrl"] = o.ApiUrl
	}
	if !IsNil(o.WebUrl) {
		toSerialize["webUrl"] = o.WebUrl
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.ContentAccessMode) {
		toSerialize["contentAccessMode"] = o.ContentAccessMode
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableImportUpstreamConsumerDTO struct {
	value *ImportUpstreamConsumerDTO
	isSet bool
}

func (v NullableImportUpstreamConsumerDTO) Get() *ImportUpstreamConsumerDTO {
	return v.value
}

func (v *NullableImportUpstreamConsumerDTO) Set(val *ImportUpstreamConsumerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableImportUpstreamConsumerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableImportUpstreamConsumerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportUpstreamConsumerDTO(val *ImportUpstreamConsumerDTO) *NullableImportUpstreamConsumerDTO {
	return &NullableImportUpstreamConsumerDTO{value: val, isSet: true}
}

func (v NullableImportUpstreamConsumerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportUpstreamConsumerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


