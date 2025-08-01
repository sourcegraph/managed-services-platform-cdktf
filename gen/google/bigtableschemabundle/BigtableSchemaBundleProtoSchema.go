package bigtableschemabundle


type BigtableSchemaBundleProtoSchema struct {
	// Base64 encoded content of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigtable_schema_bundle#proto_descriptors BigtableSchemaBundle#proto_descriptors}
	ProtoDescriptors *string `field:"required" json:"protoDescriptors" yaml:"protoDescriptors"`
}

