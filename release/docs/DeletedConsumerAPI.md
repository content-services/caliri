# \DeletedConsumerAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListByDate**](DeletedConsumerAPI.md#ListByDate) | **Get** /deleted_consumers | 



## ListByDate

> []DeletedConsumerDTO ListByDate(ctx).Date(date).Execute()





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
	date := time.Now() // time.Time | Filter deleted consumers to those on or after this date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeletedConsumerAPI.ListByDate(context.Background()).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeletedConsumerAPI.ListByDate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListByDate`: []DeletedConsumerDTO
	fmt.Fprintf(os.Stdout, "Response from `DeletedConsumerAPI.ListByDate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListByDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** | Filter deleted consumers to those on or after this date | 

### Return type

[**[]DeletedConsumerDTO**](DeletedConsumerDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

