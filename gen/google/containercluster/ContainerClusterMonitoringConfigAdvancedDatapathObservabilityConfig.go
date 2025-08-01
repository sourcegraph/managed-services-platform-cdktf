package containercluster


type ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig struct {
	// Whether or not the advanced datapath metrics are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#enable_metrics ContainerCluster#enable_metrics}
	EnableMetrics interface{} `field:"required" json:"enableMetrics" yaml:"enableMetrics"`
	// Whether or not Relay is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#enable_relay ContainerCluster#enable_relay}
	EnableRelay interface{} `field:"required" json:"enableRelay" yaml:"enableRelay"`
}

