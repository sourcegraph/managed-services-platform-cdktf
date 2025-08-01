package gkehubfeature


type GkeHubFeatureSpecClusterupgradePostConditions struct {
	// Amount of time to "soak" after a rollout has been finished before marking it COMPLETE. Cannot exceed 30 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature#soaking GkeHubFeature#soaking}
	Soaking *string `field:"required" json:"soaking" yaml:"soaking"`
}

