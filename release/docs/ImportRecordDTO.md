# ImportRecordDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**GeneratedBy** | Pointer to **string** |  | [optional] 
**GeneratedDate** | Pointer to **time.Time** |  | [optional] 
**UpstreamConsumer** | Pointer to [**ImportUpstreamConsumerDTO**](ImportUpstreamConsumerDTO.md) |  | [optional] 

## Methods

### NewImportRecordDTO

`func NewImportRecordDTO() *ImportRecordDTO`

NewImportRecordDTO instantiates a new ImportRecordDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRecordDTOWithDefaults

`func NewImportRecordDTOWithDefaults() *ImportRecordDTO`

NewImportRecordDTOWithDefaults instantiates a new ImportRecordDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ImportRecordDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ImportRecordDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ImportRecordDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ImportRecordDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ImportRecordDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ImportRecordDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ImportRecordDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ImportRecordDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *ImportRecordDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportRecordDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportRecordDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImportRecordDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ImportRecordDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportRecordDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportRecordDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImportRecordDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ImportRecordDTO) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ImportRecordDTO) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ImportRecordDTO) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ImportRecordDTO) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetFileName

`func (o *ImportRecordDTO) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ImportRecordDTO) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ImportRecordDTO) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ImportRecordDTO) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetGeneratedBy

`func (o *ImportRecordDTO) GetGeneratedBy() string`

GetGeneratedBy returns the GeneratedBy field if non-nil, zero value otherwise.

### GetGeneratedByOk

`func (o *ImportRecordDTO) GetGeneratedByOk() (*string, bool)`

GetGeneratedByOk returns a tuple with the GeneratedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedBy

`func (o *ImportRecordDTO) SetGeneratedBy(v string)`

SetGeneratedBy sets GeneratedBy field to given value.

### HasGeneratedBy

`func (o *ImportRecordDTO) HasGeneratedBy() bool`

HasGeneratedBy returns a boolean if a field has been set.

### GetGeneratedDate

`func (o *ImportRecordDTO) GetGeneratedDate() time.Time`

GetGeneratedDate returns the GeneratedDate field if non-nil, zero value otherwise.

### GetGeneratedDateOk

`func (o *ImportRecordDTO) GetGeneratedDateOk() (*time.Time, bool)`

GetGeneratedDateOk returns a tuple with the GeneratedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedDate

`func (o *ImportRecordDTO) SetGeneratedDate(v time.Time)`

SetGeneratedDate sets GeneratedDate field to given value.

### HasGeneratedDate

`func (o *ImportRecordDTO) HasGeneratedDate() bool`

HasGeneratedDate returns a boolean if a field has been set.

### GetUpstreamConsumer

`func (o *ImportRecordDTO) GetUpstreamConsumer() ImportUpstreamConsumerDTO`

GetUpstreamConsumer returns the UpstreamConsumer field if non-nil, zero value otherwise.

### GetUpstreamConsumerOk

`func (o *ImportRecordDTO) GetUpstreamConsumerOk() (*ImportUpstreamConsumerDTO, bool)`

GetUpstreamConsumerOk returns a tuple with the UpstreamConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamConsumer

`func (o *ImportRecordDTO) SetUpstreamConsumer(v ImportUpstreamConsumerDTO)`

SetUpstreamConsumer sets UpstreamConsumer field to given value.

### HasUpstreamConsumer

`func (o *ImportRecordDTO) HasUpstreamConsumer() bool`

HasUpstreamConsumer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


