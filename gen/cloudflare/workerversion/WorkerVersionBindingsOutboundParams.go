package workerversion


type WorkerVersionBindingsOutboundParams struct {
	// Name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#name WorkerVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

