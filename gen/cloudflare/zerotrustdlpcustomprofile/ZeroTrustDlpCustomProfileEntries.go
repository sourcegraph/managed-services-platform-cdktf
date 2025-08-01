package zerotrustdlpcustomprofile


type ZeroTrustDlpCustomProfileEntries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dlp_custom_profile#enabled ZeroTrustDlpCustomProfile#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dlp_custom_profile#name ZeroTrustDlpCustomProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dlp_custom_profile#pattern ZeroTrustDlpCustomProfile#pattern}.
	Pattern *ZeroTrustDlpCustomProfileEntriesPattern `field:"optional" json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dlp_custom_profile#words ZeroTrustDlpCustomProfile#words}.
	Words *[]*string `field:"optional" json:"words" yaml:"words"`
}

