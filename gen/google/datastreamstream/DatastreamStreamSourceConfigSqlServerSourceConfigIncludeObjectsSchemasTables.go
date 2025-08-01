package datastreamstream


type DatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTables struct {
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/datastream_stream#table DatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/datastream_stream#columns DatastreamStream#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
}

