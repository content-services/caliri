# ConsumerDTO

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
**RegistrationAuthenticationMethod** | Pointer to **string** |  | [optional] 
**ReleaseVer** | Pointer to [**ReleaseVerDTO**](ReleaseVerDTO.md) |  | [optional] 
**Owner** | Pointer to [**NestedOwnerDTO**](NestedOwnerDTO.md) |  | [optional] 
**Environment** | Pointer to [**EnvironmentDTO**](EnvironmentDTO.md) |  | [optional] 
**EntitlementCount** | Pointer to **int64** |  | [optional] 
**Facts** | Pointer to **map[string]string** |  | [optional] 
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
**IdCert** | Pointer to [**CertificateDTO**](CertificateDTO.md) |  | [optional] 
**GuestIds** | Pointer to [**[]GuestIdDTO**](GuestIdDTO.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**ActivationKeys** | Pointer to [**[]ConsumerActivationKeyDTO**](ConsumerActivationKeyDTO.md) |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**Environments** | Pointer to [**[]EnvironmentDTO**](EnvironmentDTO.md) |  | [optional] 

## Methods

### NewConsumerDTO

`func NewConsumerDTO() *ConsumerDTO`

NewConsumerDTO instantiates a new ConsumerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerDTOWithDefaults

`func NewConsumerDTOWithDefaults() *ConsumerDTO`

NewConsumerDTOWithDefaults instantiates a new ConsumerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ConsumerDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConsumerDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConsumerDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ConsumerDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ConsumerDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ConsumerDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ConsumerDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ConsumerDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *ConsumerDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsumerDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsumerDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsumerDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ConsumerDTO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ConsumerDTO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ConsumerDTO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ConsumerDTO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *ConsumerDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsumerDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsumerDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsumerDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *ConsumerDTO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConsumerDTO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConsumerDTO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConsumerDTO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEntitlementStatus

`func (o *ConsumerDTO) GetEntitlementStatus() string`

GetEntitlementStatus returns the EntitlementStatus field if non-nil, zero value otherwise.

### GetEntitlementStatusOk

`func (o *ConsumerDTO) GetEntitlementStatusOk() (*string, bool)`

GetEntitlementStatusOk returns a tuple with the EntitlementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementStatus

`func (o *ConsumerDTO) SetEntitlementStatus(v string)`

SetEntitlementStatus sets EntitlementStatus field to given value.

### HasEntitlementStatus

`func (o *ConsumerDTO) HasEntitlementStatus() bool`

HasEntitlementStatus returns a boolean if a field has been set.

### GetServiceLevel

`func (o *ConsumerDTO) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *ConsumerDTO) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *ConsumerDTO) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *ConsumerDTO) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetRole

`func (o *ConsumerDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ConsumerDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ConsumerDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ConsumerDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUsage

`func (o *ConsumerDTO) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ConsumerDTO) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ConsumerDTO) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ConsumerDTO) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetAddOns

`func (o *ConsumerDTO) GetAddOns() []string`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *ConsumerDTO) GetAddOnsOk() (*[]string, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *ConsumerDTO) SetAddOns(v []string)`

SetAddOns sets AddOns field to given value.

### HasAddOns

`func (o *ConsumerDTO) HasAddOns() bool`

HasAddOns returns a boolean if a field has been set.

### GetSystemPurposeStatus

`func (o *ConsumerDTO) GetSystemPurposeStatus() string`

GetSystemPurposeStatus returns the SystemPurposeStatus field if non-nil, zero value otherwise.

### GetSystemPurposeStatusOk

`func (o *ConsumerDTO) GetSystemPurposeStatusOk() (*string, bool)`

GetSystemPurposeStatusOk returns a tuple with the SystemPurposeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPurposeStatus

`func (o *ConsumerDTO) SetSystemPurposeStatus(v string)`

SetSystemPurposeStatus sets SystemPurposeStatus field to given value.

### HasSystemPurposeStatus

`func (o *ConsumerDTO) HasSystemPurposeStatus() bool`

HasSystemPurposeStatus returns a boolean if a field has been set.

### GetRegistrationAuthenticationMethod

`func (o *ConsumerDTO) GetRegistrationAuthenticationMethod() string`

GetRegistrationAuthenticationMethod returns the RegistrationAuthenticationMethod field if non-nil, zero value otherwise.

### GetRegistrationAuthenticationMethodOk

`func (o *ConsumerDTO) GetRegistrationAuthenticationMethodOk() (*string, bool)`

GetRegistrationAuthenticationMethodOk returns a tuple with the RegistrationAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationAuthenticationMethod

`func (o *ConsumerDTO) SetRegistrationAuthenticationMethod(v string)`

SetRegistrationAuthenticationMethod sets RegistrationAuthenticationMethod field to given value.

### HasRegistrationAuthenticationMethod

`func (o *ConsumerDTO) HasRegistrationAuthenticationMethod() bool`

HasRegistrationAuthenticationMethod returns a boolean if a field has been set.

### GetReleaseVer

`func (o *ConsumerDTO) GetReleaseVer() ReleaseVerDTO`

GetReleaseVer returns the ReleaseVer field if non-nil, zero value otherwise.

### GetReleaseVerOk

`func (o *ConsumerDTO) GetReleaseVerOk() (*ReleaseVerDTO, bool)`

GetReleaseVerOk returns a tuple with the ReleaseVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVer

`func (o *ConsumerDTO) SetReleaseVer(v ReleaseVerDTO)`

SetReleaseVer sets ReleaseVer field to given value.

### HasReleaseVer

`func (o *ConsumerDTO) HasReleaseVer() bool`

HasReleaseVer returns a boolean if a field has been set.

### GetOwner

`func (o *ConsumerDTO) GetOwner() NestedOwnerDTO`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ConsumerDTO) GetOwnerOk() (*NestedOwnerDTO, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ConsumerDTO) SetOwner(v NestedOwnerDTO)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ConsumerDTO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetEnvironment

`func (o *ConsumerDTO) GetEnvironment() EnvironmentDTO`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConsumerDTO) GetEnvironmentOk() (*EnvironmentDTO, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConsumerDTO) SetEnvironment(v EnvironmentDTO)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConsumerDTO) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *ConsumerDTO) GetEntitlementCount() int64`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *ConsumerDTO) GetEntitlementCountOk() (*int64, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *ConsumerDTO) SetEntitlementCount(v int64)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *ConsumerDTO) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetFacts

`func (o *ConsumerDTO) GetFacts() map[string]string`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *ConsumerDTO) GetFactsOk() (*map[string]string, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *ConsumerDTO) SetFacts(v map[string]string)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *ConsumerDTO) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetLastCheckin

`func (o *ConsumerDTO) GetLastCheckin() time.Time`

GetLastCheckin returns the LastCheckin field if non-nil, zero value otherwise.

### GetLastCheckinOk

`func (o *ConsumerDTO) GetLastCheckinOk() (*time.Time, bool)`

GetLastCheckinOk returns a tuple with the LastCheckin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckin

`func (o *ConsumerDTO) SetLastCheckin(v time.Time)`

SetLastCheckin sets LastCheckin field to given value.

### HasLastCheckin

`func (o *ConsumerDTO) HasLastCheckin() bool`

HasLastCheckin returns a boolean if a field has been set.

### GetInstalledProducts

`func (o *ConsumerDTO) GetInstalledProducts() []ConsumerInstalledProductDTO`

GetInstalledProducts returns the InstalledProducts field if non-nil, zero value otherwise.

### GetInstalledProductsOk

`func (o *ConsumerDTO) GetInstalledProductsOk() (*[]ConsumerInstalledProductDTO, bool)`

GetInstalledProductsOk returns a tuple with the InstalledProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledProducts

`func (o *ConsumerDTO) SetInstalledProducts(v []ConsumerInstalledProductDTO)`

SetInstalledProducts sets InstalledProducts field to given value.

### HasInstalledProducts

`func (o *ConsumerDTO) HasInstalledProducts() bool`

HasInstalledProducts returns a boolean if a field has been set.

### GetCanActivate

`func (o *ConsumerDTO) GetCanActivate() bool`

GetCanActivate returns the CanActivate field if non-nil, zero value otherwise.

### GetCanActivateOk

`func (o *ConsumerDTO) GetCanActivateOk() (*bool, bool)`

GetCanActivateOk returns a tuple with the CanActivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanActivate

`func (o *ConsumerDTO) SetCanActivate(v bool)`

SetCanActivate sets CanActivate field to given value.

### HasCanActivate

`func (o *ConsumerDTO) HasCanActivate() bool`

HasCanActivate returns a boolean if a field has been set.

### GetCapabilities

`func (o *ConsumerDTO) GetCapabilities() []CapabilityDTO`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ConsumerDTO) GetCapabilitiesOk() (*[]CapabilityDTO, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ConsumerDTO) SetCapabilities(v []CapabilityDTO)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ConsumerDTO) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetHypervisorId

`func (o *ConsumerDTO) GetHypervisorId() HypervisorIdDTO`

GetHypervisorId returns the HypervisorId field if non-nil, zero value otherwise.

### GetHypervisorIdOk

`func (o *ConsumerDTO) GetHypervisorIdOk() (*HypervisorIdDTO, bool)`

GetHypervisorIdOk returns a tuple with the HypervisorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorId

`func (o *ConsumerDTO) SetHypervisorId(v HypervisorIdDTO)`

SetHypervisorId sets HypervisorId field to given value.

### HasHypervisorId

`func (o *ConsumerDTO) HasHypervisorId() bool`

HasHypervisorId returns a boolean if a field has been set.

### GetContentTags

`func (o *ConsumerDTO) GetContentTags() []string`

GetContentTags returns the ContentTags field if non-nil, zero value otherwise.

### GetContentTagsOk

`func (o *ConsumerDTO) GetContentTagsOk() (*[]string, bool)`

GetContentTagsOk returns a tuple with the ContentTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTags

`func (o *ConsumerDTO) SetContentTags(v []string)`

SetContentTags sets ContentTags field to given value.

### HasContentTags

`func (o *ConsumerDTO) HasContentTags() bool`

HasContentTags returns a boolean if a field has been set.

### GetAutoheal

`func (o *ConsumerDTO) GetAutoheal() bool`

GetAutoheal returns the Autoheal field if non-nil, zero value otherwise.

### GetAutohealOk

`func (o *ConsumerDTO) GetAutohealOk() (*bool, bool)`

GetAutohealOk returns a tuple with the Autoheal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoheal

`func (o *ConsumerDTO) SetAutoheal(v bool)`

SetAutoheal sets Autoheal field to given value.

### HasAutoheal

`func (o *ConsumerDTO) HasAutoheal() bool`

HasAutoheal returns a boolean if a field has been set.

### GetAnnotations

`func (o *ConsumerDTO) GetAnnotations() string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ConsumerDTO) GetAnnotationsOk() (*string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ConsumerDTO) SetAnnotations(v string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ConsumerDTO) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetContentAccessMode

`func (o *ConsumerDTO) GetContentAccessMode() string`

GetContentAccessMode returns the ContentAccessMode field if non-nil, zero value otherwise.

### GetContentAccessModeOk

`func (o *ConsumerDTO) GetContentAccessModeOk() (*string, bool)`

GetContentAccessModeOk returns a tuple with the ContentAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAccessMode

`func (o *ConsumerDTO) SetContentAccessMode(v string)`

SetContentAccessMode sets ContentAccessMode field to given value.

### HasContentAccessMode

`func (o *ConsumerDTO) HasContentAccessMode() bool`

HasContentAccessMode returns a boolean if a field has been set.

### GetType

`func (o *ConsumerDTO) GetType() ConsumerTypeDTO`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsumerDTO) GetTypeOk() (*ConsumerTypeDTO, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsumerDTO) SetType(v ConsumerTypeDTO)`

SetType sets Type field to given value.

### HasType

`func (o *ConsumerDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIdCert

`func (o *ConsumerDTO) GetIdCert() CertificateDTO`

GetIdCert returns the IdCert field if non-nil, zero value otherwise.

### GetIdCertOk

`func (o *ConsumerDTO) GetIdCertOk() (*CertificateDTO, bool)`

GetIdCertOk returns a tuple with the IdCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCert

`func (o *ConsumerDTO) SetIdCert(v CertificateDTO)`

SetIdCert sets IdCert field to given value.

### HasIdCert

`func (o *ConsumerDTO) HasIdCert() bool`

HasIdCert returns a boolean if a field has been set.

### GetGuestIds

`func (o *ConsumerDTO) GetGuestIds() []GuestIdDTO`

GetGuestIds returns the GuestIds field if non-nil, zero value otherwise.

### GetGuestIdsOk

`func (o *ConsumerDTO) GetGuestIdsOk() (*[]GuestIdDTO, bool)`

GetGuestIdsOk returns a tuple with the GuestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestIds

`func (o *ConsumerDTO) SetGuestIds(v []GuestIdDTO)`

SetGuestIds sets GuestIds field to given value.

### HasGuestIds

`func (o *ConsumerDTO) HasGuestIds() bool`

HasGuestIds returns a boolean if a field has been set.

### GetHref

`func (o *ConsumerDTO) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConsumerDTO) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConsumerDTO) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConsumerDTO) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetActivationKeys

`func (o *ConsumerDTO) GetActivationKeys() []ConsumerActivationKeyDTO`

GetActivationKeys returns the ActivationKeys field if non-nil, zero value otherwise.

### GetActivationKeysOk

`func (o *ConsumerDTO) GetActivationKeysOk() (*[]ConsumerActivationKeyDTO, bool)`

GetActivationKeysOk returns a tuple with the ActivationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationKeys

`func (o *ConsumerDTO) SetActivationKeys(v []ConsumerActivationKeyDTO)`

SetActivationKeys sets ActivationKeys field to given value.

### HasActivationKeys

`func (o *ConsumerDTO) HasActivationKeys() bool`

HasActivationKeys returns a boolean if a field has been set.

### GetServiceType

`func (o *ConsumerDTO) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ConsumerDTO) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ConsumerDTO) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ConsumerDTO) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetEnvironments

`func (o *ConsumerDTO) GetEnvironments() []EnvironmentDTO`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *ConsumerDTO) GetEnvironmentsOk() (*[]EnvironmentDTO, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *ConsumerDTO) SetEnvironments(v []EnvironmentDTO)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *ConsumerDTO) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


