package datatferegistrymodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataTfeRegistryModuleConfig struct {
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
	// Name of the module provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module#module_provider DataTfeRegistryModule#module_provider}
	ModuleProvider *string `field:"required" json:"moduleProvider" yaml:"moduleProvider"`
	// Name of the module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module#name DataTfeRegistryModule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module#organization DataTfeRegistryModule#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The namespace of the no-code module. Uses organization name if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module#namespace DataTfeRegistryModule#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module#permissions DataTfeRegistryModule#permissions}
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Name of the registry. Valid options: "public", "private". Defaults to "private".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module#registry_name DataTfeRegistryModule#registry_name}
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// test_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module#test_config DataTfeRegistryModule#test_config}
	TestConfig interface{} `field:"optional" json:"testConfig" yaml:"testConfig"`
	// vcs_repo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module#vcs_repo DataTfeRegistryModule#vcs_repo}
	VcsRepo interface{} `field:"optional" json:"vcsRepo" yaml:"vcsRepo"`
	// version_statuses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module#version_statuses DataTfeRegistryModule#version_statuses}
	VersionStatuses interface{} `field:"optional" json:"versionStatuses" yaml:"versionStatuses"`
}

