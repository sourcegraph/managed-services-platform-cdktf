package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesRequireCommonName struct {
	// The common name to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_application#common_name ZeroTrustAccessApplication#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
}

