package zerotrustdlpsettings


type ZeroTrustDlpSettingsPayloadLogging struct {
	// Masking level for payload logs.
	//
	// - `full`: The entire payload is masked.
	// - `partial`: Only partial payload content is masked.
	// - `clear`: No masking is applied to the payload content.
	// - `default`: DLP uses its default masking behavior.
	// Available values: "full", "partial", "clear", "default".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_settings#masking_level ZeroTrustDlpSettings#masking_level}
	MaskingLevel *string `field:"optional" json:"maskingLevel" yaml:"maskingLevel"`
	// Base64-encoded public key for encrypting payload logs.
	//
	// - Set to a non-empty base64 string to enable payload logging with the given key.
	// - Set to an empty string to disable payload logging.
	// - Omit or set to null to leave unchanged (PATCH) or reset to disabled (PUT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_settings#public_key ZeroTrustDlpSettings#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
}

