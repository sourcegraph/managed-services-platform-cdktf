package datatferegistryproviders

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataTfeRegistryProvidersConfig struct {
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
	// Name of the organization. If omitted, organization must be defined in the provider config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_providers#organization DataTfeRegistryProviders#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Whether to list only public or private providers. Must be either `public` or `private`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_providers#registry_name DataTfeRegistryProviders#registry_name}
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// A query string to do a fuzzy search on provider name and namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_providers#search DataTfeRegistryProviders#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

