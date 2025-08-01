package storagetransferjob


type StorageTransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials struct {
	// Azure shared access signature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_transfer_job#sas_token StorageTransferJob#sas_token}
	SasToken *string `field:"required" json:"sasToken" yaml:"sasToken"`
}

