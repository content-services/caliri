# ProvidedProductDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** |  | 
**ProductName** | Pointer to **string** |  | [optional] 

## Methods

### NewProvidedProductDTO

`func NewProvidedProductDTO(productId string, ) *ProvidedProductDTO`

NewProvidedProductDTO instantiates a new ProvidedProductDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvidedProductDTOWithDefaults

`func NewProvidedProductDTOWithDefaults() *ProvidedProductDTO`

NewProvidedProductDTOWithDefaults instantiates a new ProvidedProductDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProvidedProductDTO) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProvidedProductDTO) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProvidedProductDTO) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetProductName

`func (o *ProvidedProductDTO) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ProvidedProductDTO) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ProvidedProductDTO) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ProvidedProductDTO) HasProductName() bool`

HasProductName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


