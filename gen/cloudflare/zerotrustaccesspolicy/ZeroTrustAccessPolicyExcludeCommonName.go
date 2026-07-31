package zerotrustaccesspolicy


type ZeroTrustAccessPolicyExcludeCommonName struct {
	// The common name to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_policy#common_name ZeroTrustAccessPolicy#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
}

