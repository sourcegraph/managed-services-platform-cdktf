package googlegeminireleasechannelsetting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGeminiReleaseChannelSettingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_gemini_release_channel_setting#location GoogleGeminiReleaseChannelSetting#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Id of the Release Channel Setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_gemini_release_channel_setting#release_channel_setting_id GoogleGeminiReleaseChannelSetting#release_channel_setting_id}
	ReleaseChannelSettingId *string `field:"required" json:"releaseChannelSettingId" yaml:"releaseChannelSettingId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_gemini_release_channel_setting#id GoogleGeminiReleaseChannelSetting#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_gemini_release_channel_setting#labels GoogleGeminiReleaseChannelSetting#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_gemini_release_channel_setting#project GoogleGeminiReleaseChannelSetting#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Release channel to be used. Possible values: STABLE EXPERIMENTAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_gemini_release_channel_setting#release_channel GoogleGeminiReleaseChannelSetting#release_channel}
	ReleaseChannel *string `field:"optional" json:"releaseChannel" yaml:"releaseChannel"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_gemini_release_channel_setting#timeouts GoogleGeminiReleaseChannelSetting#timeouts}
	Timeouts *GoogleGeminiReleaseChannelSettingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

