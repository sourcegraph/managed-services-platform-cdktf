package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyCustomProviderCloudIap struct {
	// Enable Cloud IAP at the AuthzPolicy level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_authz_policy#enabled NetworkSecurityAuthzPolicy#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

