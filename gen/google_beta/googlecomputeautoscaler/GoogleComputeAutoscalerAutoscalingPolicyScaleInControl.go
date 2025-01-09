package googlecomputeautoscaler


type GoogleComputeAutoscalerAutoscalingPolicyScaleInControl struct {
	// max_scaled_in_replicas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_compute_autoscaler#max_scaled_in_replicas GoogleComputeAutoscaler#max_scaled_in_replicas}
	MaxScaledInReplicas *GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas `field:"optional" json:"maxScaledInReplicas" yaml:"maxScaledInReplicas"`
	// How long back autoscaling should look when computing recommendations to include directives regarding slower scale down, as described above.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_compute_autoscaler#time_window_sec GoogleComputeAutoscaler#time_window_sec}
	TimeWindowSec *float64 `field:"optional" json:"timeWindowSec" yaml:"timeWindowSec"`
}

