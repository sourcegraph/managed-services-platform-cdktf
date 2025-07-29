package googledatastreamstream


type GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemas struct {
	// Schema name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_stream#schema GoogleDatastreamStream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_stream#tables GoogleDatastreamStream#tables}
	Tables interface{} `field:"optional" json:"tables" yaml:"tables"`
}

