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

> []ProductDTO GetProducts(ctx).Owner(owner).Product(product).Name(name).Active(active).Custom(custom).Execute()





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
	owner := []string{"Inner_example"} // []string | The key of an owner to use to limit the product search. If defined, the list of products returned by this endpoint will only include those available to the given owner. May be specified multiple times to filter by multiple owners. If multiple owners are provided, products available to any of the provided owners will be returned.  (optional)
	product := []string{"Inner_example"} // []string | The ID of a product to fetch. If defined, the list of products returned by this endpoint will only include those matching the given ID. May be specified multiple times to filter on multiple product IDs. If multiple IDs are provided, any products matching any of the provided IDs will be returned.  (optional)
	name := []string{"Inner_example"} // []string | The names of products to fetch. If defined, the list of products returned by this endpoint will only include those matching the given names (case-insensitive). May be specified multiple times to filter on multiple names. If multiple names are provided, any products matching any of the provided names will be returned.  (optional)
	active := "{"include":{"value":"include","description":"indicates the query should include matching active products, along with any inactive\nproducts\n"},"exclude":{"value":"exclude","description":"indicates the query should exclude, or omit, active products, returning only matching\ninactive products\n"},"exclusive":{"value":"exclusive","description":"indicates the query should only include matching active products, excluding any inactive\nproducts\n"}}" // string | A value indicating how 'active' products should be considered when fetching products, where 'active' is defined as a product that is currently in use by a subscription with a start time in the past and that has not yet expired, or in use by a product which itself is considered 'active.' Must be one of 'include', 'exclude', or 'exclusive' indicating that active products should be included along with inactive products, excluded (omitted) from the results, or exclusively selected as the only products to return. Defaults to 'exclusive'.  (optional) (default to "exclusive")
	custom := "{"include":{"value":"include","description":"indicates the query should include matching custom products, along with any matching\nglobally defined products\n"},"exclude":{"value":"exclude","description":"indicates the query should exclude, or omit, custom products, returning only matching\nglobally defined products\n"},"exclusive":{"value":"exclusive","description":"indicates the query should only include matching custom products, excluding any globally\ndefined products\n"}}" // string | A value indicating how custom products are considered when fetching products, where 'custom' is defined as a product that did not originate from a refresh operation nor manifest import. Must be one of 'include', 'exclude', or 'exclusive' indicating that custom products should be passively included, excluded or omitted from the output, or exclusively selected as the only products to return.  (optional) (default to "include")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.GetProducts(context.Background()).Owner(owner).Product(product).Name(name).Active(active).Custom(custom).Execute()
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
 **owner** | **[]string** | The key of an owner to use to limit the product search. If defined, the list of products returned by this endpoint will only include those available to the given owner. May be specified multiple times to filter by multiple owners. If multiple owners are provided, products available to any of the provided owners will be returned.  | 
 **product** | **[]string** | The ID of a product to fetch. If defined, the list of products returned by this endpoint will only include those matching the given ID. May be specified multiple times to filter on multiple product IDs. If multiple IDs are provided, any products matching any of the provided IDs will be returned.  | 
 **name** | **[]string** | The names of products to fetch. If defined, the list of products returned by this endpoint will only include those matching the given names (case-insensitive). May be specified multiple times to filter on multiple names. If multiple names are provided, any products matching any of the provided names will be returned.  | 
 **active** | **string** | A value indicating how &#39;active&#39; products should be considered when fetching products, where &#39;active&#39; is defined as a product that is currently in use by a subscription with a start time in the past and that has not yet expired, or in use by a product which itself is considered &#39;active.&#39; Must be one of &#39;include&#39;, &#39;exclude&#39;, or &#39;exclusive&#39; indicating that active products should be included along with inactive products, excluded (omitted) from the results, or exclusively selected as the only products to return. Defaults to &#39;exclusive&#39;.  | [default to &quot;exclusive&quot;]
 **custom** | **string** | A value indicating how custom products are considered when fetching products, where &#39;custom&#39; is defined as a product that did not originate from a refresh operation nor manifest import. Must be one of &#39;include&#39;, &#39;exclude&#39;, or &#39;exclusive&#39; indicating that custom products should be passively included, excluded or omitted from the output, or exclusively selected as the only products to return.  | [default to &quot;include&quot;]

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

