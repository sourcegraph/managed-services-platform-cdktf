package datacloudflareaisearchinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/datacloudflareaisearchinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instance cloudflare_ai_search_instance}.
type DataCloudflareAiSearchInstance interface {
	cdktf.TerraformDataSource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AiGatewayId() *string
	AisearchModel() *string
	Cache() cdktf.IResolvable
	CacheThreshold() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChunkOverlap() *float64
	ChunkSize() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	CreatedBy() *string
	CustomMetadata() DataCloudflareAiSearchInstanceCustomMetadataList
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EmbeddingModel() *string
	Enable() cdktf.IResolvable
	EngineVersion() *float64
	Filter() DataCloudflareAiSearchInstanceFilterOutputReference
	FilterInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FusionMethod() *string
	HybridSearchEnabled() cdktf.IResolvable
	Id() *string
	SetId(val *string)
	IdInput() *string
	IndexingOptions() DataCloudflareAiSearchInstanceIndexingOptionsOutputReference
	IndexMethod() DataCloudflareAiSearchInstanceIndexMethodOutputReference
	LastActivity() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxNumResults() *float64
	Metadata() DataCloudflareAiSearchInstanceMetadataOutputReference
	ModifiedAt() *string
	ModifiedBy() *string
	Namespace() *string
	// The tree node.
	Node() constructs.Node
	Paused() cdktf.IResolvable
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	PublicEndpointId() *string
	PublicEndpointParams() DataCloudflareAiSearchInstancePublicEndpointParamsOutputReference
	// Experimental.
	RawOverrides() interface{}
	Reranking() cdktf.IResolvable
	RerankingModel() *string
	RetrievalOptions() DataCloudflareAiSearchInstanceRetrievalOptionsOutputReference
	RewriteModel() *string
	RewriteQuery() cdktf.IResolvable
	ScoreThreshold() *float64
	Source() *string
	SourceParams() DataCloudflareAiSearchInstanceSourceParamsOutputReference
	Status() *string
	SyncInterval() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenId() *string
	Type() *string
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
	PutFilter(value *DataCloudflareAiSearchInstanceFilter)
	ResetAccountId()
	ResetFilter()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DataCloudflareAiSearchInstance
type jsiiProxy_DataCloudflareAiSearchInstance struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) AiGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aiGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) AisearchModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aisearchModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Cache() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) CacheThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) ChunkOverlap() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"chunkOverlap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) ChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"chunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) CustomMetadata() DataCloudflareAiSearchInstanceCustomMetadataList {
	var returns DataCloudflareAiSearchInstanceCustomMetadataList
	_jsii_.Get(
		j,
		"customMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) EmbeddingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Enable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) EngineVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Filter() DataCloudflareAiSearchInstanceFilterOutputReference {
	var returns DataCloudflareAiSearchInstanceFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) FusionMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fusionMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) HybridSearchEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hybridSearchEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) IndexingOptions() DataCloudflareAiSearchInstanceIndexingOptionsOutputReference {
	var returns DataCloudflareAiSearchInstanceIndexingOptionsOutputReference
	_jsii_.Get(
		j,
		"indexingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) IndexMethod() DataCloudflareAiSearchInstanceIndexMethodOutputReference {
	var returns DataCloudflareAiSearchInstanceIndexMethodOutputReference
	_jsii_.Get(
		j,
		"indexMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) LastActivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastActivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) MaxNumResults() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Metadata() DataCloudflareAiSearchInstanceMetadataOutputReference {
	var returns DataCloudflareAiSearchInstanceMetadataOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) ModifiedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Paused() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) PublicEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) PublicEndpointParams() DataCloudflareAiSearchInstancePublicEndpointParamsOutputReference {
	var returns DataCloudflareAiSearchInstancePublicEndpointParamsOutputReference
	_jsii_.Get(
		j,
		"publicEndpointParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Reranking() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reranking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) RerankingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rerankingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) RetrievalOptions() DataCloudflareAiSearchInstanceRetrievalOptionsOutputReference {
	var returns DataCloudflareAiSearchInstanceRetrievalOptionsOutputReference
	_jsii_.Get(
		j,
		"retrievalOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) RewriteModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewriteModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) RewriteQuery() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rewriteQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) ScoreThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scoreThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) SourceParams() DataCloudflareAiSearchInstanceSourceParamsOutputReference {
	var returns DataCloudflareAiSearchInstanceSourceParamsOutputReference
	_jsii_.Get(
		j,
		"sourceParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) SyncInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) TokenId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instance cloudflare_ai_search_instance} Data Source.
func NewDataCloudflareAiSearchInstance(scope constructs.Construct, id *string, config *DataCloudflareAiSearchInstanceConfig) DataCloudflareAiSearchInstance {
	_init_.Initialize()

	if err := validateNewDataCloudflareAiSearchInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareAiSearchInstance{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareAiSearchInstance.DataCloudflareAiSearchInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/ai_search_instance cloudflare_ai_search_instance} Data Source.
func NewDataCloudflareAiSearchInstance_Override(d DataCloudflareAiSearchInstance, scope constructs.Construct, id *string, config *DataCloudflareAiSearchInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareAiSearchInstance.DataCloudflareAiSearchInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataCloudflareAiSearchInstance resource upon running "cdktf plan <stack-name>".
func DataCloudflareAiSearchInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataCloudflareAiSearchInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareAiSearchInstance.DataCloudflareAiSearchInstance",
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
func DataCloudflareAiSearchInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareAiSearchInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareAiSearchInstance.DataCloudflareAiSearchInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareAiSearchInstance_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareAiSearchInstance_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareAiSearchInstance.DataCloudflareAiSearchInstance",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareAiSearchInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareAiSearchInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareAiSearchInstance.DataCloudflareAiSearchInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataCloudflareAiSearchInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.dataCloudflareAiSearchInstance.DataCloudflareAiSearchInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) PutFilter(value *DataCloudflareAiSearchInstanceFilter) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) ResetAccountId() {
	_jsii_.InvokeVoid(
		d,
		"resetAccountId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

