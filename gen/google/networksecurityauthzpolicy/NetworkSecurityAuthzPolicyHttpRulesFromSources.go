package networksecurityauthzpolicy

type NetworkSecurityAuthzPolicyHttpRulesFromSources struct {
	// principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_authz_policy#principals NetworkSecurityAuthzPolicy#principals}
	Principals any `field:"optional" json:"principals" yaml:"principals"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_authz_policy#resources NetworkSecurityAuthzPolicy#resources}
	Resources any `field:"optional" json:"resources" yaml:"resources"`
}
