# EntitlementDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Consumer** | Pointer to [**NestedConsumerDTO**](NestedConsumerDTO.md) |  | [optional] 
**Pool** | Pointer to [**PoolDTO**](PoolDTO.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Certificates** | Pointer to [**[]CertificateDTO**](CertificateDTO.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewEntitlementDTO

`func NewEntitlementDTO() *EntitlementDTO`

NewEntitlementDTO instantiates a new EntitlementDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementDTOWithDefaults

`func NewEntitlementDTOWithDefaults() *EntitlementDTO`

NewEntitlementDTOWithDefaults instantiates a new EntitlementDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *EntitlementDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntitlementDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntitlementDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EntitlementDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *EntitlementDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *EntitlementDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *EntitlementDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *EntitlementDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *EntitlementDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConsumer

`func (o *EntitlementDTO) GetConsumer() NestedConsumerDTO`

GetConsumer returns the Consumer field if non-nil, zero value otherwise.

### GetConsumerOk

`func (o *EntitlementDTO) GetConsumerOk() (*NestedConsumerDTO, bool)`

GetConsumerOk returns a tuple with the Consumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumer

`func (o *EntitlementDTO) SetConsumer(v NestedConsumerDTO)`

SetConsumer sets Consumer field to given value.

### HasConsumer

`func (o *EntitlementDTO) HasConsumer() bool`

HasConsumer returns a boolean if a field has been set.

### GetPool

`func (o *EntitlementDTO) GetPool() PoolDTO`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *EntitlementDTO) GetPoolOk() (*PoolDTO, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *EntitlementDTO) SetPool(v PoolDTO)`

SetPool sets Pool field to given value.

### HasPool

`func (o *EntitlementDTO) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetQuantity

`func (o *EntitlementDTO) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *EntitlementDTO) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *EntitlementDTO) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *EntitlementDTO) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCertificates

`func (o *EntitlementDTO) GetCertificates() []CertificateDTO`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *EntitlementDTO) GetCertificatesOk() (*[]CertificateDTO, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *EntitlementDTO) SetCertificates(v []CertificateDTO)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *EntitlementDTO) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetStartDate

`func (o *EntitlementDTO) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EntitlementDTO) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EntitlementDTO) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EntitlementDTO) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *EntitlementDTO) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EntitlementDTO) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EntitlementDTO) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EntitlementDTO) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetHref

`func (o *EntitlementDTO) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *EntitlementDTO) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *EntitlementDTO) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *EntitlementDTO) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


