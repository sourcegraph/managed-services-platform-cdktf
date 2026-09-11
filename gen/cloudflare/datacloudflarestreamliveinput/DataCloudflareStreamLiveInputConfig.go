package datacloudflarestreamliveinput

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareStreamLiveInputConfig struct {
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
	// A unique identifier for a live input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/stream_live_input#live_input_identifier DataCloudflareStreamLiveInput#live_input_identifier}
	LiveInputIdentifier *string `field:"required" json:"liveInputIdentifier" yaml:"liveInputIdentifier"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/stream_live_input#account_id DataCloudflareStreamLiveInput#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}

