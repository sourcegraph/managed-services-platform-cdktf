package googlebigqueryjob


type GoogleBigqueryJobLoadDestinationEncryptionConfiguration struct {
	// Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table.
	//
	// The BigQuery Service Account associated with your project requires access to this encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_bigquery_job#kms_key_name GoogleBigqueryJob#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

