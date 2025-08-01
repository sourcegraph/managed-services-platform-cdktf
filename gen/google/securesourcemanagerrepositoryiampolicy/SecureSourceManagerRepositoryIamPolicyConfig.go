package securesourcemanagerrepositoryiampolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecureSourceManagerRepositoryIamPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/secure_source_manager_repository_iam_policy#policy_data SecureSourceManagerRepositoryIamPolicy#policy_data}.
	PolicyData *string `field:"required" json:"policyData" yaml:"policyData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/secure_source_manager_repository_iam_policy#repository_id SecureSourceManagerRepositoryIamPolicy#repository_id}.
	RepositoryId *string `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/secure_source_manager_repository_iam_policy#id SecureSourceManagerRepositoryIamPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/secure_source_manager_repository_iam_policy#location SecureSourceManagerRepositoryIamPolicy#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/secure_source_manager_repository_iam_policy#project SecureSourceManagerRepositoryIamPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

