package tokenvalidationrules


type TokenValidationRulesSelectorInclude struct {
	// Included hostnames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/token_validation_rules#host TokenValidationRules#host}
	Host *[]*string `field:"optional" json:"host" yaml:"host"`
}

