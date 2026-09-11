package tokenvalidationrules


type TokenValidationRulesSelector struct {
	// Ignore operations that were otherwise included by `include`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_rules#exclude TokenValidationRules#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// Select all matching operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/token_validation_rules#include TokenValidationRules#include}
	Include interface{} `field:"optional" json:"include" yaml:"include"`
}

