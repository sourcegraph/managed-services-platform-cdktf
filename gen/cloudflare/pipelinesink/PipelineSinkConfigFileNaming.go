package pipelinesink


type PipelineSinkConfigFileNaming struct {
	// The prefix to use in file name. i.e prefix-<uuid>.parquet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#prefix PipelineSink#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Filename generation strategy. Available values: "serial", "uuid", "uuid_v7", "ulid".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#strategy PipelineSink#strategy}
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// This will overwrite the default file suffix. i.e .parquet, use with caution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#suffix PipelineSink#suffix}
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

