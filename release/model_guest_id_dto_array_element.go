/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.19
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GuestIdDTOArrayElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuestIdDTOArrayElement{}

// GuestIdDTOArrayElement Represents a guest ID running on a virt host consumer. Does not include the attributes field
type GuestIdDTOArrayElement struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	GuestId string `json:"guestId"`
}

type _GuestIdDTOArrayElement GuestIdDTOArrayElement

// NewGuestIdDTOArrayElement instantiates a new GuestIdDTOArrayElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuestIdDTOArrayElement(guestId string) *GuestIdDTOArrayElement {
	this := GuestIdDTOArrayElement{}
	this.GuestId = guestId
	return &this
}

// NewGuestIdDTOArrayElementWithDefaults instantiates a new GuestIdDTOArrayElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuestIdDTOArrayElementWithDefaults() *GuestIdDTOArrayElement {
	this := GuestIdDTOArrayElement{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GuestIdDTOArrayElement) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestIdDTOArrayElement) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GuestIdDTOArrayElement) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *GuestIdDTOArrayElement) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *GuestIdDTOArrayElement) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestIdDTOArrayElement) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *GuestIdDTOArrayElement) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *GuestIdDTOArrayElement) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GuestIdDTOArrayElement) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestIdDTOArrayElement) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GuestIdDTOArrayElement) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GuestIdDTOArrayElement) SetId(v string) {
	o.Id = &v
}

// GetGuestId returns the GuestId field value
func (o *GuestIdDTOArrayElement) GetGuestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuestId
}

// GetGuestIdOk returns a tuple with the GuestId field value
// and a boolean to check if the value has been set.
func (o *GuestIdDTOArrayElement) GetGuestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuestId, true
}

// SetGuestId sets field value
func (o *GuestIdDTOArrayElement) SetGuestId(v string) {
	o.GuestId = v
}

func (o GuestIdDTOArrayElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuestIdDTOArrayElement) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *GuestIdDTOArrayElement) UnmarshalJSON(data []byte) (err error) {
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

	varGuestIdDTOArrayElement := _GuestIdDTOArrayElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuestIdDTOArrayElement)

	if err != nil {
		return err
	}

	*o = GuestIdDTOArrayElement(varGuestIdDTOArrayElement)

	return err
}

type NullableGuestIdDTOArrayElement struct {
	value *GuestIdDTOArrayElement
	isSet bool
}

func (v NullableGuestIdDTOArrayElement) Get() *GuestIdDTOArrayElement {
	return v.value
}

func (v *NullableGuestIdDTOArrayElement) Set(val *GuestIdDTOArrayElement) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestIdDTOArrayElement) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestIdDTOArrayElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestIdDTOArrayElement(val *GuestIdDTOArrayElement) *NullableGuestIdDTOArrayElement {
	return &NullableGuestIdDTOArrayElement{value: val, isSet: true}
}

func (v NullableGuestIdDTOArrayElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestIdDTOArrayElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


