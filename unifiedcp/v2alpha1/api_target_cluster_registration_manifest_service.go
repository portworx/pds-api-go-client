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


// TargetClusterRegistrationManifestServiceApiService TargetClusterRegistrationManifestServiceApi service
type TargetClusterRegistrationManifestServiceApiService service

type ApiTargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestRequest struct {
	ctx context.Context
	ApiService *TargetClusterRegistrationManifestServiceApiService
	tenantId string
	targetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody *TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody
}

func (r ApiTargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestRequest) TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody(targetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody) ApiTargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestRequest {
	r.targetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody = &targetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody
	return r
}

func (r ApiTargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestRequest) Execute() (*V1TargetClusterRegistrationManifest, *http.Response, error) {
	return r.ApiService.TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestExecute(r)
}

/*
TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifest GetTargetClusterRegistrationManifest fetches registration manifest for the given request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tenantId tenanat_id is the id of the tenant for which manifest is requested
 @return ApiTargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestRequest
*/
func (a *TargetClusterRegistrationManifestServiceApiService) TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifest(ctx context.Context, tenantId string) ApiTargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestRequest {
	return ApiTargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestRequest{
		ApiService: a,
		ctx: ctx,
		tenantId: tenantId,
	}
}

// Execute executes the request
//  @return V1TargetClusterRegistrationManifest
func (a *TargetClusterRegistrationManifestServiceApiService) TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestExecute(r ApiTargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestRequest) (*V1TargetClusterRegistrationManifest, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1TargetClusterRegistrationManifest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetClusterRegistrationManifestServiceApiService.TargetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants/{tenantId}:registrationManifests"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantId"+"}", url.PathEscape(parameterToString(r.tenantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.targetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody == nil {
		return localVarReturnValue, nil, reportError("targetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody is required and must be specified")
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
	localVarPostBody = r.targetClusterRegistrationManifestServiceGenerateTargetClusterRegistrationManifestBody
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
