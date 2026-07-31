package datacloudflarepagesprojects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/datacloudflarepagesprojects/internal"
)

type DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference interface {
	cdktf.ComplexObject
	AiBindings() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewAiBindingsMap
	AlwaysUseLatestCompatibilityDate() cdktf.IResolvable
	AnalyticsEngineDatasets() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewAnalyticsEngineDatasetsMap
	Browsers() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewBrowsersMap
	BuildImageMajorVersion() *float64
	CompatibilityDate() *string
	CompatibilityFlags() *[]*string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	D1Databases() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewD1DatabasesMap
	DurableObjectNamespaces() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewDurableObjectNamespacesMap
	EnvVars() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewEnvVarsMap
	FailOpen() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	HyperdriveBindings() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewHyperdriveBindingsMap
	InternalValue() *DataCloudflarePagesProjectsResultDeploymentConfigsPreview
	SetInternalValue(val *DataCloudflarePagesProjectsResultDeploymentConfigsPreview)
	KvNamespaces() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewKvNamespacesMap
	Limits() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewLimitsOutputReference
	MtlsCertificates() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewMtlsCertificatesMap
	Placement() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewPlacementOutputReference
	QueueProducers() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewQueueProducersMap
	R2Buckets() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewR2BucketsMap
	Services() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewServicesMap
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsageModel() *string
	VectorizeBindings() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewVectorizeBindingsMap
	WranglerConfigHash() *string
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

// The jsii proxy struct for DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference
type jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) AiBindings() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewAiBindingsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewAiBindingsMap
	_jsii_.Get(
		j,
		"aiBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) AlwaysUseLatestCompatibilityDate() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"alwaysUseLatestCompatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) AnalyticsEngineDatasets() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewAnalyticsEngineDatasetsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewAnalyticsEngineDatasetsMap
	_jsii_.Get(
		j,
		"analyticsEngineDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) Browsers() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewBrowsersMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewBrowsersMap
	_jsii_.Get(
		j,
		"browsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) BuildImageMajorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildImageMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) CompatibilityDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) CompatibilityFlags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) D1Databases() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewD1DatabasesMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewD1DatabasesMap
	_jsii_.Get(
		j,
		"d1Databases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) DurableObjectNamespaces() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewDurableObjectNamespacesMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewDurableObjectNamespacesMap
	_jsii_.Get(
		j,
		"durableObjectNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) EnvVars() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewEnvVarsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewEnvVarsMap
	_jsii_.Get(
		j,
		"envVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) FailOpen() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"failOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) HyperdriveBindings() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewHyperdriveBindingsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewHyperdriveBindingsMap
	_jsii_.Get(
		j,
		"hyperdriveBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) InternalValue() *DataCloudflarePagesProjectsResultDeploymentConfigsPreview {
	var returns *DataCloudflarePagesProjectsResultDeploymentConfigsPreview
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) KvNamespaces() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewKvNamespacesMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewKvNamespacesMap
	_jsii_.Get(
		j,
		"kvNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) Limits() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewLimitsOutputReference {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewLimitsOutputReference
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) MtlsCertificates() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewMtlsCertificatesMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewMtlsCertificatesMap
	_jsii_.Get(
		j,
		"mtlsCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) Placement() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewPlacementOutputReference {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) QueueProducers() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewQueueProducersMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewQueueProducersMap
	_jsii_.Get(
		j,
		"queueProducers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) R2Buckets() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewR2BucketsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewR2BucketsMap
	_jsii_.Get(
		j,
		"r2Buckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) Services() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewServicesMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewServicesMap
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) UsageModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) VectorizeBindings() DataCloudflarePagesProjectsResultDeploymentConfigsPreviewVectorizeBindingsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsPreviewVectorizeBindingsMap
	_jsii_.Get(
		j,
		"vectorizeBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) WranglerConfigHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wranglerConfigHash",
		&returns,
	)
	return returns
}


func NewDataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflarePagesProjects.DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference_Override(d DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflarePagesProjects.DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference)SetInternalValue(val *DataCloudflarePagesProjectsResultDeploymentConfigsPreview) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsPreviewOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

