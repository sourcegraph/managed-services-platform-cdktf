package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger struct {
	// The score range is in [0,100].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dataplex_datascan#score_threshold GoogleDataplexDatascan#score_threshold}
	ScoreThreshold *float64 `field:"optional" json:"scoreThreshold" yaml:"scoreThreshold"`
}

