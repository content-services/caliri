# \OwnerProductAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContentToProduct**](OwnerProductAPI.md#AddContentToProduct) | **Post** /owners/{owner_key}/products/{product_id}/content/{content_id} | 
[**AddContentsToProduct**](OwnerProductAPI.md#AddContentsToProduct) | **Post** /owners/{owner_key}/products/{product_id}/batch_content | 
[**CreateProduct**](OwnerProductAPI.md#CreateProduct) | **Post** /owners/{owner_key}/products | 
[**GetProductById**](OwnerProductAPI.md#GetProductById) | **Get** /owners/{owner_key}/products/{product_id} | 
[**GetProductCertificateById**](OwnerProductAPI.md#GetProductCertificateById) | **Get** /owners/{owner_key}/products/{product_id}/certificate | 
[**GetProductsByOwner**](OwnerProductAPI.md#GetProductsByOwner) | **Get** /owners/{owner_key}/products | 
[**RefreshPoolsForProduct**](OwnerProductAPI.md#RefreshPoolsForProduct) | **Put** /owners/{owner_key}/products/{product_id}/subscriptions | 
[**RemoveContentFromProduct**](OwnerProductAPI.md#RemoveContentFromProduct) | **Delete** /owners/{owner_key}/products/{product_id}/content/{content_id} | 
[**RemoveContentsFromProduct**](OwnerProductAPI.md#RemoveContentsFromProduct) | **Delete** /owners/{owner_key}/products/{product_id}/batch_content | 
[**RemoveProduct**](OwnerProductAPI.md#RemoveProduct) | **Delete** /owners/{owner_key}/products/{product_id} | 
[**UpdateProduct**](OwnerProductAPI.md#UpdateProduct) | **Put** /owners/{owner_key}/products/{product_id} | 



## AddContentToProduct

> ProductDTO AddContentToProduct(ctx, ownerKey, productId, contentId).Enabled(enabled).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	productId := "productId_example" // string | Product ID
	contentId := "contentId_example" // string | Content ID
	enabled := true // bool | Content enabled flag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerProductAPI.AddContentToProduct(context.Background(), ownerKey, productId, contentId).Enabled(enabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerProductAPI.AddContentToProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddContentToProduct`: ProductDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerProductAPI.AddContentToProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 
**productId** | **string** | Product ID | 
**contentId** | **string** | Content ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddContentToProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enabled** | **bool** | Content enabled flag | 

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


## AddContentsToProduct

> ProductDTO AddContentsToProduct(ctx, ownerKey, productId).RequestBody(requestBody).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	productId := "productId_example" // string | Product ID
	requestBody := map[string]bool{"key": false} // map[string]bool | Content map

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerProductAPI.AddContentsToProduct(context.Background(), ownerKey, productId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerProductAPI.AddContentsToProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddContentsToProduct`: ProductDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerProductAPI.AddContentsToProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddContentsToProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]bool** | Content map | 

### Return type

[**ProductDTO**](ProductDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProduct

> ProductDTO CreateProduct(ctx, ownerKey).ProductDTO(productDTO).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	productDTO := *openapiclient.NewProductDTO() // ProductDTO | A product to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerProductAPI.CreateProduct(context.Background(), ownerKey).ProductDTO(productDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerProductAPI.CreateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProduct`: ProductDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerProductAPI.CreateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productDTO** | [**ProductDTO**](ProductDTO.md) | A product to be created | 

### Return type

[**ProductDTO**](ProductDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductById

> ProductDTO GetProductById(ctx, ownerKey, productId).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	productId := "productId_example" // string | Product ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerProductAPI.GetProductById(context.Background(), ownerKey, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerProductAPI.GetProductById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductById`: ProductDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerProductAPI.GetProductById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductByIdRequest struct via the builder pattern


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


## GetProductCertificateById

> ProductCertificateDTO GetProductCertificateById(ctx, ownerKey, productId).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	productId := "productId_example" // string | Product ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerProductAPI.GetProductCertificateById(context.Background(), ownerKey, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerProductAPI.GetProductCertificateById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductCertificateById`: ProductCertificateDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerProductAPI.GetProductCertificateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductCertificateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProductCertificateDTO**](ProductCertificateDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductsByOwner

> []ProductDTO GetProductsByOwner(ctx, ownerKey).Product(product).OmitGlobal(omitGlobal).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	product := []string{"Inner_example"} // []string | The ID of a product to fetch. If defined, the list of products returned by this method will only include those matching the given ID. May be specified multiple times to filter on multiple product IDs.  (optional)
	omitGlobal := true // bool | whether or not to limit the lookup to only products defined within the organization's namespace, excluding any globally defined products  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerProductAPI.GetProductsByOwner(context.Background(), ownerKey).Product(product).OmitGlobal(omitGlobal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerProductAPI.GetProductsByOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductsByOwner`: []ProductDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerProductAPI.GetProductsByOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsByOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | **[]string** | The ID of a product to fetch. If defined, the list of products returned by this method will only include those matching the given ID. May be specified multiple times to filter on multiple product IDs.  | 
 **omitGlobal** | **bool** | whether or not to limit the lookup to only products defined within the organization&#39;s namespace, excluding any globally defined products  | [default to false]

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


## RefreshPoolsForProduct

> AsyncJobStatusDTO RefreshPoolsForProduct(ctx, ownerKey, productId).LazyRegen(lazyRegen).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	productId := "productId_example" // string | Product ID
	lazyRegen := true // bool | Regenerate certificates immediatelly or allow them to be regenerated on demand (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerProductAPI.RefreshPoolsForProduct(context.Background(), ownerKey, productId).LazyRegen(lazyRegen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerProductAPI.RefreshPoolsForProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshPoolsForProduct`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerProductAPI.RefreshPoolsForProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshPoolsForProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lazyRegen** | **bool** | Regenerate certificates immediatelly or allow them to be regenerated on demand | [default to true]

### Return type

[**AsyncJobStatusDTO**](AsyncJobStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveContentFromProduct

> ProductDTO RemoveContentFromProduct(ctx, ownerKey, productId, contentId).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	productId := "productId_example" // string | Product ID
	contentId := "contentId_example" // string | Content ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerProductAPI.RemoveContentFromProduct(context.Background(), ownerKey, productId, contentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerProductAPI.RemoveContentFromProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveContentFromProduct`: ProductDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerProductAPI.RemoveContentFromProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 
**productId** | **string** | Product ID | 
**contentId** | **string** | Content ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContentFromProductRequest struct via the builder pattern


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


## RemoveContentsFromProduct

> ProductDTO RemoveContentsFromProduct(ctx, ownerKey, productId).RequestBody(requestBody).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	productId := "productId_example" // string | Product ID
	requestBody := []string{"Property_example"} // []string | Content IDs

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerProductAPI.RemoveContentsFromProduct(context.Background(), ownerKey, productId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerProductAPI.RemoveContentsFromProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveContentsFromProduct`: ProductDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerProductAPI.RemoveContentsFromProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContentsFromProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]string** | Content IDs | 

### Return type

[**ProductDTO**](ProductDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProduct

> RemoveProduct(ctx, ownerKey, productId).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	productId := "productId_example" // string | Product ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OwnerProductAPI.RemoveProduct(context.Background(), ownerKey, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerProductAPI.RemoveProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProduct

> ProductDTO UpdateProduct(ctx, ownerKey, productId).ProductDTO(productDTO).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	productId := "productId_example" // string | Product ID
	productDTO := *openapiclient.NewProductDTO() // ProductDTO | Product to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerProductAPI.UpdateProduct(context.Background(), ownerKey, productId).ProductDTO(productDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerProductAPI.UpdateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProduct`: ProductDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerProductAPI.UpdateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 
**productId** | **string** | Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productDTO** | [**ProductDTO**](ProductDTO.md) | Product to be updated | 

### Return type

[**ProductDTO**](ProductDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

