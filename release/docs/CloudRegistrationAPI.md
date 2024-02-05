# \CloudRegistrationAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelCloudAccountJobs**](CloudRegistrationAPI.md#CancelCloudAccountJobs) | **Delete** /cloud/jobs/orgsetup/{cloud_account_id} | 
[**CloudAuthorize**](CloudRegistrationAPI.md#CloudAuthorize) | **Post** /cloud/authorize | 
[**DeleteAnonymousConsumersByAccountId**](CloudRegistrationAPI.md#DeleteAnonymousConsumersByAccountId) | **Delete** /cloud/consumers/anonymous | 



## CancelCloudAccountJobs

> []string CancelCloudAccountJobs(ctx, cloudAccountId).Execute()





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
	cloudAccountId := "cloudAccountId_example" // string | The ID of the cloud account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRegistrationAPI.CancelCloudAccountJobs(context.Background(), cloudAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRegistrationAPI.CancelCloudAccountJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelCloudAccountJobs`: []string
	fmt.Fprintf(os.Stdout, "Response from `CloudRegistrationAPI.CancelCloudAccountJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudAccountId** | **string** | The ID of the cloud account | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCloudAccountJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAuthorize

> string CloudAuthorize(ctx).CloudRegistrationDTO(cloudRegistrationDTO).Version(version).Execute()





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
	cloudRegistrationDTO := *openapiclient.NewCloudRegistrationDTO("Type_example", "Metadata_example", "Signature_example") // CloudRegistrationDTO | Cloud registration data
	version := int32(56) // int32 | Version of cloud authentication (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRegistrationAPI.CloudAuthorize(context.Background()).CloudRegistrationDTO(cloudRegistrationDTO).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRegistrationAPI.CloudAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAuthorize`: string
	fmt.Fprintf(os.Stdout, "Response from `CloudRegistrationAPI.CloudAuthorize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRegistrationDTO** | [**CloudRegistrationDTO**](CloudRegistrationDTO.md) | Cloud registration data | 
 **version** | **int32** | Version of cloud authentication | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnonymousConsumersByAccountId

> DeleteAnonymousConsumersByAccountId(ctx).CloudAccountId(cloudAccountId).Execute()





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
	cloudAccountId := "cloudAccountId_example" // string | The ID of the Cloud Account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudRegistrationAPI.DeleteAnonymousConsumersByAccountId(context.Background()).CloudAccountId(cloudAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRegistrationAPI.DeleteAnonymousConsumersByAccountId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnonymousConsumersByAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAccountId** | **string** | The ID of the Cloud Account | 

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

