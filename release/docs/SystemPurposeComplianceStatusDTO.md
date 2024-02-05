# SystemPurposeComplianceStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Compliant** | Pointer to **bool** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**NonCompliantRole** | Pointer to **string** |  | [optional] 
**NonCompliantSLA** | Pointer to **string** |  | [optional] 
**NonCompliantUsage** | Pointer to **string** |  | [optional] 
**NonCompliantServiceType** | Pointer to **string** |  | [optional] 
**CompliantRole** | Pointer to [**map[string][]EntitlementDTO**](set.md) |  | [optional] 
**CompliantAddOns** | Pointer to [**map[string][]EntitlementDTO**](set.md) |  | [optional] 
**CompliantSLA** | Pointer to [**map[string][]EntitlementDTO**](set.md) |  | [optional] 
**CompliantUsage** | Pointer to [**map[string][]EntitlementDTO**](set.md) |  | [optional] 
**NonCompliantAddOns** | Pointer to **[]string** |  | [optional] 
**CompliantServiceType** | Pointer to [**map[string][]EntitlementDTO**](set.md) |  | [optional] 
**Reasons** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSystemPurposeComplianceStatusDTO

`func NewSystemPurposeComplianceStatusDTO() *SystemPurposeComplianceStatusDTO`

NewSystemPurposeComplianceStatusDTO instantiates a new SystemPurposeComplianceStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemPurposeComplianceStatusDTOWithDefaults

`func NewSystemPurposeComplianceStatusDTOWithDefaults() *SystemPurposeComplianceStatusDTO`

NewSystemPurposeComplianceStatusDTOWithDefaults instantiates a new SystemPurposeComplianceStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SystemPurposeComplianceStatusDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SystemPurposeComplianceStatusDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SystemPurposeComplianceStatusDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SystemPurposeComplianceStatusDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCompliant

`func (o *SystemPurposeComplianceStatusDTO) GetCompliant() bool`

GetCompliant returns the Compliant field if non-nil, zero value otherwise.

### GetCompliantOk

`func (o *SystemPurposeComplianceStatusDTO) GetCompliantOk() (*bool, bool)`

GetCompliantOk returns a tuple with the Compliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliant

`func (o *SystemPurposeComplianceStatusDTO) SetCompliant(v bool)`

SetCompliant sets Compliant field to given value.

### HasCompliant

`func (o *SystemPurposeComplianceStatusDTO) HasCompliant() bool`

HasCompliant returns a boolean if a field has been set.

### GetDate

`func (o *SystemPurposeComplianceStatusDTO) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SystemPurposeComplianceStatusDTO) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SystemPurposeComplianceStatusDTO) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *SystemPurposeComplianceStatusDTO) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetNonCompliantRole

`func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantRole() string`

GetNonCompliantRole returns the NonCompliantRole field if non-nil, zero value otherwise.

### GetNonCompliantRoleOk

`func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantRoleOk() (*string, bool)`

GetNonCompliantRoleOk returns a tuple with the NonCompliantRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantRole

`func (o *SystemPurposeComplianceStatusDTO) SetNonCompliantRole(v string)`

SetNonCompliantRole sets NonCompliantRole field to given value.

### HasNonCompliantRole

`func (o *SystemPurposeComplianceStatusDTO) HasNonCompliantRole() bool`

HasNonCompliantRole returns a boolean if a field has been set.

### GetNonCompliantSLA

`func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantSLA() string`

GetNonCompliantSLA returns the NonCompliantSLA field if non-nil, zero value otherwise.

### GetNonCompliantSLAOk

`func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantSLAOk() (*string, bool)`

GetNonCompliantSLAOk returns a tuple with the NonCompliantSLA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantSLA

`func (o *SystemPurposeComplianceStatusDTO) SetNonCompliantSLA(v string)`

SetNonCompliantSLA sets NonCompliantSLA field to given value.

### HasNonCompliantSLA

`func (o *SystemPurposeComplianceStatusDTO) HasNonCompliantSLA() bool`

HasNonCompliantSLA returns a boolean if a field has been set.

### GetNonCompliantUsage

`func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantUsage() string`

GetNonCompliantUsage returns the NonCompliantUsage field if non-nil, zero value otherwise.

### GetNonCompliantUsageOk

`func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantUsageOk() (*string, bool)`

GetNonCompliantUsageOk returns a tuple with the NonCompliantUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantUsage

`func (o *SystemPurposeComplianceStatusDTO) SetNonCompliantUsage(v string)`

SetNonCompliantUsage sets NonCompliantUsage field to given value.

### HasNonCompliantUsage

`func (o *SystemPurposeComplianceStatusDTO) HasNonCompliantUsage() bool`

HasNonCompliantUsage returns a boolean if a field has been set.

### GetNonCompliantServiceType

`func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantServiceType() string`

GetNonCompliantServiceType returns the NonCompliantServiceType field if non-nil, zero value otherwise.

### GetNonCompliantServiceTypeOk

`func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantServiceTypeOk() (*string, bool)`

GetNonCompliantServiceTypeOk returns a tuple with the NonCompliantServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantServiceType

`func (o *SystemPurposeComplianceStatusDTO) SetNonCompliantServiceType(v string)`

SetNonCompliantServiceType sets NonCompliantServiceType field to given value.

### HasNonCompliantServiceType

`func (o *SystemPurposeComplianceStatusDTO) HasNonCompliantServiceType() bool`

HasNonCompliantServiceType returns a boolean if a field has been set.

### GetCompliantRole

`func (o *SystemPurposeComplianceStatusDTO) GetCompliantRole() map[string][]EntitlementDTO`

GetCompliantRole returns the CompliantRole field if non-nil, zero value otherwise.

### GetCompliantRoleOk

`func (o *SystemPurposeComplianceStatusDTO) GetCompliantRoleOk() (*map[string][]EntitlementDTO, bool)`

GetCompliantRoleOk returns a tuple with the CompliantRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantRole

`func (o *SystemPurposeComplianceStatusDTO) SetCompliantRole(v map[string][]EntitlementDTO)`

SetCompliantRole sets CompliantRole field to given value.

### HasCompliantRole

`func (o *SystemPurposeComplianceStatusDTO) HasCompliantRole() bool`

HasCompliantRole returns a boolean if a field has been set.

### GetCompliantAddOns

`func (o *SystemPurposeComplianceStatusDTO) GetCompliantAddOns() map[string][]EntitlementDTO`

GetCompliantAddOns returns the CompliantAddOns field if non-nil, zero value otherwise.

### GetCompliantAddOnsOk

`func (o *SystemPurposeComplianceStatusDTO) GetCompliantAddOnsOk() (*map[string][]EntitlementDTO, bool)`

GetCompliantAddOnsOk returns a tuple with the CompliantAddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantAddOns

`func (o *SystemPurposeComplianceStatusDTO) SetCompliantAddOns(v map[string][]EntitlementDTO)`

SetCompliantAddOns sets CompliantAddOns field to given value.

### HasCompliantAddOns

`func (o *SystemPurposeComplianceStatusDTO) HasCompliantAddOns() bool`

HasCompliantAddOns returns a boolean if a field has been set.

### GetCompliantSLA

`func (o *SystemPurposeComplianceStatusDTO) GetCompliantSLA() map[string][]EntitlementDTO`

GetCompliantSLA returns the CompliantSLA field if non-nil, zero value otherwise.

### GetCompliantSLAOk

`func (o *SystemPurposeComplianceStatusDTO) GetCompliantSLAOk() (*map[string][]EntitlementDTO, bool)`

GetCompliantSLAOk returns a tuple with the CompliantSLA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantSLA

`func (o *SystemPurposeComplianceStatusDTO) SetCompliantSLA(v map[string][]EntitlementDTO)`

SetCompliantSLA sets CompliantSLA field to given value.

### HasCompliantSLA

`func (o *SystemPurposeComplianceStatusDTO) HasCompliantSLA() bool`

HasCompliantSLA returns a boolean if a field has been set.

### GetCompliantUsage

`func (o *SystemPurposeComplianceStatusDTO) GetCompliantUsage() map[string][]EntitlementDTO`

GetCompliantUsage returns the CompliantUsage field if non-nil, zero value otherwise.

### GetCompliantUsageOk

`func (o *SystemPurposeComplianceStatusDTO) GetCompliantUsageOk() (*map[string][]EntitlementDTO, bool)`

GetCompliantUsageOk returns a tuple with the CompliantUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantUsage

`func (o *SystemPurposeComplianceStatusDTO) SetCompliantUsage(v map[string][]EntitlementDTO)`

SetCompliantUsage sets CompliantUsage field to given value.

### HasCompliantUsage

`func (o *SystemPurposeComplianceStatusDTO) HasCompliantUsage() bool`

HasCompliantUsage returns a boolean if a field has been set.

### GetNonCompliantAddOns

`func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantAddOns() []string`

GetNonCompliantAddOns returns the NonCompliantAddOns field if non-nil, zero value otherwise.

### GetNonCompliantAddOnsOk

`func (o *SystemPurposeComplianceStatusDTO) GetNonCompliantAddOnsOk() (*[]string, bool)`

GetNonCompliantAddOnsOk returns a tuple with the NonCompliantAddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantAddOns

`func (o *SystemPurposeComplianceStatusDTO) SetNonCompliantAddOns(v []string)`

SetNonCompliantAddOns sets NonCompliantAddOns field to given value.

### HasNonCompliantAddOns

`func (o *SystemPurposeComplianceStatusDTO) HasNonCompliantAddOns() bool`

HasNonCompliantAddOns returns a boolean if a field has been set.

### GetCompliantServiceType

`func (o *SystemPurposeComplianceStatusDTO) GetCompliantServiceType() map[string][]EntitlementDTO`

GetCompliantServiceType returns the CompliantServiceType field if non-nil, zero value otherwise.

### GetCompliantServiceTypeOk

`func (o *SystemPurposeComplianceStatusDTO) GetCompliantServiceTypeOk() (*map[string][]EntitlementDTO, bool)`

GetCompliantServiceTypeOk returns a tuple with the CompliantServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantServiceType

`func (o *SystemPurposeComplianceStatusDTO) SetCompliantServiceType(v map[string][]EntitlementDTO)`

SetCompliantServiceType sets CompliantServiceType field to given value.

### HasCompliantServiceType

`func (o *SystemPurposeComplianceStatusDTO) HasCompliantServiceType() bool`

HasCompliantServiceType returns a boolean if a field has been set.

### GetReasons

`func (o *SystemPurposeComplianceStatusDTO) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *SystemPurposeComplianceStatusDTO) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *SystemPurposeComplianceStatusDTO) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *SystemPurposeComplianceStatusDTO) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


