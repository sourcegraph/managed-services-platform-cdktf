package zerotrustdnslocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDnsLocationConfig struct {
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
	// Specify the location name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dns_location#name ZeroTrustDnsLocation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dns_location#account_id ZeroTrustDnsLocation#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Indicate whether this location is the default location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dns_location#client_default ZeroTrustDnsLocation#client_default}
	ClientDefault interface{} `field:"optional" json:"clientDefault" yaml:"clientDefault"`
	// Specify the identifier of the pair of IPv4 addresses assigned to this location.
	//
	// When creating a location, if this field is absent or set to null, the pair of shared IPv4 addresses (0e4a32c6-6fb8-4858-9296-98f51631e8e6) is auto-assigned. When updating a location, if this field is absent or set to null, the pre-assigned pair remains unchanged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dns_location#dns_destination_ips_id ZeroTrustDnsLocation#dns_destination_ips_id}
	DnsDestinationIpsId *string `field:"optional" json:"dnsDestinationIpsId" yaml:"dnsDestinationIpsId"`
	// Indicate whether the location must resolve EDNS queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dns_location#ecs_support ZeroTrustDnsLocation#ecs_support}
	EcsSupport interface{} `field:"optional" json:"ecsSupport" yaml:"ecsSupport"`
	// Configure the destination endpoints for this location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dns_location#endpoints ZeroTrustDnsLocation#endpoints}
	Endpoints *ZeroTrustDnsLocationEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Specify the list of network ranges from which requests at this location originate.
	//
	// The list takes effect only if it is non-empty and the IPv4 endpoint is enabled for this location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dns_location#networks ZeroTrustDnsLocation#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
}

