package dataopsgenieescalation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/dataopsgenieescalation/internal"
)

type DataOpsgenieEscalationRepeatOutputReference interface {
	cdktf.ComplexObject
	CloseAlertAfterAll() any
	SetCloseAlertAfterAll(val any)
	CloseAlertAfterAllInput() any
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	ResetRecipientStates() any
	SetResetRecipientStates(val any)
	ResetRecipientStatesInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WaitInterval() *float64
	SetWaitInterval(val *float64)
	WaitIntervalInput() *float64
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
	ResetCloseAlertAfterAll()
	ResetCount()
	ResetResetRecipientStates()
	ResetWaitInterval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataOpsgenieEscalationRepeatOutputReference
type jsiiProxy_DataOpsgenieEscalationRepeatOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) CloseAlertAfterAll() any {
	var returns any
	_jsii_.Get(
		j,
		"closeAlertAfterAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) CloseAlertAfterAllInput() any {
	var returns any
	_jsii_.Get(
		j,
		"closeAlertAfterAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) ResetRecipientStates() any {
	var returns any
	_jsii_.Get(
		j,
		"resetRecipientStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) ResetRecipientStatesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"resetRecipientStatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) WaitInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) WaitIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIntervalInput",
		&returns,
	)
	return returns
}

func NewDataOpsgenieEscalationRepeatOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataOpsgenieEscalationRepeatOutputReference {
	_init_.Initialize()

	if err := validateNewDataOpsgenieEscalationRepeatOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOpsgenieEscalationRepeatOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opsgenie.dataOpsgenieEscalation.DataOpsgenieEscalationRepeatOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataOpsgenieEscalationRepeatOutputReference_Override(d DataOpsgenieEscalationRepeatOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opsgenie.dataOpsgenieEscalation.DataOpsgenieEscalationRepeatOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) SetCloseAlertAfterAll(val any) {
	if err := j.validateSetCloseAlertAfterAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"closeAlertAfterAll",
		val,
	)
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) SetCount(val *float64) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) SetResetRecipientStates(val any) {
	if err := j.validateSetResetRecipientStatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resetRecipientStates",
		val,
	)
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) SetWaitInterval(val *float64) {
	if err := j.validateSetWaitIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitInterval",
		val,
	)
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) ResetCloseAlertAfterAll() {
	_jsii_.InvokeVoid(
		d,
		"resetCloseAlertAfterAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		d,
		"resetCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) ResetResetRecipientStates() {
	_jsii_.InvokeVoid(
		d,
		"resetResetRecipientStates",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) ResetWaitInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetWaitInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOpsgenieEscalationRepeatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
