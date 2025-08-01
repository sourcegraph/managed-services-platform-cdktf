package googleprojectdefaultserviceaccounts

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleProjectDefaultServiceAccountsConfig struct {
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
	// The action to be performed in the default service accounts.
	//
	// Valid values are: DEPRIVILEGE, DELETE, DISABLE.
	// 				Note that DEPRIVILEGE action will ignore the REVERT configuration in the restore_policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_default_service_accounts#action GoogleProjectDefaultServiceAccounts#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The project ID where service accounts are created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_default_service_accounts#project GoogleProjectDefaultServiceAccounts#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_default_service_accounts#id GoogleProjectDefaultServiceAccounts#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The action to be performed in the default service accounts on the resource destroy.
	//
	// Valid values are NONE, REVERT and REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_default_service_accounts#restore_policy GoogleProjectDefaultServiceAccounts#restore_policy}
	RestorePolicy *string `field:"optional" json:"restorePolicy" yaml:"restorePolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_default_service_accounts#timeouts GoogleProjectDefaultServiceAccounts#timeouts}
	Timeouts *GoogleProjectDefaultServiceAccountsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

