package googledatabasemigrationserviceconnectionprofile


type GoogleDatabaseMigrationServiceConnectionProfileAlloydbSettings struct {
	// initial_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_database_migration_service_connection_profile#initial_user GoogleDatabaseMigrationServiceConnectionProfile#initial_user}
	InitialUser *GoogleDatabaseMigrationServiceConnectionProfileAlloydbSettingsInitialUser `field:"required" json:"initialUser" yaml:"initialUser"`
	// Required.
	//
	// The resource link for the VPC network in which cluster resources are created and from which they are accessible via Private IP. The network must belong to the same project as the cluster.
	// It is specified in the form: 'projects/{project_number}/global/networks/{network_id}'. This is required to create a cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_database_migration_service_connection_profile#vpc_network GoogleDatabaseMigrationServiceConnectionProfile#vpc_network}
	VpcNetwork *string `field:"required" json:"vpcNetwork" yaml:"vpcNetwork"`
	// Labels for the AlloyDB cluster created by DMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_database_migration_service_connection_profile#labels GoogleDatabaseMigrationServiceConnectionProfile#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// primary_instance_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_database_migration_service_connection_profile#primary_instance_settings GoogleDatabaseMigrationServiceConnectionProfile#primary_instance_settings}
	PrimaryInstanceSettings *GoogleDatabaseMigrationServiceConnectionProfileAlloydbSettingsPrimaryInstanceSettings `field:"optional" json:"primaryInstanceSettings" yaml:"primaryInstanceSettings"`
}

