package networksecurityservertlspolicy


type NetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCa struct {
	// certificate_provider_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_server_tls_policy#certificate_provider_instance NetworkSecurityServerTlsPolicy#certificate_provider_instance}
	CertificateProviderInstance *NetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance `field:"optional" json:"certificateProviderInstance" yaml:"certificateProviderInstance"`
	// grpc_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_server_tls_policy#grpc_endpoint NetworkSecurityServerTlsPolicy#grpc_endpoint}
	GrpcEndpoint *NetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint `field:"optional" json:"grpcEndpoint" yaml:"grpcEndpoint"`
}

