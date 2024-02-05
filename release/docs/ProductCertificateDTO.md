# ProductCertificateDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 

## Methods

### NewProductCertificateDTO

`func NewProductCertificateDTO() *ProductCertificateDTO`

NewProductCertificateDTO instantiates a new ProductCertificateDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCertificateDTOWithDefaults

`func NewProductCertificateDTOWithDefaults() *ProductCertificateDTO`

NewProductCertificateDTOWithDefaults instantiates a new ProductCertificateDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ProductCertificateDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ProductCertificateDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ProductCertificateDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ProductCertificateDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ProductCertificateDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ProductCertificateDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ProductCertificateDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ProductCertificateDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetKey

`func (o *ProductCertificateDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProductCertificateDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProductCertificateDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProductCertificateDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCert

`func (o *ProductCertificateDTO) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *ProductCertificateDTO) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *ProductCertificateDTO) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *ProductCertificateDTO) HasCert() bool`

HasCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


