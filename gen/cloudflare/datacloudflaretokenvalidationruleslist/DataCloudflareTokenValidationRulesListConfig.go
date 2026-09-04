package datacloudflaretokenvalidationruleslist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareTokenValidationRulesListConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules_list#zone_id DataCloudflareTokenValidationRulesList#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Action to take on requests that match operations included in `selector` and fail `expression`. Available values: "log", "block".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules_list#action DataCloudflareTokenValidationRulesList#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Toggle rule on or off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules_list#enabled DataCloudflareTokenValidationRulesList#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Select rules with this host in `include`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules_list#host DataCloudflareTokenValidationRulesList#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Select rules with this host in `include`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules_list#hostname DataCloudflareTokenValidationRulesList#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Select rules with these IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules_list#id DataCloudflareTokenValidationRulesList#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules_list#max_items DataCloudflareTokenValidationRulesList#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Select rules with these IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules_list#rule_id DataCloudflareTokenValidationRulesList#rule_id}
	RuleId *string `field:"optional" json:"ruleId" yaml:"ruleId"`
	// Select rules using any of these token configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/token_validation_rules_list#token_configuration DataCloudflareTokenValidationRulesList#token_configuration}
	TokenConfiguration *[]*string `field:"optional" json:"tokenConfiguration" yaml:"tokenConfiguration"`
}

