package datacloudflarepipelinestream


type DataCloudflarePipelineStreamFilter struct {
	// Specifies the public ID of the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/pipeline_stream#pipeline_id DataCloudflarePipelineStream#pipeline_id}
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
}

