package googleorgpolicypolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleOrgPolicyPolicyConfig struct {
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
	// Immutable.
	//
	// The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * 'projects/{project_number}/policies/{constraint_name}' * 'folders/{folder_id}/policies/{constraint_name}' * 'organizations/{organization_id}/policies/{constraint_name}' For example, "projects/123/policies/compute.disableSerialPortAccess". Note: 'projects/{project_id}/policies/{constraint_name}' is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_org_policy_policy#name GoogleOrgPolicyPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The parent of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_org_policy_policy#parent GoogleOrgPolicyPolicy#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// dry_run_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_org_policy_policy#dry_run_spec GoogleOrgPolicyPolicy#dry_run_spec}
	DryRunSpec *GoogleOrgPolicyPolicyDryRunSpec `field:"optional" json:"dryRunSpec" yaml:"dryRunSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_org_policy_policy#id GoogleOrgPolicyPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_org_policy_policy#spec GoogleOrgPolicyPolicy#spec}
	Spec *GoogleOrgPolicyPolicySpec `field:"optional" json:"spec" yaml:"spec"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_org_policy_policy#timeouts GoogleOrgPolicyPolicy#timeouts}
	Timeouts *GoogleOrgPolicyPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

