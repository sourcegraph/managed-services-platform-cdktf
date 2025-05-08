package googlecomputeinstancegroupmanager


type GoogleComputeInstanceGroupManagerParams struct {
	// Resource manager tags to bind to the managed instance group.
	//
	// The tags are key-value pairs. Keys must be in the format tagKeys/123 and values in the format tagValues/456.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance_group_manager#resource_manager_tags GoogleComputeInstanceGroupManager#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
}

