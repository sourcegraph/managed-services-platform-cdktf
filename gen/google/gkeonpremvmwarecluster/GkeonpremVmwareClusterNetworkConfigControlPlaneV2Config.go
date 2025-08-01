package gkeonpremvmwarecluster


type GkeonpremVmwareClusterNetworkConfigControlPlaneV2Config struct {
	// control_plane_ip_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_cluster#control_plane_ip_block GkeonpremVmwareCluster#control_plane_ip_block}
	ControlPlaneIpBlock *GkeonpremVmwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlock `field:"optional" json:"controlPlaneIpBlock" yaml:"controlPlaneIpBlock"`
}

