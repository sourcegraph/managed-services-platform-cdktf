package googlecomputeregioninstancegroupmanager


type GoogleComputeRegionInstanceGroupManagerVersionTargetSize struct {
	// The number of instances which are managed for this version. Conflicts with percent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_group_manager#fixed GoogleComputeRegionInstanceGroupManager#fixed}
	Fixed *float64 `field:"optional" json:"fixed" yaml:"fixed"`
	// The number of instances (calculated as percentage) which are managed for this version.
	//
	// Conflicts with fixed. Note that when using percent, rounding will be in favor of explicitly set target_size values; a managed instance group with 2 instances and 2 versions, one of which has a target_size.percent of 60 will create 2 instances of that version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_group_manager#percent GoogleComputeRegionInstanceGroupManager#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
}

