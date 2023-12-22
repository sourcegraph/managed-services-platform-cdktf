package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyAlertStrategyNotificationRateLimit struct {
	// Not more than one notification per period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_monitoring_alert_policy#period GoogleMonitoringAlertPolicy#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
}

