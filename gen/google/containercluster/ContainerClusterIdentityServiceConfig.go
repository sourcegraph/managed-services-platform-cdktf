package containercluster

type ContainerClusterIdentityServiceConfig struct {
	// Whether to enable the Identity Service component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled any `field:"optional" json:"enabled" yaml:"enabled"`
}
