/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OwnerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnerInfo{}

// OwnerInfo Represents owner information
type OwnerInfo struct {
	ConsumerCounts map[string]int32 `json:"consumerCounts"`
	ConsumerGuestCounts map[string]int32 `json:"consumerGuestCounts"`
	EntitlementsConsumedByType map[string]int32 `json:"entitlementsConsumedByType"`
	ConsumerTypeCountByPool map[string]int32 `json:"consumerTypeCountByPool"`
	EnabledConsumerTypeCountByPool map[string]int32 `json:"enabledConsumerTypeCountByPool"`
	ConsumerCountsByComplianceStatus map[string]int32 `json:"consumerCountsByComplianceStatus"`
	EntitlementsConsumedByFamily map[string]ConsumptionTypeCountsDTO `json:"entitlementsConsumedByFamily"`
}

type _OwnerInfo OwnerInfo

// NewOwnerInfo instantiates a new OwnerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerInfo(consumerCounts map[string]int32, consumerGuestCounts map[string]int32, entitlementsConsumedByType map[string]int32, consumerTypeCountByPool map[string]int32, enabledConsumerTypeCountByPool map[string]int32, consumerCountsByComplianceStatus map[string]int32, entitlementsConsumedByFamily map[string]ConsumptionTypeCountsDTO) *OwnerInfo {
	this := OwnerInfo{}
	this.ConsumerCounts = consumerCounts
	this.ConsumerGuestCounts = consumerGuestCounts
	this.EntitlementsConsumedByType = entitlementsConsumedByType
	this.ConsumerTypeCountByPool = consumerTypeCountByPool
	this.EnabledConsumerTypeCountByPool = enabledConsumerTypeCountByPool
	this.ConsumerCountsByComplianceStatus = consumerCountsByComplianceStatus
	this.EntitlementsConsumedByFamily = entitlementsConsumedByFamily
	return &this
}

// NewOwnerInfoWithDefaults instantiates a new OwnerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerInfoWithDefaults() *OwnerInfo {
	this := OwnerInfo{}
	return &this
}

// GetConsumerCounts returns the ConsumerCounts field value
func (o *OwnerInfo) GetConsumerCounts() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.ConsumerCounts
}

// GetConsumerCountsOk returns a tuple with the ConsumerCounts field value
// and a boolean to check if the value has been set.
func (o *OwnerInfo) GetConsumerCountsOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerCounts, true
}

// SetConsumerCounts sets field value
func (o *OwnerInfo) SetConsumerCounts(v map[string]int32) {
	o.ConsumerCounts = v
}

// GetConsumerGuestCounts returns the ConsumerGuestCounts field value
func (o *OwnerInfo) GetConsumerGuestCounts() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.ConsumerGuestCounts
}

// GetConsumerGuestCountsOk returns a tuple with the ConsumerGuestCounts field value
// and a boolean to check if the value has been set.
func (o *OwnerInfo) GetConsumerGuestCountsOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerGuestCounts, true
}

// SetConsumerGuestCounts sets field value
func (o *OwnerInfo) SetConsumerGuestCounts(v map[string]int32) {
	o.ConsumerGuestCounts = v
}

// GetEntitlementsConsumedByType returns the EntitlementsConsumedByType field value
func (o *OwnerInfo) GetEntitlementsConsumedByType() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.EntitlementsConsumedByType
}

// GetEntitlementsConsumedByTypeOk returns a tuple with the EntitlementsConsumedByType field value
// and a boolean to check if the value has been set.
func (o *OwnerInfo) GetEntitlementsConsumedByTypeOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementsConsumedByType, true
}

// SetEntitlementsConsumedByType sets field value
func (o *OwnerInfo) SetEntitlementsConsumedByType(v map[string]int32) {
	o.EntitlementsConsumedByType = v
}

// GetConsumerTypeCountByPool returns the ConsumerTypeCountByPool field value
func (o *OwnerInfo) GetConsumerTypeCountByPool() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.ConsumerTypeCountByPool
}

// GetConsumerTypeCountByPoolOk returns a tuple with the ConsumerTypeCountByPool field value
// and a boolean to check if the value has been set.
func (o *OwnerInfo) GetConsumerTypeCountByPoolOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerTypeCountByPool, true
}

// SetConsumerTypeCountByPool sets field value
func (o *OwnerInfo) SetConsumerTypeCountByPool(v map[string]int32) {
	o.ConsumerTypeCountByPool = v
}

// GetEnabledConsumerTypeCountByPool returns the EnabledConsumerTypeCountByPool field value
func (o *OwnerInfo) GetEnabledConsumerTypeCountByPool() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.EnabledConsumerTypeCountByPool
}

// GetEnabledConsumerTypeCountByPoolOk returns a tuple with the EnabledConsumerTypeCountByPool field value
// and a boolean to check if the value has been set.
func (o *OwnerInfo) GetEnabledConsumerTypeCountByPoolOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnabledConsumerTypeCountByPool, true
}

// SetEnabledConsumerTypeCountByPool sets field value
func (o *OwnerInfo) SetEnabledConsumerTypeCountByPool(v map[string]int32) {
	o.EnabledConsumerTypeCountByPool = v
}

// GetConsumerCountsByComplianceStatus returns the ConsumerCountsByComplianceStatus field value
func (o *OwnerInfo) GetConsumerCountsByComplianceStatus() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.ConsumerCountsByComplianceStatus
}

// GetConsumerCountsByComplianceStatusOk returns a tuple with the ConsumerCountsByComplianceStatus field value
// and a boolean to check if the value has been set.
func (o *OwnerInfo) GetConsumerCountsByComplianceStatusOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerCountsByComplianceStatus, true
}

// SetConsumerCountsByComplianceStatus sets field value
func (o *OwnerInfo) SetConsumerCountsByComplianceStatus(v map[string]int32) {
	o.ConsumerCountsByComplianceStatus = v
}

// GetEntitlementsConsumedByFamily returns the EntitlementsConsumedByFamily field value
func (o *OwnerInfo) GetEntitlementsConsumedByFamily() map[string]ConsumptionTypeCountsDTO {
	if o == nil {
		var ret map[string]ConsumptionTypeCountsDTO
		return ret
	}

	return o.EntitlementsConsumedByFamily
}

// GetEntitlementsConsumedByFamilyOk returns a tuple with the EntitlementsConsumedByFamily field value
// and a boolean to check if the value has been set.
func (o *OwnerInfo) GetEntitlementsConsumedByFamilyOk() (*map[string]ConsumptionTypeCountsDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementsConsumedByFamily, true
}

// SetEntitlementsConsumedByFamily sets field value
func (o *OwnerInfo) SetEntitlementsConsumedByFamily(v map[string]ConsumptionTypeCountsDTO) {
	o.EntitlementsConsumedByFamily = v
}

func (o OwnerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["consumerCounts"] = o.ConsumerCounts
	toSerialize["consumerGuestCounts"] = o.ConsumerGuestCounts
	toSerialize["entitlementsConsumedByType"] = o.EntitlementsConsumedByType
	toSerialize["consumerTypeCountByPool"] = o.ConsumerTypeCountByPool
	toSerialize["enabledConsumerTypeCountByPool"] = o.EnabledConsumerTypeCountByPool
	toSerialize["consumerCountsByComplianceStatus"] = o.ConsumerCountsByComplianceStatus
	toSerialize["entitlementsConsumedByFamily"] = o.EntitlementsConsumedByFamily
	return toSerialize, nil
}

func (o *OwnerInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"consumerCounts",
		"consumerGuestCounts",
		"entitlementsConsumedByType",
		"consumerTypeCountByPool",
		"enabledConsumerTypeCountByPool",
		"consumerCountsByComplianceStatus",
		"entitlementsConsumedByFamily",
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

	varOwnerInfo := _OwnerInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOwnerInfo)

	if err != nil {
		return err
	}

	*o = OwnerInfo(varOwnerInfo)

	return err
}

type NullableOwnerInfo struct {
	value *OwnerInfo
	isSet bool
}

func (v NullableOwnerInfo) Get() *OwnerInfo {
	return v.value
}

func (v *NullableOwnerInfo) Set(val *OwnerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerInfo(val *OwnerInfo) *NullableOwnerInfo {
	return &NullableOwnerInfo{value: val, isSet: true}
}

func (v NullableOwnerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


