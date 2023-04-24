/*
PDS API

Testing AccountsRoleInvitationsApiService

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

func Test_pds_AccountsRoleInvitationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccountsRoleInvitationsApiService ApiAccountRoleInvitationsIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AccountsRoleInvitationsApi.ApiAccountRoleInvitationsIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountsRoleInvitationsApiService ApiAccountRoleInvitationsIdPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AccountsRoleInvitationsApi.ApiAccountRoleInvitationsIdPatch(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
