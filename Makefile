bundle-openapi:
	swagger-cli bundle openapi/openapi.yml --outfile pkg/client/openapi.yml --type yaml
generate-client:
	oapi-codegen --config pkg/client/oapi.config.yml pkg/client/openapi.yml
build-client: bundle-openapi generate-client