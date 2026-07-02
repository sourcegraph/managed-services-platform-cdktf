package gkeonprembaremetalcluster

type GkeonpremBareMetalClusterLoadBalancerManualLbConfig struct {
	// Whether manual load balancing is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_cluster#enabled GkeonpremBareMetalCluster#enabled}
	Enabled any `field:"required" json:"enabled" yaml:"enabled"`
}
