package googlenetworksecurityservertlspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlenetworksecurityservertlspolicy/internal"
)

type GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference interface {
	cdktf.ComplexObject
	CertificateProviderInstance() GoogleNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstanceOutputReference
	CertificateProviderInstanceInput() *GoogleNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance
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
	GrpcEndpoint() GoogleNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointOutputReference
	GrpcEndpointInput() *GoogleNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint
	InternalValue() *GoogleNetworkSecurityServerTlsPolicyServerCertificate
	SetInternalValue(val *GoogleNetworkSecurityServerTlsPolicyServerCertificate)
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
	PutCertificateProviderInstance(value *GoogleNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance)
	PutGrpcEndpoint(value *GoogleNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint)
	ResetCertificateProviderInstance()
	ResetGrpcEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference
type jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) CertificateProviderInstance() GoogleNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstanceOutputReference {
	var returns GoogleNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstanceOutputReference
	_jsii_.Get(
		j,
		"certificateProviderInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) CertificateProviderInstanceInput() *GoogleNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance {
	var returns *GoogleNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance
	_jsii_.Get(
		j,
		"certificateProviderInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) GrpcEndpoint() GoogleNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointOutputReference {
	var returns GoogleNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointOutputReference
	_jsii_.Get(
		j,
		"grpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) GrpcEndpointInput() *GoogleNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint {
	var returns *GoogleNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint
	_jsii_.Get(
		j,
		"grpcEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) InternalValue() *GoogleNetworkSecurityServerTlsPolicyServerCertificate {
	var returns *GoogleNetworkSecurityServerTlsPolicyServerCertificate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleNetworkSecurityServerTlsPolicy.GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference_Override(g GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleNetworkSecurityServerTlsPolicy.GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference)SetInternalValue(val *GoogleNetworkSecurityServerTlsPolicyServerCertificate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) PutCertificateProviderInstance(value *GoogleNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance) {
	if err := g.validatePutCertificateProviderInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCertificateProviderInstance",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) PutGrpcEndpoint(value *GoogleNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint) {
	if err := g.validatePutGrpcEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGrpcEndpoint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) ResetCertificateProviderInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetCertificateProviderInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) ResetGrpcEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGrpcEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityServerTlsPolicyServerCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

