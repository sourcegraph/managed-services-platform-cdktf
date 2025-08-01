package oracledatabaseautonomousdatabase


type OracleDatabaseAutonomousDatabasePropertiesCustomerContacts struct {
	// The email address used by Oracle to send notifications regarding databases and infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/oracle_database_autonomous_database#email OracleDatabaseAutonomousDatabase#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

