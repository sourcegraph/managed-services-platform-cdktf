package metricmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/metricmonitor/internal"
)

type MetricMonitorConditionGroupConditionsOutputReference interface {
	cdktf.ComplexObject
	Comparison() *float64
	SetComparison(val *float64)
	ComparisonInput() *float64
	ComparisonSensitivity() *string
	SetComparisonSensitivity(val *string)
	ComparisonSensitivityInput() *string
	ComparisonThresholdType() *string
	SetComparisonThresholdType(val *string)
	ComparisonThresholdTypeInput() *string
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
	ConditionResult() *float64
	SetConditionResult(val *float64)
	ConditionResultInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetComparison()
	ResetComparisonSensitivity()
	ResetComparisonThresholdType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MetricMonitorConditionGroupConditionsOutputReference
type jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) Comparison() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"comparison",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ComparisonInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"comparisonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ComparisonSensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ComparisonSensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ComparisonThresholdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonThresholdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ComparisonThresholdTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonThresholdTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ConditionResult() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"conditionResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ConditionResultInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"conditionResultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewMetricMonitorConditionGroupConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MetricMonitorConditionGroupConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewMetricMonitorConditionGroupConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.metricMonitor.MetricMonitorConditionGroupConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMetricMonitorConditionGroupConditionsOutputReference_Override(m MetricMonitorConditionGroupConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.metricMonitor.MetricMonitorConditionGroupConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference)SetComparison(val *float64) {
	if err := j.validateSetComparisonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparison",
		val,
	)
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference)SetComparisonSensitivity(val *string) {
	if err := j.validateSetComparisonSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonSensitivity",
		val,
	)
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference)SetComparisonThresholdType(val *string) {
	if err := j.validateSetComparisonThresholdTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonThresholdType",
		val,
	)
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference)SetConditionResult(val *float64) {
	if err := j.validateSetConditionResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionResult",
		val,
	)
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ResetComparison() {
	_jsii_.InvokeVoid(
		m,
		"resetComparison",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ResetComparisonSensitivity() {
	_jsii_.InvokeVoid(
		m,
		"resetComparisonSensitivity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ResetComparisonThresholdType() {
	_jsii_.InvokeVoid(
		m,
		"resetComparisonThresholdType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricMonitorConditionGroupConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

