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

// checks if the ContentToPromoteDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentToPromoteDTO{}

// ContentToPromoteDTO Represents the JSON input when promoting content, or updating promoted content.
type ContentToPromoteDTO struct {
	EnvironmentId *string `json:"environmentId,omitempty"`
	ContentId *string `json:"contentId,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewContentToPromoteDTO instantiates a new ContentToPromoteDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentToPromoteDTO() *ContentToPromoteDTO {
	this := ContentToPromoteDTO{}
	return &this
}

// NewContentToPromoteDTOWithDefaults instantiates a new ContentToPromoteDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentToPromoteDTOWithDefaults() *ContentToPromoteDTO {
	this := ContentToPromoteDTO{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ContentToPromoteDTO) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentToPromoteDTO) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ContentToPromoteDTO) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *ContentToPromoteDTO) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetContentId returns the ContentId field value if set, zero value otherwise.
func (o *ContentToPromoteDTO) GetContentId() string {
	if o == nil || IsNil(o.ContentId) {
		var ret string
		return ret
	}
	return *o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentToPromoteDTO) GetContentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContentId) {
		return nil, false
	}
	return o.ContentId, true
}

// HasContentId returns a boolean if a field has been set.
func (o *ContentToPromoteDTO) HasContentId() bool {
	if o != nil && !IsNil(o.ContentId) {
		return true
	}

	return false
}

// SetContentId gets a reference to the given string and assigns it to the ContentId field.
func (o *ContentToPromoteDTO) SetContentId(v string) {
	o.ContentId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ContentToPromoteDTO) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentToPromoteDTO) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ContentToPromoteDTO) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ContentToPromoteDTO) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ContentToPromoteDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentToPromoteDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if !IsNil(o.ContentId) {
		toSerialize["contentId"] = o.ContentId
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableContentToPromoteDTO struct {
	value *ContentToPromoteDTO
	isSet bool
}

func (v NullableContentToPromoteDTO) Get() *ContentToPromoteDTO {
	return v.value
}

func (v *NullableContentToPromoteDTO) Set(val *ContentToPromoteDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableContentToPromoteDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableContentToPromoteDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentToPromoteDTO(val *ContentToPromoteDTO) *NullableContentToPromoteDTO {
	return &NullableContentToPromoteDTO{value: val, isSet: true}
}

func (v NullableContentToPromoteDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentToPromoteDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


