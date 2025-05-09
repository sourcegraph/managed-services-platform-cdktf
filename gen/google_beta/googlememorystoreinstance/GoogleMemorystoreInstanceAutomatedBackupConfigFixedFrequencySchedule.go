package googlememorystoreinstance


type GoogleMemorystoreInstanceAutomatedBackupConfigFixedFrequencySchedule struct {
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_memorystore_instance#start_time GoogleMemorystoreInstance#start_time}
	StartTime *GoogleMemorystoreInstanceAutomatedBackupConfigFixedFrequencyScheduleStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

