package databasemigrationserviceconnectionprofile


type DatabaseMigrationServiceConnectionProfilePostgresqlSsl struct {
	// Input only.
	//
	// The x509 PEM-encoded certificate of the CA that signed the source database server's certificate.
	// The replica will use this certificate to verify it's connecting to the right host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/database_migration_service_connection_profile#ca_certificate DatabaseMigrationServiceConnectionProfile#ca_certificate}
	CaCertificate *string `field:"optional" json:"caCertificate" yaml:"caCertificate"`
	// Input only.
	//
	// The x509 PEM-encoded certificate that will be used by the replica to authenticate against the source database server.
	// If this field is used then the 'clientKey' field is mandatory
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/database_migration_service_connection_profile#client_certificate DatabaseMigrationServiceConnectionProfile#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// Input only.
	//
	// The unencrypted PKCS#1 or PKCS#8 PEM-encoded private key associated with the Client Certificate.
	// If this field is used then the 'clientCertificate' field is mandatory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/database_migration_service_connection_profile#client_key DatabaseMigrationServiceConnectionProfile#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// The current connection profile state. Possible values: ["SERVER_ONLY", "SERVER_CLIENT", "REQUIRED", "NONE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/database_migration_service_connection_profile#type DatabaseMigrationServiceConnectionProfile#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

