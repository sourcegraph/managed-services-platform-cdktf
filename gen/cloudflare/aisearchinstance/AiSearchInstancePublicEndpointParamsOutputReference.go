package aisearchinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/aisearchinstance/internal"
)

type AiSearchInstancePublicEndpointParamsOutputReference interface {
	cdktf.ComplexObject
	AuthorizedHosts() *[]*string
	SetAuthorizedHosts(val *[]*string)
	AuthorizedHostsInput() *[]*string
	ChatCompletionsEndpoint() AiSearchInstancePublicEndpointParamsChatCompletionsEndpointOutputReference
	ChatCompletionsEndpointInput() interface{}
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Mcp() AiSearchInstancePublicEndpointParamsMcpOutputReference
	McpInput() interface{}
	RateLimit() AiSearchInstancePublicEndpointParamsRateLimitOutputReference
	RateLimitInput() interface{}
	SearchEndpoint() AiSearchInstancePublicEndpointParamsSearchEndpointOutputReference
	SearchEndpointInput() interface{}
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
	PutChatCompletionsEndpoint(value *AiSearchInstancePublicEndpointParamsChatCompletionsEndpoint)
	PutMcp(value *AiSearchInstancePublicEndpointParamsMcp)
	PutRateLimit(value *AiSearchInstancePublicEndpointParamsRateLimit)
	PutSearchEndpoint(value *AiSearchInstancePublicEndpointParamsSearchEndpoint)
	ResetAuthorizedHosts()
	ResetChatCompletionsEndpoint()
	ResetEnabled()
	ResetMcp()
	ResetRateLimit()
	ResetSearchEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiSearchInstancePublicEndpointParamsOutputReference
type jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) AuthorizedHosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) AuthorizedHostsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedHostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ChatCompletionsEndpoint() AiSearchInstancePublicEndpointParamsChatCompletionsEndpointOutputReference {
	var returns AiSearchInstancePublicEndpointParamsChatCompletionsEndpointOutputReference
	_jsii_.Get(
		j,
		"chatCompletionsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ChatCompletionsEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chatCompletionsEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) Mcp() AiSearchInstancePublicEndpointParamsMcpOutputReference {
	var returns AiSearchInstancePublicEndpointParamsMcpOutputReference
	_jsii_.Get(
		j,
		"mcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) McpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) RateLimit() AiSearchInstancePublicEndpointParamsRateLimitOutputReference {
	var returns AiSearchInstancePublicEndpointParamsRateLimitOutputReference
	_jsii_.Get(
		j,
		"rateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) RateLimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) SearchEndpoint() AiSearchInstancePublicEndpointParamsSearchEndpointOutputReference {
	var returns AiSearchInstancePublicEndpointParamsSearchEndpointOutputReference
	_jsii_.Get(
		j,
		"searchEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) SearchEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiSearchInstancePublicEndpointParamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AiSearchInstancePublicEndpointParamsOutputReference {
	_init_.Initialize()

	if err := validateNewAiSearchInstancePublicEndpointParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstancePublicEndpointParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiSearchInstancePublicEndpointParamsOutputReference_Override(a AiSearchInstancePublicEndpointParamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstancePublicEndpointParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference)SetAuthorizedHosts(val *[]*string) {
	if err := j.validateSetAuthorizedHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedHosts",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) PutChatCompletionsEndpoint(value *AiSearchInstancePublicEndpointParamsChatCompletionsEndpoint) {
	if err := a.validatePutChatCompletionsEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putChatCompletionsEndpoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) PutMcp(value *AiSearchInstancePublicEndpointParamsMcp) {
	if err := a.validatePutMcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMcp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) PutRateLimit(value *AiSearchInstancePublicEndpointParamsRateLimit) {
	if err := a.validatePutRateLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRateLimit",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) PutSearchEndpoint(value *AiSearchInstancePublicEndpointParamsSearchEndpoint) {
	if err := a.validatePutSearchEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSearchEndpoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ResetAuthorizedHosts() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizedHosts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ResetChatCompletionsEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetChatCompletionsEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ResetMcp() {
	_jsii_.InvokeVoid(
		a,
		"resetMcp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ResetRateLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetRateLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ResetSearchEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetSearchEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstancePublicEndpointParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

