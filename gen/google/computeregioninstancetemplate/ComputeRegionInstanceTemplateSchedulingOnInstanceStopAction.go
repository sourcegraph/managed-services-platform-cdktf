package computeregioninstancetemplate


type ComputeRegionInstanceTemplateSchedulingOnInstanceStopAction struct {
	// If true, the contents of any attached Local SSD disks will be discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_region_instance_template#discard_local_ssd ComputeRegionInstanceTemplate#discard_local_ssd}
	DiscardLocalSsd interface{} `field:"optional" json:"discardLocalSsd" yaml:"discardLocalSsd"`
}

