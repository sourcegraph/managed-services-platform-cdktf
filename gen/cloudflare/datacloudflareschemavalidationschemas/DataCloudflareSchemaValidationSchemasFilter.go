package datacloudflareschemavalidationschemas


type DataCloudflareSchemaValidationSchemasFilter struct {
	// Filter for enabled schemas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/schema_validation_schemas#validation_enabled DataCloudflareSchemaValidationSchemas#validation_enabled}
	ValidationEnabled interface{} `field:"optional" json:"validationEnabled" yaml:"validationEnabled"`
}

