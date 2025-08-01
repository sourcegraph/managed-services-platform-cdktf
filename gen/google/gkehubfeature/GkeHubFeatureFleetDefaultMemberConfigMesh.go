package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigMesh struct {
	// Whether to automatically manage Service Mesh Possible values: ["MANAGEMENT_UNSPECIFIED", "MANAGEMENT_AUTOMATIC", "MANAGEMENT_MANUAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature#management GkeHubFeature#management}
	Management *string `field:"required" json:"management" yaml:"management"`
}

