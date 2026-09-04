package zerotrustgatewaypolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewayPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#account_id ZeroTrustGatewayPolicy#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Specify the action to perform when the associated traffic, identity, and device posture expressions either absent or evaluate to `true`.
	//
	// Available values: "on", "off", "allow", "block", "scan", "noscan", "safesearch", "ytrestricted", "isolate", "noisolate", "override", "l4_override", "egress", "resolve", "quarantine", "redirect".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#action ZeroTrustGatewayPolicy#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Specify the rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#name ZeroTrustGatewayPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify the rule description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#description ZeroTrustGatewayPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specify the wirefilter expression used for device posture check.
	//
	// The API automatically formats and sanitizes expressions before storing them. To prevent Terraform state drift, use the formatted expression returned in the API response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#device_posture ZeroTrustGatewayPolicy#device_posture}
	DevicePosture *string `field:"optional" json:"devicePosture" yaml:"devicePosture"`
	// Specify whether the rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#enabled ZeroTrustGatewayPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Defines the expiration time stamp and default duration of a DNS policy.
	//
	// Takes precedence over the policy's `schedule` configuration, if any. This  does not apply to HTTP or network policies. Settable only for `dns` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#expiration ZeroTrustGatewayPolicy#expiration}
	Expiration *ZeroTrustGatewayPolicyExpiration `field:"optional" json:"expiration" yaml:"expiration"`
	// Specify the protocol or layer to evaluate the traffic, identity, and device posture expressions.
	//
	// Can only contain a single value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#filters ZeroTrustGatewayPolicy#filters}
	Filters *[]*string `field:"optional" json:"filters" yaml:"filters"`
	// Specify the wirefilter expression used for identity matching.
	//
	// The API automatically formats and sanitizes expressions before storing them. To prevent Terraform state drift, use the formatted expression returned in the API response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#identity ZeroTrustGatewayPolicy#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Set the order of your rules.
	//
	// Lower values indicate higher precedence. At each processing phase, evaluate applicable rules in ascending order of this value. Refer to [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform) to manage precedence via Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#precedence ZeroTrustGatewayPolicy#precedence}
	Precedence *float64 `field:"optional" json:"precedence" yaml:"precedence"`
	// Defines settings for this rule.
	//
	// Settings apply only to specific rule types and must use compatible selectors. If Terraform detects drift, confirm the setting supports your rule type and check whether the API modifies the value. Use API-returned values in your configuration to prevent drift.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#rule_settings ZeroTrustGatewayPolicy#rule_settings}
	RuleSettings *ZeroTrustGatewayPolicyRuleSettings `field:"optional" json:"ruleSettings" yaml:"ruleSettings"`
	// Defines the schedule for activating DNS policies. Settable only for `dns` and `dns_resolver` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#schedule ZeroTrustGatewayPolicy#schedule}
	Schedule *ZeroTrustGatewayPolicySchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Specify the wirefilter expression used for traffic matching.
	//
	// The API automatically formats and sanitizes expressions before storing them. To prevent Terraform state drift, use the formatted expression returned in the API response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#traffic ZeroTrustGatewayPolicy#traffic}
	Traffic *string `field:"optional" json:"traffic" yaml:"traffic"`
}

