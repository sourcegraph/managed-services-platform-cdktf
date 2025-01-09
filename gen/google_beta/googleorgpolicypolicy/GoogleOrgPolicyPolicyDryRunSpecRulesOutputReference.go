package googleorgpolicypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleorgpolicypolicy/internal"
)

type GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference interface {
	cdktf.ComplexObject
	AllowAll() *string
	SetAllowAll(val *string)
	AllowAllInput() *string
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
	Condition() GoogleOrgPolicyPolicyDryRunSpecRulesConditionOutputReference
	ConditionInput() *GoogleOrgPolicyPolicyDryRunSpecRulesCondition
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DenyAll() *string
	SetDenyAll(val *string)
	DenyAllInput() *string
	Enforce() *string
	SetEnforce(val *string)
	EnforceInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() GoogleOrgPolicyPolicyDryRunSpecRulesValuesOutputReference
	ValuesInput() *GoogleOrgPolicyPolicyDryRunSpecRulesValues
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
	PutCondition(value *GoogleOrgPolicyPolicyDryRunSpecRulesCondition)
	PutValues(value *GoogleOrgPolicyPolicyDryRunSpecRulesValues)
	ResetAllowAll()
	ResetCondition()
	ResetDenyAll()
	ResetEnforce()
	ResetParameters()
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference
type jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) AllowAll() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) AllowAllInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) Condition() GoogleOrgPolicyPolicyDryRunSpecRulesConditionOutputReference {
	var returns GoogleOrgPolicyPolicyDryRunSpecRulesConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ConditionInput() *GoogleOrgPolicyPolicyDryRunSpecRulesCondition {
	var returns *GoogleOrgPolicyPolicyDryRunSpecRulesCondition
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) DenyAll() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denyAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) DenyAllInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denyAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) Enforce() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) EnforceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) Values() GoogleOrgPolicyPolicyDryRunSpecRulesValuesOutputReference {
	var returns GoogleOrgPolicyPolicyDryRunSpecRulesValuesOutputReference
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ValuesInput() *GoogleOrgPolicyPolicyDryRunSpecRulesValues {
	var returns *GoogleOrgPolicyPolicyDryRunSpecRulesValues
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewGoogleOrgPolicyPolicyDryRunSpecRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOrgPolicyPolicyDryRunSpecRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleOrgPolicyPolicy.GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleOrgPolicyPolicyDryRunSpecRulesOutputReference_Override(g GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleOrgPolicyPolicy.GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference)SetAllowAll(val *string) {
	if err := j.validateSetAllowAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAll",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference)SetDenyAll(val *string) {
	if err := j.validateSetDenyAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denyAll",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference)SetEnforce(val *string) {
	if err := j.validateSetEnforceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforce",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference)SetParameters(val *string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) PutCondition(value *GoogleOrgPolicyPolicyDryRunSpecRulesCondition) {
	if err := g.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCondition",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) PutValues(value *GoogleOrgPolicyPolicyDryRunSpecRulesValues) {
	if err := g.validatePutValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putValues",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ResetAllowAll() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		g,
		"resetCondition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ResetDenyAll() {
	_jsii_.InvokeVoid(
		g,
		"resetDenyAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ResetEnforce() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforce",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		g,
		"resetValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOrgPolicyPolicyDryRunSpecRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

