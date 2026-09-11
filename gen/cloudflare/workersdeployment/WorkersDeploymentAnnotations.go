package workersdeployment


type WorkersDeploymentAnnotations struct {
	// Human-readable message about the deployment. Truncated to 1000 bytes if longer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_deployment#workers_message WorkersDeployment#workers_message}
	WorkersMessage *string `field:"optional" json:"workersMessage" yaml:"workersMessage"`
}

