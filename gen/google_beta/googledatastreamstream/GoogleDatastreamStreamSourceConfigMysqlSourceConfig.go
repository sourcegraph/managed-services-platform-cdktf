package googledatastreamstream


type GoogleDatastreamStreamSourceConfigMysqlSourceConfig struct {
	// binary_log_position block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_datastream_stream#binary_log_position GoogleDatastreamStream#binary_log_position}
	BinaryLogPosition *GoogleDatastreamStreamSourceConfigMysqlSourceConfigBinaryLogPosition `field:"optional" json:"binaryLogPosition" yaml:"binaryLogPosition"`
	// exclude_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_datastream_stream#exclude_objects GoogleDatastreamStream#exclude_objects}
	ExcludeObjects *GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects `field:"optional" json:"excludeObjects" yaml:"excludeObjects"`
	// gtid block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_datastream_stream#gtid GoogleDatastreamStream#gtid}
	Gtid *GoogleDatastreamStreamSourceConfigMysqlSourceConfigGtid `field:"optional" json:"gtid" yaml:"gtid"`
	// include_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_datastream_stream#include_objects GoogleDatastreamStream#include_objects}
	IncludeObjects *GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjects `field:"optional" json:"includeObjects" yaml:"includeObjects"`
	// Maximum number of concurrent backfill tasks.
	//
	// The number should be non negative.
	// If not set (or set to 0), the system's default value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_datastream_stream#max_concurrent_backfill_tasks GoogleDatastreamStream#max_concurrent_backfill_tasks}
	MaxConcurrentBackfillTasks *float64 `field:"optional" json:"maxConcurrentBackfillTasks" yaml:"maxConcurrentBackfillTasks"`
	// Maximum number of concurrent CDC tasks.
	//
	// The number should be non negative.
	// If not set (or set to 0), the system's default value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_datastream_stream#max_concurrent_cdc_tasks GoogleDatastreamStream#max_concurrent_cdc_tasks}
	MaxConcurrentCdcTasks *float64 `field:"optional" json:"maxConcurrentCdcTasks" yaml:"maxConcurrentCdcTasks"`
}

