# \CertificateSerialAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCertificateSerial**](CertificateSerialAPI.md#GetCertificateSerial) | **Get** /serials/{id} | 



## GetCertificateSerial

> CertificateSerialDTO GetCertificateSerial(ctx, id).Execute()





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
	id := int64(789) // int64 | The ID of the certificate serial to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateSerialAPI.GetCertificateSerial(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateSerialAPI.GetCertificateSerial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCertificateSerial`: CertificateSerialDTO
	fmt.Fprintf(os.Stdout, "Response from `CertificateSerialAPI.GetCertificateSerial`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the certificate serial to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateSerialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateSerialDTO**](CertificateSerialDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

