# \CdnAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCdn**](CdnAPI.md#CreateCdn) | **Post** /cdn | 
[**DeleteCdn**](CdnAPI.md#DeleteCdn) | **Delete** /cdn/{label} | 
[**GetContentDeliveryNetworks**](CdnAPI.md#GetContentDeliveryNetworks) | **Get** /cdn | 
[**UpdateCdn**](CdnAPI.md#UpdateCdn) | **Put** /cdn/{label} | 



## CreateCdn

> CdnDTO CreateCdn(ctx).CdnDTO(cdnDTO).Execute()





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
	cdnDTO := *openapiclient.NewCdnDTO() // CdnDTO | CDN to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CdnAPI.CreateCdn(context.Background()).CdnDTO(cdnDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CdnAPI.CreateCdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCdn`: CdnDTO
	fmt.Fprintf(os.Stdout, "Response from `CdnAPI.CreateCdn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCdnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cdnDTO** | [**CdnDTO**](CdnDTO.md) | CDN to be created | 

### Return type

[**CdnDTO**](CdnDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCdn

> DeleteCdn(ctx, label).Execute()





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
	label := "label_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CdnAPI.DeleteCdn(context.Background(), label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CdnAPI.DeleteCdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**label** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCdnRequest struct via the builder pattern


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


## GetContentDeliveryNetworks

> []CdnDTO GetContentDeliveryNetworks(ctx).Execute()





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
	resp, r, err := apiClient.CdnAPI.GetContentDeliveryNetworks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CdnAPI.GetContentDeliveryNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentDeliveryNetworks`: []CdnDTO
	fmt.Fprintf(os.Stdout, "Response from `CdnAPI.GetContentDeliveryNetworks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentDeliveryNetworksRequest struct via the builder pattern


### Return type

[**[]CdnDTO**](CdnDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCdn

> CdnDTO UpdateCdn(ctx, label).CdnDTO(cdnDTO).Execute()





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
	label := "label_example" // string | 
	cdnDTO := *openapiclient.NewCdnDTO() // CdnDTO | Fields that needs to be updated for specified CDN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CdnAPI.UpdateCdn(context.Background(), label).CdnDTO(cdnDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CdnAPI.UpdateCdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCdn`: CdnDTO
	fmt.Fprintf(os.Stdout, "Response from `CdnAPI.UpdateCdn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**label** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCdnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnDTO** | [**CdnDTO**](CdnDTO.md) | Fields that needs to be updated for specified CDN | 

### Return type

[**CdnDTO**](CdnDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

