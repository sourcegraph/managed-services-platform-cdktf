package integrationconnectorsconnection


type IntegrationConnectorsConnectionEventingConfigAdditionalVariableEncryptionKeyValue struct {
	// The [KMS key name] with which the content of the Operation is encrypted.
	//
	// The expected
	// format: projects/*\/locations/*\/keyRings/*\/cryptoKeys/*.
	// Will be empty string if google managed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/integration_connectors_connection#kms_key_name IntegrationConnectorsConnection#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Type of Encryption Key Possible values: ["GOOGLE_MANAGED", "CUSTOMER_MANAGED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/integration_connectors_connection#type IntegrationConnectorsConnection#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

