package gkehubfeature


type GkeHubFeatureSpec struct {
	// clusterupgrade block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature#clusterupgrade GkeHubFeature#clusterupgrade}
	Clusterupgrade *GkeHubFeatureSpecClusterupgrade `field:"optional" json:"clusterupgrade" yaml:"clusterupgrade"`
	// fleetobservability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature#fleetobservability GkeHubFeature#fleetobservability}
	Fleetobservability *GkeHubFeatureSpecFleetobservability `field:"optional" json:"fleetobservability" yaml:"fleetobservability"`
	// multiclusteringress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature#multiclusteringress GkeHubFeature#multiclusteringress}
	Multiclusteringress *GkeHubFeatureSpecMulticlusteringress `field:"optional" json:"multiclusteringress" yaml:"multiclusteringress"`
	// rbacrolebindingactuation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature#rbacrolebindingactuation GkeHubFeature#rbacrolebindingactuation}
	Rbacrolebindingactuation *GkeHubFeatureSpecRbacrolebindingactuation `field:"optional" json:"rbacrolebindingactuation" yaml:"rbacrolebindingactuation"`
}

