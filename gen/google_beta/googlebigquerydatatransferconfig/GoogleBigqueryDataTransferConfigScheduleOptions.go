package googlebigquerydatatransferconfig


type GoogleBigqueryDataTransferConfigScheduleOptions struct {
	// If true, automatic scheduling of data transfer runs for this configuration will be disabled.
	//
	// The runs can be started on ad-hoc
	// basis using transferConfigs.startManualRuns API. When automatic
	// scheduling is disabled, the TransferConfig.schedule field will
	// be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_data_transfer_config#disable_auto_scheduling GoogleBigqueryDataTransferConfig#disable_auto_scheduling}
	DisableAutoScheduling interface{} `field:"optional" json:"disableAutoScheduling" yaml:"disableAutoScheduling"`
	// Defines time to stop scheduling transfer runs.
	//
	// A transfer run cannot be
	// scheduled at or after the end time. The end time can be changed at any
	// moment. The time when a data transfer can be triggered manually is not
	// limited by this option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_data_transfer_config#end_time GoogleBigqueryDataTransferConfig#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Specifies time to start scheduling transfer runs.
	//
	// The first run will be
	// scheduled at or after the start time according to a recurrence pattern
	// defined in the schedule string. The start time can be changed at any
	// moment. The time when a data transfer can be triggered manually is not
	// limited by this option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_data_transfer_config#start_time GoogleBigqueryDataTransferConfig#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

