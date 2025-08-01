package firewallrule


type FirewallRuleFilter struct {
	// An informative summary of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/firewall_rule#description FirewallRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/firewall_rule#expression FirewallRule#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// When true, indicates that the filter is currently paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/firewall_rule#paused FirewallRule#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// A short reference tag. Allows you to select related filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/firewall_rule#ref FirewallRule#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

