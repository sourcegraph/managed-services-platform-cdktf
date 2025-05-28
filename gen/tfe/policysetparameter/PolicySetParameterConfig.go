package policysetparameter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicySetParameterConfig struct {
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
	// Name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/policy_set_parameter#key PolicySetParameter#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The ID of the policy set that owns the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/policy_set_parameter#policy_set_id PolicySetParameter#policy_set_id}
	PolicySetId *string `field:"required" json:"policySetId" yaml:"policySetId"`
	// Whether the value is sensitive. If true then the parameter is written once and not visible thereafter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/policy_set_parameter#sensitive PolicySetParameter#sensitive}
	Sensitive interface{} `field:"optional" json:"sensitive" yaml:"sensitive"`
	// Value of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/policy_set_parameter#value PolicySetParameter#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// Value of the parameter in write-only mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/policy_set_parameter#value_wo PolicySetParameter#value_wo}
	ValueWo *string `field:"optional" json:"valueWo" yaml:"valueWo"`
}

