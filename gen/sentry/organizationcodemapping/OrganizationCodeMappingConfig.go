package organizationcodemapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationCodeMappingConfig struct {
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
	// Default branch of your code we fall back to if you do not have commit tracking set up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_code_mapping#default_branch OrganizationCodeMapping#default_branch}
	DefaultBranch *string `field:"required" json:"defaultBranch" yaml:"defaultBranch"`
	// Sentry Organization Integration ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_code_mapping#integration_id OrganizationCodeMapping#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The slug of the organization the code mapping is under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_code_mapping#organization OrganizationCodeMapping#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Sentry Project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_code_mapping#project_id OrganizationCodeMapping#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Sentry Organization Repository ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_code_mapping#repository_id OrganizationCodeMapping#repository_id}
	RepositoryId *string `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_code_mapping#id OrganizationCodeMapping#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// https://docs.sentry.io/product/integrations/source-code-mgmt/github/#stack-trace-linking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_code_mapping#source_root OrganizationCodeMapping#source_root}
	SourceRoot *string `field:"optional" json:"sourceRoot" yaml:"sourceRoot"`
	// https://docs.sentry.io/product/integrations/source-code-mgmt/github/#stack-trace-linking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/organization_code_mapping#stack_root OrganizationCodeMapping#stack_root}
	StackRoot *string `field:"optional" json:"stackRoot" yaml:"stackRoot"`
}

