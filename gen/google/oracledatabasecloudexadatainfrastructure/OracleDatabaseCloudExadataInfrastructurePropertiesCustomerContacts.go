package oracledatabasecloudexadatainfrastructure


type OracleDatabaseCloudExadataInfrastructurePropertiesCustomerContacts struct {
	// The email address used by Oracle to send notifications regarding databases and infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/oracle_database_cloud_exadata_infrastructure#email OracleDatabaseCloudExadataInfrastructure#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

