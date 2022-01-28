package examples

import (
	"context"
	"fmt"
	"net/url"

	pds "github.com/outwarped/pds-api-go-client/pds/v1alpha"
)

func example_auth_01() {
	pds_cluster_name := "my-pds-cluster"
	pds_cluster_sa_token := "service_account_token_string"

	endpointUrl, _ := url.Parse(fmt.Sprintf("https://%s.pds-dev.io", pds_cluster_name))

	apiConf := pds.NewConfiguration()
	apiConf.Host = endpointUrl.Host
	apiConf.Scheme = endpointUrl.Scheme

	// Use Configuration or context with WithValue (below)
	apiConf.DefaultHeader = map[string]string{
		"Authorization": "Bearer " + pds_cluster_sa_token,
	}
	// Use Configuration or context with WithValue (above)
	ctx := context.WithValue(context.Background(), pds.ContextAPIKeys, map[string]pds.APIKey{"ApiKeyAuth": {Key: pds_cluster_sa_token, Prefix: "Bearer"}})

	apiClient := pds.NewAPIClient(apiConf)
	authClient := apiClient.AuthorizerApi

	req := pds.ModelsAuthorizerRequest{}
	req.SetEntityType("entity_type")
	req.SetEntityId("entity_id")
	req.SetOperation("some_operation")

	model, res, err := authClient.ApiAuthorizerPost(ctx).Body(req).Execute()
	if err != nil {
		panic(err)
	}
	if res.StatusCode >= 300 {
		panic(res.Status)
	}
	if allowed, ok := model.GetAllowOk(); !ok || !*allowed {
		panic("not allowed")
	}
}
