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

> []ContentDTO GetContents(ctx).Owner(owner).Content(content).Label(label).Active(active).Custom(custom).Execute()





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
	owner := []string{"Inner_example"} // []string | The key of an owner to use to limit the content search. If defined, the list of contents returned by this endpoint will only include those available to the given owner. May be specified multiple times to filter by multiple owners. If multiple owners are provided, contents available to any of the provided owners will be returned.  (optional)
	content := []string{"Inner_example"} // []string | The ID of a content to fetch. If defined, the list of contents returned by this method will only include those matching the given ID. May be specified multiple times to filter on multiple content IDs.  (optional)
	label := []string{"Inner_example"} // []string | The labels of content to fetch. If defined, the list of content returned by this endpoint will only include those matching the given labels (case-insensitive). May be specified multiple times to filter on multiple labels. If multiple labels are provided, any content matching any of the provided labels will be returned.  (optional)
	active := "{"include":{"value":"include","description":"indicates the query should include matching active content, along with any inactive\ncontent\n"},"exclude":{"value":"exclude","description":"indicates the query should exclude, or omit, active content, returning only matching\ninactive content\n"},"exclusive":{"value":"exclusive","description":"indicates the query should only include matching active content, excluding any inactive\ncontent\n"}}" // string | A value indicating how 'active' content should be considered when fetching content, where 'active' is defined as a content that is currently in use by a subscription with a start time in the past and that has not yet expired, or in use by a product which itself is considered 'active.' Must be one of 'include', 'exclude', or 'exclusive' indicating that active content should be included along with inactive content, excluded (omitted) from the results, or exclusively selected as the only content to return. Defaults to 'exclusive'.  (optional) (default to "exclusive")
	custom := "{"include":{"value":"include","description":"indicates the query should include matching custom content, along with any matching\nglobally defined content\n"},"exclude":{"value":"exclude","description":"indicates the query should exclude, or omit, custom content, returning only matching\nglobally defined content\n"},"exclusive":{"value":"exclusive","description":"indicates the query should only include matching custom content, excluding any globally\ndefined content\n"}}" // string | A value indicating how custom content are considered when fetching content, where 'custom' is defined as content that did not originate from a refresh operation nor manifest import. Must be one of 'include', 'exclude', or 'exclusive' indicating that custom content should be passively included, excluded or omitted from the output, or exclusively selected as the only content to return.  (optional) (default to "include")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetContents(context.Background()).Owner(owner).Content(content).Label(label).Active(active).Custom(custom).Execute()
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
 **owner** | **[]string** | The key of an owner to use to limit the content search. If defined, the list of contents returned by this endpoint will only include those available to the given owner. May be specified multiple times to filter by multiple owners. If multiple owners are provided, contents available to any of the provided owners will be returned.  | 
 **content** | **[]string** | The ID of a content to fetch. If defined, the list of contents returned by this method will only include those matching the given ID. May be specified multiple times to filter on multiple content IDs.  | 
 **label** | **[]string** | The labels of content to fetch. If defined, the list of content returned by this endpoint will only include those matching the given labels (case-insensitive). May be specified multiple times to filter on multiple labels. If multiple labels are provided, any content matching any of the provided labels will be returned.  | 
 **active** | **string** | A value indicating how &#39;active&#39; content should be considered when fetching content, where &#39;active&#39; is defined as a content that is currently in use by a subscription with a start time in the past and that has not yet expired, or in use by a product which itself is considered &#39;active.&#39; Must be one of &#39;include&#39;, &#39;exclude&#39;, or &#39;exclusive&#39; indicating that active content should be included along with inactive content, excluded (omitted) from the results, or exclusively selected as the only content to return. Defaults to &#39;exclusive&#39;.  | [default to &quot;exclusive&quot;]
 **custom** | **string** | A value indicating how custom content are considered when fetching content, where &#39;custom&#39; is defined as content that did not originate from a refresh operation nor manifest import. Must be one of &#39;include&#39;, &#39;exclude&#39;, or &#39;exclusive&#39; indicating that custom content should be passively included, excluded or omitted from the output, or exclusively selected as the only content to return.  | [default to &quot;include&quot;]

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

