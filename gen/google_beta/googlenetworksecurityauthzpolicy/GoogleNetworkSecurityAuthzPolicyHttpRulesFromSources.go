package googlenetworksecurityauthzpolicy

type GoogleNetworkSecurityAuthzPolicyHttpRulesFromSources struct {
	// principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_security_authz_policy#principals GoogleNetworkSecurityAuthzPolicy#principals}
	Principals any `field:"optional" json:"principals" yaml:"principals"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_security_authz_policy#resources GoogleNetworkSecurityAuthzPolicy#resources}
	Resources any `field:"optional" json:"resources" yaml:"resources"`
}
