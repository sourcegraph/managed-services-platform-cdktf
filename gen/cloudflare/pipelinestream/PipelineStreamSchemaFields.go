package pipelinestream


type PipelineStreamSchemaFields struct {
	// Available values: "int32", "int64", "float32", "float64", "bool", "string", "binary", "timestamp", "json".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#type PipelineStream#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#metadata_key PipelineStream#metadata_key}.
	MetadataKey *string `field:"optional" json:"metadataKey" yaml:"metadataKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#name PipelineStream#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#required PipelineStream#required}.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#sql_name PipelineStream#sql_name}.
	SqlName *string `field:"optional" json:"sqlName" yaml:"sqlName"`
	// Available values: "second", "millisecond", "microsecond", "nanosecond".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pipeline_stream#unit PipelineStream#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

