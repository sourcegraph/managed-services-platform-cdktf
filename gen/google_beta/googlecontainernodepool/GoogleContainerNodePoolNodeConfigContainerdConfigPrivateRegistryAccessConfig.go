package googlecontainernodepool

type GoogleContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig struct {
	// Whether or not private registries are configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_node_pool#enabled GoogleContainerNodePool#enabled}
	Enabled any `field:"required" json:"enabled" yaml:"enabled"`
	// certificate_authority_domain_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_node_pool#certificate_authority_domain_config GoogleContainerNodePool#certificate_authority_domain_config}
	CertificateAuthorityDomainConfig any `field:"optional" json:"certificateAuthorityDomainConfig" yaml:"certificateAuthorityDomainConfig"`
}
