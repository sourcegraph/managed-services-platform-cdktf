package spannerinstance


type SpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptions struct {
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/spanner_instance#overrides SpannerInstance#overrides}
	Overrides *SpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverrides `field:"required" json:"overrides" yaml:"overrides"`
	// replica_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/spanner_instance#replica_selection SpannerInstance#replica_selection}
	ReplicaSelection *SpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsReplicaSelection `field:"required" json:"replicaSelection" yaml:"replicaSelection"`
}

