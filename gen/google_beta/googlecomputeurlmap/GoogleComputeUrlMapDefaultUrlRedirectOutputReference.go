package googlecomputeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputeurlmap/internal"
)

type GoogleComputeUrlMapDefaultUrlRedirectOutputReference interface {
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
	HostRedirect() *string
	SetHostRedirect(val *string)
	HostRedirectInput() *string
	HttpsRedirect() interface{}
	SetHttpsRedirect(val interface{})
	HttpsRedirectInput() interface{}
	InternalValue() *GoogleComputeUrlMapDefaultUrlRedirect
	SetInternalValue(val *GoogleComputeUrlMapDefaultUrlRedirect)
	PathRedirect() *string
	SetPathRedirect(val *string)
	PathRedirectInput() *string
	PrefixRedirect() *string
	SetPrefixRedirect(val *string)
	PrefixRedirectInput() *string
	RedirectResponseCode() *string
	SetRedirectResponseCode(val *string)
	RedirectResponseCodeInput() *string
	StripQuery() interface{}
	SetStripQuery(val interface{})
	StripQueryInput() interface{}
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
	ResetHostRedirect()
	ResetHttpsRedirect()
	ResetPathRedirect()
	ResetPrefixRedirect()
	ResetRedirectResponseCode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeUrlMapDefaultUrlRedirectOutputReference
type jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) HostRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) HostRedirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) HttpsRedirect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) HttpsRedirectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) InternalValue() *GoogleComputeUrlMapDefaultUrlRedirect {
	var returns *GoogleComputeUrlMapDefaultUrlRedirect
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) PathRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) PathRedirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) PrefixRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) PrefixRedirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) RedirectResponseCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectResponseCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) RedirectResponseCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectResponseCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) StripQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) StripQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeUrlMapDefaultUrlRedirectOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeUrlMapDefaultUrlRedirectOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeUrlMapDefaultUrlRedirectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeUrlMap.GoogleComputeUrlMapDefaultUrlRedirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeUrlMapDefaultUrlRedirectOutputReference_Override(g GoogleComputeUrlMapDefaultUrlRedirectOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeUrlMap.GoogleComputeUrlMapDefaultUrlRedirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference)SetHostRedirect(val *string) {
	if err := j.validateSetHostRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostRedirect",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference)SetHttpsRedirect(val interface{}) {
	if err := j.validateSetHttpsRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsRedirect",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference)SetInternalValue(val *GoogleComputeUrlMapDefaultUrlRedirect) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference)SetPathRedirect(val *string) {
	if err := j.validateSetPathRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathRedirect",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference)SetPrefixRedirect(val *string) {
	if err := j.validateSetPrefixRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixRedirect",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference)SetRedirectResponseCode(val *string) {
	if err := j.validateSetRedirectResponseCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectResponseCode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference)SetStripQuery(val interface{}) {
	if err := j.validateSetStripQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stripQuery",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) ResetHostRedirect() {
	_jsii_.InvokeVoid(
		g,
		"resetHostRedirect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) ResetHttpsRedirect() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpsRedirect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) ResetPathRedirect() {
	_jsii_.InvokeVoid(
		g,
		"resetPathRedirect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) ResetPrefixRedirect() {
	_jsii_.InvokeVoid(
		g,
		"resetPrefixRedirect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) ResetRedirectResponseCode() {
	_jsii_.InvokeVoid(
		g,
		"resetRedirectResponseCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultUrlRedirectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

