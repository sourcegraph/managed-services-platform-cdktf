package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterControlPlaneControlPlaneNodePoolConfig struct {
	// node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gkeonprem_bare_metal_cluster#node_pool_config GoogleGkeonpremBareMetalCluster#node_pool_config}
	NodePoolConfig *GoogleGkeonpremBareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfig `field:"required" json:"nodePoolConfig" yaml:"nodePoolConfig"`
}

