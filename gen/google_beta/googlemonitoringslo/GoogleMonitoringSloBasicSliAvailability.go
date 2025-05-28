package googlemonitoringslo


type GoogleMonitoringSloBasicSliAvailability struct {
	// Whether an availability SLI is enabled or not. Must be set to true. Defaults to 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_monitoring_slo#enabled GoogleMonitoringSlo#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

