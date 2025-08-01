package datastreamstream


type DatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabases struct {
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/datastream_stream#database DatastreamStream#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// mysql_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/datastream_stream#mysql_tables DatastreamStream#mysql_tables}
	MysqlTables interface{} `field:"optional" json:"mysqlTables" yaml:"mysqlTables"`
}

