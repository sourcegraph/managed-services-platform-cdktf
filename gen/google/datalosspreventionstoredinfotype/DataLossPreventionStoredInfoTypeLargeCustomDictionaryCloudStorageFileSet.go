package datalosspreventionstoredinfotype


type DataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet struct {
	// The url, in the format 'gs://<bucket>/<path>'. Trailing wildcard in the path is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/data_loss_prevention_stored_info_type#url DataLossPreventionStoredInfoType#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

