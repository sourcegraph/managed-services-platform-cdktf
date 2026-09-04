package datacloudflarepagesprojects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/datacloudflarepagesprojects/internal"
)

type DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference interface {
	cdktf.ComplexObject
	AiBindings() DataCloudflarePagesProjectsResultDeploymentConfigsProductionAiBindingsMap
	AlwaysUseLatestCompatibilityDate() cdktf.IResolvable
	AnalyticsEngineDatasets() DataCloudflarePagesProjectsResultDeploymentConfigsProductionAnalyticsEngineDatasetsMap
	Browsers() DataCloudflarePagesProjectsResultDeploymentConfigsProductionBrowsersMap
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
	D1Databases() DataCloudflarePagesProjectsResultDeploymentConfigsProductionD1DatabasesMap
	DurableObjectNamespaces() DataCloudflarePagesProjectsResultDeploymentConfigsProductionDurableObjectNamespacesMap
	EnvVars() DataCloudflarePagesProjectsResultDeploymentConfigsProductionEnvVarsMap
	FailOpen() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	HyperdriveBindings() DataCloudflarePagesProjectsResultDeploymentConfigsProductionHyperdriveBindingsMap
	InternalValue() *DataCloudflarePagesProjectsResultDeploymentConfigsProduction
	SetInternalValue(val *DataCloudflarePagesProjectsResultDeploymentConfigsProduction)
	KvNamespaces() DataCloudflarePagesProjectsResultDeploymentConfigsProductionKvNamespacesMap
	Limits() DataCloudflarePagesProjectsResultDeploymentConfigsProductionLimitsOutputReference
	MtlsCertificates() DataCloudflarePagesProjectsResultDeploymentConfigsProductionMtlsCertificatesMap
	Placement() DataCloudflarePagesProjectsResultDeploymentConfigsProductionPlacementOutputReference
	QueueProducers() DataCloudflarePagesProjectsResultDeploymentConfigsProductionQueueProducersMap
	R2Buckets() DataCloudflarePagesProjectsResultDeploymentConfigsProductionR2BucketsMap
	Services() DataCloudflarePagesProjectsResultDeploymentConfigsProductionServicesMap
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsageModel() *string
	VectorizeBindings() DataCloudflarePagesProjectsResultDeploymentConfigsProductionVectorizeBindingsMap
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

// The jsii proxy struct for DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference
type jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) AiBindings() DataCloudflarePagesProjectsResultDeploymentConfigsProductionAiBindingsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionAiBindingsMap
	_jsii_.Get(
		j,
		"aiBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) AlwaysUseLatestCompatibilityDate() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"alwaysUseLatestCompatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) AnalyticsEngineDatasets() DataCloudflarePagesProjectsResultDeploymentConfigsProductionAnalyticsEngineDatasetsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionAnalyticsEngineDatasetsMap
	_jsii_.Get(
		j,
		"analyticsEngineDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) Browsers() DataCloudflarePagesProjectsResultDeploymentConfigsProductionBrowsersMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionBrowsersMap
	_jsii_.Get(
		j,
		"browsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) BuildImageMajorVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildImageMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) CompatibilityDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) CompatibilityFlags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) D1Databases() DataCloudflarePagesProjectsResultDeploymentConfigsProductionD1DatabasesMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionD1DatabasesMap
	_jsii_.Get(
		j,
		"d1Databases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) DurableObjectNamespaces() DataCloudflarePagesProjectsResultDeploymentConfigsProductionDurableObjectNamespacesMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionDurableObjectNamespacesMap
	_jsii_.Get(
		j,
		"durableObjectNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) EnvVars() DataCloudflarePagesProjectsResultDeploymentConfigsProductionEnvVarsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionEnvVarsMap
	_jsii_.Get(
		j,
		"envVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) FailOpen() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"failOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) HyperdriveBindings() DataCloudflarePagesProjectsResultDeploymentConfigsProductionHyperdriveBindingsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionHyperdriveBindingsMap
	_jsii_.Get(
		j,
		"hyperdriveBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) InternalValue() *DataCloudflarePagesProjectsResultDeploymentConfigsProduction {
	var returns *DataCloudflarePagesProjectsResultDeploymentConfigsProduction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) KvNamespaces() DataCloudflarePagesProjectsResultDeploymentConfigsProductionKvNamespacesMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionKvNamespacesMap
	_jsii_.Get(
		j,
		"kvNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) Limits() DataCloudflarePagesProjectsResultDeploymentConfigsProductionLimitsOutputReference {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionLimitsOutputReference
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) MtlsCertificates() DataCloudflarePagesProjectsResultDeploymentConfigsProductionMtlsCertificatesMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionMtlsCertificatesMap
	_jsii_.Get(
		j,
		"mtlsCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) Placement() DataCloudflarePagesProjectsResultDeploymentConfigsProductionPlacementOutputReference {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) QueueProducers() DataCloudflarePagesProjectsResultDeploymentConfigsProductionQueueProducersMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionQueueProducersMap
	_jsii_.Get(
		j,
		"queueProducers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) R2Buckets() DataCloudflarePagesProjectsResultDeploymentConfigsProductionR2BucketsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionR2BucketsMap
	_jsii_.Get(
		j,
		"r2Buckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) Services() DataCloudflarePagesProjectsResultDeploymentConfigsProductionServicesMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionServicesMap
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) UsageModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) VectorizeBindings() DataCloudflarePagesProjectsResultDeploymentConfigsProductionVectorizeBindingsMap {
	var returns DataCloudflarePagesProjectsResultDeploymentConfigsProductionVectorizeBindingsMap
	_jsii_.Get(
		j,
		"vectorizeBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) WranglerConfigHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wranglerConfigHash",
		&returns,
	)
	return returns
}


func NewDataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflarePagesProjects.DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference_Override(d DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflarePagesProjects.DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference)SetInternalValue(val *DataCloudflarePagesProjectsResultDeploymentConfigsProduction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultDeploymentConfigsProductionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

