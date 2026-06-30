package googleclouddeploytarget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleclouddeploytarget/internal"
)

type GoogleClouddeployTargetGkeOutputReference interface {
	cdktf.ComplexObject
	Cluster() *string
	SetCluster(val *string)
	ClusterInput() *string
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
	DnsEndpoint() any
	SetDnsEndpoint(val any)
	DnsEndpointInput() any
	// Experimental.
	Fqn() *string
	InternalIp() any
	SetInternalIp(val any)
	InternalIpInput() any
	InternalValue() *GoogleClouddeployTargetGke
	SetInternalValue(val *GoogleClouddeployTargetGke)
	ProxyUrl() *string
	SetProxyUrl(val *string)
	ProxyUrlInput() *string
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
	ResetCluster()
	ResetDnsEndpoint()
	ResetInternalIp()
	ResetProxyUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployTargetGkeOutputReference
type jsiiProxy_GoogleClouddeployTargetGkeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) DnsEndpoint() any {
	var returns any
	_jsii_.Get(
		j,
		"dnsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) DnsEndpointInput() any {
	var returns any
	_jsii_.Get(
		j,
		"dnsEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) InternalIp() any {
	var returns any
	_jsii_.Get(
		j,
		"internalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) InternalIpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"internalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) InternalValue() *GoogleClouddeployTargetGke {
	var returns *GoogleClouddeployTargetGke
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) ProxyUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) ProxyUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleClouddeployTargetGkeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleClouddeployTargetGkeOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployTargetGkeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployTargetGkeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleClouddeployTarget.GoogleClouddeployTargetGkeOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleClouddeployTargetGkeOutputReference_Override(g GoogleClouddeployTargetGkeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleClouddeployTarget.GoogleClouddeployTargetGkeOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) SetDnsEndpoint(val any) {
	if err := j.validateSetDnsEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) SetInternalIp(val any) {
	if err := j.validateSetInternalIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalIp",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) SetInternalValue(val *GoogleClouddeployTargetGke) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) SetProxyUrl(val *string) {
	if err := j.validateSetProxyUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyUrl",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) ResetCluster() {
	_jsii_.InvokeVoid(
		g,
		"resetCluster",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) ResetDnsEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) ResetInternalIp() {
	_jsii_.InvokeVoid(
		g,
		"resetInternalIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) ResetProxyUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetProxyUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleClouddeployTargetGkeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
