package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesExcludeEmailDomain struct {
	// The email domain to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#domain ZeroTrustAccessApplication#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

