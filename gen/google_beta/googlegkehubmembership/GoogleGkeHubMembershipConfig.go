package googlegkehubmembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGkeHubMembershipConfig struct {
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
	// The client-provided identifier of the membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership#membership_id GoogleGkeHubMembership#membership_id}
	MembershipId *string `field:"required" json:"membershipId" yaml:"membershipId"`
	// authority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership#authority GoogleGkeHubMembership#authority}
	Authority *GoogleGkeHubMembershipAuthority `field:"optional" json:"authority" yaml:"authority"`
	// The name of this entity type to be displayed on the console.
	//
	// This field is unavailable in v1 of the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership#description GoogleGkeHubMembership#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership#endpoint GoogleGkeHubMembership#endpoint}
	Endpoint *GoogleGkeHubMembershipEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership#id GoogleGkeHubMembership#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels to apply to this membership.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership#labels GoogleGkeHubMembership#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Location of the membership. The default value is 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership#location GoogleGkeHubMembership#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership#project GoogleGkeHubMembership#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership#timeouts GoogleGkeHubMembership#timeouts}
	Timeouts *GoogleGkeHubMembershipTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

