package teamroutingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/teamroutingrule/internal"
)

type TeamRoutingRuleTimeRestrictionRestrictionOutputReference interface {
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
	EndHour() *float64
	SetEndHour(val *float64)
	EndHourInput() *float64
	EndMin() *float64
	SetEndMin(val *float64)
	EndMinInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *TeamRoutingRuleTimeRestrictionRestriction
	SetInternalValue(val *TeamRoutingRuleTimeRestrictionRestriction)
	StartHour() *float64
	SetStartHour(val *float64)
	StartHourInput() *float64
	StartMin() *float64
	SetStartMin(val *float64)
	StartMinInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TeamRoutingRuleTimeRestrictionRestrictionOutputReference
type jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) EndHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) EndHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) EndMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) EndMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) InternalValue() *TeamRoutingRuleTimeRestrictionRestriction {
	var returns *TeamRoutingRuleTimeRestrictionRestriction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) StartHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) StartHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) StartMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) StartMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTeamRoutingRuleTimeRestrictionRestrictionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TeamRoutingRuleTimeRestrictionRestrictionOutputReference {
	_init_.Initialize()

	if err := validateNewTeamRoutingRuleTimeRestrictionRestrictionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opsgenie.teamRoutingRule.TeamRoutingRuleTimeRestrictionRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTeamRoutingRuleTimeRestrictionRestrictionOutputReference_Override(t TeamRoutingRuleTimeRestrictionRestrictionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opsgenie.teamRoutingRule.TeamRoutingRuleTimeRestrictionRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference)SetEndHour(val *float64) {
	if err := j.validateSetEndHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endHour",
		val,
	)
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference)SetEndMin(val *float64) {
	if err := j.validateSetEndMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endMin",
		val,
	)
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference)SetInternalValue(val *TeamRoutingRuleTimeRestrictionRestriction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference)SetStartHour(val *float64) {
	if err := j.validateSetStartHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startHour",
		val,
	)
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference)SetStartMin(val *float64) {
	if err := j.validateSetStartMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startMin",
		val,
	)
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamRoutingRuleTimeRestrictionRestrictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

