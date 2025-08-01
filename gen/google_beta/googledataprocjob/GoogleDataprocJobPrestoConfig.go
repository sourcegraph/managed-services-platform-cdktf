package googledataprocjob


type GoogleDataprocJobPrestoConfig struct {
	// Presto client tags to attach to this query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_job#client_tags GoogleDataprocJob#client_tags}
	ClientTags *[]*string `field:"optional" json:"clientTags" yaml:"clientTags"`
	// Whether to continue executing queries if a query fails.
	//
	// Setting to true can be useful when executing independent parallel queries. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_job#continue_on_failure GoogleDataprocJob#continue_on_failure}
	ContinueOnFailure interface{} `field:"optional" json:"continueOnFailure" yaml:"continueOnFailure"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_job#logging_config GoogleDataprocJob#logging_config}
	LoggingConfig *GoogleDataprocJobPrestoConfigLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The format in which query output will be displayed. See the Presto documentation for supported output formats.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_job#output_format GoogleDataprocJob#output_format}
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// A mapping of property names to values.
	//
	// Used to set Presto session properties Equivalent to using the --session flag in the Presto CLI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_job#properties GoogleDataprocJob#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The HCFS URI of the script that contains SQL queries. Conflicts with query_list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_job#query_file_uri GoogleDataprocJob#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// The list of SQL queries or statements to execute as part of the job. Conflicts with query_file_uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_job#query_list GoogleDataprocJob#query_list}
	QueryList *[]*string `field:"optional" json:"queryList" yaml:"queryList"`
}

