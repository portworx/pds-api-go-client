/*
PDS API

Portworx Data Services API Server

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ApplicationConfigurationTemplatesApiService ApplicationConfigurationTemplatesApi service
type ApplicationConfigurationTemplatesApiService service

type ApiApiApplicationConfigurationTemplatesIdDeleteRequest struct {
	ctx context.Context
	ApiService *ApplicationConfigurationTemplatesApiService
	id string
}

func (r ApiApiApplicationConfigurationTemplatesIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiApplicationConfigurationTemplatesIdDeleteExecute(r)
}

/*
ApiApplicationConfigurationTemplatesIdDelete Delete ApplicationConfigurationTemplates

Removes a single ApplicationConfigurationTemplate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ApplicationConfigurationTemplate ID (must be valid UUID)
 @return ApiApiApplicationConfigurationTemplatesIdDeleteRequest
*/
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdDelete(ctx context.Context, id string) ApiApiApplicationConfigurationTemplatesIdDeleteRequest {
	return ApiApiApplicationConfigurationTemplatesIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdDeleteExecute(r ApiApiApplicationConfigurationTemplatesIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConfigurationTemplatesApiService.ApiApplicationConfigurationTemplatesIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/application-configuration-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiApplicationConfigurationTemplatesIdGetRequest struct {
	ctx context.Context
	ApiService *ApplicationConfigurationTemplatesApiService
	id string
}

func (r ApiApiApplicationConfigurationTemplatesIdGetRequest) Execute() (*ModelsApplicationConfigurationTemplate, *http.Response, error) {
	return r.ApiService.ApiApplicationConfigurationTemplatesIdGetExecute(r)
}

/*
ApiApplicationConfigurationTemplatesIdGet Get ApplicationConfigurationTemplate

Fetches a single ApplicationConfigurationTemplate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ApplicationConfigurationTemplate ID (must be valid UUID)
 @return ApiApiApplicationConfigurationTemplatesIdGetRequest
*/
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdGet(ctx context.Context, id string) ApiApiApplicationConfigurationTemplatesIdGetRequest {
	return ApiApiApplicationConfigurationTemplatesIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsApplicationConfigurationTemplate
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdGetExecute(r ApiApiApplicationConfigurationTemplatesIdGetRequest) (*ModelsApplicationConfigurationTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsApplicationConfigurationTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConfigurationTemplatesApiService.ApiApplicationConfigurationTemplatesIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/application-configuration-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiApiApplicationConfigurationTemplatesIdPutRequest struct {
	ctx context.Context
	ApiService *ApplicationConfigurationTemplatesApiService
	id string
	body *ControllersUpdateApplicationConfigurationTemplateRequest
}

// Request body containing updated template
func (r ApiApiApplicationConfigurationTemplatesIdPutRequest) Body(body ControllersUpdateApplicationConfigurationTemplateRequest) ApiApiApplicationConfigurationTemplatesIdPutRequest {
	r.body = &body
	return r
}

func (r ApiApiApplicationConfigurationTemplatesIdPutRequest) Execute() (*ModelsApplicationConfigurationTemplate, *http.Response, error) {
	return r.ApiService.ApiApplicationConfigurationTemplatesIdPutExecute(r)
}

/*
ApiApplicationConfigurationTemplatesIdPut Update ApplicationConfigurationTemplate

Updates existing ApplicationConfigurationTemplate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ApplicationConfigurationTemplate ID (must be valid UUID)
 @return ApiApiApplicationConfigurationTemplatesIdPutRequest
*/
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdPut(ctx context.Context, id string) ApiApiApplicationConfigurationTemplatesIdPutRequest {
	return ApiApiApplicationConfigurationTemplatesIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsApplicationConfigurationTemplate
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdPutExecute(r ApiApiApplicationConfigurationTemplatesIdPutRequest) (*ModelsApplicationConfigurationTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsApplicationConfigurationTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConfigurationTemplatesApiService.ApiApplicationConfigurationTemplatesIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/application-configuration-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest struct {
	ctx context.Context
	ApiService *ApplicationConfigurationTemplatesApiService
	id string
	sortBy *string
	limit *string
	continuation *string
	id2 *string
	name *string
	dataServiceId *string
}

// A given ApplicationConfigurationTemplates attribute to sort results by (one of: id, name, created_at)
func (r ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest) SortBy(sortBy string) ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest {
	r.sortBy = &sortBy
	return r
}

// Maximum number of rows to return (could be less)
func (r ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest) Limit(limit string) ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest {
	r.limit = &limit
	return r
}

// Use a token returned by a previous query to continue listing with the next batch of rows
func (r ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest) Continuation(continuation string) ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest {
	r.continuation = &continuation
	return r
}

// Filter results by ApplicationConfigurationTemplates id
func (r ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest) Id2(id2 string) ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest {
	r.id2 = &id2
	return r
}

// Filter results by ApplicationConfigurationTemplates name
func (r ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest) Name(name string) ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest {
	r.name = &name
	return r
}

// Filter results by DataService ID
func (r ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest) DataServiceId(dataServiceId string) ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest {
	r.dataServiceId = &dataServiceId
	return r
}

func (r ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest) Execute() (*ModelsPaginatedResultModelsApplicationConfigurationTemplate, *http.Response, error) {
	return r.ApiService.ApiTenantsIdApplicationConfigurationTemplatesGetExecute(r)
}

/*
ApiTenantsIdApplicationConfigurationTemplatesGet List ApplicationConfigurationTemplates

Lists ApplicationConfigurationTemplates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Tenant ID (must be valid UUID)
 @return ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest
*/
func (a *ApplicationConfigurationTemplatesApiService) ApiTenantsIdApplicationConfigurationTemplatesGet(ctx context.Context, id string) ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest {
	return ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsPaginatedResultModelsApplicationConfigurationTemplate
func (a *ApplicationConfigurationTemplatesApiService) ApiTenantsIdApplicationConfigurationTemplatesGetExecute(r ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest) (*ModelsPaginatedResultModelsApplicationConfigurationTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsPaginatedResultModelsApplicationConfigurationTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConfigurationTemplatesApiService.ApiTenantsIdApplicationConfigurationTemplatesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenants/{id}/application-configuration-templates"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.continuation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "continuation", r.continuation, "")
	}
	if r.id2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id2, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.dataServiceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "data_service_id", r.dataServiceId, "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest struct {
	ctx context.Context
	ApiService *ApplicationConfigurationTemplatesApiService
	id string
	body *ControllersCreateApplicationConfigurationTemplateRequest
}

// Request body containing the application configuration template
func (r ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest) Body(body ControllersCreateApplicationConfigurationTemplateRequest) ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest {
	r.body = &body
	return r
}

func (r ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest) Execute() (*ModelsApplicationConfigurationTemplate, *http.Response, error) {
	return r.ApiService.ApiTenantsIdApplicationConfigurationTemplatesPostExecute(r)
}

/*
ApiTenantsIdApplicationConfigurationTemplatesPost Create ApplicationConfigurationTemplate

Creates a new ApplicationConfigurationTemplate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Tenant ID (must be valid UUID)
 @return ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest
*/
func (a *ApplicationConfigurationTemplatesApiService) ApiTenantsIdApplicationConfigurationTemplatesPost(ctx context.Context, id string) ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest {
	return ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsApplicationConfigurationTemplate
func (a *ApplicationConfigurationTemplatesApiService) ApiTenantsIdApplicationConfigurationTemplatesPostExecute(r ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest) (*ModelsApplicationConfigurationTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsApplicationConfigurationTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConfigurationTemplatesApiService.ApiTenantsIdApplicationConfigurationTemplatesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenants/{id}/application-configuration-templates"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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
