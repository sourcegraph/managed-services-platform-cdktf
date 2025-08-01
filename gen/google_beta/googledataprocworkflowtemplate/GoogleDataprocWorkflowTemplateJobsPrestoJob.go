package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplateJobsPrestoJob struct {
	// Optional. Presto client tags to attach to this query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_workflow_template#client_tags GoogleDataprocWorkflowTemplate#client_tags}
	ClientTags *[]*string `field:"optional" json:"clientTags" yaml:"clientTags"`
	// Optional.
	//
	// Whether to continue executing queries if a query fails. The default value is `false`. Setting to `true` can be useful when executing independent parallel queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_workflow_template#continue_on_failure GoogleDataprocWorkflowTemplate#continue_on_failure}
	ContinueOnFailure interface{} `field:"optional" json:"continueOnFailure" yaml:"continueOnFailure"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_workflow_template#logging_config GoogleDataprocWorkflowTemplate#logging_config}
	LoggingConfig *GoogleDataprocWorkflowTemplateJobsPrestoJobLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// Optional. The format in which query output will be displayed. See the Presto documentation for supported output formats.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_workflow_template#output_format GoogleDataprocWorkflowTemplate#output_format}
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Optional.
	//
	// A mapping of property names to values. Used to set Presto [session properties](https://prestodb.io/docs/current/sql/set-session.html) Equivalent to using the --session flag in the Presto CLI
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_workflow_template#properties GoogleDataprocWorkflowTemplate#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The HCFS URI of the script that contains SQL queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_workflow_template#query_file_uri GoogleDataprocWorkflowTemplate#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// query_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_workflow_template#query_list GoogleDataprocWorkflowTemplate#query_list}
	QueryList *GoogleDataprocWorkflowTemplateJobsPrestoJobQueryListStruct `field:"optional" json:"queryList" yaml:"queryList"`
}

