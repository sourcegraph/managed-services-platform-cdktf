package zerotrustorganization


type ZeroTrustOrganizationMfaConfig struct {
	// Lists the MFA methods that users can authenticate with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_organization#allowed_authenticators ZeroTrustOrganization#allowed_authenticators}
	AllowedAuthenticators *[]*string `field:"optional" json:"allowedAuthenticators" yaml:"allowedAuthenticators"`
	// Allows a user to skip MFA via Authentication Method Reference (AMR) matching when the AMR claim provided by the IdP the user used to authenticate contains "mfa".
	//
	// Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30 days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_organization#amr_matching_session_duration ZeroTrustOrganization#amr_matching_session_duration}
	AmrMatchingSessionDuration *string `field:"optional" json:"amrMatchingSessionDuration" yaml:"amrMatchingSessionDuration"`
	// Specifies a Cloudflare List of required FIDO2 authenticator device AAGUIDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_organization#required_aaguids ZeroTrustOrganization#required_aaguids}
	RequiredAaguids *string `field:"optional" json:"requiredAaguids" yaml:"requiredAaguids"`
	// Defines the duration of an MFA session.
	//
	// Must be in minutes (m) or hours (h). Minimum: 0m. Maximum: 720h (30 days). Examples:`5m` or `24h`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_organization#session_duration ZeroTrustOrganization#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
}

