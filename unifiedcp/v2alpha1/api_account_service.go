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


// AccountServiceApiService AccountServiceApi service
type AccountServiceApiService service

type ApiAccountServiceCreateAccountRequest struct {
	ctx context.Context
	ApiService *AccountServiceApiService
	v1Account *V1Account
}

// account to be created
func (r ApiAccountServiceCreateAccountRequest) V1Account(v1Account V1Account) ApiAccountServiceCreateAccountRequest {
	r.v1Account = &v1Account
	return r
}

func (r ApiAccountServiceCreateAccountRequest) Execute() (*V1Account, *http.Response, error) {
	return r.ApiService.AccountServiceCreateAccountExecute(r)
}

/*
AccountServiceCreateAccount CreateAccount API creates a new Account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAccountServiceCreateAccountRequest
*/
func (a *AccountServiceApiService) AccountServiceCreateAccount(ctx context.Context) ApiAccountServiceCreateAccountRequest {
	return ApiAccountServiceCreateAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1Account
func (a *AccountServiceApiService) AccountServiceCreateAccountExecute(r ApiAccountServiceCreateAccountRequest) (*V1Account, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountServiceApiService.AccountServiceCreateAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.v1Account == nil {
		return localVarReturnValue, nil, reportError("v1Account is required and must be specified")
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
	localVarPostBody = r.v1Account
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

type ApiAccountServiceDeleteAccountRequest struct {
	ctx context.Context
	ApiService *AccountServiceApiService
	accountId string
}

func (r ApiAccountServiceDeleteAccountRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AccountServiceDeleteAccountExecute(r)
}

/*
AccountServiceDeleteAccount DeleteAccount API deletes the specified account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId account id for which account details need to be fetched
 @return ApiAccountServiceDeleteAccountRequest
*/
func (a *AccountServiceApiService) AccountServiceDeleteAccount(ctx context.Context, accountId string) ApiAccountServiceDeleteAccountRequest {
	return ApiAccountServiceDeleteAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AccountServiceApiService) AccountServiceDeleteAccountExecute(r ApiAccountServiceDeleteAccountRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountServiceApiService.AccountServiceDeleteAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts/{accountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)

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

type ApiAccountServiceGetAccountRequest struct {
	ctx context.Context
	ApiService *AccountServiceApiService
	accountId string
}

func (r ApiAccountServiceGetAccountRequest) Execute() (*V1Account, *http.Response, error) {
	return r.ApiService.AccountServiceGetAccountExecute(r)
}

/*
AccountServiceGetAccount GetAccount API returns the info about account for given account id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId account id for which account details need to be fetched
 @return ApiAccountServiceGetAccountRequest
*/
func (a *AccountServiceApiService) AccountServiceGetAccount(ctx context.Context, accountId string) ApiAccountServiceGetAccountRequest {
	return ApiAccountServiceGetAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return V1Account
func (a *AccountServiceApiService) AccountServiceGetAccountExecute(r ApiAccountServiceGetAccountRequest) (*V1Account, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountServiceApiService.AccountServiceGetAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts/{accountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)

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

type ApiAccountServiceListAccountsRequest struct {
	ctx context.Context
	ApiService *AccountServiceApiService
	paginationPageNumber *string
	paginationPageSize *string
}

// Page number is the page number to return based on the size
func (r ApiAccountServiceListAccountsRequest) PaginationPageNumber(paginationPageNumber string) ApiAccountServiceListAccountsRequest {
	r.paginationPageNumber = &paginationPageNumber
	return r
}

// Page size is the maximum number of records to include per page
func (r ApiAccountServiceListAccountsRequest) PaginationPageSize(paginationPageSize string) ApiAccountServiceListAccountsRequest {
	r.paginationPageSize = &paginationPageSize
	return r
}

func (r ApiAccountServiceListAccountsRequest) Execute() (*V1ListAccountsResponse, *http.Response, error) {
	return r.ApiService.AccountServiceListAccountsExecute(r)
}

/*
AccountServiceListAccounts ListAccounts API lists the accounts visible to the caller

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAccountServiceListAccountsRequest
*/
func (a *AccountServiceApiService) AccountServiceListAccounts(ctx context.Context) ApiAccountServiceListAccountsRequest {
	return ApiAccountServiceListAccountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1ListAccountsResponse
func (a *AccountServiceApiService) AccountServiceListAccountsExecute(r ApiAccountServiceListAccountsRequest) (*V1ListAccountsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1ListAccountsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountServiceApiService.AccountServiceListAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiAccountServiceUpdateAccountRequest struct {
	ctx context.Context
	ApiService *AccountServiceApiService
	accountMetaUid string
	accountForWhichNameNeedsToBeUpdated *AccountForWhichNameNeedsToBeUpdated
}

// account  for which name needs to be updated
func (r ApiAccountServiceUpdateAccountRequest) AccountForWhichNameNeedsToBeUpdated(accountForWhichNameNeedsToBeUpdated AccountForWhichNameNeedsToBeUpdated) ApiAccountServiceUpdateAccountRequest {
	r.accountForWhichNameNeedsToBeUpdated = &accountForWhichNameNeedsToBeUpdated
	return r
}

func (r ApiAccountServiceUpdateAccountRequest) Execute() (*V1Account, *http.Response, error) {
	return r.ApiService.AccountServiceUpdateAccountExecute(r)
}

/*
AccountServiceUpdateAccount UpdateAccount API updates the specified account.This API requires the caller to first get an Account from GetAccount().

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountMetaUid UID of the resource of the format <resource prefix>-<uuid>.
 @return ApiAccountServiceUpdateAccountRequest
*/
func (a *AccountServiceApiService) AccountServiceUpdateAccount(ctx context.Context, accountMetaUid string) ApiAccountServiceUpdateAccountRequest {
	return ApiAccountServiceUpdateAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountMetaUid: accountMetaUid,
	}
}

// Execute executes the request
//  @return V1Account
func (a *AccountServiceApiService) AccountServiceUpdateAccountExecute(r ApiAccountServiceUpdateAccountRequest) (*V1Account, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountServiceApiService.AccountServiceUpdateAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/accounts/{account.meta.uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"account.meta.uid"+"}", url.PathEscape(parameterToString(r.accountMetaUid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountForWhichNameNeedsToBeUpdated == nil {
		return localVarReturnValue, nil, reportError("accountForWhichNameNeedsToBeUpdated is required and must be specified")
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
	localVarPostBody = r.accountForWhichNameNeedsToBeUpdated
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
