package sqldatabaseinstance


type SqlDatabaseInstanceSettingsConnectionPoolConfigFlags struct {
	// Name of the flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#name SqlDatabaseInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#value SqlDatabaseInstance#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

