# UeberCertificateDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**Serial** | Pointer to [**CertificateSerialDTO**](CertificateSerialDTO.md) |  | [optional] 
**Owner** | Pointer to [**NestedOwnerDTO**](NestedOwnerDTO.md) |  | [optional] 

## Methods

### NewUeberCertificateDTO

`func NewUeberCertificateDTO() *UeberCertificateDTO`

NewUeberCertificateDTO instantiates a new UeberCertificateDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeberCertificateDTOWithDefaults

`func NewUeberCertificateDTOWithDefaults() *UeberCertificateDTO`

NewUeberCertificateDTOWithDefaults instantiates a new UeberCertificateDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UeberCertificateDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UeberCertificateDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UeberCertificateDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UeberCertificateDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *UeberCertificateDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UeberCertificateDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UeberCertificateDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *UeberCertificateDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *UeberCertificateDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UeberCertificateDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UeberCertificateDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UeberCertificateDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *UeberCertificateDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UeberCertificateDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UeberCertificateDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UeberCertificateDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCert

`func (o *UeberCertificateDTO) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *UeberCertificateDTO) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *UeberCertificateDTO) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *UeberCertificateDTO) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetSerial

`func (o *UeberCertificateDTO) GetSerial() CertificateSerialDTO`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *UeberCertificateDTO) GetSerialOk() (*CertificateSerialDTO, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *UeberCertificateDTO) SetSerial(v CertificateSerialDTO)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *UeberCertificateDTO) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetOwner

`func (o *UeberCertificateDTO) GetOwner() NestedOwnerDTO`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UeberCertificateDTO) GetOwnerOk() (*NestedOwnerDTO, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UeberCertificateDTO) SetOwner(v NestedOwnerDTO)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UeberCertificateDTO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


