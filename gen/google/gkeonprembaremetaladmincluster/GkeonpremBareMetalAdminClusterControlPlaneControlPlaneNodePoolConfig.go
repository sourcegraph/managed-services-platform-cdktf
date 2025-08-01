package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfig struct {
	// node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_admin_cluster#node_pool_config GkeonpremBareMetalAdminCluster#node_pool_config}
	NodePoolConfig *GkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfig `field:"required" json:"nodePoolConfig" yaml:"nodePoolConfig"`
}

