package googlecomputeresourcepolicy


type GoogleComputeResourcePolicyInstanceSchedulePolicy struct {
	// Specifies the time zone to be used in interpreting the schedule.
	//
	// The value of this field must be a time zone name
	// from the tz database: http://en.wikipedia.org/wiki/Tz_database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_resource_policy#time_zone GoogleComputeResourcePolicy#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// The expiration time of the schedule. The timestamp is an RFC3339 string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_resource_policy#expiration_time GoogleComputeResourcePolicy#expiration_time}
	ExpirationTime *string `field:"optional" json:"expirationTime" yaml:"expirationTime"`
	// The start time of the schedule. The timestamp is an RFC3339 string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_resource_policy#start_time GoogleComputeResourcePolicy#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// vm_start_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_resource_policy#vm_start_schedule GoogleComputeResourcePolicy#vm_start_schedule}
	VmStartSchedule *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule `field:"optional" json:"vmStartSchedule" yaml:"vmStartSchedule"`
	// vm_stop_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_resource_policy#vm_stop_schedule GoogleComputeResourcePolicy#vm_stop_schedule}
	VmStopSchedule *GoogleComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule `field:"optional" json:"vmStopSchedule" yaml:"vmStopSchedule"`
}

