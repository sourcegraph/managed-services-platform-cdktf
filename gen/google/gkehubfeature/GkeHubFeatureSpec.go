package gkehubfeature


type GkeHubFeatureSpec struct {
	// fleetobservability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gke_hub_feature#fleetobservability GkeHubFeature#fleetobservability}
	Fleetobservability *GkeHubFeatureSpecFleetobservability `field:"optional" json:"fleetobservability" yaml:"fleetobservability"`
	// multiclusteringress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gke_hub_feature#multiclusteringress GkeHubFeature#multiclusteringress}
	Multiclusteringress *GkeHubFeatureSpecMulticlusteringress `field:"optional" json:"multiclusteringress" yaml:"multiclusteringress"`
}

