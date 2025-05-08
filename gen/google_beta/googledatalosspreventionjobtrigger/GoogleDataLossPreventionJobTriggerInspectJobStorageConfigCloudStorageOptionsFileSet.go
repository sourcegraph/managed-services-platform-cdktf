package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet struct {
	// regex_file_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_data_loss_prevention_job_trigger#regex_file_set GoogleDataLossPreventionJobTrigger#regex_file_set}
	RegexFileSet *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetRegexFileSet `field:"optional" json:"regexFileSet" yaml:"regexFileSet"`
	// The Cloud Storage url of the file(s) to scan, in the format 'gs://<bucket>/<path>'. Trailing wildcard in the path is allowed.
	//
	// If the url ends in a trailing slash, the bucket or directory represented by the url will be scanned
	// non-recursively (content in sub-directories will not be scanned). This means that 'gs://mybucket/' is
	// equivalent to 'gs://mybucket/*', and 'gs://mybucket/directory/' is equivalent to 'gs://mybucket/directory/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_data_loss_prevention_job_trigger#url GoogleDataLossPreventionJobTrigger#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

