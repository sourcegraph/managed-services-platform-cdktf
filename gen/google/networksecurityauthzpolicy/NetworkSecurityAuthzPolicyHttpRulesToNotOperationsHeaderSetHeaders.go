package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHeaderSetHeaders struct {
	// Specifies the name of the header in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_authz_policy#name NetworkSecurityAuthzPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_authz_policy#value NetworkSecurityAuthzPolicy#value}
	Value *NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHeaderSetHeadersValue `field:"optional" json:"value" yaml:"value"`
}

