/*
public/portworx/pds/tasks/apiv1/tasks.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdsclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DeploymentServiceAPIService DeploymentServiceAPI service
type DeploymentServiceAPIService service

type ApiDeploymentServiceCreateDeploymentRequest struct {
	ctx context.Context
	ApiService *DeploymentServiceAPIService
	namespaceId string
	v1Deployment *V1Deployment
}

// Deployment resource.
func (r ApiDeploymentServiceCreateDeploymentRequest) V1Deployment(v1Deployment V1Deployment) ApiDeploymentServiceCreateDeploymentRequest {
	r.v1Deployment = &v1Deployment
	return r
}

func (r ApiDeploymentServiceCreateDeploymentRequest) Execute() (*V1Deployment, *http.Response, error) {
	return r.ApiService.DeploymentServiceCreateDeploymentExecute(r)
}

/*
DeploymentServiceCreateDeployment CreateDeployment API creates the Deployment resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param namespaceId UID of the namespace resource where this deployment will be created. (-- api-linter: core::0133::request-unknown-fields=disabled     aip.dev/not-precedent: We need this field for to support creation of     the resource in the namespace context. --)
 @return ApiDeploymentServiceCreateDeploymentRequest
*/
func (a *DeploymentServiceAPIService) DeploymentServiceCreateDeployment(ctx context.Context, namespaceId string) ApiDeploymentServiceCreateDeploymentRequest {
	return ApiDeploymentServiceCreateDeploymentRequest{
		ApiService: a,
		ctx: ctx,
		namespaceId: namespaceId,
	}
}

// Execute executes the request
//  @return V1Deployment
func (a *DeploymentServiceAPIService) DeploymentServiceCreateDeploymentExecute(r ApiDeploymentServiceCreateDeploymentRequest) (*V1Deployment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1Deployment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentServiceAPIService.DeploymentServiceCreateDeployment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pds/v1/namespace/{namespaceId}/deployment"
	localVarPath = strings.Replace(localVarPath, "{"+"namespaceId"+"}", url.PathEscape(parameterValueToString(r.namespaceId, "namespaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.v1Deployment == nil {
		return localVarReturnValue, nil, reportError("v1Deployment is required and must be specified")
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
	localVarPostBody = r.v1Deployment
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
			var v GooglerpcStatus1
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

type ApiDeploymentServiceDeleteDeploymentRequest struct {
	ctx context.Context
	ApiService *DeploymentServiceAPIService
	id string
}

func (r ApiDeploymentServiceDeleteDeploymentRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeploymentServiceDeleteDeploymentExecute(r)
}

/*
DeploymentServiceDeleteDeployment DeleteDeployment API deletes the Deployment resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id UID of the Deployment.
 @return ApiDeploymentServiceDeleteDeploymentRequest
*/
func (a *DeploymentServiceAPIService) DeploymentServiceDeleteDeployment(ctx context.Context, id string) ApiDeploymentServiceDeleteDeploymentRequest {
	return ApiDeploymentServiceDeleteDeploymentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DeploymentServiceAPIService) DeploymentServiceDeleteDeploymentExecute(r ApiDeploymentServiceDeleteDeploymentRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentServiceAPIService.DeploymentServiceDeleteDeployment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pds/v1/deployments/{id}"
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
			var v GooglerpcStatus1
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

type ApiDeploymentServiceGetDeploymentRequest struct {
	ctx context.Context
	ApiService *DeploymentServiceAPIService
	id string
}

func (r ApiDeploymentServiceGetDeploymentRequest) Execute() (*V1Deployment, *http.Response, error) {
	return r.ApiService.DeploymentServiceGetDeploymentExecute(r)
}

/*
DeploymentServiceGetDeployment GetDeployment API returns the Deployment resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id UID of the Deployment.
 @return ApiDeploymentServiceGetDeploymentRequest
*/
func (a *DeploymentServiceAPIService) DeploymentServiceGetDeployment(ctx context.Context, id string) ApiDeploymentServiceGetDeploymentRequest {
	return ApiDeploymentServiceGetDeploymentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return V1Deployment
func (a *DeploymentServiceAPIService) DeploymentServiceGetDeploymentExecute(r ApiDeploymentServiceGetDeploymentRequest) (*V1Deployment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1Deployment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentServiceAPIService.DeploymentServiceGetDeployment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pds/v1/deployments/{id}"
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
			var v GooglerpcStatus1
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

type ApiDeploymentServiceGetDeploymentCredentialsRequest struct {
	ctx context.Context
	ApiService *DeploymentServiceAPIService
	id string
}

func (r ApiDeploymentServiceGetDeploymentCredentialsRequest) Execute() (*V1DeploymentCredentials, *http.Response, error) {
	return r.ApiService.DeploymentServiceGetDeploymentCredentialsExecute(r)
}

/*
DeploymentServiceGetDeploymentCredentials GetDeploymentCredentials API returns the Credentials to be used to access the Deployment .

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id UID of the Deployment.
 @return ApiDeploymentServiceGetDeploymentCredentialsRequest
*/
func (a *DeploymentServiceAPIService) DeploymentServiceGetDeploymentCredentials(ctx context.Context, id string) ApiDeploymentServiceGetDeploymentCredentialsRequest {
	return ApiDeploymentServiceGetDeploymentCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return V1DeploymentCredentials
func (a *DeploymentServiceAPIService) DeploymentServiceGetDeploymentCredentialsExecute(r ApiDeploymentServiceGetDeploymentCredentialsRequest) (*V1DeploymentCredentials, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1DeploymentCredentials
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentServiceAPIService.DeploymentServiceGetDeploymentCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pds/v1/deployments/{id}:credentials"
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
			var v GooglerpcStatus1
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

type ApiDeploymentServiceListDeploymentsRequest struct {
	ctx context.Context
	ApiService *DeploymentServiceAPIService
	accountId *string
	tenantId *string
	clusterId *string
	namespaceId *string
	projectId *string
	paginationPageNumber *string
	paginationPageSize *string
	sortSortBy *string
	sortSortOrder *string
}

// UID of the account.
func (r ApiDeploymentServiceListDeploymentsRequest) AccountId(accountId string) ApiDeploymentServiceListDeploymentsRequest {
	r.accountId = &accountId
	return r
}

// UID of the tenant.
func (r ApiDeploymentServiceListDeploymentsRequest) TenantId(tenantId string) ApiDeploymentServiceListDeploymentsRequest {
	r.tenantId = &tenantId
	return r
}

// UID of the target cluster.
func (r ApiDeploymentServiceListDeploymentsRequest) ClusterId(clusterId string) ApiDeploymentServiceListDeploymentsRequest {
	r.clusterId = &clusterId
	return r
}

// UID of the namespace.
func (r ApiDeploymentServiceListDeploymentsRequest) NamespaceId(namespaceId string) ApiDeploymentServiceListDeploymentsRequest {
	r.namespaceId = &namespaceId
	return r
}

// UID of the project.
func (r ApiDeploymentServiceListDeploymentsRequest) ProjectId(projectId string) ApiDeploymentServiceListDeploymentsRequest {
	r.projectId = &projectId
	return r
}

// Page number is the page number to return based on the size
func (r ApiDeploymentServiceListDeploymentsRequest) PaginationPageNumber(paginationPageNumber string) ApiDeploymentServiceListDeploymentsRequest {
	r.paginationPageNumber = &paginationPageNumber
	return r
}

// Page size is the maximum number of records to include per page
func (r ApiDeploymentServiceListDeploymentsRequest) PaginationPageSize(paginationPageSize string) ApiDeploymentServiceListDeploymentsRequest {
	r.paginationPageSize = &paginationPageSize
	return r
}

// Name of the attribute to sort results by.   - FIELD_UNSPECIFIED: Unspecified, do not use.  - NAME: Sorting based on the name of the resource.  - CREATED_AT: Sorting on create time of the resource.  - UPDATED_AT: Sorting on update time of the resource.  - PHASE: Sorting on phase of the resource.
func (r ApiDeploymentServiceListDeploymentsRequest) SortSortBy(sortSortBy string) ApiDeploymentServiceListDeploymentsRequest {
	r.sortSortBy = &sortSortBy
	return r
}

// Order of sorting to be applied on requested list. If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.   - VALUE_UNSPECIFIED: Unspecified, do not use.  - ASC: Sort order ascending.  - DESC: Sort order descending.
func (r ApiDeploymentServiceListDeploymentsRequest) SortSortOrder(sortSortOrder string) ApiDeploymentServiceListDeploymentsRequest {
	r.sortSortOrder = &sortSortOrder
	return r
}

func (r ApiDeploymentServiceListDeploymentsRequest) Execute() (*V1ListDeploymentsResponse, *http.Response, error) {
	return r.ApiService.DeploymentServiceListDeploymentsExecute(r)
}

/*
DeploymentServiceListDeployments ListDeployments API lists the Deployment resources.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeploymentServiceListDeploymentsRequest
*/
func (a *DeploymentServiceAPIService) DeploymentServiceListDeployments(ctx context.Context) ApiDeploymentServiceListDeploymentsRequest {
	return ApiDeploymentServiceListDeploymentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1ListDeploymentsResponse
func (a *DeploymentServiceAPIService) DeploymentServiceListDeploymentsExecute(r ApiDeploymentServiceListDeploymentsRequest) (*V1ListDeploymentsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1ListDeploymentsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentServiceAPIService.DeploymentServiceListDeployments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pds/v1/deployments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountId", r.accountId, "")
	}
	if r.tenantId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tenantId", r.tenantId, "")
	}
	if r.clusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clusterId", r.clusterId, "")
	}
	if r.namespaceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespaceId", r.namespaceId, "")
	}
	if r.projectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "projectId", r.projectId, "")
	}
	if r.paginationPageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagination.pageNumber", r.paginationPageNumber, "")
	}
	if r.paginationPageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagination.pageSize", r.paginationPageSize, "")
	}
	if r.sortSortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort.sortBy", r.sortSortBy, "")
	} else {
		var defaultValue string = "FIELD_UNSPECIFIED"
		r.sortSortBy = &defaultValue
	}
	if r.sortSortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort.sortOrder", r.sortSortOrder, "")
	} else {
		var defaultValue string = "VALUE_UNSPECIFIED"
		r.sortSortOrder = &defaultValue
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
			var v GooglerpcStatus1
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

type ApiDeploymentServiceUpdateDeploymentRequest struct {
	ctx context.Context
	ApiService *DeploymentServiceAPIService
	v1Deployment *V1Deployment
}

// Deployment resource.
func (r ApiDeploymentServiceUpdateDeploymentRequest) V1Deployment(v1Deployment V1Deployment) ApiDeploymentServiceUpdateDeploymentRequest {
	r.v1Deployment = &v1Deployment
	return r
}

func (r ApiDeploymentServiceUpdateDeploymentRequest) Execute() (*V1Deployment, *http.Response, error) {
	return r.ApiService.DeploymentServiceUpdateDeploymentExecute(r)
}

/*
DeploymentServiceUpdateDeployment UpdateDeployment API updates the Deployment resource.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeploymentServiceUpdateDeploymentRequest
*/
func (a *DeploymentServiceAPIService) DeploymentServiceUpdateDeployment(ctx context.Context) ApiDeploymentServiceUpdateDeploymentRequest {
	return ApiDeploymentServiceUpdateDeploymentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1Deployment
func (a *DeploymentServiceAPIService) DeploymentServiceUpdateDeploymentExecute(r ApiDeploymentServiceUpdateDeploymentRequest) (*V1Deployment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1Deployment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentServiceAPIService.DeploymentServiceUpdateDeployment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pds/v1/deployments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.v1Deployment == nil {
		return localVarReturnValue, nil, reportError("v1Deployment is required and must be specified")
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
	localVarPostBody = r.v1Deployment
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
			var v GooglerpcStatus1
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
