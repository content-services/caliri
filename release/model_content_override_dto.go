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

// checks if the ContentOverrideDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentOverrideDTO{}

// ContentOverrideDTO Represents a content override for an activation key or consumer
type ContentOverrideDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Name *string `json:"name,omitempty"`
	ContentLabel *string `json:"contentLabel,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewContentOverrideDTO instantiates a new ContentOverrideDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentOverrideDTO() *ContentOverrideDTO {
	this := ContentOverrideDTO{}
	return &this
}

// NewContentOverrideDTOWithDefaults instantiates a new ContentOverrideDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentOverrideDTOWithDefaults() *ContentOverrideDTO {
	this := ContentOverrideDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ContentOverrideDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentOverrideDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ContentOverrideDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ContentOverrideDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ContentOverrideDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentOverrideDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ContentOverrideDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *ContentOverrideDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContentOverrideDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentOverrideDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContentOverrideDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContentOverrideDTO) SetName(v string) {
	o.Name = &v
}

// GetContentLabel returns the ContentLabel field value if set, zero value otherwise.
func (o *ContentOverrideDTO) GetContentLabel() string {
	if o == nil || IsNil(o.ContentLabel) {
		var ret string
		return ret
	}
	return *o.ContentLabel
}

// GetContentLabelOk returns a tuple with the ContentLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentOverrideDTO) GetContentLabelOk() (*string, bool) {
	if o == nil || IsNil(o.ContentLabel) {
		return nil, false
	}
	return o.ContentLabel, true
}

// HasContentLabel returns a boolean if a field has been set.
func (o *ContentOverrideDTO) HasContentLabel() bool {
	if o != nil && !IsNil(o.ContentLabel) {
		return true
	}

	return false
}

// SetContentLabel gets a reference to the given string and assigns it to the ContentLabel field.
func (o *ContentOverrideDTO) SetContentLabel(v string) {
	o.ContentLabel = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ContentOverrideDTO) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentOverrideDTO) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ContentOverrideDTO) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ContentOverrideDTO) SetValue(v string) {
	o.Value = &v
}

func (o ContentOverrideDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentOverrideDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ContentLabel) {
		toSerialize["contentLabel"] = o.ContentLabel
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableContentOverrideDTO struct {
	value *ContentOverrideDTO
	isSet bool
}

func (v NullableContentOverrideDTO) Get() *ContentOverrideDTO {
	return v.value
}

func (v *NullableContentOverrideDTO) Set(val *ContentOverrideDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableContentOverrideDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableContentOverrideDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentOverrideDTO(val *ContentOverrideDTO) *NullableContentOverrideDTO {
	return &NullableContentOverrideDTO{value: val, isSet: true}
}

func (v NullableContentOverrideDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentOverrideDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


