# ConsumerTypeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Manifest** | Pointer to **bool** |  | [optional] 

## Methods

### NewConsumerTypeDTO

`func NewConsumerTypeDTO() *ConsumerTypeDTO`

NewConsumerTypeDTO instantiates a new ConsumerTypeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerTypeDTOWithDefaults

`func NewConsumerTypeDTOWithDefaults() *ConsumerTypeDTO`

NewConsumerTypeDTOWithDefaults instantiates a new ConsumerTypeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ConsumerTypeDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConsumerTypeDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConsumerTypeDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ConsumerTypeDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ConsumerTypeDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ConsumerTypeDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ConsumerTypeDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ConsumerTypeDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *ConsumerTypeDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsumerTypeDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsumerTypeDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsumerTypeDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ConsumerTypeDTO) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConsumerTypeDTO) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConsumerTypeDTO) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConsumerTypeDTO) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetManifest

`func (o *ConsumerTypeDTO) GetManifest() bool`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *ConsumerTypeDTO) GetManifestOk() (*bool, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *ConsumerTypeDTO) SetManifest(v bool)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *ConsumerTypeDTO) HasManifest() bool`

HasManifest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


