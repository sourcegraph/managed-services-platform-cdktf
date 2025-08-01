package googlecontainercluster


type GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy struct {
	// Number of blue nodes to drain in a batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#batch_node_count GoogleContainerCluster#batch_node_count}
	BatchNodeCount *float64 `field:"optional" json:"batchNodeCount" yaml:"batchNodeCount"`
	// Percentage of the bool pool nodes to drain in a batch.
	//
	// The range of this field should be (0.0, 1.0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#batch_percentage GoogleContainerCluster#batch_percentage}
	BatchPercentage *float64 `field:"optional" json:"batchPercentage" yaml:"batchPercentage"`
	// Soak time after each batch gets drained.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#batch_soak_duration GoogleContainerCluster#batch_soak_duration}
	BatchSoakDuration *string `field:"optional" json:"batchSoakDuration" yaml:"batchSoakDuration"`
}

