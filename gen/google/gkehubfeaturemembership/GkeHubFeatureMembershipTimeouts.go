package gkehubfeaturemembership


type GkeHubFeatureMembershipTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature_membership#create GkeHubFeatureMembership#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature_membership#delete GkeHubFeatureMembership#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature_membership#update GkeHubFeatureMembership#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

