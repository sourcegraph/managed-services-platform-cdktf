package dlpprofile


type DlpProfileContextAwarenessSkip struct {
	// Return all matches, regardless of context analysis result, if the data is a file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/dlp_profile#files DlpProfile#files}
	Files interface{} `field:"required" json:"files" yaml:"files"`
}

