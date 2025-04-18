package zerotrustaccessapplication


type ZeroTrustAccessApplicationFooterLinks struct {
	// The name of the footer link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The URL of the footer link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_application#url ZeroTrustAccessApplication#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

