/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
	"time"
)

// checks if the AbstractCertificateDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AbstractCertificateDTO{}

// AbstractCertificateDTO Represents the base of most Candlepin certificates presented to the API (exceptions include ProductCertificate which has its own DTO). 
type AbstractCertificateDTO struct {
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
	Cert *string `json:"cert,omitempty"`
	Serial *CertificateSerialDTO `json:"serial,omitempty"`
}

// NewAbstractCertificateDTO instantiates a new AbstractCertificateDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbstractCertificateDTO() *AbstractCertificateDTO {
	this := AbstractCertificateDTO{}
	return &this
}

// NewAbstractCertificateDTOWithDefaults instantiates a new AbstractCertificateDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbstractCertificateDTOWithDefaults() *AbstractCertificateDTO {
	this := AbstractCertificateDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AbstractCertificateDTO) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractCertificateDTO) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AbstractCertificateDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AbstractCertificateDTO) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AbstractCertificateDTO) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractCertificateDTO) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AbstractCertificateDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *AbstractCertificateDTO) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AbstractCertificateDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractCertificateDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AbstractCertificateDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AbstractCertificateDTO) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AbstractCertificateDTO) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractCertificateDTO) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AbstractCertificateDTO) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AbstractCertificateDTO) SetKey(v string) {
	o.Key = &v
}

// GetCert returns the Cert field value if set, zero value otherwise.
func (o *AbstractCertificateDTO) GetCert() string {
	if o == nil || IsNil(o.Cert) {
		var ret string
		return ret
	}
	return *o.Cert
}

// GetCertOk returns a tuple with the Cert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractCertificateDTO) GetCertOk() (*string, bool) {
	if o == nil || IsNil(o.Cert) {
		return nil, false
	}
	return o.Cert, true
}

// HasCert returns a boolean if a field has been set.
func (o *AbstractCertificateDTO) HasCert() bool {
	if o != nil && !IsNil(o.Cert) {
		return true
	}

	return false
}

// SetCert gets a reference to the given string and assigns it to the Cert field.
func (o *AbstractCertificateDTO) SetCert(v string) {
	o.Cert = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *AbstractCertificateDTO) GetSerial() CertificateSerialDTO {
	if o == nil || IsNil(o.Serial) {
		var ret CertificateSerialDTO
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractCertificateDTO) GetSerialOk() (*CertificateSerialDTO, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *AbstractCertificateDTO) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given CertificateSerialDTO and assigns it to the Serial field.
func (o *AbstractCertificateDTO) SetSerial(v CertificateSerialDTO) {
	o.Serial = &v
}

func (o AbstractCertificateDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AbstractCertificateDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Cert) {
		toSerialize["cert"] = o.Cert
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return toSerialize, nil
}

type NullableAbstractCertificateDTO struct {
	value *AbstractCertificateDTO
	isSet bool
}

func (v NullableAbstractCertificateDTO) Get() *AbstractCertificateDTO {
	return v.value
}

func (v *NullableAbstractCertificateDTO) Set(val *AbstractCertificateDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAbstractCertificateDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAbstractCertificateDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbstractCertificateDTO(val *AbstractCertificateDTO) *NullableAbstractCertificateDTO {
	return &NullableAbstractCertificateDTO{value: val, isSet: true}
}

func (v NullableAbstractCertificateDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbstractCertificateDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

