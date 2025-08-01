package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesFromNotSources struct {
	// principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_authz_policy#principals NetworkSecurityAuthzPolicy#principals}
	Principals interface{} `field:"optional" json:"principals" yaml:"principals"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_authz_policy#resources NetworkSecurityAuthzPolicy#resources}
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
}

