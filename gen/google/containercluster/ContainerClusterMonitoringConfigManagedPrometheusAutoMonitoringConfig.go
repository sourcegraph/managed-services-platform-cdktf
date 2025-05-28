package containercluster


type ContainerClusterMonitoringConfigManagedPrometheusAutoMonitoringConfig struct {
	// The scope of auto-monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/container_cluster#scope ContainerCluster#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

