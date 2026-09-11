package pipelinestream


type PipelineStreamSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#fields PipelineStream#fields}.
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#format PipelineStream#format}.
	Format *PipelineStreamSchemaFormat `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#inferred PipelineStream#inferred}.
	Inferred interface{} `field:"optional" json:"inferred" yaml:"inferred"`
}

