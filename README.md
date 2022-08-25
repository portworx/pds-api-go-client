# PDS API Go Client

This is a Go client for the Portworx Data Services [API](https://prod.pds.portworx.com/swagger/index.html).

Iinformation about Portworx Data Services in general can be found [here](https://portworx.com/products/portworx-data-services/).

## Documentation

Full documentation with installation and examples can be found [here](https://github.com/portworx/pds-api-go-client/blob/master/pds/v1alpha1/README.md).




## Manually generating the client

It's possible to manually generate this client using the following command:

```bash
docker run --rm -v ${PWD}/pds/v1alpha1:/local openapitools/openapi-generator-cli generate -i /local/api/swagger.json -g go -o /local/pds/v1alpha1 --package-name pds
```

