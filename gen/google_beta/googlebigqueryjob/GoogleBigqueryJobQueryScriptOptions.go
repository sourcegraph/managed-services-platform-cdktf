package googlebigqueryjob


type GoogleBigqueryJobQueryScriptOptions struct {
	// Determines which statement in the script represents the "key result", used to populate the schema and query results of the script job.
	//
	// Possible values: ["LAST", "FIRST_SELECT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_job#key_result_statement GoogleBigqueryJob#key_result_statement}
	KeyResultStatement *string `field:"optional" json:"keyResultStatement" yaml:"keyResultStatement"`
	// Limit on the number of bytes billed per statement. Exceeding this budget results in an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_job#statement_byte_budget GoogleBigqueryJob#statement_byte_budget}
	StatementByteBudget *string `field:"optional" json:"statementByteBudget" yaml:"statementByteBudget"`
	// Timeout period for each statement in a script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_job#statement_timeout_ms GoogleBigqueryJob#statement_timeout_ms}
	StatementTimeoutMs *string `field:"optional" json:"statementTimeoutMs" yaml:"statementTimeoutMs"`
}

