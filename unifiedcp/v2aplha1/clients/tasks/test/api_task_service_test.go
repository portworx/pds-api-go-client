/*
public/portworx/pds/tasks/apiv1/tasks.proto

Testing TaskServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tasks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/portworx/pds-api-go-client/unifiedcp/v2alpha2/clients/tasks"
)

func Test_tasks_TaskServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaskServiceAPIService TaskServiceGetTask", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TaskServiceAPI.TaskServiceGetTask(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskServiceAPIService TaskServiceListTasks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TaskServiceAPI.TaskServiceListTasks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
