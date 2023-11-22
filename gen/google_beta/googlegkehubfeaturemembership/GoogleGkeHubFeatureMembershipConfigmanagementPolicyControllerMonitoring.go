package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring struct {
	// Specifies the list of backends Policy Controller will export to. Specifying an empty value `[]` disables metrics export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_gke_hub_feature_membership#backends GoogleGkeHubFeatureMembership#backends}
	Backends *[]*string `field:"optional" json:"backends" yaml:"backends"`
}

