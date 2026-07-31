package workersscript


type WorkersScriptLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#cpu_ms WorkersScript#cpu_ms}
	CpuMs *float64 `field:"optional" json:"cpuMs" yaml:"cpuMs"`
}

