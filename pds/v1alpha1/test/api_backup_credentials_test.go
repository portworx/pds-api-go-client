/*
PDS API

Testing BackupCredentialsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pds

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_pds_BackupCredentialsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BackupCredentialsApiService ApiBackupCredentialsIdCredentialsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BackupCredentialsApi.ApiBackupCredentialsIdCredentialsGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupCredentialsApiService ApiBackupCredentialsIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.BackupCredentialsApi.ApiBackupCredentialsIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupCredentialsApiService ApiBackupCredentialsIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BackupCredentialsApi.ApiBackupCredentialsIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupCredentialsApiService ApiBackupCredentialsIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BackupCredentialsApi.ApiBackupCredentialsIdPut(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupCredentialsApiService ApiTenantsIdBackupCredentialsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BackupCredentialsApi.ApiTenantsIdBackupCredentialsGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupCredentialsApiService ApiTenantsIdBackupCredentialsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BackupCredentialsApi.ApiTenantsIdBackupCredentialsPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
