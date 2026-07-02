package googlecomputeregionsecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputeregionsecuritypolicyrule/internal"
)

type GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference interface {
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
	RequestCookie() GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList
	RequestCookieInput() any
	RequestHeader() GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderList
	RequestHeaderInput() any
	RequestQueryParam() GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList
	RequestQueryParamInput() any
	RequestUri() GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriList
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

// The jsii proxy struct for GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference
type jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestCookie() GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList {
	var returns GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieList
	_jsii_.Get(
		j,
		"requestCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestCookieInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestHeader() GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderList {
	var returns GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderList
	_jsii_.Get(
		j,
		"requestHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestHeaderInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestQueryParam() GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList {
	var returns GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamList
	_jsii_.Get(
		j,
		"requestQueryParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestQueryParamInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestQueryParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestUri() GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriList {
	var returns GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriList
	_jsii_.Get(
		j,
		"requestUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) RequestUriInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requestUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TargetRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TargetRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TargetRuleSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TargetRuleSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeRegionSecurityPolicyRule.GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference_Override(g GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeRegionSecurityPolicyRule.GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetTargetRuleIds(val *[]*string) {
	if err := j.validateSetTargetRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRuleIds",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetTargetRuleSet(val *string) {
	if err := j.validateSetTargetRuleSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRuleSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) PutRequestCookie(value any) {
	if err := g.validatePutRequestCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestCookie",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) PutRequestHeader(value any) {
	if err := g.validatePutRequestHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestHeader",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) PutRequestQueryParam(value any) {
	if err := g.validatePutRequestQueryParamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestQueryParam",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) PutRequestUri(value any) {
	if err := g.validatePutRequestUriParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestUri",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ResetRequestCookie() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestCookie",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ResetRequestHeader() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestHeader",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ResetRequestQueryParam() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestQueryParam",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ResetRequestUri() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ResetTargetRuleIds() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetRuleIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulePreconfiguredWafConfigExclusionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
