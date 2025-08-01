package googlecloudbuildworkerpool


type GoogleCloudbuildWorkerPoolWorkerConfig struct {
	// Size of the disk attached to the worker, in GB.
	//
	// See [Worker pool config file](https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). Specify a value of up to 1000. If `0` is specified, Cloud Build will use a standard disk size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_worker_pool#disk_size_gb GoogleCloudbuildWorkerPool#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Machine type of a worker, such as `n1-standard-1`.
	//
	// See [Worker pool config file](https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). If left blank, Cloud Build will use `n1-standard-1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_worker_pool#machine_type GoogleCloudbuildWorkerPool#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// If true, workers are created without any public address, which prevents network egress to public IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_worker_pool#no_external_ip GoogleCloudbuildWorkerPool#no_external_ip}
	NoExternalIp interface{} `field:"optional" json:"noExternalIp" yaml:"noExternalIp"`
}

