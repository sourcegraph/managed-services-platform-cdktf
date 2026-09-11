package zerotrustaccesspolicy


type ZeroTrustAccessPolicyMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#allowed_authenticators ZeroTrustAccessPolicy#allowed_authenticators}
	AllowedAuthenticators *[]*string `field:"optional" json:"allowedAuthenticators" yaml:"allowedAuthenticators"`
	// Indicates whether to disable MFA for this resource. This option is available at the application and policy level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#mfa_disabled ZeroTrustAccessPolicy#mfa_disabled}
	MfaDisabled interface{} `field:"optional" json:"mfaDisabled" yaml:"mfaDisabled"`
	// Defines the duration of an MFA session.
	//
	// Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_policy#session_duration ZeroTrustAccessPolicy#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
}

