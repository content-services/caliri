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

// checks if the GuestIdDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuestIdDTO{}

// GuestIdDTO Represents a guest ID running on a virt host consumer
type GuestIdDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	GuestId string `json:"guestId"`
	Attributes *map[string]string `json:"attributes,omitempty"`
}

type _GuestIdDTO GuestIdDTO

// NewGuestIdDTO instantiates a new GuestIdDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuestIdDTO(guestId string) *GuestIdDTO {
	this := GuestIdDTO{}
	this.GuestId = guestId
	return &this
}

// NewGuestIdDTOWithDefaults instantiates a new GuestIdDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuestIdDTOWithDefaults() *GuestIdDTO {
	this := GuestIdDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GuestIdDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestIdDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GuestIdDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *GuestIdDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *GuestIdDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestIdDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *GuestIdDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *GuestIdDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GuestIdDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestIdDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GuestIdDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GuestIdDTO) SetId(v string) {
	o.Id = &v
}

// GetGuestId returns the GuestId field value
func (o *GuestIdDTO) GetGuestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuestId
}

// GetGuestIdOk returns a tuple with the GuestId field value
// and a boolean to check if the value has been set.
func (o *GuestIdDTO) GetGuestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuestId, true
}

// SetGuestId sets field value
func (o *GuestIdDTO) SetGuestId(v string) {
	o.GuestId = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GuestIdDTO) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestIdDTO) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GuestIdDTO) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *GuestIdDTO) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

func (o GuestIdDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuestIdDTO) ToMap() (map[string]interface{}, error) {
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
	toSerialize["guestId"] = o.GuestId
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *GuestIdDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"guestId",
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

	varGuestIdDTO := _GuestIdDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuestIdDTO)

	if err != nil {
		return err
	}

	*o = GuestIdDTO(varGuestIdDTO)

	return err
}

type NullableGuestIdDTO struct {
	value *GuestIdDTO
	isSet bool
}

func (v NullableGuestIdDTO) Get() *GuestIdDTO {
	return v.value
}

func (v *NullableGuestIdDTO) Set(val *GuestIdDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestIdDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestIdDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestIdDTO(val *GuestIdDTO) *NullableGuestIdDTO {
	return &NullableGuestIdDTO{value: val, isSet: true}
}

func (v NullableGuestIdDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestIdDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


