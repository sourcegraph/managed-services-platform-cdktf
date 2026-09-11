package workerversion


type WorkerVersionLimits struct {
	// CPU time limit in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#cpu_ms WorkerVersion#cpu_ms}
	CpuMs *float64 `field:"optional" json:"cpuMs" yaml:"cpuMs"`
	// Subrequest limit per request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#subrequests WorkerVersion#subrequests}
	Subrequests *float64 `field:"optional" json:"subrequests" yaml:"subrequests"`
}

