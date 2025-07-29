package googlevertexaiendpointwithmodelgardendeployment


type GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeExec struct {
	// Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem.
	//
	// The command is simply exec'd, it is not run inside a shell, so
	// traditional shell instructions ('|', etc) won't work. To use a shell, you
	// need to explicitly call out to that shell. Exit status of 0 is treated as
	// live/healthy and non-zero is unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#command GoogleVertexAiEndpointWithModelGardenDeployment#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
}

