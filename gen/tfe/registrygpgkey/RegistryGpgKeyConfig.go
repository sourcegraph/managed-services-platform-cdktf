package registrygpgkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RegistryGpgKeyConfig struct {
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
	// ASCII-armored representation of the GPG key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/registry_gpg_key#ascii_armor RegistryGpgKey#ascii_armor}
	AsciiArmor *string `field:"required" json:"asciiArmor" yaml:"asciiArmor"`
	// Name of the organization. If omitted, organization must be defined in the provider config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/registry_gpg_key#organization RegistryGpgKey#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
}

