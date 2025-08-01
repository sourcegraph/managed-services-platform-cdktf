package computeinstancegroupmanager


type ComputeInstanceGroupManagerResourcePolicies struct {
	// The URL of the workload policy that is specified for this managed instance group.
	//
	// It can be a full or partial URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_instance_group_manager#workload_policy ComputeInstanceGroupManager#workload_policy}
	WorkloadPolicy *string `field:"optional" json:"workloadPolicy" yaml:"workloadPolicy"`
}

