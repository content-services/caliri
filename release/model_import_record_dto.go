/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the ImportRecordDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportRecordDTO{}

// ImportRecordDTO Represents a import record details
type ImportRecordDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusMessage *string `json:"statusMessage,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	GeneratedBy *string `json:"generatedBy,omitempty"`
	GeneratedDate *string `json:"generatedDate,omitempty"`
	UpstreamConsumer *ImportUpstreamConsumerDTO `json:"upstreamConsumer,omitempty"`
}

// NewImportRecordDTO instantiates a new ImportRecordDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRecordDTO() *ImportRecordDTO {
	this := ImportRecordDTO{}
	return &this
}

// NewImportRecordDTOWithDefaults instantiates a new ImportRecordDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRecordDTOWithDefaults() *ImportRecordDTO {
	this := ImportRecordDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ImportRecordDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRecordDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ImportRecordDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ImportRecordDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ImportRecordDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRecordDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ImportRecordDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *ImportRecordDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportRecordDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRecordDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportRecordDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImportRecordDTO) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ImportRecordDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRecordDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ImportRecordDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ImportRecordDTO) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ImportRecordDTO) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRecordDTO) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ImportRecordDTO) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ImportRecordDTO) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ImportRecordDTO) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRecordDTO) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ImportRecordDTO) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ImportRecordDTO) SetFileName(v string) {
	o.FileName = &v
}

// GetGeneratedBy returns the GeneratedBy field value if set, zero value otherwise.
func (o *ImportRecordDTO) GetGeneratedBy() string {
	if o == nil || IsNil(o.GeneratedBy) {
		var ret string
		return ret
	}
	return *o.GeneratedBy
}

// GetGeneratedByOk returns a tuple with the GeneratedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRecordDTO) GetGeneratedByOk() (*string, bool) {
	if o == nil || IsNil(o.GeneratedBy) {
		return nil, false
	}
	return o.GeneratedBy, true
}

// HasGeneratedBy returns a boolean if a field has been set.
func (o *ImportRecordDTO) HasGeneratedBy() bool {
	if o != nil && !IsNil(o.GeneratedBy) {
		return true
	}

	return false
}

// SetGeneratedBy gets a reference to the given string and assigns it to the GeneratedBy field.
func (o *ImportRecordDTO) SetGeneratedBy(v string) {
	o.GeneratedBy = &v
}

// GetGeneratedDate returns the GeneratedDate field value if set, zero value otherwise.
func (o *ImportRecordDTO) GetGeneratedDate() string {
	if o == nil || IsNil(o.GeneratedDate) {
		var ret string
		return ret
	}
	return *o.GeneratedDate
}

// GetGeneratedDateOk returns a tuple with the GeneratedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRecordDTO) GetGeneratedDateOk() (*string, bool) {
	if o == nil || IsNil(o.GeneratedDate) {
		return nil, false
	}
	return o.GeneratedDate, true
}

// HasGeneratedDate returns a boolean if a field has been set.
func (o *ImportRecordDTO) HasGeneratedDate() bool {
	if o != nil && !IsNil(o.GeneratedDate) {
		return true
	}

	return false
}

// SetGeneratedDate gets a reference to the given string and assigns it to the GeneratedDate field.
func (o *ImportRecordDTO) SetGeneratedDate(v string) {
	o.GeneratedDate = &v
}

// GetUpstreamConsumer returns the UpstreamConsumer field value if set, zero value otherwise.
func (o *ImportRecordDTO) GetUpstreamConsumer() ImportUpstreamConsumerDTO {
	if o == nil || IsNil(o.UpstreamConsumer) {
		var ret ImportUpstreamConsumerDTO
		return ret
	}
	return *o.UpstreamConsumer
}

// GetUpstreamConsumerOk returns a tuple with the UpstreamConsumer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRecordDTO) GetUpstreamConsumerOk() (*ImportUpstreamConsumerDTO, bool) {
	if o == nil || IsNil(o.UpstreamConsumer) {
		return nil, false
	}
	return o.UpstreamConsumer, true
}

// HasUpstreamConsumer returns a boolean if a field has been set.
func (o *ImportRecordDTO) HasUpstreamConsumer() bool {
	if o != nil && !IsNil(o.UpstreamConsumer) {
		return true
	}

	return false
}

// SetUpstreamConsumer gets a reference to the given ImportUpstreamConsumerDTO and assigns it to the UpstreamConsumer field.
func (o *ImportRecordDTO) SetUpstreamConsumer(v ImportUpstreamConsumerDTO) {
	o.UpstreamConsumer = &v
}

func (o ImportRecordDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportRecordDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.GeneratedBy) {
		toSerialize["generatedBy"] = o.GeneratedBy
	}
	if !IsNil(o.GeneratedDate) {
		toSerialize["generatedDate"] = o.GeneratedDate
	}
	if !IsNil(o.UpstreamConsumer) {
		toSerialize["upstreamConsumer"] = o.UpstreamConsumer
	}
	return toSerialize, nil
}

type NullableImportRecordDTO struct {
	value *ImportRecordDTO
	isSet bool
}

func (v NullableImportRecordDTO) Get() *ImportRecordDTO {
	return v.value
}

func (v *NullableImportRecordDTO) Set(val *ImportRecordDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRecordDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRecordDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRecordDTO(val *ImportRecordDTO) *NullableImportRecordDTO {
	return &NullableImportRecordDTO{value: val, isSet: true}
}

func (v NullableImportRecordDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRecordDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


