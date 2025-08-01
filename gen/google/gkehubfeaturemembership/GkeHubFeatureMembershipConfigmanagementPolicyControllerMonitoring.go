package gkehubfeaturemembership


type GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring struct {
	// Specifies the list of backends Policy Controller will export to. Specifying an empty value `[]` disables metrics export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_hub_feature_membership#backends GkeHubFeatureMembership#backends}
	Backends *[]*string `field:"optional" json:"backends" yaml:"backends"`
}

