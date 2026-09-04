package connectivitydirectoryservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/connectivitydirectoryservice/internal"
)

type ConnectivityDirectoryServiceHostOutputReference interface {
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
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv4() *string
	SetIpv4(val *string)
	Ipv4Input() *string
	Ipv6() *string
	SetIpv6(val *string)
	Ipv6Input() *string
	Network() ConnectivityDirectoryServiceHostNetworkOutputReference
	NetworkInput() interface{}
	ResolverNetwork() ConnectivityDirectoryServiceHostResolverNetworkOutputReference
	ResolverNetworkInput() interface{}
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
	PutNetwork(value *ConnectivityDirectoryServiceHostNetwork)
	PutResolverNetwork(value *ConnectivityDirectoryServiceHostResolverNetwork)
	ResetHostname()
	ResetIpv4()
	ResetIpv6()
	ResetNetwork()
	ResetResolverNetwork()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectivityDirectoryServiceHostOutputReference
type jsiiProxy_ConnectivityDirectoryServiceHostOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) Ipv4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) Ipv4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) Ipv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) Ipv6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) Network() ConnectivityDirectoryServiceHostNetworkOutputReference {
	var returns ConnectivityDirectoryServiceHostNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) NetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) ResolverNetwork() ConnectivityDirectoryServiceHostResolverNetworkOutputReference {
	var returns ConnectivityDirectoryServiceHostResolverNetworkOutputReference
	_jsii_.Get(
		j,
		"resolverNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) ResolverNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolverNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConnectivityDirectoryServiceHostOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConnectivityDirectoryServiceHostOutputReference {
	_init_.Initialize()

	if err := validateNewConnectivityDirectoryServiceHostOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectivityDirectoryServiceHostOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.connectivityDirectoryService.ConnectivityDirectoryServiceHostOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectivityDirectoryServiceHostOutputReference_Override(c ConnectivityDirectoryServiceHostOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.connectivityDirectoryService.ConnectivityDirectoryServiceHostOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference)SetIpv4(val *string) {
	if err := j.validateSetIpv4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4",
		val,
	)
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference)SetIpv6(val *string) {
	if err := j.validateSetIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6",
		val,
	)
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) PutNetwork(value *ConnectivityDirectoryServiceHostNetwork) {
	if err := c.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetwork",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) PutResolverNetwork(value *ConnectivityDirectoryServiceHostResolverNetwork) {
	if err := c.validatePutResolverNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResolverNetwork",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		c,
		"resetHostname",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) ResetIpv4() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv4",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) ResetIpv6() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv6",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) ResetResolverNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetResolverNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectivityDirectoryServiceHostOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

