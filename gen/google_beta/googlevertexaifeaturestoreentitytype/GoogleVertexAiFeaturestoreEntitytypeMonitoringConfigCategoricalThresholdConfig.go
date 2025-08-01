package googlevertexaifeaturestoreentitytype


type GoogleVertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig struct {
	// Specify a threshold value that can trigger the alert.
	//
	// For categorical feature, the distribution distance is calculated by L-inifinity norm. Each feature must have a non-zero threshold if they need to be monitored. Otherwise no alert will be triggered for that feature. The default value is 0.3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_featurestore_entitytype#value GoogleVertexAiFeaturestoreEntitytype#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

