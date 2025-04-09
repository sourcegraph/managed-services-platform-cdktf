package googlecomputeregionautoscaler


type GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas struct {
	// Specifies a fixed number of VM instances. This must be a positive integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_region_autoscaler#fixed GoogleComputeRegionAutoscaler#fixed}
	Fixed *float64 `field:"optional" json:"fixed" yaml:"fixed"`
	// Specifies a percentage of instances between 0 to 100%, inclusive. For example, specify 80 for 80%.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_compute_region_autoscaler#percent GoogleComputeRegionAutoscaler#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
}

