package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsInsightsConfig struct {
	// True if Query Insights feature is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#query_insights_enabled GoogleSqlDatabaseInstance#query_insights_enabled}
	QueryInsightsEnabled interface{} `field:"optional" json:"queryInsightsEnabled" yaml:"queryInsightsEnabled"`
	// Number of query execution plans captured by Insights per minute for all queries combined.
	//
	// Between 0 and 20. Default to 5. For Enterprise Plus instances, from 0 to 200.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#query_plans_per_minute GoogleSqlDatabaseInstance#query_plans_per_minute}
	QueryPlansPerMinute *float64 `field:"optional" json:"queryPlansPerMinute" yaml:"queryPlansPerMinute"`
	// Maximum query length stored in bytes.
	//
	// Between 256 and 4500. Default to 1024. For Enterprise Plus instances, from 1 to 1048576.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#query_string_length GoogleSqlDatabaseInstance#query_string_length}
	QueryStringLength *float64 `field:"optional" json:"queryStringLength" yaml:"queryStringLength"`
	// True if Query Insights will record application tags from query when enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#record_application_tags GoogleSqlDatabaseInstance#record_application_tags}
	RecordApplicationTags interface{} `field:"optional" json:"recordApplicationTags" yaml:"recordApplicationTags"`
	// True if Query Insights will record client address when enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_sql_database_instance#record_client_address GoogleSqlDatabaseInstance#record_client_address}
	RecordClientAddress interface{} `field:"optional" json:"recordClientAddress" yaml:"recordClientAddress"`
}

