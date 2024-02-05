# ContentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**ContentUrl** | Pointer to **string** |  | [optional] 
**RequiredTags** | Pointer to **string** |  | [optional] 
**ReleaseVer** | Pointer to **string** |  | [optional] 
**GpgUrl** | Pointer to **string** |  | [optional] 
**ModifiedProductIds** | Pointer to **[]string** |  | [optional] 
**Arches** | Pointer to **string** |  | [optional] 
**MetadataExpire** | Pointer to **int64** |  | [optional] 

## Methods

### NewContentDTO

`func NewContentDTO() *ContentDTO`

NewContentDTO instantiates a new ContentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentDTOWithDefaults

`func NewContentDTOWithDefaults() *ContentDTO`

NewContentDTOWithDefaults instantiates a new ContentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ContentDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ContentDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ContentDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ContentDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ContentDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ContentDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ContentDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ContentDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUuid

`func (o *ContentDTO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ContentDTO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ContentDTO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ContentDTO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetId

`func (o *ContentDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContentDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ContentDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLabel

`func (o *ContentDTO) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContentDTO) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContentDTO) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ContentDTO) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *ContentDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVendor

`func (o *ContentDTO) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ContentDTO) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ContentDTO) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ContentDTO) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetContentUrl

`func (o *ContentDTO) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *ContentDTO) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *ContentDTO) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *ContentDTO) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### GetRequiredTags

`func (o *ContentDTO) GetRequiredTags() string`

GetRequiredTags returns the RequiredTags field if non-nil, zero value otherwise.

### GetRequiredTagsOk

`func (o *ContentDTO) GetRequiredTagsOk() (*string, bool)`

GetRequiredTagsOk returns a tuple with the RequiredTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredTags

`func (o *ContentDTO) SetRequiredTags(v string)`

SetRequiredTags sets RequiredTags field to given value.

### HasRequiredTags

`func (o *ContentDTO) HasRequiredTags() bool`

HasRequiredTags returns a boolean if a field has been set.

### GetReleaseVer

`func (o *ContentDTO) GetReleaseVer() string`

GetReleaseVer returns the ReleaseVer field if non-nil, zero value otherwise.

### GetReleaseVerOk

`func (o *ContentDTO) GetReleaseVerOk() (*string, bool)`

GetReleaseVerOk returns a tuple with the ReleaseVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVer

`func (o *ContentDTO) SetReleaseVer(v string)`

SetReleaseVer sets ReleaseVer field to given value.

### HasReleaseVer

`func (o *ContentDTO) HasReleaseVer() bool`

HasReleaseVer returns a boolean if a field has been set.

### GetGpgUrl

`func (o *ContentDTO) GetGpgUrl() string`

GetGpgUrl returns the GpgUrl field if non-nil, zero value otherwise.

### GetGpgUrlOk

`func (o *ContentDTO) GetGpgUrlOk() (*string, bool)`

GetGpgUrlOk returns a tuple with the GpgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgUrl

`func (o *ContentDTO) SetGpgUrl(v string)`

SetGpgUrl sets GpgUrl field to given value.

### HasGpgUrl

`func (o *ContentDTO) HasGpgUrl() bool`

HasGpgUrl returns a boolean if a field has been set.

### GetModifiedProductIds

`func (o *ContentDTO) GetModifiedProductIds() []string`

GetModifiedProductIds returns the ModifiedProductIds field if non-nil, zero value otherwise.

### GetModifiedProductIdsOk

`func (o *ContentDTO) GetModifiedProductIdsOk() (*[]string, bool)`

GetModifiedProductIdsOk returns a tuple with the ModifiedProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedProductIds

`func (o *ContentDTO) SetModifiedProductIds(v []string)`

SetModifiedProductIds sets ModifiedProductIds field to given value.

### HasModifiedProductIds

`func (o *ContentDTO) HasModifiedProductIds() bool`

HasModifiedProductIds returns a boolean if a field has been set.

### GetArches

`func (o *ContentDTO) GetArches() string`

GetArches returns the Arches field if non-nil, zero value otherwise.

### GetArchesOk

`func (o *ContentDTO) GetArchesOk() (*string, bool)`

GetArchesOk returns a tuple with the Arches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArches

`func (o *ContentDTO) SetArches(v string)`

SetArches sets Arches field to given value.

### HasArches

`func (o *ContentDTO) HasArches() bool`

HasArches returns a boolean if a field has been set.

### GetMetadataExpire

`func (o *ContentDTO) GetMetadataExpire() int64`

GetMetadataExpire returns the MetadataExpire field if non-nil, zero value otherwise.

### GetMetadataExpireOk

`func (o *ContentDTO) GetMetadataExpireOk() (*int64, bool)`

GetMetadataExpireOk returns a tuple with the MetadataExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataExpire

`func (o *ContentDTO) SetMetadataExpire(v int64)`

SetMetadataExpire sets MetadataExpire field to given value.

### HasMetadataExpire

`func (o *ContentDTO) HasMetadataExpire() bool`

HasMetadataExpire returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


