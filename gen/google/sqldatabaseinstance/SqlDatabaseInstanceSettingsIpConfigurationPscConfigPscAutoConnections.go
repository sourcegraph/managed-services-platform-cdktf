package sqldatabaseinstance


type SqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnections struct {
	// The consumer network of this consumer endpoint.
	//
	// This must be a resource path that includes both the host project and the network name. The consumer host project of this network might be different from the consumer service project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#consumer_network SqlDatabaseInstance#consumer_network}
	ConsumerNetwork *string `field:"required" json:"consumerNetwork" yaml:"consumerNetwork"`
	// The project ID of consumer service project of this consumer endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#consumer_service_project_id SqlDatabaseInstance#consumer_service_project_id}
	ConsumerServiceProjectId *string `field:"optional" json:"consumerServiceProjectId" yaml:"consumerServiceProjectId"`
}

