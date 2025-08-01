package databasemigrationservicemigrationjob


type DatabaseMigrationServiceMigrationJobVpcPeeringConnectivity struct {
	// The name of the VPC network to peer with the Cloud SQL private network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/database_migration_service_migration_job#vpc DatabaseMigrationServiceMigrationJob#vpc}
	Vpc *string `field:"optional" json:"vpc" yaml:"vpc"`
}

