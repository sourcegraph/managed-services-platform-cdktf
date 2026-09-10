package hyperdriveconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HyperdriveConfigConfig struct {
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
	// The name of the Hyperdrive configuration. Used to identify the configuration in the Cloudflare dashboard and API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#name HyperdriveConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#origin HyperdriveConfig#origin}.
	Origin *HyperdriveConfigOrigin `field:"required" json:"origin" yaml:"origin"`
	// Define configurations using a unique string identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#account_id HyperdriveConfig#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#caching HyperdriveConfig#caching}.
	Caching *HyperdriveConfigCaching `field:"optional" json:"caching" yaml:"caching"`
	// mTLS configuration for the origin connection.
	//
	// Cannot be used with VPC Service origins; TLS must be managed on the VPC Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#mtls HyperdriveConfig#mtls}
	Mtls *HyperdriveConfigMtls `field:"optional" json:"mtls" yaml:"mtls"`
	// The (soft) maximum number of connections the Hyperdrive is allowed to make to the origin database.
	//
	// Maximum allowed: 20 for free tier accounts, 100 for paid tier accounts.
	// If not specified, defaults to 20 for free tier and 60 for paid tier.
	// Contact Cloudflare if you need a higher limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hyperdrive_config#origin_connection_limit HyperdriveConfig#origin_connection_limit}
	OriginConnectionLimit *float64 `field:"optional" json:"originConnectionLimit" yaml:"originConnectionLimit"`
}

