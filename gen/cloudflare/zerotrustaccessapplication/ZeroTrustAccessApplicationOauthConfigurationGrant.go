package zerotrustaccessapplication


type ZeroTrustAccessApplicationOauthConfigurationGrant struct {
	// The lifetime of the access token.
	//
	// Must be in the format `300ms` or `2h45m`. Valid time units are ns, us (or µs), ms, s, m, h.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_application#access_token_lifetime ZeroTrustAccessApplication#access_token_lifetime}
	AccessTokenLifetime *string `field:"optional" json:"accessTokenLifetime" yaml:"accessTokenLifetime"`
	// The duration of the OAuth session.
	//
	// Must be in the format `300ms` or `2h45m`. Valid time units are ns, us (or µs), ms, s, m, h.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_application#session_duration ZeroTrustAccessApplication#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
}

