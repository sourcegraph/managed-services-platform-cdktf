package workerversion


type WorkerVersionBindingsOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#params WorkerVersion#params}
	Params interface{} `field:"optional" json:"params" yaml:"params"`
	// Outbound worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#worker WorkerVersion#worker}
	Worker *WorkerVersionBindingsOutboundWorker `field:"optional" json:"worker" yaml:"worker"`
}

