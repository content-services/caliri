/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// HypervisorsAPIService HypervisorsAPI service
type HypervisorsAPIService service

type ApiHypervisorHeartbeatUpdateRequest struct {
	ctx context.Context
	ApiService *HypervisorsAPIService
	owner string
	reporterId *string
}

func (r ApiHypervisorHeartbeatUpdateRequest) ReporterId(reporterId string) ApiHypervisorHeartbeatUpdateRequest {
	r.reporterId = &reporterId
	return r
}

func (r ApiHypervisorHeartbeatUpdateRequest) Execute() (*AsyncJobStatusDTO, *http.Response, error) {
	return r.ApiService.HypervisorHeartbeatUpdateExecute(r)
}

/*
HypervisorHeartbeatUpdate Method for HypervisorHeartbeatUpdate

Updates last check in date of all consumers of the given reporterId.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner Owner key
 @return ApiHypervisorHeartbeatUpdateRequest
*/
func (a *HypervisorsAPIService) HypervisorHeartbeatUpdate(ctx context.Context, owner string) ApiHypervisorHeartbeatUpdateRequest {
	return ApiHypervisorHeartbeatUpdateRequest{
		ApiService: a,
		ctx: ctx,
		owner: owner,
	}
}

// Execute executes the request
//  @return AsyncJobStatusDTO
func (a *HypervisorsAPIService) HypervisorHeartbeatUpdateExecute(r ApiHypervisorHeartbeatUpdateRequest) (*AsyncJobStatusDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncJobStatusDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HypervisorsAPIService.HypervisorHeartbeatUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/hypervisors/{owner}/heartbeat"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterValueToString(r.owner, "owner")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.reporterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reporter_id", r.reporterId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiHypervisorUpdateAsyncRequest struct {
	ctx context.Context
	ApiService *HypervisorsAPIService
	owner string
	createMissing *bool
	reporterId *string
	body *string
}

// Specify whether or not to create missing hypervisors. Default is true.  If false is specified, hypervisorIds that are not found will result in failed entries in the resulting HypervisorCheckInResult.
func (r ApiHypervisorUpdateAsyncRequest) CreateMissing(createMissing bool) ApiHypervisorUpdateAsyncRequest {
	r.createMissing = &createMissing
	return r
}

func (r ApiHypervisorUpdateAsyncRequest) ReporterId(reporterId string) ApiHypervisorUpdateAsyncRequest {
	r.reporterId = &reporterId
	return r
}

// Host and Guest mapping details
func (r ApiHypervisorUpdateAsyncRequest) Body(body string) ApiHypervisorUpdateAsyncRequest {
	r.body = &body
	return r
}

func (r ApiHypervisorUpdateAsyncRequest) Execute() (*AsyncJobStatusDTO, *http.Response, error) {
	return r.ApiService.HypervisorUpdateAsyncExecute(r)
}

/*
HypervisorUpdateAsync Method for HypervisorUpdateAsync

Creates or Updates the list of Hypervisor hosts Allows agents such as virt-who to update hosts'
information . This is typically used when a host is unable to register to candlepin via subscription
manager. In situations where consumers already exist it is probably best not to allow creation of new
hypervisor consumers.  Most consumers do not have a hypervisorId attribute, so that should be added
manually when necessary by the management environment. Default is true. If false is specified,
hypervisorIds that are not found will result in a failed state of the job.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner Owner key
 @return ApiHypervisorUpdateAsyncRequest
*/
func (a *HypervisorsAPIService) HypervisorUpdateAsync(ctx context.Context, owner string) ApiHypervisorUpdateAsyncRequest {
	return ApiHypervisorUpdateAsyncRequest{
		ApiService: a,
		ctx: ctx,
		owner: owner,
	}
}

// Execute executes the request
//  @return AsyncJobStatusDTO
func (a *HypervisorsAPIService) HypervisorUpdateAsyncExecute(r ApiHypervisorUpdateAsyncRequest) (*AsyncJobStatusDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncJobStatusDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HypervisorsAPIService.HypervisorUpdateAsync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/hypervisors/{owner}"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterValueToString(r.owner, "owner")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createMissing != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "create_missing", r.createMissing, "form", "")
	} else {
		var defaultValue bool = true
		r.createMissing = &defaultValue
	}
	if r.reporterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reporter_id", r.reporterId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
