package teamslocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamsLocationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_location#account_id TeamsLocation#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the teams location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_location#name TeamsLocation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicator that this is the default location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_location#client_default TeamsLocation#client_default}
	ClientDefault interface{} `field:"optional" json:"clientDefault" yaml:"clientDefault"`
	// IPv4 binding assigned to this location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_location#dns_destination_ips_id TeamsLocation#dns_destination_ips_id}
	DnsDestinationIpsId *string `field:"optional" json:"dnsDestinationIpsId" yaml:"dnsDestinationIpsId"`
	// IPv6 block binding assigned to this location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_location#dns_destination_ipv6_block_id TeamsLocation#dns_destination_ipv6_block_id}
	DnsDestinationIpv6BlockId *string `field:"optional" json:"dnsDestinationIpv6BlockId" yaml:"dnsDestinationIpv6BlockId"`
	// Indicator that this location needs to resolve EDNS queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_location#ecs_support TeamsLocation#ecs_support}
	EcsSupport interface{} `field:"optional" json:"ecsSupport" yaml:"ecsSupport"`
	// endpoints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_location#endpoints TeamsLocation#endpoints}
	Endpoints *TeamsLocationEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_location#id TeamsLocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The networks CIDRs that comprise the location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_location#networks TeamsLocation#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
}

