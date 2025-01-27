package teamslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/teamslocation/internal"
)

type TeamsLocationEndpointsOutputReference interface {
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
	Doh() TeamsLocationEndpointsDohOutputReference
	DohInput() *TeamsLocationEndpointsDoh
	Dot() TeamsLocationEndpointsDotOutputReference
	DotInput() *TeamsLocationEndpointsDot
	// Experimental.
	Fqn() *string
	InternalValue() *TeamsLocationEndpoints
	SetInternalValue(val *TeamsLocationEndpoints)
	Ipv4() TeamsLocationEndpointsIpv4OutputReference
	Ipv4Input() *TeamsLocationEndpointsIpv4
	Ipv6() TeamsLocationEndpointsIpv6OutputReference
	Ipv6Input() *TeamsLocationEndpointsIpv6
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
	PutDoh(value *TeamsLocationEndpointsDoh)
	PutDot(value *TeamsLocationEndpointsDot)
	PutIpv4(value *TeamsLocationEndpointsIpv4)
	PutIpv6(value *TeamsLocationEndpointsIpv6)
	ResetDoh()
	ResetDot()
	ResetIpv4()
	ResetIpv6()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TeamsLocationEndpointsOutputReference
type jsiiProxy_TeamsLocationEndpointsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) Doh() TeamsLocationEndpointsDohOutputReference {
	var returns TeamsLocationEndpointsDohOutputReference
	_jsii_.Get(
		j,
		"doh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) DohInput() *TeamsLocationEndpointsDoh {
	var returns *TeamsLocationEndpointsDoh
	_jsii_.Get(
		j,
		"dohInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) Dot() TeamsLocationEndpointsDotOutputReference {
	var returns TeamsLocationEndpointsDotOutputReference
	_jsii_.Get(
		j,
		"dot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) DotInput() *TeamsLocationEndpointsDot {
	var returns *TeamsLocationEndpointsDot
	_jsii_.Get(
		j,
		"dotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) InternalValue() *TeamsLocationEndpoints {
	var returns *TeamsLocationEndpoints
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) Ipv4() TeamsLocationEndpointsIpv4OutputReference {
	var returns TeamsLocationEndpointsIpv4OutputReference
	_jsii_.Get(
		j,
		"ipv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) Ipv4Input() *TeamsLocationEndpointsIpv4 {
	var returns *TeamsLocationEndpointsIpv4
	_jsii_.Get(
		j,
		"ipv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) Ipv6() TeamsLocationEndpointsIpv6OutputReference {
	var returns TeamsLocationEndpointsIpv6OutputReference
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) Ipv6Input() *TeamsLocationEndpointsIpv6 {
	var returns *TeamsLocationEndpointsIpv6
	_jsii_.Get(
		j,
		"ipv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTeamsLocationEndpointsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TeamsLocationEndpointsOutputReference {
	_init_.Initialize()

	if err := validateNewTeamsLocationEndpointsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamsLocationEndpointsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsLocation.TeamsLocationEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTeamsLocationEndpointsOutputReference_Override(t TeamsLocationEndpointsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsLocation.TeamsLocationEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference)SetInternalValue(val *TeamsLocationEndpoints) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamsLocationEndpointsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) PutDoh(value *TeamsLocationEndpointsDoh) {
	if err := t.validatePutDohParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDoh",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) PutDot(value *TeamsLocationEndpointsDot) {
	if err := t.validatePutDotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDot",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) PutIpv4(value *TeamsLocationEndpointsIpv4) {
	if err := t.validatePutIpv4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpv4",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) PutIpv6(value *TeamsLocationEndpointsIpv6) {
	if err := t.validatePutIpv6Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpv6",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) ResetDoh() {
	_jsii_.InvokeVoid(
		t,
		"resetDoh",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) ResetDot() {
	_jsii_.InvokeVoid(
		t,
		"resetDot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) ResetIpv4() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv4",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) ResetIpv6() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsLocationEndpointsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

