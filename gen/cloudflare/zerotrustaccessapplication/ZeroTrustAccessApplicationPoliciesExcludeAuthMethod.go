package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesExcludeAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#auth_method ZeroTrustAccessApplication#auth_method}
	AuthMethod *string `field:"required" json:"authMethod" yaml:"authMethod"`
}

