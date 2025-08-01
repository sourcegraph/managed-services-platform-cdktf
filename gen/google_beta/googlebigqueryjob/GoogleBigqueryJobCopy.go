package googlebigqueryjob


type GoogleBigqueryJobCopy struct {
	// source_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_job#source_tables GoogleBigqueryJob#source_tables}
	SourceTables interface{} `field:"required" json:"sourceTables" yaml:"sourceTables"`
	// Specifies whether the job is allowed to create new tables.
	//
	// The following values are supported:
	// CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the table.
	// CREATE_NEVER: The table must already exist. If it does not, a 'notFound' error is returned in the job result.
	// Creation, truncation and append actions occur as one atomic update upon job completion Default value: "CREATE_IF_NEEDED" Possible values: ["CREATE_IF_NEEDED", "CREATE_NEVER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_job#create_disposition GoogleBigqueryJob#create_disposition}
	CreateDisposition *string `field:"optional" json:"createDisposition" yaml:"createDisposition"`
	// destination_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_job#destination_encryption_configuration GoogleBigqueryJob#destination_encryption_configuration}
	DestinationEncryptionConfiguration *GoogleBigqueryJobCopyDestinationEncryptionConfiguration `field:"optional" json:"destinationEncryptionConfiguration" yaml:"destinationEncryptionConfiguration"`
	// destination_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_job#destination_table GoogleBigqueryJob#destination_table}
	DestinationTable *GoogleBigqueryJobCopyDestinationTable `field:"optional" json:"destinationTable" yaml:"destinationTable"`
	// Specifies the action that occurs if the destination table already exists.
	//
	// The following values are supported:
	// WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the table data and uses the schema from the query result.
	// WRITE_APPEND: If the table already exists, BigQuery appends the data to the table.
	// WRITE_EMPTY: If the table already exists and contains data, a 'duplicate' error is returned in the job result.
	// Each action is atomic and only occurs if BigQuery is able to complete the job successfully.
	// Creation, truncation and append actions occur as one atomic update upon job completion. Default value: "WRITE_EMPTY" Possible values: ["WRITE_TRUNCATE", "WRITE_APPEND", "WRITE_EMPTY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_job#write_disposition GoogleBigqueryJob#write_disposition}
	WriteDisposition *string `field:"optional" json:"writeDisposition" yaml:"writeDisposition"`
}

