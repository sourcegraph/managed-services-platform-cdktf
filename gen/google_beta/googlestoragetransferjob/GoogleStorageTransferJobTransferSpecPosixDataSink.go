package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecPosixDataSink struct {
	// Root directory path to the filesystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_storage_transfer_job#root_directory GoogleStorageTransferJob#root_directory}
	RootDirectory *string `field:"required" json:"rootDirectory" yaml:"rootDirectory"`
}
