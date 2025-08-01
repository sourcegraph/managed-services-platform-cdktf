package workersdeployment


type WorkersDeploymentAnnotations struct {
	// Human-readable message about the deployment. Truncated to 100 bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_deployment#workers_message WorkersDeployment#workers_message}
	WorkersMessage *string `field:"optional" json:"workersMessage" yaml:"workersMessage"`
}

