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


// AccountsRoleInvitationsApiService AccountsRoleInvitationsApi service
type AccountsRoleInvitationsApiService service

type ApiApiAccountRoleInvitationsIdDeleteRequest struct {
	ctx context.Context
	ApiService *AccountsRoleInvitationsApiService
	id string
}

func (r ApiApiAccountRoleInvitationsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiAccountRoleInvitationsIdDeleteExecute(r)
}

/*
ApiAccountRoleInvitationsIdDelete Delete AccountRoleInvitation

Removes a single AccountRoleInvitation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id AccountRoleInvitation ID (must be valid UUID)
 @return ApiApiAccountRoleInvitationsIdDeleteRequest
*/
func (a *AccountsRoleInvitationsApiService) ApiAccountRoleInvitationsIdDelete(ctx context.Context, id string) ApiApiAccountRoleInvitationsIdDeleteRequest {
	return ApiApiAccountRoleInvitationsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *AccountsRoleInvitationsApiService) ApiAccountRoleInvitationsIdDeleteExecute(r ApiApiAccountRoleInvitationsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsRoleInvitationsApiService.ApiAccountRoleInvitationsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/account-role-invitations/{id}"
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

type ApiApiAccountRoleInvitationsIdPatchRequest struct {
	ctx context.Context
	ApiService *AccountsRoleInvitationsApiService
	id string
	body *RequestsPatchAccountRoleInvitationRequest
}

// Object with patched role name
func (r ApiApiAccountRoleInvitationsIdPatchRequest) Body(body RequestsPatchAccountRoleInvitationRequest) ApiApiAccountRoleInvitationsIdPatchRequest {
	r.body = &body
	return r
}

func (r ApiApiAccountRoleInvitationsIdPatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiAccountRoleInvitationsIdPatchExecute(r)
}

/*
ApiAccountRoleInvitationsIdPatch Patch AccountRoleInvitation role_name

Patches AccountRoleInvitation role_name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id AccountRoleInvitation ID (must be valid UUID)
 @return ApiApiAccountRoleInvitationsIdPatchRequest
*/
func (a *AccountsRoleInvitationsApiService) ApiAccountRoleInvitationsIdPatch(ctx context.Context, id string) ApiApiAccountRoleInvitationsIdPatchRequest {
	return ApiApiAccountRoleInvitationsIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *AccountsRoleInvitationsApiService) ApiAccountRoleInvitationsIdPatchExecute(r ApiApiAccountRoleInvitationsIdPatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsRoleInvitationsApiService.ApiAccountRoleInvitationsIdPatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/account-role-invitations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
