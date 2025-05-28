package googlegeminiloggingsetting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGeminiLoggingSettingConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gemini_logging_setting#location GoogleGeminiLoggingSetting#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Id of the Logging Setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gemini_logging_setting#logging_setting_id GoogleGeminiLoggingSetting#logging_setting_id}
	LoggingSettingId *string `field:"required" json:"loggingSettingId" yaml:"loggingSettingId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gemini_logging_setting#id GoogleGeminiLoggingSetting#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gemini_logging_setting#labels GoogleGeminiLoggingSetting#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Whether to log metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gemini_logging_setting#log_metadata GoogleGeminiLoggingSetting#log_metadata}
	LogMetadata interface{} `field:"optional" json:"logMetadata" yaml:"logMetadata"`
	// Whether to log prompts and responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gemini_logging_setting#log_prompts_and_responses GoogleGeminiLoggingSetting#log_prompts_and_responses}
	LogPromptsAndResponses interface{} `field:"optional" json:"logPromptsAndResponses" yaml:"logPromptsAndResponses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gemini_logging_setting#project GoogleGeminiLoggingSetting#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gemini_logging_setting#timeouts GoogleGeminiLoggingSetting#timeouts}
	Timeouts *GoogleGeminiLoggingSettingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

