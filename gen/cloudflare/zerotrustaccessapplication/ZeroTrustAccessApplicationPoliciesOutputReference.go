package zerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/zerotrustaccessapplication/internal"
)

type ZeroTrustAccessApplicationPoliciesOutputReference interface {
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
	ConnectionRules() ZeroTrustAccessApplicationPoliciesConnectionRulesOutputReference
	ConnectionRulesInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Decision() *string
	SetDecision(val *string)
	DecisionInput() *string
	Exclude() ZeroTrustAccessApplicationPoliciesExcludeList
	ExcludeInput() any
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Include() ZeroTrustAccessApplicationPoliciesIncludeList
	IncludeInput() any
	InternalValue() any
	SetInternalValue(val any)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Precedence() *float64
	SetPrecedence(val *float64)
	PrecedenceInput() *float64
	Require() ZeroTrustAccessApplicationPoliciesRequireList
	RequireInput() any
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
	PutConnectionRules(value *ZeroTrustAccessApplicationPoliciesConnectionRules)
	PutExclude(value any)
	PutInclude(value any)
	PutRequire(value any)
	ResetConnectionRules()
	ResetDecision()
	ResetExclude()
	ResetId()
	ResetInclude()
	ResetName()
	ResetPrecedence()
	ResetRequire()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustAccessApplicationPoliciesOutputReference
type jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ConnectionRules() ZeroTrustAccessApplicationPoliciesConnectionRulesOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesConnectionRulesOutputReference
	_jsii_.Get(
		j,
		"connectionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ConnectionRulesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"connectionRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) Decision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) DecisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) Exclude() ZeroTrustAccessApplicationPoliciesExcludeList {
	var returns ZeroTrustAccessApplicationPoliciesExcludeList
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ExcludeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) Include() ZeroTrustAccessApplicationPoliciesIncludeList {
	var returns ZeroTrustAccessApplicationPoliciesIncludeList
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) IncludeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) Precedence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) PrecedenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) Require() ZeroTrustAccessApplicationPoliciesRequireList {
	var returns ZeroTrustAccessApplicationPoliciesRequireList
	_jsii_.Get(
		j,
		"require",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) RequireInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requireInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewZeroTrustAccessApplicationPoliciesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustAccessApplicationPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessApplicationPoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustAccessApplicationPoliciesOutputReference_Override(z ZeroTrustAccessApplicationPoliciesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) SetDecision(val *string) {
	if err := j.validateSetDecisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decision",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) SetPrecedence(val *float64) {
	if err := j.validateSetPrecedenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precedence",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) PutConnectionRules(value *ZeroTrustAccessApplicationPoliciesConnectionRules) {
	if err := z.validatePutConnectionRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putConnectionRules",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) PutExclude(value any) {
	if err := z.validatePutExcludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExclude",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) PutInclude(value any) {
	if err := z.validatePutIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putInclude",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) PutRequire(value any) {
	if err := z.validatePutRequireParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putRequire",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ResetConnectionRules() {
	_jsii_.InvokeVoid(
		z,
		"resetConnectionRules",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ResetDecision() {
	_jsii_.InvokeVoid(
		z,
		"resetDecision",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ResetExclude() {
	_jsii_.InvokeVoid(
		z,
		"resetExclude",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		z,
		"resetId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ResetInclude() {
	_jsii_.InvokeVoid(
		z,
		"resetInclude",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		z,
		"resetName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ResetPrecedence() {
	_jsii_.InvokeVoid(
		z,
		"resetPrecedence",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ResetRequire() {
	_jsii_.InvokeVoid(
		z,
		"resetRequire",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := z.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		z,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
