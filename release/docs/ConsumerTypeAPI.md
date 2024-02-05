# \ConsumerTypeAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConsumerType**](ConsumerTypeAPI.md#CreateConsumerType) | **Post** /consumertypes | 
[**DeleteConsumerType**](ConsumerTypeAPI.md#DeleteConsumerType) | **Delete** /consumertypes/{id} | 
[**GetConsumerType**](ConsumerTypeAPI.md#GetConsumerType) | **Get** /consumertypes/{id} | 
[**GetConsumerTypes**](ConsumerTypeAPI.md#GetConsumerTypes) | **Get** /consumertypes | 
[**UpdateConsumerType**](ConsumerTypeAPI.md#UpdateConsumerType) | **Put** /consumertypes/{id} | 



## CreateConsumerType

> ConsumerTypeDTO CreateConsumerType(ctx).ConsumerTypeDTO(consumerTypeDTO).Execute()





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
	consumerTypeDTO := *openapiclient.NewConsumerTypeDTO() // ConsumerTypeDTO | The consumer type to create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerTypeAPI.CreateConsumerType(context.Background()).ConsumerTypeDTO(consumerTypeDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerTypeAPI.CreateConsumerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConsumerType`: ConsumerTypeDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerTypeAPI.CreateConsumerType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConsumerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consumerTypeDTO** | [**ConsumerTypeDTO**](ConsumerTypeDTO.md) | The consumer type to create | 

### Return type

[**ConsumerTypeDTO**](ConsumerTypeDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConsumerType

> DeleteConsumerType(ctx, id).Execute()





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
	id := "id_example" // string | The ID of the consumer type to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumerTypeAPI.DeleteConsumerType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerTypeAPI.DeleteConsumerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the consumer type to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsumerTypeRequest struct via the builder pattern


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


## GetConsumerType

> ConsumerTypeDTO GetConsumerType(ctx, id).Execute()





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
	id := "id_example" // string | The ID of the consumer type to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerTypeAPI.GetConsumerType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerTypeAPI.GetConsumerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConsumerType`: ConsumerTypeDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerTypeAPI.GetConsumerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the consumer type to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsumerTypeDTO**](ConsumerTypeDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsumerTypes

> []ConsumerTypeDTO GetConsumerTypes(ctx).Execute()





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
	resp, r, err := apiClient.ConsumerTypeAPI.GetConsumerTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerTypeAPI.GetConsumerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConsumerTypes`: []ConsumerTypeDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerTypeAPI.GetConsumerTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerTypesRequest struct via the builder pattern


### Return type

[**[]ConsumerTypeDTO**](ConsumerTypeDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConsumerType

> ConsumerTypeDTO UpdateConsumerType(ctx, id).ConsumerTypeDTO(consumerTypeDTO).Execute()





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
	id := "id_example" // string | The ID of the consumer type to update
	consumerTypeDTO := *openapiclient.NewConsumerTypeDTO() // ConsumerTypeDTO | The fields and updated values to apply to the specified consumer type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerTypeAPI.UpdateConsumerType(context.Background(), id).ConsumerTypeDTO(consumerTypeDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerTypeAPI.UpdateConsumerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConsumerType`: ConsumerTypeDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerTypeAPI.UpdateConsumerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the consumer type to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConsumerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumerTypeDTO** | [**ConsumerTypeDTO**](ConsumerTypeDTO.md) | The fields and updated values to apply to the specified consumer type | 

### Return type

[**ConsumerTypeDTO**](ConsumerTypeDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

