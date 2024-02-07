/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the OwnerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnerDTO{}

// OwnerDTO DTO representing an owner/organization
type OwnerDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Key *string `json:"key,omitempty"`
	ContentPrefix *string `json:"contentPrefix,omitempty"`
	DefaultServiceLevel *string `json:"defaultServiceLevel,omitempty"`
	LogLevel *string `json:"logLevel,omitempty"`
	ContentAccessMode *string `json:"contentAccessMode,omitempty"`
	ContentAccessModeList *string `json:"contentAccessModeList,omitempty"`
	AutobindHypervisorDisabled *bool `json:"autobindHypervisorDisabled,omitempty"`
	AutobindDisabled *bool `json:"autobindDisabled,omitempty"`
	LastRefreshed *string `json:"lastRefreshed,omitempty"`
	ParentOwner *NestedOwnerDTO `json:"parentOwner,omitempty"`
	UpstreamConsumer *UpstreamConsumerDTO `json:"upstreamConsumer,omitempty"`
	Anonymous *bool `json:"anonymous,omitempty"`
	Claimed *bool `json:"claimed,omitempty"`
	ClaimantOwner *string `json:"claimantOwner,omitempty"`
}

// NewOwnerDTO instantiates a new OwnerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerDTO() *OwnerDTO {
	this := OwnerDTO{}
	return &this
}

// NewOwnerDTOWithDefaults instantiates a new OwnerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerDTOWithDefaults() *OwnerDTO {
	this := OwnerDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OwnerDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OwnerDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *OwnerDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *OwnerDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *OwnerDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *OwnerDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OwnerDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OwnerDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OwnerDTO) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OwnerDTO) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OwnerDTO) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OwnerDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *OwnerDTO) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *OwnerDTO) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *OwnerDTO) SetKey(v string) {
	o.Key = &v
}

// GetContentPrefix returns the ContentPrefix field value if set, zero value otherwise.
func (o *OwnerDTO) GetContentPrefix() string {
	if o == nil || IsNil(o.ContentPrefix) {
		var ret string
		return ret
	}
	return *o.ContentPrefix
}

// GetContentPrefixOk returns a tuple with the ContentPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetContentPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.ContentPrefix) {
		return nil, false
	}
	return o.ContentPrefix, true
}

// HasContentPrefix returns a boolean if a field has been set.
func (o *OwnerDTO) HasContentPrefix() bool {
	if o != nil && !IsNil(o.ContentPrefix) {
		return true
	}

	return false
}

// SetContentPrefix gets a reference to the given string and assigns it to the ContentPrefix field.
func (o *OwnerDTO) SetContentPrefix(v string) {
	o.ContentPrefix = &v
}

// GetDefaultServiceLevel returns the DefaultServiceLevel field value if set, zero value otherwise.
func (o *OwnerDTO) GetDefaultServiceLevel() string {
	if o == nil || IsNil(o.DefaultServiceLevel) {
		var ret string
		return ret
	}
	return *o.DefaultServiceLevel
}

// GetDefaultServiceLevelOk returns a tuple with the DefaultServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetDefaultServiceLevelOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultServiceLevel) {
		return nil, false
	}
	return o.DefaultServiceLevel, true
}

// HasDefaultServiceLevel returns a boolean if a field has been set.
func (o *OwnerDTO) HasDefaultServiceLevel() bool {
	if o != nil && !IsNil(o.DefaultServiceLevel) {
		return true
	}

	return false
}

// SetDefaultServiceLevel gets a reference to the given string and assigns it to the DefaultServiceLevel field.
func (o *OwnerDTO) SetDefaultServiceLevel(v string) {
	o.DefaultServiceLevel = &v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *OwnerDTO) GetLogLevel() string {
	if o == nil || IsNil(o.LogLevel) {
		var ret string
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetLogLevelOk() (*string, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *OwnerDTO) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given string and assigns it to the LogLevel field.
func (o *OwnerDTO) SetLogLevel(v string) {
	o.LogLevel = &v
}

// GetContentAccessMode returns the ContentAccessMode field value if set, zero value otherwise.
func (o *OwnerDTO) GetContentAccessMode() string {
	if o == nil || IsNil(o.ContentAccessMode) {
		var ret string
		return ret
	}
	return *o.ContentAccessMode
}

// GetContentAccessModeOk returns a tuple with the ContentAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetContentAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentAccessMode) {
		return nil, false
	}
	return o.ContentAccessMode, true
}

// HasContentAccessMode returns a boolean if a field has been set.
func (o *OwnerDTO) HasContentAccessMode() bool {
	if o != nil && !IsNil(o.ContentAccessMode) {
		return true
	}

	return false
}

// SetContentAccessMode gets a reference to the given string and assigns it to the ContentAccessMode field.
func (o *OwnerDTO) SetContentAccessMode(v string) {
	o.ContentAccessMode = &v
}

// GetContentAccessModeList returns the ContentAccessModeList field value if set, zero value otherwise.
func (o *OwnerDTO) GetContentAccessModeList() string {
	if o == nil || IsNil(o.ContentAccessModeList) {
		var ret string
		return ret
	}
	return *o.ContentAccessModeList
}

// GetContentAccessModeListOk returns a tuple with the ContentAccessModeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetContentAccessModeListOk() (*string, bool) {
	if o == nil || IsNil(o.ContentAccessModeList) {
		return nil, false
	}
	return o.ContentAccessModeList, true
}

// HasContentAccessModeList returns a boolean if a field has been set.
func (o *OwnerDTO) HasContentAccessModeList() bool {
	if o != nil && !IsNil(o.ContentAccessModeList) {
		return true
	}

	return false
}

// SetContentAccessModeList gets a reference to the given string and assigns it to the ContentAccessModeList field.
func (o *OwnerDTO) SetContentAccessModeList(v string) {
	o.ContentAccessModeList = &v
}

// GetAutobindHypervisorDisabled returns the AutobindHypervisorDisabled field value if set, zero value otherwise.
func (o *OwnerDTO) GetAutobindHypervisorDisabled() bool {
	if o == nil || IsNil(o.AutobindHypervisorDisabled) {
		var ret bool
		return ret
	}
	return *o.AutobindHypervisorDisabled
}

// GetAutobindHypervisorDisabledOk returns a tuple with the AutobindHypervisorDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetAutobindHypervisorDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutobindHypervisorDisabled) {
		return nil, false
	}
	return o.AutobindHypervisorDisabled, true
}

// HasAutobindHypervisorDisabled returns a boolean if a field has been set.
func (o *OwnerDTO) HasAutobindHypervisorDisabled() bool {
	if o != nil && !IsNil(o.AutobindHypervisorDisabled) {
		return true
	}

	return false
}

// SetAutobindHypervisorDisabled gets a reference to the given bool and assigns it to the AutobindHypervisorDisabled field.
func (o *OwnerDTO) SetAutobindHypervisorDisabled(v bool) {
	o.AutobindHypervisorDisabled = &v
}

// GetAutobindDisabled returns the AutobindDisabled field value if set, zero value otherwise.
func (o *OwnerDTO) GetAutobindDisabled() bool {
	if o == nil || IsNil(o.AutobindDisabled) {
		var ret bool
		return ret
	}
	return *o.AutobindDisabled
}

// GetAutobindDisabledOk returns a tuple with the AutobindDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetAutobindDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutobindDisabled) {
		return nil, false
	}
	return o.AutobindDisabled, true
}

// HasAutobindDisabled returns a boolean if a field has been set.
func (o *OwnerDTO) HasAutobindDisabled() bool {
	if o != nil && !IsNil(o.AutobindDisabled) {
		return true
	}

	return false
}

// SetAutobindDisabled gets a reference to the given bool and assigns it to the AutobindDisabled field.
func (o *OwnerDTO) SetAutobindDisabled(v bool) {
	o.AutobindDisabled = &v
}

// GetLastRefreshed returns the LastRefreshed field value if set, zero value otherwise.
func (o *OwnerDTO) GetLastRefreshed() string {
	if o == nil || IsNil(o.LastRefreshed) {
		var ret string
		return ret
	}
	return *o.LastRefreshed
}

// GetLastRefreshedOk returns a tuple with the LastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetLastRefreshedOk() (*string, bool) {
	if o == nil || IsNil(o.LastRefreshed) {
		return nil, false
	}
	return o.LastRefreshed, true
}

// HasLastRefreshed returns a boolean if a field has been set.
func (o *OwnerDTO) HasLastRefreshed() bool {
	if o != nil && !IsNil(o.LastRefreshed) {
		return true
	}

	return false
}

// SetLastRefreshed gets a reference to the given string and assigns it to the LastRefreshed field.
func (o *OwnerDTO) SetLastRefreshed(v string) {
	o.LastRefreshed = &v
}

// GetParentOwner returns the ParentOwner field value if set, zero value otherwise.
func (o *OwnerDTO) GetParentOwner() NestedOwnerDTO {
	if o == nil || IsNil(o.ParentOwner) {
		var ret NestedOwnerDTO
		return ret
	}
	return *o.ParentOwner
}

// GetParentOwnerOk returns a tuple with the ParentOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetParentOwnerOk() (*NestedOwnerDTO, bool) {
	if o == nil || IsNil(o.ParentOwner) {
		return nil, false
	}
	return o.ParentOwner, true
}

// HasParentOwner returns a boolean if a field has been set.
func (o *OwnerDTO) HasParentOwner() bool {
	if o != nil && !IsNil(o.ParentOwner) {
		return true
	}

	return false
}

// SetParentOwner gets a reference to the given NestedOwnerDTO and assigns it to the ParentOwner field.
func (o *OwnerDTO) SetParentOwner(v NestedOwnerDTO) {
	o.ParentOwner = &v
}

// GetUpstreamConsumer returns the UpstreamConsumer field value if set, zero value otherwise.
func (o *OwnerDTO) GetUpstreamConsumer() UpstreamConsumerDTO {
	if o == nil || IsNil(o.UpstreamConsumer) {
		var ret UpstreamConsumerDTO
		return ret
	}
	return *o.UpstreamConsumer
}

// GetUpstreamConsumerOk returns a tuple with the UpstreamConsumer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetUpstreamConsumerOk() (*UpstreamConsumerDTO, bool) {
	if o == nil || IsNil(o.UpstreamConsumer) {
		return nil, false
	}
	return o.UpstreamConsumer, true
}

// HasUpstreamConsumer returns a boolean if a field has been set.
func (o *OwnerDTO) HasUpstreamConsumer() bool {
	if o != nil && !IsNil(o.UpstreamConsumer) {
		return true
	}

	return false
}

// SetUpstreamConsumer gets a reference to the given UpstreamConsumerDTO and assigns it to the UpstreamConsumer field.
func (o *OwnerDTO) SetUpstreamConsumer(v UpstreamConsumerDTO) {
	o.UpstreamConsumer = &v
}

// GetAnonymous returns the Anonymous field value if set, zero value otherwise.
func (o *OwnerDTO) GetAnonymous() bool {
	if o == nil || IsNil(o.Anonymous) {
		var ret bool
		return ret
	}
	return *o.Anonymous
}

// GetAnonymousOk returns a tuple with the Anonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetAnonymousOk() (*bool, bool) {
	if o == nil || IsNil(o.Anonymous) {
		return nil, false
	}
	return o.Anonymous, true
}

// HasAnonymous returns a boolean if a field has been set.
func (o *OwnerDTO) HasAnonymous() bool {
	if o != nil && !IsNil(o.Anonymous) {
		return true
	}

	return false
}

// SetAnonymous gets a reference to the given bool and assigns it to the Anonymous field.
func (o *OwnerDTO) SetAnonymous(v bool) {
	o.Anonymous = &v
}

// GetClaimed returns the Claimed field value if set, zero value otherwise.
func (o *OwnerDTO) GetClaimed() bool {
	if o == nil || IsNil(o.Claimed) {
		var ret bool
		return ret
	}
	return *o.Claimed
}

// GetClaimedOk returns a tuple with the Claimed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetClaimedOk() (*bool, bool) {
	if o == nil || IsNil(o.Claimed) {
		return nil, false
	}
	return o.Claimed, true
}

// HasClaimed returns a boolean if a field has been set.
func (o *OwnerDTO) HasClaimed() bool {
	if o != nil && !IsNil(o.Claimed) {
		return true
	}

	return false
}

// SetClaimed gets a reference to the given bool and assigns it to the Claimed field.
func (o *OwnerDTO) SetClaimed(v bool) {
	o.Claimed = &v
}

// GetClaimantOwner returns the ClaimantOwner field value if set, zero value otherwise.
func (o *OwnerDTO) GetClaimantOwner() string {
	if o == nil || IsNil(o.ClaimantOwner) {
		var ret string
		return ret
	}
	return *o.ClaimantOwner
}

// GetClaimantOwnerOk returns a tuple with the ClaimantOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerDTO) GetClaimantOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.ClaimantOwner) {
		return nil, false
	}
	return o.ClaimantOwner, true
}

// HasClaimantOwner returns a boolean if a field has been set.
func (o *OwnerDTO) HasClaimantOwner() bool {
	if o != nil && !IsNil(o.ClaimantOwner) {
		return true
	}

	return false
}

// SetClaimantOwner gets a reference to the given string and assigns it to the ClaimantOwner field.
func (o *OwnerDTO) SetClaimantOwner(v string) {
	o.ClaimantOwner = &v
}

func (o OwnerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnerDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.ContentPrefix) {
		toSerialize["contentPrefix"] = o.ContentPrefix
	}
	if !IsNil(o.DefaultServiceLevel) {
		toSerialize["defaultServiceLevel"] = o.DefaultServiceLevel
	}
	if !IsNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !IsNil(o.ContentAccessMode) {
		toSerialize["contentAccessMode"] = o.ContentAccessMode
	}
	if !IsNil(o.ContentAccessModeList) {
		toSerialize["contentAccessModeList"] = o.ContentAccessModeList
	}
	if !IsNil(o.AutobindHypervisorDisabled) {
		toSerialize["autobindHypervisorDisabled"] = o.AutobindHypervisorDisabled
	}
	if !IsNil(o.AutobindDisabled) {
		toSerialize["autobindDisabled"] = o.AutobindDisabled
	}
	if !IsNil(o.LastRefreshed) {
		toSerialize["lastRefreshed"] = o.LastRefreshed
	}
	if !IsNil(o.ParentOwner) {
		toSerialize["parentOwner"] = o.ParentOwner
	}
	if !IsNil(o.UpstreamConsumer) {
		toSerialize["upstreamConsumer"] = o.UpstreamConsumer
	}
	if !IsNil(o.Anonymous) {
		toSerialize["anonymous"] = o.Anonymous
	}
	if !IsNil(o.Claimed) {
		toSerialize["claimed"] = o.Claimed
	}
	if !IsNil(o.ClaimantOwner) {
		toSerialize["claimantOwner"] = o.ClaimantOwner
	}
	return toSerialize, nil
}

type NullableOwnerDTO struct {
	value *OwnerDTO
	isSet bool
}

func (v NullableOwnerDTO) Get() *OwnerDTO {
	return v.value
}

func (v *NullableOwnerDTO) Set(val *OwnerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerDTO(val *OwnerDTO) *NullableOwnerDTO {
	return &NullableOwnerDTO{value: val, isSet: true}
}

func (v NullableOwnerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


