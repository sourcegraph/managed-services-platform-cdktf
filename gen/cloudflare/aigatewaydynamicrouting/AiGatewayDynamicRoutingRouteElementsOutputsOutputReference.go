package aigatewaydynamicrouting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/aigatewaydynamicrouting/internal"
)

type AiGatewayDynamicRoutingRouteElementsOutputsOutputReference interface {
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
	Fallback() AiGatewayDynamicRoutingRouteElementsOutputsFallbackOutputReference
	False() AiGatewayDynamicRoutingRouteElementsOutputsFalseOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *AiGatewayDynamicRoutingRouteElementsOutputs
	SetInternalValue(val *AiGatewayDynamicRoutingRouteElementsOutputs)
	Next() AiGatewayDynamicRoutingRouteElementsOutputsNextOutputReference
	Success() AiGatewayDynamicRoutingRouteElementsOutputsSuccessOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	True() AiGatewayDynamicRoutingRouteElementsOutputsTrueOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiGatewayDynamicRoutingRouteElementsOutputsOutputReference
type jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) ElementId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) Fallback() AiGatewayDynamicRoutingRouteElementsOutputsFallbackOutputReference {
	var returns AiGatewayDynamicRoutingRouteElementsOutputsFallbackOutputReference
	_jsii_.Get(
		j,
		"fallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) False() AiGatewayDynamicRoutingRouteElementsOutputsFalseOutputReference {
	var returns AiGatewayDynamicRoutingRouteElementsOutputsFalseOutputReference
	_jsii_.Get(
		j,
		"false",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) InternalValue() *AiGatewayDynamicRoutingRouteElementsOutputs {
	var returns *AiGatewayDynamicRoutingRouteElementsOutputs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) Next() AiGatewayDynamicRoutingRouteElementsOutputsNextOutputReference {
	var returns AiGatewayDynamicRoutingRouteElementsOutputsNextOutputReference
	_jsii_.Get(
		j,
		"next",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) Success() AiGatewayDynamicRoutingRouteElementsOutputsSuccessOutputReference {
	var returns AiGatewayDynamicRoutingRouteElementsOutputsSuccessOutputReference
	_jsii_.Get(
		j,
		"success",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) True() AiGatewayDynamicRoutingRouteElementsOutputsTrueOutputReference {
	var returns AiGatewayDynamicRoutingRouteElementsOutputsTrueOutputReference
	_jsii_.Get(
		j,
		"true",
		&returns,
	)
	return returns
}


func NewAiGatewayDynamicRoutingRouteElementsOutputsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AiGatewayDynamicRoutingRouteElementsOutputsOutputReference {
	_init_.Initialize()

	if err := validateNewAiGatewayDynamicRoutingRouteElementsOutputsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiGatewayDynamicRouting.AiGatewayDynamicRoutingRouteElementsOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiGatewayDynamicRoutingRouteElementsOutputsOutputReference_Override(a AiGatewayDynamicRoutingRouteElementsOutputsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiGatewayDynamicRouting.AiGatewayDynamicRoutingRouteElementsOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference)SetInternalValue(val *AiGatewayDynamicRoutingRouteElementsOutputs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiGatewayDynamicRoutingRouteElementsOutputsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

