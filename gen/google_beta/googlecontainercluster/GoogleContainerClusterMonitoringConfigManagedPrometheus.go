package googlecontainercluster


type GoogleContainerClusterMonitoringConfigManagedPrometheus struct {
	// Whether or not the managed collection is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// auto_monitoring_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_cluster#auto_monitoring_config GoogleContainerCluster#auto_monitoring_config}
	AutoMonitoringConfig *GoogleContainerClusterMonitoringConfigManagedPrometheusAutoMonitoringConfig `field:"optional" json:"autoMonitoringConfig" yaml:"autoMonitoringConfig"`
}

