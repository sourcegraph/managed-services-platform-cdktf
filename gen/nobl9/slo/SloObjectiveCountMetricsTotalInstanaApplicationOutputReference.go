package slo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/nobl9/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/nobl9/slo/internal"
)

type SloObjectiveCountMetricsTotalInstanaApplicationOutputReference interface {
	cdktf.ComplexObject
	Aggregation() *string
	SetAggregation(val *string)
	AggregationInput() *string
	ApiQuery() *string
	SetApiQuery(val *string)
	ApiQueryInput() *string
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
	// Experimental.
	Fqn() *string
	GroupBy() SloObjectiveCountMetricsTotalInstanaApplicationGroupByList
	GroupByInput() interface{}
	IncludeInternal() interface{}
	SetIncludeInternal(val interface{})
	IncludeInternalInput() interface{}
	IncludeSynthetic() interface{}
	SetIncludeSynthetic(val interface{})
	IncludeSyntheticInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricId() *string
	SetMetricId(val *string)
	MetricIdInput() *string
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
	PutGroupBy(value interface{})
	ResetIncludeInternal()
	ResetIncludeSynthetic()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SloObjectiveCountMetricsTotalInstanaApplicationOutputReference
type jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) Aggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) AggregationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) ApiQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) ApiQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) GroupBy() SloObjectiveCountMetricsTotalInstanaApplicationGroupByList {
	var returns SloObjectiveCountMetricsTotalInstanaApplicationGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) IncludeInternal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInternal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) IncludeInternalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInternalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) IncludeSynthetic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSynthetic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) IncludeSyntheticInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSyntheticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) MetricId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) MetricIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSloObjectiveCountMetricsTotalInstanaApplicationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SloObjectiveCountMetricsTotalInstanaApplicationOutputReference {
	_init_.Initialize()

	if err := validateNewSloObjectiveCountMetricsTotalInstanaApplicationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsTotalInstanaApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSloObjectiveCountMetricsTotalInstanaApplicationOutputReference_Override(s SloObjectiveCountMetricsTotalInstanaApplicationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-nobl9.slo.SloObjectiveCountMetricsTotalInstanaApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference)SetAggregation(val *string) {
	if err := j.validateSetAggregationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregation",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference)SetApiQuery(val *string) {
	if err := j.validateSetApiQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiQuery",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference)SetIncludeInternal(val interface{}) {
	if err := j.validateSetIncludeInternalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeInternal",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference)SetIncludeSynthetic(val interface{}) {
	if err := j.validateSetIncludeSyntheticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSynthetic",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference)SetMetricId(val *string) {
	if err := j.validateSetMetricIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricId",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) PutGroupBy(value interface{}) {
	if err := s.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) ResetIncludeInternal() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeInternal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) ResetIncludeSynthetic() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeSynthetic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SloObjectiveCountMetricsTotalInstanaApplicationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

