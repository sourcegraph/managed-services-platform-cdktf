package googleorgpolicypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleorgpolicypolicy/internal"
)

type GoogleOrgPolicyPolicyDryRunSpecOutputReference interface {
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
	Etag() *string
	// Experimental.
	Fqn() *string
	InheritFromParent() any
	SetInheritFromParent(val any)
	InheritFromParentInput() any
	InternalValue() *GoogleOrgPolicyPolicyDryRunSpec
	SetInternalValue(val *GoogleOrgPolicyPolicyDryRunSpec)
	Reset() any
	SetReset(val any)
	ResetInput() any
	Rules() GoogleOrgPolicyPolicyDryRunSpecRulesList
	RulesInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdateTime() *string
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
	PutRules(value any)
	ResetInheritFromParent()
	ResetReset()
	ResetRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOrgPolicyPolicyDryRunSpecOutputReference
type jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) InheritFromParent() any {
	var returns any
	_jsii_.Get(
		j,
		"inheritFromParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) InheritFromParentInput() any {
	var returns any
	_jsii_.Get(
		j,
		"inheritFromParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) InternalValue() *GoogleOrgPolicyPolicyDryRunSpec {
	var returns *GoogleOrgPolicyPolicyDryRunSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) Reset() any {
	var returns any
	_jsii_.Get(
		j,
		"reset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) ResetInput() any {
	var returns any
	_jsii_.Get(
		j,
		"resetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) Rules() GoogleOrgPolicyPolicyDryRunSpecRulesList {
	var returns GoogleOrgPolicyPolicyDryRunSpecRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) RulesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func NewGoogleOrgPolicyPolicyDryRunSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOrgPolicyPolicyDryRunSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOrgPolicyPolicyDryRunSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleOrgPolicyPolicy.GoogleOrgPolicyPolicyDryRunSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOrgPolicyPolicyDryRunSpecOutputReference_Override(g GoogleOrgPolicyPolicyDryRunSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleOrgPolicyPolicy.GoogleOrgPolicyPolicyDryRunSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) SetInheritFromParent(val any) {
	if err := j.validateSetInheritFromParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inheritFromParent",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) SetInternalValue(val *GoogleOrgPolicyPolicyDryRunSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) SetReset(val any) {
	if err := j.validateSetResetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reset",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) PutRules(value any) {
	if err := g.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRules",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) ResetInheritFromParent() {
	_jsii_.InvokeVoid(
		g,
		"resetInheritFromParent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) ResetReset() {
	_jsii_.InvokeVoid(
		g,
		"resetReset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) ResetRules() {
	_jsii_.InvokeVoid(
		g,
		"resetRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
