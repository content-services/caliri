# OwnerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerCounts** | **map[string]int32** |  | 
**ConsumerGuestCounts** | **map[string]int32** |  | 
**EntitlementsConsumedByType** | **map[string]int32** |  | 
**ConsumerTypeCountByPool** | **map[string]int32** |  | 
**EnabledConsumerTypeCountByPool** | **map[string]int32** |  | 
**ConsumerCountsByComplianceStatus** | **map[string]int32** |  | 
**EntitlementsConsumedByFamily** | [**map[string]ConsumptionTypeCountsDTO**](ConsumptionTypeCountsDTO.md) |  | 

## Methods

### NewOwnerInfo

`func NewOwnerInfo(consumerCounts map[string]int32, consumerGuestCounts map[string]int32, entitlementsConsumedByType map[string]int32, consumerTypeCountByPool map[string]int32, enabledConsumerTypeCountByPool map[string]int32, consumerCountsByComplianceStatus map[string]int32, entitlementsConsumedByFamily map[string]ConsumptionTypeCountsDTO, ) *OwnerInfo`

NewOwnerInfo instantiates a new OwnerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerInfoWithDefaults

`func NewOwnerInfoWithDefaults() *OwnerInfo`

NewOwnerInfoWithDefaults instantiates a new OwnerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerCounts

`func (o *OwnerInfo) GetConsumerCounts() map[string]int32`

GetConsumerCounts returns the ConsumerCounts field if non-nil, zero value otherwise.

### GetConsumerCountsOk

`func (o *OwnerInfo) GetConsumerCountsOk() (*map[string]int32, bool)`

GetConsumerCountsOk returns a tuple with the ConsumerCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerCounts

`func (o *OwnerInfo) SetConsumerCounts(v map[string]int32)`

SetConsumerCounts sets ConsumerCounts field to given value.


### GetConsumerGuestCounts

`func (o *OwnerInfo) GetConsumerGuestCounts() map[string]int32`

GetConsumerGuestCounts returns the ConsumerGuestCounts field if non-nil, zero value otherwise.

### GetConsumerGuestCountsOk

`func (o *OwnerInfo) GetConsumerGuestCountsOk() (*map[string]int32, bool)`

GetConsumerGuestCountsOk returns a tuple with the ConsumerGuestCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerGuestCounts

`func (o *OwnerInfo) SetConsumerGuestCounts(v map[string]int32)`

SetConsumerGuestCounts sets ConsumerGuestCounts field to given value.


### GetEntitlementsConsumedByType

`func (o *OwnerInfo) GetEntitlementsConsumedByType() map[string]int32`

GetEntitlementsConsumedByType returns the EntitlementsConsumedByType field if non-nil, zero value otherwise.

### GetEntitlementsConsumedByTypeOk

`func (o *OwnerInfo) GetEntitlementsConsumedByTypeOk() (*map[string]int32, bool)`

GetEntitlementsConsumedByTypeOk returns a tuple with the EntitlementsConsumedByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsConsumedByType

`func (o *OwnerInfo) SetEntitlementsConsumedByType(v map[string]int32)`

SetEntitlementsConsumedByType sets EntitlementsConsumedByType field to given value.


### GetConsumerTypeCountByPool

`func (o *OwnerInfo) GetConsumerTypeCountByPool() map[string]int32`

GetConsumerTypeCountByPool returns the ConsumerTypeCountByPool field if non-nil, zero value otherwise.

### GetConsumerTypeCountByPoolOk

`func (o *OwnerInfo) GetConsumerTypeCountByPoolOk() (*map[string]int32, bool)`

GetConsumerTypeCountByPoolOk returns a tuple with the ConsumerTypeCountByPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerTypeCountByPool

`func (o *OwnerInfo) SetConsumerTypeCountByPool(v map[string]int32)`

SetConsumerTypeCountByPool sets ConsumerTypeCountByPool field to given value.


### GetEnabledConsumerTypeCountByPool

`func (o *OwnerInfo) GetEnabledConsumerTypeCountByPool() map[string]int32`

GetEnabledConsumerTypeCountByPool returns the EnabledConsumerTypeCountByPool field if non-nil, zero value otherwise.

### GetEnabledConsumerTypeCountByPoolOk

`func (o *OwnerInfo) GetEnabledConsumerTypeCountByPoolOk() (*map[string]int32, bool)`

GetEnabledConsumerTypeCountByPoolOk returns a tuple with the EnabledConsumerTypeCountByPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledConsumerTypeCountByPool

`func (o *OwnerInfo) SetEnabledConsumerTypeCountByPool(v map[string]int32)`

SetEnabledConsumerTypeCountByPool sets EnabledConsumerTypeCountByPool field to given value.


### GetConsumerCountsByComplianceStatus

`func (o *OwnerInfo) GetConsumerCountsByComplianceStatus() map[string]int32`

GetConsumerCountsByComplianceStatus returns the ConsumerCountsByComplianceStatus field if non-nil, zero value otherwise.

### GetConsumerCountsByComplianceStatusOk

`func (o *OwnerInfo) GetConsumerCountsByComplianceStatusOk() (*map[string]int32, bool)`

GetConsumerCountsByComplianceStatusOk returns a tuple with the ConsumerCountsByComplianceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerCountsByComplianceStatus

`func (o *OwnerInfo) SetConsumerCountsByComplianceStatus(v map[string]int32)`

SetConsumerCountsByComplianceStatus sets ConsumerCountsByComplianceStatus field to given value.


### GetEntitlementsConsumedByFamily

`func (o *OwnerInfo) GetEntitlementsConsumedByFamily() map[string]ConsumptionTypeCountsDTO`

GetEntitlementsConsumedByFamily returns the EntitlementsConsumedByFamily field if non-nil, zero value otherwise.

### GetEntitlementsConsumedByFamilyOk

`func (o *OwnerInfo) GetEntitlementsConsumedByFamilyOk() (*map[string]ConsumptionTypeCountsDTO, bool)`

GetEntitlementsConsumedByFamilyOk returns a tuple with the EntitlementsConsumedByFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsConsumedByFamily

`func (o *OwnerInfo) SetEntitlementsConsumedByFamily(v map[string]ConsumptionTypeCountsDTO)`

SetEntitlementsConsumedByFamily sets EntitlementsConsumedByFamily field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


