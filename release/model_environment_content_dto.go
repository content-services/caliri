/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.19
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the EnvironmentContentDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentContentDTO{}

// EnvironmentContentDTO EnvironmentContent represents the promotion of content into a particular environment.
type EnvironmentContentDTO struct {
	ContentId *string `json:"contentId,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewEnvironmentContentDTO instantiates a new EnvironmentContentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentContentDTO() *EnvironmentContentDTO {
	this := EnvironmentContentDTO{}
	return &this
}

// NewEnvironmentContentDTOWithDefaults instantiates a new EnvironmentContentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentContentDTOWithDefaults() *EnvironmentContentDTO {
	this := EnvironmentContentDTO{}
	return &this
}

// GetContentId returns the ContentId field value if set, zero value otherwise.
func (o *EnvironmentContentDTO) GetContentId() string {
	if o == nil || IsNil(o.ContentId) {
		var ret string
		return ret
	}
	return *o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContentDTO) GetContentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContentId) {
		return nil, false
	}
	return o.ContentId, true
}

// HasContentId returns a boolean if a field has been set.
func (o *EnvironmentContentDTO) HasContentId() bool {
	if o != nil && !IsNil(o.ContentId) {
		return true
	}

	return false
}

// SetContentId gets a reference to the given string and assigns it to the ContentId field.
func (o *EnvironmentContentDTO) SetContentId(v string) {
	o.ContentId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EnvironmentContentDTO) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentContentDTO) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EnvironmentContentDTO) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EnvironmentContentDTO) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o EnvironmentContentDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentContentDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentId) {
		toSerialize["contentId"] = o.ContentId
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableEnvironmentContentDTO struct {
	value *EnvironmentContentDTO
	isSet bool
}

func (v NullableEnvironmentContentDTO) Get() *EnvironmentContentDTO {
	return v.value
}

func (v *NullableEnvironmentContentDTO) Set(val *EnvironmentContentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentContentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentContentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentContentDTO(val *EnvironmentContentDTO) *NullableEnvironmentContentDTO {
	return &NullableEnvironmentContentDTO{value: val, isSet: true}
}

func (v NullableEnvironmentContentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentContentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


