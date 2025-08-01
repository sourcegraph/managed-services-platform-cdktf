package storagebatchoperationsjob


type StorageBatchOperationsJobBucketListStruct struct {
	// buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_batch_operations_job#buckets StorageBatchOperationsJob#buckets}
	Buckets *StorageBatchOperationsJobBucketListBuckets `field:"required" json:"buckets" yaml:"buckets"`
}

