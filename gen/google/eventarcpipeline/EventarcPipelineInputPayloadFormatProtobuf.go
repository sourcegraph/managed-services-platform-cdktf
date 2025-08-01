package eventarcpipeline


type EventarcPipelineInputPayloadFormatProtobuf struct {
	// The entire schema definition is stored in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/eventarc_pipeline#schema_definition EventarcPipeline#schema_definition}
	SchemaDefinition *string `field:"optional" json:"schemaDefinition" yaml:"schemaDefinition"`
}

