# \ProductsAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductByUuid**](ProductsAPI.md#GetProductByUuid) | **Get** /products/{product_uuid} | 
[**GetProducts**](ProductsAPI.md#GetProducts) | **Get** /products | 
[**RefreshPoolsForProducts**](ProductsAPI.md#RefreshPoolsForProducts) | **Put** /products/subscriptions | 



## GetProductByUuid

> ProductDTO GetProductByUuid(ctx, productUuid).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/caliri/release/v4"
)

func main() {
	productUuid := "productUuid_example" // string | The product UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.GetProductByUuid(context.Background(), productUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetProductByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductByUuid`: ProductDTO
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetProductByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productUuid** | **string** | The product UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductDTO**](ProductDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProducts

> []ProductDTO GetProducts(ctx).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/caliri/release/v4"
)

func main() {
	page := int32(2) // int32 | Page index to return (optional)
	perPage := int32(10) // int32 | Number of items to return per page (optional)
	order := "asc" // string | Direction of ordering (optional)
	sortBy := "name" // string | Property to use for ordering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.GetProducts(context.Background()).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProducts`: []ProductDTO
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page index to return | 
 **perPage** | **int32** | Number of items to return per page | 
 **order** | **string** | Direction of ordering | 
 **sortBy** | **string** | Property to use for ordering | 

### Return type

[**[]ProductDTO**](ProductDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshPoolsForProducts

> []AsyncJobStatusDTO RefreshPoolsForProducts(ctx).Product(product).LazyRegen(lazyRegen).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/content-services/caliri/release/v4"
)

func main() {
	product := []string{"Inner_example"} // []string | Multiple product Ids
	lazyRegen := true // bool | Regenerate certificates immediatelly or allow them to be regenerated on demand (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.RefreshPoolsForProducts(context.Background()).Product(product).LazyRegen(lazyRegen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.RefreshPoolsForProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshPoolsForProducts`: []AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.RefreshPoolsForProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshPoolsForProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **[]string** | Multiple product Ids | 
 **lazyRegen** | **bool** | Regenerate certificates immediatelly or allow them to be regenerated on demand | [default to true]

### Return type

[**[]AsyncJobStatusDTO**](AsyncJobStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

