package googlecomputeinstance


type GoogleComputeInstanceScheduling struct {
	// Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#automatic_restart GoogleComputeInstance#automatic_restart}
	AutomaticRestart interface{} `field:"optional" json:"automaticRestart" yaml:"automaticRestart"`
	// Specifies the availability domain, which this instance should be scheduled on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#availability_domain GoogleComputeInstance#availability_domain}
	AvailabilityDomain *float64 `field:"optional" json:"availabilityDomain" yaml:"availabilityDomain"`
	// graceful_shutdown block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#graceful_shutdown GoogleComputeInstance#graceful_shutdown}
	GracefulShutdown *GoogleComputeInstanceSchedulingGracefulShutdown `field:"optional" json:"gracefulShutdown" yaml:"gracefulShutdown"`
	// Specify the time in seconds for host error detection, the value must be within the range of [90, 330] with the increment of 30, if unset, the default behavior of host error recovery will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#host_error_timeout_seconds GoogleComputeInstance#host_error_timeout_seconds}
	HostErrorTimeoutSeconds *float64 `field:"optional" json:"hostErrorTimeoutSeconds" yaml:"hostErrorTimeoutSeconds"`
	// Specifies the action GCE should take when SPOT VM is preempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#instance_termination_action GoogleComputeInstance#instance_termination_action}
	InstanceTerminationAction *string `field:"optional" json:"instanceTerminationAction" yaml:"instanceTerminationAction"`
	// local_ssd_recovery_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#local_ssd_recovery_timeout GoogleComputeInstance#local_ssd_recovery_timeout}
	LocalSsdRecoveryTimeout *GoogleComputeInstanceSchedulingLocalSsdRecoveryTimeout `field:"optional" json:"localSsdRecoveryTimeout" yaml:"localSsdRecoveryTimeout"`
	// Specifies the frequency of planned maintenance events. The accepted values are: PERIODIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#maintenance_interval GoogleComputeInstance#maintenance_interval}
	MaintenanceInterval *string `field:"optional" json:"maintenanceInterval" yaml:"maintenanceInterval"`
	// max_run_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#max_run_duration GoogleComputeInstance#max_run_duration}
	MaxRunDuration *GoogleComputeInstanceSchedulingMaxRunDuration `field:"optional" json:"maxRunDuration" yaml:"maxRunDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#min_node_cpus GoogleComputeInstance#min_node_cpus}.
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
	// node_affinities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#node_affinities GoogleComputeInstance#node_affinities}
	NodeAffinities interface{} `field:"optional" json:"nodeAffinities" yaml:"nodeAffinities"`
	// Describes maintenance behavior for the instance. One of MIGRATE or TERMINATE,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#on_host_maintenance GoogleComputeInstance#on_host_maintenance}
	OnHostMaintenance *string `field:"optional" json:"onHostMaintenance" yaml:"onHostMaintenance"`
	// on_instance_stop_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#on_instance_stop_action GoogleComputeInstance#on_instance_stop_action}
	OnInstanceStopAction *GoogleComputeInstanceSchedulingOnInstanceStopAction `field:"optional" json:"onInstanceStopAction" yaml:"onInstanceStopAction"`
	// Whether the instance is preemptible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#preemptible GoogleComputeInstance#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Whether the instance is spot. If this is set as SPOT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#provisioning_model GoogleComputeInstance#provisioning_model}
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
	// Specifies the timestamp, when the instance will be terminated, in RFC3339 text format.
	//
	// If specified, the instance termination action
	// will be performed at the termination time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instance#termination_time GoogleComputeInstance#termination_time}
	TerminationTime *string `field:"optional" json:"terminationTime" yaml:"terminationTime"`
}

