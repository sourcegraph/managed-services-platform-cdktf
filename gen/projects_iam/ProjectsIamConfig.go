package projects_iam

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectsIamConfig struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// Map of role (key) and list of members (value) to add the IAM policies/bindings.
	// Default: [object Object]
	// The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}
	//
	Bindings *map[string]*[]*string `field:"optional" json:"bindings" yaml:"bindings"`
	// List of maps of role and respective conditions, and the members to add the IAM policies/bindings.
	ConditionalBindings interface{} `field:"optional" json:"conditionalBindings" yaml:"conditionalBindings"`
	// Mode for adding the IAM policies/bindings, additive and authoritative.
	// Default: additive.
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Projects list to add the IAM policies/bindings.
	Projects *[]*string `field:"optional" json:"projects" yaml:"projects"`
}

