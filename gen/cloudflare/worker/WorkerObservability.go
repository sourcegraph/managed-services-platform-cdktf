package worker


type WorkerObservability struct {
	// Whether observability is enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker#enabled Worker#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The sampling rate for observability. From 0 to 1 (1 = 100%, 0.1 = 10%).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker#head_sampling_rate Worker#head_sampling_rate}
	HeadSamplingRate *float64 `field:"optional" json:"headSamplingRate" yaml:"headSamplingRate"`
	// Log settings for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker#logs Worker#logs}
	Logs *WorkerObservabilityLogs `field:"optional" json:"logs" yaml:"logs"`
	// Trace settings for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker#traces Worker#traces}
	Traces *WorkerObservabilityTraces `field:"optional" json:"traces" yaml:"traces"`
}

