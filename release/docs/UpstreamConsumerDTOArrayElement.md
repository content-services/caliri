# UpstreamConsumerDTOArrayElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **string** |  | [optional] 
**WebUrl** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**ContentAccessMode** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ConsumerTypeDTO**](ConsumerTypeDTO.md) |  | [optional] 

## Methods

### NewUpstreamConsumerDTOArrayElement

`func NewUpstreamConsumerDTOArrayElement() *UpstreamConsumerDTOArrayElement`

NewUpstreamConsumerDTOArrayElement instantiates a new UpstreamConsumerDTOArrayElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamConsumerDTOArrayElementWithDefaults

`func NewUpstreamConsumerDTOArrayElementWithDefaults() *UpstreamConsumerDTOArrayElement`

NewUpstreamConsumerDTOArrayElementWithDefaults instantiates a new UpstreamConsumerDTOArrayElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UpstreamConsumerDTOArrayElement) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UpstreamConsumerDTOArrayElement) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UpstreamConsumerDTOArrayElement) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UpstreamConsumerDTOArrayElement) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *UpstreamConsumerDTOArrayElement) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UpstreamConsumerDTOArrayElement) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UpstreamConsumerDTOArrayElement) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *UpstreamConsumerDTOArrayElement) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *UpstreamConsumerDTOArrayElement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpstreamConsumerDTOArrayElement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpstreamConsumerDTOArrayElement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpstreamConsumerDTOArrayElement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *UpstreamConsumerDTOArrayElement) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UpstreamConsumerDTOArrayElement) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UpstreamConsumerDTOArrayElement) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UpstreamConsumerDTOArrayElement) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *UpstreamConsumerDTOArrayElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpstreamConsumerDTOArrayElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpstreamConsumerDTOArrayElement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpstreamConsumerDTOArrayElement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApiUrl

`func (o *UpstreamConsumerDTOArrayElement) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *UpstreamConsumerDTOArrayElement) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *UpstreamConsumerDTOArrayElement) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *UpstreamConsumerDTOArrayElement) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetWebUrl

`func (o *UpstreamConsumerDTOArrayElement) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *UpstreamConsumerDTOArrayElement) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *UpstreamConsumerDTOArrayElement) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *UpstreamConsumerDTOArrayElement) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetOwnerId

`func (o *UpstreamConsumerDTOArrayElement) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UpstreamConsumerDTOArrayElement) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UpstreamConsumerDTOArrayElement) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UpstreamConsumerDTOArrayElement) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetContentAccessMode

`func (o *UpstreamConsumerDTOArrayElement) GetContentAccessMode() string`

GetContentAccessMode returns the ContentAccessMode field if non-nil, zero value otherwise.

### GetContentAccessModeOk

`func (o *UpstreamConsumerDTOArrayElement) GetContentAccessModeOk() (*string, bool)`

GetContentAccessModeOk returns a tuple with the ContentAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAccessMode

`func (o *UpstreamConsumerDTOArrayElement) SetContentAccessMode(v string)`

SetContentAccessMode sets ContentAccessMode field to given value.

### HasContentAccessMode

`func (o *UpstreamConsumerDTOArrayElement) HasContentAccessMode() bool`

HasContentAccessMode returns a boolean if a field has been set.

### GetType

`func (o *UpstreamConsumerDTOArrayElement) GetType() ConsumerTypeDTO`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpstreamConsumerDTOArrayElement) GetTypeOk() (*ConsumerTypeDTO, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpstreamConsumerDTOArrayElement) SetType(v ConsumerTypeDTO)`

SetType sets Type field to given value.

### HasType

`func (o *UpstreamConsumerDTOArrayElement) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


