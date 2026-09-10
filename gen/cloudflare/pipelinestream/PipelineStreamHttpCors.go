package pipelinestream


type PipelineStreamHttpCors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#origins PipelineStream#origins}.
	Origins *[]*string `field:"optional" json:"origins" yaml:"origins"`
}

