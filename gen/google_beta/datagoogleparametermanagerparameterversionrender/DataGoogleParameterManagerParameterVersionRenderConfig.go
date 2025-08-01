package datagoogleparametermanagerparameterversionrender

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleParameterManagerParameterVersionRenderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_parameter_manager_parameter_version_render#parameter DataGoogleParameterManagerParameterVersionRender#parameter}.
	Parameter *string `field:"required" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_parameter_manager_parameter_version_render#parameter_version_id DataGoogleParameterManagerParameterVersionRender#parameter_version_id}.
	ParameterVersionId *string `field:"required" json:"parameterVersionId" yaml:"parameterVersionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_parameter_manager_parameter_version_render#id DataGoogleParameterManagerParameterVersionRender#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/data-sources/google_parameter_manager_parameter_version_render#project DataGoogleParameterManagerParameterVersionRender#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

