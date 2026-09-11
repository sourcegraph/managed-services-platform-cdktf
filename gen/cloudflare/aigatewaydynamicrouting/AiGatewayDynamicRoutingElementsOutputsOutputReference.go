package aigatewaydynamicrouting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/aigatewaydynamicrouting/internal"
)

type AiGatewayDynamicRoutingElementsOutputsOutputReference interface {
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
	ElementId() *string
	SetElementId(val *string)
	ElementIdInput() *string
	Fallback() AiGatewayDynamicRoutingElementsOutputsFallbackOutputReference
	FallbackInput() interface{}
	False() AiGatewayDynamicRoutingElementsOutputsFalseOutputReference
	FalseInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Next() AiGatewayDynamicRoutingElementsOutputsNextOutputReference
	NextInput() interface{}
	Success() AiGatewayDynamicRoutingElementsOutputsSuccessOutputReference
	SuccessInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	True() AiGatewayDynamicRoutingElementsOutputsTrueOutputReference
	TrueInput() interface{}
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
	PutFallback(value *AiGatewayDynamicRoutingElementsOutputsFallback)
	PutFalse(value *AiGatewayDynamicRoutingElementsOutputsFalse)
	PutNext(value *AiGatewayDynamicRoutingElementsOutputsNext)
	PutSuccess(value *AiGatewayDynamicRoutingElementsOutputsSuccess)
	PutTrue(value *AiGatewayDynamicRoutingElementsOutputsTrue)
	ResetElementId()
	ResetFallback()
	ResetFalse()
	ResetNext()
	ResetSuccess()
	ResetTrue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiGatewayDynamicRoutingElementsOutputsOutputReference
type jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ElementId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ElementIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) Fallback() AiGatewayDynamicRoutingElementsOutputsFallbackOutputReference {
	var returns AiGatewayDynamicRoutingElementsOutputsFallbackOutputReference
	_jsii_.Get(
		j,
		"fallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) FallbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) False() AiGatewayDynamicRoutingElementsOutputsFalseOutputReference {
	var returns AiGatewayDynamicRoutingElementsOutputsFalseOutputReference
	_jsii_.Get(
		j,
		"false",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) FalseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"falseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) Next() AiGatewayDynamicRoutingElementsOutputsNextOutputReference {
	var returns AiGatewayDynamicRoutingElementsOutputsNextOutputReference
	_jsii_.Get(
		j,
		"next",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) NextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) Success() AiGatewayDynamicRoutingElementsOutputsSuccessOutputReference {
	var returns AiGatewayDynamicRoutingElementsOutputsSuccessOutputReference
	_jsii_.Get(
		j,
		"success",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) SuccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) True() AiGatewayDynamicRoutingElementsOutputsTrueOutputReference {
	var returns AiGatewayDynamicRoutingElementsOutputsTrueOutputReference
	_jsii_.Get(
		j,
		"true",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) TrueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trueInput",
		&returns,
	)
	return returns
}


func NewAiGatewayDynamicRoutingElementsOutputsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AiGatewayDynamicRoutingElementsOutputsOutputReference {
	_init_.Initialize()

	if err := validateNewAiGatewayDynamicRoutingElementsOutputsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiGatewayDynamicRouting.AiGatewayDynamicRoutingElementsOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiGatewayDynamicRoutingElementsOutputsOutputReference_Override(a AiGatewayDynamicRoutingElementsOutputsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiGatewayDynamicRouting.AiGatewayDynamicRoutingElementsOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference)SetElementId(val *string) {
	if err := j.validateSetElementIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elementId",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) PutFallback(value *AiGatewayDynamicRoutingElementsOutputsFallback) {
	if err := a.validatePutFallbackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFallback",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) PutFalse(value *AiGatewayDynamicRoutingElementsOutputsFalse) {
	if err := a.validatePutFalseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFalse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) PutNext(value *AiGatewayDynamicRoutingElementsOutputsNext) {
	if err := a.validatePutNextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNext",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) PutSuccess(value *AiGatewayDynamicRoutingElementsOutputsSuccess) {
	if err := a.validatePutSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSuccess",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) PutTrue(value *AiGatewayDynamicRoutingElementsOutputsTrue) {
	if err := a.validatePutTrueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrue",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ResetElementId() {
	_jsii_.InvokeVoid(
		a,
		"resetElementId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ResetFallback() {
	_jsii_.InvokeVoid(
		a,
		"resetFallback",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ResetFalse() {
	_jsii_.InvokeVoid(
		a,
		"resetFalse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ResetNext() {
	_jsii_.InvokeVoid(
		a,
		"resetNext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ResetSuccess() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ResetTrue() {
	_jsii_.InvokeVoid(
		a,
		"resetTrue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingElementsOutputsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

