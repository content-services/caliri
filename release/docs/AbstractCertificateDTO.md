# AbstractCertificateDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**Serial** | Pointer to [**CertificateSerialDTO**](CertificateSerialDTO.md) |  | [optional] 

## Methods

### NewAbstractCertificateDTO

`func NewAbstractCertificateDTO() *AbstractCertificateDTO`

NewAbstractCertificateDTO instantiates a new AbstractCertificateDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractCertificateDTOWithDefaults

`func NewAbstractCertificateDTOWithDefaults() *AbstractCertificateDTO`

NewAbstractCertificateDTOWithDefaults instantiates a new AbstractCertificateDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AbstractCertificateDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AbstractCertificateDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AbstractCertificateDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AbstractCertificateDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *AbstractCertificateDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AbstractCertificateDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AbstractCertificateDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AbstractCertificateDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *AbstractCertificateDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbstractCertificateDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbstractCertificateDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AbstractCertificateDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *AbstractCertificateDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AbstractCertificateDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AbstractCertificateDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AbstractCertificateDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCert

`func (o *AbstractCertificateDTO) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *AbstractCertificateDTO) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *AbstractCertificateDTO) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *AbstractCertificateDTO) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetSerial

`func (o *AbstractCertificateDTO) GetSerial() CertificateSerialDTO`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *AbstractCertificateDTO) GetSerialOk() (*CertificateSerialDTO, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *AbstractCertificateDTO) SetSerial(v CertificateSerialDTO)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *AbstractCertificateDTO) HasSerial() bool`

HasSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


