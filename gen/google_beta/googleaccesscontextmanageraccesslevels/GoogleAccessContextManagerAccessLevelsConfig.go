package googleaccesscontextmanageraccesslevels

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleAccessContextManagerAccessLevelsConfig struct {
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
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_access_levels#parent GoogleAccessContextManagerAccessLevels#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// access_levels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_access_levels#access_levels GoogleAccessContextManagerAccessLevels#access_levels}
	AccessLevels interface{} `field:"optional" json:"accessLevels" yaml:"accessLevels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_access_levels#id GoogleAccessContextManagerAccessLevels#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_access_levels#timeouts GoogleAccessContextManagerAccessLevels#timeouts}
	Timeouts *GoogleAccessContextManagerAccessLevelsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

