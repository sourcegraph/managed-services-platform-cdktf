package datastreamstream


type DatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjects struct {
	// postgresql_schemas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/datastream_stream#postgresql_schemas DatastreamStream#postgresql_schemas}
	PostgresqlSchemas interface{} `field:"required" json:"postgresqlSchemas" yaml:"postgresqlSchemas"`
}

