package sshkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SshKeyConfig struct {
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
	// The name of the SSH key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/ssh_key#name SshKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The text of the SSH private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/ssh_key#key SshKey#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The text of the SSH private key, guaranteed not to be written to state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/ssh_key#key_wo SshKey#key_wo}
	KeyWo *string `field:"optional" json:"keyWo" yaml:"keyWo"`
	// The name of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/ssh_key#organization SshKey#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
}

