package teamtoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamTokenConfig struct {
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
	// ID of the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_token#team_id TeamToken#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// The description of the token, which must be unique per team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_token#description TeamToken#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The token's expiration date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_token#expired_at TeamToken#expired_at}
	ExpiredAt *string `field:"optional" json:"expiredAt" yaml:"expiredAt"`
	// When set to true will force the audit trail token to be recreated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_token#force_regenerate TeamToken#force_regenerate}
	ForceRegenerate interface{} `field:"optional" json:"forceRegenerate" yaml:"forceRegenerate"`
}

