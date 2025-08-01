package gkeonpremvmwarecluster


type GkeonpremVmwareClusterNetworkConfigStaticIpConfigIpBlocks struct {
	// The network gateway used by the VMware User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_cluster#gateway GkeonpremVmwareCluster#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// ips block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_cluster#ips GkeonpremVmwareCluster#ips}
	Ips interface{} `field:"required" json:"ips" yaml:"ips"`
	// The netmask used by the VMware User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_cluster#netmask GkeonpremVmwareCluster#netmask}
	Netmask *string `field:"required" json:"netmask" yaml:"netmask"`
}

