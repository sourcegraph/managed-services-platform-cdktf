package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigWindowsNodeConfig struct {
	// The OS Version of the windows nodepool.Values are OS_VERSION_UNSPECIFIED,OS_VERSION_LTSC2019 and OS_VERSION_LTSC2022.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_container_node_pool#osversion GoogleContainerNodePool#osversion}
	Osversion *string `field:"optional" json:"osversion" yaml:"osversion"`
}

