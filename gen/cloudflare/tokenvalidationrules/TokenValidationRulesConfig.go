package tokenvalidationrules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TokenValidationRulesConfig struct {
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
	// Action to take on requests that match operations included in `selector` and fail `expression`. Available values: "log", "block".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_rules#action TokenValidationRules#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// A human-readable description that gives more details than `title`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_rules#description TokenValidationRules#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Toggle rule on or off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_rules#enabled TokenValidationRules#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Rule expression. Requests that fail to match this expression will be subject to `action`.
	//
	// For details on expressions, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_rules#expression TokenValidationRules#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Select operations covered by this rule.
	//
	// For details on selectors, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_rules#selector TokenValidationRules#selector}
	Selector *TokenValidationRulesSelector `field:"required" json:"selector" yaml:"selector"`
	// A human-readable name for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_rules#title TokenValidationRules#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_rules#zone_id TokenValidationRules#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Update rule order among zone rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_rules#position TokenValidationRules#position}
	Position *TokenValidationRulesPosition `field:"optional" json:"position" yaml:"position"`
}

