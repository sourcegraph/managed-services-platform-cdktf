package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipConfigmanagement struct {
	// binauthz block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_feature_membership#binauthz GoogleGkeHubFeatureMembership#binauthz}
	Binauthz *GoogleGkeHubFeatureMembershipConfigmanagementBinauthz `field:"optional" json:"binauthz" yaml:"binauthz"`
	// config_sync block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_feature_membership#config_sync GoogleGkeHubFeatureMembership#config_sync}
	ConfigSync *GoogleGkeHubFeatureMembershipConfigmanagementConfigSync `field:"optional" json:"configSync" yaml:"configSync"`
	// hierarchy_controller block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_feature_membership#hierarchy_controller GoogleGkeHubFeatureMembership#hierarchy_controller}
	HierarchyController *GoogleGkeHubFeatureMembershipConfigmanagementHierarchyController `field:"optional" json:"hierarchyController" yaml:"hierarchyController"`
	// Set this field to MANAGEMENT_AUTOMATIC to enable Config Sync auto-upgrades, and set this field to MANAGEMENT_MANUAL or MANAGEMENT_UNSPECIFIED to disable Config Sync auto-upgrades.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_feature_membership#management GoogleGkeHubFeatureMembership#management}
	Management *string `field:"optional" json:"management" yaml:"management"`
	// policy_controller block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_feature_membership#policy_controller GoogleGkeHubFeatureMembership#policy_controller}
	PolicyController *GoogleGkeHubFeatureMembershipConfigmanagementPolicyController `field:"optional" json:"policyController" yaml:"policyController"`
	// Optional. Version of ACM to install. Defaults to the latest version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_feature_membership#version GoogleGkeHubFeatureMembership#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

