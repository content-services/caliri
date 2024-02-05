# \DistributorVersionsAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](DistributorVersionsAPI.md#Create) | **Post** /distributor_versions | 
[**Delete**](DistributorVersionsAPI.md#Delete) | **Delete** /distributor_versions/{id} | 
[**GetVersions**](DistributorVersionsAPI.md#GetVersions) | **Get** /distributor_versions | 
[**Update**](DistributorVersionsAPI.md#Update) | **Put** /distributor_versions/{id} | 



## Create

> DistributorVersionDTO Create(ctx).DistributorVersionDTO(distributorVersionDTO).Execute()





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
	distributorVersionDTO := *openapiclient.NewDistributorVersionDTO() // DistributorVersionDTO | A new distributor version to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributorVersionsAPI.Create(context.Background()).DistributorVersionDTO(distributorVersionDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributorVersionsAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: DistributorVersionDTO
	fmt.Fprintf(os.Stdout, "Response from `DistributorVersionsAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **distributorVersionDTO** | [**DistributorVersionDTO**](DistributorVersionDTO.md) | A new distributor version to be created | 

### Return type

[**DistributorVersionDTO**](DistributorVersionDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DistributorVersionsAPI.Delete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributorVersionsAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


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


## GetVersions

> []DistributorVersionDTO GetVersions(ctx).NameSearch(nameSearch).Capability(capability).Execute()





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
	nameSearch := "nameSearch_example" // string |  (optional)
	capability := "capability_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributorVersionsAPI.GetVersions(context.Background()).NameSearch(nameSearch).Capability(capability).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributorVersionsAPI.GetVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersions`: []DistributorVersionDTO
	fmt.Fprintf(os.Stdout, "Response from `DistributorVersionsAPI.GetVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameSearch** | **string** |  | 
 **capability** | **string** |  | 

### Return type

[**[]DistributorVersionDTO**](DistributorVersionDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> DistributorVersionDTO Update(ctx, id).DistributorVersionDTO(distributorVersionDTO).Execute()





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
	id := "id_example" // string | 
	distributorVersionDTO := *openapiclient.NewDistributorVersionDTO() // DistributorVersionDTO | The fields and updated values to apply to the specified distributor version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributorVersionsAPI.Update(context.Background(), id).DistributorVersionDTO(distributorVersionDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributorVersionsAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: DistributorVersionDTO
	fmt.Fprintf(os.Stdout, "Response from `DistributorVersionsAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distributorVersionDTO** | [**DistributorVersionDTO**](DistributorVersionDTO.md) | The fields and updated values to apply to the specified distributor version | 

### Return type

[**DistributorVersionDTO**](DistributorVersionDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

