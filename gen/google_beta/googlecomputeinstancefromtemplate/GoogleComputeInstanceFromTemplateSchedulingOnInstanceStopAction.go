package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateSchedulingOnInstanceStopAction struct {
	// If true, the contents of any attached Local SSD disks will be discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance_from_template#discard_local_ssd GoogleComputeInstanceFromTemplate#discard_local_ssd}
	DiscardLocalSsd interface{} `field:"optional" json:"discardLocalSsd" yaml:"discardLocalSsd"`
}

