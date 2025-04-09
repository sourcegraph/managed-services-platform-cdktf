package googlestoragecontrolprojectintelligenceconfig


type GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets struct {
	// List of bucket id regexes to exclude in the storage intelligence plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_storage_control_project_intelligence_config#bucket_id_regexes GoogleStorageControlProjectIntelligenceConfig#bucket_id_regexes}
	BucketIdRegexes *[]*string `field:"required" json:"bucketIdRegexes" yaml:"bucketIdRegexes"`
}

