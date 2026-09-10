package pipelinesink


type PipelineSinkConfigRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#file_size_bytes PipelineSink#file_size_bytes}
	FileSizeBytes *float64 `field:"optional" json:"fileSizeBytes" yaml:"fileSizeBytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#inactivity_seconds PipelineSink#inactivity_seconds}
	InactivitySeconds *float64 `field:"optional" json:"inactivitySeconds" yaml:"inactivitySeconds"`
	// Number of seconds to wait before rolling over to a new file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#interval_seconds PipelineSink#interval_seconds}
	IntervalSeconds *float64 `field:"optional" json:"intervalSeconds" yaml:"intervalSeconds"`
}

