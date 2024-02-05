# ConsumerInstalledProductDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | **string** |  | 
**ProductName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Arch** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewConsumerInstalledProductDTO

`func NewConsumerInstalledProductDTO(productId string, ) *ConsumerInstalledProductDTO`

NewConsumerInstalledProductDTO instantiates a new ConsumerInstalledProductDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerInstalledProductDTOWithDefaults

`func NewConsumerInstalledProductDTOWithDefaults() *ConsumerInstalledProductDTO`

NewConsumerInstalledProductDTOWithDefaults instantiates a new ConsumerInstalledProductDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ConsumerInstalledProductDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConsumerInstalledProductDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConsumerInstalledProductDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ConsumerInstalledProductDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ConsumerInstalledProductDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ConsumerInstalledProductDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ConsumerInstalledProductDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ConsumerInstalledProductDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *ConsumerInstalledProductDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsumerInstalledProductDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsumerInstalledProductDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsumerInstalledProductDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *ConsumerInstalledProductDTO) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ConsumerInstalledProductDTO) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ConsumerInstalledProductDTO) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetProductName

`func (o *ConsumerInstalledProductDTO) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ConsumerInstalledProductDTO) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ConsumerInstalledProductDTO) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ConsumerInstalledProductDTO) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetVersion

`func (o *ConsumerInstalledProductDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConsumerInstalledProductDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConsumerInstalledProductDTO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConsumerInstalledProductDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetArch

`func (o *ConsumerInstalledProductDTO) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ConsumerInstalledProductDTO) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ConsumerInstalledProductDTO) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ConsumerInstalledProductDTO) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetStatus

`func (o *ConsumerInstalledProductDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConsumerInstalledProductDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConsumerInstalledProductDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConsumerInstalledProductDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *ConsumerInstalledProductDTO) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ConsumerInstalledProductDTO) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ConsumerInstalledProductDTO) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ConsumerInstalledProductDTO) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ConsumerInstalledProductDTO) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ConsumerInstalledProductDTO) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ConsumerInstalledProductDTO) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ConsumerInstalledProductDTO) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


