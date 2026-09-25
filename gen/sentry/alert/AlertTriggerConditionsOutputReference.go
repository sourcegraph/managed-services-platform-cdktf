package alert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/alert/internal"
)

type AlertTriggerConditionsOutputReference interface {
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
	EventFrequencyCount() AlertTriggerConditionsEventFrequencyCountOutputReference
	EventFrequencyCountInput() interface{}
	FirstSeenEvent() AlertTriggerConditionsFirstSeenEventOutputReference
	FirstSeenEventInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IssueResolvedTrigger() AlertTriggerConditionsIssueResolvedTriggerOutputReference
	IssueResolvedTriggerInput() interface{}
	ReappearedEvent() AlertTriggerConditionsReappearedEventOutputReference
	ReappearedEventInput() interface{}
	RegressionEvent() AlertTriggerConditionsRegressionEventOutputReference
	RegressionEventInput() interface{}
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
	PutEventFrequencyCount(value *AlertTriggerConditionsEventFrequencyCount)
	PutFirstSeenEvent(value *AlertTriggerConditionsFirstSeenEvent)
	PutIssueResolvedTrigger(value *AlertTriggerConditionsIssueResolvedTrigger)
	PutReappearedEvent(value *AlertTriggerConditionsReappearedEvent)
	PutRegressionEvent(value *AlertTriggerConditionsRegressionEvent)
	ResetEventFrequencyCount()
	ResetFirstSeenEvent()
	ResetIssueResolvedTrigger()
	ResetReappearedEvent()
	ResetRegressionEvent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertTriggerConditionsOutputReference
type jsiiProxy_AlertTriggerConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) EventFrequencyCount() AlertTriggerConditionsEventFrequencyCountOutputReference {
	var returns AlertTriggerConditionsEventFrequencyCountOutputReference
	_jsii_.Get(
		j,
		"eventFrequencyCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) EventFrequencyCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventFrequencyCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) FirstSeenEvent() AlertTriggerConditionsFirstSeenEventOutputReference {
	var returns AlertTriggerConditionsFirstSeenEventOutputReference
	_jsii_.Get(
		j,
		"firstSeenEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) FirstSeenEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstSeenEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) IssueResolvedTrigger() AlertTriggerConditionsIssueResolvedTriggerOutputReference {
	var returns AlertTriggerConditionsIssueResolvedTriggerOutputReference
	_jsii_.Get(
		j,
		"issueResolvedTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) IssueResolvedTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueResolvedTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) ReappearedEvent() AlertTriggerConditionsReappearedEventOutputReference {
	var returns AlertTriggerConditionsReappearedEventOutputReference
	_jsii_.Get(
		j,
		"reappearedEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) ReappearedEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reappearedEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) RegressionEvent() AlertTriggerConditionsRegressionEventOutputReference {
	var returns AlertTriggerConditionsRegressionEventOutputReference
	_jsii_.Get(
		j,
		"regressionEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) RegressionEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regressionEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlertTriggerConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlertTriggerConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewAlertTriggerConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertTriggerConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.alert.AlertTriggerConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlertTriggerConditionsOutputReference_Override(a AlertTriggerConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.alert.AlertTriggerConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertTriggerConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) PutEventFrequencyCount(value *AlertTriggerConditionsEventFrequencyCount) {
	if err := a.validatePutEventFrequencyCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventFrequencyCount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) PutFirstSeenEvent(value *AlertTriggerConditionsFirstSeenEvent) {
	if err := a.validatePutFirstSeenEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirstSeenEvent",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) PutIssueResolvedTrigger(value *AlertTriggerConditionsIssueResolvedTrigger) {
	if err := a.validatePutIssueResolvedTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIssueResolvedTrigger",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) PutReappearedEvent(value *AlertTriggerConditionsReappearedEvent) {
	if err := a.validatePutReappearedEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReappearedEvent",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) PutRegressionEvent(value *AlertTriggerConditionsRegressionEvent) {
	if err := a.validatePutRegressionEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegressionEvent",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) ResetEventFrequencyCount() {
	_jsii_.InvokeVoid(
		a,
		"resetEventFrequencyCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) ResetFirstSeenEvent() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstSeenEvent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) ResetIssueResolvedTrigger() {
	_jsii_.InvokeVoid(
		a,
		"resetIssueResolvedTrigger",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) ResetReappearedEvent() {
	_jsii_.InvokeVoid(
		a,
		"resetReappearedEvent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) ResetRegressionEvent() {
	_jsii_.InvokeVoid(
		a,
		"resetRegressionEvent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlertTriggerConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

