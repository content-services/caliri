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

// checks if the ExportResultDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportResultDTO{}

// ExportResultDTO Represents the result of an export job
type ExportResultDTO struct {
	ExportedConsumer *string `json:"exportedConsumer,omitempty"`
	ExportId *string `json:"exportId,omitempty"`
	Href *string `json:"href,omitempty"`
}

// NewExportResultDTO instantiates a new ExportResultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportResultDTO() *ExportResultDTO {
	this := ExportResultDTO{}
	return &this
}

// NewExportResultDTOWithDefaults instantiates a new ExportResultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportResultDTOWithDefaults() *ExportResultDTO {
	this := ExportResultDTO{}
	return &this
}

// GetExportedConsumer returns the ExportedConsumer field value if set, zero value otherwise.
func (o *ExportResultDTO) GetExportedConsumer() string {
	if o == nil || IsNil(o.ExportedConsumer) {
		var ret string
		return ret
	}
	return *o.ExportedConsumer
}

// GetExportedConsumerOk returns a tuple with the ExportedConsumer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportResultDTO) GetExportedConsumerOk() (*string, bool) {
	if o == nil || IsNil(o.ExportedConsumer) {
		return nil, false
	}
	return o.ExportedConsumer, true
}

// HasExportedConsumer returns a boolean if a field has been set.
func (o *ExportResultDTO) HasExportedConsumer() bool {
	if o != nil && !IsNil(o.ExportedConsumer) {
		return true
	}

	return false
}

// SetExportedConsumer gets a reference to the given string and assigns it to the ExportedConsumer field.
func (o *ExportResultDTO) SetExportedConsumer(v string) {
	o.ExportedConsumer = &v
}

// GetExportId returns the ExportId field value if set, zero value otherwise.
func (o *ExportResultDTO) GetExportId() string {
	if o == nil || IsNil(o.ExportId) {
		var ret string
		return ret
	}
	return *o.ExportId
}

// GetExportIdOk returns a tuple with the ExportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportResultDTO) GetExportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExportId) {
		return nil, false
	}
	return o.ExportId, true
}

// HasExportId returns a boolean if a field has been set.
func (o *ExportResultDTO) HasExportId() bool {
	if o != nil && !IsNil(o.ExportId) {
		return true
	}

	return false
}

// SetExportId gets a reference to the given string and assigns it to the ExportId field.
func (o *ExportResultDTO) SetExportId(v string) {
	o.ExportId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ExportResultDTO) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportResultDTO) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ExportResultDTO) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ExportResultDTO) SetHref(v string) {
	o.Href = &v
}

func (o ExportResultDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportResultDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExportedConsumer) {
		toSerialize["exportedConsumer"] = o.ExportedConsumer
	}
	if !IsNil(o.ExportId) {
		toSerialize["exportId"] = o.ExportId
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableExportResultDTO struct {
	value *ExportResultDTO
	isSet bool
}

func (v NullableExportResultDTO) Get() *ExportResultDTO {
	return v.value
}

func (v *NullableExportResultDTO) Set(val *ExportResultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableExportResultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableExportResultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportResultDTO(val *ExportResultDTO) *NullableExportResultDTO {
	return &NullableExportResultDTO{value: val, isSet: true}
}

func (v NullableExportResultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportResultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


