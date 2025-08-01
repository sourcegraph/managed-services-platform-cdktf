package cloudrunv2service


type CloudRunV2ServiceTemplateNodeSelector struct {
	// The GPU to attach to an instance. See https://cloud.google.com/run/docs/configuring/services/gpu for configuring GPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/cloud_run_v2_service#accelerator CloudRunV2Service#accelerator}
	Accelerator *string `field:"required" json:"accelerator" yaml:"accelerator"`
}

