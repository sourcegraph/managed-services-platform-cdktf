package zerotrustaccessapplication


type ZeroTrustAccessApplicationFooterLinks struct {
	// The hypertext in the footer link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// the hyperlink in the footer link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#url ZeroTrustAccessApplication#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

