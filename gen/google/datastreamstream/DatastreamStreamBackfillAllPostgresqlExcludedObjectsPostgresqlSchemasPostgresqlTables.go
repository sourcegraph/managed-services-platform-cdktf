package datastreamstream


type DatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTables struct {
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/datastream_stream#table DatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// postgresql_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/datastream_stream#postgresql_columns DatastreamStream#postgresql_columns}
	PostgresqlColumns interface{} `field:"optional" json:"postgresqlColumns" yaml:"postgresqlColumns"`
}

