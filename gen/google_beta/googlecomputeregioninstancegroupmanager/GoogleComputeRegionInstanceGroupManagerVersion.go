package googlecomputeregioninstancegroupmanager


type GoogleComputeRegionInstanceGroupManagerVersion struct {
	// The full URL to an instance template from which all new instances of this version will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_group_manager#instance_template GoogleComputeRegionInstanceGroupManager#instance_template}
	InstanceTemplate *string `field:"required" json:"instanceTemplate" yaml:"instanceTemplate"`
	// Version name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_group_manager#name GoogleComputeRegionInstanceGroupManager#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// target_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_instance_group_manager#target_size GoogleComputeRegionInstanceGroupManager#target_size}
	TargetSize *GoogleComputeRegionInstanceGroupManagerVersionTargetSize `field:"optional" json:"targetSize" yaml:"targetSize"`
}

