package googleclouddeployautomation


type GoogleClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry struct {
	// Required.
	//
	// Total number of retries. Retry is skipped if set to 0; The minimum value is 1, and the maximum value is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_clouddeploy_automation#attempts GoogleClouddeployAutomation#attempts}
	Attempts *string `field:"required" json:"attempts" yaml:"attempts"`
	// Optional.
	//
	// The pattern of how wait time will be increased. Default is linear. Backoff mode will be ignored if wait is 0. Possible values: ["BACKOFF_MODE_UNSPECIFIED", "BACKOFF_MODE_LINEAR", "BACKOFF_MODE_EXPONENTIAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_clouddeploy_automation#backoff_mode GoogleClouddeployAutomation#backoff_mode}
	BackoffMode *string `field:"optional" json:"backoffMode" yaml:"backoffMode"`
	// Optional.
	//
	// How long to wait for the first retry. Default is 0, and the maximum value is 14d. A duration in seconds with up to nine fractional digits, ending with 's'. Example: '3.5s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_clouddeploy_automation#wait GoogleClouddeployAutomation#wait}
	Wait *string `field:"optional" json:"wait" yaml:"wait"`
}

