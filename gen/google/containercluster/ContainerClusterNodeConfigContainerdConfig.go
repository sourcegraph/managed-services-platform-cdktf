package containercluster


type ContainerClusterNodeConfigContainerdConfig struct {
	// private_registry_access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/container_cluster#private_registry_access_config ContainerCluster#private_registry_access_config}
	PrivateRegistryAccessConfig *ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfig `field:"optional" json:"privateRegistryAccessConfig" yaml:"privateRegistryAccessConfig"`
}

