package tokenvalidationrules


type TokenValidationRulesSelectorExclude struct {
	// Excluded operation IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_rules#operation_ids TokenValidationRules#operation_ids}
	OperationIds *[]*string `field:"optional" json:"operationIds" yaml:"operationIds"`
}

