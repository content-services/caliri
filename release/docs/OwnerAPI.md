# \OwnerAPI

All URIs are relative to */candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Claim**](OwnerAPI.md#Claim) | **Put** /owners/{anonymous_owner_key}/claim | 
[**CountConsumers**](OwnerAPI.md#CountConsumers) | **Get** /owners/{owner_key}/consumers/count | 
[**CreateActivationKey**](OwnerAPI.md#CreateActivationKey) | **Post** /owners/{owner_key}/activation_keys | 
[**CreateEnv**](OwnerAPI.md#CreateEnv) | **Post** /owners/{owner_key}/environments | 
[**CreateOwner**](OwnerAPI.md#CreateOwner) | **Post** /owners | 
[**CreatePool**](OwnerAPI.md#CreatePool) | **Post** /owners/{owner_key}/pools | 
[**CreateUeberCertificate**](OwnerAPI.md#CreateUeberCertificate) | **Post** /owners/{owner_key}/uebercert | 
[**DeleteLogLevel**](OwnerAPI.md#DeleteLogLevel) | **Delete** /owners/{owner_key}/log | 
[**DeleteOwner**](OwnerAPI.md#DeleteOwner) | **Delete** /owners/{owner_key} | 
[**GetConsumersSyspurpose**](OwnerAPI.md#GetConsumersSyspurpose) | **Get** /owners/{owner_key}/consumers_system_purpose | 
[**GetHypervisors**](OwnerAPI.md#GetHypervisors) | **Get** /owners/{owner_key}/hypervisors | 
[**GetImports**](OwnerAPI.md#GetImports) | **Get** /owners/{owner_key}/imports | 
[**GetOwner**](OwnerAPI.md#GetOwner) | **Get** /owners/{owner_key} | 
[**GetOwnerContentAccess**](OwnerAPI.md#GetOwnerContentAccess) | **Get** /owners/{owner_key}/content_access | 
[**GetOwnerInfo**](OwnerAPI.md#GetOwnerInfo) | **Get** /owners/{owner_key}/info | 
[**GetOwnerSubscriptions**](OwnerAPI.md#GetOwnerSubscriptions) | **Get** /owners/{owner_key}/subscriptions | 
[**GetSyspurpose**](OwnerAPI.md#GetSyspurpose) | **Get** /owners/{owner_key}/system_purpose | 
[**GetUeberCertificate**](OwnerAPI.md#GetUeberCertificate) | **Get** /owners/{owner_key}/uebercert | 
[**GetUpstreamConsumers**](OwnerAPI.md#GetUpstreamConsumers) | **Get** /owners/{owner_key}/upstream_consumers | 
[**HealEntire**](OwnerAPI.md#HealEntire) | **Post** /owners/{owner_key}/entitlements | 
[**ImportManifestAsync**](OwnerAPI.md#ImportManifestAsync) | **Post** /owners/{owner_key}/imports/async | 
[**ListConsumers**](OwnerAPI.md#ListConsumers) | **Get** /owners/{owner_key}/consumers | 
[**ListEnvironments**](OwnerAPI.md#ListEnvironments) | **Get** /owners/{owner_key}/environments | 
[**ListOwnerPools**](OwnerAPI.md#ListOwnerPools) | **Get** /owners/{owner_key}/pools | 
[**ListOwners**](OwnerAPI.md#ListOwners) | **Get** /owners | 
[**OwnerActivationKeys**](OwnerAPI.md#OwnerActivationKeys) | **Get** /owners/{owner_key}/activation_keys | 
[**OwnerEntitlements**](OwnerAPI.md#OwnerEntitlements) | **Get** /owners/{owner_key}/entitlements | 
[**OwnerServiceLevels**](OwnerAPI.md#OwnerServiceLevels) | **Get** /owners/{owner_key}/servicelevels | 
[**RefreshPools**](OwnerAPI.md#RefreshPools) | **Put** /owners/{owner_key}/subscriptions | 
[**SetLogLevel**](OwnerAPI.md#SetLogLevel) | **Put** /owners/{owner_key}/log | 
[**UndoImports**](OwnerAPI.md#UndoImports) | **Delete** /owners/{owner_key}/imports | 
[**UpdateOwner**](OwnerAPI.md#UpdateOwner) | **Put** /owners/{owner_key} | 
[**UpdatePool**](OwnerAPI.md#UpdatePool) | **Put** /owners/{owner_key}/pools | 



## Claim

> AsyncJobStatusDTO Claim(ctx, anonymousOwnerKey).ClaimantOwner(claimantOwner).Execute()





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
	anonymousOwnerKey := "anonymousOwnerKey_example" // string | The key of the owner
	claimantOwner := *openapiclient.NewClaimantOwner("ClaimantOwnerKey_example") // ClaimantOwner | Claimant owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.Claim(context.Background(), anonymousOwnerKey).ClaimantOwner(claimantOwner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.Claim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Claim`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.Claim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**anonymousOwnerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimantOwner** | [**ClaimantOwner**](ClaimantOwner.md) | Claimant owner | 

### Return type

[**AsyncJobStatusDTO**](AsyncJobStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountConsumers

> int32 CountConsumers(ctx, ownerKey).Username(username).Type_(type_).Uuid(uuid).HypervisorId(hypervisorId).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	username := "username_example" // string | The username of the consumer (optional)
	type_ := []string{"Inner_example"} // []string | The consumer type (optional)
	uuid := []string{"Inner_example"} // []string | The UUID of consumers (optional)
	hypervisorId := []string{"Inner_example"} // []string | The hypervisor IDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.CountConsumers(context.Background(), ownerKey).Username(username).Type_(type_).Uuid(uuid).HypervisorId(hypervisorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.CountConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountConsumers`: int32
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.CountConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of the consumer | 
 **type_** | **[]string** | The consumer type | 
 **uuid** | **[]string** | The UUID of consumers | 
 **hypervisorId** | **[]string** | The hypervisor IDs | 

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


## CreateActivationKey

> ActivationKeyDTO CreateActivationKey(ctx, ownerKey).ActivationKeyDTO(activationKeyDTO).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	activationKeyDTO := *openapiclient.NewActivationKeyDTO() // ActivationKeyDTO | Activation key to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.CreateActivationKey(context.Background(), ownerKey).ActivationKeyDTO(activationKeyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.CreateActivationKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateActivationKey`: ActivationKeyDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.CreateActivationKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateActivationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activationKeyDTO** | [**ActivationKeyDTO**](ActivationKeyDTO.md) | Activation key to be created | 

### Return type

[**ActivationKeyDTO**](ActivationKeyDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnv

> EnvironmentDTO CreateEnv(ctx, ownerKey).EnvironmentDTO(environmentDTO).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	environmentDTO := *openapiclient.NewEnvironmentDTO() // EnvironmentDTO | Environment to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.CreateEnv(context.Background(), ownerKey).EnvironmentDTO(environmentDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.CreateEnv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnv`: EnvironmentDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.CreateEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentDTO** | [**EnvironmentDTO**](EnvironmentDTO.md) | Environment to be created | 

### Return type

[**EnvironmentDTO**](EnvironmentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOwner

> OwnerDTO CreateOwner(ctx).OwnerDTO(ownerDTO).Execute()





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
	ownerDTO := *openapiclient.NewOwnerDTO() // OwnerDTO | Owner to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.CreateOwner(context.Background()).OwnerDTO(ownerDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.CreateOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOwner`: OwnerDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.CreateOwner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerDTO** | [**OwnerDTO**](OwnerDTO.md) | Owner to be created | 

### Return type

[**OwnerDTO**](OwnerDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePool

> PoolDTO CreatePool(ctx, ownerKey).PoolDTO(poolDTO).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	poolDTO := *openapiclient.NewPoolDTO() // PoolDTO | A pool to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.CreatePool(context.Background(), ownerKey).PoolDTO(poolDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.CreatePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePool`: PoolDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.CreatePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **poolDTO** | [**PoolDTO**](PoolDTO.md) | A pool to be created | 

### Return type

[**PoolDTO**](PoolDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUeberCertificate

> UeberCertificateDTO CreateUeberCertificate(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.CreateUeberCertificate(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.CreateUeberCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUeberCertificate`: UeberCertificateDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.CreateUeberCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUeberCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UeberCertificateDTO**](UeberCertificateDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogLevel

> DeleteLogLevel(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The owner key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OwnerAPI.DeleteLogLevel(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.DeleteLogLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The owner key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogLevelRequest struct via the builder pattern


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


## DeleteOwner

> DeleteOwner(ctx, ownerKey).Revoke(revoke).Force(force).Execute()





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
	ownerKey := "ownerKey_example" // string | The owner key
	revoke := true // bool | Boolean value to revoke an owner (optional) (default to true)
	force := true // bool | Boolean value to forcefully delete an owner. This is not in use. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OwnerAPI.DeleteOwner(context.Background(), ownerKey).Revoke(revoke).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.DeleteOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The owner key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revoke** | **bool** | Boolean value to revoke an owner | [default to true]
 **force** | **bool** | Boolean value to forcefully delete an owner. This is not in use. | [default to false]

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


## GetConsumersSyspurpose

> SystemPurposeAttributesDTO GetConsumersSyspurpose(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.GetConsumersSyspurpose(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.GetConsumersSyspurpose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConsumersSyspurpose`: SystemPurposeAttributesDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.GetConsumersSyspurpose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumersSyspurposeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemPurposeAttributesDTO**](SystemPurposeAttributesDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHypervisors

> []ConsumerDTOArrayElement GetHypervisors(ctx, ownerKey).HypervisorId(hypervisorId).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	hypervisorId := []string{"Inner_example"} // []string | The list of hypervisor Ids (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.GetHypervisors(context.Background(), ownerKey).HypervisorId(hypervisorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.GetHypervisors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHypervisors`: []ConsumerDTOArrayElement
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.GetHypervisors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHypervisorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hypervisorId** | **[]string** | The list of hypervisor Ids | 

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


## GetImports

> []ImportRecordDTO GetImports(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.GetImports(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.GetImports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImports`: []ImportRecordDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.GetImports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ImportRecordDTO**](ImportRecordDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwner

> OwnerDTO GetOwner(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.GetOwner(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.GetOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwner`: OwnerDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.GetOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerRequest struct via the builder pattern


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


## GetOwnerContentAccess

> ContentAccessDTO GetOwnerContentAccess(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.GetOwnerContentAccess(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.GetOwnerContentAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnerContentAccess`: ContentAccessDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.GetOwnerContentAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerContentAccessRequest struct via the builder pattern


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


## GetOwnerInfo

> OwnerInfo GetOwnerInfo(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.GetOwnerInfo(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.GetOwnerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnerInfo`: OwnerInfo
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.GetOwnerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OwnerInfo**](OwnerInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnerSubscriptions

> []SubscriptionDTO GetOwnerSubscriptions(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.GetOwnerSubscriptions(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.GetOwnerSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnerSubscriptions`: []SubscriptionDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.GetOwnerSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SubscriptionDTO**](SubscriptionDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyspurpose

> SystemPurposeAttributesDTO GetSyspurpose(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.GetSyspurpose(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.GetSyspurpose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyspurpose`: SystemPurposeAttributesDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.GetSyspurpose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyspurposeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemPurposeAttributesDTO**](SystemPurposeAttributesDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUeberCertificate

> UeberCertificateDTO GetUeberCertificate(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.GetUeberCertificate(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.GetUeberCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUeberCertificate`: UeberCertificateDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.GetUeberCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUeberCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UeberCertificateDTO**](UeberCertificateDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpstreamConsumers

> []UpstreamConsumerDTOArrayElement GetUpstreamConsumers(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.GetUpstreamConsumers(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.GetUpstreamConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpstreamConsumers`: []UpstreamConsumerDTOArrayElement
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.GetUpstreamConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpstreamConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UpstreamConsumerDTOArrayElement**](UpstreamConsumerDTOArrayElement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealEntire

> AsyncJobStatusDTO HealEntire(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.HealEntire(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.HealEntire``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealEntire`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.HealEntire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiHealEntireRequest struct via the builder pattern


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


## ImportManifestAsync

> AsyncJobStatusDTO ImportManifestAsync(ctx, ownerKey).Force(force).Input(input).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	force := []string{"Inner_example"} // []string | This is used to override the conflicts (optional)
	input := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.ImportManifestAsync(context.Background(), ownerKey).Force(force).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.ImportManifestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportManifestAsync`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.ImportManifestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportManifestAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **[]string** | This is used to override the conflicts | 
 **input** | ***os.File** |  | 

### Return type

[**AsyncJobStatusDTO**](AsyncJobStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConsumers

> []ConsumerDTOArrayElement ListConsumers(ctx, ownerKey).Username(username).Type_(type_).Uuid(uuid).HypervisorId(hypervisorId).Fact(fact).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	username := "username_example" // string | The username of the consumer (optional)
	type_ := []string{"Inner_example"} // []string | The consumer type (optional)
	uuid := []string{"Inner_example"} // []string | The UUID of consumers (optional)
	hypervisorId := []string{"Inner_example"} // []string | The hypervisor IDs (optional)
	fact := []string{"Inner_example"} // []string | The consumer facts (optional)
	page := int32(2) // int32 | Page index to return (optional)
	perPage := int32(10) // int32 | Number of items to return per page (optional)
	order := "asc" // string | Direction of ordering (optional)
	sortBy := "name" // string | Property to use for ordering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.ListConsumers(context.Background(), ownerKey).Username(username).Type_(type_).Uuid(uuid).HypervisorId(hypervisorId).Fact(fact).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.ListConsumers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConsumers`: []ConsumerDTOArrayElement
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.ListConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of the consumer | 
 **type_** | **[]string** | The consumer type | 
 **uuid** | **[]string** | The UUID of consumers | 
 **hypervisorId** | **[]string** | The hypervisor IDs | 
 **fact** | **[]string** | The consumer facts | 
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


## ListEnvironments

> []EnvironmentDTO ListEnvironments(ctx, ownerKey).Name(name).Type_(type_).ListAll(listAll).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	name := "name_example" // string | The name of the environment (optional)
	type_ := []string{"Inner_example"} // []string | The type of the environments (optional)
	listAll := true // bool | List all boolean for environments of all types (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.ListEnvironments(context.Background(), ownerKey).Name(name).Type_(type_).ListAll(listAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.ListEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironments`: []EnvironmentDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.ListEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the environment | 
 **type_** | **[]string** | The type of the environments | 
 **listAll** | **bool** | List all boolean for environments of all types | [default to false]

### Return type

[**[]EnvironmentDTO**](EnvironmentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOwnerPools

> []PoolDTO ListOwnerPools(ctx, ownerKey).Consumer(consumer).ActivationKey(activationKey).Product(product).Subscription(subscription).Listall(listall).Activeon(activeon).Matches(matches).Attribute(attribute).AddFuture(addFuture).OnlyFuture(onlyFuture).After(after).Poolid(poolid).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	consumer := "consumer_example" // string | The consumer UUID (optional)
	activationKey := "activationKey_example" // string | The activation key name (optional)
	product := "product_example" // string | The product id (optional)
	subscription := "subscription_example" // string | The subscription id (optional)
	listall := true // bool | Includes pools which are not applicable based on some of the consumer's current system facts (i.e. system architecture, sockets, cores, ram, storage_band, instance_multiplier), but will still filter pools based on other factors, such as pools restricted to a consumer of specific types, usernames, or uuids; pools restricted to guests of specific hosts; non-multi-entitlement pools that the consumer has already attached; unmapped guest pools for which the consumer is ineligible; and expired pools.  (optional) (default to false)
	activeon := time.Now() // time.Time | Active on date (optional)
	matches := []string{"Inner_example"} // []string | Find pools matching the given pattern in a variety of fields * and ? wildcards are supported; may be specified multiple times  (optional)
	attribute := []string{"Inner_example"} // []string | The attributes to return based on the specified types (optional)
	addFuture := true // bool | When set to true, it will add future dated pools to the result, based on the activeon date  (optional) (default to false)
	onlyFuture := true // bool | When set to true, it will return only future dated pools to the result, based on the activeon date  (optional) (default to false)
	after := time.Now() // time.Time | Will only return pools with a start date after the supplied date. Overrides the activeOn date  (optional)
	poolid := []string{"Inner_example"} // []string | One or more pool IDs to use to filter the output; only pools with IDs matching those provided will be returned; may be specified multiple times  (optional)
	page := int32(2) // int32 | Page index to return (optional)
	perPage := int32(10) // int32 | Number of items to return per page (optional)
	order := "asc" // string | Direction of ordering (optional)
	sortBy := "name" // string | Property to use for ordering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.ListOwnerPools(context.Background(), ownerKey).Consumer(consumer).ActivationKey(activationKey).Product(product).Subscription(subscription).Listall(listall).Activeon(activeon).Matches(matches).Attribute(attribute).AddFuture(addFuture).OnlyFuture(onlyFuture).After(after).Poolid(poolid).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.ListOwnerPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOwnerPools`: []PoolDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.ListOwnerPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOwnerPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumer** | **string** | The consumer UUID | 
 **activationKey** | **string** | The activation key name | 
 **product** | **string** | The product id | 
 **subscription** | **string** | The subscription id | 
 **listall** | **bool** | Includes pools which are not applicable based on some of the consumer&#39;s current system facts (i.e. system architecture, sockets, cores, ram, storage_band, instance_multiplier), but will still filter pools based on other factors, such as pools restricted to a consumer of specific types, usernames, or uuids; pools restricted to guests of specific hosts; non-multi-entitlement pools that the consumer has already attached; unmapped guest pools for which the consumer is ineligible; and expired pools.  | [default to false]
 **activeon** | **time.Time** | Active on date | 
 **matches** | **[]string** | Find pools matching the given pattern in a variety of fields * and ? wildcards are supported; may be specified multiple times  | 
 **attribute** | **[]string** | The attributes to return based on the specified types | 
 **addFuture** | **bool** | When set to true, it will add future dated pools to the result, based on the activeon date  | [default to false]
 **onlyFuture** | **bool** | When set to true, it will return only future dated pools to the result, based on the activeon date  | [default to false]
 **after** | **time.Time** | Will only return pools with a start date after the supplied date. Overrides the activeOn date  | 
 **poolid** | **[]string** | One or more pool IDs to use to filter the output; only pools with IDs matching those provided will be returned; may be specified multiple times  | 
 **page** | **int32** | Page index to return | 
 **perPage** | **int32** | Number of items to return per page | 
 **order** | **string** | Direction of ordering | 
 **sortBy** | **string** | Property to use for ordering | 

### Return type

[**[]PoolDTO**](PoolDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOwners

> []OwnerDTO ListOwners(ctx).Key(key).Execute()





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
	key := "key_example" // string | The owner key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.ListOwners(context.Background()).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.ListOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOwners`: []OwnerDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.ListOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | The owner key | 

### Return type

[**[]OwnerDTO**](OwnerDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OwnerActivationKeys

> []ActivationKeyDTO OwnerActivationKeys(ctx, ownerKey).Name(name).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	name := "name_example" // string | The name of the activation key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.OwnerActivationKeys(context.Background(), ownerKey).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.OwnerActivationKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OwnerActivationKeys`: []ActivationKeyDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.OwnerActivationKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiOwnerActivationKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the activation key | 

### Return type

[**[]ActivationKeyDTO**](ActivationKeyDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OwnerEntitlements

> []EntitlementDTO OwnerEntitlements(ctx, ownerKey).Product(product).Attribute(attribute).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	product := "product_example" // string | The product id (optional)
	attribute := []string{"Inner_example"} // []string | Attribute filters (optional)
	page := int32(2) // int32 | Page index to return (optional)
	perPage := int32(10) // int32 | Number of items to return per page (optional)
	order := "asc" // string | Direction of ordering (optional)
	sortBy := "name" // string | Property to use for ordering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.OwnerEntitlements(context.Background(), ownerKey).Product(product).Attribute(attribute).Page(page).PerPage(perPage).Order(order).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.OwnerEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OwnerEntitlements`: []EntitlementDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.OwnerEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiOwnerEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | **string** | The product id | 
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


## OwnerServiceLevels

> []string OwnerServiceLevels(ctx, ownerKey).Exempt(exempt).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	exempt := "exempt_example" // string | exempt (optional) (default to "false")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.OwnerServiceLevels(context.Background(), ownerKey).Exempt(exempt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.OwnerServiceLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OwnerServiceLevels`: []string
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.OwnerServiceLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiOwnerServiceLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exempt** | **string** | exempt | [default to &quot;false&quot;]

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


## RefreshPools

> AsyncJobStatusDTO RefreshPools(ctx, ownerKey).AutoCreateOwner(autoCreateOwner).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	autoCreateOwner := true // bool | Specify whether or not to create an owner automatically. Default is false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.RefreshPools(context.Background(), ownerKey).AutoCreateOwner(autoCreateOwner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.RefreshPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshPools`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.RefreshPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoCreateOwner** | **bool** | Specify whether or not to create an owner automatically. Default is false. | [default to false]

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


## SetLogLevel

> OwnerDTO SetLogLevel(ctx, ownerKey).Level(level).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	level := "level_example" // string | The log level (optional) (default to "DEBUG")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.SetLogLevel(context.Background(), ownerKey).Level(level).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.SetLogLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetLogLevel`: OwnerDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.SetLogLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLogLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **level** | **string** | The log level | [default to &quot;DEBUG&quot;]

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


## UndoImports

> AsyncJobStatusDTO UndoImports(ctx, ownerKey).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.UndoImports(context.Background(), ownerKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.UndoImports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UndoImports`: AsyncJobStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.UndoImports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndoImportsRequest struct via the builder pattern


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


## UpdateOwner

> OwnerDTO UpdateOwner(ctx, ownerKey).OwnerDTO(ownerDTO).Execute()





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
	ownerKey := "ownerKey_example" // string | The key of the owner
	ownerDTO := *openapiclient.NewOwnerDTO() // OwnerDTO | Owner to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OwnerAPI.UpdateOwner(context.Background(), ownerKey).OwnerDTO(ownerDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.UpdateOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOwner`: OwnerDTO
	fmt.Fprintf(os.Stdout, "Response from `OwnerAPI.UpdateOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | The key of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ownerDTO** | [**OwnerDTO**](OwnerDTO.md) | Owner to be updated | 

### Return type

[**OwnerDTO**](OwnerDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePool

> UpdatePool(ctx, ownerKey).PoolDTO(poolDTO).Execute()





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
	ownerKey := "ownerKey_example" // string | Owner key
	poolDTO := *openapiclient.NewPoolDTO() // PoolDTO | A pool to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OwnerAPI.UpdatePool(context.Background(), ownerKey).PoolDTO(poolDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OwnerAPI.UpdatePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerKey** | **string** | Owner key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **poolDTO** | [**PoolDTO**](PoolDTO.md) | A pool to be updated | 

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

