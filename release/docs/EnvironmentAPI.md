# \EnvironmentAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConsumerInEnvironment**](EnvironmentAPI.md#CreateConsumerInEnvironment) | **Post** /environments/{env_id}/consumers | 
[**DeleteEnvironment**](EnvironmentAPI.md#DeleteEnvironment) | **Delete** /environments/{env_id} | 
[**DeleteEnvironmentContentOverrides**](EnvironmentAPI.md#DeleteEnvironmentContentOverrides) | **Delete** /environments/{environment_id}/content_overrides | 
[**DemoteContent**](EnvironmentAPI.md#DemoteContent) | **Delete** /environments/{env_id}/content | 
[**GetEnvironment**](EnvironmentAPI.md#GetEnvironment) | **Get** /environments/{env_id} | 
[**GetEnvironmentContentOverrides**](EnvironmentAPI.md#GetEnvironmentContentOverrides) | **Get** /environments/{environment_id}/content_overrides | 
[**PromoteContent**](EnvironmentAPI.md#PromoteContent) | **Post** /environments/{env_id}/content | 
[**PutEnvironmentContentOverrides**](EnvironmentAPI.md#PutEnvironmentContentOverrides) | **Put** /environments/{environment_id}/content_overrides | 
[**UpdateEnvironment**](EnvironmentAPI.md#UpdateEnvironment) | **Put** /environments/{env_id} | 



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

> DeleteEnvironment(ctx, envId).RetainConsumers(retainConsumers).Execute()





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
	retainConsumers := true // bool | Whether or not to retain affected consumers with no remaining environments. If set to true, affected consumers will be moved to the organization's default content view.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentAPI.DeleteEnvironment(context.Background(), envId).RetainConsumers(retainConsumers).Execute()
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

 **retainConsumers** | **bool** | Whether or not to retain affected consumers with no remaining environments. If set to true, affected consumers will be moved to the organization&#39;s default content view.  | [default to false]

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


## DeleteEnvironmentContentOverrides

> []ContentOverrideDTO DeleteEnvironmentContentOverrides(ctx, environmentId).ContentOverrideDTO(contentOverrideDTO).Execute()





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
	environmentId := "environmentId_example" // string | The ID of the environment from which to remove content overrides 
	contentOverrideDTO := []openapiclient.ContentOverrideDTO{*openapiclient.NewContentOverrideDTO()} // []ContentOverrideDTO | A list containing the content overrides to remove from the environment. The overrides need not be fully populated, as the value is ignored entirely during removal; and depending on the behavior desired, the name or label fields may not be needed, either. See the description for the operation itself for more details.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.DeleteEnvironmentContentOverrides(context.Background(), environmentId).ContentOverrideDTO(contentOverrideDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.DeleteEnvironmentContentOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEnvironmentContentOverrides`: []ContentOverrideDTO
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.DeleteEnvironmentContentOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment from which to remove content overrides  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentContentOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentOverrideDTO** | [**[]ContentOverrideDTO**](ContentOverrideDTO.md) | A list containing the content overrides to remove from the environment. The overrides need not be fully populated, as the value is ignored entirely during removal; and depending on the behavior desired, the name or label fields may not be needed, either. See the description for the operation itself for more details.  | 

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


## GetEnvironmentContentOverrides

> []ContentOverrideDTO GetEnvironmentContentOverrides(ctx, environmentId).Execute()





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
	environmentId := "environmentId_example" // string | The ID of the environment for which to fetch content overrides 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.GetEnvironmentContentOverrides(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.GetEnvironmentContentOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentContentOverrides`: []ContentOverrideDTO
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.GetEnvironmentContentOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment for which to fetch content overrides  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentContentOverridesRequest struct via the builder pattern


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


## PutEnvironmentContentOverrides

> []ContentOverrideDTO PutEnvironmentContentOverrides(ctx, environmentId).ContentOverrideDTO(contentOverrideDTO).Execute()





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
	environmentId := "environmentId_example" // string | The ID of the environment in which to add or update content overrides 
	contentOverrideDTO := []openapiclient.ContentOverrideDTO{*openapiclient.NewContentOverrideDTO()} // []ContentOverrideDTO | A list containing the content overrides to apply to the environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.PutEnvironmentContentOverrides(context.Background(), environmentId).ContentOverrideDTO(contentOverrideDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.PutEnvironmentContentOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutEnvironmentContentOverrides`: []ContentOverrideDTO
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.PutEnvironmentContentOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment in which to add or update content overrides  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutEnvironmentContentOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentOverrideDTO** | [**[]ContentOverrideDTO**](ContentOverrideDTO.md) | A list containing the content overrides to apply to the environment | 

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


## UpdateEnvironment

> EnvironmentDTO UpdateEnvironment(ctx, envId).EnvironmentDTO(environmentDTO).Execute()





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
	environmentDTO := *openapiclient.NewEnvironmentDTO() // EnvironmentDTO | An EnvironmentDTO containing the data to use to update the target environment. The object need not be fully populated, as some fields cannot be changed by an update operation and will be ignored. Any field which is not intended to be updated should be left undefined or null. An empty, non-null value should be used to clear any existing value in a given field. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.UpdateEnvironment(context.Background(), envId).EnvironmentDTO(environmentDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.UpdateEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEnvironment`: EnvironmentDTO
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.UpdateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentDTO** | [**EnvironmentDTO**](EnvironmentDTO.md) | An EnvironmentDTO containing the data to use to update the target environment. The object need not be fully populated, as some fields cannot be changed by an update operation and will be ignored. Any field which is not intended to be updated should be left undefined or null. An empty, non-null value should be used to clear any existing value in a given field.  | 

### Return type

[**EnvironmentDTO**](EnvironmentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

