# ImportUpstreamConsumerDTO

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

### NewImportUpstreamConsumerDTO

`func NewImportUpstreamConsumerDTO() *ImportUpstreamConsumerDTO`

NewImportUpstreamConsumerDTO instantiates a new ImportUpstreamConsumerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportUpstreamConsumerDTOWithDefaults

`func NewImportUpstreamConsumerDTOWithDefaults() *ImportUpstreamConsumerDTO`

NewImportUpstreamConsumerDTOWithDefaults instantiates a new ImportUpstreamConsumerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ImportUpstreamConsumerDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ImportUpstreamConsumerDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ImportUpstreamConsumerDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ImportUpstreamConsumerDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ImportUpstreamConsumerDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ImportUpstreamConsumerDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ImportUpstreamConsumerDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ImportUpstreamConsumerDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *ImportUpstreamConsumerDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportUpstreamConsumerDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportUpstreamConsumerDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImportUpstreamConsumerDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ImportUpstreamConsumerDTO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ImportUpstreamConsumerDTO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ImportUpstreamConsumerDTO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ImportUpstreamConsumerDTO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *ImportUpstreamConsumerDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportUpstreamConsumerDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportUpstreamConsumerDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImportUpstreamConsumerDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApiUrl

`func (o *ImportUpstreamConsumerDTO) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ImportUpstreamConsumerDTO) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ImportUpstreamConsumerDTO) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ImportUpstreamConsumerDTO) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetWebUrl

`func (o *ImportUpstreamConsumerDTO) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *ImportUpstreamConsumerDTO) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *ImportUpstreamConsumerDTO) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *ImportUpstreamConsumerDTO) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetOwnerId

`func (o *ImportUpstreamConsumerDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ImportUpstreamConsumerDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ImportUpstreamConsumerDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ImportUpstreamConsumerDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetContentAccessMode

`func (o *ImportUpstreamConsumerDTO) GetContentAccessMode() string`

GetContentAccessMode returns the ContentAccessMode field if non-nil, zero value otherwise.

### GetContentAccessModeOk

`func (o *ImportUpstreamConsumerDTO) GetContentAccessModeOk() (*string, bool)`

GetContentAccessModeOk returns a tuple with the ContentAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAccessMode

`func (o *ImportUpstreamConsumerDTO) SetContentAccessMode(v string)`

SetContentAccessMode sets ContentAccessMode field to given value.

### HasContentAccessMode

`func (o *ImportUpstreamConsumerDTO) HasContentAccessMode() bool`

HasContentAccessMode returns a boolean if a field has been set.

### GetType

`func (o *ImportUpstreamConsumerDTO) GetType() ConsumerTypeDTO`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportUpstreamConsumerDTO) GetTypeOk() (*ConsumerTypeDTO, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportUpstreamConsumerDTO) SetType(v ConsumerTypeDTO)`

SetType sets Type field to given value.

### HasType

`func (o *ImportUpstreamConsumerDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


