package teamsaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/teamsaccount/internal"
)

type TeamsAccountProxyOutputReference interface {
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
	DisableForTime() *float64
	SetDisableForTime(val *float64)
	DisableForTimeInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *TeamsAccountProxy
	SetInternalValue(val *TeamsAccountProxy)
	RootCa() interface{}
	SetRootCa(val interface{})
	RootCaInput() interface{}
	Tcp() interface{}
	SetTcp(val interface{})
	TcpInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Udp() interface{}
	SetUdp(val interface{})
	UdpInput() interface{}
	VirtualIp() interface{}
	SetVirtualIp(val interface{})
	VirtualIpInput() interface{}
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

// The jsii proxy struct for TeamsAccountProxyOutputReference
type jsiiProxy_TeamsAccountProxyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) DisableForTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disableForTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) DisableForTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disableForTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) InternalValue() *TeamsAccountProxy {
	var returns *TeamsAccountProxy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) RootCa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) RootCaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) Tcp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) TcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) Udp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"udp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) UdpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"udpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) VirtualIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference) VirtualIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virtualIpInput",
		&returns,
	)
	return returns
}


func NewTeamsAccountProxyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TeamsAccountProxyOutputReference {
	_init_.Initialize()

	if err := validateNewTeamsAccountProxyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamsAccountProxyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccountProxyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTeamsAccountProxyOutputReference_Override(t TeamsAccountProxyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.teamsAccount.TeamsAccountProxyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference)SetDisableForTime(val *float64) {
	if err := j.validateSetDisableForTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableForTime",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference)SetInternalValue(val *TeamsAccountProxy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference)SetRootCa(val interface{}) {
	if err := j.validateSetRootCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootCa",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference)SetTcp(val interface{}) {
	if err := j.validateSetTcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcp",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference)SetUdp(val interface{}) {
	if err := j.validateSetUdpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"udp",
		val,
	)
}

func (j *jsiiProxy_TeamsAccountProxyOutputReference)SetVirtualIp(val interface{}) {
	if err := j.validateSetVirtualIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualIp",
		val,
	)
}

func (t *jsiiProxy_TeamsAccountProxyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountProxyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TeamsAccountProxyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TeamsAccountProxyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TeamsAccountProxyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TeamsAccountProxyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TeamsAccountProxyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TeamsAccountProxyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TeamsAccountProxyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TeamsAccountProxyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TeamsAccountProxyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamsAccountProxyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TeamsAccountProxyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TeamsAccountProxyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

