package pipelinesink


type PipelineSinkSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#fields PipelineSink#fields}.
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#format PipelineSink#format}.
	Format *PipelineSinkSchemaFormat `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#inferred PipelineSink#inferred}.
	Inferred interface{} `field:"optional" json:"inferred" yaml:"inferred"`
}

