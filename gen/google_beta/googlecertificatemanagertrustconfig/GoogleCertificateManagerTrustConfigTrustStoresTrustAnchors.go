package googlecertificatemanagertrustconfig


type GoogleCertificateManagerTrustConfigTrustStoresTrustAnchors struct {
	// PEM root certificate of the PKI used for validation. Each certificate provided in PEM format may occupy up to 5kB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_certificate_manager_trust_config#pem_certificate GoogleCertificateManagerTrustConfig#pem_certificate}
	PemCertificate *string `field:"optional" json:"pemCertificate" yaml:"pemCertificate"`
}

