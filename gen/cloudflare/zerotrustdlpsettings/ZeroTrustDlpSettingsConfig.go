package zerotrustdlpsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDlpSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_settings#account_id ZeroTrustDlpSettings#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Whether AI context analysis is enabled at the account level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_settings#ai_context_analysis ZeroTrustDlpSettings#ai_context_analysis}
	AiContextAnalysis interface{} `field:"optional" json:"aiContextAnalysis" yaml:"aiContextAnalysis"`
	// Whether OCR is enabled at the account level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_settings#ocr ZeroTrustDlpSettings#ocr}
	Ocr interface{} `field:"optional" json:"ocr" yaml:"ocr"`
	// Request model for payload log settings within the DLP settings endpoint.
	//
	// Unlike the legacy endpoint, null and missing are treated identically here
	// (both mean "not provided" for PATCH, "reset to default" for PUT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_settings#payload_logging ZeroTrustDlpSettings#payload_logging}
	PayloadLogging *ZeroTrustDlpSettingsPayloadLogging `field:"optional" json:"payloadLogging" yaml:"payloadLogging"`
}

