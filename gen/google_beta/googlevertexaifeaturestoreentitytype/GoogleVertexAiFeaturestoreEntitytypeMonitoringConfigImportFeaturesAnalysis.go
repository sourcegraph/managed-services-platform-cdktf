package googlevertexaifeaturestoreentitytype


type GoogleVertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis struct {
	// Defines the baseline to do anomaly detection for feature values imported by each [entityTypes.importFeatureValues][] operation. The value must be one of the values below: LATEST_STATS: Choose the later one statistics generated by either most recent snapshot analysis or previous import features analysis. If non of them exists, skip anomaly detection and only generate a statistics. MOST_RECENT_SNAPSHOT_STATS: Use the statistics generated by the most recent snapshot analysis if exists. PREVIOUS_IMPORT_FEATURES_STATS: Use the statistics generated by the previous import features analysis if exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_vertex_ai_featurestore_entitytype#anomaly_detection_baseline GoogleVertexAiFeaturestoreEntitytype#anomaly_detection_baseline}
	AnomalyDetectionBaseline *string `field:"optional" json:"anomalyDetectionBaseline" yaml:"anomalyDetectionBaseline"`
	// Whether to enable / disable / inherite default hebavior for import features analysis.
	//
	// The value must be one of the values below:
	// DEFAULT: The default behavior of whether to enable the monitoring. EntityType-level config: disabled.
	// ENABLED: Explicitly enables import features analysis. EntityType-level config: by default enables import features analysis for all Features under it.
	// DISABLED: Explicitly disables import features analysis. EntityType-level config: by default disables import features analysis for all Features under it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_vertex_ai_featurestore_entitytype#state GoogleVertexAiFeaturestoreEntitytype#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}
