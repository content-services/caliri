/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.16
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


// GuestIdsAPIService GuestIdsAPI service
type GuestIdsAPIService service

type ApiDeleteGuestRequest struct {
	ctx context.Context
	ApiService *GuestIdsAPIService
	consumerUuid string
	guestId string
	unregister *bool
}

// Unregister the consumer
func (r ApiDeleteGuestRequest) Unregister(unregister bool) ApiDeleteGuestRequest {
	r.unregister = &unregister
	return r
}

func (r ApiDeleteGuestRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGuestExecute(r)
}

/*
DeleteGuest Method for DeleteGuest

Removes the guest from the consumer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param consumerUuid The UUID of the consumer who owns or hosts the guest
 @param guestId The UUID of the guest to be deleted
 @return ApiDeleteGuestRequest
*/
func (a *GuestIdsAPIService) DeleteGuest(ctx context.Context, consumerUuid string, guestId string) ApiDeleteGuestRequest {
	return ApiDeleteGuestRequest{
		ApiService: a,
		ctx: ctx,
		consumerUuid: consumerUuid,
		guestId: guestId,
	}
}

// Execute executes the request
func (a *GuestIdsAPIService) DeleteGuestExecute(r ApiDeleteGuestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestIdsAPIService.DeleteGuest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/consumers/{consumer_uuid}/guestids/{guest_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumer_uuid"+"}", url.PathEscape(parameterValueToString(r.consumerUuid, "consumerUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"guest_id"+"}", url.PathEscape(parameterValueToString(r.guestId, "guestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unregister == nil {
		return nil, reportError("unregister is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "unregister", r.unregister, "form", "")
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 410 {
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetGuestIdRequest struct {
	ctx context.Context
	ApiService *GuestIdsAPIService
	consumerUuid string
	guestId string
}

func (r ApiGetGuestIdRequest) Execute() (*GuestIdDTO, *http.Response, error) {
	return r.ApiService.GetGuestIdExecute(r)
}

/*
GetGuestId Method for GetGuestId

Retrieves a single Guest by its consumer and the guest UUID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param consumerUuid The UUID of the consumer to retrieve guest
 @param guestId The UUID of the guest to retrieve
 @return ApiGetGuestIdRequest
*/
func (a *GuestIdsAPIService) GetGuestId(ctx context.Context, consumerUuid string, guestId string) ApiGetGuestIdRequest {
	return ApiGetGuestIdRequest{
		ApiService: a,
		ctx: ctx,
		consumerUuid: consumerUuid,
		guestId: guestId,
	}
}

// Execute executes the request
//  @return GuestIdDTO
func (a *GuestIdsAPIService) GetGuestIdExecute(r ApiGetGuestIdRequest) (*GuestIdDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GuestIdDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestIdsAPIService.GetGuestId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/consumers/{consumer_uuid}/guestids/{guest_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumer_uuid"+"}", url.PathEscape(parameterValueToString(r.consumerUuid, "consumerUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"guest_id"+"}", url.PathEscape(parameterValueToString(r.guestId, "guestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetGuestIdsRequest struct {
	ctx context.Context
	ApiService *GuestIdsAPIService
	consumerUuid string
}

func (r ApiGetGuestIdsRequest) Execute() ([]GuestIdDTOArrayElement, *http.Response, error) {
	return r.ApiService.GetGuestIdsExecute(r)
}

/*
GetGuestIds Method for GetGuestIds

Retrieves the list of a Consumer's Guests

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param consumerUuid
 @return ApiGetGuestIdsRequest
*/
func (a *GuestIdsAPIService) GetGuestIds(ctx context.Context, consumerUuid string) ApiGetGuestIdsRequest {
	return ApiGetGuestIdsRequest{
		ApiService: a,
		ctx: ctx,
		consumerUuid: consumerUuid,
	}
}

// Execute executes the request
//  @return []GuestIdDTOArrayElement
func (a *GuestIdsAPIService) GetGuestIdsExecute(r ApiGetGuestIdsRequest) ([]GuestIdDTOArrayElement, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GuestIdDTOArrayElement
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestIdsAPIService.GetGuestIds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/consumers/{consumer_uuid}/guestids"
	localVarPath = strings.Replace(localVarPath, "{"+"consumer_uuid"+"}", url.PathEscape(parameterValueToString(r.consumerUuid, "consumerUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiUpdateGuestRequest struct {
	ctx context.Context
	ApiService *GuestIdsAPIService
	consumerUuid string
	guestId string
	guestIdDTO *GuestIdDTO
}

// Updated guest data
func (r ApiUpdateGuestRequest) GuestIdDTO(guestIdDTO GuestIdDTO) ApiUpdateGuestRequest {
	r.guestIdDTO = &guestIdDTO
	return r
}

func (r ApiUpdateGuestRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateGuestExecute(r)
}

/*
UpdateGuest Method for UpdateGuest

Updates a single Guest on a Consumer. Allows virt-who to avoid uploading an entire list of guests

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param consumerUuid The UUID of the consumer who owns or hosts the guest
 @param guestId The UUID of the guest to be updated
 @return ApiUpdateGuestRequest
*/
func (a *GuestIdsAPIService) UpdateGuest(ctx context.Context, consumerUuid string, guestId string) ApiUpdateGuestRequest {
	return ApiUpdateGuestRequest{
		ApiService: a,
		ctx: ctx,
		consumerUuid: consumerUuid,
		guestId: guestId,
	}
}

// Execute executes the request
func (a *GuestIdsAPIService) UpdateGuestExecute(r ApiUpdateGuestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestIdsAPIService.UpdateGuest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/consumers/{consumer_uuid}/guestids/{guest_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"consumer_uuid"+"}", url.PathEscape(parameterValueToString(r.consumerUuid, "consumerUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"guest_id"+"}", url.PathEscape(parameterValueToString(r.guestId, "guestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.guestIdDTO == nil {
		return nil, reportError("guestIdDTO is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.guestIdDTO
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateGuestsRequest struct {
	ctx context.Context
	ApiService *GuestIdsAPIService
	consumerUuid string
	guestIdDTO *[]GuestIdDTO
}

// The list of the guests to be updated
func (r ApiUpdateGuestsRequest) GuestIdDTO(guestIdDTO []GuestIdDTO) ApiUpdateGuestsRequest {
	r.guestIdDTO = &guestIdDTO
	return r
}

func (r ApiUpdateGuestsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateGuestsExecute(r)
}

/*
UpdateGuests Method for UpdateGuests

Updates the list of guests on a consumer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param consumerUuid The UUID of the consumer who owns or hosts the guest
 @return ApiUpdateGuestsRequest
*/
func (a *GuestIdsAPIService) UpdateGuests(ctx context.Context, consumerUuid string) ApiUpdateGuestsRequest {
	return ApiUpdateGuestsRequest{
		ApiService: a,
		ctx: ctx,
		consumerUuid: consumerUuid,
	}
}

// Execute executes the request
func (a *GuestIdsAPIService) UpdateGuestsExecute(r ApiUpdateGuestsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestIdsAPIService.UpdateGuests")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/consumers/{consumer_uuid}/guestids"
	localVarPath = strings.Replace(localVarPath, "{"+"consumer_uuid"+"}", url.PathEscape(parameterValueToString(r.consumerUuid, "consumerUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.guestIdDTO == nil {
		return nil, reportError("guestIdDTO is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.guestIdDTO
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ExceptionMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
