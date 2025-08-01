package certificatemanagertrustconfig


type CertificateManagerTrustConfigAllowlistedCertificates struct {
	// PEM certificate that is allowlisted.
	//
	// The certificate can be up to 5k bytes, and must be a parseable X.509 certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/certificate_manager_trust_config#pem_certificate CertificateManagerTrustConfig#pem_certificate}
	PemCertificate *string `field:"required" json:"pemCertificate" yaml:"pemCertificate"`
}

