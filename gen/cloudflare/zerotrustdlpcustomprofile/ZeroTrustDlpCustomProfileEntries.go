package zerotrustdlpcustomprofile


type ZeroTrustDlpCustomProfileEntries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_custom_profile#enabled ZeroTrustDlpCustomProfile#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_custom_profile#name ZeroTrustDlpCustomProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_custom_profile#pattern ZeroTrustDlpCustomProfile#pattern}.
	Pattern *ZeroTrustDlpCustomProfileEntriesPattern `field:"required" json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_custom_profile#description ZeroTrustDlpCustomProfile#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_custom_profile#entry_id ZeroTrustDlpCustomProfile#entry_id}.
	EntryId *string `field:"optional" json:"entryId" yaml:"entryId"`
}

