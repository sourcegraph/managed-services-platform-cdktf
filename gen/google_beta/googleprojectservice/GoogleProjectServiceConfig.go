package googleprojectservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleProjectServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_service#service GoogleProjectService#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_service#check_if_service_has_usage_on_destroy GoogleProjectService#check_if_service_has_usage_on_destroy}.
	CheckIfServiceHasUsageOnDestroy interface{} `field:"optional" json:"checkIfServiceHasUsageOnDestroy" yaml:"checkIfServiceHasUsageOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_service#disable_dependent_services GoogleProjectService#disable_dependent_services}.
	DisableDependentServices interface{} `field:"optional" json:"disableDependentServices" yaml:"disableDependentServices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_service#disable_on_destroy GoogleProjectService#disable_on_destroy}.
	DisableOnDestroy interface{} `field:"optional" json:"disableOnDestroy" yaml:"disableOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_service#id GoogleProjectService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_service#project GoogleProjectService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_project_service#timeouts GoogleProjectService#timeouts}
	Timeouts *GoogleProjectServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

