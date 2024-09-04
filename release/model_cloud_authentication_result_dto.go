/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CloudAuthenticationResultDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudAuthenticationResultDTO{}

// CloudAuthenticationResultDTO Contains the result of authenticating the provided cloud registration information.
type CloudAuthenticationResultDTO struct {
	OwnerKey *string `json:"ownerKey,omitempty"`
	AnonymousConsumerUuid *string `json:"anonymousConsumerUuid,omitempty"`
	Token string `json:"token"`
	TokenType string `json:"tokenType"`
}

type _CloudAuthenticationResultDTO CloudAuthenticationResultDTO

// NewCloudAuthenticationResultDTO instantiates a new CloudAuthenticationResultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAuthenticationResultDTO(token string, tokenType string) *CloudAuthenticationResultDTO {
	this := CloudAuthenticationResultDTO{}
	this.Token = token
	this.TokenType = tokenType
	return &this
}

// NewCloudAuthenticationResultDTOWithDefaults instantiates a new CloudAuthenticationResultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAuthenticationResultDTOWithDefaults() *CloudAuthenticationResultDTO {
	this := CloudAuthenticationResultDTO{}
	return &this
}

// GetOwnerKey returns the OwnerKey field value if set, zero value otherwise.
func (o *CloudAuthenticationResultDTO) GetOwnerKey() string {
	if o == nil || IsNil(o.OwnerKey) {
		var ret string
		return ret
	}
	return *o.OwnerKey
}

// GetOwnerKeyOk returns a tuple with the OwnerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAuthenticationResultDTO) GetOwnerKeyOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerKey) {
		return nil, false
	}
	return o.OwnerKey, true
}

// HasOwnerKey returns a boolean if a field has been set.
func (o *CloudAuthenticationResultDTO) HasOwnerKey() bool {
	if o != nil && !IsNil(o.OwnerKey) {
		return true
	}

	return false
}

// SetOwnerKey gets a reference to the given string and assigns it to the OwnerKey field.
func (o *CloudAuthenticationResultDTO) SetOwnerKey(v string) {
	o.OwnerKey = &v
}

// GetAnonymousConsumerUuid returns the AnonymousConsumerUuid field value if set, zero value otherwise.
func (o *CloudAuthenticationResultDTO) GetAnonymousConsumerUuid() string {
	if o == nil || IsNil(o.AnonymousConsumerUuid) {
		var ret string
		return ret
	}
	return *o.AnonymousConsumerUuid
}

// GetAnonymousConsumerUuidOk returns a tuple with the AnonymousConsumerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAuthenticationResultDTO) GetAnonymousConsumerUuidOk() (*string, bool) {
	if o == nil || IsNil(o.AnonymousConsumerUuid) {
		return nil, false
	}
	return o.AnonymousConsumerUuid, true
}

// HasAnonymousConsumerUuid returns a boolean if a field has been set.
func (o *CloudAuthenticationResultDTO) HasAnonymousConsumerUuid() bool {
	if o != nil && !IsNil(o.AnonymousConsumerUuid) {
		return true
	}

	return false
}

// SetAnonymousConsumerUuid gets a reference to the given string and assigns it to the AnonymousConsumerUuid field.
func (o *CloudAuthenticationResultDTO) SetAnonymousConsumerUuid(v string) {
	o.AnonymousConsumerUuid = &v
}

// GetToken returns the Token field value
func (o *CloudAuthenticationResultDTO) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CloudAuthenticationResultDTO) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CloudAuthenticationResultDTO) SetToken(v string) {
	o.Token = v
}

// GetTokenType returns the TokenType field value
func (o *CloudAuthenticationResultDTO) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *CloudAuthenticationResultDTO) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *CloudAuthenticationResultDTO) SetTokenType(v string) {
	o.TokenType = v
}

func (o CloudAuthenticationResultDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudAuthenticationResultDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OwnerKey) {
		toSerialize["ownerKey"] = o.OwnerKey
	}
	if !IsNil(o.AnonymousConsumerUuid) {
		toSerialize["anonymousConsumerUuid"] = o.AnonymousConsumerUuid
	}
	toSerialize["token"] = o.Token
	toSerialize["tokenType"] = o.TokenType
	return toSerialize, nil
}

func (o *CloudAuthenticationResultDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
		"tokenType",
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

	varCloudAuthenticationResultDTO := _CloudAuthenticationResultDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudAuthenticationResultDTO)

	if err != nil {
		return err
	}

	*o = CloudAuthenticationResultDTO(varCloudAuthenticationResultDTO)

	return err
}

type NullableCloudAuthenticationResultDTO struct {
	value *CloudAuthenticationResultDTO
	isSet bool
}

func (v NullableCloudAuthenticationResultDTO) Get() *CloudAuthenticationResultDTO {
	return v.value
}

func (v *NullableCloudAuthenticationResultDTO) Set(val *CloudAuthenticationResultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAuthenticationResultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAuthenticationResultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAuthenticationResultDTO(val *CloudAuthenticationResultDTO) *NullableCloudAuthenticationResultDTO {
	return &NullableCloudAuthenticationResultDTO{value: val, isSet: true}
}

func (v NullableCloudAuthenticationResultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAuthenticationResultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


