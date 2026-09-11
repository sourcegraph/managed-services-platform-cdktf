package workerversion


type WorkerVersionContainers struct {
	// Select which Durable Object class should get this container attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#class_name WorkerVersion#class_name}
	ClassName *string `field:"required" json:"className" yaml:"className"`
}

