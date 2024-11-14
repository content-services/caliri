/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.19
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
	"reflect"
)


// OwnerContentAPIService OwnerContentAPI service
type OwnerContentAPIService service

type ApiCreateContentRequest struct {
	ctx context.Context
	ApiService *OwnerContentAPIService
	ownerKey string
	contentDTO *ContentDTO
}

// Content to be created
func (r ApiCreateContentRequest) ContentDTO(contentDTO ContentDTO) ApiCreateContentRequest {
	r.contentDTO = &contentDTO
	return r
}

func (r ApiCreateContentRequest) Execute() (*ContentDTO, *http.Response, error) {
	return r.ApiService.CreateContentExecute(r)
}

/*
CreateContent Method for CreateContent

Creates a content for owner

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerKey The key of the owner
 @return ApiCreateContentRequest
*/
func (a *OwnerContentAPIService) CreateContent(ctx context.Context, ownerKey string) ApiCreateContentRequest {
	return ApiCreateContentRequest{
		ApiService: a,
		ctx: ctx,
		ownerKey: ownerKey,
	}
}

// Execute executes the request
//  @return ContentDTO
func (a *OwnerContentAPIService) CreateContentExecute(r ApiCreateContentRequest) (*ContentDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContentDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnerContentAPIService.CreateContent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners/{owner_key}/content"
	localVarPath = strings.Replace(localVarPath, "{"+"owner_key"+"}", url.PathEscape(parameterValueToString(r.ownerKey, "ownerKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentDTO == nil {
		return localVarReturnValue, nil, reportError("contentDTO is required and must be specified")
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
	localVarPostBody = r.contentDTO
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

type ApiCreateContentBatchRequest struct {
	ctx context.Context
	ApiService *OwnerContentAPIService
	ownerKey string
	contentDTO *[]ContentDTO
}

// List of contents to be created
func (r ApiCreateContentBatchRequest) ContentDTO(contentDTO []ContentDTO) ApiCreateContentBatchRequest {
	r.contentDTO = &contentDTO
	return r
}

func (r ApiCreateContentBatchRequest) Execute() ([]ContentDTO, *http.Response, error) {
	return r.ApiService.CreateContentBatchExecute(r)
}

/*
CreateContentBatch Method for CreateContentBatch

Creates Contents in bulk

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerKey key of the owner whose content has to be created
 @return ApiCreateContentBatchRequest
*/
func (a *OwnerContentAPIService) CreateContentBatch(ctx context.Context, ownerKey string) ApiCreateContentBatchRequest {
	return ApiCreateContentBatchRequest{
		ApiService: a,
		ctx: ctx,
		ownerKey: ownerKey,
	}
}

// Execute executes the request
//  @return []ContentDTO
func (a *OwnerContentAPIService) CreateContentBatchExecute(r ApiCreateContentBatchRequest) ([]ContentDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ContentDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnerContentAPIService.CreateContentBatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners/{owner_key}/content/batch"
	localVarPath = strings.Replace(localVarPath, "{"+"owner_key"+"}", url.PathEscape(parameterValueToString(r.ownerKey, "ownerKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentDTO == nil {
		return localVarReturnValue, nil, reportError("contentDTO is required and must be specified")
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
	localVarPostBody = r.contentDTO
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

type ApiGetContentByIdRequest struct {
	ctx context.Context
	ApiService *OwnerContentAPIService
	ownerKey string
	contentId string
}

func (r ApiGetContentByIdRequest) Execute() (*ContentDTO, *http.Response, error) {
	return r.ApiService.GetContentByIdExecute(r)
}

/*
GetContentById Method for GetContentById

Retrieves a single content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerKey The key of the owner whose content has to be retrieved
 @param contentId The id of the content to be retrieved
 @return ApiGetContentByIdRequest
*/
func (a *OwnerContentAPIService) GetContentById(ctx context.Context, ownerKey string, contentId string) ApiGetContentByIdRequest {
	return ApiGetContentByIdRequest{
		ApiService: a,
		ctx: ctx,
		ownerKey: ownerKey,
		contentId: contentId,
	}
}

// Execute executes the request
//  @return ContentDTO
func (a *OwnerContentAPIService) GetContentByIdExecute(r ApiGetContentByIdRequest) (*ContentDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContentDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnerContentAPIService.GetContentById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners/{owner_key}/content/{content_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"owner_key"+"}", url.PathEscape(parameterValueToString(r.ownerKey, "ownerKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"content_id"+"}", url.PathEscape(parameterValueToString(r.contentId, "contentId")), -1)

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

type ApiGetContentsByOwnerRequest struct {
	ctx context.Context
	ApiService *OwnerContentAPIService
	ownerKey string
	content *[]string
	label *[]string
	active *string
	custom *string
}

// The ID of a content to fetch. If defined, the list of contents returned by this method will only include those matching the given ID. May be specified multiple times to filter on multiple content IDs. 
func (r ApiGetContentsByOwnerRequest) Content(content []string) ApiGetContentsByOwnerRequest {
	r.content = &content
	return r
}

// The labels of content to fetch. If defined, the list of content returned by this endpoint will only include those matching the given labels (case-insensitive). May be specified multiple times to filter on multiple labels. If multiple labels are provided, any content matching any of the provided labels will be returned. 
func (r ApiGetContentsByOwnerRequest) Label(label []string) ApiGetContentsByOwnerRequest {
	r.label = &label
	return r
}

// A value indicating how &#39;active&#39; content should be considered when fetching content, where &#39;active&#39; is defined as a content that is currently in use by a subscription with a start time in the past and that has not yet expired, or in use by a product which itself is considered &#39;active.&#39; Must be one of &#39;include&#39;, &#39;exclude&#39;, or &#39;exclusive&#39; indicating that active content should be included along with inactive content, excluded (omitted) from the results, or exclusively selected as the only content to return. Defaults to &#39;exclusive&#39;. 
func (r ApiGetContentsByOwnerRequest) Active(active string) ApiGetContentsByOwnerRequest {
	r.active = &active
	return r
}

// A value indicating how custom content are considered when fetching content, where &#39;custom&#39; is defined as content that did not originate from a refresh operation nor manifest import. Must be one of &#39;include&#39;, &#39;exclude&#39;, or &#39;exclusive&#39; indicating that custom content should be passively included, excluded or omitted from the output, or exclusively selected as the only content to return. 
func (r ApiGetContentsByOwnerRequest) Custom(custom string) ApiGetContentsByOwnerRequest {
	r.custom = &custom
	return r
}

func (r ApiGetContentsByOwnerRequest) Execute() ([]ContentDTO, *http.Response, error) {
	return r.ApiService.GetContentsByOwnerExecute(r)
}

/*
GetContentsByOwner Method for GetContentsByOwner

Retrieves a list of contents within the given organization's namespace, optionally filtered
by a list of content IDs.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerKey Owner key
 @return ApiGetContentsByOwnerRequest
*/
func (a *OwnerContentAPIService) GetContentsByOwner(ctx context.Context, ownerKey string) ApiGetContentsByOwnerRequest {
	return ApiGetContentsByOwnerRequest{
		ApiService: a,
		ctx: ctx,
		ownerKey: ownerKey,
	}
}

// Execute executes the request
//  @return []ContentDTO
func (a *OwnerContentAPIService) GetContentsByOwnerExecute(r ApiGetContentsByOwnerRequest) ([]ContentDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ContentDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnerContentAPIService.GetContentsByOwner")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners/{owner_key}/content"
	localVarPath = strings.Replace(localVarPath, "{"+"owner_key"+"}", url.PathEscape(parameterValueToString(r.ownerKey, "ownerKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.content != nil {
		t := *r.content
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "content", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "content", t, "form", "multi")
		}
	}
	if r.label != nil {
		t := *r.label
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "label", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "label", t, "form", "multi")
		}
	}
	if r.active != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "active", r.active, "form", "")
	} else {
		var defaultValue string = "exclusive"
		r.active = &defaultValue
	}
	if r.custom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "custom", r.custom, "form", "")
	} else {
		var defaultValue string = "include"
		r.custom = &defaultValue
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

type ApiRemoveContentRequest struct {
	ctx context.Context
	ApiService *OwnerContentAPIService
	ownerKey string
	contentId string
}

func (r ApiRemoveContentRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveContentExecute(r)
}

/*
RemoveContent Method for RemoveContent

Deletes a Content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerKey The key of the owner whose content has to be deleted
 @param contentId The ID of the content to be deleted
 @return ApiRemoveContentRequest
*/
func (a *OwnerContentAPIService) RemoveContent(ctx context.Context, ownerKey string, contentId string) ApiRemoveContentRequest {
	return ApiRemoveContentRequest{
		ApiService: a,
		ctx: ctx,
		ownerKey: ownerKey,
		contentId: contentId,
	}
}

// Execute executes the request
func (a *OwnerContentAPIService) RemoveContentExecute(r ApiRemoveContentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnerContentAPIService.RemoveContent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners/{owner_key}/content/{content_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"owner_key"+"}", url.PathEscape(parameterValueToString(r.ownerKey, "ownerKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"content_id"+"}", url.PathEscape(parameterValueToString(r.contentId, "contentId")), -1)

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

type ApiUpdateContentRequest struct {
	ctx context.Context
	ApiService *OwnerContentAPIService
	ownerKey string
	contentId string
	contentDTO *ContentDTO
}

// Content to be updated
func (r ApiUpdateContentRequest) ContentDTO(contentDTO ContentDTO) ApiUpdateContentRequest {
	r.contentDTO = &contentDTO
	return r
}

func (r ApiUpdateContentRequest) Execute() (*ContentDTO, *http.Response, error) {
	return r.ApiService.UpdateContentExecute(r)
}

/*
UpdateContent Method for UpdateContent

Updates a Content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerKey The key of the owner to be updated
 @param contentId The ID of the Content to be updated
 @return ApiUpdateContentRequest
*/
func (a *OwnerContentAPIService) UpdateContent(ctx context.Context, ownerKey string, contentId string) ApiUpdateContentRequest {
	return ApiUpdateContentRequest{
		ApiService: a,
		ctx: ctx,
		ownerKey: ownerKey,
		contentId: contentId,
	}
}

// Execute executes the request
//  @return ContentDTO
func (a *OwnerContentAPIService) UpdateContentExecute(r ApiUpdateContentRequest) (*ContentDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContentDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OwnerContentAPIService.UpdateContent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/owners/{owner_key}/content/{content_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"owner_key"+"}", url.PathEscape(parameterValueToString(r.ownerKey, "ownerKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"content_id"+"}", url.PathEscape(parameterValueToString(r.contentId, "contentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentDTO == nil {
		return localVarReturnValue, nil, reportError("contentDTO is required and must be specified")
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
	localVarPostBody = r.contentDTO
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
