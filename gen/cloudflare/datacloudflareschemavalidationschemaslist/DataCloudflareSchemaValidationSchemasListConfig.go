package datacloudflareschemavalidationschemaslist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareSchemaValidationSchemasListConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/schema_validation_schemas_list#zone_id DataCloudflareSchemaValidationSchemasList#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/schema_validation_schemas_list#max_items DataCloudflareSchemaValidationSchemasList#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Omit the source-files of schemas and only retrieve their meta-data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/schema_validation_schemas_list#omit_source DataCloudflareSchemaValidationSchemasList#omit_source}
	OmitSource interface{} `field:"optional" json:"omitSource" yaml:"omitSource"`
	// Filter for enabled schemas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/schema_validation_schemas_list#validation_enabled DataCloudflareSchemaValidationSchemasList#validation_enabled}
	ValidationEnabled interface{} `field:"optional" json:"validationEnabled" yaml:"validationEnabled"`
}

