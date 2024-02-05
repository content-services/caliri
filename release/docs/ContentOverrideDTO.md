# ContentOverrideDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ContentLabel** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewContentOverrideDTO

`func NewContentOverrideDTO() *ContentOverrideDTO`

NewContentOverrideDTO instantiates a new ContentOverrideDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentOverrideDTOWithDefaults

`func NewContentOverrideDTOWithDefaults() *ContentOverrideDTO`

NewContentOverrideDTOWithDefaults instantiates a new ContentOverrideDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ContentOverrideDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ContentOverrideDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ContentOverrideDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ContentOverrideDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ContentOverrideDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ContentOverrideDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ContentOverrideDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ContentOverrideDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetName

`func (o *ContentOverrideDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentOverrideDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentOverrideDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentOverrideDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContentLabel

`func (o *ContentOverrideDTO) GetContentLabel() string`

GetContentLabel returns the ContentLabel field if non-nil, zero value otherwise.

### GetContentLabelOk

`func (o *ContentOverrideDTO) GetContentLabelOk() (*string, bool)`

GetContentLabelOk returns a tuple with the ContentLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLabel

`func (o *ContentOverrideDTO) SetContentLabel(v string)`

SetContentLabel sets ContentLabel field to given value.

### HasContentLabel

`func (o *ContentOverrideDTO) HasContentLabel() bool`

HasContentLabel returns a boolean if a field has been set.

### GetValue

`func (o *ContentOverrideDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentOverrideDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentOverrideDTO) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContentOverrideDTO) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


