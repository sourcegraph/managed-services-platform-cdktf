package googlecomputeinstancetemplate


type GoogleComputeInstanceTemplateSchedulingGracefulShutdown struct {
	// Opts-in for graceful shutdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#enabled GoogleComputeInstanceTemplate#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// max_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_instance_template#max_duration GoogleComputeInstanceTemplate#max_duration}
	MaxDuration *GoogleComputeInstanceTemplateSchedulingGracefulShutdownMaxDuration `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

