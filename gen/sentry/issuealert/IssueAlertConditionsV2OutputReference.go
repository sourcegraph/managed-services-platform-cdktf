package issuealert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/issuealert/internal"
)

type IssueAlertConditionsV2OutputReference interface {
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
	EventFrequency() IssueAlertConditionsV2EventFrequencyOutputReference
	EventFrequencyInput() interface{}
	EventFrequencyPercent() IssueAlertConditionsV2EventFrequencyPercentOutputReference
	EventFrequencyPercentInput() interface{}
	EventUniqueUserFrequency() IssueAlertConditionsV2EventUniqueUserFrequencyOutputReference
	EventUniqueUserFrequencyInput() interface{}
	ExistingHighPriorityIssue() IssueAlertConditionsV2ExistingHighPriorityIssueOutputReference
	ExistingHighPriorityIssueInput() interface{}
	FirstSeenEvent() IssueAlertConditionsV2FirstSeenEventOutputReference
	FirstSeenEventInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NewHighPriorityIssue() IssueAlertConditionsV2NewHighPriorityIssueOutputReference
	NewHighPriorityIssueInput() interface{}
	ReappearedEvent() IssueAlertConditionsV2ReappearedEventOutputReference
	ReappearedEventInput() interface{}
	RegressionEvent() IssueAlertConditionsV2RegressionEventOutputReference
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
	PutEventFrequency(value *IssueAlertConditionsV2EventFrequency)
	PutEventFrequencyPercent(value *IssueAlertConditionsV2EventFrequencyPercent)
	PutEventUniqueUserFrequency(value *IssueAlertConditionsV2EventUniqueUserFrequency)
	PutExistingHighPriorityIssue(value *IssueAlertConditionsV2ExistingHighPriorityIssue)
	PutFirstSeenEvent(value *IssueAlertConditionsV2FirstSeenEvent)
	PutNewHighPriorityIssue(value *IssueAlertConditionsV2NewHighPriorityIssue)
	PutReappearedEvent(value *IssueAlertConditionsV2ReappearedEvent)
	PutRegressionEvent(value *IssueAlertConditionsV2RegressionEvent)
	ResetEventFrequency()
	ResetEventFrequencyPercent()
	ResetEventUniqueUserFrequency()
	ResetExistingHighPriorityIssue()
	ResetFirstSeenEvent()
	ResetNewHighPriorityIssue()
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

// The jsii proxy struct for IssueAlertConditionsV2OutputReference
type jsiiProxy_IssueAlertConditionsV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) EventFrequency() IssueAlertConditionsV2EventFrequencyOutputReference {
	var returns IssueAlertConditionsV2EventFrequencyOutputReference
	_jsii_.Get(
		j,
		"eventFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) EventFrequencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) EventFrequencyPercent() IssueAlertConditionsV2EventFrequencyPercentOutputReference {
	var returns IssueAlertConditionsV2EventFrequencyPercentOutputReference
	_jsii_.Get(
		j,
		"eventFrequencyPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) EventFrequencyPercentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventFrequencyPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) EventUniqueUserFrequency() IssueAlertConditionsV2EventUniqueUserFrequencyOutputReference {
	var returns IssueAlertConditionsV2EventUniqueUserFrequencyOutputReference
	_jsii_.Get(
		j,
		"eventUniqueUserFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) EventUniqueUserFrequencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventUniqueUserFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) ExistingHighPriorityIssue() IssueAlertConditionsV2ExistingHighPriorityIssueOutputReference {
	var returns IssueAlertConditionsV2ExistingHighPriorityIssueOutputReference
	_jsii_.Get(
		j,
		"existingHighPriorityIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) ExistingHighPriorityIssueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"existingHighPriorityIssueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) FirstSeenEvent() IssueAlertConditionsV2FirstSeenEventOutputReference {
	var returns IssueAlertConditionsV2FirstSeenEventOutputReference
	_jsii_.Get(
		j,
		"firstSeenEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) FirstSeenEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstSeenEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) NewHighPriorityIssue() IssueAlertConditionsV2NewHighPriorityIssueOutputReference {
	var returns IssueAlertConditionsV2NewHighPriorityIssueOutputReference
	_jsii_.Get(
		j,
		"newHighPriorityIssue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) NewHighPriorityIssueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newHighPriorityIssueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) ReappearedEvent() IssueAlertConditionsV2ReappearedEventOutputReference {
	var returns IssueAlertConditionsV2ReappearedEventOutputReference
	_jsii_.Get(
		j,
		"reappearedEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) ReappearedEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reappearedEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) RegressionEvent() IssueAlertConditionsV2RegressionEventOutputReference {
	var returns IssueAlertConditionsV2RegressionEventOutputReference
	_jsii_.Get(
		j,
		"regressionEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) RegressionEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regressionEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIssueAlertConditionsV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IssueAlertConditionsV2OutputReference {
	_init_.Initialize()

	if err := validateNewIssueAlertConditionsV2OutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IssueAlertConditionsV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.issueAlert.IssueAlertConditionsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIssueAlertConditionsV2OutputReference_Override(i IssueAlertConditionsV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.issueAlert.IssueAlertConditionsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IssueAlertConditionsV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) PutEventFrequency(value *IssueAlertConditionsV2EventFrequency) {
	if err := i.validatePutEventFrequencyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEventFrequency",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) PutEventFrequencyPercent(value *IssueAlertConditionsV2EventFrequencyPercent) {
	if err := i.validatePutEventFrequencyPercentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEventFrequencyPercent",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) PutEventUniqueUserFrequency(value *IssueAlertConditionsV2EventUniqueUserFrequency) {
	if err := i.validatePutEventUniqueUserFrequencyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEventUniqueUserFrequency",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) PutExistingHighPriorityIssue(value *IssueAlertConditionsV2ExistingHighPriorityIssue) {
	if err := i.validatePutExistingHighPriorityIssueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putExistingHighPriorityIssue",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) PutFirstSeenEvent(value *IssueAlertConditionsV2FirstSeenEvent) {
	if err := i.validatePutFirstSeenEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirstSeenEvent",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) PutNewHighPriorityIssue(value *IssueAlertConditionsV2NewHighPriorityIssue) {
	if err := i.validatePutNewHighPriorityIssueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNewHighPriorityIssue",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) PutReappearedEvent(value *IssueAlertConditionsV2ReappearedEvent) {
	if err := i.validatePutReappearedEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putReappearedEvent",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) PutRegressionEvent(value *IssueAlertConditionsV2RegressionEvent) {
	if err := i.validatePutRegressionEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRegressionEvent",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) ResetEventFrequency() {
	_jsii_.InvokeVoid(
		i,
		"resetEventFrequency",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) ResetEventFrequencyPercent() {
	_jsii_.InvokeVoid(
		i,
		"resetEventFrequencyPercent",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) ResetEventUniqueUserFrequency() {
	_jsii_.InvokeVoid(
		i,
		"resetEventUniqueUserFrequency",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) ResetExistingHighPriorityIssue() {
	_jsii_.InvokeVoid(
		i,
		"resetExistingHighPriorityIssue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) ResetFirstSeenEvent() {
	_jsii_.InvokeVoid(
		i,
		"resetFirstSeenEvent",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) ResetNewHighPriorityIssue() {
	_jsii_.InvokeVoid(
		i,
		"resetNewHighPriorityIssue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) ResetReappearedEvent() {
	_jsii_.InvokeVoid(
		i,
		"resetReappearedEvent",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) ResetRegressionEvent() {
	_jsii_.InvokeVoid(
		i,
		"resetRegressionEvent",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertConditionsV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

