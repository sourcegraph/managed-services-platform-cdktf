package computeregioninstancetemplate


type ComputeRegionInstanceTemplateNetworkInterfaceIpv6AccessConfig struct {
	// The service-level to be provided for IPv6 traffic when the subnet has an external subnet.
	//
	// Only PREMIUM tier is valid for IPv6
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_instance_template#network_tier ComputeRegionInstanceTemplate#network_tier}
	NetworkTier *string `field:"required" json:"networkTier" yaml:"networkTier"`
}

