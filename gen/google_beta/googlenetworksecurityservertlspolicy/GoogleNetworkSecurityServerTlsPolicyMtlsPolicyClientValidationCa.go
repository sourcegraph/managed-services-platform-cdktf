package googlenetworksecurityservertlspolicy


type GoogleNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCa struct {
	// certificate_provider_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_network_security_server_tls_policy#certificate_provider_instance GoogleNetworkSecurityServerTlsPolicy#certificate_provider_instance}
	CertificateProviderInstance *GoogleNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance `field:"optional" json:"certificateProviderInstance" yaml:"certificateProviderInstance"`
	// grpc_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_network_security_server_tls_policy#grpc_endpoint GoogleNetworkSecurityServerTlsPolicy#grpc_endpoint}
	GrpcEndpoint *GoogleNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint `field:"optional" json:"grpcEndpoint" yaml:"grpcEndpoint"`
}

