package googledatastreamstream


type GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTables struct {
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_datastream_stream#table GoogleDatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// oracle_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_datastream_stream#oracle_columns GoogleDatastreamStream#oracle_columns}
	OracleColumns interface{} `field:"optional" json:"oracleColumns" yaml:"oracleColumns"`
}

