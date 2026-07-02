package googlegkeonprembaremetalcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlegkeonprembaremetalcluster/internal"
)

type GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig
	SetInternalValue(val *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	NodeConfigs() GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfigsList
	NodeConfigsInput() any
	OperatingSystem() *string
	SetOperatingSystem(val *string)
	OperatingSystemInput() *string
	Taints() GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaintsList
	TaintsInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutNodeConfigs(value any)
	PutTaints(value any)
	ResetLabels()
	ResetNodeConfigs()
	ResetOperatingSystem()
	ResetTaints()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference
type jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) InternalValue() *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig {
	var returns *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) NodeConfigs() GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfigsList {
	var returns GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfigsList
	_jsii_.Get(
		j,
		"nodeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) NodeConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"nodeConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) OperatingSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) Taints() GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaintsList {
	var returns GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) TaintsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference_Override(g GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetInternalValue(val *GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetOperatingSystem(val *string) {
	if err := j.validateSetOperatingSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) PutNodeConfigs(value any) {
	if err := g.validatePutNodeConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeConfigs",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) PutTaints(value any) {
	if err := g.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTaints",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ResetNodeConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ResetOperatingSystem() {
	_jsii_.InvokeVoid(
		g,
		"resetOperatingSystem",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ResetTaints() {
	_jsii_.InvokeVoid(
		g,
		"resetTaints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
