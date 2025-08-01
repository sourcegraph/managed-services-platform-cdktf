package containercluster


type ContainerClusterNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfig struct {
	// List of fully-qualified-domain-names. IPv4s and port specification are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#fqdns ContainerCluster#fqdns}
	Fqdns *[]*string `field:"required" json:"fqdns" yaml:"fqdns"`
	// gcp_secret_manager_certificate_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#gcp_secret_manager_certificate_config ContainerCluster#gcp_secret_manager_certificate_config}
	GcpSecretManagerCertificateConfig *ContainerClusterNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfig `field:"required" json:"gcpSecretManagerCertificateConfig" yaml:"gcpSecretManagerCertificateConfig"`
}

