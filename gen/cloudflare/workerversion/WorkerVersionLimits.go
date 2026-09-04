package workerversion


type WorkerVersionLimits struct {
	// CPU time limit in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/worker_version#cpu_ms WorkerVersion#cpu_ms}
	CpuMs *float64 `field:"required" json:"cpuMs" yaml:"cpuMs"`
}

