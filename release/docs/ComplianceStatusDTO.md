# ComplianceStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Compliant** | Pointer to **bool** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**CompliantUntil** | Pointer to **time.Time** |  | [optional] 
**CompliantProducts** | Pointer to [**map[string][]EntitlementDTO**](set.md) |  | [optional] 
**PartiallyCompliantProducts** | Pointer to [**map[string][]EntitlementDTO**](set.md) |  | [optional] 
**PartialStacks** | Pointer to [**map[string][]EntitlementDTO**](set.md) |  | [optional] 
**NonCompliantProducts** | Pointer to **[]string** |  | [optional] 
**Reasons** | Pointer to [**[]ComplianceReasonDTO**](ComplianceReasonDTO.md) |  | [optional] 
**ProductComplianceDateRanges** | Pointer to [**map[string]DateRange**](DateRange.md) |  | [optional] 

## Methods

### NewComplianceStatusDTO

`func NewComplianceStatusDTO() *ComplianceStatusDTO`

NewComplianceStatusDTO instantiates a new ComplianceStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceStatusDTOWithDefaults

`func NewComplianceStatusDTOWithDefaults() *ComplianceStatusDTO`

NewComplianceStatusDTOWithDefaults instantiates a new ComplianceStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ComplianceStatusDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComplianceStatusDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComplianceStatusDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComplianceStatusDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCompliant

`func (o *ComplianceStatusDTO) GetCompliant() bool`

GetCompliant returns the Compliant field if non-nil, zero value otherwise.

### GetCompliantOk

`func (o *ComplianceStatusDTO) GetCompliantOk() (*bool, bool)`

GetCompliantOk returns a tuple with the Compliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliant

`func (o *ComplianceStatusDTO) SetCompliant(v bool)`

SetCompliant sets Compliant field to given value.

### HasCompliant

`func (o *ComplianceStatusDTO) HasCompliant() bool`

HasCompliant returns a boolean if a field has been set.

### GetDate

`func (o *ComplianceStatusDTO) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ComplianceStatusDTO) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ComplianceStatusDTO) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ComplianceStatusDTO) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCompliantUntil

`func (o *ComplianceStatusDTO) GetCompliantUntil() time.Time`

GetCompliantUntil returns the CompliantUntil field if non-nil, zero value otherwise.

### GetCompliantUntilOk

`func (o *ComplianceStatusDTO) GetCompliantUntilOk() (*time.Time, bool)`

GetCompliantUntilOk returns a tuple with the CompliantUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantUntil

`func (o *ComplianceStatusDTO) SetCompliantUntil(v time.Time)`

SetCompliantUntil sets CompliantUntil field to given value.

### HasCompliantUntil

`func (o *ComplianceStatusDTO) HasCompliantUntil() bool`

HasCompliantUntil returns a boolean if a field has been set.

### GetCompliantProducts

`func (o *ComplianceStatusDTO) GetCompliantProducts() map[string][]EntitlementDTO`

GetCompliantProducts returns the CompliantProducts field if non-nil, zero value otherwise.

### GetCompliantProductsOk

`func (o *ComplianceStatusDTO) GetCompliantProductsOk() (*map[string][]EntitlementDTO, bool)`

GetCompliantProductsOk returns a tuple with the CompliantProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantProducts

`func (o *ComplianceStatusDTO) SetCompliantProducts(v map[string][]EntitlementDTO)`

SetCompliantProducts sets CompliantProducts field to given value.

### HasCompliantProducts

`func (o *ComplianceStatusDTO) HasCompliantProducts() bool`

HasCompliantProducts returns a boolean if a field has been set.

### GetPartiallyCompliantProducts

`func (o *ComplianceStatusDTO) GetPartiallyCompliantProducts() map[string][]EntitlementDTO`

GetPartiallyCompliantProducts returns the PartiallyCompliantProducts field if non-nil, zero value otherwise.

### GetPartiallyCompliantProductsOk

`func (o *ComplianceStatusDTO) GetPartiallyCompliantProductsOk() (*map[string][]EntitlementDTO, bool)`

GetPartiallyCompliantProductsOk returns a tuple with the PartiallyCompliantProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartiallyCompliantProducts

`func (o *ComplianceStatusDTO) SetPartiallyCompliantProducts(v map[string][]EntitlementDTO)`

SetPartiallyCompliantProducts sets PartiallyCompliantProducts field to given value.

### HasPartiallyCompliantProducts

`func (o *ComplianceStatusDTO) HasPartiallyCompliantProducts() bool`

HasPartiallyCompliantProducts returns a boolean if a field has been set.

### GetPartialStacks

`func (o *ComplianceStatusDTO) GetPartialStacks() map[string][]EntitlementDTO`

GetPartialStacks returns the PartialStacks field if non-nil, zero value otherwise.

### GetPartialStacksOk

`func (o *ComplianceStatusDTO) GetPartialStacksOk() (*map[string][]EntitlementDTO, bool)`

GetPartialStacksOk returns a tuple with the PartialStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialStacks

`func (o *ComplianceStatusDTO) SetPartialStacks(v map[string][]EntitlementDTO)`

SetPartialStacks sets PartialStacks field to given value.

### HasPartialStacks

`func (o *ComplianceStatusDTO) HasPartialStacks() bool`

HasPartialStacks returns a boolean if a field has been set.

### GetNonCompliantProducts

`func (o *ComplianceStatusDTO) GetNonCompliantProducts() []string`

GetNonCompliantProducts returns the NonCompliantProducts field if non-nil, zero value otherwise.

### GetNonCompliantProductsOk

`func (o *ComplianceStatusDTO) GetNonCompliantProductsOk() (*[]string, bool)`

GetNonCompliantProductsOk returns a tuple with the NonCompliantProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantProducts

`func (o *ComplianceStatusDTO) SetNonCompliantProducts(v []string)`

SetNonCompliantProducts sets NonCompliantProducts field to given value.

### HasNonCompliantProducts

`func (o *ComplianceStatusDTO) HasNonCompliantProducts() bool`

HasNonCompliantProducts returns a boolean if a field has been set.

### GetReasons

`func (o *ComplianceStatusDTO) GetReasons() []ComplianceReasonDTO`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *ComplianceStatusDTO) GetReasonsOk() (*[]ComplianceReasonDTO, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *ComplianceStatusDTO) SetReasons(v []ComplianceReasonDTO)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *ComplianceStatusDTO) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetProductComplianceDateRanges

`func (o *ComplianceStatusDTO) GetProductComplianceDateRanges() map[string]DateRange`

GetProductComplianceDateRanges returns the ProductComplianceDateRanges field if non-nil, zero value otherwise.

### GetProductComplianceDateRangesOk

`func (o *ComplianceStatusDTO) GetProductComplianceDateRangesOk() (*map[string]DateRange, bool)`

GetProductComplianceDateRangesOk returns a tuple with the ProductComplianceDateRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductComplianceDateRanges

`func (o *ComplianceStatusDTO) SetProductComplianceDateRanges(v map[string]DateRange)`

SetProductComplianceDateRanges sets ProductComplianceDateRanges field to given value.

### HasProductComplianceDateRanges

`func (o *ComplianceStatusDTO) HasProductComplianceDateRanges() bool`

HasProductComplianceDateRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


