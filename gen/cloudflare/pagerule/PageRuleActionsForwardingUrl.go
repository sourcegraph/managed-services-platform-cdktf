package pagerule


type PageRuleActionsForwardingUrl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/page_rule#status_code PageRule#status_code}.
	StatusCode *float64 `field:"required" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/page_rule#url PageRule#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

