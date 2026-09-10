package zerotrustaccesspolicy


type ZeroTrustAccessPolicyIncludeEmailDomain struct {
	// The email domain to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#domain ZeroTrustAccessPolicy#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

