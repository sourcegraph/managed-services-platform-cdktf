package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions struct {
	// table_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_job_trigger#table_reference GoogleDataLossPreventionJobTrigger#table_reference}
	TableReference *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference `field:"required" json:"tableReference" yaml:"tableReference"`
	// excluded_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_job_trigger#excluded_fields GoogleDataLossPreventionJobTrigger#excluded_fields}
	ExcludedFields interface{} `field:"optional" json:"excludedFields" yaml:"excludedFields"`
	// identifying_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_job_trigger#identifying_fields GoogleDataLossPreventionJobTrigger#identifying_fields}
	IdentifyingFields interface{} `field:"optional" json:"identifyingFields" yaml:"identifyingFields"`
	// included_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_job_trigger#included_fields GoogleDataLossPreventionJobTrigger#included_fields}
	IncludedFields interface{} `field:"optional" json:"includedFields" yaml:"includedFields"`
	// Max number of rows to scan.
	//
	// If the table has more rows than this value, the rest of the rows are omitted.
	// If not set, or if set to 0, all rows will be scanned. Only one of rowsLimit and rowsLimitPercent can be
	// specified. Cannot be used in conjunction with TimespanConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_job_trigger#rows_limit GoogleDataLossPreventionJobTrigger#rows_limit}
	RowsLimit *float64 `field:"optional" json:"rowsLimit" yaml:"rowsLimit"`
	// Max percentage of rows to scan.
	//
	// The rest are omitted. The number of rows scanned is rounded down.
	// Must be between 0 and 100, inclusively. Both 0 and 100 means no limit. Defaults to 0. Only one of
	// rowsLimit and rowsLimitPercent can be specified. Cannot be used in conjunction with TimespanConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_job_trigger#rows_limit_percent GoogleDataLossPreventionJobTrigger#rows_limit_percent}
	RowsLimitPercent *float64 `field:"optional" json:"rowsLimitPercent" yaml:"rowsLimitPercent"`
	// How to sample rows if not all rows are scanned.
	//
	// Meaningful only when used in conjunction with either
	// rowsLimit or rowsLimitPercent. If not specified, rows are scanned in the order BigQuery reads them.
	// If TimespanConfig is set, set this to an empty string to avoid using the default value. Default value: "TOP" Possible values: ["TOP", "RANDOM_START"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_job_trigger#sample_method GoogleDataLossPreventionJobTrigger#sample_method}
	SampleMethod *string `field:"optional" json:"sampleMethod" yaml:"sampleMethod"`
}

