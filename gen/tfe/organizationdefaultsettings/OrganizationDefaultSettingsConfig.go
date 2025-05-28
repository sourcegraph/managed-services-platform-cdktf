package organizationdefaultsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationDefaultSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/organization_default_settings#default_execution_mode OrganizationDefaultSettings#default_execution_mode}.
	DefaultExecutionMode *string `field:"required" json:"defaultExecutionMode" yaml:"defaultExecutionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/organization_default_settings#default_agent_pool_id OrganizationDefaultSettings#default_agent_pool_id}.
	DefaultAgentPoolId *string `field:"optional" json:"defaultAgentPoolId" yaml:"defaultAgentPoolId"`
	// The name of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/organization_default_settings#organization OrganizationDefaultSettings#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
}

