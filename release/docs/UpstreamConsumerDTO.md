# UpstreamConsumerDTO

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
**IdCert** | Pointer to [**CertificateDTO**](CertificateDTO.md) |  | [optional] 

## Methods

### NewUpstreamConsumerDTO

`func NewUpstreamConsumerDTO() *UpstreamConsumerDTO`

NewUpstreamConsumerDTO instantiates a new UpstreamConsumerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamConsumerDTOWithDefaults

`func NewUpstreamConsumerDTOWithDefaults() *UpstreamConsumerDTO`

NewUpstreamConsumerDTOWithDefaults instantiates a new UpstreamConsumerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UpstreamConsumerDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UpstreamConsumerDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UpstreamConsumerDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UpstreamConsumerDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *UpstreamConsumerDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UpstreamConsumerDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UpstreamConsumerDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *UpstreamConsumerDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *UpstreamConsumerDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpstreamConsumerDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpstreamConsumerDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpstreamConsumerDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *UpstreamConsumerDTO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UpstreamConsumerDTO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UpstreamConsumerDTO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UpstreamConsumerDTO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *UpstreamConsumerDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpstreamConsumerDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpstreamConsumerDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpstreamConsumerDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApiUrl

`func (o *UpstreamConsumerDTO) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *UpstreamConsumerDTO) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *UpstreamConsumerDTO) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *UpstreamConsumerDTO) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetWebUrl

`func (o *UpstreamConsumerDTO) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *UpstreamConsumerDTO) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *UpstreamConsumerDTO) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *UpstreamConsumerDTO) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetOwnerId

`func (o *UpstreamConsumerDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UpstreamConsumerDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UpstreamConsumerDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UpstreamConsumerDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetContentAccessMode

`func (o *UpstreamConsumerDTO) GetContentAccessMode() string`

GetContentAccessMode returns the ContentAccessMode field if non-nil, zero value otherwise.

### GetContentAccessModeOk

`func (o *UpstreamConsumerDTO) GetContentAccessModeOk() (*string, bool)`

GetContentAccessModeOk returns a tuple with the ContentAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAccessMode

`func (o *UpstreamConsumerDTO) SetContentAccessMode(v string)`

SetContentAccessMode sets ContentAccessMode field to given value.

### HasContentAccessMode

`func (o *UpstreamConsumerDTO) HasContentAccessMode() bool`

HasContentAccessMode returns a boolean if a field has been set.

### GetType

`func (o *UpstreamConsumerDTO) GetType() ConsumerTypeDTO`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpstreamConsumerDTO) GetTypeOk() (*ConsumerTypeDTO, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpstreamConsumerDTO) SetType(v ConsumerTypeDTO)`

SetType sets Type field to given value.

### HasType

`func (o *UpstreamConsumerDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIdCert

`func (o *UpstreamConsumerDTO) GetIdCert() CertificateDTO`

GetIdCert returns the IdCert field if non-nil, zero value otherwise.

### GetIdCertOk

`func (o *UpstreamConsumerDTO) GetIdCertOk() (*CertificateDTO, bool)`

GetIdCertOk returns a tuple with the IdCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCert

`func (o *UpstreamConsumerDTO) SetIdCert(v CertificateDTO)`

SetIdCert sets IdCert field to given value.

### HasIdCert

`func (o *UpstreamConsumerDTO) HasIdCert() bool`

HasIdCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


