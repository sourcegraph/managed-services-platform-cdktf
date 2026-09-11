package pipelinesink


type PipelineSinkConfigPartitioning struct {
	// The pattern of the date string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#time_pattern PipelineSink#time_pattern}
	TimePattern *string `field:"optional" json:"timePattern" yaml:"timePattern"`
}

