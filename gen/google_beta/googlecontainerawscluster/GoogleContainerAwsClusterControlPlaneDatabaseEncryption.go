package googlecontainerawscluster


type GoogleContainerAwsClusterControlPlaneDatabaseEncryption struct {
	// The ARN of the AWS KMS key used to encrypt cluster secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_container_aws_cluster#kms_key_arn GoogleContainerAwsCluster#kms_key_arn}
	KmsKeyArn *string `field:"required" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

