/*
public/portworx/pds/catalog/dataservices/apiv1/dataservices.proto

Testing DataServiceVersionServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package catalog

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/portworx/pds-api-go-client/unifiedcp/v2alpha2/clients/catalog"
)

func Test_catalog_DataServiceVersionServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DataServiceVersionServiceAPIService DataServiceVersionServiceGetDataServiceVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DataServiceVersionServiceAPI.DataServiceVersionServiceGetDataServiceVersion(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataServiceVersionServiceAPIService DataServiceVersionServiceListCompatibleDataServiceVersions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DataServiceVersionServiceAPI.DataServiceVersionServiceListCompatibleDataServiceVersions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataServiceVersionServiceAPIService DataServiceVersionServiceListDataServiceVersions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dataServiceId string

		resp, httpRes, err := apiClient.DataServiceVersionServiceAPI.DataServiceVersionServiceListDataServiceVersions(context.Background(), dataServiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
