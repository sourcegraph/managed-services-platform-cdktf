package googledatastreamstream


type GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTables struct {
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_stream#table GoogleDatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// postgresql_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_stream#postgresql_columns GoogleDatastreamStream#postgresql_columns}
	PostgresqlColumns interface{} `field:"optional" json:"postgresqlColumns" yaml:"postgresqlColumns"`
}

