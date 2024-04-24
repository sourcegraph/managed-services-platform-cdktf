package googlecontainercluster


type GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings struct {
	// Time needed after draining entire blue pool. After this period, blue pool will be cleaned up.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_container_cluster#node_pool_soak_duration GoogleContainerCluster#node_pool_soak_duration}
	NodePoolSoakDuration *string `field:"optional" json:"nodePoolSoakDuration" yaml:"nodePoolSoakDuration"`
	// standard_rollout_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_container_cluster#standard_rollout_policy GoogleContainerCluster#standard_rollout_policy}
	StandardRolloutPolicy *GoogleContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy `field:"optional" json:"standardRolloutPolicy" yaml:"standardRolloutPolicy"`
}

