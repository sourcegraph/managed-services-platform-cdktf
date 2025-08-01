package zerotrustaccessgroup


type ZeroTrustAccessGroupExcludeGeo struct {
	// The country code that should be matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#country_code ZeroTrustAccessGroup#country_code}
	CountryCode *string `field:"required" json:"countryCode" yaml:"countryCode"`
}

