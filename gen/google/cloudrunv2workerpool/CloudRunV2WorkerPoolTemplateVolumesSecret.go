package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplateVolumesSecret struct {
	// The name of the secret in Cloud Secret Manager.
	//
	// Format: {secret} if the secret is in the same project. projects/{project}/secrets/{secret} if the secret is in a different project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/cloud_run_v2_worker_pool#secret CloudRunV2WorkerPool#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// Integer representation of mode bits to use on created files by default.
	//
	// Must be a value between 0000 and 0777 (octal), defaulting to 0444. Directories within the path are not affected by this setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/cloud_run_v2_worker_pool#default_mode CloudRunV2WorkerPool#default_mode}
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/cloud_run_v2_worker_pool#items CloudRunV2WorkerPool#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
}

