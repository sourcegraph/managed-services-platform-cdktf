package project


type ProjectClientSecurity struct {
	// A list of allowed domains. Examples: https://example.com, *, *.example.com, *:80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#allowed_domains Project#allowed_domains}
	AllowedDomains *[]*string `field:"optional" json:"allowedDomains" yaml:"allowedDomains"`
	// Enable JavaScript source fetching. Allow Sentry to scrape missing JavaScript source context when possible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#scrape_javascript Project#scrape_javascript}
	ScrapeJavascript interface{} `field:"optional" json:"scrapeJavascript" yaml:"scrapeJavascript"`
	// Security Token. Outbound requests matching Allowed Domains will have the header "{security_token_header}: {security_token}" appended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#security_token Project#security_token}
	SecurityToken *string `field:"optional" json:"securityToken" yaml:"securityToken"`
	// Security Token Header. Outbound requests matching Allowed Domains will have the header "{security_token_header}: {security_token}" appended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#security_token_header Project#security_token_header}
	SecurityTokenHeader *string `field:"optional" json:"securityTokenHeader" yaml:"securityTokenHeader"`
	// Verify TLS/SSL. Outbound requests will verify TLS (sometimes known as SSL) connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project#verify_tls_ssl Project#verify_tls_ssl}
	VerifyTlsSsl interface{} `field:"optional" json:"verifyTlsSsl" yaml:"verifyTlsSsl"`
}

