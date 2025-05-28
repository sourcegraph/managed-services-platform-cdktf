package googlenetworksecurityclienttlspolicy


type GoogleNetworkSecurityClientTlsPolicyServerValidationCa struct {
	// certificate_provider_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_network_security_client_tls_policy#certificate_provider_instance GoogleNetworkSecurityClientTlsPolicy#certificate_provider_instance}
	CertificateProviderInstance *GoogleNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance `field:"optional" json:"certificateProviderInstance" yaml:"certificateProviderInstance"`
	// grpc_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_network_security_client_tls_policy#grpc_endpoint GoogleNetworkSecurityClientTlsPolicy#grpc_endpoint}
	GrpcEndpoint *GoogleNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint `field:"optional" json:"grpcEndpoint" yaml:"grpcEndpoint"`
}

