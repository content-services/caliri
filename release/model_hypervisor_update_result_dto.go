/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HypervisorUpdateResultDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorUpdateResultDTO{}

// HypervisorUpdateResultDTO struct for HypervisorUpdateResultDTO
type HypervisorUpdateResultDTO struct {
	Created []HypervisorConsumerDTO `json:"created"`
	Updated []HypervisorConsumerDTO `json:"updated"`
	Unchanged []HypervisorConsumerDTO `json:"unchanged"`
	FailedUpdate []string `json:"failedUpdate"`
}

type _HypervisorUpdateResultDTO HypervisorUpdateResultDTO

// NewHypervisorUpdateResultDTO instantiates a new HypervisorUpdateResultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorUpdateResultDTO(created []HypervisorConsumerDTO, updated []HypervisorConsumerDTO, unchanged []HypervisorConsumerDTO, failedUpdate []string) *HypervisorUpdateResultDTO {
	this := HypervisorUpdateResultDTO{}
	this.Created = created
	this.Updated = updated
	this.Unchanged = unchanged
	this.FailedUpdate = failedUpdate
	return &this
}

// NewHypervisorUpdateResultDTOWithDefaults instantiates a new HypervisorUpdateResultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorUpdateResultDTOWithDefaults() *HypervisorUpdateResultDTO {
	this := HypervisorUpdateResultDTO{}
	return &this
}

// GetCreated returns the Created field value
func (o *HypervisorUpdateResultDTO) GetCreated() []HypervisorConsumerDTO {
	if o == nil {
		var ret []HypervisorConsumerDTO
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *HypervisorUpdateResultDTO) GetCreatedOk() ([]HypervisorConsumerDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created, true
}

// SetCreated sets field value
func (o *HypervisorUpdateResultDTO) SetCreated(v []HypervisorConsumerDTO) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *HypervisorUpdateResultDTO) GetUpdated() []HypervisorConsumerDTO {
	if o == nil {
		var ret []HypervisorConsumerDTO
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *HypervisorUpdateResultDTO) GetUpdatedOk() ([]HypervisorConsumerDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Updated, true
}

// SetUpdated sets field value
func (o *HypervisorUpdateResultDTO) SetUpdated(v []HypervisorConsumerDTO) {
	o.Updated = v
}

// GetUnchanged returns the Unchanged field value
func (o *HypervisorUpdateResultDTO) GetUnchanged() []HypervisorConsumerDTO {
	if o == nil {
		var ret []HypervisorConsumerDTO
		return ret
	}

	return o.Unchanged
}

// GetUnchangedOk returns a tuple with the Unchanged field value
// and a boolean to check if the value has been set.
func (o *HypervisorUpdateResultDTO) GetUnchangedOk() ([]HypervisorConsumerDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unchanged, true
}

// SetUnchanged sets field value
func (o *HypervisorUpdateResultDTO) SetUnchanged(v []HypervisorConsumerDTO) {
	o.Unchanged = v
}

// GetFailedUpdate returns the FailedUpdate field value
func (o *HypervisorUpdateResultDTO) GetFailedUpdate() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FailedUpdate
}

// GetFailedUpdateOk returns a tuple with the FailedUpdate field value
// and a boolean to check if the value has been set.
func (o *HypervisorUpdateResultDTO) GetFailedUpdateOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedUpdate, true
}

// SetFailedUpdate sets field value
func (o *HypervisorUpdateResultDTO) SetFailedUpdate(v []string) {
	o.FailedUpdate = v
}

func (o HypervisorUpdateResultDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorUpdateResultDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	toSerialize["unchanged"] = o.Unchanged
	toSerialize["failedUpdate"] = o.FailedUpdate
	return toSerialize, nil
}

func (o *HypervisorUpdateResultDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"updated",
		"unchanged",
		"failedUpdate",
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

	varHypervisorUpdateResultDTO := _HypervisorUpdateResultDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHypervisorUpdateResultDTO)

	if err != nil {
		return err
	}

	*o = HypervisorUpdateResultDTO(varHypervisorUpdateResultDTO)

	return err
}

type NullableHypervisorUpdateResultDTO struct {
	value *HypervisorUpdateResultDTO
	isSet bool
}

func (v NullableHypervisorUpdateResultDTO) Get() *HypervisorUpdateResultDTO {
	return v.value
}

func (v *NullableHypervisorUpdateResultDTO) Set(val *HypervisorUpdateResultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorUpdateResultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorUpdateResultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorUpdateResultDTO(val *HypervisorUpdateResultDTO) *NullableHypervisorUpdateResultDTO {
	return &NullableHypervisorUpdateResultDTO{value: val, isSet: true}
}

func (v NullableHypervisorUpdateResultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorUpdateResultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


