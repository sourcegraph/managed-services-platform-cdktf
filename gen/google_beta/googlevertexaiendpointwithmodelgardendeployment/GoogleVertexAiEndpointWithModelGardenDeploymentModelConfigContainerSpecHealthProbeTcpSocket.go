package googlevertexaiendpointwithmodelgardendeployment


type GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeTcpSocket struct {
	// Optional: Host name to connect to, defaults to the model serving container's IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#host GoogleVertexAiEndpointWithModelGardenDeployment#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Number of the port to access on the container. Number must be in the range 1 to 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#port GoogleVertexAiEndpointWithModelGardenDeployment#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

