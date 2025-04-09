package googledatastreamstream


type GoogleDatastreamStreamSourceConfigSqlServerSourceConfig struct {
	// change_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_datastream_stream#change_tables GoogleDatastreamStream#change_tables}
	ChangeTables *GoogleDatastreamStreamSourceConfigSqlServerSourceConfigChangeTables `field:"optional" json:"changeTables" yaml:"changeTables"`
	// exclude_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_datastream_stream#exclude_objects GoogleDatastreamStream#exclude_objects}
	ExcludeObjects *GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjects `field:"optional" json:"excludeObjects" yaml:"excludeObjects"`
	// include_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_datastream_stream#include_objects GoogleDatastreamStream#include_objects}
	IncludeObjects *GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjects `field:"optional" json:"includeObjects" yaml:"includeObjects"`
	// Max concurrent backfill tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_datastream_stream#max_concurrent_backfill_tasks GoogleDatastreamStream#max_concurrent_backfill_tasks}
	MaxConcurrentBackfillTasks *float64 `field:"optional" json:"maxConcurrentBackfillTasks" yaml:"maxConcurrentBackfillTasks"`
	// Max concurrent CDC tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_datastream_stream#max_concurrent_cdc_tasks GoogleDatastreamStream#max_concurrent_cdc_tasks}
	MaxConcurrentCdcTasks *float64 `field:"optional" json:"maxConcurrentCdcTasks" yaml:"maxConcurrentCdcTasks"`
	// transaction_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_datastream_stream#transaction_logs GoogleDatastreamStream#transaction_logs}
	TransactionLogs *GoogleDatastreamStreamSourceConfigSqlServerSourceConfigTransactionLogs `field:"optional" json:"transactionLogs" yaml:"transactionLogs"`
}

