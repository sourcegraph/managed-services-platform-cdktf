package googlestoragetransferjob


type GoogleStorageTransferJobReplicationSpecTransferOptions struct {
	// Whether objects should be deleted from the source after they are transferred to the sink.
	//
	// Note that this option and delete_objects_unique_in_sink are mutually exclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_storage_transfer_job#delete_objects_from_source_after_transfer GoogleStorageTransferJob#delete_objects_from_source_after_transfer}
	DeleteObjectsFromSourceAfterTransfer interface{} `field:"optional" json:"deleteObjectsFromSourceAfterTransfer" yaml:"deleteObjectsFromSourceAfterTransfer"`
	// Whether objects that exist only in the sink should be deleted.
	//
	// Note that this option and delete_objects_from_source_after_transfer are mutually exclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_storage_transfer_job#delete_objects_unique_in_sink GoogleStorageTransferJob#delete_objects_unique_in_sink}
	DeleteObjectsUniqueInSink interface{} `field:"optional" json:"deleteObjectsUniqueInSink" yaml:"deleteObjectsUniqueInSink"`
	// metadata_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_storage_transfer_job#metadata_options GoogleStorageTransferJob#metadata_options}
	MetadataOptions *GoogleStorageTransferJobReplicationSpecTransferOptionsMetadataOptions `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// Whether overwriting objects that already exist in the sink is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_storage_transfer_job#overwrite_objects_already_existing_in_sink GoogleStorageTransferJob#overwrite_objects_already_existing_in_sink}
	OverwriteObjectsAlreadyExistingInSink interface{} `field:"optional" json:"overwriteObjectsAlreadyExistingInSink" yaml:"overwriteObjectsAlreadyExistingInSink"`
	// When to overwrite objects that already exist in the sink. If not set, overwrite behavior is determined by overwriteObjectsAlreadyExistingInSink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_storage_transfer_job#overwrite_when GoogleStorageTransferJob#overwrite_when}
	OverwriteWhen *string `field:"optional" json:"overwriteWhen" yaml:"overwriteWhen"`
}

