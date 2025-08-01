package dataplexdatascan


type DataplexDatascanExecutionSpecTriggerSchedule struct {
	// Cron schedule for running scans periodically. This field is required for Schedule scans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataplex_datascan#cron DataplexDatascan#cron}
	Cron *string `field:"required" json:"cron" yaml:"cron"`
}

