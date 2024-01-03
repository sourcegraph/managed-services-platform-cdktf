package usergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UsergroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pablovarela/slack/1.2.2/docs/resources/usergroup#name Usergroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pablovarela/slack/1.2.2/docs/resources/usergroup#channels Usergroup#channels}.
	Channels *[]*string `field:"optional" json:"channels" yaml:"channels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pablovarela/slack/1.2.2/docs/resources/usergroup#description Usergroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pablovarela/slack/1.2.2/docs/resources/usergroup#handle Usergroup#handle}.
	Handle *string `field:"optional" json:"handle" yaml:"handle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pablovarela/slack/1.2.2/docs/resources/usergroup#id Usergroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pablovarela/slack/1.2.2/docs/resources/usergroup#users Usergroup#users}.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}

