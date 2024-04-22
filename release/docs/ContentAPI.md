# \ContentAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContentByUuid**](ContentAPI.md#GetContentByUuid) | **Get** /content/{content_uuid} | 
[**GetContents**](ContentAPI.md#GetContents) | **Get** /content | 



## GetContentByUuid

> ContentDTO GetContentByUuid(ctx, contentUuid).Execute()





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
	contentUuid := "contentUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetContentByUuid(context.Background(), contentUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.GetContentByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentByUuid`: ContentDTO
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.GetContentByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentDTO**](ContentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContents

> []ContentDTO GetContents(ctx).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()





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
	page := int32(2) // int32 | Page index to return (optional)
	perPage := int32(10) // int32 | Number of items to return per page (optional)
	order := "asc" // string | Direction of ordering (optional)
	sortBy := "name" // string | Property to use for ordering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetContents(context.Background()).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.GetContents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContents`: []ContentDTO
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.GetContents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page index to return | 
 **perPage** | **int32** | Number of items to return per page | 
 **order** | **string** | Direction of ordering | 
 **sortBy** | **string** | Property to use for ordering | 

### Return type

[**[]ContentDTO**](ContentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

