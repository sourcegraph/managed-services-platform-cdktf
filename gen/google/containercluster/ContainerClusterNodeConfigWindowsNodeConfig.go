package containercluster


type ContainerClusterNodeConfigWindowsNodeConfig struct {
	// The OS Version of the windows nodepool.Values are OS_VERSION_UNSPECIFIED,OS_VERSION_LTSC2019 and OS_VERSION_LTSC2022.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#osversion ContainerCluster#osversion}
	Osversion *string `field:"optional" json:"osversion" yaml:"osversion"`
}

