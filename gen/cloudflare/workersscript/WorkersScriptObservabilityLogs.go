package workersscript


type WorkersScriptObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#enabled WorkersScript#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Whether [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs) are enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#invocation_logs WorkersScript#invocation_logs}
	InvocationLogs interface{} `field:"required" json:"invocationLogs" yaml:"invocationLogs"`
	// A list of destinations where logs will be exported to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#destinations WorkersScript#destinations}
	Destinations *[]*string `field:"optional" json:"destinations" yaml:"destinations"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#head_sampling_rate WorkersScript#head_sampling_rate}
	HeadSamplingRate *float64 `field:"optional" json:"headSamplingRate" yaml:"headSamplingRate"`
	// Whether log persistence is enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#persist WorkersScript#persist}
	Persist interface{} `field:"optional" json:"persist" yaml:"persist"`
}

