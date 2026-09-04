package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesExcludeGeo struct {
	// The country code that should be matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_access_application#country_code ZeroTrustAccessApplication#country_code}
	CountryCode *string `field:"required" json:"countryCode" yaml:"countryCode"`
}

