package datatfeorganizationruntaskglobalsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataTfeOrganizationRunTaskGlobalSettingsConfig struct {
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
	// The id of the run task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/organization_run_task_global_settings#task_id DataTfeOrganizationRunTaskGlobalSettings#task_id}
	TaskId *string `field:"required" json:"taskId" yaml:"taskId"`
	// Whether the run task will be applied globally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/organization_run_task_global_settings#enabled DataTfeOrganizationRunTaskGlobalSettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The enforcement level of the global task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/organization_run_task_global_settings#enforcement_level DataTfeOrganizationRunTaskGlobalSettings#enforcement_level}
	EnforcementLevel *string `field:"optional" json:"enforcementLevel" yaml:"enforcementLevel"`
	// Which stages the task will run in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/organization_run_task_global_settings#stages DataTfeOrganizationRunTaskGlobalSettings#stages}
	Stages *[]*string `field:"optional" json:"stages" yaml:"stages"`
}

