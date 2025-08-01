package googlecomputeresourcepolicy


type GoogleComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule struct {
	// Specifies the frequency for the operation, using the unix-cron format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_resource_policy#schedule GoogleComputeResourcePolicy#schedule}
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
}

