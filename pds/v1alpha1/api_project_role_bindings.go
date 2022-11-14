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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// ProjectRoleBindingsApiService ProjectRoleBindingsApi service
type ProjectRoleBindingsApiService service

type ApiApiProjectsIdInvitationsPostRequest struct {
	ctx context.Context
	ApiService *ProjectRoleBindingsApiService
	id string
	body *ControllersInvitationRequest
}

// Request body containing the invitation details.
func (r ApiApiProjectsIdInvitationsPostRequest) Body(body ControllersInvitationRequest) ApiApiProjectsIdInvitationsPostRequest {
	r.body = &body
	return r
}

func (r ApiApiProjectsIdInvitationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiProjectsIdInvitationsPostExecute(r)
}

/*
ApiProjectsIdInvitationsPost Create Project Invitation

Adds project role binding to existing user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Project ID (must be valid UUID)
 @return ApiApiProjectsIdInvitationsPostRequest
*/
func (a *ProjectRoleBindingsApiService) ApiProjectsIdInvitationsPost(ctx context.Context, id string) ApiApiProjectsIdInvitationsPostRequest {
	return ApiApiProjectsIdInvitationsPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ProjectRoleBindingsApiService) ApiProjectsIdInvitationsPostExecute(r ApiApiProjectsIdInvitationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectRoleBindingsApiService.ApiProjectsIdInvitationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/projects/{id}/invitations"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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

type ApiApiProjectsIdRoleBindingsDeleteRequest struct {
	ctx context.Context
	ApiService *ProjectRoleBindingsApiService
	id string
	body *RequestsDeleteRoleBindingRequest
}

// Request body containing the project role binding
func (r ApiApiProjectsIdRoleBindingsDeleteRequest) Body(body RequestsDeleteRoleBindingRequest) ApiApiProjectsIdRoleBindingsDeleteRequest {
	r.body = &body
	return r
}

func (r ApiApiProjectsIdRoleBindingsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiProjectsIdRoleBindingsDeleteExecute(r)
}

/*
ApiProjectsIdRoleBindingsDelete Delete ProjectRoleBinding

Removes a single ProjectRoleBinding

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Project ID (must be valid UUID)
 @return ApiApiProjectsIdRoleBindingsDeleteRequest
*/
func (a *ProjectRoleBindingsApiService) ApiProjectsIdRoleBindingsDelete(ctx context.Context, id string) ApiApiProjectsIdRoleBindingsDeleteRequest {
	return ApiApiProjectsIdRoleBindingsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ProjectRoleBindingsApiService) ApiProjectsIdRoleBindingsDeleteExecute(r ApiApiProjectsIdRoleBindingsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectRoleBindingsApiService.ApiProjectsIdRoleBindingsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/projects/{id}/role-bindings"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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

type ApiApiProjectsIdRoleBindingsGetRequest struct {
	ctx context.Context
	ApiService *ProjectRoleBindingsApiService
	id string
	sortBy *string
	roleName *string
	actorId *string
	actorType *string
}

// A given ProjectRoleBinding attribute to sort results by (one of: role_name, actor_id)
func (r ApiApiProjectsIdRoleBindingsGetRequest) SortBy(sortBy string) ApiApiProjectsIdRoleBindingsGetRequest {
	r.sortBy = &sortBy
	return r
}
// Filter results by ProjectRoleBinding assigned role name
func (r ApiApiProjectsIdRoleBindingsGetRequest) RoleName(roleName string) ApiApiProjectsIdRoleBindingsGetRequest {
	r.roleName = &roleName
	return r
}
// Filter results by ProjectRoleBinding actor id
func (r ApiApiProjectsIdRoleBindingsGetRequest) ActorId(actorId string) ApiApiProjectsIdRoleBindingsGetRequest {
	r.actorId = &actorId
	return r
}
// Filter results by ProjectRoleBinding actor type
func (r ApiApiProjectsIdRoleBindingsGetRequest) ActorType(actorType string) ApiApiProjectsIdRoleBindingsGetRequest {
	r.actorType = &actorType
	return r
}

func (r ApiApiProjectsIdRoleBindingsGetRequest) Execute() (*ControllersPaginatedProjectRoleBindings, *http.Response, error) {
	return r.ApiService.ApiProjectsIdRoleBindingsGetExecute(r)
}

/*
ApiProjectsIdRoleBindingsGet List ProjectRoleBindings

Lists ProjectRoleBindings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Project ID (must be valid UUID)
 @return ApiApiProjectsIdRoleBindingsGetRequest
*/
func (a *ProjectRoleBindingsApiService) ApiProjectsIdRoleBindingsGet(ctx context.Context, id string) ApiApiProjectsIdRoleBindingsGetRequest {
	return ApiApiProjectsIdRoleBindingsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ControllersPaginatedProjectRoleBindings
func (a *ProjectRoleBindingsApiService) ApiProjectsIdRoleBindingsGetExecute(r ApiApiProjectsIdRoleBindingsGetRequest) (*ControllersPaginatedProjectRoleBindings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ControllersPaginatedProjectRoleBindings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectRoleBindingsApiService.ApiProjectsIdRoleBindingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/projects/{id}/role-bindings"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.roleName != nil {
		localVarQueryParams.Add("role_name", parameterToString(*r.roleName, ""))
	}
	if r.actorId != nil {
		localVarQueryParams.Add("actor_id", parameterToString(*r.actorId, ""))
	}
	if r.actorType != nil {
		localVarQueryParams.Add("actor_type", parameterToString(*r.actorType, ""))
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

type ApiApiProjectsIdRoleBindingsPutRequest struct {
	ctx context.Context
	ApiService *ProjectRoleBindingsApiService
	id string
	body *ControllersUpsertProjectRoleBindingRequest
}

// Request body containing the project role binding
func (r ApiApiProjectsIdRoleBindingsPutRequest) Body(body ControllersUpsertProjectRoleBindingRequest) ApiApiProjectsIdRoleBindingsPutRequest {
	r.body = &body
	return r
}

func (r ApiApiProjectsIdRoleBindingsPutRequest) Execute() (*ModelsProjectRoleBinding, *http.Response, error) {
	return r.ApiService.ApiProjectsIdRoleBindingsPutExecute(r)
}

/*
ApiProjectsIdRoleBindingsPut Create ProjectRoleBinding

Creates a new ProjectRoleBinding

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Project ID (must be valid UUID)
 @return ApiApiProjectsIdRoleBindingsPutRequest
*/
func (a *ProjectRoleBindingsApiService) ApiProjectsIdRoleBindingsPut(ctx context.Context, id string) ApiApiProjectsIdRoleBindingsPutRequest {
	return ApiApiProjectsIdRoleBindingsPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsProjectRoleBinding
func (a *ProjectRoleBindingsApiService) ApiProjectsIdRoleBindingsPutExecute(r ApiApiProjectsIdRoleBindingsPutRequest) (*ModelsProjectRoleBinding, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsProjectRoleBinding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectRoleBindingsApiService.ApiProjectsIdRoleBindingsPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/projects/{id}/role-bindings"
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
