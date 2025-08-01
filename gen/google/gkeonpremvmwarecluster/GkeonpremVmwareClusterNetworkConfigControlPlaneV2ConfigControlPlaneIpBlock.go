package gkeonpremvmwarecluster


type GkeonpremVmwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlock struct {
	// The network gateway used by the VMware User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_cluster#gateway GkeonpremVmwareCluster#gateway}
	Gateway *string `field:"optional" json:"gateway" yaml:"gateway"`
	// ips block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_cluster#ips GkeonpremVmwareCluster#ips}
	Ips interface{} `field:"optional" json:"ips" yaml:"ips"`
	// The netmask used by the VMware User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_cluster#netmask GkeonpremVmwareCluster#netmask}
	Netmask *string `field:"optional" json:"netmask" yaml:"netmask"`
}

