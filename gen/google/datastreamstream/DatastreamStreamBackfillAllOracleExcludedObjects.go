package datastreamstream


type DatastreamStreamBackfillAllOracleExcludedObjects struct {
	// oracle_schemas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/datastream_stream#oracle_schemas DatastreamStream#oracle_schemas}
	OracleSchemas interface{} `field:"required" json:"oracleSchemas" yaml:"oracleSchemas"`
}

