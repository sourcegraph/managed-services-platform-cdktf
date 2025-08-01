package googlemonitoringslo


type GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency struct {
	// A duration string, e.g. 10s. Good service is defined to be the count of requests made to this service that return in no more than threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_slo#threshold GoogleMonitoringSlo#threshold}
	Threshold *string `field:"required" json:"threshold" yaml:"threshold"`
}

