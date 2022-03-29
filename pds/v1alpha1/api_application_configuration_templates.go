/*
PDS API

Portworx Data Services API Server

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ApplicationConfigurationTemplatesApiService ApplicationConfigurationTemplatesApi service
type ApplicationConfigurationTemplatesApiService service

type ApiApiApplicationConfigurationTemplatesIdDeleteRequest struct {
	ctx _context.Context
	ApiService *ApplicationConfigurationTemplatesApiService
	id string
}


func (r ApiApiApplicationConfigurationTemplatesIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ApiApplicationConfigurationTemplatesIdDeleteExecute(r)
}

/*
ApiApplicationConfigurationTemplatesIdDelete Delete ApplicationConfigurationTemplates

Removes a single ApplicationConfigurationTemplate

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ApplicationConfigurationTemplate ID (must be valid UUID)
 @return ApiApiApplicationConfigurationTemplatesIdDeleteRequest
*/
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdDelete(ctx _context.Context, id string) ApiApiApplicationConfigurationTemplatesIdDeleteRequest {
	return ApiApiApplicationConfigurationTemplatesIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdDeleteExecute(r ApiApiApplicationConfigurationTemplatesIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConfigurationTemplatesApiService.ApiApplicationConfigurationTemplatesIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/application-configuration-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiApplicationConfigurationTemplatesIdGetRequest struct {
	ctx _context.Context
	ApiService *ApplicationConfigurationTemplatesApiService
	id string
}


func (r ApiApiApplicationConfigurationTemplatesIdGetRequest) Execute() (ModelsApplicationConfigurationTemplate, *_nethttp.Response, error) {
	return r.ApiService.ApiApplicationConfigurationTemplatesIdGetExecute(r)
}

/*
ApiApplicationConfigurationTemplatesIdGet Get ApplicationConfigurationTemplate

Fetches a single ApplicationConfigurationTemplate

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ApplicationConfigurationTemplate ID (must be valid UUID)
 @return ApiApiApplicationConfigurationTemplatesIdGetRequest
*/
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdGet(ctx _context.Context, id string) ApiApiApplicationConfigurationTemplatesIdGetRequest {
	return ApiApiApplicationConfigurationTemplatesIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsApplicationConfigurationTemplate
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdGetExecute(r ApiApiApplicationConfigurationTemplatesIdGetRequest) (ModelsApplicationConfigurationTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ModelsApplicationConfigurationTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConfigurationTemplatesApiService.ApiApplicationConfigurationTemplatesIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/application-configuration-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiApplicationConfigurationTemplatesIdPutRequest struct {
	ctx _context.Context
	ApiService *ApplicationConfigurationTemplatesApiService
	id string
	body *ControllersUpdateApplicationConfigurationTemplateRequest
}

// Request body containing updated template
func (r ApiApiApplicationConfigurationTemplatesIdPutRequest) Body(body ControllersUpdateApplicationConfigurationTemplateRequest) ApiApiApplicationConfigurationTemplatesIdPutRequest {
	r.body = &body
	return r
}

func (r ApiApiApplicationConfigurationTemplatesIdPutRequest) Execute() (ModelsApplicationConfigurationTemplate, *_nethttp.Response, error) {
	return r.ApiService.ApiApplicationConfigurationTemplatesIdPutExecute(r)
}

/*
ApiApplicationConfigurationTemplatesIdPut Update ApplicationConfigurationTemplate

Updates existing ApplicationConfigurationTemplate

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ApplicationConfigurationTemplate ID (must be valid UUID)
 @return ApiApiApplicationConfigurationTemplatesIdPutRequest
*/
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdPut(ctx _context.Context, id string) ApiApiApplicationConfigurationTemplatesIdPutRequest {
	return ApiApiApplicationConfigurationTemplatesIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsApplicationConfigurationTemplate
func (a *ApplicationConfigurationTemplatesApiService) ApiApplicationConfigurationTemplatesIdPutExecute(r ApiApiApplicationConfigurationTemplatesIdPutRequest) (ModelsApplicationConfigurationTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ModelsApplicationConfigurationTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConfigurationTemplatesApiService.ApiApplicationConfigurationTemplatesIdPut")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/application-configuration-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest struct {
	ctx _context.Context
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

func (r ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest) Execute() (ControllersPaginatedApplicationConfigurationTemplates, *_nethttp.Response, error) {
	return r.ApiService.ApiTenantsIdApplicationConfigurationTemplatesGetExecute(r)
}

/*
ApiTenantsIdApplicationConfigurationTemplatesGet List ApplicationConfigurationTemplates

Lists ApplicationConfigurationTemplates

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Tenant ID (must be valid UUID)
 @return ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest
*/
func (a *ApplicationConfigurationTemplatesApiService) ApiTenantsIdApplicationConfigurationTemplatesGet(ctx _context.Context, id string) ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest {
	return ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ControllersPaginatedApplicationConfigurationTemplates
func (a *ApplicationConfigurationTemplatesApiService) ApiTenantsIdApplicationConfigurationTemplatesGetExecute(r ApiApiTenantsIdApplicationConfigurationTemplatesGetRequest) (ControllersPaginatedApplicationConfigurationTemplates, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ControllersPaginatedApplicationConfigurationTemplates
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConfigurationTemplatesApiService.ApiTenantsIdApplicationConfigurationTemplatesGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenants/{id}/application-configuration-templates"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest struct {
	ctx _context.Context
	ApiService *ApplicationConfigurationTemplatesApiService
	id string
	body *ControllersCreateApplicationConfigurationTemplatesRequest
}

// Request body containing the application configuration template
func (r ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest) Body(body ControllersCreateApplicationConfigurationTemplatesRequest) ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest {
	r.body = &body
	return r
}

func (r ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest) Execute() (ModelsApplicationConfigurationTemplate, *_nethttp.Response, error) {
	return r.ApiService.ApiTenantsIdApplicationConfigurationTemplatesPostExecute(r)
}

/*
ApiTenantsIdApplicationConfigurationTemplatesPost Create ApplicationConfigurationTemplates

Creates a new ApplicationConfigurationTemplates

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Tenant ID (must be valid UUID)
 @return ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest
*/
func (a *ApplicationConfigurationTemplatesApiService) ApiTenantsIdApplicationConfigurationTemplatesPost(ctx _context.Context, id string) ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest {
	return ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsApplicationConfigurationTemplate
func (a *ApplicationConfigurationTemplatesApiService) ApiTenantsIdApplicationConfigurationTemplatesPostExecute(r ApiApiTenantsIdApplicationConfigurationTemplatesPostRequest) (ModelsApplicationConfigurationTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ModelsApplicationConfigurationTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConfigurationTemplatesApiService.ApiTenantsIdApplicationConfigurationTemplatesPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenants/{id}/application-configuration-templates"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
