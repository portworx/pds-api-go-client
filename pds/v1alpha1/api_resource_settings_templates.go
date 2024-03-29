/*
PDS API

Portworx Data Services API Server

API version: 121
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// ResourceSettingsTemplatesApiService ResourceSettingsTemplatesApi service
type ResourceSettingsTemplatesApiService service

type ApiApiResourceSettingsTemplatesIdDeleteRequest struct {
	ctx context.Context
	ApiService *ResourceSettingsTemplatesApiService
	id string
}


func (r ApiApiResourceSettingsTemplatesIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiResourceSettingsTemplatesIdDeleteExecute(r)
}

/*
ApiResourceSettingsTemplatesIdDelete Delete ResourceSettingsTemplates

Removes a single ResourceSettingsTemplate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ResourceSettingsTemplate ID (must be valid UUID)
 @return ApiApiResourceSettingsTemplatesIdDeleteRequest
*/
func (a *ResourceSettingsTemplatesApiService) ApiResourceSettingsTemplatesIdDelete(ctx context.Context, id string) ApiApiResourceSettingsTemplatesIdDeleteRequest {
	return ApiApiResourceSettingsTemplatesIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ResourceSettingsTemplatesApiService) ApiResourceSettingsTemplatesIdDeleteExecute(r ApiApiResourceSettingsTemplatesIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceSettingsTemplatesApiService.ApiResourceSettingsTemplatesIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resource-settings-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiApiResourceSettingsTemplatesIdGetRequest struct {
	ctx context.Context
	ApiService *ResourceSettingsTemplatesApiService
	id string
}


func (r ApiApiResourceSettingsTemplatesIdGetRequest) Execute() (*ModelsResourceSettingsTemplate, *http.Response, error) {
	return r.ApiService.ApiResourceSettingsTemplatesIdGetExecute(r)
}

/*
ApiResourceSettingsTemplatesIdGet Get ResourceSettingsTemplate

Fetches a single ResourceSettingsTemplate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ResourceSettingsTemplate ID (must be valid UUID)
 @return ApiApiResourceSettingsTemplatesIdGetRequest
*/
func (a *ResourceSettingsTemplatesApiService) ApiResourceSettingsTemplatesIdGet(ctx context.Context, id string) ApiApiResourceSettingsTemplatesIdGetRequest {
	return ApiApiResourceSettingsTemplatesIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsResourceSettingsTemplate
func (a *ResourceSettingsTemplatesApiService) ApiResourceSettingsTemplatesIdGetExecute(r ApiApiResourceSettingsTemplatesIdGetRequest) (*ModelsResourceSettingsTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsResourceSettingsTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceSettingsTemplatesApiService.ApiResourceSettingsTemplatesIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resource-settings-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiApiResourceSettingsTemplatesIdPutRequest struct {
	ctx context.Context
	ApiService *ResourceSettingsTemplatesApiService
	id string
	body *ControllersUpdateResourceSettingsTemplateRequest
}

// Request body containing updated template
func (r ApiApiResourceSettingsTemplatesIdPutRequest) Body(body ControllersUpdateResourceSettingsTemplateRequest) ApiApiResourceSettingsTemplatesIdPutRequest {
	r.body = &body
	return r
}

func (r ApiApiResourceSettingsTemplatesIdPutRequest) Execute() (*ModelsResourceSettingsTemplate, *http.Response, error) {
	return r.ApiService.ApiResourceSettingsTemplatesIdPutExecute(r)
}

/*
ApiResourceSettingsTemplatesIdPut Update ResourceSettingsTemplate

Updates existing ResourceSettingsTemplate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ResourceSettingsTemplate ID (must be valid UUID)
 @return ApiApiResourceSettingsTemplatesIdPutRequest
*/
func (a *ResourceSettingsTemplatesApiService) ApiResourceSettingsTemplatesIdPut(ctx context.Context, id string) ApiApiResourceSettingsTemplatesIdPutRequest {
	return ApiApiResourceSettingsTemplatesIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsResourceSettingsTemplate
func (a *ResourceSettingsTemplatesApiService) ApiResourceSettingsTemplatesIdPutExecute(r ApiApiResourceSettingsTemplatesIdPutRequest) (*ModelsResourceSettingsTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsResourceSettingsTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceSettingsTemplatesApiService.ApiResourceSettingsTemplatesIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/resource-settings-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiApiTenantsIdResourceSettingsTemplatesGetRequest struct {
	ctx context.Context
	ApiService *ResourceSettingsTemplatesApiService
	id string
	sortBy *string
	limit *string
	continuation *string
	id2 *string
	name *string
	dataServiceId *string
}

// A given ResourceSettingsTemplates attribute to sort results by (one of: id, name, created_at)
func (r ApiApiTenantsIdResourceSettingsTemplatesGetRequest) SortBy(sortBy string) ApiApiTenantsIdResourceSettingsTemplatesGetRequest {
	r.sortBy = &sortBy
	return r
}
// Maximum number of rows to return (could be less)
func (r ApiApiTenantsIdResourceSettingsTemplatesGetRequest) Limit(limit string) ApiApiTenantsIdResourceSettingsTemplatesGetRequest {
	r.limit = &limit
	return r
}
// Use a token returned by a previous query to continue listing with the next batch of rows
func (r ApiApiTenantsIdResourceSettingsTemplatesGetRequest) Continuation(continuation string) ApiApiTenantsIdResourceSettingsTemplatesGetRequest {
	r.continuation = &continuation
	return r
}
// Filter results by ResourceSettingsTemplates id
func (r ApiApiTenantsIdResourceSettingsTemplatesGetRequest) Id2(id2 string) ApiApiTenantsIdResourceSettingsTemplatesGetRequest {
	r.id2 = &id2
	return r
}
// Filter results by ResourceSettingsTemplates name
func (r ApiApiTenantsIdResourceSettingsTemplatesGetRequest) Name(name string) ApiApiTenantsIdResourceSettingsTemplatesGetRequest {
	r.name = &name
	return r
}
// Filter results by DataService ID
func (r ApiApiTenantsIdResourceSettingsTemplatesGetRequest) DataServiceId(dataServiceId string) ApiApiTenantsIdResourceSettingsTemplatesGetRequest {
	r.dataServiceId = &dataServiceId
	return r
}

func (r ApiApiTenantsIdResourceSettingsTemplatesGetRequest) Execute() (*ModelsPaginatedResultModelsResourceSettingsTemplate, *http.Response, error) {
	return r.ApiService.ApiTenantsIdResourceSettingsTemplatesGetExecute(r)
}

/*
ApiTenantsIdResourceSettingsTemplatesGet List ResourceSettingsTemplates

Lists ResourceSettingsTemplates

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Tenant ID (must be valid UUID)
 @return ApiApiTenantsIdResourceSettingsTemplatesGetRequest
*/
func (a *ResourceSettingsTemplatesApiService) ApiTenantsIdResourceSettingsTemplatesGet(ctx context.Context, id string) ApiApiTenantsIdResourceSettingsTemplatesGetRequest {
	return ApiApiTenantsIdResourceSettingsTemplatesGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsPaginatedResultModelsResourceSettingsTemplate
func (a *ResourceSettingsTemplatesApiService) ApiTenantsIdResourceSettingsTemplatesGetExecute(r ApiApiTenantsIdResourceSettingsTemplatesGetRequest) (*ModelsPaginatedResultModelsResourceSettingsTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsPaginatedResultModelsResourceSettingsTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceSettingsTemplatesApiService.ApiTenantsIdResourceSettingsTemplatesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenants/{id}/resource-settings-templates"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.continuation != nil {
		localVarQueryParams.Add("continuation", parameterToString(*r.continuation, ""))
	}
	if r.id2 != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id2, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.dataServiceId != nil {
		localVarQueryParams.Add("data_service_id", parameterToString(*r.dataServiceId, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiApiTenantsIdResourceSettingsTemplatesPostRequest struct {
	ctx context.Context
	ApiService *ResourceSettingsTemplatesApiService
	id string
	body *ControllersCreateResourceSettingsTemplateRequest
}

// Request body containing the resource settings template
func (r ApiApiTenantsIdResourceSettingsTemplatesPostRequest) Body(body ControllersCreateResourceSettingsTemplateRequest) ApiApiTenantsIdResourceSettingsTemplatesPostRequest {
	r.body = &body
	return r
}

func (r ApiApiTenantsIdResourceSettingsTemplatesPostRequest) Execute() (*ModelsResourceSettingsTemplate, *http.Response, error) {
	return r.ApiService.ApiTenantsIdResourceSettingsTemplatesPostExecute(r)
}

/*
ApiTenantsIdResourceSettingsTemplatesPost Create ResourceSettingsTemplate

Creates a new ResourceSettingsTemplate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Tenant ID (must be valid UUID)
 @return ApiApiTenantsIdResourceSettingsTemplatesPostRequest
*/
func (a *ResourceSettingsTemplatesApiService) ApiTenantsIdResourceSettingsTemplatesPost(ctx context.Context, id string) ApiApiTenantsIdResourceSettingsTemplatesPostRequest {
	return ApiApiTenantsIdResourceSettingsTemplatesPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsResourceSettingsTemplate
func (a *ResourceSettingsTemplatesApiService) ApiTenantsIdResourceSettingsTemplatesPostExecute(r ApiApiTenantsIdResourceSettingsTemplatesPostRequest) (*ModelsResourceSettingsTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsResourceSettingsTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceSettingsTemplatesApiService.ApiTenantsIdResourceSettingsTemplatesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenants/{id}/resource-settings-templates"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
