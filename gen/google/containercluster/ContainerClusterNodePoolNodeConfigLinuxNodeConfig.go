package containercluster


type ContainerClusterNodePoolNodeConfigLinuxNodeConfig struct {
	// cgroupMode specifies the cgroup mode to be used on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#cgroup_mode ContainerCluster#cgroup_mode}
	CgroupMode *string `field:"optional" json:"cgroupMode" yaml:"cgroupMode"`
	// hugepages_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#hugepages_config ContainerCluster#hugepages_config}
	HugepagesConfig *ContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfig `field:"optional" json:"hugepagesConfig" yaml:"hugepagesConfig"`
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#sysctls ContainerCluster#sysctls}
	Sysctls *map[string]*string `field:"optional" json:"sysctls" yaml:"sysctls"`
}

