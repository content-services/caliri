# GuestIdDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**GuestId** | **string** |  | 
**Attributes** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGuestIdDTO

`func NewGuestIdDTO(guestId string, ) *GuestIdDTO`

NewGuestIdDTO instantiates a new GuestIdDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestIdDTOWithDefaults

`func NewGuestIdDTOWithDefaults() *GuestIdDTO`

NewGuestIdDTOWithDefaults instantiates a new GuestIdDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GuestIdDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GuestIdDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GuestIdDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GuestIdDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *GuestIdDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GuestIdDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GuestIdDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GuestIdDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *GuestIdDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuestIdDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuestIdDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GuestIdDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGuestId

`func (o *GuestIdDTO) GetGuestId() string`

GetGuestId returns the GuestId field if non-nil, zero value otherwise.

### GetGuestIdOk

`func (o *GuestIdDTO) GetGuestIdOk() (*string, bool)`

GetGuestIdOk returns a tuple with the GuestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestId

`func (o *GuestIdDTO) SetGuestId(v string)`

SetGuestId sets GuestId field to given value.


### GetAttributes

`func (o *GuestIdDTO) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GuestIdDTO) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GuestIdDTO) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GuestIdDTO) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


