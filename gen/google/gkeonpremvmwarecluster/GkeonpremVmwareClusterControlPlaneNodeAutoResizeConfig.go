package gkeonpremvmwarecluster


type GkeonpremVmwareClusterControlPlaneNodeAutoResizeConfig struct {
	// Whether to enable control plane node auto resizing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_vmware_cluster#enabled GkeonpremVmwareCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

