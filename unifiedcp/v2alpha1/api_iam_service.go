/*
public/portworx/pds/backupconfig/apiv1/backupconfig.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// IAMServiceApiService IAMServiceApi service
type IAMServiceApiService service

type ApiIAMServiceCreateIAMRequest struct {
	ctx context.Context
	ApiService *IAMServiceApiService
	v1IAM *V1IAM
}

// IAM to be created
func (r ApiIAMServiceCreateIAMRequest) V1IAM(v1IAM V1IAM) ApiIAMServiceCreateIAMRequest {
	r.v1IAM = &v1IAM
	return r
}

func (r ApiIAMServiceCreateIAMRequest) Execute() (*V1IAM, *http.Response, error) {
	return r.ApiService.IAMServiceCreateIAMExecute(r)
}

/*
IAMServiceCreateIAM CreateIAM API creates a new IAM rolebinding

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIAMServiceCreateIAMRequest
*/
func (a *IAMServiceApiService) IAMServiceCreateIAM(ctx context.Context) ApiIAMServiceCreateIAMRequest {
	return ApiIAMServiceCreateIAMRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1IAM
func (a *IAMServiceApiService) IAMServiceCreateIAMExecute(r ApiIAMServiceCreateIAMRequest) (*V1IAM, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1IAM
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IAMServiceApiService.IAMServiceCreateIAM")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/iam"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.v1IAM == nil {
		return localVarReturnValue, nil, reportError("v1IAM is required and must be specified")
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
	localVarPostBody = r.v1IAM
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
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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

type ApiIAMServiceDeleteIAMRequest struct {
	ctx context.Context
	ApiService *IAMServiceApiService
	actorId string
}

func (r ApiIAMServiceDeleteIAMRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.IAMServiceDeleteIAMExecute(r)
}

/*
IAMServiceDeleteIAM DeleteIAM API delete IAM, currently required only for name

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID for which the IAM should be deleted.
 @return ApiIAMServiceDeleteIAMRequest
*/
func (a *IAMServiceApiService) IAMServiceDeleteIAM(ctx context.Context, actorId string) ApiIAMServiceDeleteIAMRequest {
	return ApiIAMServiceDeleteIAMRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *IAMServiceApiService) IAMServiceDeleteIAMExecute(r ApiIAMServiceDeleteIAMRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IAMServiceApiService.IAMServiceDeleteIAM")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/iam/{actorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterToString(r.actorId, "")), -1)

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
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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

type ApiIAMServiceGetIAMRequest struct {
	ctx context.Context
	ApiService *IAMServiceApiService
	actorId string
}

func (r ApiIAMServiceGetIAMRequest) Execute() (*V1IAM, *http.Response, error) {
	return r.ApiService.IAMServiceGetIAMExecute(r)
}

/*
IAMServiceGetIAM GetIAM API returns the info about IAM for given IAM id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param actorId Actor ID for which the details need to be fetched
 @return ApiIAMServiceGetIAMRequest
*/
func (a *IAMServiceApiService) IAMServiceGetIAM(ctx context.Context, actorId string) ApiIAMServiceGetIAMRequest {
	return ApiIAMServiceGetIAMRequest{
		ApiService: a,
		ctx: ctx,
		actorId: actorId,
	}
}

// Execute executes the request
//  @return V1IAM
func (a *IAMServiceApiService) IAMServiceGetIAMExecute(r ApiIAMServiceGetIAMRequest) (*V1IAM, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1IAM
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IAMServiceApiService.IAMServiceGetIAM")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/iam/{actorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"actorId"+"}", url.PathEscape(parameterToString(r.actorId, "")), -1)

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
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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

type ApiIAMServiceGrantIAMRequest struct {
	ctx context.Context
	ApiService *IAMServiceApiService
	iamConfigActorId string
	iAMServiceGrantIAMBody *IAMServiceGrantIAMBody
}

func (r ApiIAMServiceGrantIAMRequest) IAMServiceGrantIAMBody(iAMServiceGrantIAMBody IAMServiceGrantIAMBody) ApiIAMServiceGrantIAMRequest {
	r.iAMServiceGrantIAMBody = &iAMServiceGrantIAMBody
	return r
}

func (r ApiIAMServiceGrantIAMRequest) Execute() (*V1GrantIAMResponse, *http.Response, error) {
	return r.ApiService.IAMServiceGrantIAMExecute(r)
}

/*
IAMServiceGrantIAM GrantIAM API creates new IAM rolebinding at tenant, project and account level

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param iamConfigActorId Actor ID for the associated actor
 @return ApiIAMServiceGrantIAMRequest
*/
func (a *IAMServiceApiService) IAMServiceGrantIAM(ctx context.Context, iamConfigActorId string) ApiIAMServiceGrantIAMRequest {
	return ApiIAMServiceGrantIAMRequest{
		ApiService: a,
		ctx: ctx,
		iamConfigActorId: iamConfigActorId,
	}
}

// Execute executes the request
//  @return V1GrantIAMResponse
func (a *IAMServiceApiService) IAMServiceGrantIAMExecute(r ApiIAMServiceGrantIAMRequest) (*V1GrantIAMResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1GrantIAMResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IAMServiceApiService.IAMServiceGrantIAM")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/iam/{iam.config.actorId}:grant"
	localVarPath = strings.Replace(localVarPath, "{"+"iam.config.actorId"+"}", url.PathEscape(parameterToString(r.iamConfigActorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.iAMServiceGrantIAMBody == nil {
		return localVarReturnValue, nil, reportError("iAMServiceGrantIAMBody is required and must be specified")
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
	localVarPostBody = r.iAMServiceGrantIAMBody
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
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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

type ApiIAMServiceListIAMRequest struct {
	ctx context.Context
	ApiService *IAMServiceApiService
	actorId *string
	accountId *string
	tenantId *string
	projectId *string
	paginationPageNumber *string
	paginationPageSize *string
}

// Actor ID for which the IAM should be listed
func (r ApiIAMServiceListIAMRequest) ActorId(actorId string) ApiIAMServiceListIAMRequest {
	r.actorId = &actorId
	return r
}

// Account UID for which the IAM should be listed
func (r ApiIAMServiceListIAMRequest) AccountId(accountId string) ApiIAMServiceListIAMRequest {
	r.accountId = &accountId
	return r
}

// Tenant UID for which the IAM should be listed
func (r ApiIAMServiceListIAMRequest) TenantId(tenantId string) ApiIAMServiceListIAMRequest {
	r.tenantId = &tenantId
	return r
}

// Project UID for which the IAM should be listed
func (r ApiIAMServiceListIAMRequest) ProjectId(projectId string) ApiIAMServiceListIAMRequest {
	r.projectId = &projectId
	return r
}

// Page number is the page number to return based on the size
func (r ApiIAMServiceListIAMRequest) PaginationPageNumber(paginationPageNumber string) ApiIAMServiceListIAMRequest {
	r.paginationPageNumber = &paginationPageNumber
	return r
}

// Page size is the maximum number of records to include per page
func (r ApiIAMServiceListIAMRequest) PaginationPageSize(paginationPageSize string) ApiIAMServiceListIAMRequest {
	r.paginationPageSize = &paginationPageSize
	return r
}

func (r ApiIAMServiceListIAMRequest) Execute() (*V1ListIAMResponse, *http.Response, error) {
	return r.ApiService.IAMServiceListIAMExecute(r)
}

/*
IAMServiceListIAM ListIAM API lists the role bindings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIAMServiceListIAMRequest
*/
func (a *IAMServiceApiService) IAMServiceListIAM(ctx context.Context) ApiIAMServiceListIAMRequest {
	return ApiIAMServiceListIAMRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1ListIAMResponse
func (a *IAMServiceApiService) IAMServiceListIAMExecute(r ApiIAMServiceListIAMRequest) (*V1ListIAMResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1ListIAMResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IAMServiceApiService.IAMServiceListIAM")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/iam"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.actorId != nil {
		localVarQueryParams.Add("actorId", parameterToString(*r.actorId, ""))
	}
	if r.accountId != nil {
		localVarQueryParams.Add("accountId", parameterToString(*r.accountId, ""))
	}
	if r.tenantId != nil {
		localVarQueryParams.Add("tenantId", parameterToString(*r.tenantId, ""))
	}
	if r.projectId != nil {
		localVarQueryParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.paginationPageNumber != nil {
		localVarQueryParams.Add("pagination.pageNumber", parameterToString(*r.paginationPageNumber, ""))
	}
	if r.paginationPageSize != nil {
		localVarQueryParams.Add("pagination.pageSize", parameterToString(*r.paginationPageSize, ""))
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
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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

type ApiIAMServiceRevokeIAMRequest struct {
	ctx context.Context
	ApiService *IAMServiceApiService
	iamConfigActorId string
	iAMServiceRevokeIAMBody *IAMServiceRevokeIAMBody
}

func (r ApiIAMServiceRevokeIAMRequest) IAMServiceRevokeIAMBody(iAMServiceRevokeIAMBody IAMServiceRevokeIAMBody) ApiIAMServiceRevokeIAMRequest {
	r.iAMServiceRevokeIAMBody = &iAMServiceRevokeIAMBody
	return r
}

func (r ApiIAMServiceRevokeIAMRequest) Execute() (*V1RevokeIAMResponse, *http.Response, error) {
	return r.ApiService.IAMServiceRevokeIAMExecute(r)
}

/*
IAMServiceRevokeIAM RevokeIAM API delete IAM rolebinding at tenant, project and account level

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param iamConfigActorId Actor ID for the associated actor
 @return ApiIAMServiceRevokeIAMRequest
*/
func (a *IAMServiceApiService) IAMServiceRevokeIAM(ctx context.Context, iamConfigActorId string) ApiIAMServiceRevokeIAMRequest {
	return ApiIAMServiceRevokeIAMRequest{
		ApiService: a,
		ctx: ctx,
		iamConfigActorId: iamConfigActorId,
	}
}

// Execute executes the request
//  @return V1RevokeIAMResponse
func (a *IAMServiceApiService) IAMServiceRevokeIAMExecute(r ApiIAMServiceRevokeIAMRequest) (*V1RevokeIAMResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1RevokeIAMResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IAMServiceApiService.IAMServiceRevokeIAM")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/iam/{iam.config.actorId}:revoke"
	localVarPath = strings.Replace(localVarPath, "{"+"iam.config.actorId"+"}", url.PathEscape(parameterToString(r.iamConfigActorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.iAMServiceRevokeIAMBody == nil {
		return localVarReturnValue, nil, reportError("iAMServiceRevokeIAMBody is required and must be specified")
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
	localVarPostBody = r.iAMServiceRevokeIAMBody
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
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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

type ApiIAMServiceUpdateIAMRequest struct {
	ctx context.Context
	ApiService *IAMServiceApiService
	iamMetaUid string
	iAMToBeUpdated *IAMToBeUpdated
}

// IAM to be updated
func (r ApiIAMServiceUpdateIAMRequest) IAMToBeUpdated(iAMToBeUpdated IAMToBeUpdated) ApiIAMServiceUpdateIAMRequest {
	r.iAMToBeUpdated = &iAMToBeUpdated
	return r
}

func (r ApiIAMServiceUpdateIAMRequest) Execute() (*V1IAM, *http.Response, error) {
	return r.ApiService.IAMServiceUpdateIAMExecute(r)
}

/*
IAMServiceUpdateIAM UpdateIAM API updates IAM with the new set of role bindings. The request replaces the existing set of bindings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param iamMetaUid UID of the resource of the format <resource prefix>-<uuid>.
 @return ApiIAMServiceUpdateIAMRequest
*/
func (a *IAMServiceApiService) IAMServiceUpdateIAM(ctx context.Context, iamMetaUid string) ApiIAMServiceUpdateIAMRequest {
	return ApiIAMServiceUpdateIAMRequest{
		ApiService: a,
		ctx: ctx,
		iamMetaUid: iamMetaUid,
	}
}

// Execute executes the request
//  @return V1IAM
func (a *IAMServiceApiService) IAMServiceUpdateIAMExecute(r ApiIAMServiceUpdateIAMRequest) (*V1IAM, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1IAM
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IAMServiceApiService.IAMServiceUpdateIAM")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/iam/{iam.meta.uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"iam.meta.uid"+"}", url.PathEscape(parameterToString(r.iamMetaUid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.iAMToBeUpdated == nil {
		return localVarReturnValue, nil, reportError("iAMToBeUpdated is required and must be specified")
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
	localVarPostBody = r.iAMToBeUpdated
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
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
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
