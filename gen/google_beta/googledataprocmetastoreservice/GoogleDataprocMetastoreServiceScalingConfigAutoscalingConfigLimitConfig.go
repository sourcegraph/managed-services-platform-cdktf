package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig struct {
	// The maximum scaling factor that the service will autoscale to. The default value is 6.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_metastore_service#max_scaling_factor GoogleDataprocMetastoreService#max_scaling_factor}
	MaxScalingFactor *float64 `field:"optional" json:"maxScalingFactor" yaml:"maxScalingFactor"`
	// The minimum scaling factor that the service will autoscale to. The default value is 0.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_metastore_service#min_scaling_factor GoogleDataprocMetastoreService#min_scaling_factor}
	MinScalingFactor *float64 `field:"optional" json:"minScalingFactor" yaml:"minScalingFactor"`
}

