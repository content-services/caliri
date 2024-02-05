# \GuestIdsAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGuest**](GuestIdsAPI.md#DeleteGuest) | **Delete** /consumers/{consumer_uuid}/guestids/{guest_id} | 
[**GetGuestId**](GuestIdsAPI.md#GetGuestId) | **Get** /consumers/{consumer_uuid}/guestids/{guest_id} | 
[**GetGuestIds**](GuestIdsAPI.md#GetGuestIds) | **Get** /consumers/{consumer_uuid}/guestids | 
[**UpdateGuest**](GuestIdsAPI.md#UpdateGuest) | **Put** /consumers/{consumer_uuid}/guestids/{guest_id} | 
[**UpdateGuests**](GuestIdsAPI.md#UpdateGuests) | **Put** /consumers/{consumer_uuid}/guestids | 



## DeleteGuest

> DeleteGuest(ctx, consumerUuid, guestId).Unregister(unregister).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the consumer who owns or hosts the guest
	guestId := "guestId_example" // string | The UUID of the guest to be deleted
	unregister := true // bool | Unregister the consumer (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GuestIdsAPI.DeleteGuest(context.Background(), consumerUuid, guestId).Unregister(unregister).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestIdsAPI.DeleteGuest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the consumer who owns or hosts the guest | 
**guestId** | **string** | The UUID of the guest to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unregister** | **bool** | Unregister the consumer | [default to false]

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


## GetGuestId

> GuestIdDTO GetGuestId(ctx, consumerUuid, guestId).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the consumer to retrieve guest
	guestId := "guestId_example" // string | The UUID of the guest to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestIdsAPI.GetGuestId(context.Background(), consumerUuid, guestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestIdsAPI.GetGuestId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuestId`: GuestIdDTO
	fmt.Fprintf(os.Stdout, "Response from `GuestIdsAPI.GetGuestId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the consumer to retrieve guest | 
**guestId** | **string** | The UUID of the guest to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuestIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GuestIdDTO**](GuestIdDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuestIds

> []GuestIdDTOArrayElement GetGuestIds(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestIdsAPI.GetGuestIds(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestIdsAPI.GetGuestIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuestIds`: []GuestIdDTOArrayElement
	fmt.Fprintf(os.Stdout, "Response from `GuestIdsAPI.GetGuestIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuestIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GuestIdDTOArrayElement**](GuestIdDTOArrayElement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuest

> UpdateGuest(ctx, consumerUuid, guestId).GuestIdDTO(guestIdDTO).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the consumer who owns or hosts the guest
	guestId := "guestId_example" // string | The UUID of the guest to be updated
	guestIdDTO := *openapiclient.NewGuestIdDTO("GuestId_example") // GuestIdDTO | Updated guest data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GuestIdsAPI.UpdateGuest(context.Background(), consumerUuid, guestId).GuestIdDTO(guestIdDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestIdsAPI.UpdateGuest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the consumer who owns or hosts the guest | 
**guestId** | **string** | The UUID of the guest to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **guestIdDTO** | [**GuestIdDTO**](GuestIdDTO.md) | Updated guest data | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGuests

> UpdateGuests(ctx, consumerUuid).GuestIdDTO(guestIdDTO).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the consumer who owns or hosts the guest
	guestIdDTO := []openapiclient.GuestIdDTO{*openapiclient.NewGuestIdDTO("GuestId_example")} // []GuestIdDTO | The list of the guests to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GuestIdsAPI.UpdateGuests(context.Background(), consumerUuid).GuestIdDTO(guestIdDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestIdsAPI.UpdateGuests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the consumer who owns or hosts the guest | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGuestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guestIdDTO** | [**[]GuestIdDTO**](GuestIdDTO.md) | The list of the guests to be updated | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

