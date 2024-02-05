# OwnerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**ContentPrefix** | Pointer to **string** |  | [optional] 
**DefaultServiceLevel** | Pointer to **string** |  | [optional] 
**LogLevel** | Pointer to **string** |  | [optional] 
**ContentAccessMode** | Pointer to **string** |  | [optional] 
**ContentAccessModeList** | Pointer to **string** |  | [optional] 
**AutobindHypervisorDisabled** | Pointer to **bool** |  | [optional] 
**AutobindDisabled** | Pointer to **bool** |  | [optional] 
**LastRefreshed** | Pointer to **time.Time** |  | [optional] 
**ParentOwner** | Pointer to [**NestedOwnerDTO**](NestedOwnerDTO.md) |  | [optional] 
**UpstreamConsumer** | Pointer to [**UpstreamConsumerDTO**](UpstreamConsumerDTO.md) |  | [optional] 
**Anonymous** | Pointer to **bool** |  | [optional] 
**Claimed** | Pointer to **bool** |  | [optional] 
**ClaimantOwner** | Pointer to **string** |  | [optional] 

## Methods

### NewOwnerDTO

`func NewOwnerDTO() *OwnerDTO`

NewOwnerDTO instantiates a new OwnerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerDTOWithDefaults

`func NewOwnerDTOWithDefaults() *OwnerDTO`

NewOwnerDTOWithDefaults instantiates a new OwnerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *OwnerDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OwnerDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OwnerDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OwnerDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *OwnerDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *OwnerDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *OwnerDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *OwnerDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *OwnerDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OwnerDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OwnerDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OwnerDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *OwnerDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OwnerDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OwnerDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OwnerDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetKey

`func (o *OwnerDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OwnerDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OwnerDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OwnerDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetContentPrefix

`func (o *OwnerDTO) GetContentPrefix() string`

GetContentPrefix returns the ContentPrefix field if non-nil, zero value otherwise.

### GetContentPrefixOk

`func (o *OwnerDTO) GetContentPrefixOk() (*string, bool)`

GetContentPrefixOk returns a tuple with the ContentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPrefix

`func (o *OwnerDTO) SetContentPrefix(v string)`

SetContentPrefix sets ContentPrefix field to given value.

### HasContentPrefix

`func (o *OwnerDTO) HasContentPrefix() bool`

HasContentPrefix returns a boolean if a field has been set.

### GetDefaultServiceLevel

`func (o *OwnerDTO) GetDefaultServiceLevel() string`

GetDefaultServiceLevel returns the DefaultServiceLevel field if non-nil, zero value otherwise.

### GetDefaultServiceLevelOk

`func (o *OwnerDTO) GetDefaultServiceLevelOk() (*string, bool)`

GetDefaultServiceLevelOk returns a tuple with the DefaultServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceLevel

`func (o *OwnerDTO) SetDefaultServiceLevel(v string)`

SetDefaultServiceLevel sets DefaultServiceLevel field to given value.

### HasDefaultServiceLevel

`func (o *OwnerDTO) HasDefaultServiceLevel() bool`

HasDefaultServiceLevel returns a boolean if a field has been set.

### GetLogLevel

`func (o *OwnerDTO) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *OwnerDTO) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *OwnerDTO) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *OwnerDTO) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetContentAccessMode

`func (o *OwnerDTO) GetContentAccessMode() string`

GetContentAccessMode returns the ContentAccessMode field if non-nil, zero value otherwise.

### GetContentAccessModeOk

`func (o *OwnerDTO) GetContentAccessModeOk() (*string, bool)`

GetContentAccessModeOk returns a tuple with the ContentAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAccessMode

`func (o *OwnerDTO) SetContentAccessMode(v string)`

SetContentAccessMode sets ContentAccessMode field to given value.

### HasContentAccessMode

`func (o *OwnerDTO) HasContentAccessMode() bool`

HasContentAccessMode returns a boolean if a field has been set.

### GetContentAccessModeList

`func (o *OwnerDTO) GetContentAccessModeList() string`

GetContentAccessModeList returns the ContentAccessModeList field if non-nil, zero value otherwise.

### GetContentAccessModeListOk

`func (o *OwnerDTO) GetContentAccessModeListOk() (*string, bool)`

GetContentAccessModeListOk returns a tuple with the ContentAccessModeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAccessModeList

`func (o *OwnerDTO) SetContentAccessModeList(v string)`

SetContentAccessModeList sets ContentAccessModeList field to given value.

### HasContentAccessModeList

`func (o *OwnerDTO) HasContentAccessModeList() bool`

HasContentAccessModeList returns a boolean if a field has been set.

### GetAutobindHypervisorDisabled

`func (o *OwnerDTO) GetAutobindHypervisorDisabled() bool`

GetAutobindHypervisorDisabled returns the AutobindHypervisorDisabled field if non-nil, zero value otherwise.

### GetAutobindHypervisorDisabledOk

`func (o *OwnerDTO) GetAutobindHypervisorDisabledOk() (*bool, bool)`

GetAutobindHypervisorDisabledOk returns a tuple with the AutobindHypervisorDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutobindHypervisorDisabled

`func (o *OwnerDTO) SetAutobindHypervisorDisabled(v bool)`

SetAutobindHypervisorDisabled sets AutobindHypervisorDisabled field to given value.

### HasAutobindHypervisorDisabled

`func (o *OwnerDTO) HasAutobindHypervisorDisabled() bool`

HasAutobindHypervisorDisabled returns a boolean if a field has been set.

### GetAutobindDisabled

`func (o *OwnerDTO) GetAutobindDisabled() bool`

GetAutobindDisabled returns the AutobindDisabled field if non-nil, zero value otherwise.

### GetAutobindDisabledOk

`func (o *OwnerDTO) GetAutobindDisabledOk() (*bool, bool)`

GetAutobindDisabledOk returns a tuple with the AutobindDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutobindDisabled

`func (o *OwnerDTO) SetAutobindDisabled(v bool)`

SetAutobindDisabled sets AutobindDisabled field to given value.

### HasAutobindDisabled

`func (o *OwnerDTO) HasAutobindDisabled() bool`

HasAutobindDisabled returns a boolean if a field has been set.

### GetLastRefreshed

`func (o *OwnerDTO) GetLastRefreshed() time.Time`

GetLastRefreshed returns the LastRefreshed field if non-nil, zero value otherwise.

### GetLastRefreshedOk

`func (o *OwnerDTO) GetLastRefreshedOk() (*time.Time, bool)`

GetLastRefreshedOk returns a tuple with the LastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshed

`func (o *OwnerDTO) SetLastRefreshed(v time.Time)`

SetLastRefreshed sets LastRefreshed field to given value.

### HasLastRefreshed

`func (o *OwnerDTO) HasLastRefreshed() bool`

HasLastRefreshed returns a boolean if a field has been set.

### GetParentOwner

`func (o *OwnerDTO) GetParentOwner() NestedOwnerDTO`

GetParentOwner returns the ParentOwner field if non-nil, zero value otherwise.

### GetParentOwnerOk

`func (o *OwnerDTO) GetParentOwnerOk() (*NestedOwnerDTO, bool)`

GetParentOwnerOk returns a tuple with the ParentOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOwner

`func (o *OwnerDTO) SetParentOwner(v NestedOwnerDTO)`

SetParentOwner sets ParentOwner field to given value.

### HasParentOwner

`func (o *OwnerDTO) HasParentOwner() bool`

HasParentOwner returns a boolean if a field has been set.

### GetUpstreamConsumer

`func (o *OwnerDTO) GetUpstreamConsumer() UpstreamConsumerDTO`

GetUpstreamConsumer returns the UpstreamConsumer field if non-nil, zero value otherwise.

### GetUpstreamConsumerOk

`func (o *OwnerDTO) GetUpstreamConsumerOk() (*UpstreamConsumerDTO, bool)`

GetUpstreamConsumerOk returns a tuple with the UpstreamConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamConsumer

`func (o *OwnerDTO) SetUpstreamConsumer(v UpstreamConsumerDTO)`

SetUpstreamConsumer sets UpstreamConsumer field to given value.

### HasUpstreamConsumer

`func (o *OwnerDTO) HasUpstreamConsumer() bool`

HasUpstreamConsumer returns a boolean if a field has been set.

### GetAnonymous

`func (o *OwnerDTO) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *OwnerDTO) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *OwnerDTO) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *OwnerDTO) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetClaimed

`func (o *OwnerDTO) GetClaimed() bool`

GetClaimed returns the Claimed field if non-nil, zero value otherwise.

### GetClaimedOk

`func (o *OwnerDTO) GetClaimedOk() (*bool, bool)`

GetClaimedOk returns a tuple with the Claimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimed

`func (o *OwnerDTO) SetClaimed(v bool)`

SetClaimed sets Claimed field to given value.

### HasClaimed

`func (o *OwnerDTO) HasClaimed() bool`

HasClaimed returns a boolean if a field has been set.

### GetClaimantOwner

`func (o *OwnerDTO) GetClaimantOwner() string`

GetClaimantOwner returns the ClaimantOwner field if non-nil, zero value otherwise.

### GetClaimantOwnerOk

`func (o *OwnerDTO) GetClaimantOwnerOk() (*string, bool)`

GetClaimantOwnerOk returns a tuple with the ClaimantOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimantOwner

`func (o *OwnerDTO) SetClaimantOwner(v string)`

SetClaimantOwner sets ClaimantOwner field to given value.

### HasClaimantOwner

`func (o *OwnerDTO) HasClaimantOwner() bool`

HasClaimantOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


