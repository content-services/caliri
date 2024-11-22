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

// checks if the HypervisorIdDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorIdDTO{}

// HypervisorIdDTO Represents a HypervisorId details
type HypervisorIdDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	HypervisorId *string `json:"hypervisorId,omitempty"`
	ReporterId *string `json:"reporterId,omitempty"`
}

// NewHypervisorIdDTO instantiates a new HypervisorIdDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorIdDTO() *HypervisorIdDTO {
	this := HypervisorIdDTO{}
	return &this
}

// NewHypervisorIdDTOWithDefaults instantiates a new HypervisorIdDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorIdDTOWithDefaults() *HypervisorIdDTO {
	this := HypervisorIdDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *HypervisorIdDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorIdDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *HypervisorIdDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *HypervisorIdDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *HypervisorIdDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorIdDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *HypervisorIdDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *HypervisorIdDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HypervisorIdDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorIdDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HypervisorIdDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HypervisorIdDTO) SetId(v string) {
	o.Id = &v
}

// GetHypervisorId returns the HypervisorId field value if set, zero value otherwise.
func (o *HypervisorIdDTO) GetHypervisorId() string {
	if o == nil || IsNil(o.HypervisorId) {
		var ret string
		return ret
	}
	return *o.HypervisorId
}

// GetHypervisorIdOk returns a tuple with the HypervisorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorIdDTO) GetHypervisorIdOk() (*string, bool) {
	if o == nil || IsNil(o.HypervisorId) {
		return nil, false
	}
	return o.HypervisorId, true
}

// HasHypervisorId returns a boolean if a field has been set.
func (o *HypervisorIdDTO) HasHypervisorId() bool {
	if o != nil && !IsNil(o.HypervisorId) {
		return true
	}

	return false
}

// SetHypervisorId gets a reference to the given string and assigns it to the HypervisorId field.
func (o *HypervisorIdDTO) SetHypervisorId(v string) {
	o.HypervisorId = &v
}

// GetReporterId returns the ReporterId field value if set, zero value otherwise.
func (o *HypervisorIdDTO) GetReporterId() string {
	if o == nil || IsNil(o.ReporterId) {
		var ret string
		return ret
	}
	return *o.ReporterId
}

// GetReporterIdOk returns a tuple with the ReporterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorIdDTO) GetReporterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReporterId) {
		return nil, false
	}
	return o.ReporterId, true
}

// HasReporterId returns a boolean if a field has been set.
func (o *HypervisorIdDTO) HasReporterId() bool {
	if o != nil && !IsNil(o.ReporterId) {
		return true
	}

	return false
}

// SetReporterId gets a reference to the given string and assigns it to the ReporterId field.
func (o *HypervisorIdDTO) SetReporterId(v string) {
	o.ReporterId = &v
}

func (o HypervisorIdDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorIdDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.HypervisorId) {
		toSerialize["hypervisorId"] = o.HypervisorId
	}
	if !IsNil(o.ReporterId) {
		toSerialize["reporterId"] = o.ReporterId
	}
	return toSerialize, nil
}

type NullableHypervisorIdDTO struct {
	value *HypervisorIdDTO
	isSet bool
}

func (v NullableHypervisorIdDTO) Get() *HypervisorIdDTO {
	return v.value
}

func (v *NullableHypervisorIdDTO) Set(val *HypervisorIdDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorIdDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorIdDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorIdDTO(val *HypervisorIdDTO) *NullableHypervisorIdDTO {
	return &NullableHypervisorIdDTO{value: val, isSet: true}
}

func (v NullableHypervisorIdDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorIdDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


