package googlecomputevpntunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputevpntunnel/internal"
)

type GoogleComputeVpnTunnelCipherSuiteOutputReference interface {
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
	InternalValue() *GoogleComputeVpnTunnelCipherSuite
	SetInternalValue(val *GoogleComputeVpnTunnelCipherSuite)
	Phase1() GoogleComputeVpnTunnelCipherSuitePhase1OutputReference
	Phase1Input() *GoogleComputeVpnTunnelCipherSuitePhase1
	Phase2() GoogleComputeVpnTunnelCipherSuitePhase2OutputReference
	Phase2Input() *GoogleComputeVpnTunnelCipherSuitePhase2
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
	PutPhase1(value *GoogleComputeVpnTunnelCipherSuitePhase1)
	PutPhase2(value *GoogleComputeVpnTunnelCipherSuitePhase2)
	ResetPhase1()
	ResetPhase2()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeVpnTunnelCipherSuiteOutputReference
type jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) InternalValue() *GoogleComputeVpnTunnelCipherSuite {
	var returns *GoogleComputeVpnTunnelCipherSuite
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) Phase1() GoogleComputeVpnTunnelCipherSuitePhase1OutputReference {
	var returns GoogleComputeVpnTunnelCipherSuitePhase1OutputReference
	_jsii_.Get(
		j,
		"phase1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) Phase1Input() *GoogleComputeVpnTunnelCipherSuitePhase1 {
	var returns *GoogleComputeVpnTunnelCipherSuitePhase1
	_jsii_.Get(
		j,
		"phase1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) Phase2() GoogleComputeVpnTunnelCipherSuitePhase2OutputReference {
	var returns GoogleComputeVpnTunnelCipherSuitePhase2OutputReference
	_jsii_.Get(
		j,
		"phase2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) Phase2Input() *GoogleComputeVpnTunnelCipherSuitePhase2 {
	var returns *GoogleComputeVpnTunnelCipherSuitePhase2
	_jsii_.Get(
		j,
		"phase2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeVpnTunnelCipherSuiteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeVpnTunnelCipherSuiteOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeVpnTunnelCipherSuiteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnelCipherSuiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeVpnTunnelCipherSuiteOutputReference_Override(g GoogleComputeVpnTunnelCipherSuiteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnelCipherSuiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference)SetInternalValue(val *GoogleComputeVpnTunnelCipherSuite) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) PutPhase1(value *GoogleComputeVpnTunnelCipherSuitePhase1) {
	if err := g.validatePutPhase1Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPhase1",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) PutPhase2(value *GoogleComputeVpnTunnelCipherSuitePhase2) {
	if err := g.validatePutPhase2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPhase2",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) ResetPhase1() {
	_jsii_.InvokeVoid(
		g,
		"resetPhase1",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) ResetPhase2() {
	_jsii_.InvokeVoid(
		g,
		"resetPhase2",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

