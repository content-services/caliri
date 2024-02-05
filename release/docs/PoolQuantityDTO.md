# PoolQuantityDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **int32** |  | [optional] 
**Pool** | Pointer to [**PoolDTO**](PoolDTO.md) |  | [optional] 

## Methods

### NewPoolQuantityDTO

`func NewPoolQuantityDTO() *PoolQuantityDTO`

NewPoolQuantityDTO instantiates a new PoolQuantityDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolQuantityDTOWithDefaults

`func NewPoolQuantityDTOWithDefaults() *PoolQuantityDTO`

NewPoolQuantityDTOWithDefaults instantiates a new PoolQuantityDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *PoolQuantityDTO) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PoolQuantityDTO) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PoolQuantityDTO) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PoolQuantityDTO) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPool

`func (o *PoolQuantityDTO) GetPool() PoolDTO`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *PoolQuantityDTO) GetPoolOk() (*PoolDTO, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *PoolQuantityDTO) SetPool(v PoolDTO)`

SetPool sets Pool field to given value.

### HasPool

`func (o *PoolQuantityDTO) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


