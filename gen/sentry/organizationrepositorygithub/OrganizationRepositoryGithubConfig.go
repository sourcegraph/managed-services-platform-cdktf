package organizationrepositorygithub

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationRepositoryGithubConfig struct {
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
	// The repo identifier. For Github it is {github_org}/{github_repo}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/organization_repository_github#identifier OrganizationRepositoryGithub#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The organization integration ID for Github.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/organization_repository_github#integration_id OrganizationRepositoryGithub#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The slug of the Sentry organization this resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/organization_repository_github#organization OrganizationRepositoryGithub#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/organization_repository_github#id OrganizationRepositoryGithub#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

