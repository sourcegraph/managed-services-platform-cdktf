package customhostname


type CustomHostnameSslCustomCertBundle struct {
	// If a custom uploaded certificate is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/custom_hostname#custom_certificate CustomHostname#custom_certificate}
	CustomCertificate *string `field:"required" json:"customCertificate" yaml:"customCertificate"`
	// The key for a custom uploaded certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/custom_hostname#custom_key CustomHostname#custom_key}
	CustomKey *string `field:"required" json:"customKey" yaml:"customKey"`
}

