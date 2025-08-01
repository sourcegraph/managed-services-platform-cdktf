package googlecontainercluster


type GoogleContainerClusterMasterAuth struct {
	// client_certificate_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#client_certificate_config GoogleContainerCluster#client_certificate_config}
	ClientCertificateConfig *GoogleContainerClusterMasterAuthClientCertificateConfig `field:"required" json:"clientCertificateConfig" yaml:"clientCertificateConfig"`
}

