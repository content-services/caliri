# TimestampedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 

## Methods

### NewTimestampedEntity

`func NewTimestampedEntity() *TimestampedEntity`

NewTimestampedEntity instantiates a new TimestampedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimestampedEntityWithDefaults

`func NewTimestampedEntityWithDefaults() *TimestampedEntity`

NewTimestampedEntityWithDefaults instantiates a new TimestampedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *TimestampedEntity) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TimestampedEntity) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TimestampedEntity) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TimestampedEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *TimestampedEntity) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TimestampedEntity) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TimestampedEntity) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *TimestampedEntity) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


