# \JobsAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelJob**](JobsAPI.md#CancelJob) | **Delete** /jobs/{id} | 
[**CleanupTerminalJobs**](JobsAPI.md#CleanupTerminalJobs) | **Delete** /jobs | 
[**GetJobStatus**](JobsAPI.md#GetJobStatus) | **Get** /jobs/{id} | 
[**GetSchedulerStatus**](JobsAPI.md#GetSchedulerStatus) | **Get** /jobs/scheduler | 
[**ListJobStatuses**](JobsAPI.md#ListJobStatuses) | **Get** /jobs | 
[**ScheduleJob**](JobsAPI.md#ScheduleJob) | **Post** /jobs/schedule/{jobKey} | 
[**SetSchedulerStatus**](JobsAPI.md#SetSchedulerStatus) | **Post** /jobs/scheduler | 



## CancelJob

> AsyncJobStatusDTO CancelJob(ctx, id).Execute()





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
	id := "id_example" // string | The ID of the job to cancel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.CancelJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.CancelJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelJob`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.CancelJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the job to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CleanupTerminalJobs

> int32 CleanupTerminalJobs(ctx).Id(id).Key(key).State(state).Owner(owner).Principal(principal).Origin(origin).Executor(executor).After(after).Before(before).Force(force).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/content-services/caliri/release/v4"
)

func main() {
	id := []string{"Inner_example"} // []string | Cleans up jobs based on ids (optional)
	key := []string{"Inner_example"} // []string | Cleans up jobs based on keys (optional)
	state := []string{"Inner_example"} // []string | Cleans up jobs based on statuses (optional)
	owner := []string{"Inner_example"} // []string | Cleans up jobs based on owners (optional)
	principal := []string{"Inner_example"} // []string | Cleans up jobs based on principals (optional)
	origin := []string{"Inner_example"} // []string | Cleans up jobs based on origins (optional)
	executor := []string{"Inner_example"} // []string | Cleans up jobs based on executors (optional)
	after := time.Now() // time.Time | Cleans up jobs to those on or after this date (optional)
	before := time.Now() // time.Time | Cleans up jobs to those on or before this date (optional)
	force := true // bool | Cleans up job forcefully (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.CleanupTerminalJobs(context.Background()).Id(id).Key(key).State(state).Owner(owner).Principal(principal).Origin(origin).Executor(executor).After(after).Before(before).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.CleanupTerminalJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CleanupTerminalJobs`: int32
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.CleanupTerminalJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCleanupTerminalJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Cleans up jobs based on ids | 
 **key** | **[]string** | Cleans up jobs based on keys | 
 **state** | **[]string** | Cleans up jobs based on statuses | 
 **owner** | **[]string** | Cleans up jobs based on owners | 
 **principal** | **[]string** | Cleans up jobs based on principals | 
 **origin** | **[]string** | Cleans up jobs based on origins | 
 **executor** | **[]string** | Cleans up jobs based on executors | 
 **after** | **time.Time** | Cleans up jobs to those on or after this date | 
 **before** | **time.Time** | Cleans up jobs to those on or before this date | 
 **force** | **bool** | Cleans up job forcefully | [default to false]

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobStatus

> AsyncJobStatusDTO GetJobStatus(ctx, id).Execute()





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
	id := "id_example" // string | The ID of the job to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.GetJobStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.GetJobStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobStatus`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.GetJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the job to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSchedulerStatus

> SchedulerStatusDTO GetSchedulerStatus(ctx).Execute()





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
	resp, r, err := apiClient.JobsAPI.GetSchedulerStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.GetSchedulerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchedulerStatus`: SchedulerStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.GetSchedulerStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulerStatusRequest struct via the builder pattern


### Return type

[**SchedulerStatusDTO**](SchedulerStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJobStatuses

> []AsyncJobStatusDTO ListJobStatuses(ctx).Id(id).Key(key).State(state).Owner(owner).Principal(principal).Origin(origin).Executor(executor).After(after).Before(before).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/content-services/caliri/release/v4"
)

func main() {
	id := []string{"Inner_example"} // []string | Filter jobs based on ids (optional)
	key := []string{"Inner_example"} // []string | Filter jobs based on keys (optional)
	state := []string{"Inner_example"} // []string | Filter jobs based on statuses (optional)
	owner := []string{"Inner_example"} // []string | Filter jobs based on owners (optional)
	principal := []string{"Inner_example"} // []string | Filter jobs based on principals (optional)
	origin := []string{"Inner_example"} // []string | Filter jobs based on origins (optional)
	executor := []string{"Inner_example"} // []string | Filter jobs based on executors (optional)
	after := time.Now() // time.Time | Filter jobs to those on or after this date (optional)
	before := time.Now() // time.Time | Filter jobs to those on or before this date (optional)
	page := int32(2) // int32 | Page index to return (optional)
	perPage := int32(10) // int32 | Number of items to return per page (optional)
	order := "asc" // string | Direction of ordering (optional)
	sortBy := "name" // string | Property to use for ordering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.ListJobStatuses(context.Background()).Id(id).Key(key).State(state).Owner(owner).Principal(principal).Origin(origin).Executor(executor).After(after).Before(before).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.ListJobStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJobStatuses`: []AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.ListJobStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListJobStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Filter jobs based on ids | 
 **key** | **[]string** | Filter jobs based on keys | 
 **state** | **[]string** | Filter jobs based on statuses | 
 **owner** | **[]string** | Filter jobs based on owners | 
 **principal** | **[]string** | Filter jobs based on principals | 
 **origin** | **[]string** | Filter jobs based on origins | 
 **executor** | **[]string** | Filter jobs based on executors | 
 **after** | **time.Time** | Filter jobs to those on or after this date | 
 **before** | **time.Time** | Filter jobs to those on or before this date | 
 **page** | **int32** | Page index to return | 
 **perPage** | **int32** | Number of items to return per page | 
 **order** | **string** | Direction of ordering | 
 **sortBy** | **string** | Property to use for ordering | 

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


## ScheduleJob

> AsyncJobStatusDTO ScheduleJob(ctx, jobKey).Execute()





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
	jobKey := "jobKey_example" // string | Job key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.ScheduleJob(context.Background(), jobKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.ScheduleJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleJob`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.ScheduleJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobKey** | **string** | Job key | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SetSchedulerStatus

> SchedulerStatusDTO SetSchedulerStatus(ctx).Running(running).Execute()





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
	running := true // bool | Boolean value to set running status (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.SetSchedulerStatus(context.Background()).Running(running).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.SetSchedulerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSchedulerStatus`: SchedulerStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.SetSchedulerStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSchedulerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **running** | **bool** | Boolean value to set running status | [default to true]

### Return type

[**SchedulerStatusDTO**](SchedulerStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

