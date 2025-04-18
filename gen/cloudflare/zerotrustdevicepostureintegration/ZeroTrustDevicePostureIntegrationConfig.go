package zerotrustdevicepostureintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDevicePostureIntegrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_device_posture_integration#account_id ZeroTrustDevicePostureIntegration#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the device posture integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_device_posture_integration#name ZeroTrustDevicePostureIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The device posture integration type. Available values: `workspace_one`, `uptycs`, `crowdstrike_s2s`, `intune`, `kolide`, `sentinelone_s2s`, `tanium_s2s`, `custom_s2s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_device_posture_integration#type ZeroTrustDevicePostureIntegration#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_device_posture_integration#config ZeroTrustDevicePostureIntegration#config}
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_device_posture_integration#id ZeroTrustDevicePostureIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_device_posture_integration#identifier ZeroTrustDevicePostureIntegration#identifier}.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// Indicates the frequency with which to poll the third-party API. Must be in the format `1h` or `30m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_device_posture_integration#interval ZeroTrustDevicePostureIntegration#interval}
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
}

