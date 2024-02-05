# StatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |  | [optional] 
**ModeReason** | Pointer to **string** |  | [optional] 
**ModeChangeTime** | Pointer to **time.Time** |  | [optional] 
**Result** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**Standalone** | Pointer to **bool** |  | [optional] 
**TimeUTC** | Pointer to **time.Time** |  | [optional] 
**RulesSource** | Pointer to **string** |  | [optional] 
**RulesVersion** | Pointer to **string** |  | [optional] 
**ManagerCapabilities** | Pointer to **[]string** |  | [optional] 
**KeycloakRealm** | Pointer to **string** |  | [optional] 
**KeycloakAuthUrl** | Pointer to **string** |  | [optional] 
**KeycloakResource** | Pointer to **string** |  | [optional] 
**DeviceAuthRealm** | Pointer to **string** |  | [optional] 
**DeviceAuthUrl** | Pointer to **string** |  | [optional] 
**DeviceAuthClientId** | Pointer to **string** |  | [optional] 
**DeviceAuthScope** | Pointer to **string** |  | [optional] 

## Methods

### NewStatusDTO

`func NewStatusDTO() *StatusDTO`

NewStatusDTO instantiates a new StatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusDTOWithDefaults

`func NewStatusDTOWithDefaults() *StatusDTO`

NewStatusDTOWithDefaults instantiates a new StatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *StatusDTO) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *StatusDTO) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *StatusDTO) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *StatusDTO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetModeReason

`func (o *StatusDTO) GetModeReason() string`

GetModeReason returns the ModeReason field if non-nil, zero value otherwise.

### GetModeReasonOk

`func (o *StatusDTO) GetModeReasonOk() (*string, bool)`

GetModeReasonOk returns a tuple with the ModeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeReason

`func (o *StatusDTO) SetModeReason(v string)`

SetModeReason sets ModeReason field to given value.

### HasModeReason

`func (o *StatusDTO) HasModeReason() bool`

HasModeReason returns a boolean if a field has been set.

### GetModeChangeTime

`func (o *StatusDTO) GetModeChangeTime() time.Time`

GetModeChangeTime returns the ModeChangeTime field if non-nil, zero value otherwise.

### GetModeChangeTimeOk

`func (o *StatusDTO) GetModeChangeTimeOk() (*time.Time, bool)`

GetModeChangeTimeOk returns a tuple with the ModeChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeChangeTime

`func (o *StatusDTO) SetModeChangeTime(v time.Time)`

SetModeChangeTime sets ModeChangeTime field to given value.

### HasModeChangeTime

`func (o *StatusDTO) HasModeChangeTime() bool`

HasModeChangeTime returns a boolean if a field has been set.

### GetResult

`func (o *StatusDTO) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StatusDTO) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StatusDTO) SetResult(v bool)`

SetResult sets Result field to given value.

### HasResult

`func (o *StatusDTO) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetVersion

`func (o *StatusDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StatusDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StatusDTO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StatusDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRelease

`func (o *StatusDTO) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *StatusDTO) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *StatusDTO) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *StatusDTO) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetStandalone

`func (o *StatusDTO) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *StatusDTO) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *StatusDTO) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *StatusDTO) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetTimeUTC

`func (o *StatusDTO) GetTimeUTC() time.Time`

GetTimeUTC returns the TimeUTC field if non-nil, zero value otherwise.

### GetTimeUTCOk

`func (o *StatusDTO) GetTimeUTCOk() (*time.Time, bool)`

GetTimeUTCOk returns a tuple with the TimeUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUTC

`func (o *StatusDTO) SetTimeUTC(v time.Time)`

SetTimeUTC sets TimeUTC field to given value.

### HasTimeUTC

`func (o *StatusDTO) HasTimeUTC() bool`

HasTimeUTC returns a boolean if a field has been set.

### GetRulesSource

`func (o *StatusDTO) GetRulesSource() string`

GetRulesSource returns the RulesSource field if non-nil, zero value otherwise.

### GetRulesSourceOk

`func (o *StatusDTO) GetRulesSourceOk() (*string, bool)`

GetRulesSourceOk returns a tuple with the RulesSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesSource

`func (o *StatusDTO) SetRulesSource(v string)`

SetRulesSource sets RulesSource field to given value.

### HasRulesSource

`func (o *StatusDTO) HasRulesSource() bool`

HasRulesSource returns a boolean if a field has been set.

### GetRulesVersion

`func (o *StatusDTO) GetRulesVersion() string`

GetRulesVersion returns the RulesVersion field if non-nil, zero value otherwise.

### GetRulesVersionOk

`func (o *StatusDTO) GetRulesVersionOk() (*string, bool)`

GetRulesVersionOk returns a tuple with the RulesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesVersion

`func (o *StatusDTO) SetRulesVersion(v string)`

SetRulesVersion sets RulesVersion field to given value.

### HasRulesVersion

`func (o *StatusDTO) HasRulesVersion() bool`

HasRulesVersion returns a boolean if a field has been set.

### GetManagerCapabilities

`func (o *StatusDTO) GetManagerCapabilities() []string`

GetManagerCapabilities returns the ManagerCapabilities field if non-nil, zero value otherwise.

### GetManagerCapabilitiesOk

`func (o *StatusDTO) GetManagerCapabilitiesOk() (*[]string, bool)`

GetManagerCapabilitiesOk returns a tuple with the ManagerCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerCapabilities

`func (o *StatusDTO) SetManagerCapabilities(v []string)`

SetManagerCapabilities sets ManagerCapabilities field to given value.

### HasManagerCapabilities

`func (o *StatusDTO) HasManagerCapabilities() bool`

HasManagerCapabilities returns a boolean if a field has been set.

### GetKeycloakRealm

`func (o *StatusDTO) GetKeycloakRealm() string`

GetKeycloakRealm returns the KeycloakRealm field if non-nil, zero value otherwise.

### GetKeycloakRealmOk

`func (o *StatusDTO) GetKeycloakRealmOk() (*string, bool)`

GetKeycloakRealmOk returns a tuple with the KeycloakRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeycloakRealm

`func (o *StatusDTO) SetKeycloakRealm(v string)`

SetKeycloakRealm sets KeycloakRealm field to given value.

### HasKeycloakRealm

`func (o *StatusDTO) HasKeycloakRealm() bool`

HasKeycloakRealm returns a boolean if a field has been set.

### GetKeycloakAuthUrl

`func (o *StatusDTO) GetKeycloakAuthUrl() string`

GetKeycloakAuthUrl returns the KeycloakAuthUrl field if non-nil, zero value otherwise.

### GetKeycloakAuthUrlOk

`func (o *StatusDTO) GetKeycloakAuthUrlOk() (*string, bool)`

GetKeycloakAuthUrlOk returns a tuple with the KeycloakAuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeycloakAuthUrl

`func (o *StatusDTO) SetKeycloakAuthUrl(v string)`

SetKeycloakAuthUrl sets KeycloakAuthUrl field to given value.

### HasKeycloakAuthUrl

`func (o *StatusDTO) HasKeycloakAuthUrl() bool`

HasKeycloakAuthUrl returns a boolean if a field has been set.

### GetKeycloakResource

`func (o *StatusDTO) GetKeycloakResource() string`

GetKeycloakResource returns the KeycloakResource field if non-nil, zero value otherwise.

### GetKeycloakResourceOk

`func (o *StatusDTO) GetKeycloakResourceOk() (*string, bool)`

GetKeycloakResourceOk returns a tuple with the KeycloakResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeycloakResource

`func (o *StatusDTO) SetKeycloakResource(v string)`

SetKeycloakResource sets KeycloakResource field to given value.

### HasKeycloakResource

`func (o *StatusDTO) HasKeycloakResource() bool`

HasKeycloakResource returns a boolean if a field has been set.

### GetDeviceAuthRealm

`func (o *StatusDTO) GetDeviceAuthRealm() string`

GetDeviceAuthRealm returns the DeviceAuthRealm field if non-nil, zero value otherwise.

### GetDeviceAuthRealmOk

`func (o *StatusDTO) GetDeviceAuthRealmOk() (*string, bool)`

GetDeviceAuthRealmOk returns a tuple with the DeviceAuthRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthRealm

`func (o *StatusDTO) SetDeviceAuthRealm(v string)`

SetDeviceAuthRealm sets DeviceAuthRealm field to given value.

### HasDeviceAuthRealm

`func (o *StatusDTO) HasDeviceAuthRealm() bool`

HasDeviceAuthRealm returns a boolean if a field has been set.

### GetDeviceAuthUrl

`func (o *StatusDTO) GetDeviceAuthUrl() string`

GetDeviceAuthUrl returns the DeviceAuthUrl field if non-nil, zero value otherwise.

### GetDeviceAuthUrlOk

`func (o *StatusDTO) GetDeviceAuthUrlOk() (*string, bool)`

GetDeviceAuthUrlOk returns a tuple with the DeviceAuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthUrl

`func (o *StatusDTO) SetDeviceAuthUrl(v string)`

SetDeviceAuthUrl sets DeviceAuthUrl field to given value.

### HasDeviceAuthUrl

`func (o *StatusDTO) HasDeviceAuthUrl() bool`

HasDeviceAuthUrl returns a boolean if a field has been set.

### GetDeviceAuthClientId

`func (o *StatusDTO) GetDeviceAuthClientId() string`

GetDeviceAuthClientId returns the DeviceAuthClientId field if non-nil, zero value otherwise.

### GetDeviceAuthClientIdOk

`func (o *StatusDTO) GetDeviceAuthClientIdOk() (*string, bool)`

GetDeviceAuthClientIdOk returns a tuple with the DeviceAuthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthClientId

`func (o *StatusDTO) SetDeviceAuthClientId(v string)`

SetDeviceAuthClientId sets DeviceAuthClientId field to given value.

### HasDeviceAuthClientId

`func (o *StatusDTO) HasDeviceAuthClientId() bool`

HasDeviceAuthClientId returns a boolean if a field has been set.

### GetDeviceAuthScope

`func (o *StatusDTO) GetDeviceAuthScope() string`

GetDeviceAuthScope returns the DeviceAuthScope field if non-nil, zero value otherwise.

### GetDeviceAuthScopeOk

`func (o *StatusDTO) GetDeviceAuthScopeOk() (*string, bool)`

GetDeviceAuthScopeOk returns a tuple with the DeviceAuthScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthScope

`func (o *StatusDTO) SetDeviceAuthScope(v string)`

SetDeviceAuthScope sets DeviceAuthScope field to given value.

### HasDeviceAuthScope

`func (o *StatusDTO) HasDeviceAuthScope() bool`

HasDeviceAuthScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


