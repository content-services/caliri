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

// checks if the ExceptionMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExceptionMessage{}

// ExceptionMessage An exception has occurred
type ExceptionMessage struct {
	DisplayMessage *string `json:"displayMessage,omitempty"`
	RequestUuid *string `json:"requestUuid,omitempty"`
}

// NewExceptionMessage instantiates a new ExceptionMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionMessage() *ExceptionMessage {
	this := ExceptionMessage{}
	return &this
}

// NewExceptionMessageWithDefaults instantiates a new ExceptionMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionMessageWithDefaults() *ExceptionMessage {
	this := ExceptionMessage{}
	return &this
}

// GetDisplayMessage returns the DisplayMessage field value if set, zero value otherwise.
func (o *ExceptionMessage) GetDisplayMessage() string {
	if o == nil || IsNil(o.DisplayMessage) {
		var ret string
		return ret
	}
	return *o.DisplayMessage
}

// GetDisplayMessageOk returns a tuple with the DisplayMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionMessage) GetDisplayMessageOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayMessage) {
		return nil, false
	}
	return o.DisplayMessage, true
}

// HasDisplayMessage returns a boolean if a field has been set.
func (o *ExceptionMessage) HasDisplayMessage() bool {
	if o != nil && !IsNil(o.DisplayMessage) {
		return true
	}

	return false
}

// SetDisplayMessage gets a reference to the given string and assigns it to the DisplayMessage field.
func (o *ExceptionMessage) SetDisplayMessage(v string) {
	o.DisplayMessage = &v
}

// GetRequestUuid returns the RequestUuid field value if set, zero value otherwise.
func (o *ExceptionMessage) GetRequestUuid() string {
	if o == nil || IsNil(o.RequestUuid) {
		var ret string
		return ret
	}
	return *o.RequestUuid
}

// GetRequestUuidOk returns a tuple with the RequestUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionMessage) GetRequestUuidOk() (*string, bool) {
	if o == nil || IsNil(o.RequestUuid) {
		return nil, false
	}
	return o.RequestUuid, true
}

// HasRequestUuid returns a boolean if a field has been set.
func (o *ExceptionMessage) HasRequestUuid() bool {
	if o != nil && !IsNil(o.RequestUuid) {
		return true
	}

	return false
}

// SetRequestUuid gets a reference to the given string and assigns it to the RequestUuid field.
func (o *ExceptionMessage) SetRequestUuid(v string) {
	o.RequestUuid = &v
}

func (o ExceptionMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExceptionMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayMessage) {
		toSerialize["displayMessage"] = o.DisplayMessage
	}
	if !IsNil(o.RequestUuid) {
		toSerialize["requestUuid"] = o.RequestUuid
	}
	return toSerialize, nil
}

type NullableExceptionMessage struct {
	value *ExceptionMessage
	isSet bool
}

func (v NullableExceptionMessage) Get() *ExceptionMessage {
	return v.value
}

func (v *NullableExceptionMessage) Set(val *ExceptionMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionMessage(val *ExceptionMessage) *NullableExceptionMessage {
	return &NullableExceptionMessage{value: val, isSet: true}
}

func (v NullableExceptionMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


