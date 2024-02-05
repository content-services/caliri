# DeletedConsumerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ConsumerUuid** | Pointer to **string** |  | [optional] 
**ConsumerName** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerKey** | Pointer to **string** |  | [optional] 
**OwnerDisplayName** | Pointer to **string** |  | [optional] 
**PrincipalName** | Pointer to **string** |  | [optional] 

## Methods

### NewDeletedConsumerDTO

`func NewDeletedConsumerDTO() *DeletedConsumerDTO`

NewDeletedConsumerDTO instantiates a new DeletedConsumerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedConsumerDTOWithDefaults

`func NewDeletedConsumerDTOWithDefaults() *DeletedConsumerDTO`

NewDeletedConsumerDTOWithDefaults instantiates a new DeletedConsumerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DeletedConsumerDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeletedConsumerDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeletedConsumerDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DeletedConsumerDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *DeletedConsumerDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *DeletedConsumerDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *DeletedConsumerDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *DeletedConsumerDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *DeletedConsumerDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeletedConsumerDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeletedConsumerDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeletedConsumerDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConsumerUuid

`func (o *DeletedConsumerDTO) GetConsumerUuid() string`

GetConsumerUuid returns the ConsumerUuid field if non-nil, zero value otherwise.

### GetConsumerUuidOk

`func (o *DeletedConsumerDTO) GetConsumerUuidOk() (*string, bool)`

GetConsumerUuidOk returns a tuple with the ConsumerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerUuid

`func (o *DeletedConsumerDTO) SetConsumerUuid(v string)`

SetConsumerUuid sets ConsumerUuid field to given value.

### HasConsumerUuid

`func (o *DeletedConsumerDTO) HasConsumerUuid() bool`

HasConsumerUuid returns a boolean if a field has been set.

### GetConsumerName

`func (o *DeletedConsumerDTO) GetConsumerName() string`

GetConsumerName returns the ConsumerName field if non-nil, zero value otherwise.

### GetConsumerNameOk

`func (o *DeletedConsumerDTO) GetConsumerNameOk() (*string, bool)`

GetConsumerNameOk returns a tuple with the ConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerName

`func (o *DeletedConsumerDTO) SetConsumerName(v string)`

SetConsumerName sets ConsumerName field to given value.

### HasConsumerName

`func (o *DeletedConsumerDTO) HasConsumerName() bool`

HasConsumerName returns a boolean if a field has been set.

### GetOwnerId

`func (o *DeletedConsumerDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DeletedConsumerDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DeletedConsumerDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *DeletedConsumerDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerKey

`func (o *DeletedConsumerDTO) GetOwnerKey() string`

GetOwnerKey returns the OwnerKey field if non-nil, zero value otherwise.

### GetOwnerKeyOk

`func (o *DeletedConsumerDTO) GetOwnerKeyOk() (*string, bool)`

GetOwnerKeyOk returns a tuple with the OwnerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerKey

`func (o *DeletedConsumerDTO) SetOwnerKey(v string)`

SetOwnerKey sets OwnerKey field to given value.

### HasOwnerKey

`func (o *DeletedConsumerDTO) HasOwnerKey() bool`

HasOwnerKey returns a boolean if a field has been set.

### GetOwnerDisplayName

`func (o *DeletedConsumerDTO) GetOwnerDisplayName() string`

GetOwnerDisplayName returns the OwnerDisplayName field if non-nil, zero value otherwise.

### GetOwnerDisplayNameOk

`func (o *DeletedConsumerDTO) GetOwnerDisplayNameOk() (*string, bool)`

GetOwnerDisplayNameOk returns a tuple with the OwnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerDisplayName

`func (o *DeletedConsumerDTO) SetOwnerDisplayName(v string)`

SetOwnerDisplayName sets OwnerDisplayName field to given value.

### HasOwnerDisplayName

`func (o *DeletedConsumerDTO) HasOwnerDisplayName() bool`

HasOwnerDisplayName returns a boolean if a field has been set.

### GetPrincipalName

`func (o *DeletedConsumerDTO) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *DeletedConsumerDTO) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *DeletedConsumerDTO) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *DeletedConsumerDTO) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


