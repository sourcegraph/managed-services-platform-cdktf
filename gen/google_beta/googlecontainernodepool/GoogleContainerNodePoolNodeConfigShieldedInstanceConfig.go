package googlecontainernodepool

type GoogleContainerNodePoolNodeConfigShieldedInstanceConfig struct {
	// Defines whether the instance has integrity monitoring enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_node_pool#enable_integrity_monitoring GoogleContainerNodePool#enable_integrity_monitoring}
	EnableIntegrityMonitoring any `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Defines whether the instance has Secure Boot enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_node_pool#enable_secure_boot GoogleContainerNodePool#enable_secure_boot}
	EnableSecureBoot any `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
}
