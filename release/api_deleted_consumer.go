/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DeletedConsumerAPIService DeletedConsumerAPI service
type DeletedConsumerAPIService service

type ApiListByDateRequest struct {
	ctx context.Context
	ApiService *DeletedConsumerAPIService
	date *string
	page *int32
	perPage *int32
	order *string
	sortBy *string
}

// Filter deleted consumers to those on or after this date
func (r ApiListByDateRequest) Date(date string) ApiListByDateRequest {
	r.date = &date
	return r
}

// Page index to return
func (r ApiListByDateRequest) Page(page int32) ApiListByDateRequest {
	r.page = &page
	return r
}

// Number of items to return per page
func (r ApiListByDateRequest) PerPage(perPage int32) ApiListByDateRequest {
	r.perPage = &perPage
	return r
}

// Direction of ordering
func (r ApiListByDateRequest) Order(order string) ApiListByDateRequest {
	r.order = &order
	return r
}

// Property to use for ordering
func (r ApiListByDateRequest) SortBy(sortBy string) ApiListByDateRequest {
	r.sortBy = &sortBy
	return r
}

func (r ApiListByDateRequest) Execute() ([]DeletedConsumerDTO, *http.Response, error) {
	return r.ApiService.ListByDateExecute(r)
}

/*
ListByDate Method for ListByDate

Retrieves a list of deleted consumers by deletion date or all. List returned is the deleted Consumers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListByDateRequest
*/
func (a *DeletedConsumerAPIService) ListByDate(ctx context.Context) ApiListByDateRequest {
	return ApiListByDateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []DeletedConsumerDTO
func (a *DeletedConsumerAPIService) ListByDateExecute(r ApiListByDateRequest) ([]DeletedConsumerDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DeletedConsumerDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeletedConsumerAPIService.ListByDate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deleted_consumers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "form", "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "form", "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "form", "")
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
