package gkeonpremvmwarecluster


type GkeonpremVmwareClusterNetworkConfigDhcpIpConfig struct {
	// enabled is a flag to mark if DHCP IP allocation is used for VMware user clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_cluster#enabled GkeonpremVmwareCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

