package gkehubfeature


type GkeHubFeatureSpecClusterupgradeGkeUpgradeOverrides struct {
	// post_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature#post_conditions GkeHubFeature#post_conditions}
	PostConditions *GkeHubFeatureSpecClusterupgradeGkeUpgradeOverridesPostConditions `field:"required" json:"postConditions" yaml:"postConditions"`
	// upgrade block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature#upgrade GkeHubFeature#upgrade}
	Upgrade *GkeHubFeatureSpecClusterupgradeGkeUpgradeOverridesUpgrade `field:"required" json:"upgrade" yaml:"upgrade"`
}

