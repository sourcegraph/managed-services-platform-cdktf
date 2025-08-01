package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaints struct {
	// Specifies the nodes operating system (default: LINUX). Possible values: ["EFFECT_UNSPECIFIED", "PREFER_NO_SCHEDULE", "NO_EXECUTE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_cluster#effect GkeonpremBareMetalCluster#effect}
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Key associated with the effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_cluster#key GkeonpremBareMetalCluster#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Value associated with the effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_cluster#value GkeonpremBareMetalCluster#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

