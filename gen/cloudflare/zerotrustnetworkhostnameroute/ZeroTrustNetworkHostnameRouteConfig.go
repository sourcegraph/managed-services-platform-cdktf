package zerotrustnetworkhostnameroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustNetworkHostnameRouteConfig struct {
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
	// Cloudflare account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_network_hostname_route#account_id ZeroTrustNetworkHostnameRoute#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// An optional description of the hostname route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_network_hostname_route#comment ZeroTrustNetworkHostnameRoute#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The hostname of the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_network_hostname_route#hostname ZeroTrustNetworkHostnameRoute#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// UUID of the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_network_hostname_route#tunnel_id ZeroTrustNetworkHostnameRoute#tunnel_id}
	TunnelId *string `field:"optional" json:"tunnelId" yaml:"tunnelId"`
}

