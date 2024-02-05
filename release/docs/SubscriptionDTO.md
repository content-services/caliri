# SubscriptionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NestedOwnerDTO**](NestedOwnerDTO.md) |  | [optional] 
**Product** | Pointer to [**ProductDTO**](ProductDTO.md) |  | [optional] 
**DerivedProduct** | Pointer to [**ProductDTO**](ProductDTO.md) |  | [optional] 
**ProvidedProducts** | Pointer to [**[]ProductDTO**](ProductDTO.md) |  | [optional] 
**DerivedProvidedProducts** | Pointer to [**[]ProductDTO**](ProductDTO.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**ContractNumber** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**OrderNumber** | Pointer to **string** |  | [optional] 
**UpstreamPoolId** | Pointer to **string** |  | [optional] 
**UpstreamEntitlementId** | Pointer to **string** |  | [optional] 
**UpstreamConsumerId** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to [**CertificateDTO**](CertificateDTO.md) |  | [optional] 
**Cdn** | Pointer to [**CdnDTO**](CdnDTO.md) |  | [optional] 
**Stacked** | Pointer to **bool** |  | [optional] 
**StackId** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionDTO

`func NewSubscriptionDTO() *SubscriptionDTO`

NewSubscriptionDTO instantiates a new SubscriptionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDTOWithDefaults

`func NewSubscriptionDTOWithDefaults() *SubscriptionDTO`

NewSubscriptionDTOWithDefaults instantiates a new SubscriptionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SubscriptionDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubscriptionDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubscriptionDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SubscriptionDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *SubscriptionDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SubscriptionDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SubscriptionDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SubscriptionDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *SubscriptionDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwner

`func (o *SubscriptionDTO) GetOwner() NestedOwnerDTO`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SubscriptionDTO) GetOwnerOk() (*NestedOwnerDTO, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SubscriptionDTO) SetOwner(v NestedOwnerDTO)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SubscriptionDTO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProduct

`func (o *SubscriptionDTO) GetProduct() ProductDTO`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *SubscriptionDTO) GetProductOk() (*ProductDTO, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *SubscriptionDTO) SetProduct(v ProductDTO)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *SubscriptionDTO) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetDerivedProduct

`func (o *SubscriptionDTO) GetDerivedProduct() ProductDTO`

GetDerivedProduct returns the DerivedProduct field if non-nil, zero value otherwise.

### GetDerivedProductOk

`func (o *SubscriptionDTO) GetDerivedProductOk() (*ProductDTO, bool)`

GetDerivedProductOk returns a tuple with the DerivedProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedProduct

`func (o *SubscriptionDTO) SetDerivedProduct(v ProductDTO)`

SetDerivedProduct sets DerivedProduct field to given value.

### HasDerivedProduct

`func (o *SubscriptionDTO) HasDerivedProduct() bool`

HasDerivedProduct returns a boolean if a field has been set.

### GetProvidedProducts

`func (o *SubscriptionDTO) GetProvidedProducts() []ProductDTO`

GetProvidedProducts returns the ProvidedProducts field if non-nil, zero value otherwise.

### GetProvidedProductsOk

`func (o *SubscriptionDTO) GetProvidedProductsOk() (*[]ProductDTO, bool)`

GetProvidedProductsOk returns a tuple with the ProvidedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidedProducts

`func (o *SubscriptionDTO) SetProvidedProducts(v []ProductDTO)`

SetProvidedProducts sets ProvidedProducts field to given value.

### HasProvidedProducts

`func (o *SubscriptionDTO) HasProvidedProducts() bool`

HasProvidedProducts returns a boolean if a field has been set.

### GetDerivedProvidedProducts

`func (o *SubscriptionDTO) GetDerivedProvidedProducts() []ProductDTO`

GetDerivedProvidedProducts returns the DerivedProvidedProducts field if non-nil, zero value otherwise.

### GetDerivedProvidedProductsOk

`func (o *SubscriptionDTO) GetDerivedProvidedProductsOk() (*[]ProductDTO, bool)`

GetDerivedProvidedProductsOk returns a tuple with the DerivedProvidedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedProvidedProducts

`func (o *SubscriptionDTO) SetDerivedProvidedProducts(v []ProductDTO)`

SetDerivedProvidedProducts sets DerivedProvidedProducts field to given value.

### HasDerivedProvidedProducts

`func (o *SubscriptionDTO) HasDerivedProvidedProducts() bool`

HasDerivedProvidedProducts returns a boolean if a field has been set.

### GetQuantity

`func (o *SubscriptionDTO) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionDTO) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionDTO) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SubscriptionDTO) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartDate

`func (o *SubscriptionDTO) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SubscriptionDTO) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SubscriptionDTO) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SubscriptionDTO) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *SubscriptionDTO) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SubscriptionDTO) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SubscriptionDTO) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SubscriptionDTO) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetContractNumber

`func (o *SubscriptionDTO) GetContractNumber() string`

GetContractNumber returns the ContractNumber field if non-nil, zero value otherwise.

### GetContractNumberOk

`func (o *SubscriptionDTO) GetContractNumberOk() (*string, bool)`

GetContractNumberOk returns a tuple with the ContractNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractNumber

`func (o *SubscriptionDTO) SetContractNumber(v string)`

SetContractNumber sets ContractNumber field to given value.

### HasContractNumber

`func (o *SubscriptionDTO) HasContractNumber() bool`

HasContractNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *SubscriptionDTO) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *SubscriptionDTO) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *SubscriptionDTO) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *SubscriptionDTO) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetModified

`func (o *SubscriptionDTO) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *SubscriptionDTO) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *SubscriptionDTO) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *SubscriptionDTO) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetLastModified

`func (o *SubscriptionDTO) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *SubscriptionDTO) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *SubscriptionDTO) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *SubscriptionDTO) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetOrderNumber

`func (o *SubscriptionDTO) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *SubscriptionDTO) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *SubscriptionDTO) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *SubscriptionDTO) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetUpstreamPoolId

`func (o *SubscriptionDTO) GetUpstreamPoolId() string`

GetUpstreamPoolId returns the UpstreamPoolId field if non-nil, zero value otherwise.

### GetUpstreamPoolIdOk

`func (o *SubscriptionDTO) GetUpstreamPoolIdOk() (*string, bool)`

GetUpstreamPoolIdOk returns a tuple with the UpstreamPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamPoolId

`func (o *SubscriptionDTO) SetUpstreamPoolId(v string)`

SetUpstreamPoolId sets UpstreamPoolId field to given value.

### HasUpstreamPoolId

`func (o *SubscriptionDTO) HasUpstreamPoolId() bool`

HasUpstreamPoolId returns a boolean if a field has been set.

### GetUpstreamEntitlementId

`func (o *SubscriptionDTO) GetUpstreamEntitlementId() string`

GetUpstreamEntitlementId returns the UpstreamEntitlementId field if non-nil, zero value otherwise.

### GetUpstreamEntitlementIdOk

`func (o *SubscriptionDTO) GetUpstreamEntitlementIdOk() (*string, bool)`

GetUpstreamEntitlementIdOk returns a tuple with the UpstreamEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamEntitlementId

`func (o *SubscriptionDTO) SetUpstreamEntitlementId(v string)`

SetUpstreamEntitlementId sets UpstreamEntitlementId field to given value.

### HasUpstreamEntitlementId

`func (o *SubscriptionDTO) HasUpstreamEntitlementId() bool`

HasUpstreamEntitlementId returns a boolean if a field has been set.

### GetUpstreamConsumerId

`func (o *SubscriptionDTO) GetUpstreamConsumerId() string`

GetUpstreamConsumerId returns the UpstreamConsumerId field if non-nil, zero value otherwise.

### GetUpstreamConsumerIdOk

`func (o *SubscriptionDTO) GetUpstreamConsumerIdOk() (*string, bool)`

GetUpstreamConsumerIdOk returns a tuple with the UpstreamConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamConsumerId

`func (o *SubscriptionDTO) SetUpstreamConsumerId(v string)`

SetUpstreamConsumerId sets UpstreamConsumerId field to given value.

### HasUpstreamConsumerId

`func (o *SubscriptionDTO) HasUpstreamConsumerId() bool`

HasUpstreamConsumerId returns a boolean if a field has been set.

### GetCertificate

`func (o *SubscriptionDTO) GetCertificate() CertificateDTO`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SubscriptionDTO) GetCertificateOk() (*CertificateDTO, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SubscriptionDTO) SetCertificate(v CertificateDTO)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SubscriptionDTO) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCdn

`func (o *SubscriptionDTO) GetCdn() CdnDTO`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *SubscriptionDTO) GetCdnOk() (*CdnDTO, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *SubscriptionDTO) SetCdn(v CdnDTO)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *SubscriptionDTO) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### GetStacked

`func (o *SubscriptionDTO) GetStacked() bool`

GetStacked returns the Stacked field if non-nil, zero value otherwise.

### GetStackedOk

`func (o *SubscriptionDTO) GetStackedOk() (*bool, bool)`

GetStackedOk returns a tuple with the Stacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacked

`func (o *SubscriptionDTO) SetStacked(v bool)`

SetStacked sets Stacked field to given value.

### HasStacked

`func (o *SubscriptionDTO) HasStacked() bool`

HasStacked returns a boolean if a field has been set.

### GetStackId

`func (o *SubscriptionDTO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *SubscriptionDTO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *SubscriptionDTO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *SubscriptionDTO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


