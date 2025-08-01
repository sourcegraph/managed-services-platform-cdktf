package googleapigeeenvironmentaddonsconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApigeeEnvironmentAddonsConfigConfig struct {
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
	// The Apigee environment group associated with the Apigee environment, in the format 'organizations/{{org_name}}/environments/{{env_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_apigee_environment_addons_config#env_id GoogleApigeeEnvironmentAddonsConfig#env_id}
	EnvId *string `field:"required" json:"envId" yaml:"envId"`
	// Flag to enable/disable Analytics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_apigee_environment_addons_config#analytics_enabled GoogleApigeeEnvironmentAddonsConfig#analytics_enabled}
	AnalyticsEnabled interface{} `field:"optional" json:"analyticsEnabled" yaml:"analyticsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_apigee_environment_addons_config#id GoogleApigeeEnvironmentAddonsConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_apigee_environment_addons_config#timeouts GoogleApigeeEnvironmentAddonsConfig#timeouts}
	Timeouts *GoogleApigeeEnvironmentAddonsConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

