package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterControlPlaneControlPlaneNodePoolConfig struct {
	// node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_cluster#node_pool_config GkeonpremBareMetalCluster#node_pool_config}
	NodePoolConfig *GkeonpremBareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfig `field:"required" json:"nodePoolConfig" yaml:"nodePoolConfig"`
}

