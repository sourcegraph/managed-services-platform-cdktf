package pagesproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/pagesproject/internal"
)

type PagesProjectDeploymentConfigsPreviewOutputReference interface {
	cdktf.ComplexObject
	AiBindings() PagesProjectDeploymentConfigsPreviewAiBindingsMap
	AiBindingsInput() any
	AnalyticsEngineDatasets() PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsMap
	AnalyticsEngineDatasetsInput() any
	Browsers() PagesProjectDeploymentConfigsPreviewBrowsersMap
	BrowsersInput() any
	CompatibilityDate() *string
	SetCompatibilityDate(val *string)
	CompatibilityDateInput() *string
	CompatibilityFlags() *[]*string
	SetCompatibilityFlags(val *[]*string)
	CompatibilityFlagsInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	D1Databases() PagesProjectDeploymentConfigsPreviewD1DatabasesMap
	D1DatabasesInput() any
	DurableObjectNamespaces() PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesMap
	DurableObjectNamespacesInput() any
	EnvVars() PagesProjectDeploymentConfigsPreviewEnvVarsMap
	EnvVarsInput() any
	// Experimental.
	Fqn() *string
	HyperdriveBindings() PagesProjectDeploymentConfigsPreviewHyperdriveBindingsMap
	HyperdriveBindingsInput() any
	InternalValue() any
	SetInternalValue(val any)
	KvNamespaces() PagesProjectDeploymentConfigsPreviewKvNamespacesMap
	KvNamespacesInput() any
	MtlsCertificates() PagesProjectDeploymentConfigsPreviewMtlsCertificatesMap
	MtlsCertificatesInput() any
	Placement() PagesProjectDeploymentConfigsPreviewPlacementOutputReference
	PlacementInput() any
	QueueProducers() PagesProjectDeploymentConfigsPreviewQueueProducersMap
	QueueProducersInput() any
	R2Buckets() PagesProjectDeploymentConfigsPreviewR2BucketsMap
	R2BucketsInput() any
	Services() PagesProjectDeploymentConfigsPreviewServicesMap
	ServicesInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VectorizeBindings() PagesProjectDeploymentConfigsPreviewVectorizeBindingsMap
	VectorizeBindingsInput() any
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutAiBindings(value any)
	PutAnalyticsEngineDatasets(value any)
	PutBrowsers(value any)
	PutD1Databases(value any)
	PutDurableObjectNamespaces(value any)
	PutEnvVars(value any)
	PutHyperdriveBindings(value any)
	PutKvNamespaces(value any)
	PutMtlsCertificates(value any)
	PutPlacement(value *PagesProjectDeploymentConfigsPreviewPlacement)
	PutQueueProducers(value any)
	PutR2Buckets(value any)
	PutServices(value any)
	PutVectorizeBindings(value any)
	ResetAiBindings()
	ResetAnalyticsEngineDatasets()
	ResetBrowsers()
	ResetCompatibilityDate()
	ResetCompatibilityFlags()
	ResetD1Databases()
	ResetDurableObjectNamespaces()
	ResetEnvVars()
	ResetHyperdriveBindings()
	ResetKvNamespaces()
	ResetMtlsCertificates()
	ResetPlacement()
	ResetQueueProducers()
	ResetR2Buckets()
	ResetServices()
	ResetVectorizeBindings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PagesProjectDeploymentConfigsPreviewOutputReference
type jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) AiBindings() PagesProjectDeploymentConfigsPreviewAiBindingsMap {
	var returns PagesProjectDeploymentConfigsPreviewAiBindingsMap
	_jsii_.Get(
		j,
		"aiBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) AiBindingsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"aiBindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) AnalyticsEngineDatasets() PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsMap {
	var returns PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsMap
	_jsii_.Get(
		j,
		"analyticsEngineDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) AnalyticsEngineDatasetsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"analyticsEngineDatasetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) Browsers() PagesProjectDeploymentConfigsPreviewBrowsersMap {
	var returns PagesProjectDeploymentConfigsPreviewBrowsersMap
	_jsii_.Get(
		j,
		"browsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) BrowsersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"browsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) CompatibilityDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) CompatibilityDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) CompatibilityFlags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) CompatibilityFlagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) D1Databases() PagesProjectDeploymentConfigsPreviewD1DatabasesMap {
	var returns PagesProjectDeploymentConfigsPreviewD1DatabasesMap
	_jsii_.Get(
		j,
		"d1Databases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) D1DatabasesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"d1DatabasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) DurableObjectNamespaces() PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesMap {
	var returns PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesMap
	_jsii_.Get(
		j,
		"durableObjectNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) DurableObjectNamespacesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"durableObjectNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) EnvVars() PagesProjectDeploymentConfigsPreviewEnvVarsMap {
	var returns PagesProjectDeploymentConfigsPreviewEnvVarsMap
	_jsii_.Get(
		j,
		"envVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) EnvVarsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"envVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) HyperdriveBindings() PagesProjectDeploymentConfigsPreviewHyperdriveBindingsMap {
	var returns PagesProjectDeploymentConfigsPreviewHyperdriveBindingsMap
	_jsii_.Get(
		j,
		"hyperdriveBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) HyperdriveBindingsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"hyperdriveBindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) KvNamespaces() PagesProjectDeploymentConfigsPreviewKvNamespacesMap {
	var returns PagesProjectDeploymentConfigsPreviewKvNamespacesMap
	_jsii_.Get(
		j,
		"kvNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) KvNamespacesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"kvNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) MtlsCertificates() PagesProjectDeploymentConfigsPreviewMtlsCertificatesMap {
	var returns PagesProjectDeploymentConfigsPreviewMtlsCertificatesMap
	_jsii_.Get(
		j,
		"mtlsCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) MtlsCertificatesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"mtlsCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) Placement() PagesProjectDeploymentConfigsPreviewPlacementOutputReference {
	var returns PagesProjectDeploymentConfigsPreviewPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PlacementInput() any {
	var returns any
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) QueueProducers() PagesProjectDeploymentConfigsPreviewQueueProducersMap {
	var returns PagesProjectDeploymentConfigsPreviewQueueProducersMap
	_jsii_.Get(
		j,
		"queueProducers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) QueueProducersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"queueProducersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) R2Buckets() PagesProjectDeploymentConfigsPreviewR2BucketsMap {
	var returns PagesProjectDeploymentConfigsPreviewR2BucketsMap
	_jsii_.Get(
		j,
		"r2Buckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) R2BucketsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"r2BucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) Services() PagesProjectDeploymentConfigsPreviewServicesMap {
	var returns PagesProjectDeploymentConfigsPreviewServicesMap
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ServicesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"servicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) VectorizeBindings() PagesProjectDeploymentConfigsPreviewVectorizeBindingsMap {
	var returns PagesProjectDeploymentConfigsPreviewVectorizeBindingsMap
	_jsii_.Get(
		j,
		"vectorizeBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) VectorizeBindingsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"vectorizeBindingsInput",
		&returns,
	)
	return returns
}

func NewPagesProjectDeploymentConfigsPreviewOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PagesProjectDeploymentConfigsPreviewOutputReference {
	_init_.Initialize()

	if err := validateNewPagesProjectDeploymentConfigsPreviewOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectDeploymentConfigsPreviewOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPagesProjectDeploymentConfigsPreviewOutputReference_Override(p PagesProjectDeploymentConfigsPreviewOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pagesProject.PagesProjectDeploymentConfigsPreviewOutputReference",
		[]any{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) SetCompatibilityDate(val *string) {
	if err := j.validateSetCompatibilityDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityDate",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) SetCompatibilityFlags(val *[]*string) {
	if err := j.validateSetCompatibilityFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityFlags",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutAiBindings(value any) {
	if err := p.validatePutAiBindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAiBindings",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutAnalyticsEngineDatasets(value any) {
	if err := p.validatePutAnalyticsEngineDatasetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAnalyticsEngineDatasets",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutBrowsers(value any) {
	if err := p.validatePutBrowsersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putBrowsers",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutD1Databases(value any) {
	if err := p.validatePutD1DatabasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putD1Databases",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutDurableObjectNamespaces(value any) {
	if err := p.validatePutDurableObjectNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDurableObjectNamespaces",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutEnvVars(value any) {
	if err := p.validatePutEnvVarsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnvVars",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutHyperdriveBindings(value any) {
	if err := p.validatePutHyperdriveBindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHyperdriveBindings",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutKvNamespaces(value any) {
	if err := p.validatePutKvNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putKvNamespaces",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutMtlsCertificates(value any) {
	if err := p.validatePutMtlsCertificatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMtlsCertificates",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutPlacement(value *PagesProjectDeploymentConfigsPreviewPlacement) {
	if err := p.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPlacement",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutQueueProducers(value any) {
	if err := p.validatePutQueueProducersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQueueProducers",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutR2Buckets(value any) {
	if err := p.validatePutR2BucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putR2Buckets",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutServices(value any) {
	if err := p.validatePutServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putServices",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) PutVectorizeBindings(value any) {
	if err := p.validatePutVectorizeBindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVectorizeBindings",
		[]any{value},
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetAiBindings() {
	_jsii_.InvokeVoid(
		p,
		"resetAiBindings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetAnalyticsEngineDatasets() {
	_jsii_.InvokeVoid(
		p,
		"resetAnalyticsEngineDatasets",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetBrowsers() {
	_jsii_.InvokeVoid(
		p,
		"resetBrowsers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetCompatibilityDate() {
	_jsii_.InvokeVoid(
		p,
		"resetCompatibilityDate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetCompatibilityFlags() {
	_jsii_.InvokeVoid(
		p,
		"resetCompatibilityFlags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetD1Databases() {
	_jsii_.InvokeVoid(
		p,
		"resetD1Databases",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetDurableObjectNamespaces() {
	_jsii_.InvokeVoid(
		p,
		"resetDurableObjectNamespaces",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetEnvVars() {
	_jsii_.InvokeVoid(
		p,
		"resetEnvVars",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetHyperdriveBindings() {
	_jsii_.InvokeVoid(
		p,
		"resetHyperdriveBindings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetKvNamespaces() {
	_jsii_.InvokeVoid(
		p,
		"resetKvNamespaces",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetMtlsCertificates() {
	_jsii_.InvokeVoid(
		p,
		"resetMtlsCertificates",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetPlacement() {
	_jsii_.InvokeVoid(
		p,
		"resetPlacement",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetQueueProducers() {
	_jsii_.InvokeVoid(
		p,
		"resetQueueProducers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetR2Buckets() {
	_jsii_.InvokeVoid(
		p,
		"resetR2Buckets",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetServices() {
	_jsii_.InvokeVoid(
		p,
		"resetServices",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ResetVectorizeBindings() {
	_jsii_.InvokeVoid(
		p,
		"resetVectorizeBindings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		p,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagesProjectDeploymentConfigsPreviewOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
