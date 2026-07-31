package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyExpiration struct {
	// Show the timestamp when the policy expires and stops applying.
	//
	// The value must follow RFC 3339 and include a UTC offset.  The system accepts non-zero offsets but converts them to the equivalent UTC+00:00  value and returns timestamps with a trailing Z. Expiration policies ignore client  timezones and expire globally at the specified expires_at time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#expires_at ZeroTrustGatewayPolicy#expires_at}
	ExpiresAt *string `field:"required" json:"expiresAt" yaml:"expiresAt"`
	// Defines the default duration a policy active in minutes.
	//
	// Must set in order to use the `reset_expiration` endpoint on this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#duration ZeroTrustGatewayPolicy#duration}
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
}

