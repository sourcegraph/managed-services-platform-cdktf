package containercluster


type ContainerClusterNodeConfigLocalNvmeSsdBlockConfig struct {
	// Number of raw-block local NVMe SSD disks to be attached to the node.
	//
	// Each local SSD is 375 GB in size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/container_cluster#local_ssd_count ContainerCluster#local_ssd_count}
	LocalSsdCount *float64 `field:"required" json:"localSsdCount" yaml:"localSsdCount"`
}
