package accesspolicy


type AccessPolicyConnectionRules struct {
	// ssh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_policy#ssh AccessPolicy#ssh}
	Ssh *AccessPolicyConnectionRulesSsh `field:"required" json:"ssh" yaml:"ssh"`
}

