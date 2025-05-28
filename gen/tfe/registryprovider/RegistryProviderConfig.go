package registryprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RegistryProviderConfig struct {
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
	// Name of the provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/registry_provider#name RegistryProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the provider. For private providers this is the same as the oraganization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/registry_provider#namespace RegistryProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Name of the organization. If omitted, organization must be defined in the provider config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/registry_provider#organization RegistryProvider#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Whether this is a publicly maintained provider or private. Must be either `public` or `private`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/registry_provider#registry_name RegistryProvider#registry_name}
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
}

