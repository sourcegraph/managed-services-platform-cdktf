package googlegkehubmembershipbinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGkeHubMembershipBindingConfig struct {
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
	// Location of the membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership_binding#location GoogleGkeHubMembershipBinding#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The client-provided identifier of the membership binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership_binding#membership_binding_id GoogleGkeHubMembershipBinding#membership_binding_id}
	MembershipBindingId *string `field:"required" json:"membershipBindingId" yaml:"membershipBindingId"`
	// Id of the membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership_binding#membership_id GoogleGkeHubMembershipBinding#membership_id}
	MembershipId *string `field:"required" json:"membershipId" yaml:"membershipId"`
	// A Workspace resource name in the format 'projects/* /locations/* /scopes/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership_binding#scope GoogleGkeHubMembershipBinding#scope}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership_binding#id GoogleGkeHubMembershipBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels for this Membership binding.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership_binding#labels GoogleGkeHubMembershipBinding#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership_binding#project GoogleGkeHubMembershipBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership_binding#timeouts GoogleGkeHubMembershipBinding#timeouts}
	Timeouts *GoogleGkeHubMembershipBindingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

