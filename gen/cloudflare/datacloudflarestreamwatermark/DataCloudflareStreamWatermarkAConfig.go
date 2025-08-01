package datacloudflarestreamwatermark

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareStreamWatermarkAConfig struct {
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
	// The account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/stream_watermark#account_id DataCloudflareStreamWatermarkA#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The unique identifier for a watermark profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/stream_watermark#identifier DataCloudflareStreamWatermarkA#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

