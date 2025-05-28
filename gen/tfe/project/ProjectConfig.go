package project

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectConfig struct {
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
	// Name of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/project#name Project#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Duration after which the project will be auto-destroyed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/project#auto_destroy_activity_duration Project#auto_destroy_activity_duration}
	AutoDestroyActivityDuration *string `field:"optional" json:"autoDestroyActivityDuration" yaml:"autoDestroyActivityDuration"`
	// Description of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/project#description Project#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Explicitly ignores tags created outside of Terraform so they will not be overwritten by tags defined in configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/project#ignore_additional_tags Project#ignore_additional_tags}
	IgnoreAdditionalTags interface{} `field:"optional" json:"ignoreAdditionalTags" yaml:"ignoreAdditionalTags"`
	// Name of the organization to which the project belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/project#organization Project#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// A map of key-value tags to add to the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/project#tags Project#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

