package datacloudflarezerotrustdeviceposturerules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustDevicePostureRulesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zero_trust_device_posture_rules#account_id DataCloudflareZeroTrustDevicePostureRules#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zero_trust_device_posture_rules#max_items DataCloudflareZeroTrustDevicePostureRules#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
}

