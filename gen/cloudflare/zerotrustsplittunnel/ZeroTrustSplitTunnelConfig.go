package zerotrustsplittunnel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustSplitTunnelConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_split_tunnel#account_id ZeroTrustSplitTunnel#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The mode of the split tunnel policy. Available values: `include`, `exclude`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_split_tunnel#mode ZeroTrustSplitTunnel#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// tunnels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_split_tunnel#tunnels ZeroTrustSplitTunnel#tunnels}
	Tunnels interface{} `field:"required" json:"tunnels" yaml:"tunnels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_split_tunnel#id ZeroTrustSplitTunnel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The settings policy for which to configure this split tunnel policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_split_tunnel#policy_id ZeroTrustSplitTunnel#policy_id}
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
}

