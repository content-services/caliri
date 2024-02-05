# ConsumerDTOArrayElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**EntitlementStatus** | Pointer to **string** |  | [optional] 
**ServiceLevel** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to **string** |  | [optional] 
**AddOns** | Pointer to **[]string** |  | [optional] 
**SystemPurposeStatus** | Pointer to **string** |  | [optional] 
**ReleaseVer** | Pointer to [**ReleaseVerDTO**](ReleaseVerDTO.md) |  | [optional] 
**Owner** | Pointer to [**NestedOwnerDTO**](NestedOwnerDTO.md) |  | [optional] 
**EntitlementCount** | Pointer to **int64** |  | [optional] 
**LastCheckin** | Pointer to **time.Time** |  | [optional] 
**InstalledProducts** | Pointer to [**[]ConsumerInstalledProductDTO**](ConsumerInstalledProductDTO.md) |  | [optional] 
**CanActivate** | Pointer to **bool** |  | [optional] 
**Capabilities** | Pointer to [**[]CapabilityDTO**](CapabilityDTO.md) |  | [optional] 
**HypervisorId** | Pointer to [**HypervisorIdDTO**](HypervisorIdDTO.md) |  | [optional] 
**ContentTags** | Pointer to **[]string** |  | [optional] 
**Autoheal** | Pointer to **bool** |  | [optional] 
**Annotations** | Pointer to **string** |  | [optional] 
**ContentAccessMode** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ConsumerTypeDTO**](ConsumerTypeDTO.md) |  | [optional] 
**GuestIds** | Pointer to [**[]GuestIdDTOArrayElement**](GuestIdDTOArrayElement.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 

## Methods

### NewConsumerDTOArrayElement

`func NewConsumerDTOArrayElement() *ConsumerDTOArrayElement`

NewConsumerDTOArrayElement instantiates a new ConsumerDTOArrayElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerDTOArrayElementWithDefaults

`func NewConsumerDTOArrayElementWithDefaults() *ConsumerDTOArrayElement`

NewConsumerDTOArrayElementWithDefaults instantiates a new ConsumerDTOArrayElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ConsumerDTOArrayElement) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConsumerDTOArrayElement) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConsumerDTOArrayElement) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ConsumerDTOArrayElement) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ConsumerDTOArrayElement) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ConsumerDTOArrayElement) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ConsumerDTOArrayElement) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ConsumerDTOArrayElement) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *ConsumerDTOArrayElement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsumerDTOArrayElement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsumerDTOArrayElement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsumerDTOArrayElement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ConsumerDTOArrayElement) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ConsumerDTOArrayElement) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ConsumerDTOArrayElement) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ConsumerDTOArrayElement) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *ConsumerDTOArrayElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsumerDTOArrayElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsumerDTOArrayElement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsumerDTOArrayElement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *ConsumerDTOArrayElement) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConsumerDTOArrayElement) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConsumerDTOArrayElement) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConsumerDTOArrayElement) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEntitlementStatus

`func (o *ConsumerDTOArrayElement) GetEntitlementStatus() string`

GetEntitlementStatus returns the EntitlementStatus field if non-nil, zero value otherwise.

### GetEntitlementStatusOk

`func (o *ConsumerDTOArrayElement) GetEntitlementStatusOk() (*string, bool)`

GetEntitlementStatusOk returns a tuple with the EntitlementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementStatus

`func (o *ConsumerDTOArrayElement) SetEntitlementStatus(v string)`

SetEntitlementStatus sets EntitlementStatus field to given value.

### HasEntitlementStatus

`func (o *ConsumerDTOArrayElement) HasEntitlementStatus() bool`

HasEntitlementStatus returns a boolean if a field has been set.

### GetServiceLevel

`func (o *ConsumerDTOArrayElement) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *ConsumerDTOArrayElement) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *ConsumerDTOArrayElement) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *ConsumerDTOArrayElement) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetRole

`func (o *ConsumerDTOArrayElement) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ConsumerDTOArrayElement) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ConsumerDTOArrayElement) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ConsumerDTOArrayElement) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUsage

`func (o *ConsumerDTOArrayElement) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ConsumerDTOArrayElement) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ConsumerDTOArrayElement) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ConsumerDTOArrayElement) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetAddOns

`func (o *ConsumerDTOArrayElement) GetAddOns() []string`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *ConsumerDTOArrayElement) GetAddOnsOk() (*[]string, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *ConsumerDTOArrayElement) SetAddOns(v []string)`

SetAddOns sets AddOns field to given value.

### HasAddOns

`func (o *ConsumerDTOArrayElement) HasAddOns() bool`

HasAddOns returns a boolean if a field has been set.

### GetSystemPurposeStatus

`func (o *ConsumerDTOArrayElement) GetSystemPurposeStatus() string`

GetSystemPurposeStatus returns the SystemPurposeStatus field if non-nil, zero value otherwise.

### GetSystemPurposeStatusOk

`func (o *ConsumerDTOArrayElement) GetSystemPurposeStatusOk() (*string, bool)`

GetSystemPurposeStatusOk returns a tuple with the SystemPurposeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPurposeStatus

`func (o *ConsumerDTOArrayElement) SetSystemPurposeStatus(v string)`

SetSystemPurposeStatus sets SystemPurposeStatus field to given value.

### HasSystemPurposeStatus

`func (o *ConsumerDTOArrayElement) HasSystemPurposeStatus() bool`

HasSystemPurposeStatus returns a boolean if a field has been set.

### GetReleaseVer

`func (o *ConsumerDTOArrayElement) GetReleaseVer() ReleaseVerDTO`

GetReleaseVer returns the ReleaseVer field if non-nil, zero value otherwise.

### GetReleaseVerOk

`func (o *ConsumerDTOArrayElement) GetReleaseVerOk() (*ReleaseVerDTO, bool)`

GetReleaseVerOk returns a tuple with the ReleaseVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVer

`func (o *ConsumerDTOArrayElement) SetReleaseVer(v ReleaseVerDTO)`

SetReleaseVer sets ReleaseVer field to given value.

### HasReleaseVer

`func (o *ConsumerDTOArrayElement) HasReleaseVer() bool`

HasReleaseVer returns a boolean if a field has been set.

### GetOwner

`func (o *ConsumerDTOArrayElement) GetOwner() NestedOwnerDTO`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ConsumerDTOArrayElement) GetOwnerOk() (*NestedOwnerDTO, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ConsumerDTOArrayElement) SetOwner(v NestedOwnerDTO)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ConsumerDTOArrayElement) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *ConsumerDTOArrayElement) GetEntitlementCount() int64`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *ConsumerDTOArrayElement) GetEntitlementCountOk() (*int64, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *ConsumerDTOArrayElement) SetEntitlementCount(v int64)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *ConsumerDTOArrayElement) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetLastCheckin

`func (o *ConsumerDTOArrayElement) GetLastCheckin() time.Time`

GetLastCheckin returns the LastCheckin field if non-nil, zero value otherwise.

### GetLastCheckinOk

`func (o *ConsumerDTOArrayElement) GetLastCheckinOk() (*time.Time, bool)`

GetLastCheckinOk returns a tuple with the LastCheckin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckin

`func (o *ConsumerDTOArrayElement) SetLastCheckin(v time.Time)`

SetLastCheckin sets LastCheckin field to given value.

### HasLastCheckin

`func (o *ConsumerDTOArrayElement) HasLastCheckin() bool`

HasLastCheckin returns a boolean if a field has been set.

### GetInstalledProducts

`func (o *ConsumerDTOArrayElement) GetInstalledProducts() []ConsumerInstalledProductDTO`

GetInstalledProducts returns the InstalledProducts field if non-nil, zero value otherwise.

### GetInstalledProductsOk

`func (o *ConsumerDTOArrayElement) GetInstalledProductsOk() (*[]ConsumerInstalledProductDTO, bool)`

GetInstalledProductsOk returns a tuple with the InstalledProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledProducts

`func (o *ConsumerDTOArrayElement) SetInstalledProducts(v []ConsumerInstalledProductDTO)`

SetInstalledProducts sets InstalledProducts field to given value.

### HasInstalledProducts

`func (o *ConsumerDTOArrayElement) HasInstalledProducts() bool`

HasInstalledProducts returns a boolean if a field has been set.

### GetCanActivate

`func (o *ConsumerDTOArrayElement) GetCanActivate() bool`

GetCanActivate returns the CanActivate field if non-nil, zero value otherwise.

### GetCanActivateOk

`func (o *ConsumerDTOArrayElement) GetCanActivateOk() (*bool, bool)`

GetCanActivateOk returns a tuple with the CanActivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanActivate

`func (o *ConsumerDTOArrayElement) SetCanActivate(v bool)`

SetCanActivate sets CanActivate field to given value.

### HasCanActivate

`func (o *ConsumerDTOArrayElement) HasCanActivate() bool`

HasCanActivate returns a boolean if a field has been set.

### GetCapabilities

`func (o *ConsumerDTOArrayElement) GetCapabilities() []CapabilityDTO`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ConsumerDTOArrayElement) GetCapabilitiesOk() (*[]CapabilityDTO, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ConsumerDTOArrayElement) SetCapabilities(v []CapabilityDTO)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ConsumerDTOArrayElement) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetHypervisorId

`func (o *ConsumerDTOArrayElement) GetHypervisorId() HypervisorIdDTO`

GetHypervisorId returns the HypervisorId field if non-nil, zero value otherwise.

### GetHypervisorIdOk

`func (o *ConsumerDTOArrayElement) GetHypervisorIdOk() (*HypervisorIdDTO, bool)`

GetHypervisorIdOk returns a tuple with the HypervisorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorId

`func (o *ConsumerDTOArrayElement) SetHypervisorId(v HypervisorIdDTO)`

SetHypervisorId sets HypervisorId field to given value.

### HasHypervisorId

`func (o *ConsumerDTOArrayElement) HasHypervisorId() bool`

HasHypervisorId returns a boolean if a field has been set.

### GetContentTags

`func (o *ConsumerDTOArrayElement) GetContentTags() []string`

GetContentTags returns the ContentTags field if non-nil, zero value otherwise.

### GetContentTagsOk

`func (o *ConsumerDTOArrayElement) GetContentTagsOk() (*[]string, bool)`

GetContentTagsOk returns a tuple with the ContentTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTags

`func (o *ConsumerDTOArrayElement) SetContentTags(v []string)`

SetContentTags sets ContentTags field to given value.

### HasContentTags

`func (o *ConsumerDTOArrayElement) HasContentTags() bool`

HasContentTags returns a boolean if a field has been set.

### GetAutoheal

`func (o *ConsumerDTOArrayElement) GetAutoheal() bool`

GetAutoheal returns the Autoheal field if non-nil, zero value otherwise.

### GetAutohealOk

`func (o *ConsumerDTOArrayElement) GetAutohealOk() (*bool, bool)`

GetAutohealOk returns a tuple with the Autoheal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoheal

`func (o *ConsumerDTOArrayElement) SetAutoheal(v bool)`

SetAutoheal sets Autoheal field to given value.

### HasAutoheal

`func (o *ConsumerDTOArrayElement) HasAutoheal() bool`

HasAutoheal returns a boolean if a field has been set.

### GetAnnotations

`func (o *ConsumerDTOArrayElement) GetAnnotations() string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ConsumerDTOArrayElement) GetAnnotationsOk() (*string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ConsumerDTOArrayElement) SetAnnotations(v string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ConsumerDTOArrayElement) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetContentAccessMode

`func (o *ConsumerDTOArrayElement) GetContentAccessMode() string`

GetContentAccessMode returns the ContentAccessMode field if non-nil, zero value otherwise.

### GetContentAccessModeOk

`func (o *ConsumerDTOArrayElement) GetContentAccessModeOk() (*string, bool)`

GetContentAccessModeOk returns a tuple with the ContentAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAccessMode

`func (o *ConsumerDTOArrayElement) SetContentAccessMode(v string)`

SetContentAccessMode sets ContentAccessMode field to given value.

### HasContentAccessMode

`func (o *ConsumerDTOArrayElement) HasContentAccessMode() bool`

HasContentAccessMode returns a boolean if a field has been set.

### GetType

`func (o *ConsumerDTOArrayElement) GetType() ConsumerTypeDTO`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsumerDTOArrayElement) GetTypeOk() (*ConsumerTypeDTO, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsumerDTOArrayElement) SetType(v ConsumerTypeDTO)`

SetType sets Type field to given value.

### HasType

`func (o *ConsumerDTOArrayElement) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGuestIds

`func (o *ConsumerDTOArrayElement) GetGuestIds() []GuestIdDTOArrayElement`

GetGuestIds returns the GuestIds field if non-nil, zero value otherwise.

### GetGuestIdsOk

`func (o *ConsumerDTOArrayElement) GetGuestIdsOk() (*[]GuestIdDTOArrayElement, bool)`

GetGuestIdsOk returns a tuple with the GuestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestIds

`func (o *ConsumerDTOArrayElement) SetGuestIds(v []GuestIdDTOArrayElement)`

SetGuestIds sets GuestIds field to given value.

### HasGuestIds

`func (o *ConsumerDTOArrayElement) HasGuestIds() bool`

HasGuestIds returns a boolean if a field has been set.

### GetHref

`func (o *ConsumerDTOArrayElement) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConsumerDTOArrayElement) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConsumerDTOArrayElement) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConsumerDTOArrayElement) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetServiceType

`func (o *ConsumerDTOArrayElement) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ConsumerDTOArrayElement) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ConsumerDTOArrayElement) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ConsumerDTOArrayElement) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


