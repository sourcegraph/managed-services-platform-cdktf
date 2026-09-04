package aisearchinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/aisearchinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance cloudflare_ai_search_instance}.
type AiSearchInstance interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AccountTag() *string
	AiGatewayId() *string
	SetAiGatewayId(val *string)
	AiGatewayIdInput() *string
	AisearchModel() *string
	SetAisearchModel(val *string)
	AisearchModelInput() *string
	Cache() interface{}
	SetCache(val interface{})
	CacheInput() interface{}
	CacheThreshold() *string
	SetCacheThreshold(val *string)
	CacheThresholdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Chunk() interface{}
	SetChunk(val interface{})
	ChunkInput() interface{}
	ChunkOverlap() *float64
	SetChunkOverlap(val *float64)
	ChunkOverlapInput() *float64
	ChunkSize() *float64
	SetChunkSize(val *float64)
	ChunkSizeInput() *float64
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	CreatedBy() *string
	CustomMetadata() AiSearchInstanceCustomMetadataList
	CustomMetadataInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EmbeddingModel() *string
	SetEmbeddingModel(val *string)
	EmbeddingModelInput() *string
	Enable() cdktf.IResolvable
	EngineVersion() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HybridSearchEnabled() interface{}
	SetHybridSearchEnabled(val interface{})
	HybridSearchEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalId() *string
	LastActivity() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxNumResults() *float64
	SetMaxNumResults(val *float64)
	MaxNumResultsInput() *float64
	Metadata() AiSearchInstanceMetadataOutputReference
	MetadataInput() interface{}
	ModifiedAt() *string
	ModifiedBy() *string
	// The tree node.
	Node() constructs.Node
	Paused() interface{}
	SetPaused(val interface{})
	PausedInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicEndpointId() *string
	PublicEndpointParams() AiSearchInstancePublicEndpointParamsOutputReference
	PublicEndpointParamsInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Reranking() interface{}
	SetReranking(val interface{})
	RerankingInput() interface{}
	RerankingModel() *string
	SetRerankingModel(val *string)
	RerankingModelInput() *string
	RewriteModel() *string
	SetRewriteModel(val *string)
	RewriteModelInput() *string
	RewriteQuery() interface{}
	SetRewriteQuery(val interface{})
	RewriteQueryInput() interface{}
	ScoreThreshold() *float64
	SetScoreThreshold(val *float64)
	ScoreThresholdInput() *float64
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	SourceParams() AiSearchInstanceSourceParamsOutputReference
	SourceParamsInput() interface{}
	Status() *string
	Summarization() interface{}
	SetSummarization(val interface{})
	SummarizationInput() interface{}
	SummarizationModel() *string
	SetSummarizationModel(val *string)
	SummarizationModelInput() *string
	SystemPromptAisearch() *string
	SetSystemPromptAisearch(val *string)
	SystemPromptAisearchInput() *string
	SystemPromptIndexSummarization() *string
	SetSystemPromptIndexSummarization(val *string)
	SystemPromptIndexSummarizationInput() *string
	SystemPromptRewriteQuery() *string
	SetSystemPromptRewriteQuery(val *string)
	SystemPromptRewriteQueryInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenId() *string
	SetTokenId(val *string)
	TokenIdInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VectorizeActiveNamespace() *string
	VectorizeName() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCustomMetadata(value interface{})
	PutMetadata(value *AiSearchInstanceMetadata)
	PutPublicEndpointParams(value *AiSearchInstancePublicEndpointParams)
	PutSourceParams(value *AiSearchInstanceSourceParams)
	ResetAiGatewayId()
	ResetAisearchModel()
	ResetCache()
	ResetCacheThreshold()
	ResetChunk()
	ResetChunkOverlap()
	ResetChunkSize()
	ResetCustomMetadata()
	ResetEmbeddingModel()
	ResetHybridSearchEnabled()
	ResetMaxNumResults()
	ResetMetadata()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPaused()
	ResetPublicEndpointParams()
	ResetReranking()
	ResetRerankingModel()
	ResetRewriteModel()
	ResetRewriteQuery()
	ResetScoreThreshold()
	ResetSourceParams()
	ResetSummarization()
	ResetSummarizationModel()
	ResetSystemPromptAisearch()
	ResetSystemPromptIndexSummarization()
	ResetSystemPromptRewriteQuery()
	ResetTokenId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for AiSearchInstance
type jsiiProxy_AiSearchInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AiSearchInstance) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) AccountTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) AiGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aiGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) AiGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aiGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) AisearchModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aisearchModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) AisearchModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aisearchModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Cache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) CacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) CacheThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) CacheThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Chunk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chunk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) ChunkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chunkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) ChunkOverlap() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"chunkOverlap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) ChunkOverlapInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"chunkOverlapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) ChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"chunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) ChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"chunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) CustomMetadata() AiSearchInstanceCustomMetadataList {
	var returns AiSearchInstanceCustomMetadataList
	_jsii_.Get(
		j,
		"customMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) CustomMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) EmbeddingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) EmbeddingModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Enable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) EngineVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) HybridSearchEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hybridSearchEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) HybridSearchEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hybridSearchEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) InternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) LastActivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastActivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) MaxNumResults() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) MaxNumResultsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Metadata() AiSearchInstanceMetadataOutputReference {
	var returns AiSearchInstanceMetadataOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) ModifiedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Paused() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) PausedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pausedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) PublicEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) PublicEndpointParams() AiSearchInstancePublicEndpointParamsOutputReference {
	var returns AiSearchInstancePublicEndpointParamsOutputReference
	_jsii_.Get(
		j,
		"publicEndpointParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) PublicEndpointParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicEndpointParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Reranking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reranking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) RerankingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rerankingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) RerankingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rerankingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) RerankingModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rerankingModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) RewriteModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewriteModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) RewriteModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewriteModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) RewriteQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rewriteQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) RewriteQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rewriteQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) ScoreThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scoreThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) ScoreThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scoreThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SourceParams() AiSearchInstanceSourceParamsOutputReference {
	var returns AiSearchInstanceSourceParamsOutputReference
	_jsii_.Get(
		j,
		"sourceParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SourceParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Summarization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"summarization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SummarizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"summarizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SummarizationModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summarizationModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SummarizationModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summarizationModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SystemPromptAisearch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemPromptAisearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SystemPromptAisearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemPromptAisearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SystemPromptIndexSummarization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemPromptIndexSummarization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SystemPromptIndexSummarizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemPromptIndexSummarizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SystemPromptRewriteQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemPromptRewriteQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) SystemPromptRewriteQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemPromptRewriteQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) TokenId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) TokenIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) VectorizeActiveNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorizeActiveNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstance) VectorizeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorizeName",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance cloudflare_ai_search_instance} Resource.
func NewAiSearchInstance(scope constructs.Construct, id *string, config *AiSearchInstanceConfig) AiSearchInstance {
	_init_.Initialize()

	if err := validateNewAiSearchInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchInstance{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance cloudflare_ai_search_instance} Resource.
func NewAiSearchInstance_Override(a AiSearchInstance, scope constructs.Construct, id *string, config *AiSearchInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstance",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetAiGatewayId(val *string) {
	if err := j.validateSetAiGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aiGatewayId",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetAisearchModel(val *string) {
	if err := j.validateSetAisearchModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aisearchModel",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetCache(val interface{}) {
	if err := j.validateSetCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cache",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetCacheThreshold(val *string) {
	if err := j.validateSetCacheThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheThreshold",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetChunk(val interface{}) {
	if err := j.validateSetChunkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chunk",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetChunkOverlap(val *float64) {
	if err := j.validateSetChunkOverlapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chunkOverlap",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetChunkSize(val *float64) {
	if err := j.validateSetChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chunkSize",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetEmbeddingModel(val *string) {
	if err := j.validateSetEmbeddingModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingModel",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetHybridSearchEnabled(val interface{}) {
	if err := j.validateSetHybridSearchEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hybridSearchEnabled",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetMaxNumResults(val *float64) {
	if err := j.validateSetMaxNumResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNumResults",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetPaused(val interface{}) {
	if err := j.validateSetPausedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paused",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetReranking(val interface{}) {
	if err := j.validateSetRerankingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reranking",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetRerankingModel(val *string) {
	if err := j.validateSetRerankingModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rerankingModel",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetRewriteModel(val *string) {
	if err := j.validateSetRewriteModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rewriteModel",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetRewriteQuery(val interface{}) {
	if err := j.validateSetRewriteQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rewriteQuery",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetScoreThreshold(val *float64) {
	if err := j.validateSetScoreThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scoreThreshold",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetSummarization(val interface{}) {
	if err := j.validateSetSummarizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summarization",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetSummarizationModel(val *string) {
	if err := j.validateSetSummarizationModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summarizationModel",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetSystemPromptAisearch(val *string) {
	if err := j.validateSetSystemPromptAisearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemPromptAisearch",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetSystemPromptIndexSummarization(val *string) {
	if err := j.validateSetSystemPromptIndexSummarizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemPromptIndexSummarization",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetSystemPromptRewriteQuery(val *string) {
	if err := j.validateSetSystemPromptRewriteQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemPromptRewriteQuery",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetTokenId(val *string) {
	if err := j.validateSetTokenIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenId",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstance)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a AiSearchInstance resource upon running "cdktf plan <stack-name>".
func AiSearchInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAiSearchInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstance",
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
func AiSearchInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiSearchInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AiSearchInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiSearchInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AiSearchInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAiSearchInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AiSearchInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AiSearchInstance) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AiSearchInstance) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AiSearchInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AiSearchInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AiSearchInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AiSearchInstance) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AiSearchInstance) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AiSearchInstance) PutCustomMetadata(value interface{}) {
	if err := a.validatePutCustomMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomMetadata",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchInstance) PutMetadata(value *AiSearchInstanceMetadata) {
	if err := a.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetadata",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchInstance) PutPublicEndpointParams(value *AiSearchInstancePublicEndpointParams) {
	if err := a.validatePutPublicEndpointParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPublicEndpointParams",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchInstance) PutSourceParams(value *AiSearchInstanceSourceParams) {
	if err := a.validatePutSourceParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceParams",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetAiGatewayId() {
	_jsii_.InvokeVoid(
		a,
		"resetAiGatewayId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetAisearchModel() {
	_jsii_.InvokeVoid(
		a,
		"resetAisearchModel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetCache() {
	_jsii_.InvokeVoid(
		a,
		"resetCache",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetCacheThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetChunk() {
	_jsii_.InvokeVoid(
		a,
		"resetChunk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetChunkOverlap() {
	_jsii_.InvokeVoid(
		a,
		"resetChunkOverlap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetChunkSize() {
	_jsii_.InvokeVoid(
		a,
		"resetChunkSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetCustomMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetEmbeddingModel() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbeddingModel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetHybridSearchEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetHybridSearchEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetMaxNumResults() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxNumResults",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetPaused() {
	_jsii_.InvokeVoid(
		a,
		"resetPaused",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetPublicEndpointParams() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicEndpointParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetReranking() {
	_jsii_.InvokeVoid(
		a,
		"resetReranking",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetRerankingModel() {
	_jsii_.InvokeVoid(
		a,
		"resetRerankingModel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetRewriteModel() {
	_jsii_.InvokeVoid(
		a,
		"resetRewriteModel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetRewriteQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetRewriteQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetScoreThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetScoreThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetSourceParams() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetSummarization() {
	_jsii_.InvokeVoid(
		a,
		"resetSummarization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetSummarizationModel() {
	_jsii_.InvokeVoid(
		a,
		"resetSummarizationModel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetSystemPromptAisearch() {
	_jsii_.InvokeVoid(
		a,
		"resetSystemPromptAisearch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetSystemPromptIndexSummarization() {
	_jsii_.InvokeVoid(
		a,
		"resetSystemPromptIndexSummarization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetSystemPromptRewriteQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetSystemPromptRewriteQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) ResetTokenId() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

