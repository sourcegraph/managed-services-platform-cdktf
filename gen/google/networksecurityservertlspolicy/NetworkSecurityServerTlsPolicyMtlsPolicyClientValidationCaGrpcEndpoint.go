package networksecurityservertlspolicy


type NetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint struct {
	// The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_security_server_tls_policy#target_uri NetworkSecurityServerTlsPolicy#target_uri}
	TargetUri *string `field:"required" json:"targetUri" yaml:"targetUri"`
}

