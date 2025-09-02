package googlecomputeinstancegroupmanager


type GoogleComputeInstanceGroupManagerResourcePolicies struct {
	// The URL of the workload policy that is specified for this managed instance group.
	//
	// It can be a full or partial URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_group_manager#workload_policy GoogleComputeInstanceGroupManager#workload_policy}
	WorkloadPolicy *string `field:"optional" json:"workloadPolicy" yaml:"workloadPolicy"`
}

