package googleaccesscontextmanageraccesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleAccessContextManagerAccessPolicyConfig struct {
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
	// The parent of this AccessPolicy in the Cloud Resource Hierarchy. Format: 'organizations/{{organization_id}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_access_policy#parent GoogleAccessContextManagerAccessPolicy#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Human readable title. Does not affect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_access_policy#title GoogleAccessContextManagerAccessPolicy#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_access_policy#id GoogleAccessContextManagerAccessPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Folder or project on which this policy is applicable. Format: 'folders/{{folder_id}}' or 'projects/{{project_number}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_access_policy#scopes GoogleAccessContextManagerAccessPolicy#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_access_policy#timeouts GoogleAccessContextManagerAccessPolicy#timeouts}
	Timeouts *GoogleAccessContextManagerAccessPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

