package googlestoragebucket


type GoogleStorageBucketIpFilter struct {
	// The mode of the IP filter. Valid values are 'Enabled' and 'Disabled'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_storage_bucket#mode GoogleStorageBucket#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// public_network_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_storage_bucket#public_network_source GoogleStorageBucket#public_network_source}
	PublicNetworkSource *GoogleStorageBucketIpFilterPublicNetworkSource `field:"optional" json:"publicNetworkSource" yaml:"publicNetworkSource"`
	// vpc_network_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_storage_bucket#vpc_network_sources GoogleStorageBucket#vpc_network_sources}
	VpcNetworkSources interface{} `field:"optional" json:"vpcNetworkSources" yaml:"vpcNetworkSources"`
}

