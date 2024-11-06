/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CapabilityDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityDTO{}

// CapabilityDTO Represents a consumer capability
type CapabilityDTO struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
}

type _CapabilityDTO CapabilityDTO

// NewCapabilityDTO instantiates a new CapabilityDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityDTO(name string) *CapabilityDTO {
	this := CapabilityDTO{}
	this.Name = name
	return &this
}

// NewCapabilityDTOWithDefaults instantiates a new CapabilityDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityDTOWithDefaults() *CapabilityDTO {
	this := CapabilityDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CapabilityDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CapabilityDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CapabilityDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *CapabilityDTO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CapabilityDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CapabilityDTO) SetName(v string) {
	o.Name = v
}

func (o CapabilityDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *CapabilityDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCapabilityDTO := _CapabilityDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCapabilityDTO)

	if err != nil {
		return err
	}

	*o = CapabilityDTO(varCapabilityDTO)

	return err
}

type NullableCapabilityDTO struct {
	value *CapabilityDTO
	isSet bool
}

func (v NullableCapabilityDTO) Get() *CapabilityDTO {
	return v.value
}

func (v *NullableCapabilityDTO) Set(val *CapabilityDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityDTO(val *CapabilityDTO) *NullableCapabilityDTO {
	return &NullableCapabilityDTO{value: val, isSet: true}
}

func (v NullableCapabilityDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


