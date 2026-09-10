package pipelinesink


type PipelineSinkSchemaFields struct {
	// Available values: "int32", "int64", "float32", "float64", "bool", "string", "binary", "timestamp", "json".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#type PipelineSink#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#metadata_key PipelineSink#metadata_key}.
	MetadataKey *string `field:"optional" json:"metadataKey" yaml:"metadataKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#name PipelineSink#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#required PipelineSink#required}.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#sql_name PipelineSink#sql_name}.
	SqlName *string `field:"optional" json:"sqlName" yaml:"sqlName"`
	// Available values: "second", "millisecond", "microsecond", "nanosecond".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_sink#unit PipelineSink#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

