package googledataplexzone


type GoogleDataplexZoneDiscoverySpecCsvOptions struct {
	// Optional. The delimiter being used to separate values. This defaults to ','.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataplex_zone#delimiter GoogleDataplexZone#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Optional.
	//
	// Whether to disable the inference of data type for CSV data. If true, all columns will be registered as strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataplex_zone#disable_type_inference GoogleDataplexZone#disable_type_inference}
	DisableTypeInference interface{} `field:"optional" json:"disableTypeInference" yaml:"disableTypeInference"`
	// Optional. The character encoding of the data. The default is UTF-8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataplex_zone#encoding GoogleDataplexZone#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Optional. The number of rows to interpret as header rows that should be skipped when reading data rows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataplex_zone#header_rows GoogleDataplexZone#header_rows}
	HeaderRows *float64 `field:"optional" json:"headerRows" yaml:"headerRows"`
}

