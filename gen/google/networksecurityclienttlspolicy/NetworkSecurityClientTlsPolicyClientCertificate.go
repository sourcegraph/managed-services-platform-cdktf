package networksecurityclienttlspolicy


type NetworkSecurityClientTlsPolicyClientCertificate struct {
	// certificate_provider_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_client_tls_policy#certificate_provider_instance NetworkSecurityClientTlsPolicy#certificate_provider_instance}
	CertificateProviderInstance *NetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstance `field:"optional" json:"certificateProviderInstance" yaml:"certificateProviderInstance"`
	// grpc_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_client_tls_policy#grpc_endpoint NetworkSecurityClientTlsPolicy#grpc_endpoint}
	GrpcEndpoint *NetworkSecurityClientTlsPolicyClientCertificateGrpcEndpoint `field:"optional" json:"grpcEndpoint" yaml:"grpcEndpoint"`
}

