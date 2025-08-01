package datacloudflareschemavalidationschemas


type DataCloudflareSchemaValidationSchemasFilter struct {
	// Omit the source-files of schemas and only retrieve their meta-data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/schema_validation_schemas#omit_source DataCloudflareSchemaValidationSchemas#omit_source}
	OmitSource interface{} `field:"optional" json:"omitSource" yaml:"omitSource"`
	// Filter for enabled schemas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/schema_validation_schemas#validation_enabled DataCloudflareSchemaValidationSchemas#validation_enabled}
	ValidationEnabled interface{} `field:"optional" json:"validationEnabled" yaml:"validationEnabled"`
}

