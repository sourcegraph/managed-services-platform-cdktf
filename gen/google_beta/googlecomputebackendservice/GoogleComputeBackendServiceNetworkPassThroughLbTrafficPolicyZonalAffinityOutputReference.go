package googlecomputebackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputebackendservice/internal"
)

type GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinity
	SetInternalValue(val *GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinity)
	Spillover() *string
	SetSpillover(val *string)
	SpilloverInput() *string
	SpilloverRatio() *float64
	SetSpilloverRatio(val *float64)
	SpilloverRatioInput() *float64
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
	ResetSpillover()
	ResetSpilloverRatio()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference
type jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) InternalValue() *GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinity {
	var returns *GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) Spillover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spillover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) SpilloverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spilloverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) SpilloverRatio() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spilloverRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) SpilloverRatioInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spilloverRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeBackendService.GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference_Override(g GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeBackendService.GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference)SetInternalValue(val *GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference)SetSpillover(val *string) {
	if err := j.validateSetSpilloverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spillover",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference)SetSpilloverRatio(val *float64) {
	if err := j.validateSetSpilloverRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spilloverRatio",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) ResetSpillover() {
	_jsii_.InvokeVoid(
		g,
		"resetSpillover",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) ResetSpilloverRatio() {
	_jsii_.InvokeVoid(
		g,
		"resetSpilloverRatio",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceNetworkPassThroughLbTrafficPolicyZonalAffinityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

