package aigatewaydynamicrouting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/aigatewaydynamicrouting/internal"
)

type AiGatewayDynamicRoutingElementsPropertiesOutputReference interface {
	cdktf.ComplexObject
	AiGatewayDynamicRoutingProvider() *string
	SetAiGatewayDynamicRoutingProvider(val *string)
	AiGatewayDynamicRoutingProviderInput() *string
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
	Conditions() *string
	SetConditions(val *string)
	ConditionsInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	Limit() *float64
	SetLimit(val *float64)
	LimitInput() *float64
	LimitType() *string
	SetLimitType(val *string)
	LimitTypeInput() *string
	Model() *string
	SetModel(val *string)
	ModelInput() *string
	Retries() *float64
	SetRetries(val *float64)
	RetriesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Window() *float64
	SetWindow(val *float64)
	WindowInput() *float64
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
	ResetAiGatewayDynamicRoutingProvider()
	ResetConditions()
	ResetKey()
	ResetLimit()
	ResetLimitType()
	ResetModel()
	ResetRetries()
	ResetTimeout()
	ResetWindow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiGatewayDynamicRoutingElementsPropertiesOutputReference
type jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) AiGatewayDynamicRoutingProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aiGatewayDynamicRoutingProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) AiGatewayDynamicRoutingProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aiGatewayDynamicRoutingProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) Conditions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ConditionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) LimitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) LimitTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) Model() *string {
	var returns *string
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) Retries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) RetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) Window() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"window",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) WindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"windowInput",
		&returns,
	)
	return returns
}


func NewAiGatewayDynamicRoutingElementsPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AiGatewayDynamicRoutingElementsPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewAiGatewayDynamicRoutingElementsPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiGatewayDynamicRouting.AiGatewayDynamicRoutingElementsPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiGatewayDynamicRoutingElementsPropertiesOutputReference_Override(a AiGatewayDynamicRoutingElementsPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiGatewayDynamicRouting.AiGatewayDynamicRoutingElementsPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetAiGatewayDynamicRoutingProvider(val *string) {
	if err := j.validateSetAiGatewayDynamicRoutingProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aiGatewayDynamicRoutingProvider",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetConditions(val *string) {
	if err := j.validateSetConditionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditions",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetLimitType(val *string) {
	if err := j.validateSetLimitTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitType",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetModel(val *string) {
	if err := j.validateSetModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"model",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetRetries(val *float64) {
	if err := j.validateSetRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retries",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference)SetWindow(val *float64) {
	if err := j.validateSetWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"window",
		val,
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ResetAiGatewayDynamicRoutingProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetAiGatewayDynamicRoutingProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ResetConditions() {
	_jsii_.InvokeVoid(
		a,
		"resetConditions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		a,
		"resetKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ResetLimitType() {
	_jsii_.InvokeVoid(
		a,
		"resetLimitType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ResetModel() {
	_jsii_.InvokeVoid(
		a,
		"resetModel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ResetRetries() {
	_jsii_.InvokeVoid(
		a,
		"resetRetries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ResetWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

