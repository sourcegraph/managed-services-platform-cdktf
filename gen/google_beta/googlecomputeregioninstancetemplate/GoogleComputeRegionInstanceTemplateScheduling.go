package googlecomputeregioninstancetemplate


type GoogleComputeRegionInstanceTemplateScheduling struct {
	// Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user).
	//
	// This defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#automatic_restart GoogleComputeRegionInstanceTemplate#automatic_restart}
	AutomaticRestart interface{} `field:"optional" json:"automaticRestart" yaml:"automaticRestart"`
	// Specifies the availability domain, which this instance should be scheduled on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#availability_domain GoogleComputeRegionInstanceTemplate#availability_domain}
	AvailabilityDomain *float64 `field:"optional" json:"availabilityDomain" yaml:"availabilityDomain"`
	// graceful_shutdown block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#graceful_shutdown GoogleComputeRegionInstanceTemplate#graceful_shutdown}
	GracefulShutdown *GoogleComputeRegionInstanceTemplateSchedulingGracefulShutdown `field:"optional" json:"gracefulShutdown" yaml:"gracefulShutdown"`
	// Specify the time in seconds for host error detection, the value must be within the range of [90, 330] with the increment of 30, if unset, the default behavior of host error recovery will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#host_error_timeout_seconds GoogleComputeRegionInstanceTemplate#host_error_timeout_seconds}
	HostErrorTimeoutSeconds *float64 `field:"optional" json:"hostErrorTimeoutSeconds" yaml:"hostErrorTimeoutSeconds"`
	// Specifies the action GCE should take when SPOT VM is preempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#instance_termination_action GoogleComputeRegionInstanceTemplate#instance_termination_action}
	InstanceTerminationAction *string `field:"optional" json:"instanceTerminationAction" yaml:"instanceTerminationAction"`
	// local_ssd_recovery_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#local_ssd_recovery_timeout GoogleComputeRegionInstanceTemplate#local_ssd_recovery_timeout}
	LocalSsdRecoveryTimeout interface{} `field:"optional" json:"localSsdRecoveryTimeout" yaml:"localSsdRecoveryTimeout"`
	// Specifies the frequency of planned maintenance events. The accepted values are: PERIODIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#maintenance_interval GoogleComputeRegionInstanceTemplate#maintenance_interval}
	MaintenanceInterval *string `field:"optional" json:"maintenanceInterval" yaml:"maintenanceInterval"`
	// max_run_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#max_run_duration GoogleComputeRegionInstanceTemplate#max_run_duration}
	MaxRunDuration *GoogleComputeRegionInstanceTemplateSchedulingMaxRunDuration `field:"optional" json:"maxRunDuration" yaml:"maxRunDuration"`
	// Minimum number of cpus for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#min_node_cpus GoogleComputeRegionInstanceTemplate#min_node_cpus}
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
	// node_affinities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#node_affinities GoogleComputeRegionInstanceTemplate#node_affinities}
	NodeAffinities interface{} `field:"optional" json:"nodeAffinities" yaml:"nodeAffinities"`
	// Defines the maintenance behavior for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#on_host_maintenance GoogleComputeRegionInstanceTemplate#on_host_maintenance}
	OnHostMaintenance *string `field:"optional" json:"onHostMaintenance" yaml:"onHostMaintenance"`
	// on_instance_stop_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#on_instance_stop_action GoogleComputeRegionInstanceTemplate#on_instance_stop_action}
	OnInstanceStopAction *GoogleComputeRegionInstanceTemplateSchedulingOnInstanceStopAction `field:"optional" json:"onInstanceStopAction" yaml:"onInstanceStopAction"`
	// Allows instance to be preempted. This defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#preemptible GoogleComputeRegionInstanceTemplate#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Whether the instance is spot. If this is set as SPOT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#provisioning_model GoogleComputeRegionInstanceTemplate#provisioning_model}
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
	// Specifies the timestamp, when the instance will be terminated, in RFC3339 text format.
	//
	// If specified, the instance termination action
	// will be performed at the termination time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_region_instance_template#termination_time GoogleComputeRegionInstanceTemplate#termination_time}
	TerminationTime *string `field:"optional" json:"terminationTime" yaml:"terminationTime"`
}

