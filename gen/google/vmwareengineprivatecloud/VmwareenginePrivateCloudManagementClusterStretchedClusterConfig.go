package vmwareengineprivatecloud


type VmwareenginePrivateCloudManagementClusterStretchedClusterConfig struct {
	// Zone that will remain operational when connection between the two zones is lost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/vmwareengine_private_cloud#preferred_location VmwareenginePrivateCloud#preferred_location}
	PreferredLocation *string `field:"optional" json:"preferredLocation" yaml:"preferredLocation"`
	// Additional zone for a higher level of availability and load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/vmwareengine_private_cloud#secondary_location VmwareenginePrivateCloud#secondary_location}
	SecondaryLocation *string `field:"optional" json:"secondaryLocation" yaml:"secondaryLocation"`
}

