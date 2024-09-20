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
	"reflect"
)


// ContentAPIService ContentAPI service
type ContentAPIService service

type ApiGetContentByUuidRequest struct {
	ctx context.Context
	ApiService *ContentAPIService
	contentUuid string
}

func (r ApiGetContentByUuidRequest) Execute() (*ContentDTO, *http.Response, error) {
	return r.ApiService.GetContentByUuidExecute(r)
}

/*
GetContentByUuid Method for GetContentByUuid

Retrieves a single Content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param contentUuid
 @return ApiGetContentByUuidRequest
*/
func (a *ContentAPIService) GetContentByUuid(ctx context.Context, contentUuid string) ApiGetContentByUuidRequest {
	return ApiGetContentByUuidRequest{
		ApiService: a,
		ctx: ctx,
		contentUuid: contentUuid,
	}
}

// Execute executes the request
//  @return ContentDTO
func (a *ContentAPIService) GetContentByUuidExecute(r ApiGetContentByUuidRequest) (*ContentDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContentDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentAPIService.GetContentByUuid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/content/{content_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"content_uuid"+"}", url.PathEscape(parameterValueToString(r.contentUuid, "contentUuid")), -1)

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

type ApiGetContentsRequest struct {
	ctx context.Context
	ApiService *ContentAPIService
	owner *[]string
	content *[]string
	label *[]string
	active *string
	custom *string
}

// The key of an owner to use to limit the content search. If defined, the list of contents returned by this endpoint will only include those available to the given owner. May be specified multiple times to filter by multiple owners. If multiple owners are provided, contents available to any of the provided owners will be returned. 
func (r ApiGetContentsRequest) Owner(owner []string) ApiGetContentsRequest {
	r.owner = &owner
	return r
}

// The ID of a content to fetch. If defined, the list of contents returned by this method will only include those matching the given ID. May be specified multiple times to filter on multiple content IDs. 
func (r ApiGetContentsRequest) Content(content []string) ApiGetContentsRequest {
	r.content = &content
	return r
}

// The labels of content to fetch. If defined, the list of content returned by this endpoint will only include those matching the given labels (case-insensitive). May be specified multiple times to filter on multiple labels. If multiple labels are provided, any content matching any of the provided labels will be returned. 
func (r ApiGetContentsRequest) Label(label []string) ApiGetContentsRequest {
	r.label = &label
	return r
}

// A value indicating how &#39;active&#39; content should be considered when fetching content, where &#39;active&#39; is defined as a content that is currently in use by a subscription with a start time in the past and that has not yet expired, or in use by a product which itself is considered &#39;active.&#39; Must be one of &#39;include&#39;, &#39;exclude&#39;, or &#39;exclusive&#39; indicating that active content should be included along with inactive content, excluded (omitted) from the results, or exclusively selected as the only content to return. Defaults to &#39;exclusive&#39;. 
func (r ApiGetContentsRequest) Active(active string) ApiGetContentsRequest {
	r.active = &active
	return r
}

// A value indicating how custom content are considered when fetching content, where &#39;custom&#39; is defined as content that did not originate from a refresh operation nor manifest import. Must be one of &#39;include&#39;, &#39;exclude&#39;, or &#39;exclusive&#39; indicating that custom content should be passively included, excluded or omitted from the output, or exclusively selected as the only content to return. 
func (r ApiGetContentsRequest) Custom(custom string) ApiGetContentsRequest {
	r.custom = &custom
	return r
}

func (r ApiGetContentsRequest) Execute() ([]ContentDTO, *http.Response, error) {
	return r.ApiService.GetContentsExecute(r)
}

/*
GetContents Method for GetContents

Retrieves list of Content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetContentsRequest
*/
func (a *ContentAPIService) GetContents(ctx context.Context) ApiGetContentsRequest {
	return ApiGetContentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ContentDTO
func (a *ContentAPIService) GetContentsExecute(r ApiGetContentsRequest) ([]ContentDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ContentDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentAPIService.GetContents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/content"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.owner != nil {
		t := *r.owner
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "owner", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "owner", t, "form", "multi")
		}
	}
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
