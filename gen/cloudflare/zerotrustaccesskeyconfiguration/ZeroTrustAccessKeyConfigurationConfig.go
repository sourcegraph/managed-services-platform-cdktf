package zerotrustaccesskeyconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessKeyConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_key_configuration#account_id ZeroTrustAccessKeyConfiguration#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The number of days between key rotations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_key_configuration#key_rotation_interval_days ZeroTrustAccessKeyConfiguration#key_rotation_interval_days}
	KeyRotationIntervalDays *float64 `field:"required" json:"keyRotationIntervalDays" yaml:"keyRotationIntervalDays"`
}

