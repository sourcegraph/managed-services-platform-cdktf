package googleloggingfolderbucketconfig


type GoogleLoggingFolderBucketConfigIndexConfigs struct {
	// The LogEntry field path to index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_logging_folder_bucket_config#field_path GoogleLoggingFolderBucketConfig#field_path}
	FieldPath *string `field:"required" json:"fieldPath" yaml:"fieldPath"`
	// The type of data in this index Note that some paths are automatically indexed, and other paths are not eligible for indexing.
	//
	// See [indexing documentation]( https://cloud.google.com/logging/docs/view/advanced-queries#indexed-fields) for details.
	// For example: jsonPayload.request.status
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_logging_folder_bucket_config#type GoogleLoggingFolderBucketConfig#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

