package datacloudflareaisearchinstances

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/datacloudflareaisearchinstances/internal"
)

type DataCloudflareAiSearchInstancesResultOutputReference interface {
	cdktf.ComplexObject
	AiGatewayId() *string
	AisearchModel() *string
	Cache() cdktf.IResolvable
	CacheThreshold() *string
	ChunkOverlap() *float64
	ChunkSize() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CreatedAt() *string
	CreatedBy() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomMetadata() DataCloudflareAiSearchInstancesResultCustomMetadataList
	EmbeddingModel() *string
	Enable() cdktf.IResolvable
	EngineVersion() *float64
	// Experimental.
	Fqn() *string
	FusionMethod() *string
	HybridSearchEnabled() cdktf.IResolvable
	Id() *string
	IndexingOptions() DataCloudflareAiSearchInstancesResultIndexingOptionsOutputReference
	IndexMethod() DataCloudflareAiSearchInstancesResultIndexMethodOutputReference
	InternalValue() *DataCloudflareAiSearchInstancesResult
	SetInternalValue(val *DataCloudflareAiSearchInstancesResult)
	LastActivity() *string
	MaxNumResults() *float64
	Metadata() DataCloudflareAiSearchInstancesResultMetadataOutputReference
	ModifiedAt() *string
	ModifiedBy() *string
	Namespace() *string
	Paused() cdktf.IResolvable
	PublicEndpointId() *string
	PublicEndpointParams() DataCloudflareAiSearchInstancesResultPublicEndpointParamsOutputReference
	Reranking() cdktf.IResolvable
	RerankingModel() *string
	RetrievalOptions() DataCloudflareAiSearchInstancesResultRetrievalOptionsOutputReference
	RewriteModel() *string
	RewriteQuery() cdktf.IResolvable
	ScoreThreshold() *float64
	Source() *string
	SourceParams() DataCloudflareAiSearchInstancesResultSourceParamsOutputReference
	Status() *string
	SyncInterval() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenId() *string
	Type() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareAiSearchInstancesResultOutputReference
type jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) AiGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aiGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) AisearchModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aisearchModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Cache() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) CacheThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ChunkOverlap() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"chunkOverlap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"chunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) CustomMetadata() DataCloudflareAiSearchInstancesResultCustomMetadataList {
	var returns DataCloudflareAiSearchInstancesResultCustomMetadataList
	_jsii_.Get(
		j,
		"customMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) EmbeddingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"embeddingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Enable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) EngineVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) FusionMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fusionMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) HybridSearchEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hybridSearchEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) IndexingOptions() DataCloudflareAiSearchInstancesResultIndexingOptionsOutputReference {
	var returns DataCloudflareAiSearchInstancesResultIndexingOptionsOutputReference
	_jsii_.Get(
		j,
		"indexingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) IndexMethod() DataCloudflareAiSearchInstancesResultIndexMethodOutputReference {
	var returns DataCloudflareAiSearchInstancesResultIndexMethodOutputReference
	_jsii_.Get(
		j,
		"indexMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) InternalValue() *DataCloudflareAiSearchInstancesResult {
	var returns *DataCloudflareAiSearchInstancesResult
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) LastActivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastActivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) MaxNumResults() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNumResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Metadata() DataCloudflareAiSearchInstancesResultMetadataOutputReference {
	var returns DataCloudflareAiSearchInstancesResultMetadataOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ModifiedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Paused() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) PublicEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) PublicEndpointParams() DataCloudflareAiSearchInstancesResultPublicEndpointParamsOutputReference {
	var returns DataCloudflareAiSearchInstancesResultPublicEndpointParamsOutputReference
	_jsii_.Get(
		j,
		"publicEndpointParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Reranking() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reranking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) RerankingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rerankingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) RetrievalOptions() DataCloudflareAiSearchInstancesResultRetrievalOptionsOutputReference {
	var returns DataCloudflareAiSearchInstancesResultRetrievalOptionsOutputReference
	_jsii_.Get(
		j,
		"retrievalOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) RewriteModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rewriteModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) RewriteQuery() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rewriteQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ScoreThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scoreThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) SourceParams() DataCloudflareAiSearchInstancesResultSourceParamsOutputReference {
	var returns DataCloudflareAiSearchInstancesResultSourceParamsOutputReference
	_jsii_.Get(
		j,
		"sourceParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) SyncInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) TokenId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewDataCloudflareAiSearchInstancesResultOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareAiSearchInstancesResultOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareAiSearchInstancesResultOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareAiSearchInstances.DataCloudflareAiSearchInstancesResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareAiSearchInstancesResultOutputReference_Override(d DataCloudflareAiSearchInstancesResultOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareAiSearchInstances.DataCloudflareAiSearchInstancesResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference)SetInternalValue(val *DataCloudflareAiSearchInstancesResult) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiSearchInstancesResultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

