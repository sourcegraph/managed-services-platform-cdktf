package zerotrustaccessapplication


type ZeroTrustAccessApplicationOauthConfigurationDynamicClientRegistration struct {
	// Allows any client with redirect URIs on localhost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_application#allow_any_on_localhost ZeroTrustAccessApplication#allow_any_on_localhost}
	AllowAnyOnLocalhost interface{} `field:"optional" json:"allowAnyOnLocalhost" yaml:"allowAnyOnLocalhost"`
	// Allows any client with redirect URIs on 127.0.0.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_application#allow_any_on_loopback ZeroTrustAccessApplication#allow_any_on_loopback}
	AllowAnyOnLoopback interface{} `field:"optional" json:"allowAnyOnLoopback" yaml:"allowAnyOnLoopback"`
	// The URIs that are allowed as redirect URIs for dynamically registered clients.
	//
	// Must use the `https` protocol. Paths may end in `/*` to match all sub-paths.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_application#allowed_uris ZeroTrustAccessApplication#allowed_uris}
	AllowedUris *[]*string `field:"optional" json:"allowedUris" yaml:"allowedUris"`
	// Whether dynamic client registration is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_application#enabled ZeroTrustAccessApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

