package organizationrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationRepositoryConfig struct {
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
	// The identifier of the repository.
	//
	// For GitHub, GitLab and BitBucket, it is `{organization}/{repository}`. For VSTS, it is the [repository ID](https://learn.microsoft.com/en-us/rest/api/azure/devops/git/repositories/get#get-a-repository-by-repositoryid).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_repository#identifier OrganizationRepository#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The ID of the organization integration. Source from the URL `https://<organization>.sentry.io/settings/integrations/<integration-type>/<integration-id>/` or use the `sentry_organization_integration` data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_repository#integration_id OrganizationRepository#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The type of the organization integration. Supported values are `github`, `github_enterprise`, `gitlab`, `vsts` (Azure DevOps), `bitbucket`, and `bitbucket_server`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_repository#integration_type OrganizationRepository#integration_type}
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
	// The organization of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_repository#organization OrganizationRepository#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
}

