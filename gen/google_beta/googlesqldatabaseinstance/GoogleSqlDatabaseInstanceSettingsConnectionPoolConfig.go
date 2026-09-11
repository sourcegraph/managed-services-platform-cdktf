package googlesqldatabaseinstance

type GoogleSqlDatabaseInstanceSettingsConnectionPoolConfig struct {
	// Whether Managed Connection Pool is enabled for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#connection_pooling_enabled GoogleSqlDatabaseInstance#connection_pooling_enabled}
	ConnectionPoolingEnabled any `field:"optional" json:"connectionPoolingEnabled" yaml:"connectionPoolingEnabled"`
	// flags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#flags GoogleSqlDatabaseInstance#flags}
	Flags any `field:"optional" json:"flags" yaml:"flags"`
}
