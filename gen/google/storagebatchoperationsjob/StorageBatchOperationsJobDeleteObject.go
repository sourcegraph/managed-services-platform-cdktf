package storagebatchoperationsjob


type StorageBatchOperationsJobDeleteObject struct {
	// enable flag to permanently delete object and all object versions if versioning is enabled on bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_batch_operations_job#permanent_object_deletion_enabled StorageBatchOperationsJob#permanent_object_deletion_enabled}
	PermanentObjectDeletionEnabled interface{} `field:"required" json:"permanentObjectDeletionEnabled" yaml:"permanentObjectDeletionEnabled"`
}

