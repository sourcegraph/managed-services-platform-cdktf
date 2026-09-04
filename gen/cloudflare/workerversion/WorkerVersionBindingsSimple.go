package workerversion


type WorkerVersionBindingsSimple struct {
	// The limit (requests per period).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/worker_version#limit WorkerVersion#limit}
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// The period in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/worker_version#period WorkerVersion#period}
	Period *float64 `field:"required" json:"period" yaml:"period"`
}

