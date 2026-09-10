package pipelinestream


type PipelineStreamSchemaFormat struct {
	// Available values: "json", "parquet".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#type PipelineStream#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Available values: "uncompressed", "snappy", "gzip", "zstd", "lz4".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#compression PipelineStream#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Available values: "number", "string", "bytes".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#decimal_encoding PipelineStream#decimal_encoding}
	DecimalEncoding *string `field:"optional" json:"decimalEncoding" yaml:"decimalEncoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#row_group_bytes PipelineStream#row_group_bytes}.
	RowGroupBytes *float64 `field:"optional" json:"rowGroupBytes" yaml:"rowGroupBytes"`
	// Available values: "rfc3339", "unix_millis".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#timestamp_format PipelineStream#timestamp_format}
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#unstructured PipelineStream#unstructured}.
	Unstructured interface{} `field:"optional" json:"unstructured" yaml:"unstructured"`
}

