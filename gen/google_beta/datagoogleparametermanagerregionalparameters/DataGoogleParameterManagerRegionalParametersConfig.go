package datagoogleparametermanagerregionalparameters

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleParameterManagerRegionalParametersConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_parameter_manager_regional_parameters#location DataGoogleParameterManagerRegionalParameters#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Filter string, adhering to the rules in List-operation filtering.
	//
	// List only parameters matching the filter.
	// If filter is empty, all regional parameters are listed from specific location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_parameter_manager_regional_parameters#filter DataGoogleParameterManagerRegionalParameters#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_parameter_manager_regional_parameters#id DataGoogleParameterManagerRegionalParameters#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_parameter_manager_regional_parameters#project DataGoogleParameterManagerRegionalParameters#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

