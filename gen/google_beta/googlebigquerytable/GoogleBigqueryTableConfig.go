package googlebigquerytable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigqueryTableConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The dataset ID to create the table in. Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#dataset_id GoogleBigqueryTable#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// A unique ID for the resource. Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#table_id GoogleBigqueryTable#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// biglake_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#biglake_configuration GoogleBigqueryTable#biglake_configuration}
	BiglakeConfiguration *GoogleBigqueryTableBiglakeConfiguration `field:"optional" json:"biglakeConfiguration" yaml:"biglakeConfiguration"`
	// Specifies column names to use for data clustering.
	//
	// Up to four top-level columns are allowed, and should be specified in descending priority order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#clustering GoogleBigqueryTable#clustering}
	Clustering *[]*string `field:"optional" json:"clustering" yaml:"clustering"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// When the field is set to true or unset in Terraform state, a terraform apply or terraform destroy that would delete the table will fail. When the field is set to false, deleting the table is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#deletion_protection GoogleBigqueryTable#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The field description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#description GoogleBigqueryTable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#encryption_configuration GoogleBigqueryTable#encryption_configuration}
	EncryptionConfiguration *GoogleBigqueryTableEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The time when this table expires, in milliseconds since the epoch.
	//
	// If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#expiration_time GoogleBigqueryTable#expiration_time}
	ExpirationTime *float64 `field:"optional" json:"expirationTime" yaml:"expirationTime"`
	// external_catalog_table_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#external_catalog_table_options GoogleBigqueryTable#external_catalog_table_options}
	ExternalCatalogTableOptions *GoogleBigqueryTableExternalCatalogTableOptions `field:"optional" json:"externalCatalogTableOptions" yaml:"externalCatalogTableOptions"`
	// external_data_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#external_data_configuration GoogleBigqueryTable#external_data_configuration}
	ExternalDataConfiguration *GoogleBigqueryTableExternalDataConfiguration `field:"optional" json:"externalDataConfiguration" yaml:"externalDataConfiguration"`
	// A descriptive name for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#friendly_name GoogleBigqueryTable#friendly_name}
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#id GoogleBigqueryTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether Terraform will prevent implicitly added columns in schema from showing diff.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#ignore_auto_generated_schema GoogleBigqueryTable#ignore_auto_generated_schema}
	IgnoreAutoGeneratedSchema interface{} `field:"optional" json:"ignoreAutoGeneratedSchema" yaml:"ignoreAutoGeneratedSchema"`
	// Mention which fields in schema are to be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#ignore_schema_changes GoogleBigqueryTable#ignore_schema_changes}
	IgnoreSchemaChanges *[]*string `field:"optional" json:"ignoreSchemaChanges" yaml:"ignoreSchemaChanges"`
	// A mapping of labels to assign to the resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// 				Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#labels GoogleBigqueryTable#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// materialized_view block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#materialized_view GoogleBigqueryTable#materialized_view}
	MaterializedView *GoogleBigqueryTableMaterializedView `field:"optional" json:"materializedView" yaml:"materializedView"`
	// The maximum staleness of data that could be returned when the table (or stale MV) is queried.
	//
	// Staleness encoded as a string encoding of [SQL IntervalValue type](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#max_staleness GoogleBigqueryTable#max_staleness}
	MaxStaleness *string `field:"optional" json:"maxStaleness" yaml:"maxStaleness"`
	// The ID of the project in which the resource belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#project GoogleBigqueryTable#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// range_partitioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#range_partitioning GoogleBigqueryTable#range_partitioning}
	RangePartitioning *GoogleBigqueryTableRangePartitioning `field:"optional" json:"rangePartitioning" yaml:"rangePartitioning"`
	// If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#require_partition_filter GoogleBigqueryTable#require_partition_filter}
	RequirePartitionFilter interface{} `field:"optional" json:"requirePartitionFilter" yaml:"requirePartitionFilter"`
	// The tags attached to this table.
	//
	// Tag keys are globally unique. Tag key is expected to be in the namespaced format, for example "123456789012/environment" where 123456789012 is the ID of the parent organization or project resource for this tag key. Tag value is expected to be the short name, for example "Production".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#resource_tags GoogleBigqueryTable#resource_tags}
	ResourceTags *map[string]*string `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// A JSON schema for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#schema GoogleBigqueryTable#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// schema_foreign_type_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#schema_foreign_type_info GoogleBigqueryTable#schema_foreign_type_info}
	SchemaForeignTypeInfo *GoogleBigqueryTableSchemaForeignTypeInfo `field:"optional" json:"schemaForeignTypeInfo" yaml:"schemaForeignTypeInfo"`
	// table_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#table_constraints GoogleBigqueryTable#table_constraints}
	TableConstraints *GoogleBigqueryTableTableConstraints `field:"optional" json:"tableConstraints" yaml:"tableConstraints"`
	// View sets the optional parameter "view": Specifies the view that determines which table information is returned.
	//
	// By default, basic table information and storage statistics (STORAGE_STATS) are returned. Possible values: TABLE_METADATA_VIEW_UNSPECIFIED, BASIC, STORAGE_STATS, FULL
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#table_metadata_view GoogleBigqueryTable#table_metadata_view}
	TableMetadataView *string `field:"optional" json:"tableMetadataView" yaml:"tableMetadataView"`
	// table_replication_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#table_replication_info GoogleBigqueryTable#table_replication_info}
	TableReplicationInfo *GoogleBigqueryTableTableReplicationInfo `field:"optional" json:"tableReplicationInfo" yaml:"tableReplicationInfo"`
	// time_partitioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#time_partitioning GoogleBigqueryTable#time_partitioning}
	TimePartitioning *GoogleBigqueryTableTimePartitioning `field:"optional" json:"timePartitioning" yaml:"timePartitioning"`
	// view block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_table#view GoogleBigqueryTable#view}
	View *GoogleBigqueryTableView `field:"optional" json:"view" yaml:"view"`
}

