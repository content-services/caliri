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

// checks if the ConsumerDTOArrayElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsumerDTOArrayElement{}

// ConsumerDTOArrayElement Represents a reduced view of a consumer (excluding data like facts and identify certificate)
type ConsumerDTOArrayElement struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Name *string `json:"name,omitempty"`
	Username *string `json:"username,omitempty"`
	EntitlementStatus *string `json:"entitlementStatus,omitempty"`
	ServiceLevel *string `json:"serviceLevel,omitempty"`
	Role *string `json:"role,omitempty"`
	Usage *string `json:"usage,omitempty"`
	AddOns []string `json:"addOns,omitempty"`
	SystemPurposeStatus *string `json:"systemPurposeStatus,omitempty"`
	ReleaseVer *ReleaseVerDTO `json:"releaseVer,omitempty"`
	Owner *NestedOwnerDTO `json:"owner,omitempty"`
	EntitlementCount *int64 `json:"entitlementCount,omitempty"`
	LastCheckin *string `json:"lastCheckin,omitempty"`
	InstalledProducts []ConsumerInstalledProductDTO `json:"installedProducts,omitempty"`
	CanActivate *bool `json:"canActivate,omitempty"`
	Capabilities []CapabilityDTO `json:"capabilities,omitempty"`
	HypervisorId *HypervisorIdDTO `json:"hypervisorId,omitempty"`
	ContentTags []string `json:"contentTags,omitempty"`
	Autoheal *bool `json:"autoheal,omitempty"`
	Annotations *string `json:"annotations,omitempty"`
	ContentAccessMode *string `json:"contentAccessMode,omitempty"`
	Type *ConsumerTypeDTO `json:"type,omitempty"`
	GuestIds []GuestIdDTOArrayElement `json:"guestIds,omitempty"`
	Href *string `json:"href,omitempty"`
	ServiceType *string `json:"serviceType,omitempty"`
}

// NewConsumerDTOArrayElement instantiates a new ConsumerDTOArrayElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerDTOArrayElement() *ConsumerDTOArrayElement {
	this := ConsumerDTOArrayElement{}
	return &this
}

// NewConsumerDTOArrayElementWithDefaults instantiates a new ConsumerDTOArrayElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerDTOArrayElementWithDefaults() *ConsumerDTOArrayElement {
	this := ConsumerDTOArrayElement{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ConsumerDTOArrayElement) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *ConsumerDTOArrayElement) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConsumerDTOArrayElement) SetId(v string) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ConsumerDTOArrayElement) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConsumerDTOArrayElement) SetName(v string) {
	o.Name = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ConsumerDTOArrayElement) SetUsername(v string) {
	o.Username = &v
}

// GetEntitlementStatus returns the EntitlementStatus field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetEntitlementStatus() string {
	if o == nil || IsNil(o.EntitlementStatus) {
		var ret string
		return ret
	}
	return *o.EntitlementStatus
}

// GetEntitlementStatusOk returns a tuple with the EntitlementStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetEntitlementStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementStatus) {
		return nil, false
	}
	return o.EntitlementStatus, true
}

// HasEntitlementStatus returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasEntitlementStatus() bool {
	if o != nil && !IsNil(o.EntitlementStatus) {
		return true
	}

	return false
}

// SetEntitlementStatus gets a reference to the given string and assigns it to the EntitlementStatus field.
func (o *ConsumerDTOArrayElement) SetEntitlementStatus(v string) {
	o.EntitlementStatus = &v
}

// GetServiceLevel returns the ServiceLevel field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetServiceLevel() string {
	if o == nil || IsNil(o.ServiceLevel) {
		var ret string
		return ret
	}
	return *o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetServiceLevelOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLevel) {
		return nil, false
	}
	return o.ServiceLevel, true
}

// HasServiceLevel returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasServiceLevel() bool {
	if o != nil && !IsNil(o.ServiceLevel) {
		return true
	}

	return false
}

// SetServiceLevel gets a reference to the given string and assigns it to the ServiceLevel field.
func (o *ConsumerDTOArrayElement) SetServiceLevel(v string) {
	o.ServiceLevel = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ConsumerDTOArrayElement) SetRole(v string) {
	o.Role = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetUsage() string {
	if o == nil || IsNil(o.Usage) {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetUsageOk() (*string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *ConsumerDTOArrayElement) SetUsage(v string) {
	o.Usage = &v
}

// GetAddOns returns the AddOns field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetAddOns() []string {
	if o == nil || IsNil(o.AddOns) {
		var ret []string
		return ret
	}
	return o.AddOns
}

// GetAddOnsOk returns a tuple with the AddOns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetAddOnsOk() ([]string, bool) {
	if o == nil || IsNil(o.AddOns) {
		return nil, false
	}
	return o.AddOns, true
}

// HasAddOns returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasAddOns() bool {
	if o != nil && !IsNil(o.AddOns) {
		return true
	}

	return false
}

// SetAddOns gets a reference to the given []string and assigns it to the AddOns field.
func (o *ConsumerDTOArrayElement) SetAddOns(v []string) {
	o.AddOns = v
}

// GetSystemPurposeStatus returns the SystemPurposeStatus field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetSystemPurposeStatus() string {
	if o == nil || IsNil(o.SystemPurposeStatus) {
		var ret string
		return ret
	}
	return *o.SystemPurposeStatus
}

// GetSystemPurposeStatusOk returns a tuple with the SystemPurposeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetSystemPurposeStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SystemPurposeStatus) {
		return nil, false
	}
	return o.SystemPurposeStatus, true
}

// HasSystemPurposeStatus returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasSystemPurposeStatus() bool {
	if o != nil && !IsNil(o.SystemPurposeStatus) {
		return true
	}

	return false
}

// SetSystemPurposeStatus gets a reference to the given string and assigns it to the SystemPurposeStatus field.
func (o *ConsumerDTOArrayElement) SetSystemPurposeStatus(v string) {
	o.SystemPurposeStatus = &v
}

// GetReleaseVer returns the ReleaseVer field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetReleaseVer() ReleaseVerDTO {
	if o == nil || IsNil(o.ReleaseVer) {
		var ret ReleaseVerDTO
		return ret
	}
	return *o.ReleaseVer
}

// GetReleaseVerOk returns a tuple with the ReleaseVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetReleaseVerOk() (*ReleaseVerDTO, bool) {
	if o == nil || IsNil(o.ReleaseVer) {
		return nil, false
	}
	return o.ReleaseVer, true
}

// HasReleaseVer returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasReleaseVer() bool {
	if o != nil && !IsNil(o.ReleaseVer) {
		return true
	}

	return false
}

// SetReleaseVer gets a reference to the given ReleaseVerDTO and assigns it to the ReleaseVer field.
func (o *ConsumerDTOArrayElement) SetReleaseVer(v ReleaseVerDTO) {
	o.ReleaseVer = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetOwner() NestedOwnerDTO {
	if o == nil || IsNil(o.Owner) {
		var ret NestedOwnerDTO
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetOwnerOk() (*NestedOwnerDTO, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NestedOwnerDTO and assigns it to the Owner field.
func (o *ConsumerDTOArrayElement) SetOwner(v NestedOwnerDTO) {
	o.Owner = &v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetEntitlementCount() int64 {
	if o == nil || IsNil(o.EntitlementCount) {
		var ret int64
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetEntitlementCountOk() (*int64, bool) {
	if o == nil || IsNil(o.EntitlementCount) {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasEntitlementCount() bool {
	if o != nil && !IsNil(o.EntitlementCount) {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int64 and assigns it to the EntitlementCount field.
func (o *ConsumerDTOArrayElement) SetEntitlementCount(v int64) {
	o.EntitlementCount = &v
}

// GetLastCheckin returns the LastCheckin field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetLastCheckin() string {
	if o == nil || IsNil(o.LastCheckin) {
		var ret string
		return ret
	}
	return *o.LastCheckin
}

// GetLastCheckinOk returns a tuple with the LastCheckin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetLastCheckinOk() (*string, bool) {
	if o == nil || IsNil(o.LastCheckin) {
		return nil, false
	}
	return o.LastCheckin, true
}

// HasLastCheckin returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasLastCheckin() bool {
	if o != nil && !IsNil(o.LastCheckin) {
		return true
	}

	return false
}

// SetLastCheckin gets a reference to the given string and assigns it to the LastCheckin field.
func (o *ConsumerDTOArrayElement) SetLastCheckin(v string) {
	o.LastCheckin = &v
}

// GetInstalledProducts returns the InstalledProducts field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetInstalledProducts() []ConsumerInstalledProductDTO {
	if o == nil || IsNil(o.InstalledProducts) {
		var ret []ConsumerInstalledProductDTO
		return ret
	}
	return o.InstalledProducts
}

// GetInstalledProductsOk returns a tuple with the InstalledProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetInstalledProductsOk() ([]ConsumerInstalledProductDTO, bool) {
	if o == nil || IsNil(o.InstalledProducts) {
		return nil, false
	}
	return o.InstalledProducts, true
}

// HasInstalledProducts returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasInstalledProducts() bool {
	if o != nil && !IsNil(o.InstalledProducts) {
		return true
	}

	return false
}

// SetInstalledProducts gets a reference to the given []ConsumerInstalledProductDTO and assigns it to the InstalledProducts field.
func (o *ConsumerDTOArrayElement) SetInstalledProducts(v []ConsumerInstalledProductDTO) {
	o.InstalledProducts = v
}

// GetCanActivate returns the CanActivate field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetCanActivate() bool {
	if o == nil || IsNil(o.CanActivate) {
		var ret bool
		return ret
	}
	return *o.CanActivate
}

// GetCanActivateOk returns a tuple with the CanActivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetCanActivateOk() (*bool, bool) {
	if o == nil || IsNil(o.CanActivate) {
		return nil, false
	}
	return o.CanActivate, true
}

// HasCanActivate returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasCanActivate() bool {
	if o != nil && !IsNil(o.CanActivate) {
		return true
	}

	return false
}

// SetCanActivate gets a reference to the given bool and assigns it to the CanActivate field.
func (o *ConsumerDTOArrayElement) SetCanActivate(v bool) {
	o.CanActivate = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetCapabilities() []CapabilityDTO {
	if o == nil || IsNil(o.Capabilities) {
		var ret []CapabilityDTO
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetCapabilitiesOk() ([]CapabilityDTO, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []CapabilityDTO and assigns it to the Capabilities field.
func (o *ConsumerDTOArrayElement) SetCapabilities(v []CapabilityDTO) {
	o.Capabilities = v
}

// GetHypervisorId returns the HypervisorId field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetHypervisorId() HypervisorIdDTO {
	if o == nil || IsNil(o.HypervisorId) {
		var ret HypervisorIdDTO
		return ret
	}
	return *o.HypervisorId
}

// GetHypervisorIdOk returns a tuple with the HypervisorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetHypervisorIdOk() (*HypervisorIdDTO, bool) {
	if o == nil || IsNil(o.HypervisorId) {
		return nil, false
	}
	return o.HypervisorId, true
}

// HasHypervisorId returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasHypervisorId() bool {
	if o != nil && !IsNil(o.HypervisorId) {
		return true
	}

	return false
}

// SetHypervisorId gets a reference to the given HypervisorIdDTO and assigns it to the HypervisorId field.
func (o *ConsumerDTOArrayElement) SetHypervisorId(v HypervisorIdDTO) {
	o.HypervisorId = &v
}

// GetContentTags returns the ContentTags field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetContentTags() []string {
	if o == nil || IsNil(o.ContentTags) {
		var ret []string
		return ret
	}
	return o.ContentTags
}

// GetContentTagsOk returns a tuple with the ContentTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetContentTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentTags) {
		return nil, false
	}
	return o.ContentTags, true
}

// HasContentTags returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasContentTags() bool {
	if o != nil && !IsNil(o.ContentTags) {
		return true
	}

	return false
}

// SetContentTags gets a reference to the given []string and assigns it to the ContentTags field.
func (o *ConsumerDTOArrayElement) SetContentTags(v []string) {
	o.ContentTags = v
}

// GetAutoheal returns the Autoheal field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetAutoheal() bool {
	if o == nil || IsNil(o.Autoheal) {
		var ret bool
		return ret
	}
	return *o.Autoheal
}

// GetAutohealOk returns a tuple with the Autoheal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetAutohealOk() (*bool, bool) {
	if o == nil || IsNil(o.Autoheal) {
		return nil, false
	}
	return o.Autoheal, true
}

// HasAutoheal returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasAutoheal() bool {
	if o != nil && !IsNil(o.Autoheal) {
		return true
	}

	return false
}

// SetAutoheal gets a reference to the given bool and assigns it to the Autoheal field.
func (o *ConsumerDTOArrayElement) SetAutoheal(v bool) {
	o.Autoheal = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetAnnotations() string {
	if o == nil || IsNil(o.Annotations) {
		var ret string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetAnnotationsOk() (*string, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given string and assigns it to the Annotations field.
func (o *ConsumerDTOArrayElement) SetAnnotations(v string) {
	o.Annotations = &v
}

// GetContentAccessMode returns the ContentAccessMode field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetContentAccessMode() string {
	if o == nil || IsNil(o.ContentAccessMode) {
		var ret string
		return ret
	}
	return *o.ContentAccessMode
}

// GetContentAccessModeOk returns a tuple with the ContentAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetContentAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentAccessMode) {
		return nil, false
	}
	return o.ContentAccessMode, true
}

// HasContentAccessMode returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasContentAccessMode() bool {
	if o != nil && !IsNil(o.ContentAccessMode) {
		return true
	}

	return false
}

// SetContentAccessMode gets a reference to the given string and assigns it to the ContentAccessMode field.
func (o *ConsumerDTOArrayElement) SetContentAccessMode(v string) {
	o.ContentAccessMode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetType() ConsumerTypeDTO {
	if o == nil || IsNil(o.Type) {
		var ret ConsumerTypeDTO
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetTypeOk() (*ConsumerTypeDTO, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ConsumerTypeDTO and assigns it to the Type field.
func (o *ConsumerDTOArrayElement) SetType(v ConsumerTypeDTO) {
	o.Type = &v
}

// GetGuestIds returns the GuestIds field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetGuestIds() []GuestIdDTOArrayElement {
	if o == nil || IsNil(o.GuestIds) {
		var ret []GuestIdDTOArrayElement
		return ret
	}
	return o.GuestIds
}

// GetGuestIdsOk returns a tuple with the GuestIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetGuestIdsOk() ([]GuestIdDTOArrayElement, bool) {
	if o == nil || IsNil(o.GuestIds) {
		return nil, false
	}
	return o.GuestIds, true
}

// HasGuestIds returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasGuestIds() bool {
	if o != nil && !IsNil(o.GuestIds) {
		return true
	}

	return false
}

// SetGuestIds gets a reference to the given []GuestIdDTOArrayElement and assigns it to the GuestIds field.
func (o *ConsumerDTOArrayElement) SetGuestIds(v []GuestIdDTOArrayElement) {
	o.GuestIds = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ConsumerDTOArrayElement) SetHref(v string) {
	o.Href = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *ConsumerDTOArrayElement) GetServiceType() string {
	if o == nil || IsNil(o.ServiceType) {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerDTOArrayElement) GetServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *ConsumerDTOArrayElement) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *ConsumerDTOArrayElement) SetServiceType(v string) {
	o.ServiceType = &v
}

func (o ConsumerDTOArrayElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsumerDTOArrayElement) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.EntitlementStatus) {
		toSerialize["entitlementStatus"] = o.EntitlementStatus
	}
	if !IsNil(o.ServiceLevel) {
		toSerialize["serviceLevel"] = o.ServiceLevel
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.AddOns) {
		toSerialize["addOns"] = o.AddOns
	}
	if !IsNil(o.SystemPurposeStatus) {
		toSerialize["systemPurposeStatus"] = o.SystemPurposeStatus
	}
	if !IsNil(o.ReleaseVer) {
		toSerialize["releaseVer"] = o.ReleaseVer
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.EntitlementCount) {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if !IsNil(o.LastCheckin) {
		toSerialize["lastCheckin"] = o.LastCheckin
	}
	if !IsNil(o.InstalledProducts) {
		toSerialize["installedProducts"] = o.InstalledProducts
	}
	if !IsNil(o.CanActivate) {
		toSerialize["canActivate"] = o.CanActivate
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.HypervisorId) {
		toSerialize["hypervisorId"] = o.HypervisorId
	}
	if !IsNil(o.ContentTags) {
		toSerialize["contentTags"] = o.ContentTags
	}
	if !IsNil(o.Autoheal) {
		toSerialize["autoheal"] = o.Autoheal
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.ContentAccessMode) {
		toSerialize["contentAccessMode"] = o.ContentAccessMode
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.GuestIds) {
		toSerialize["guestIds"] = o.GuestIds
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.ServiceType) {
		toSerialize["serviceType"] = o.ServiceType
	}
	return toSerialize, nil
}

type NullableConsumerDTOArrayElement struct {
	value *ConsumerDTOArrayElement
	isSet bool
}

func (v NullableConsumerDTOArrayElement) Get() *ConsumerDTOArrayElement {
	return v.value
}

func (v *NullableConsumerDTOArrayElement) Set(val *ConsumerDTOArrayElement) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerDTOArrayElement) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerDTOArrayElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerDTOArrayElement(val *ConsumerDTOArrayElement) *NullableConsumerDTOArrayElement {
	return &NullableConsumerDTOArrayElement{value: val, isSet: true}
}

func (v NullableConsumerDTOArrayElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerDTOArrayElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


