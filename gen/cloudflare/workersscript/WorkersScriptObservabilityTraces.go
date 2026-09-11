package workersscript


type WorkersScriptObservabilityTraces struct {
	// A list of destinations where traces will be exported to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_script#destinations WorkersScript#destinations}
	Destinations *[]*string `field:"optional" json:"destinations" yaml:"destinations"`
	// Whether traces are enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_script#enabled WorkersScript#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_script#head_sampling_rate WorkersScript#head_sampling_rate}
	HeadSamplingRate *float64 `field:"optional" json:"headSamplingRate" yaml:"headSamplingRate"`
	// Whether trace persistence is enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_script#persist WorkersScript#persist}
	Persist interface{} `field:"optional" json:"persist" yaml:"persist"`
}

