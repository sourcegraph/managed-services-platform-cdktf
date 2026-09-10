package pipelinestream


type PipelineStreamWorkerBinding struct {
	// Indicates that the worker binding is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#enabled PipelineStream#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

