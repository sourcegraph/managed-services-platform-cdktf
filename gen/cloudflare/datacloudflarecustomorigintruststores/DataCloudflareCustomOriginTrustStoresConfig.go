package datacloudflarecustomorigintruststores

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareCustomOriginTrustStoresConfig struct {
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
	// Limit to the number of records returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_origin_trust_stores#limit DataCloudflareCustomOriginTrustStores#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_origin_trust_stores#max_items DataCloudflareCustomOriginTrustStores#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Offset the results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_origin_trust_stores#offset DataCloudflareCustomOriginTrustStores#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_origin_trust_stores#zone_id DataCloudflareCustomOriginTrustStores#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

