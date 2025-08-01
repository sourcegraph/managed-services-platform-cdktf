package zerotrustdlppredefinedprofile


type ZeroTrustDlpPredefinedProfileContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dlp_predefined_profile#files ZeroTrustDlpPredefinedProfile#files}
	Files interface{} `field:"required" json:"files" yaml:"files"`
}

