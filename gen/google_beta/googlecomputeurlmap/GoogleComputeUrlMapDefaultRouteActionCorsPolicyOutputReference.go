package googlecomputeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputeurlmap/internal"
)

type GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference interface {
	cdktf.ComplexObject
	AllowCredentials() interface{}
	SetAllowCredentials(val interface{})
	AllowCredentialsInput() interface{}
	AllowHeaders() *[]*string
	SetAllowHeaders(val *[]*string)
	AllowHeadersInput() *[]*string
	AllowMethods() *[]*string
	SetAllowMethods(val *[]*string)
	AllowMethodsInput() *[]*string
	AllowOriginRegexes() *[]*string
	SetAllowOriginRegexes(val *[]*string)
	AllowOriginRegexesInput() *[]*string
	AllowOrigins() *[]*string
	SetAllowOrigins(val *[]*string)
	AllowOriginsInput() *[]*string
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
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	ExposeHeaders() *[]*string
	SetExposeHeaders(val *[]*string)
	ExposeHeadersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeUrlMapDefaultRouteActionCorsPolicy
	SetInternalValue(val *GoogleComputeUrlMapDefaultRouteActionCorsPolicy)
	MaxAge() *float64
	SetMaxAge(val *float64)
	MaxAgeInput() *float64
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
	ResetAllowCredentials()
	ResetAllowHeaders()
	ResetAllowMethods()
	ResetAllowOriginRegexes()
	ResetAllowOrigins()
	ResetDisabled()
	ResetExposeHeaders()
	ResetMaxAge()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference
type jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) AllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) AllowCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) AllowHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) AllowHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) AllowMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) AllowMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) AllowOriginRegexes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOriginRegexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) AllowOriginRegexesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOriginRegexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) AllowOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) AllowOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ExposeHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposeHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ExposeHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposeHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) InternalValue() *GoogleComputeUrlMapDefaultRouteActionCorsPolicy {
	var returns *GoogleComputeUrlMapDefaultRouteActionCorsPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) MaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) MaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeUrlMap.GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference_Override(g GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeUrlMap.GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetAllowCredentials(val interface{}) {
	if err := j.validateSetAllowCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCredentials",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetAllowHeaders(val *[]*string) {
	if err := j.validateSetAllowHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetAllowMethods(val *[]*string) {
	if err := j.validateSetAllowMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMethods",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetAllowOriginRegexes(val *[]*string) {
	if err := j.validateSetAllowOriginRegexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOriginRegexes",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetAllowOrigins(val *[]*string) {
	if err := j.validateSetAllowOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOrigins",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetExposeHeaders(val *[]*string) {
	if err := j.validateSetExposeHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exposeHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetInternalValue(val *GoogleComputeUrlMapDefaultRouteActionCorsPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetMaxAge(val *float64) {
	if err := j.validateSetMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAge",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ResetAllowCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ResetAllowHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ResetAllowMethods() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowMethods",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ResetAllowOriginRegexes() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowOriginRegexes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ResetAllowOrigins() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowOrigins",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ResetExposeHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetExposeHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ResetMaxAge() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxAge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapDefaultRouteActionCorsPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

