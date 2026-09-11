package googlecomputeregionsecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputeregionsecuritypolicy/internal"
)

type GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference interface {
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
	InternalValue() any
	SetInternalValue(val any)
	RequestCookie() GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestCookieList
	RequestCookieInput() any
	RequestHeader() GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestHeaderList
	RequestHeaderInput() any
	RequestQueryParam() GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestQueryParamList
	RequestQueryParamInput() any
	RequestUri() GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestUriList
	RequestUriInput() any
	TargetRuleIds() *[]*string
	SetTargetRuleIds(val *[]*string)
	TargetRuleIdsInput() *[]*string
	TargetRuleSet() *string
	SetTargetRuleSet(val *string)
	TargetRuleSetInput() *string
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
	PutRequestCookie(value any)
	PutRequestHeader(value any)
	PutRequestQueryParam(value any)
	PutRequestUri(value any)
	ResetRequestCookie()
	ResetRequestHeader()
	ResetRequestQueryParam()
	ResetRequestUri()
	ResetTargetRuleIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference
type jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) RequestCookie() GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestCookieList {
	var returns GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestCookieList
	_jsii_.Get(
		j,
		"requestCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) RequestCookieInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) RequestHeader() GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestHeaderList {
	var returns GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestHeaderList
	_jsii_.Get(
		j,
		"requestHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) RequestHeaderInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) RequestQueryParam() GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestQueryParamList {
	var returns GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestQueryParamList
	_jsii_.Get(
		j,
		"requestQueryParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) RequestQueryParamInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestQueryParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) RequestUri() GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestUriList {
	var returns GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestUriList
	_jsii_.Get(
		j,
		"requestUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) RequestUriInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) TargetRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) TargetRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) TargetRuleSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) TargetRuleSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeRegionSecurityPolicy.GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference_Override(g GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeRegionSecurityPolicy.GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) SetTargetRuleIds(val *[]*string) {
	if err := j.validateSetTargetRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRuleIds",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) SetTargetRuleSet(val *string) {
	if err := j.validateSetTargetRuleSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRuleSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) PutRequestCookie(value any) {
	if err := g.validatePutRequestCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestCookie",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) PutRequestHeader(value any) {
	if err := g.validatePutRequestHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestHeader",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) PutRequestQueryParam(value any) {
	if err := g.validatePutRequestQueryParamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestQueryParam",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) PutRequestUri(value any) {
	if err := g.validatePutRequestUriParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestUri",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) ResetRequestCookie() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestCookie",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) ResetRequestHeader() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestHeader",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) ResetRequestQueryParam() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestQueryParam",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) ResetRequestUri() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) ResetTargetRuleIds() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetRuleIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
