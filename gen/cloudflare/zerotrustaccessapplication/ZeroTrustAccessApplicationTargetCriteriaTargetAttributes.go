package zerotrustaccessapplication


type ZeroTrustAccessApplicationTargetCriteriaTargetAttributes struct {
	// The key of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_application#values ZeroTrustAccessApplication#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

