# ProductDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Multiplier** | Pointer to **int64** |  | [optional] 
**Attributes** | Pointer to [**[]AttributeDTO**](AttributeDTO.md) |  | [optional] 
**ProductContent** | Pointer to [**[]ProductContentDTO**](ProductContentDTO.md) |  | [optional] 
**DependentProductIds** | Pointer to **[]string** |  | [optional] 
**Branding** | Pointer to [**[]BrandingDTO**](BrandingDTO.md) |  | [optional] 
**DerivedProduct** | Pointer to [**ProductDTO**](ProductDTO.md) |  | [optional] 
**ProvidedProducts** | Pointer to [**[]ProductDTO**](ProductDTO.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewProductDTO

`func NewProductDTO() *ProductDTO`

NewProductDTO instantiates a new ProductDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDTOWithDefaults

`func NewProductDTOWithDefaults() *ProductDTO`

NewProductDTOWithDefaults instantiates a new ProductDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ProductDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ProductDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ProductDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ProductDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *ProductDTO) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ProductDTO) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ProductDTO) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ProductDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetId

`func (o *ProductDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProductDTO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProductDTO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProductDTO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProductDTO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *ProductDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMultiplier

`func (o *ProductDTO) GetMultiplier() int64`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *ProductDTO) GetMultiplierOk() (*int64, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *ProductDTO) SetMultiplier(v int64)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *ProductDTO) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetAttributes

`func (o *ProductDTO) GetAttributes() []AttributeDTO`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProductDTO) GetAttributesOk() (*[]AttributeDTO, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProductDTO) SetAttributes(v []AttributeDTO)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ProductDTO) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetProductContent

`func (o *ProductDTO) GetProductContent() []ProductContentDTO`

GetProductContent returns the ProductContent field if non-nil, zero value otherwise.

### GetProductContentOk

`func (o *ProductDTO) GetProductContentOk() (*[]ProductContentDTO, bool)`

GetProductContentOk returns a tuple with the ProductContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductContent

`func (o *ProductDTO) SetProductContent(v []ProductContentDTO)`

SetProductContent sets ProductContent field to given value.

### HasProductContent

`func (o *ProductDTO) HasProductContent() bool`

HasProductContent returns a boolean if a field has been set.

### GetDependentProductIds

`func (o *ProductDTO) GetDependentProductIds() []string`

GetDependentProductIds returns the DependentProductIds field if non-nil, zero value otherwise.

### GetDependentProductIdsOk

`func (o *ProductDTO) GetDependentProductIdsOk() (*[]string, bool)`

GetDependentProductIdsOk returns a tuple with the DependentProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentProductIds

`func (o *ProductDTO) SetDependentProductIds(v []string)`

SetDependentProductIds sets DependentProductIds field to given value.

### HasDependentProductIds

`func (o *ProductDTO) HasDependentProductIds() bool`

HasDependentProductIds returns a boolean if a field has been set.

### GetBranding

`func (o *ProductDTO) GetBranding() []BrandingDTO`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *ProductDTO) GetBrandingOk() (*[]BrandingDTO, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *ProductDTO) SetBranding(v []BrandingDTO)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *ProductDTO) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetDerivedProduct

`func (o *ProductDTO) GetDerivedProduct() ProductDTO`

GetDerivedProduct returns the DerivedProduct field if non-nil, zero value otherwise.

### GetDerivedProductOk

`func (o *ProductDTO) GetDerivedProductOk() (*ProductDTO, bool)`

GetDerivedProductOk returns a tuple with the DerivedProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedProduct

`func (o *ProductDTO) SetDerivedProduct(v ProductDTO)`

SetDerivedProduct sets DerivedProduct field to given value.

### HasDerivedProduct

`func (o *ProductDTO) HasDerivedProduct() bool`

HasDerivedProduct returns a boolean if a field has been set.

### GetProvidedProducts

`func (o *ProductDTO) GetProvidedProducts() []ProductDTO`

GetProvidedProducts returns the ProvidedProducts field if non-nil, zero value otherwise.

### GetProvidedProductsOk

`func (o *ProductDTO) GetProvidedProductsOk() (*[]ProductDTO, bool)`

GetProvidedProductsOk returns a tuple with the ProvidedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvidedProducts

`func (o *ProductDTO) SetProvidedProducts(v []ProductDTO)`

SetProvidedProducts sets ProvidedProducts field to given value.

### HasProvidedProducts

`func (o *ProductDTO) HasProvidedProducts() bool`

HasProvidedProducts returns a boolean if a field has been set.

### GetHref

`func (o *ProductDTO) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ProductDTO) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ProductDTO) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ProductDTO) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


