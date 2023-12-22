package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfig struct {
	// List of consumer projects that are allow-listed for PSC connections to this instance.
	//
	// This instance can be connected to with PSC from any network in these projects. Each consumer project in this list may be represented by a project number (numeric) or by a project id (alphanumeric).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_sql_database_instance#allowed_consumer_projects GoogleSqlDatabaseInstance#allowed_consumer_projects}
	AllowedConsumerProjects *[]*string `field:"optional" json:"allowedConsumerProjects" yaml:"allowedConsumerProjects"`
	// Whether PSC connectivity is enabled for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_sql_database_instance#psc_enabled GoogleSqlDatabaseInstance#psc_enabled}
	PscEnabled interface{} `field:"optional" json:"pscEnabled" yaml:"pscEnabled"`
}

