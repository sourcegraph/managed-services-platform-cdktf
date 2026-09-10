package snippetrules


type SnippetRulesRules struct {
	// Define the expression that determines which traffic matches the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/snippet_rules#expression SnippetRules#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Identify the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/snippet_rules#snippet_name SnippetRules#snippet_name}
	SnippetName *string `field:"required" json:"snippetName" yaml:"snippetName"`
	// Provide an informative description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/snippet_rules#description SnippetRules#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicate whether to execute the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/snippet_rules#enabled SnippetRules#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

