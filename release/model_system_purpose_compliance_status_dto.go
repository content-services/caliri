/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the SystemPurposeComplianceStatusDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemPurposeComplianceStatusDTO{}

// SystemPurposeComplianceStatusDTO DTO representing system purpose compliance status
type SystemPurposeComplianceStatusDTO struct {
	Status *string `json:"status,omitempty"`
	Compliant *bool `json:"compliant,omitempty"`
	Date *string `json:"date,omitempty"`
	NonCompliantRole *string `json:"nonCompliantRole,omitempty"`
	NonCompliantSLA *string `json:"nonCompliantSLA,omitempty"`
	NonCompliantUsage *string `json:"nonCompliantUsage,omitempty"`
	NonCompliantServiceType *string `json:"nonCompliantServiceType,omitempty"`
	CompliantRole *map[string][]EntitlementDTO `json:"compliantRole,omitempty"`
	CompliantAddOns *map[string][]EntitlementDTO `json:"compliantAddOns,omitempty"`
	CompliantSLA *map[string][]EntitlementDTO `json:"compliantSLA,omitempty"`
	CompliantUsage *map[string][]EntitlementDTO `json:"compliantUsage,omitempty"`
	NonCompliantAddOns []string `json:"nonCompliantAddOns,omitempty"`
	CompliantServiceType *map[string][]EntitlementDTO `json:"compliantServiceType,omitempty"`
	Reasons []string `json:"reasons,omitempty"`
}

// NewSystemPurposeComplianceStatusDTO instantiates a new SystemPurposeComplianceStatusDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemPurposeComplianceStatusDTO() *SystemPurposeComplianceStatusDTO {
	this := SystemPurposeComplianceStatusDTO{}
	return &this
}

// NewSystemPurposeComplianceStatusDTOWithDefaults instantiates a new SystemPurposeComplianceStatusDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemPurposeComplianceStatusDTOWithDefaults() *SystemPurposeComplianceStatusDTO {
	this := SystemPurposeComplianceStatusDTO{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SystemPurposeComplianceStatusDTO) SetStatus(v string) {
	o.Status = &v
}

// GetCompliant returns the Compliant field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetCompliant() bool {
	if o == nil || IsNil(o.Compliant) {
		var ret bool
		return ret
	}
	return *o.Compliant
}

// GetCompliantOk returns a tuple with the Compliant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetCompliantOk() (*bool, bool) {
	if o == nil || IsNil(o.Compliant) {
		return nil, false
	}
	return o.Compliant, true
}

// HasCompliant returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasCompliant() bool {
	if o != nil && !IsNil(o.Compliant) {
		return true
	}

	return false
}

// SetCompliant gets a reference to the given bool and assigns it to the Compliant field.
func (o *SystemPurposeComplianceStatusDTO) SetCompliant(v bool) {
	o.Compliant = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *SystemPurposeComplianceStatusDTO) SetDate(v string) {
	o.Date = &v
}

// GetNonCompliantRole returns the NonCompliantRole field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantRole() string {
	if o == nil || IsNil(o.NonCompliantRole) {
		var ret string
		return ret
	}
	return *o.NonCompliantRole
}

// GetNonCompliantRoleOk returns a tuple with the NonCompliantRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantRoleOk() (*string, bool) {
	if o == nil || IsNil(o.NonCompliantRole) {
		return nil, false
	}
	return o.NonCompliantRole, true
}

// HasNonCompliantRole returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasNonCompliantRole() bool {
	if o != nil && !IsNil(o.NonCompliantRole) {
		return true
	}

	return false
}

// SetNonCompliantRole gets a reference to the given string and assigns it to the NonCompliantRole field.
func (o *SystemPurposeComplianceStatusDTO) SetNonCompliantRole(v string) {
	o.NonCompliantRole = &v
}

// GetNonCompliantSLA returns the NonCompliantSLA field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantSLA() string {
	if o == nil || IsNil(o.NonCompliantSLA) {
		var ret string
		return ret
	}
	return *o.NonCompliantSLA
}

// GetNonCompliantSLAOk returns a tuple with the NonCompliantSLA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantSLAOk() (*string, bool) {
	if o == nil || IsNil(o.NonCompliantSLA) {
		return nil, false
	}
	return o.NonCompliantSLA, true
}

// HasNonCompliantSLA returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasNonCompliantSLA() bool {
	if o != nil && !IsNil(o.NonCompliantSLA) {
		return true
	}

	return false
}

// SetNonCompliantSLA gets a reference to the given string and assigns it to the NonCompliantSLA field.
func (o *SystemPurposeComplianceStatusDTO) SetNonCompliantSLA(v string) {
	o.NonCompliantSLA = &v
}

// GetNonCompliantUsage returns the NonCompliantUsage field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantUsage() string {
	if o == nil || IsNil(o.NonCompliantUsage) {
		var ret string
		return ret
	}
	return *o.NonCompliantUsage
}

// GetNonCompliantUsageOk returns a tuple with the NonCompliantUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantUsageOk() (*string, bool) {
	if o == nil || IsNil(o.NonCompliantUsage) {
		return nil, false
	}
	return o.NonCompliantUsage, true
}

// HasNonCompliantUsage returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasNonCompliantUsage() bool {
	if o != nil && !IsNil(o.NonCompliantUsage) {
		return true
	}

	return false
}

// SetNonCompliantUsage gets a reference to the given string and assigns it to the NonCompliantUsage field.
func (o *SystemPurposeComplianceStatusDTO) SetNonCompliantUsage(v string) {
	o.NonCompliantUsage = &v
}

// GetNonCompliantServiceType returns the NonCompliantServiceType field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantServiceType() string {
	if o == nil || IsNil(o.NonCompliantServiceType) {
		var ret string
		return ret
	}
	return *o.NonCompliantServiceType
}

// GetNonCompliantServiceTypeOk returns a tuple with the NonCompliantServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NonCompliantServiceType) {
		return nil, false
	}
	return o.NonCompliantServiceType, true
}

// HasNonCompliantServiceType returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasNonCompliantServiceType() bool {
	if o != nil && !IsNil(o.NonCompliantServiceType) {
		return true
	}

	return false
}

// SetNonCompliantServiceType gets a reference to the given string and assigns it to the NonCompliantServiceType field.
func (o *SystemPurposeComplianceStatusDTO) SetNonCompliantServiceType(v string) {
	o.NonCompliantServiceType = &v
}

// GetCompliantRole returns the CompliantRole field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetCompliantRole() map[string][]EntitlementDTO {
	if o == nil || IsNil(o.CompliantRole) {
		var ret map[string][]EntitlementDTO
		return ret
	}
	return *o.CompliantRole
}

// GetCompliantRoleOk returns a tuple with the CompliantRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetCompliantRoleOk() (*map[string][]EntitlementDTO, bool) {
	if o == nil || IsNil(o.CompliantRole) {
		return nil, false
	}
	return o.CompliantRole, true
}

// HasCompliantRole returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasCompliantRole() bool {
	if o != nil && !IsNil(o.CompliantRole) {
		return true
	}

	return false
}

// SetCompliantRole gets a reference to the given map[string][]EntitlementDTO and assigns it to the CompliantRole field.
func (o *SystemPurposeComplianceStatusDTO) SetCompliantRole(v map[string][]EntitlementDTO) {
	o.CompliantRole = &v
}

// GetCompliantAddOns returns the CompliantAddOns field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetCompliantAddOns() map[string][]EntitlementDTO {
	if o == nil || IsNil(o.CompliantAddOns) {
		var ret map[string][]EntitlementDTO
		return ret
	}
	return *o.CompliantAddOns
}

// GetCompliantAddOnsOk returns a tuple with the CompliantAddOns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetCompliantAddOnsOk() (*map[string][]EntitlementDTO, bool) {
	if o == nil || IsNil(o.CompliantAddOns) {
		return nil, false
	}
	return o.CompliantAddOns, true
}

// HasCompliantAddOns returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasCompliantAddOns() bool {
	if o != nil && !IsNil(o.CompliantAddOns) {
		return true
	}

	return false
}

// SetCompliantAddOns gets a reference to the given map[string][]EntitlementDTO and assigns it to the CompliantAddOns field.
func (o *SystemPurposeComplianceStatusDTO) SetCompliantAddOns(v map[string][]EntitlementDTO) {
	o.CompliantAddOns = &v
}

// GetCompliantSLA returns the CompliantSLA field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetCompliantSLA() map[string][]EntitlementDTO {
	if o == nil || IsNil(o.CompliantSLA) {
		var ret map[string][]EntitlementDTO
		return ret
	}
	return *o.CompliantSLA
}

// GetCompliantSLAOk returns a tuple with the CompliantSLA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetCompliantSLAOk() (*map[string][]EntitlementDTO, bool) {
	if o == nil || IsNil(o.CompliantSLA) {
		return nil, false
	}
	return o.CompliantSLA, true
}

// HasCompliantSLA returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasCompliantSLA() bool {
	if o != nil && !IsNil(o.CompliantSLA) {
		return true
	}

	return false
}

// SetCompliantSLA gets a reference to the given map[string][]EntitlementDTO and assigns it to the CompliantSLA field.
func (o *SystemPurposeComplianceStatusDTO) SetCompliantSLA(v map[string][]EntitlementDTO) {
	o.CompliantSLA = &v
}

// GetCompliantUsage returns the CompliantUsage field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetCompliantUsage() map[string][]EntitlementDTO {
	if o == nil || IsNil(o.CompliantUsage) {
		var ret map[string][]EntitlementDTO
		return ret
	}
	return *o.CompliantUsage
}

// GetCompliantUsageOk returns a tuple with the CompliantUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetCompliantUsageOk() (*map[string][]EntitlementDTO, bool) {
	if o == nil || IsNil(o.CompliantUsage) {
		return nil, false
	}
	return o.CompliantUsage, true
}

// HasCompliantUsage returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasCompliantUsage() bool {
	if o != nil && !IsNil(o.CompliantUsage) {
		return true
	}

	return false
}

// SetCompliantUsage gets a reference to the given map[string][]EntitlementDTO and assigns it to the CompliantUsage field.
func (o *SystemPurposeComplianceStatusDTO) SetCompliantUsage(v map[string][]EntitlementDTO) {
	o.CompliantUsage = &v
}

// GetNonCompliantAddOns returns the NonCompliantAddOns field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantAddOns() []string {
	if o == nil || IsNil(o.NonCompliantAddOns) {
		var ret []string
		return ret
	}
	return o.NonCompliantAddOns
}

// GetNonCompliantAddOnsOk returns a tuple with the NonCompliantAddOns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantAddOnsOk() ([]string, bool) {
	if o == nil || IsNil(o.NonCompliantAddOns) {
		return nil, false
	}
	return o.NonCompliantAddOns, true
}

// HasNonCompliantAddOns returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasNonCompliantAddOns() bool {
	if o != nil && !IsNil(o.NonCompliantAddOns) {
		return true
	}

	return false
}

// SetNonCompliantAddOns gets a reference to the given []string and assigns it to the NonCompliantAddOns field.
func (o *SystemPurposeComplianceStatusDTO) SetNonCompliantAddOns(v []string) {
	o.NonCompliantAddOns = v
}

// GetCompliantServiceType returns the CompliantServiceType field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetCompliantServiceType() map[string][]EntitlementDTO {
	if o == nil || IsNil(o.CompliantServiceType) {
		var ret map[string][]EntitlementDTO
		return ret
	}
	return *o.CompliantServiceType
}

// GetCompliantServiceTypeOk returns a tuple with the CompliantServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetCompliantServiceTypeOk() (*map[string][]EntitlementDTO, bool) {
	if o == nil || IsNil(o.CompliantServiceType) {
		return nil, false
	}
	return o.CompliantServiceType, true
}

// HasCompliantServiceType returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasCompliantServiceType() bool {
	if o != nil && !IsNil(o.CompliantServiceType) {
		return true
	}

	return false
}

// SetCompliantServiceType gets a reference to the given map[string][]EntitlementDTO and assigns it to the CompliantServiceType field.
func (o *SystemPurposeComplianceStatusDTO) SetCompliantServiceType(v map[string][]EntitlementDTO) {
	o.CompliantServiceType = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *SystemPurposeComplianceStatusDTO) GetReasons() []string {
	if o == nil || IsNil(o.Reasons) {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeComplianceStatusDTO) GetReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *SystemPurposeComplianceStatusDTO) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *SystemPurposeComplianceStatusDTO) SetReasons(v []string) {
	o.Reasons = v
}

func (o SystemPurposeComplianceStatusDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemPurposeComplianceStatusDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Compliant) {
		toSerialize["compliant"] = o.Compliant
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.NonCompliantRole) {
		toSerialize["nonCompliantRole"] = o.NonCompliantRole
	}
	if !IsNil(o.NonCompliantSLA) {
		toSerialize["nonCompliantSLA"] = o.NonCompliantSLA
	}
	if !IsNil(o.NonCompliantUsage) {
		toSerialize["nonCompliantUsage"] = o.NonCompliantUsage
	}
	if !IsNil(o.NonCompliantServiceType) {
		toSerialize["nonCompliantServiceType"] = o.NonCompliantServiceType
	}
	if !IsNil(o.CompliantRole) {
		toSerialize["compliantRole"] = o.CompliantRole
	}
	if !IsNil(o.CompliantAddOns) {
		toSerialize["compliantAddOns"] = o.CompliantAddOns
	}
	if !IsNil(o.CompliantSLA) {
		toSerialize["compliantSLA"] = o.CompliantSLA
	}
	if !IsNil(o.CompliantUsage) {
		toSerialize["compliantUsage"] = o.CompliantUsage
	}
	if !IsNil(o.NonCompliantAddOns) {
		toSerialize["nonCompliantAddOns"] = o.NonCompliantAddOns
	}
	if !IsNil(o.CompliantServiceType) {
		toSerialize["compliantServiceType"] = o.CompliantServiceType
	}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

type NullableSystemPurposeComplianceStatusDTO struct {
	value *SystemPurposeComplianceStatusDTO
	isSet bool
}

func (v NullableSystemPurposeComplianceStatusDTO) Get() *SystemPurposeComplianceStatusDTO {
	return v.value
}

func (v *NullableSystemPurposeComplianceStatusDTO) Set(val *SystemPurposeComplianceStatusDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemPurposeComplianceStatusDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemPurposeComplianceStatusDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemPurposeComplianceStatusDTO(val *SystemPurposeComplianceStatusDTO) *NullableSystemPurposeComplianceStatusDTO {
	return &NullableSystemPurposeComplianceStatusDTO{value: val, isSet: true}
}

func (v NullableSystemPurposeComplianceStatusDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemPurposeComplianceStatusDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


