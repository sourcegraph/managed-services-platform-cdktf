package zerotrustdevicesubnet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDeviceSubnetConfig struct {
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
	// A user-friendly name for the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_device_subnet#name ZeroTrustDeviceSubnet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_device_subnet#network ZeroTrustDeviceSubnet#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Cloudflare account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_device_subnet#account_id ZeroTrustDeviceSubnet#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// An optional description of the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_device_subnet#comment ZeroTrustDeviceSubnet#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// If `true`, this is the default subnet for the account. There can only be one default subnet per account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_device_subnet#is_default_network ZeroTrustDeviceSubnet#is_default_network}
	IsDefaultNetwork interface{} `field:"optional" json:"isDefaultNetwork" yaml:"isDefaultNetwork"`
}

