package containercluster


type ContainerClusterControlPlaneEndpointsConfigIpEndpointsConfig struct {
	// Controls whether to allow direct IP access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

