/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the CertificateSerialDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateSerialDTO{}

// CertificateSerialDTO Represents a database sequence used to ensure certificates receive unique serial numbers
type CertificateSerialDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Serial *int64 `json:"serial,omitempty"`
	Expiration *string `json:"expiration,omitempty"`
	Revoked *bool `json:"revoked,omitempty"`
}

// NewCertificateSerialDTO instantiates a new CertificateSerialDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateSerialDTO() *CertificateSerialDTO {
	this := CertificateSerialDTO{}
	return &this
}

// NewCertificateSerialDTOWithDefaults instantiates a new CertificateSerialDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateSerialDTOWithDefaults() *CertificateSerialDTO {
	this := CertificateSerialDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CertificateSerialDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSerialDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CertificateSerialDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *CertificateSerialDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *CertificateSerialDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSerialDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *CertificateSerialDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *CertificateSerialDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CertificateSerialDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSerialDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CertificateSerialDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CertificateSerialDTO) SetId(v int64) {
	o.Id = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *CertificateSerialDTO) GetSerial() int64 {
	if o == nil || IsNil(o.Serial) {
		var ret int64
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSerialDTO) GetSerialOk() (*int64, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *CertificateSerialDTO) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given int64 and assigns it to the Serial field.
func (o *CertificateSerialDTO) SetSerial(v int64) {
	o.Serial = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *CertificateSerialDTO) GetExpiration() string {
	if o == nil || IsNil(o.Expiration) {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSerialDTO) GetExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *CertificateSerialDTO) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *CertificateSerialDTO) SetExpiration(v string) {
	o.Expiration = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *CertificateSerialDTO) GetRevoked() bool {
	if o == nil || IsNil(o.Revoked) {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateSerialDTO) GetRevokedOk() (*bool, bool) {
	if o == nil || IsNil(o.Revoked) {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *CertificateSerialDTO) HasRevoked() bool {
	if o != nil && !IsNil(o.Revoked) {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *CertificateSerialDTO) SetRevoked(v bool) {
	o.Revoked = &v
}

func (o CertificateSerialDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateSerialDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.Revoked) {
		toSerialize["revoked"] = o.Revoked
	}
	return toSerialize, nil
}

type NullableCertificateSerialDTO struct {
	value *CertificateSerialDTO
	isSet bool
}

func (v NullableCertificateSerialDTO) Get() *CertificateSerialDTO {
	return v.value
}

func (v *NullableCertificateSerialDTO) Set(val *CertificateSerialDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateSerialDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateSerialDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateSerialDTO(val *CertificateSerialDTO) *NullableCertificateSerialDTO {
	return &NullableCertificateSerialDTO{value: val, isSet: true}
}

func (v NullableCertificateSerialDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateSerialDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


