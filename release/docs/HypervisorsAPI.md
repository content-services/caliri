# \HypervisorsAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HypervisorHeartbeatUpdate**](HypervisorsAPI.md#HypervisorHeartbeatUpdate) | **Put** /hypervisors/{owner}/heartbeat | 
[**HypervisorUpdateAsync**](HypervisorsAPI.md#HypervisorUpdateAsync) | **Post** /hypervisors/{owner} | 



## HypervisorHeartbeatUpdate

> AsyncJobStatusDTO HypervisorHeartbeatUpdate(ctx, owner).ReporterId(reporterId).Execute()





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
	owner := "owner_example" // string | Owner key
	reporterId := "reporterId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HypervisorsAPI.HypervisorHeartbeatUpdate(context.Background(), owner).ReporterId(reporterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPI.HypervisorHeartbeatUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HypervisorHeartbeatUpdate`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPI.HypervisorHeartbeatUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owner key | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorHeartbeatUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reporterId** | **string** |  | 

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


## HypervisorUpdateAsync

> AsyncJobStatusDTO HypervisorUpdateAsync(ctx, owner).CreateMissing(createMissing).ReporterId(reporterId).Body(body).Execute()





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
	owner := "owner_example" // string | Owner key
	createMissing := true // bool | Specify whether or not to create missing hypervisors. Default is true.  If false is specified, hypervisorIds that are not found will result in failed entries in the resulting HypervisorCheckInResult. (optional) (default to true)
	reporterId := "reporterId_example" // string |  (optional)
	body := "body_example" // string | Host and Guest mapping details (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HypervisorsAPI.HypervisorUpdateAsync(context.Background(), owner).CreateMissing(createMissing).ReporterId(reporterId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HypervisorsAPI.HypervisorUpdateAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HypervisorUpdateAsync`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `HypervisorsAPI.HypervisorUpdateAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owner key | 

### Other Parameters

Other parameters are passed through a pointer to a apiHypervisorUpdateAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMissing** | **bool** | Specify whether or not to create missing hypervisors. Default is true.  If false is specified, hypervisorIds that are not found will result in failed entries in the resulting HypervisorCheckInResult. | [default to true]
 **reporterId** | **string** |  | 
 **body** | **string** | Host and Guest mapping details | 

### Return type

[**AsyncJobStatusDTO**](AsyncJobStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

