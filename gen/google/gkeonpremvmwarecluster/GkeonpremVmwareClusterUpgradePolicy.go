package gkeonpremvmwarecluster


type GkeonpremVmwareClusterUpgradePolicy struct {
	// Controls whether the upgrade applies to the control plane only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_cluster#control_plane_only GkeonpremVmwareCluster#control_plane_only}
	ControlPlaneOnly interface{} `field:"optional" json:"controlPlaneOnly" yaml:"controlPlaneOnly"`
}

