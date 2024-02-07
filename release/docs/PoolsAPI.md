# \PoolsAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePool**](PoolsAPI.md#DeletePool) | **Delete** /pools/{pool_id} | 
[**GetPool**](PoolsAPI.md#GetPool) | **Get** /pools/{pool_id} | 
[**GetPoolCdn**](PoolsAPI.md#GetPoolCdn) | **Get** /pools/{pool_id}/cdn | 
[**GetPoolEntitlements**](PoolsAPI.md#GetPoolEntitlements) | **Get** /pools/{pool_id}/entitlements | 
[**GetSubCert**](PoolsAPI.md#GetSubCert) | **Get** /pools/{pool_id}/cert | 
[**ListEntitledConsumerUuids**](PoolsAPI.md#ListEntitledConsumerUuids) | **Get** /pools/{pool_id}/entitlements/consumer_uuids | 
[**ListPools**](PoolsAPI.md#ListPools) | **Get** /pools | 



## DeletePool

> DeletePool(ctx, poolId).Execute()





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
	poolId := "poolId_example" // string | Pool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoolsAPI.DeletePool(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.DeletePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPool

> PoolDTO GetPool(ctx, poolId).Consumer(consumer).Activeon(activeon).Execute()





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
	poolId := "poolId_example" // string | Pool ID
	consumer := "consumer_example" // string | Consumer UUID (optional)
	activeon := "activeon_example" // string | Uses ISO 8601 format (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPool(context.Background(), poolId).Consumer(consumer).Activeon(activeon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPool`: PoolDTO
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumer** | **string** | Consumer UUID | 
 **activeon** | **string** | Uses ISO 8601 format | 

### Return type

[**PoolDTO**](PoolDTO.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolCdn

> CdnDTO GetPoolCdn(ctx, poolId).Execute()





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
	poolId := "poolId_example" // string | Pool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPoolCdn(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPoolCdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolCdn`: CdnDTO
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPoolCdn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolCdnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CdnDTO**](CdnDTO.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolEntitlements

> []EntitlementDTO GetPoolEntitlements(ctx, poolId).Execute()





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
	poolId := "poolId_example" // string | Pool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPoolEntitlements(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPoolEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPoolEntitlements`: []EntitlementDTO
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPoolEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EntitlementDTO**](EntitlementDTO.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubCert

> map[string]interface{} GetSubCert(ctx, poolId).Execute()





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
	poolId := "poolId_example" // string | Pool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetSubCert(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetSubCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubCert`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetSubCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitledConsumerUuids

> []string ListEntitledConsumerUuids(ctx, poolId).Execute()





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
	poolId := "poolId_example" // string | Pool ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.ListEntitledConsumerUuids(context.Background(), poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.ListEntitledConsumerUuids``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntitledConsumerUuids`: []string
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.ListEntitledConsumerUuids`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitledConsumerUuidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPools

> []PoolDTO ListPools(ctx).Owner(owner).Consumer(consumer).Product(product).Listall(listall).Activeon(activeon).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()





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
	owner := "owner_example" // string | Owner ID (optional)
	consumer := "consumer_example" // string | Consumer UUID (optional)
	product := "product_example" // string | Product ID (optional)
	listall := true // bool | Use with consumerUuid to list all pools available to the consumer. This will include pools which would otherwise be omitted due to a rules warning. (i.e. not recommended) Pools that trigger an error however will still be omitted. (no entitlements available, consumer type mismatch, etc)  (optional) (default to false)
	activeon := "activeon_example" // string | Uses ISO 8601 format (optional)
	page := int32(2) // int32 | Page index to return (optional)
	perPage := int32(10) // int32 | Number of items to return per page (optional)
	order := "asc" // string | Direction of ordering (optional)
	sortBy := "name" // string | Property to use for ordering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.ListPools(context.Background()).Owner(owner).Consumer(consumer).Product(product).Listall(listall).Activeon(activeon).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.ListPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPools`: []PoolDTO
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.ListPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Owner ID | 
 **consumer** | **string** | Consumer UUID | 
 **product** | **string** | Product ID | 
 **listall** | **bool** | Use with consumerUuid to list all pools available to the consumer. This will include pools which would otherwise be omitted due to a rules warning. (i.e. not recommended) Pools that trigger an error however will still be omitted. (no entitlements available, consumer type mismatch, etc)  | [default to false]
 **activeon** | **string** | Uses ISO 8601 format | 
 **page** | **int32** | Page index to return | 
 **perPage** | **int32** | Number of items to return per page | 
 **order** | **string** | Direction of ordering | 
 **sortBy** | **string** | Property to use for ordering | 

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

