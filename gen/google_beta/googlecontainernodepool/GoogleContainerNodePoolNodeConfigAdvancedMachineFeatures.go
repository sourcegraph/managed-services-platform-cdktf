package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigAdvancedMachineFeatures struct {
	// The number of threads per physical core.
	//
	// To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_container_node_pool#threads_per_core GoogleContainerNodePool#threads_per_core}
	ThreadsPerCore *float64 `field:"required" json:"threadsPerCore" yaml:"threadsPerCore"`
}

