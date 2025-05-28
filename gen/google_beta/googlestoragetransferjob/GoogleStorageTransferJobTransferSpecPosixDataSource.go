package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecPosixDataSource struct {
	// Root directory path to the filesystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_storage_transfer_job#root_directory GoogleStorageTransferJob#root_directory}
	RootDirectory *string `field:"required" json:"rootDirectory" yaml:"rootDirectory"`
}

