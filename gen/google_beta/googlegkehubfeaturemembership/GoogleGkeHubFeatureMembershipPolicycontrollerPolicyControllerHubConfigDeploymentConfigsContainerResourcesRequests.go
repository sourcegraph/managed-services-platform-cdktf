package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesRequests struct {
	// CPU requirement expressed in Kubernetes resource units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gke_hub_feature_membership#cpu GoogleGkeHubFeatureMembership#cpu}
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Memory requirement expressed in Kubernetes resource units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_gke_hub_feature_membership#memory GoogleGkeHubFeatureMembership#memory}
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

