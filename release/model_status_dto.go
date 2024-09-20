/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the StatusDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusDTO{}

// StatusDTO Version and Status information about running Candlepin server
type StatusDTO struct {
	Mode *string `json:"mode,omitempty"`
	ModeReason *string `json:"modeReason,omitempty"`
	ModeChangeTime *string `json:"modeChangeTime,omitempty"`
	Result *bool `json:"result,omitempty"`
	Version *string `json:"version,omitempty"`
	Release *string `json:"release,omitempty"`
	Standalone *bool `json:"standalone,omitempty"`
	TimeUTC *string `json:"timeUTC,omitempty"`
	RulesSource *string `json:"rulesSource,omitempty"`
	RulesVersion *string `json:"rulesVersion,omitempty"`
	ManagerCapabilities []string `json:"managerCapabilities,omitempty"`
	KeycloakRealm *string `json:"keycloakRealm,omitempty"`
	KeycloakAuthUrl *string `json:"keycloakAuthUrl,omitempty"`
	KeycloakResource *string `json:"keycloakResource,omitempty"`
	DeviceAuthRealm *string `json:"deviceAuthRealm,omitempty"`
	DeviceAuthUrl *string `json:"deviceAuthUrl,omitempty"`
	DeviceAuthClientId *string `json:"deviceAuthClientId,omitempty"`
	DeviceAuthScope *string `json:"deviceAuthScope,omitempty"`
}

// NewStatusDTO instantiates a new StatusDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusDTO() *StatusDTO {
	this := StatusDTO{}
	return &this
}

// NewStatusDTOWithDefaults instantiates a new StatusDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusDTOWithDefaults() *StatusDTO {
	this := StatusDTO{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *StatusDTO) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *StatusDTO) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *StatusDTO) SetMode(v string) {
	o.Mode = &v
}

// GetModeReason returns the ModeReason field value if set, zero value otherwise.
func (o *StatusDTO) GetModeReason() string {
	if o == nil || IsNil(o.ModeReason) {
		var ret string
		return ret
	}
	return *o.ModeReason
}

// GetModeReasonOk returns a tuple with the ModeReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetModeReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ModeReason) {
		return nil, false
	}
	return o.ModeReason, true
}

// HasModeReason returns a boolean if a field has been set.
func (o *StatusDTO) HasModeReason() bool {
	if o != nil && !IsNil(o.ModeReason) {
		return true
	}

	return false
}

// SetModeReason gets a reference to the given string and assigns it to the ModeReason field.
func (o *StatusDTO) SetModeReason(v string) {
	o.ModeReason = &v
}

// GetModeChangeTime returns the ModeChangeTime field value if set, zero value otherwise.
func (o *StatusDTO) GetModeChangeTime() string {
	if o == nil || IsNil(o.ModeChangeTime) {
		var ret string
		return ret
	}
	return *o.ModeChangeTime
}

// GetModeChangeTimeOk returns a tuple with the ModeChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetModeChangeTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ModeChangeTime) {
		return nil, false
	}
	return o.ModeChangeTime, true
}

// HasModeChangeTime returns a boolean if a field has been set.
func (o *StatusDTO) HasModeChangeTime() bool {
	if o != nil && !IsNil(o.ModeChangeTime) {
		return true
	}

	return false
}

// SetModeChangeTime gets a reference to the given string and assigns it to the ModeChangeTime field.
func (o *StatusDTO) SetModeChangeTime(v string) {
	o.ModeChangeTime = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *StatusDTO) GetResult() bool {
	if o == nil || IsNil(o.Result) {
		var ret bool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetResultOk() (*bool, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *StatusDTO) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given bool and assigns it to the Result field.
func (o *StatusDTO) SetResult(v bool) {
	o.Result = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StatusDTO) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StatusDTO) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StatusDTO) SetVersion(v string) {
	o.Version = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *StatusDTO) GetRelease() string {
	if o == nil || IsNil(o.Release) {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetReleaseOk() (*string, bool) {
	if o == nil || IsNil(o.Release) {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *StatusDTO) HasRelease() bool {
	if o != nil && !IsNil(o.Release) {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *StatusDTO) SetRelease(v string) {
	o.Release = &v
}

// GetStandalone returns the Standalone field value if set, zero value otherwise.
func (o *StatusDTO) GetStandalone() bool {
	if o == nil || IsNil(o.Standalone) {
		var ret bool
		return ret
	}
	return *o.Standalone
}

// GetStandaloneOk returns a tuple with the Standalone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetStandaloneOk() (*bool, bool) {
	if o == nil || IsNil(o.Standalone) {
		return nil, false
	}
	return o.Standalone, true
}

// HasStandalone returns a boolean if a field has been set.
func (o *StatusDTO) HasStandalone() bool {
	if o != nil && !IsNil(o.Standalone) {
		return true
	}

	return false
}

// SetStandalone gets a reference to the given bool and assigns it to the Standalone field.
func (o *StatusDTO) SetStandalone(v bool) {
	o.Standalone = &v
}

// GetTimeUTC returns the TimeUTC field value if set, zero value otherwise.
func (o *StatusDTO) GetTimeUTC() string {
	if o == nil || IsNil(o.TimeUTC) {
		var ret string
		return ret
	}
	return *o.TimeUTC
}

// GetTimeUTCOk returns a tuple with the TimeUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetTimeUTCOk() (*string, bool) {
	if o == nil || IsNil(o.TimeUTC) {
		return nil, false
	}
	return o.TimeUTC, true
}

// HasTimeUTC returns a boolean if a field has been set.
func (o *StatusDTO) HasTimeUTC() bool {
	if o != nil && !IsNil(o.TimeUTC) {
		return true
	}

	return false
}

// SetTimeUTC gets a reference to the given string and assigns it to the TimeUTC field.
func (o *StatusDTO) SetTimeUTC(v string) {
	o.TimeUTC = &v
}

// GetRulesSource returns the RulesSource field value if set, zero value otherwise.
func (o *StatusDTO) GetRulesSource() string {
	if o == nil || IsNil(o.RulesSource) {
		var ret string
		return ret
	}
	return *o.RulesSource
}

// GetRulesSourceOk returns a tuple with the RulesSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetRulesSourceOk() (*string, bool) {
	if o == nil || IsNil(o.RulesSource) {
		return nil, false
	}
	return o.RulesSource, true
}

// HasRulesSource returns a boolean if a field has been set.
func (o *StatusDTO) HasRulesSource() bool {
	if o != nil && !IsNil(o.RulesSource) {
		return true
	}

	return false
}

// SetRulesSource gets a reference to the given string and assigns it to the RulesSource field.
func (o *StatusDTO) SetRulesSource(v string) {
	o.RulesSource = &v
}

// GetRulesVersion returns the RulesVersion field value if set, zero value otherwise.
func (o *StatusDTO) GetRulesVersion() string {
	if o == nil || IsNil(o.RulesVersion) {
		var ret string
		return ret
	}
	return *o.RulesVersion
}

// GetRulesVersionOk returns a tuple with the RulesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetRulesVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RulesVersion) {
		return nil, false
	}
	return o.RulesVersion, true
}

// HasRulesVersion returns a boolean if a field has been set.
func (o *StatusDTO) HasRulesVersion() bool {
	if o != nil && !IsNil(o.RulesVersion) {
		return true
	}

	return false
}

// SetRulesVersion gets a reference to the given string and assigns it to the RulesVersion field.
func (o *StatusDTO) SetRulesVersion(v string) {
	o.RulesVersion = &v
}

// GetManagerCapabilities returns the ManagerCapabilities field value if set, zero value otherwise.
func (o *StatusDTO) GetManagerCapabilities() []string {
	if o == nil || IsNil(o.ManagerCapabilities) {
		var ret []string
		return ret
	}
	return o.ManagerCapabilities
}

// GetManagerCapabilitiesOk returns a tuple with the ManagerCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetManagerCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.ManagerCapabilities) {
		return nil, false
	}
	return o.ManagerCapabilities, true
}

// HasManagerCapabilities returns a boolean if a field has been set.
func (o *StatusDTO) HasManagerCapabilities() bool {
	if o != nil && !IsNil(o.ManagerCapabilities) {
		return true
	}

	return false
}

// SetManagerCapabilities gets a reference to the given []string and assigns it to the ManagerCapabilities field.
func (o *StatusDTO) SetManagerCapabilities(v []string) {
	o.ManagerCapabilities = v
}

// GetKeycloakRealm returns the KeycloakRealm field value if set, zero value otherwise.
func (o *StatusDTO) GetKeycloakRealm() string {
	if o == nil || IsNil(o.KeycloakRealm) {
		var ret string
		return ret
	}
	return *o.KeycloakRealm
}

// GetKeycloakRealmOk returns a tuple with the KeycloakRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetKeycloakRealmOk() (*string, bool) {
	if o == nil || IsNil(o.KeycloakRealm) {
		return nil, false
	}
	return o.KeycloakRealm, true
}

// HasKeycloakRealm returns a boolean if a field has been set.
func (o *StatusDTO) HasKeycloakRealm() bool {
	if o != nil && !IsNil(o.KeycloakRealm) {
		return true
	}

	return false
}

// SetKeycloakRealm gets a reference to the given string and assigns it to the KeycloakRealm field.
func (o *StatusDTO) SetKeycloakRealm(v string) {
	o.KeycloakRealm = &v
}

// GetKeycloakAuthUrl returns the KeycloakAuthUrl field value if set, zero value otherwise.
func (o *StatusDTO) GetKeycloakAuthUrl() string {
	if o == nil || IsNil(o.KeycloakAuthUrl) {
		var ret string
		return ret
	}
	return *o.KeycloakAuthUrl
}

// GetKeycloakAuthUrlOk returns a tuple with the KeycloakAuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetKeycloakAuthUrlOk() (*string, bool) {
	if o == nil || IsNil(o.KeycloakAuthUrl) {
		return nil, false
	}
	return o.KeycloakAuthUrl, true
}

// HasKeycloakAuthUrl returns a boolean if a field has been set.
func (o *StatusDTO) HasKeycloakAuthUrl() bool {
	if o != nil && !IsNil(o.KeycloakAuthUrl) {
		return true
	}

	return false
}

// SetKeycloakAuthUrl gets a reference to the given string and assigns it to the KeycloakAuthUrl field.
func (o *StatusDTO) SetKeycloakAuthUrl(v string) {
	o.KeycloakAuthUrl = &v
}

// GetKeycloakResource returns the KeycloakResource field value if set, zero value otherwise.
func (o *StatusDTO) GetKeycloakResource() string {
	if o == nil || IsNil(o.KeycloakResource) {
		var ret string
		return ret
	}
	return *o.KeycloakResource
}

// GetKeycloakResourceOk returns a tuple with the KeycloakResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetKeycloakResourceOk() (*string, bool) {
	if o == nil || IsNil(o.KeycloakResource) {
		return nil, false
	}
	return o.KeycloakResource, true
}

// HasKeycloakResource returns a boolean if a field has been set.
func (o *StatusDTO) HasKeycloakResource() bool {
	if o != nil && !IsNil(o.KeycloakResource) {
		return true
	}

	return false
}

// SetKeycloakResource gets a reference to the given string and assigns it to the KeycloakResource field.
func (o *StatusDTO) SetKeycloakResource(v string) {
	o.KeycloakResource = &v
}

// GetDeviceAuthRealm returns the DeviceAuthRealm field value if set, zero value otherwise.
func (o *StatusDTO) GetDeviceAuthRealm() string {
	if o == nil || IsNil(o.DeviceAuthRealm) {
		var ret string
		return ret
	}
	return *o.DeviceAuthRealm
}

// GetDeviceAuthRealmOk returns a tuple with the DeviceAuthRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetDeviceAuthRealmOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceAuthRealm) {
		return nil, false
	}
	return o.DeviceAuthRealm, true
}

// HasDeviceAuthRealm returns a boolean if a field has been set.
func (o *StatusDTO) HasDeviceAuthRealm() bool {
	if o != nil && !IsNil(o.DeviceAuthRealm) {
		return true
	}

	return false
}

// SetDeviceAuthRealm gets a reference to the given string and assigns it to the DeviceAuthRealm field.
func (o *StatusDTO) SetDeviceAuthRealm(v string) {
	o.DeviceAuthRealm = &v
}

// GetDeviceAuthUrl returns the DeviceAuthUrl field value if set, zero value otherwise.
func (o *StatusDTO) GetDeviceAuthUrl() string {
	if o == nil || IsNil(o.DeviceAuthUrl) {
		var ret string
		return ret
	}
	return *o.DeviceAuthUrl
}

// GetDeviceAuthUrlOk returns a tuple with the DeviceAuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetDeviceAuthUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceAuthUrl) {
		return nil, false
	}
	return o.DeviceAuthUrl, true
}

// HasDeviceAuthUrl returns a boolean if a field has been set.
func (o *StatusDTO) HasDeviceAuthUrl() bool {
	if o != nil && !IsNil(o.DeviceAuthUrl) {
		return true
	}

	return false
}

// SetDeviceAuthUrl gets a reference to the given string and assigns it to the DeviceAuthUrl field.
func (o *StatusDTO) SetDeviceAuthUrl(v string) {
	o.DeviceAuthUrl = &v
}

// GetDeviceAuthClientId returns the DeviceAuthClientId field value if set, zero value otherwise.
func (o *StatusDTO) GetDeviceAuthClientId() string {
	if o == nil || IsNil(o.DeviceAuthClientId) {
		var ret string
		return ret
	}
	return *o.DeviceAuthClientId
}

// GetDeviceAuthClientIdOk returns a tuple with the DeviceAuthClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetDeviceAuthClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceAuthClientId) {
		return nil, false
	}
	return o.DeviceAuthClientId, true
}

// HasDeviceAuthClientId returns a boolean if a field has been set.
func (o *StatusDTO) HasDeviceAuthClientId() bool {
	if o != nil && !IsNil(o.DeviceAuthClientId) {
		return true
	}

	return false
}

// SetDeviceAuthClientId gets a reference to the given string and assigns it to the DeviceAuthClientId field.
func (o *StatusDTO) SetDeviceAuthClientId(v string) {
	o.DeviceAuthClientId = &v
}

// GetDeviceAuthScope returns the DeviceAuthScope field value if set, zero value otherwise.
func (o *StatusDTO) GetDeviceAuthScope() string {
	if o == nil || IsNil(o.DeviceAuthScope) {
		var ret string
		return ret
	}
	return *o.DeviceAuthScope
}

// GetDeviceAuthScopeOk returns a tuple with the DeviceAuthScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusDTO) GetDeviceAuthScopeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceAuthScope) {
		return nil, false
	}
	return o.DeviceAuthScope, true
}

// HasDeviceAuthScope returns a boolean if a field has been set.
func (o *StatusDTO) HasDeviceAuthScope() bool {
	if o != nil && !IsNil(o.DeviceAuthScope) {
		return true
	}

	return false
}

// SetDeviceAuthScope gets a reference to the given string and assigns it to the DeviceAuthScope field.
func (o *StatusDTO) SetDeviceAuthScope(v string) {
	o.DeviceAuthScope = &v
}

func (o StatusDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.ModeReason) {
		toSerialize["modeReason"] = o.ModeReason
	}
	if !IsNil(o.ModeChangeTime) {
		toSerialize["modeChangeTime"] = o.ModeChangeTime
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Release) {
		toSerialize["release"] = o.Release
	}
	if !IsNil(o.Standalone) {
		toSerialize["standalone"] = o.Standalone
	}
	if !IsNil(o.TimeUTC) {
		toSerialize["timeUTC"] = o.TimeUTC
	}
	if !IsNil(o.RulesSource) {
		toSerialize["rulesSource"] = o.RulesSource
	}
	if !IsNil(o.RulesVersion) {
		toSerialize["rulesVersion"] = o.RulesVersion
	}
	if !IsNil(o.ManagerCapabilities) {
		toSerialize["managerCapabilities"] = o.ManagerCapabilities
	}
	if !IsNil(o.KeycloakRealm) {
		toSerialize["keycloakRealm"] = o.KeycloakRealm
	}
	if !IsNil(o.KeycloakAuthUrl) {
		toSerialize["keycloakAuthUrl"] = o.KeycloakAuthUrl
	}
	if !IsNil(o.KeycloakResource) {
		toSerialize["keycloakResource"] = o.KeycloakResource
	}
	if !IsNil(o.DeviceAuthRealm) {
		toSerialize["deviceAuthRealm"] = o.DeviceAuthRealm
	}
	if !IsNil(o.DeviceAuthUrl) {
		toSerialize["deviceAuthUrl"] = o.DeviceAuthUrl
	}
	if !IsNil(o.DeviceAuthClientId) {
		toSerialize["deviceAuthClientId"] = o.DeviceAuthClientId
	}
	if !IsNil(o.DeviceAuthScope) {
		toSerialize["deviceAuthScope"] = o.DeviceAuthScope
	}
	return toSerialize, nil
}

type NullableStatusDTO struct {
	value *StatusDTO
	isSet bool
}

func (v NullableStatusDTO) Get() *StatusDTO {
	return v.value
}

func (v *NullableStatusDTO) Set(val *StatusDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusDTO(val *StatusDTO) *NullableStatusDTO {
	return &NullableStatusDTO{value: val, isSet: true}
}

func (v NullableStatusDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


