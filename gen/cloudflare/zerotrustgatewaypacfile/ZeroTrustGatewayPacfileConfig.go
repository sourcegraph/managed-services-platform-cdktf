package zerotrustgatewaypacfile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewayPacfileConfig struct {
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
	// Actual contents of the PAC file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_pacfile#contents ZeroTrustGatewayPacfile#contents}
	Contents *string `field:"required" json:"contents" yaml:"contents"`
	// Name of the PAC file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_pacfile#name ZeroTrustGatewayPacfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_pacfile#account_id ZeroTrustGatewayPacfile#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Detailed description of the PAC file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_pacfile#description ZeroTrustGatewayPacfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// URL-friendly version of the PAC file name. If not provided, it will be auto-generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_pacfile#slug ZeroTrustGatewayPacfile#slug}
	Slug *string `field:"optional" json:"slug" yaml:"slug"`
}

