package computeregioninstancetemplate


type ComputeRegionInstanceTemplateNetworkPerformanceConfig struct {
	// The egress bandwidth tier to enable. Possible values:TIER_1, DEFAULT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_instance_template#total_egress_bandwidth_tier ComputeRegionInstanceTemplate#total_egress_bandwidth_tier}
	TotalEgressBandwidthTier *string `field:"required" json:"totalEgressBandwidthTier" yaml:"totalEgressBandwidthTier"`
}

