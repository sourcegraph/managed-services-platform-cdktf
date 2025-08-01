package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplateContainersEnvValueSource struct {
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/cloud_run_v2_worker_pool#secret_key_ref CloudRunV2WorkerPool#secret_key_ref}
	SecretKeyRef *CloudRunV2WorkerPoolTemplateContainersEnvValueSourceSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

