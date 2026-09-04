package workerversion


type WorkerVersionBindingsOutboundWorker struct {
	// Environment of the outbound worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/worker_version#environment WorkerVersion#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Name of the outbound worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/worker_version#service WorkerVersion#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

