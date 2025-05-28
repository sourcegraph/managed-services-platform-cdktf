package googledatastreamstream


type GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemas struct {
	// Schema name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_datastream_stream#schema GoogleDatastreamStream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// oracle_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_datastream_stream#oracle_tables GoogleDatastreamStream#oracle_tables}
	OracleTables interface{} `field:"optional" json:"oracleTables" yaml:"oracleTables"`
}

