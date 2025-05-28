package teammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamMemberConfig struct {
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
	// The ID of the member to add to the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/team_member#member_id TeamMember#member_id}
	MemberId *string `field:"required" json:"memberId" yaml:"memberId"`
	// The organization of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/team_member#organization TeamMember#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The slug of the team to add the member to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/team_member#team TeamMember#team}
	Team *string `field:"required" json:"team" yaml:"team"`
	// The role of the member in the team.
	//
	// When not set, resolve to the minimum team role given by this member's organization role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/team_member#role TeamMember#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

