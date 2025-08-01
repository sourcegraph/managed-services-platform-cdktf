package bigquerydatatransferconfig


type BigqueryDataTransferConfigEncryptionConfiguration struct {
	// The name of the KMS key used for encrypting BigQuery data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_data_transfer_config#kms_key_name BigqueryDataTransferConfig#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

