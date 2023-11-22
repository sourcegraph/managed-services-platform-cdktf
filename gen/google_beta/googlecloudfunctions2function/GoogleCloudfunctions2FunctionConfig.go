package googlecloudfunctions2function

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudfunctions2FunctionConfig struct {
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
	// The location of this cloud function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudfunctions2_function#location GoogleCloudfunctions2Function#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A user-defined name of the function. Function names must be unique globally and match pattern 'projects/*\/locations/*\/functions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudfunctions2_function#name GoogleCloudfunctions2Function#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// build_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudfunctions2_function#build_config GoogleCloudfunctions2Function#build_config}
	BuildConfig *GoogleCloudfunctions2FunctionBuildConfig `field:"optional" json:"buildConfig" yaml:"buildConfig"`
	// User-provided description of a function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudfunctions2_function#description GoogleCloudfunctions2Function#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// event_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudfunctions2_function#event_trigger GoogleCloudfunctions2Function#event_trigger}
	EventTrigger *GoogleCloudfunctions2FunctionEventTrigger `field:"optional" json:"eventTrigger" yaml:"eventTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudfunctions2_function#id GoogleCloudfunctions2Function#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
	//
	// It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudfunctions2_function#kms_key_name GoogleCloudfunctions2Function#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// A set of key/value label pairs associated with this Cloud Function.
	//
	// *Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudfunctions2_function#labels GoogleCloudfunctions2Function#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudfunctions2_function#project GoogleCloudfunctions2Function#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// service_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudfunctions2_function#service_config GoogleCloudfunctions2Function#service_config}
	ServiceConfig *GoogleCloudfunctions2FunctionServiceConfig `field:"optional" json:"serviceConfig" yaml:"serviceConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_cloudfunctions2_function#timeouts GoogleCloudfunctions2Function#timeouts}
	Timeouts *GoogleCloudfunctions2FunctionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

