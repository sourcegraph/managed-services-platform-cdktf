package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyCustomProviderAuthzExtension struct {
	// A list of references to authorization extensions that will be invoked for requests matching this policy.
	//
	// Limited to 1 custom provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_authz_policy#resources NetworkSecurityAuthzPolicy#resources}
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

