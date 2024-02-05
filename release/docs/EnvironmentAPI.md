# \EnvironmentAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConsumerInEnvironment**](EnvironmentAPI.md#CreateConsumerInEnvironment) | **Post** /environments/{env_id}/consumers | 
[**DeleteEnvironment**](EnvironmentAPI.md#DeleteEnvironment) | **Delete** /environments/{env_id} | 
[**DemoteContent**](EnvironmentAPI.md#DemoteContent) | **Delete** /environments/{env_id}/content | 
[**GetEnvironment**](EnvironmentAPI.md#GetEnvironment) | **Get** /environments/{env_id} | 
[**PromoteContent**](EnvironmentAPI.md#PromoteContent) | **Post** /environments/{env_id}/content | 



## CreateConsumerInEnvironment

> ConsumerDTO CreateConsumerInEnvironment(ctx, envId).ConsumerDTO(consumerDTO).Username(username).ActivationKeys(activationKeys).Execute()





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
	envId := "envId_example" // string | 
	consumerDTO := *openapiclient.NewConsumerDTO() // ConsumerDTO | 
	username := "username_example" // string |  (optional)
	activationKeys := "activationKeys_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.CreateConsumerInEnvironment(context.Background(), envId).ConsumerDTO(consumerDTO).Username(username).ActivationKeys(activationKeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.CreateConsumerInEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConsumerInEnvironment`: ConsumerDTO
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.CreateConsumerInEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConsumerInEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumerDTO** | [**ConsumerDTO**](ConsumerDTO.md) |  | 
 **username** | **string** |  | 
 **activationKeys** | **string** |  | 

### Return type

[**ConsumerDTO**](ConsumerDTO.md)

### Authorization

[ActivationKey](../README.md#ActivationKey), [ActivationKeyOwner](../README.md#ActivationKeyOwner)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironment(ctx, envId).Execute()





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
	envId := "envId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentAPI.DeleteEnvironment(context.Background(), envId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.DeleteEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


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


## DemoteContent

> AsyncJobStatusDTO DemoteContent(ctx, envId).Content(content).LazyRegen(lazyRegen).Execute()





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
	envId := "envId_example" // string | 
	content := []string{"Inner_example"} // []string | 
	lazyRegen := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.DemoteContent(context.Background(), envId).Content(content).LazyRegen(lazyRegen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.DemoteContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DemoteContent`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.DemoteContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDemoteContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **[]string** |  | 
 **lazyRegen** | **bool** |  | [default to true]

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


## GetEnvironment

> EnvironmentDTO GetEnvironment(ctx, envId).Execute()





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
	envId := "envId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.GetEnvironment(context.Background(), envId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.GetEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironment`: EnvironmentDTO
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.GetEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentDTO**](EnvironmentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromoteContent

> AsyncJobStatusDTO PromoteContent(ctx, envId).ContentToPromoteDTO(contentToPromoteDTO).LazyRegen(lazyRegen).Execute()





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
	envId := "envId_example" // string | 
	contentToPromoteDTO := []openapiclient.ContentToPromoteDTO{*openapiclient.NewContentToPromoteDTO()} // []ContentToPromoteDTO | Contents to promote
	lazyRegen := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.PromoteContent(context.Background(), envId).ContentToPromoteDTO(contentToPromoteDTO).LazyRegen(lazyRegen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.PromoteContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromoteContent`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.PromoteContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromoteContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentToPromoteDTO** | [**[]ContentToPromoteDTO**](ContentToPromoteDTO.md) | Contents to promote | 
 **lazyRegen** | **bool** |  | [default to true]

### Return type

[**AsyncJobStatusDTO**](AsyncJobStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

