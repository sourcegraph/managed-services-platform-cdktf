package zerotrustdlpprofile


type ZeroTrustDlpProfileContextAwareness struct {
	// Scan the context of predefined entries to only return matches surrounded by keywords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_dlp_profile#enabled ZeroTrustDlpProfile#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// skip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_dlp_profile#skip ZeroTrustDlpProfile#skip}
	Skip *ZeroTrustDlpProfileContextAwarenessSkip `field:"required" json:"skip" yaml:"skip"`
}

