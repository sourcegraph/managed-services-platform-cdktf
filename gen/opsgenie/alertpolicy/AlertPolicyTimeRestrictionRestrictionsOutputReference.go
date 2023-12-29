package alertpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/alertpolicy/internal"
)

type AlertPolicyTimeRestrictionRestrictionsOutputReference interface {
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
	EndDay() *string
	SetEndDay(val *string)
	EndDayInput() *string
	EndHour() *float64
	SetEndHour(val *float64)
	EndHourInput() *float64
	EndMin() *float64
	SetEndMin(val *float64)
	EndMinInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StartDay() *string
	SetStartDay(val *string)
	StartDayInput() *string
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

// The jsii proxy struct for AlertPolicyTimeRestrictionRestrictionsOutputReference
type jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) EndDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) EndDayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) EndHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) EndHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) EndMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) EndMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) StartDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) StartDayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) StartHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) StartHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) StartMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) StartMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlertPolicyTimeRestrictionRestrictionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlertPolicyTimeRestrictionRestrictionsOutputReference {
	_init_.Initialize()

	if err := validateNewAlertPolicyTimeRestrictionRestrictionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicyTimeRestrictionRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlertPolicyTimeRestrictionRestrictionsOutputReference_Override(a AlertPolicyTimeRestrictionRestrictionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicyTimeRestrictionRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference)SetEndDay(val *string) {
	if err := j.validateSetEndDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endDay",
		val,
	)
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference)SetEndHour(val *float64) {
	if err := j.validateSetEndHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endHour",
		val,
	)
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference)SetEndMin(val *float64) {
	if err := j.validateSetEndMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endMin",
		val,
	)
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference)SetStartDay(val *string) {
	if err := j.validateSetStartDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startDay",
		val,
	)
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference)SetStartHour(val *float64) {
	if err := j.validateSetStartHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startHour",
		val,
	)
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference)SetStartMin(val *float64) {
	if err := j.validateSetStartMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startMin",
		val,
	)
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlertPolicyTimeRestrictionRestrictionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

