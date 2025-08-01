package gkehubfeaturemembership


type GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigDeploymentConfigsContainerResourcesLimits struct {
	// CPU requirement expressed in Kubernetes resource units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature_membership#cpu GkeHubFeatureMembership#cpu}
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Memory requirement expressed in Kubernetes resource units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature_membership#memory GkeHubFeatureMembership#memory}
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

