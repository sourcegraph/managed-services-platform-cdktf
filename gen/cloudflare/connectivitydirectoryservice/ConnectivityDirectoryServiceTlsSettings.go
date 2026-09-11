package connectivitydirectoryservice


type ConnectivityDirectoryServiceTlsSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/connectivity_directory_service#cert_verification_mode ConnectivityDirectoryService#cert_verification_mode}
	CertVerificationMode *string `field:"required" json:"certVerificationMode" yaml:"certVerificationMode"`
}

