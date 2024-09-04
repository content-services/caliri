# \ConsumerAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConsumerContentOverrides**](ConsumerAPI.md#AddConsumerContentOverrides) | **Put** /consumers/{consumer_uuid}/content_overrides | 
[**Bind**](ConsumerAPI.md#Bind) | **Post** /consumers/{consumer_uuid}/entitlements | 
[**ConsumerExists**](ConsumerAPI.md#ConsumerExists) | **Head** /consumers/{consumer_uuid}/exists | 
[**ConsumerExistsBulk**](ConsumerAPI.md#ConsumerExistsBulk) | **Post** /consumers/exists | 
[**CreateConsumer**](ConsumerAPI.md#CreateConsumer) | **Post** /consumers | 
[**DeleteConsumer**](ConsumerAPI.md#DeleteConsumer) | **Delete** /consumers/{consumer_uuid} | 
[**DeleteConsumerContentOverrides**](ConsumerAPI.md#DeleteConsumerContentOverrides) | **Delete** /consumers/{consumer_uuid}/content_overrides | 
[**DownloadExistingExport**](ConsumerAPI.md#DownloadExistingExport) | **Get** /consumers/{consumer_uuid}/export/{export_id} | 
[**DryBind**](ConsumerAPI.md#DryBind) | **Get** /consumers/{consumer_uuid}/entitlements/dry-run | 
[**ExportCertificates**](ConsumerAPI.md#ExportCertificates) | **Get** /consumers/{consumer_uuid}/certificates | 
[**ExportData**](ConsumerAPI.md#ExportData) | **Get** /consumers/{consumer_uuid}/export | 
[**ExportDataAsync**](ConsumerAPI.md#ExportDataAsync) | **Get** /consumers/{consumer_uuid}/export/async | 
[**GetComplianceStatus**](ConsumerAPI.md#GetComplianceStatus) | **Get** /consumers/{consumer_uuid}/compliance | 
[**GetComplianceStatusList**](ConsumerAPI.md#GetComplianceStatusList) | **Get** /consumers/compliance | 
[**GetConsumer**](ConsumerAPI.md#GetConsumer) | **Get** /consumers/{consumer_uuid} | 
[**GetContentAccessBody**](ConsumerAPI.md#GetContentAccessBody) | **Get** /consumers/{consumer_uuid}/accessible_content | 
[**GetContentAccessForConsumer**](ConsumerAPI.md#GetContentAccessForConsumer) | **Get** /consumers/{consumer_uuid}/content_access | 
[**GetEntitlementCertificateSerials**](ConsumerAPI.md#GetEntitlementCertificateSerials) | **Get** /consumers/{consumer_uuid}/certificates/serials | 
[**GetGuests**](ConsumerAPI.md#GetGuests) | **Get** /consumers/{consumer_uuid}/guests | 
[**GetHost**](ConsumerAPI.md#GetHost) | **Get** /consumers/{consumer_uuid}/host | 
[**GetOwnerByConsumerUuid**](ConsumerAPI.md#GetOwnerByConsumerUuid) | **Get** /consumers/{consumer_uuid}/owner | 
[**GetRelease**](ConsumerAPI.md#GetRelease) | **Get** /consumers/{consumer_uuid}/release | 
[**GetSystemPurposeComplianceStatus**](ConsumerAPI.md#GetSystemPurposeComplianceStatus) | **Get** /consumers/{consumer_uuid}/purpose_compliance | 
[**ListConsumerContentOverrides**](ConsumerAPI.md#ListConsumerContentOverrides) | **Get** /consumers/{consumer_uuid}/content_overrides | 
[**ListEntitlements**](ConsumerAPI.md#ListEntitlements) | **Get** /consumers/{consumer_uuid}/entitlements | 
[**RegenerateEntitlementCertificates**](ConsumerAPI.md#RegenerateEntitlementCertificates) | **Put** /consumers/{consumer_uuid}/certificates | 
[**RegenerateIdentityCertificates**](ConsumerAPI.md#RegenerateIdentityCertificates) | **Post** /consumers/{consumer_uuid} | 
[**RemoveDeletionRecord**](ConsumerAPI.md#RemoveDeletionRecord) | **Delete** /consumers/{consumer_uuid}/deletionrecord | 
[**SearchConsumers**](ConsumerAPI.md#SearchConsumers) | **Get** /consumers | 
[**UnbindAll**](ConsumerAPI.md#UnbindAll) | **Delete** /consumers/{consumer_uuid}/entitlements | 
[**UnbindByEntitlementId**](ConsumerAPI.md#UnbindByEntitlementId) | **Delete** /consumers/{consumer_uuid}/entitlements/{dbid} | 
[**UnbindByPool**](ConsumerAPI.md#UnbindByPool) | **Delete** /consumers/{consumer_uuid}/entitlements/pool/{pool_id} | 
[**UnbindBySerial**](ConsumerAPI.md#UnbindBySerial) | **Delete** /consumers/{consumer_uuid}/certificates/{serial} | 
[**UpdateConsumer**](ConsumerAPI.md#UpdateConsumer) | **Put** /consumers/{consumer_uuid} | 



## AddConsumerContentOverrides

> []ContentOverrideDTO AddConsumerContentOverrides(ctx, consumerUuid).ContentOverrideDTO(contentOverrideDTO).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The ID of the Consumer
	contentOverrideDTO := []openapiclient.ContentOverrideDTO{*openapiclient.NewContentOverrideDTO()} // []ContentOverrideDTO | The list of the content overrides

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.AddConsumerContentOverrides(context.Background(), consumerUuid).ContentOverrideDTO(contentOverrideDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.AddConsumerContentOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddConsumerContentOverrides`: []ContentOverrideDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.AddConsumerContentOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The ID of the Consumer | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddConsumerContentOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentOverrideDTO** | [**[]ContentOverrideDTO**](ContentOverrideDTO.md) | The list of the content overrides | 

### Return type

[**[]ContentOverrideDTO**](ContentOverrideDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Bind

> string Bind(ctx, consumerUuid).Pool(pool).Product(product).Quantity(quantity).Email(email).EmailLocale(emailLocale).Async(async).EntitleDate(entitleDate).FromPool(fromPool).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID
	pool := "pool_example" // string | Pool ID (optional)
	product := []string{"Inner_example"} // []string | Product array (optional)
	quantity := int32(56) // int32 | Quantity (optional)
	email := "email_example" // string | Email address (optional)
	emailLocale := "emailLocale_example" // string | Email locale (optional)
	async := true // bool | Operation Type async or sync (optional) (default to false)
	entitleDate := "entitleDate_example" // string | Entitlement date (optional)
	fromPool := []string{"Inner_example"} // []string | From pool (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.Bind(context.Background(), consumerUuid).Pool(pool).Product(product).Quantity(quantity).Email(email).EmailLocale(emailLocale).Async(async).EntitleDate(entitleDate).FromPool(fromPool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.Bind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Bind`: string
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.Bind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pool** | **string** | Pool ID | 
 **product** | **[]string** | Product array | 
 **quantity** | **int32** | Quantity | 
 **email** | **string** | Email address | 
 **emailLocale** | **string** | Email locale | 
 **async** | **bool** | Operation Type async or sync | [default to false]
 **entitleDate** | **string** | Entitlement date | 
 **fromPool** | **[]string** | From pool | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumerExists

> ConsumerExists(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the Consumer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumerAPI.ConsumerExists(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.ConsumerExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the Consumer | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumerExistsRequest struct via the builder pattern


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


## ConsumerExistsBulk

> ConsumerExistsBulk(ctx).RequestBody(requestBody).Execute()





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
	requestBody := []string{"Property_example"} // []string | Array of Consumer UUIDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumerAPI.ConsumerExistsBulk(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.ConsumerExistsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsumerExistsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | Array of Consumer UUIDs | 

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


## CreateConsumer

> ConsumerDTO CreateConsumer(ctx).ConsumerDTO(consumerDTO).Username(username).Owner(owner).ActivationKeys(activationKeys).IdentityCertCreation(identityCertCreation).Execute()





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
	consumerDTO := *openapiclient.NewConsumerDTO() // ConsumerDTO | Consumer to be created
	username := "username_example" // string | User name (optional)
	owner := "owner_example" // string | Owner key (optional)
	activationKeys := "activationKeys_example" // string | Activation key (optional)
	identityCertCreation := true // bool | Boolean flag for identity cert generation. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.CreateConsumer(context.Background()).ConsumerDTO(consumerDTO).Username(username).Owner(owner).ActivationKeys(activationKeys).IdentityCertCreation(identityCertCreation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.CreateConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConsumer`: ConsumerDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.CreateConsumer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consumerDTO** | [**ConsumerDTO**](ConsumerDTO.md) | Consumer to be created | 
 **username** | **string** | User name | 
 **owner** | **string** | Owner key | 
 **activationKeys** | **string** | Activation key | 
 **identityCertCreation** | **bool** | Boolean flag for identity cert generation. | [default to true]

### Return type

[**ConsumerDTO**](ConsumerDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConsumer

> DeleteConsumer(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the Consumer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumerAPI.DeleteConsumer(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.DeleteConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the Consumer | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsumerRequest struct via the builder pattern


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


## DeleteConsumerContentOverrides

> []ContentOverrideDTO DeleteConsumerContentOverrides(ctx, consumerUuid).ContentOverrideDTO(contentOverrideDTO).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The ID of the Consumer
	contentOverrideDTO := []openapiclient.ContentOverrideDTO{*openapiclient.NewContentOverrideDTO()} // []ContentOverrideDTO | The list of the content overrides

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.DeleteConsumerContentOverrides(context.Background(), consumerUuid).ContentOverrideDTO(contentOverrideDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.DeleteConsumerContentOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteConsumerContentOverrides`: []ContentOverrideDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.DeleteConsumerContentOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The ID of the Consumer | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsumerContentOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentOverrideDTO** | [**[]ContentOverrideDTO**](ContentOverrideDTO.md) | The list of the content overrides | 

### Return type

[**[]ContentOverrideDTO**](ContentOverrideDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadExistingExport

> *os.File DownloadExistingExport(ctx, consumerUuid, exportId).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the consumer
	exportId := "exportId_example" // string | export ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.DownloadExistingExport(context.Background(), consumerUuid, exportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.DownloadExistingExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadExistingExport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.DownloadExistingExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the consumer | 
**exportId** | **string** | export ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadExistingExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DryBind

> []PoolQuantityDTO DryBind(ctx, consumerUuid).ServiceLevel(serviceLevel).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID
	serviceLevel := "serviceLevel_example" // string | Service level (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.DryBind(context.Background(), consumerUuid).ServiceLevel(serviceLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.DryBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DryBind`: []PoolQuantityDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.DryBind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDryBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceLevel** | **string** | Service level | 

### Return type

[**[]PoolQuantityDTO**](PoolQuantityDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCertificates

> map[string]interface{} ExportCertificates(ctx, consumerUuid).Serials(serials).Execute()





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
	serials := "serials_example" // string | Certificate serials (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.ExportCertificates(context.Background(), consumerUuid).Serials(serials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.ExportCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCertificates`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.ExportCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the consumer to retrieve guest | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **string** | Certificate serials | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportData

> *os.File ExportData(ctx, consumerUuid).CdnLabel(cdnLabel).WebappPrefix(webappPrefix).ApiUrl(apiUrl).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the target consumer.
	cdnLabel := "cdnLabel_example" // string | The label of the target CDN. (optional)
	webappPrefix := "webappPrefix_example" // string | The URL pointing to the manifest's originating web application. (optional)
	apiUrl := "apiUrl_example" // string | The URL pointing to the manifest's originating candlepin API. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.ExportData(context.Background(), consumerUuid).CdnLabel(cdnLabel).WebappPrefix(webappPrefix).ApiUrl(apiUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.ExportData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportData`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.ExportData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the target consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnLabel** | **string** | The label of the target CDN. | 
 **webappPrefix** | **string** | The URL pointing to the manifest&#39;s originating web application. | 
 **apiUrl** | **string** | The URL pointing to the manifest&#39;s originating candlepin API. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportDataAsync

> AsyncJobStatusDTO ExportDataAsync(ctx, consumerUuid).CdnLabel(cdnLabel).WebappPrefix(webappPrefix).ApiUrl(apiUrl).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the target consumer.
	cdnLabel := "cdnLabel_example" // string | The lable of the target CDN. (optional)
	webappPrefix := "webappPrefix_example" // string | The URL pointing to the manifest's originating web application. (optional)
	apiUrl := "apiUrl_example" // string | The URL pointing to the manifest's originating candlepin API. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.ExportDataAsync(context.Background(), consumerUuid).CdnLabel(cdnLabel).WebappPrefix(webappPrefix).ApiUrl(apiUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.ExportDataAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDataAsync`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.ExportDataAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the target consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportDataAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnLabel** | **string** | The lable of the target CDN. | 
 **webappPrefix** | **string** | The URL pointing to the manifest&#39;s originating web application. | 
 **apiUrl** | **string** | The URL pointing to the manifest&#39;s originating candlepin API. | 

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


## GetComplianceStatus

> ComplianceStatusDTO GetComplianceStatus(ctx, consumerUuid).OnDate(onDate).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID
	onDate := "onDate_example" // string | Date to get compliance information for, default is now. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.GetComplianceStatus(context.Background(), consumerUuid).OnDate(onDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.GetComplianceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceStatus`: ComplianceStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.GetComplianceStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onDate** | **string** | Date to get compliance information for, default is now. | 

### Return type

[**ComplianceStatusDTO**](ComplianceStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceStatusList

> map[string]ComplianceStatusDTO GetComplianceStatusList(ctx).Uuid(uuid).Execute()





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
	uuid := []string{"Inner_example"} // []string | Consumers UUIDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.GetComplianceStatusList(context.Background()).Uuid(uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.GetComplianceStatusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplianceStatusList`: map[string]ComplianceStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.GetComplianceStatusList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceStatusListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **[]string** | Consumers UUIDs | 

### Return type

[**map[string]ComplianceStatusDTO**](ComplianceStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsumer

> ConsumerDTO GetConsumer(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.GetConsumer(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.GetConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConsumer`: ConsumerDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.GetConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsumerDTO**](ConsumerDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentAccessBody

> string GetContentAccessBody(ctx, consumerUuid).IfModifiedSince(ifModifiedSince).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the consumer
	ifModifiedSince := "ifModifiedSince_example" // string | Modified date. Accepted format EEE, dd MMM yyyy HH:mm:ss z (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.GetContentAccessBody(context.Background(), consumerUuid).IfModifiedSince(ifModifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.GetContentAccessBody``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentAccessBody`: string
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.GetContentAccessBody`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the consumer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentAccessBodyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **string** | Modified date. Accepted format EEE, dd MMM yyyy HH:mm:ss z | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentAccessForConsumer

> ContentAccessDTO GetContentAccessForConsumer(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.GetContentAccessForConsumer(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.GetContentAccessForConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentAccessForConsumer`: ContentAccessDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.GetContentAccessForConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentAccessForConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentAccessDTO**](ContentAccessDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementCertificateSerials

> []CertificateSerialDTO GetEntitlementCertificateSerials(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.GetEntitlementCertificateSerials(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.GetEntitlementCertificateSerials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementCertificateSerials`: []CertificateSerialDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.GetEntitlementCertificateSerials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementCertificateSerialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CertificateSerialDTO**](CertificateSerialDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuests

> []ConsumerDTOArrayElement GetGuests(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.GetGuests(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.GetGuests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuests`: []ConsumerDTOArrayElement
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.GetGuests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConsumerDTOArrayElement**](ConsumerDTOArrayElement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHost

> ConsumerDTO GetHost(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.GetHost(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.GetHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHost`: ConsumerDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.GetHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsumerDTO**](ConsumerDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnerByConsumerUuid

> OwnerDTO GetOwnerByConsumerUuid(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.GetOwnerByConsumerUuid(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.GetOwnerByConsumerUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnerByConsumerUuid`: OwnerDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.GetOwnerByConsumerUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerByConsumerUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OwnerDTO**](OwnerDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelease

> ReleaseVerDTO GetRelease(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.GetRelease(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.GetRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRelease`: ReleaseVerDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.GetRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseVerDTO**](ReleaseVerDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemPurposeComplianceStatus

> SystemPurposeComplianceStatusDTO GetSystemPurposeComplianceStatus(ctx, consumerUuid).OnDate(onDate).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID
	onDate := "onDate_example" // string | Date to get compliance information for, default is now. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.GetSystemPurposeComplianceStatus(context.Background(), consumerUuid).OnDate(onDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.GetSystemPurposeComplianceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemPurposeComplianceStatus`: SystemPurposeComplianceStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.GetSystemPurposeComplianceStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemPurposeComplianceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onDate** | **string** | Date to get compliance information for, default is now. | 

### Return type

[**SystemPurposeComplianceStatusDTO**](SystemPurposeComplianceStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConsumerContentOverrides

> []ContentOverrideDTO ListConsumerContentOverrides(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The ID of the consumer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.ListConsumerContentOverrides(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.ListConsumerContentOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConsumerContentOverrides`: []ContentOverrideDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.ListConsumerContentOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The ID of the consumer | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConsumerContentOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ContentOverrideDTO**](ContentOverrideDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlements

> []EntitlementDTO ListEntitlements(ctx, consumerUuid).Product(product).Regen(regen).Attribute(attribute).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID
	product := "product_example" // string | ID of a Product (optional)
	regen := true // bool | Boolean flag to regenerate entitlements (optional) (default to true)
	attribute := []string{"Inner_example"} // []string | Attribute filters (optional)
	page := int32(2) // int32 | Page index to return (optional)
	perPage := int32(10) // int32 | Number of items to return per page (optional)
	order := "asc" // string | Direction of ordering (optional)
	sortBy := "name" // string | Property to use for ordering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.ListEntitlements(context.Background(), consumerUuid).Product(product).Regen(regen).Attribute(attribute).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.ListEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntitlements`: []EntitlementDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.ListEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | **string** | ID of a Product | 
 **regen** | **bool** | Boolean flag to regenerate entitlements | [default to true]
 **attribute** | **[]string** | Attribute filters | 
 **page** | **int32** | Page index to return | 
 **perPage** | **int32** | Number of items to return per page | 
 **order** | **string** | Direction of ordering | 
 **sortBy** | **string** | Property to use for ordering | 

### Return type

[**[]EntitlementDTO**](EntitlementDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateEntitlementCertificates

> RegenerateEntitlementCertificates(ctx, consumerUuid).Entitlement(entitlement).LazyRegen(lazyRegen).CleanupEntitlements(cleanupEntitlements).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the Consumer
	entitlement := "entitlement_example" // string | Entitlement ID (optional)
	lazyRegen := true // bool | Lazy regeneration of entitlement certs (optional) (default to true)
	cleanupEntitlements := true // bool | Whether or not to remove unnecessary or unused entitlements for the consumer before regenerating certificates  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumerAPI.RegenerateEntitlementCertificates(context.Background(), consumerUuid).Entitlement(entitlement).LazyRegen(lazyRegen).CleanupEntitlements(cleanupEntitlements).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.RegenerateEntitlementCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the Consumer | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateEntitlementCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitlement** | **string** | Entitlement ID | 
 **lazyRegen** | **bool** | Lazy regeneration of entitlement certs | [default to true]
 **cleanupEntitlements** | **bool** | Whether or not to remove unnecessary or unused entitlements for the consumer before regenerating certificates  | [default to false]

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


## RegenerateIdentityCertificates

> ConsumerDTO RegenerateIdentityCertificates(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | Consumer UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.RegenerateIdentityCertificates(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.RegenerateIdentityCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateIdentityCertificates`: ConsumerDTO
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.RegenerateIdentityCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | Consumer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateIdentityCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsumerDTO**](ConsumerDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDeletionRecord

> RemoveDeletionRecord(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the Consumer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumerAPI.RemoveDeletionRecord(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.RemoveDeletionRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the Consumer | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDeletionRecordRequest struct via the builder pattern


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


## SearchConsumers

> []ConsumerDTOArrayElement SearchConsumers(ctx).Username(username).Type_(type_).Owner(owner).Uuid(uuid).HypervisorId(hypervisorId).RegistrationAuthenticationMethod(registrationAuthenticationMethod).Fact(fact).EnvironmentId(environmentId).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()





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
	username := "username_example" // string | Username (optional)
	type_ := []string{"Inner_example"} // []string | Consumer type (optional)
	owner := "owner_example" // string | Owner key (optional)
	uuid := []string{"Inner_example"} // []string | The UUID of consumers (optional)
	hypervisorId := []string{"Inner_example"} // []string | Hypervisor IDs (optional)
	registrationAuthenticationMethod := "registrationAuthenticationMethod_example" // string | Registration Authentication Method (optional)
	fact := []string{"Inner_example"} // []string | The consumer facts (optional)
	environmentId := "environmentId_example" // string | Environment identifier (optional)
	page := int32(2) // int32 | Page index to return (optional)
	perPage := int32(10) // int32 | Number of items to return per page (optional)
	order := "asc" // string | Direction of ordering (optional)
	sortBy := "name" // string | Property to use for ordering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.SearchConsumers(context.Background()).Username(username).Type_(type_).Owner(owner).Uuid(uuid).HypervisorId(hypervisorId).RegistrationAuthenticationMethod(registrationAuthenticationMethod).Fact(fact).EnvironmentId(environmentId).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.SearchConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchConsumers`: []ConsumerDTOArrayElement
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.SearchConsumers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** | Username | 
 **type_** | **[]string** | Consumer type | 
 **owner** | **string** | Owner key | 
 **uuid** | **[]string** | The UUID of consumers | 
 **hypervisorId** | **[]string** | Hypervisor IDs | 
 **registrationAuthenticationMethod** | **string** | Registration Authentication Method | 
 **fact** | **[]string** | The consumer facts | 
 **environmentId** | **string** | Environment identifier | 
 **page** | **int32** | Page index to return | 
 **perPage** | **int32** | Number of items to return per page | 
 **order** | **string** | Direction of ordering | 
 **sortBy** | **string** | Property to use for ordering | 

### Return type

[**[]ConsumerDTOArrayElement**](ConsumerDTOArrayElement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindAll

> DeleteResult UnbindAll(ctx, consumerUuid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the Consumer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumerAPI.UnbindAll(context.Background(), consumerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.UnbindAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbindAll`: DeleteResult
	fmt.Fprintf(os.Stdout, "Response from `ConsumerAPI.UnbindAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the Consumer | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResult**](DeleteResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindByEntitlementId

> UnbindByEntitlementId(ctx, consumerUuid, dbid).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the Consumer
	dbid := "dbid_example" // string | Entitlement ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumerAPI.UnbindByEntitlementId(context.Background(), consumerUuid, dbid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.UnbindByEntitlementId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the Consumer | 
**dbid** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindByEntitlementIdRequest struct via the builder pattern


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


## UnbindByPool

> UnbindByPool(ctx, consumerUuid, poolId).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The ID of the Consumer
	poolId := "poolId_example" // string | The ID of the Consumer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumerAPI.UnbindByPool(context.Background(), consumerUuid, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.UnbindByPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The ID of the Consumer | 
**poolId** | **string** | The ID of the Consumer | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindByPoolRequest struct via the builder pattern


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


## UnbindBySerial

> UnbindBySerial(ctx, consumerUuid, serial).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the Consumer
	serial := int64(789) // int64 | certificate serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumerAPI.UnbindBySerial(context.Background(), consumerUuid, serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.UnbindBySerial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the Consumer | 
**serial** | **int64** | certificate serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindBySerialRequest struct via the builder pattern


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


## UpdateConsumer

> UpdateConsumer(ctx, consumerUuid).ConsumerDTO(consumerDTO).Execute()





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
	consumerUuid := "consumerUuid_example" // string | The UUID of the Consumer
	consumerDTO := *openapiclient.NewConsumerDTO() // ConsumerDTO | Consumer to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumerAPI.UpdateConsumer(context.Background(), consumerUuid).ConsumerDTO(consumerDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumerAPI.UpdateConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerUuid** | **string** | The UUID of the Consumer | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumerDTO** | [**ConsumerDTO**](ConsumerDTO.md) | Consumer to be updated | 

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

