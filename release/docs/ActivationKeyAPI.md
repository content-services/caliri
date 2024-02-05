# \ActivationKeyAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddActivationKeyContentOverrides**](ActivationKeyAPI.md#AddActivationKeyContentOverrides) | **Put** /activation_keys/{activation_key_id}/content_overrides | 
[**AddPoolToKey**](ActivationKeyAPI.md#AddPoolToKey) | **Post** /activation_keys/{activation_key_id}/pools/{pool_id} | 
[**AddProductIdToKey**](ActivationKeyAPI.md#AddProductIdToKey) | **Post** /activation_keys/{activation_key_id}/product/{product_id} | 
[**DeleteActivationKey**](ActivationKeyAPI.md#DeleteActivationKey) | **Delete** /activation_keys/{activation_key_id} | 
[**DeleteActivationKeyContentOverrides**](ActivationKeyAPI.md#DeleteActivationKeyContentOverrides) | **Delete** /activation_keys/{activation_key_id}/content_overrides | 
[**FindActivationKey**](ActivationKeyAPI.md#FindActivationKey) | **Get** /activation_keys | 
[**GetActivationKey**](ActivationKeyAPI.md#GetActivationKey) | **Get** /activation_keys/{activation_key_id} | 
[**GetActivationKeyPools**](ActivationKeyAPI.md#GetActivationKeyPools) | **Get** /activation_keys/{activation_key_id}/pools | 
[**ListActivationKeyContentOverrides**](ActivationKeyAPI.md#ListActivationKeyContentOverrides) | **Get** /activation_keys/{activation_key_id}/content_overrides | 
[**RemovePoolFromKey**](ActivationKeyAPI.md#RemovePoolFromKey) | **Delete** /activation_keys/{activation_key_id}/pools/{pool_id} | 
[**RemoveProductIdFromKey**](ActivationKeyAPI.md#RemoveProductIdFromKey) | **Delete** /activation_keys/{activation_key_id}/product/{product_id} | 
[**UpdateActivationKey**](ActivationKeyAPI.md#UpdateActivationKey) | **Put** /activation_keys/{activation_key_id} | 



## AddActivationKeyContentOverrides

> []ContentOverrideDTO AddActivationKeyContentOverrides(ctx, activationKeyId).ContentOverrideDTO(contentOverrideDTO).Execute()





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
	activationKeyId := "activationKeyId_example" // string | The ID of the activation key
	contentOverrideDTO := []openapiclient.ContentOverrideDTO{*openapiclient.NewContentOverrideDTO()} // []ContentOverrideDTO | The list of the content overrides

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationKeyAPI.AddActivationKeyContentOverrides(context.Background(), activationKeyId).ContentOverrideDTO(contentOverrideDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.AddActivationKeyContentOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddActivationKeyContentOverrides`: []ContentOverrideDTO
	fmt.Fprintf(os.Stdout, "Response from `ActivationKeyAPI.AddActivationKeyContentOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationKeyId** | **string** | The ID of the activation key | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddActivationKeyContentOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentOverrideDTO** | [**[]ContentOverrideDTO**](ContentOverrideDTO.md) | The list of the content overrides | 

### Return type

[**[]ContentOverrideDTO**](ContentOverrideDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPoolToKey

> ActivationKeyDTO AddPoolToKey(ctx, activationKeyId, poolId).Quantity(quantity).Execute()





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
	activationKeyId := "activationKeyId_example" // string | ID of the activation key
	poolId := "poolId_example" // string | ID of the pool
	quantity := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationKeyAPI.AddPoolToKey(context.Background(), activationKeyId, poolId).Quantity(quantity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.AddPoolToKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPoolToKey`: ActivationKeyDTO
	fmt.Fprintf(os.Stdout, "Response from `ActivationKeyAPI.AddPoolToKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationKeyId** | **string** | ID of the activation key | 
**poolId** | **string** | ID of the pool | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPoolToKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **quantity** | **int64** |  | 

### Return type

[**ActivationKeyDTO**](ActivationKeyDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddProductIdToKey

> ActivationKeyDTO AddProductIdToKey(ctx, activationKeyId, productId).Execute()





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
	activationKeyId := "activationKeyId_example" // string | ID of the activation key
	productId := "productId_example" // string | ID of the product

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationKeyAPI.AddProductIdToKey(context.Background(), activationKeyId, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.AddProductIdToKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddProductIdToKey`: ActivationKeyDTO
	fmt.Fprintf(os.Stdout, "Response from `ActivationKeyAPI.AddProductIdToKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationKeyId** | **string** | ID of the activation key | 
**productId** | **string** | ID of the product | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProductIdToKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActivationKeyDTO**](ActivationKeyDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteActivationKey

> DeleteActivationKey(ctx, activationKeyId).Execute()





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
	activationKeyId := "activationKeyId_example" // string | The ID of the activation key to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActivationKeyAPI.DeleteActivationKey(context.Background(), activationKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.DeleteActivationKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationKeyId** | **string** | The ID of the activation key to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteActivationKeyRequest struct via the builder pattern


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


## DeleteActivationKeyContentOverrides

> []ContentOverrideDTO DeleteActivationKeyContentOverrides(ctx, activationKeyId).ContentOverrideDTO(contentOverrideDTO).Execute()





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
	activationKeyId := "activationKeyId_example" // string | The ID of the activation key
	contentOverrideDTO := []openapiclient.ContentOverrideDTO{*openapiclient.NewContentOverrideDTO()} // []ContentOverrideDTO | The list of the content overrides

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationKeyAPI.DeleteActivationKeyContentOverrides(context.Background(), activationKeyId).ContentOverrideDTO(contentOverrideDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.DeleteActivationKeyContentOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteActivationKeyContentOverrides`: []ContentOverrideDTO
	fmt.Fprintf(os.Stdout, "Response from `ActivationKeyAPI.DeleteActivationKeyContentOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationKeyId** | **string** | The ID of the activation key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteActivationKeyContentOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentOverrideDTO** | [**[]ContentOverrideDTO**](ContentOverrideDTO.md) | The list of the content overrides | 

### Return type

[**[]ContentOverrideDTO**](ContentOverrideDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindActivationKey

> []ActivationKeyDTO FindActivationKey(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationKeyAPI.FindActivationKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.FindActivationKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindActivationKey`: []ActivationKeyDTO
	fmt.Fprintf(os.Stdout, "Response from `ActivationKeyAPI.FindActivationKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindActivationKeyRequest struct via the builder pattern


### Return type

[**[]ActivationKeyDTO**](ActivationKeyDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivationKey

> ActivationKeyDTO GetActivationKey(ctx, activationKeyId).Execute()





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
	activationKeyId := "activationKeyId_example" // string | The ID of the activation key to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationKeyAPI.GetActivationKey(context.Background(), activationKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.GetActivationKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivationKey`: ActivationKeyDTO
	fmt.Fprintf(os.Stdout, "Response from `ActivationKeyAPI.GetActivationKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationKeyId** | **string** | The ID of the activation key to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivationKeyDTO**](ActivationKeyDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivationKeyPools

> []PoolDTO GetActivationKeyPools(ctx, activationKeyId).Execute()





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
	activationKeyId := "activationKeyId_example" // string | The ID of the activation key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationKeyAPI.GetActivationKeyPools(context.Background(), activationKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.GetActivationKeyPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivationKeyPools`: []PoolDTO
	fmt.Fprintf(os.Stdout, "Response from `ActivationKeyAPI.GetActivationKeyPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationKeyId** | **string** | The ID of the activation key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivationKeyPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PoolDTO**](PoolDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActivationKeyContentOverrides

> []ContentOverrideDTO ListActivationKeyContentOverrides(ctx, activationKeyId).Execute()





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
	activationKeyId := "activationKeyId_example" // string | The ID of the activation key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationKeyAPI.ListActivationKeyContentOverrides(context.Background(), activationKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.ListActivationKeyContentOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActivationKeyContentOverrides`: []ContentOverrideDTO
	fmt.Fprintf(os.Stdout, "Response from `ActivationKeyAPI.ListActivationKeyContentOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationKeyId** | **string** | The ID of the activation key | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActivationKeyContentOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ContentOverrideDTO**](ContentOverrideDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePoolFromKey

> ActivationKeyDTO RemovePoolFromKey(ctx, activationKeyId, poolId).Execute()





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
	activationKeyId := "activationKeyId_example" // string | The ID of the activation key
	poolId := "poolId_example" // string | The ID of a pool to be removed

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationKeyAPI.RemovePoolFromKey(context.Background(), activationKeyId, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.RemovePoolFromKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemovePoolFromKey`: ActivationKeyDTO
	fmt.Fprintf(os.Stdout, "Response from `ActivationKeyAPI.RemovePoolFromKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationKeyId** | **string** | The ID of the activation key | 
**poolId** | **string** | The ID of a pool to be removed | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePoolFromKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActivationKeyDTO**](ActivationKeyDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProductIdFromKey

> ActivationKeyDTO RemoveProductIdFromKey(ctx, activationKeyId, productId).Execute()





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
	activationKeyId := "activationKeyId_example" // string | The ID of the activation key
	productId := "productId_example" // string | The ID of the product

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationKeyAPI.RemoveProductIdFromKey(context.Background(), activationKeyId, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.RemoveProductIdFromKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveProductIdFromKey`: ActivationKeyDTO
	fmt.Fprintf(os.Stdout, "Response from `ActivationKeyAPI.RemoveProductIdFromKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationKeyId** | **string** | The ID of the activation key | 
**productId** | **string** | The ID of the product | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProductIdFromKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ActivationKeyDTO**](ActivationKeyDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActivationKey

> ActivationKeyDTO UpdateActivationKey(ctx, activationKeyId).ActivationKeyDTO(activationKeyDTO).Execute()





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
	activationKeyId := "activationKeyId_example" // string | The ID of the activation key to be updated
	activationKeyDTO := *openapiclient.NewActivationKeyDTO() // ActivationKeyDTO | Activation key to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationKeyAPI.UpdateActivationKey(context.Background(), activationKeyId).ActivationKeyDTO(activationKeyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationKeyAPI.UpdateActivationKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateActivationKey`: ActivationKeyDTO
	fmt.Fprintf(os.Stdout, "Response from `ActivationKeyAPI.UpdateActivationKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationKeyId** | **string** | The ID of the activation key to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActivationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activationKeyDTO** | [**ActivationKeyDTO**](ActivationKeyDTO.md) | Activation key to be updated | 

### Return type

[**ActivationKeyDTO**](ActivationKeyDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

