package datacloudflareclientcertificates

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareClientCertificatesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/client_certificates#limit DataCloudflareClientCertificates#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/client_certificates#max_items DataCloudflareClientCertificates#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Offset the results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/client_certificates#offset DataCloudflareClientCertificates#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Client Certitifcate Status to filter results by. Available values: "all", "active", "pending_reactivation", "pending_revocation", "revoked".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/client_certificates#status DataCloudflareClientCertificates#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/client_certificates#zone_id DataCloudflareClientCertificates#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

