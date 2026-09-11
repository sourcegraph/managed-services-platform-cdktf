package zerotrustgatewaylogging

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewayLoggingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_logging#account_id ZeroTrustGatewayLogging#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Indicate whether to redact personally identifiable information from activity logging (PII fields include source IP, user email, user ID, device ID, URL, referrer, and user agent).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_logging#redact_pii ZeroTrustGatewayLogging#redact_pii}
	RedactPii interface{} `field:"optional" json:"redactPii" yaml:"redactPii"`
	// Configure logging settings for each rule type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_logging#settings_by_rule_type ZeroTrustGatewayLogging#settings_by_rule_type}
	SettingsByRuleType *ZeroTrustGatewayLoggingSettingsByRuleType `field:"optional" json:"settingsByRuleType" yaml:"settingsByRuleType"`
}

