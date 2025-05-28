package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersPorts struct {
	// Port number the container listens on. This must be a valid TCP port number, 0 < containerPort < 65536.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_cloud_run_v2_service#container_port GoogleCloudRunV2Service#container_port}
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// If specified, used to specify which protocol to use. Allowed values are "http1" and "h2c".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_cloud_run_v2_service#name GoogleCloudRunV2Service#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

