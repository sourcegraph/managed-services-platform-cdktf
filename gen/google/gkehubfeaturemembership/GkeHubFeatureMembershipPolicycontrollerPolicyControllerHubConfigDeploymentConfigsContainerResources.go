package gkehubfeaturemembership


type GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResources struct {
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature_membership#limits GkeHubFeatureMembership#limits}
	Limits *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	// requests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature_membership#requests GkeHubFeatureMembership#requests}
	Requests *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

