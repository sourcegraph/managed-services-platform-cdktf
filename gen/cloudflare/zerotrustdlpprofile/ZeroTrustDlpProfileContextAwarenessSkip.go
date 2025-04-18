package zerotrustdlpprofile


type ZeroTrustDlpProfileContextAwarenessSkip struct {
	// Return all matches, regardless of context analysis result, if the data is a file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_dlp_profile#files ZeroTrustDlpProfile#files}
	Files interface{} `field:"required" json:"files" yaml:"files"`
}

