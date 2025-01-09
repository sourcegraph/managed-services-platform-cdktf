package testvariable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TestVariableConfig struct {
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
	// Whether this is a Terraform or environment variable. Valid values are "terraform" or "env".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/test_variable#category TestVariable#category}
	Category *string `field:"required" json:"category" yaml:"category"`
	// Name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/test_variable#key TestVariable#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/test_variable#module_name TestVariable#module_name}.
	ModuleName *string `field:"required" json:"moduleName" yaml:"moduleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/test_variable#module_provider TestVariable#module_provider}.
	ModuleProvider *string `field:"required" json:"moduleProvider" yaml:"moduleProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/test_variable#organization TestVariable#organization}.
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/test_variable#description TestVariable#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/test_variable#hcl TestVariable#hcl}.
	Hcl interface{} `field:"optional" json:"hcl" yaml:"hcl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/test_variable#sensitive TestVariable#sensitive}.
	Sensitive interface{} `field:"optional" json:"sensitive" yaml:"sensitive"`
	// Value of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/test_variable#value TestVariable#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

