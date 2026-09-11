package pipelinesink


type PipelineSinkFormat struct {
	// Available values: "json", "parquet".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#type PipelineSink#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Available values: "uncompressed", "snappy", "gzip", "zstd", "lz4".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#compression PipelineSink#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Available values: "number", "string", "bytes".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#decimal_encoding PipelineSink#decimal_encoding}
	DecimalEncoding *string `field:"optional" json:"decimalEncoding" yaml:"decimalEncoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#row_group_bytes PipelineSink#row_group_bytes}.
	RowGroupBytes *float64 `field:"optional" json:"rowGroupBytes" yaml:"rowGroupBytes"`
	// Available values: "rfc3339", "unix_millis".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#timestamp_format PipelineSink#timestamp_format}
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#unstructured PipelineSink#unstructured}.
	Unstructured interface{} `field:"optional" json:"unstructured" yaml:"unstructured"`
}

