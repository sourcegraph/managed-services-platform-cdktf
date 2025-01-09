package googledatabasemigrationserviceprivateconnection


type GoogleDatabaseMigrationServicePrivateConnectionVpcPeeringConfig struct {
	// A free subnet for peering. (CIDR of /29).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_database_migration_service_private_connection#subnet GoogleDatabaseMigrationServicePrivateConnection#subnet}
	Subnet *string `field:"required" json:"subnet" yaml:"subnet"`
	// Fully qualified name of the VPC that Database Migration Service will peer to. Format: projects/{project}/global/{networks}/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_database_migration_service_private_connection#vpc_name GoogleDatabaseMigrationServicePrivateConnection#vpc_name}
	VpcName *string `field:"required" json:"vpcName" yaml:"vpcName"`
}

