package googlemanagedkafkaacl


type GoogleManagedKafkaAclAclEntries struct {
	// The operation type.
	//
	// Allowed values are (case insensitive): ALL, READ,
	// WRITE, CREATE, DELETE, ALTER, DESCRIBE, CLUSTER_ACTION, DESCRIBE_CONFIGS,
	// ALTER_CONFIGS, and IDEMPOTENT_WRITE. See https://kafka.apache.org/documentation/#operations_resources_and_protocols
	// for valid combinations of resource_type and operation for different Kafka API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_managed_kafka_acl#operation GoogleManagedKafkaAcl#operation}
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// The principal.
	//
	// Specified as Google Cloud account, with the Kafka StandardAuthorizer prefix User:". For example: "User:test-kafka-client@test-project.iam.gserviceaccount.com". Can be the wildcard "User:*" to refer to all users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_managed_kafka_acl#principal GoogleManagedKafkaAcl#principal}
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// The host. Must be set to "*" for Managed Service for Apache Kafka.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_managed_kafka_acl#host GoogleManagedKafkaAcl#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The permission type. Accepted values are (case insensitive): ALLOW, DENY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_managed_kafka_acl#permission_type GoogleManagedKafkaAcl#permission_type}
	PermissionType *string `field:"optional" json:"permissionType" yaml:"permissionType"`
}

