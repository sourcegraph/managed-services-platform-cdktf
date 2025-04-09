package googlebigtableinstanceiambinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigtableInstanceIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_bigtable_instance_iam_binding#instance GoogleBigtableInstanceIamBinding#instance}.
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_bigtable_instance_iam_binding#members GoogleBigtableInstanceIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_bigtable_instance_iam_binding#role GoogleBigtableInstanceIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_bigtable_instance_iam_binding#condition GoogleBigtableInstanceIamBinding#condition}
	Condition *GoogleBigtableInstanceIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_bigtable_instance_iam_binding#id GoogleBigtableInstanceIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_bigtable_instance_iam_binding#project GoogleBigtableInstanceIamBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

