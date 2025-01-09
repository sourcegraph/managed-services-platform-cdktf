package datasentryissuealert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/datasentryissuealert/internal"
)

type DataSentryIssueAlertFiltersV2OutputReference interface {
	cdktf.ComplexObject
	AgeComparison() DataSentryIssueAlertFiltersV2AgeComparisonOutputReference
	AssignedTo() DataSentryIssueAlertFiltersV2AssignedToOutputReference
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
	EventAttribute() DataSentryIssueAlertFiltersV2EventAttributeOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataSentryIssueAlertFiltersV2
	SetInternalValue(val *DataSentryIssueAlertFiltersV2)
	IssueCategory() DataSentryIssueAlertFiltersV2IssueCategoryOutputReference
	IssueOccurrences() DataSentryIssueAlertFiltersV2IssueOccurrencesOutputReference
	LatestAdoptedRelease() DataSentryIssueAlertFiltersV2LatestAdoptedReleaseOutputReference
	LatestRelease() DataSentryIssueAlertFiltersV2LatestReleaseOutputReference
	Level() DataSentryIssueAlertFiltersV2LevelOutputReference
	TaggedEvent() DataSentryIssueAlertFiltersV2TaggedEventOutputReference
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

// The jsii proxy struct for DataSentryIssueAlertFiltersV2OutputReference
type jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) AgeComparison() DataSentryIssueAlertFiltersV2AgeComparisonOutputReference {
	var returns DataSentryIssueAlertFiltersV2AgeComparisonOutputReference
	_jsii_.Get(
		j,
		"ageComparison",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) AssignedTo() DataSentryIssueAlertFiltersV2AssignedToOutputReference {
	var returns DataSentryIssueAlertFiltersV2AssignedToOutputReference
	_jsii_.Get(
		j,
		"assignedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) EventAttribute() DataSentryIssueAlertFiltersV2EventAttributeOutputReference {
	var returns DataSentryIssueAlertFiltersV2EventAttributeOutputReference
	_jsii_.Get(
		j,
		"eventAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) InternalValue() *DataSentryIssueAlertFiltersV2 {
	var returns *DataSentryIssueAlertFiltersV2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) IssueCategory() DataSentryIssueAlertFiltersV2IssueCategoryOutputReference {
	var returns DataSentryIssueAlertFiltersV2IssueCategoryOutputReference
	_jsii_.Get(
		j,
		"issueCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) IssueOccurrences() DataSentryIssueAlertFiltersV2IssueOccurrencesOutputReference {
	var returns DataSentryIssueAlertFiltersV2IssueOccurrencesOutputReference
	_jsii_.Get(
		j,
		"issueOccurrences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) LatestAdoptedRelease() DataSentryIssueAlertFiltersV2LatestAdoptedReleaseOutputReference {
	var returns DataSentryIssueAlertFiltersV2LatestAdoptedReleaseOutputReference
	_jsii_.Get(
		j,
		"latestAdoptedRelease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) LatestRelease() DataSentryIssueAlertFiltersV2LatestReleaseOutputReference {
	var returns DataSentryIssueAlertFiltersV2LatestReleaseOutputReference
	_jsii_.Get(
		j,
		"latestRelease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) Level() DataSentryIssueAlertFiltersV2LevelOutputReference {
	var returns DataSentryIssueAlertFiltersV2LevelOutputReference
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) TaggedEvent() DataSentryIssueAlertFiltersV2TaggedEventOutputReference {
	var returns DataSentryIssueAlertFiltersV2TaggedEventOutputReference
	_jsii_.Get(
		j,
		"taggedEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataSentryIssueAlertFiltersV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSentryIssueAlertFiltersV2OutputReference {
	_init_.Initialize()

	if err := validateNewDataSentryIssueAlertFiltersV2OutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.dataSentryIssueAlert.DataSentryIssueAlertFiltersV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSentryIssueAlertFiltersV2OutputReference_Override(d DataSentryIssueAlertFiltersV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.dataSentryIssueAlert.DataSentryIssueAlertFiltersV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference)SetInternalValue(val *DataSentryIssueAlertFiltersV2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertFiltersV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

