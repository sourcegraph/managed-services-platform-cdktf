package zerotrustaccessgroup


type ZeroTrustAccessGroupRequireCommonName struct {
	// The common name to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#common_name ZeroTrustAccessGroup#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
}

