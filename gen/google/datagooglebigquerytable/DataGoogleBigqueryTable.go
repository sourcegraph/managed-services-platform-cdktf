package datagooglebigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/datagooglebigquerytable/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/data-sources/bigquery_table google_bigquery_table}.
type DataGoogleBigqueryTable interface {
	cdktf.TerraformDataSource
	BiglakeConfiguration() DataGoogleBigqueryTableBiglakeConfigurationList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Clustering() *[]*string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *float64
	DatasetId() *string
	SetDatasetId(val *string)
	DatasetIdInput() *string
	DeletionProtection() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	EffectiveLabels() cdktf.StringMap
	EncryptionConfiguration() DataGoogleBigqueryTableEncryptionConfigurationList
	Etag() *string
	ExpirationTime() *float64
	ExternalCatalogTableOptions() DataGoogleBigqueryTableExternalCatalogTableOptionsList
	ExternalDataConfiguration() DataGoogleBigqueryTableExternalDataConfigurationList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	FriendlyName() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreAutoGeneratedSchema() cdktf.IResolvable
	IgnoreSchemaChanges() *[]*string
	Labels() cdktf.StringMap
	LastModifiedTime() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	MaterializedView() DataGoogleBigqueryTableMaterializedViewList
	MaxStaleness() *string
	// The tree node.
	Node() constructs.Node
	NumBytes() *float64
	NumLongTermBytes() *float64
	NumRows() *float64
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	RangePartitioning() DataGoogleBigqueryTableRangePartitioningList
	// Experimental.
	RawOverrides() interface{}
	RequirePartitionFilter() cdktf.IResolvable
	ResourceTags() cdktf.StringMap
	Schema() *string
	SchemaForeignTypeInfo() DataGoogleBigqueryTableSchemaForeignTypeInfoList
	SelfLink() *string
	TableConstraints() DataGoogleBigqueryTableTableConstraintsList
	TableId() *string
	SetTableId(val *string)
	TableIdInput() *string
	TableMetadataView() *string
	TableReplicationInfo() DataGoogleBigqueryTableTableReplicationInfoList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimePartitioning() DataGoogleBigqueryTableTimePartitioningList
	Type() *string
	View() DataGoogleBigqueryTableViewList
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataGoogleBigqueryTable
type jsiiProxy_DataGoogleBigqueryTable struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGoogleBigqueryTable) BiglakeConfiguration() DataGoogleBigqueryTableBiglakeConfigurationList {
	var returns DataGoogleBigqueryTableBiglakeConfigurationList
	_jsii_.Get(
		j,
		"biglakeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Clustering() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clustering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) CreationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) DatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) DatasetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) DeletionProtection() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) EncryptionConfiguration() DataGoogleBigqueryTableEncryptionConfigurationList {
	var returns DataGoogleBigqueryTableEncryptionConfigurationList
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) ExpirationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) ExternalCatalogTableOptions() DataGoogleBigqueryTableExternalCatalogTableOptionsList {
	var returns DataGoogleBigqueryTableExternalCatalogTableOptionsList
	_jsii_.Get(
		j,
		"externalCatalogTableOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) ExternalDataConfiguration() DataGoogleBigqueryTableExternalDataConfigurationList {
	var returns DataGoogleBigqueryTableExternalDataConfigurationList
	_jsii_.Get(
		j,
		"externalDataConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) FriendlyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) IgnoreAutoGeneratedSchema() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ignoreAutoGeneratedSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) IgnoreSchemaChanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreSchemaChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Labels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) LastModifiedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) MaterializedView() DataGoogleBigqueryTableMaterializedViewList {
	var returns DataGoogleBigqueryTableMaterializedViewList
	_jsii_.Get(
		j,
		"materializedView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) MaxStaleness() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxStaleness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) NumBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) NumLongTermBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numLongTermBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) NumRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) RangePartitioning() DataGoogleBigqueryTableRangePartitioningList {
	var returns DataGoogleBigqueryTableRangePartitioningList
	_jsii_.Get(
		j,
		"rangePartitioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) RequirePartitionFilter() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requirePartitionFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) ResourceTags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) SchemaForeignTypeInfo() DataGoogleBigqueryTableSchemaForeignTypeInfoList {
	var returns DataGoogleBigqueryTableSchemaForeignTypeInfoList
	_jsii_.Get(
		j,
		"schemaForeignTypeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) TableConstraints() DataGoogleBigqueryTableTableConstraintsList {
	var returns DataGoogleBigqueryTableTableConstraintsList
	_jsii_.Get(
		j,
		"tableConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) TableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) TableIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) TableMetadataView() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableMetadataView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) TableReplicationInfo() DataGoogleBigqueryTableTableReplicationInfoList {
	var returns DataGoogleBigqueryTableTableReplicationInfoList
	_jsii_.Get(
		j,
		"tableReplicationInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) TimePartitioning() DataGoogleBigqueryTableTimePartitioningList {
	var returns DataGoogleBigqueryTableTimePartitioningList
	_jsii_.Get(
		j,
		"timePartitioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTable) View() DataGoogleBigqueryTableViewList {
	var returns DataGoogleBigqueryTableViewList
	_jsii_.Get(
		j,
		"view",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/data-sources/bigquery_table google_bigquery_table} Data Source.
func NewDataGoogleBigqueryTable(scope constructs.Construct, id *string, config *DataGoogleBigqueryTableConfig) DataGoogleBigqueryTable {
	_init_.Initialize()

	if err := validateNewDataGoogleBigqueryTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleBigqueryTable{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleBigqueryTable.DataGoogleBigqueryTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/data-sources/bigquery_table google_bigquery_table} Data Source.
func NewDataGoogleBigqueryTable_Override(d DataGoogleBigqueryTable, scope constructs.Construct, id *string, config *DataGoogleBigqueryTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleBigqueryTable.DataGoogleBigqueryTable",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTable)SetDatasetId(val *string) {
	if err := j.validateSetDatasetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetId",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTable)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTable)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTable)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTable)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTable)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTable)SetTableId(val *string) {
	if err := j.validateSetTableIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableId",
		val,
	)
}

// Generates CDKTF code for importing a DataGoogleBigqueryTable resource upon running "cdktf plan <stack-name>".
func DataGoogleBigqueryTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataGoogleBigqueryTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleBigqueryTable.DataGoogleBigqueryTable",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataGoogleBigqueryTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleBigqueryTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleBigqueryTable.DataGoogleBigqueryTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleBigqueryTable_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleBigqueryTable_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleBigqueryTable.DataGoogleBigqueryTable",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleBigqueryTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleBigqueryTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleBigqueryTable.DataGoogleBigqueryTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGoogleBigqueryTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dataGoogleBigqueryTable.DataGoogleBigqueryTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGoogleBigqueryTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGoogleBigqueryTable) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleBigqueryTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleBigqueryTable) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleBigqueryTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

