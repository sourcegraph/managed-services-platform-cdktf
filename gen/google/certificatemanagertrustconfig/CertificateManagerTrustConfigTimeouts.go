package certificatemanagertrustconfig


type CertificateManagerTrustConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/certificate_manager_trust_config#create CertificateManagerTrustConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/certificate_manager_trust_config#delete CertificateManagerTrustConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/certificate_manager_trust_config#update CertificateManagerTrustConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

