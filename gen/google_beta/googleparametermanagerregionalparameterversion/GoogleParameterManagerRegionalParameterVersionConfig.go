package googleparametermanagerregionalparameterversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleParameterManagerRegionalParameterVersionConfig struct {
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
	// Parameter Manager Regional Parameter resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_parameter_manager_regional_parameter_version#parameter GoogleParameterManagerRegionalParameterVersion#parameter}
	Parameter *string `field:"required" json:"parameter" yaml:"parameter"`
	// The Regional Parameter data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_parameter_manager_regional_parameter_version#parameter_data GoogleParameterManagerRegionalParameterVersion#parameter_data}
	ParameterData *string `field:"required" json:"parameterData" yaml:"parameterData"`
	// Version ID of the Regional Parameter Version Resource. This must be unique within the Regional Parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_parameter_manager_regional_parameter_version#parameter_version_id GoogleParameterManagerRegionalParameterVersion#parameter_version_id}
	ParameterVersionId *string `field:"required" json:"parameterVersionId" yaml:"parameterVersionId"`
	// The current state of Regional Parameter Version. This field is only applicable for updating Regional Parameter Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_parameter_manager_regional_parameter_version#disabled GoogleParameterManagerRegionalParameterVersion#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_parameter_manager_regional_parameter_version#id GoogleParameterManagerRegionalParameterVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_parameter_manager_regional_parameter_version#timeouts GoogleParameterManagerRegionalParameterVersion#timeouts}
	Timeouts *GoogleParameterManagerRegionalParameterVersionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

