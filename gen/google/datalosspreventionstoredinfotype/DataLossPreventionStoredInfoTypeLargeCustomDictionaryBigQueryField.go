package datalosspreventionstoredinfotype


type DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField struct {
	// field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/data_loss_prevention_stored_info_type#field DataLossPreventionStoredInfoType#field}
	Field *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField `field:"required" json:"field" yaml:"field"`
	// table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/data_loss_prevention_stored_info_type#table DataLossPreventionStoredInfoType#table}
	Table *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable `field:"required" json:"table" yaml:"table"`
}

