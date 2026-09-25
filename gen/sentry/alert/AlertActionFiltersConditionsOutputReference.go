package alert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/alert/internal"
)

type AlertActionFiltersConditionsOutputReference interface {
	cdktf.ComplexObject
	AgeComparison() AlertActionFiltersConditionsAgeComparisonOutputReference
	AgeComparisonInput() interface{}
	AssignedTo() AlertActionFiltersConditionsAssignedToOutputReference
	AssignedToInput() interface{}
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
	EventAttribute() AlertActionFiltersConditionsEventAttributeOutputReference
	EventAttributeInput() interface{}
	EventFrequencyCount() AlertActionFiltersConditionsEventFrequencyCountOutputReference
	EventFrequencyCountInput() interface{}
	EventFrequencyPercent() AlertActionFiltersConditionsEventFrequencyPercentOutputReference
	EventFrequencyPercentInput() interface{}
	EventUniqueUserFrequencyCount() AlertActionFiltersConditionsEventUniqueUserFrequencyCountOutputReference
	EventUniqueUserFrequencyCountInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IssueCategory() AlertActionFiltersConditionsIssueCategoryOutputReference
	IssueCategoryInput() interface{}
	IssueOccurrences() AlertActionFiltersConditionsIssueOccurrencesOutputReference
	IssueOccurrencesInput() interface{}
	IssuePriorityDeescalating() AlertActionFiltersConditionsIssuePriorityDeescalatingOutputReference
	IssuePriorityDeescalatingInput() interface{}
	IssuePriorityGreaterOrEqual() AlertActionFiltersConditionsIssuePriorityGreaterOrEqualOutputReference
	IssuePriorityGreaterOrEqualInput() interface{}
	IssueType() AlertActionFiltersConditionsIssueTypeOutputReference
	IssueTypeInput() interface{}
	LatestAdoptedRelease() AlertActionFiltersConditionsLatestAdoptedReleaseOutputReference
	LatestAdoptedReleaseInput() interface{}
	LatestRelease() AlertActionFiltersConditionsLatestReleaseOutputReference
	LatestReleaseInput() interface{}
	Level() AlertActionFiltersConditionsLevelOutputReference
	LevelInput() interface{}
	PercentSessionsCount() AlertActionFiltersConditionsPercentSessionsCountOutputReference
	PercentSessionsCountInput() interface{}
	PercentSessionsPercent() AlertActionFiltersConditionsPercentSessionsPercentOutputReference
	PercentSessionsPercentInput() interface{}
	TaggedEvent() AlertActionFiltersConditionsTaggedEventOutputReference
	TaggedEventInput() interface{}
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
	PutAgeComparison(value *AlertActionFiltersConditionsAgeComparison)
	PutAssignedTo(value *AlertActionFiltersConditionsAssignedTo)
	PutEventAttribute(value *AlertActionFiltersConditionsEventAttribute)
	PutEventFrequencyCount(value *AlertActionFiltersConditionsEventFrequencyCount)
	PutEventFrequencyPercent(value *AlertActionFiltersConditionsEventFrequencyPercent)
	PutEventUniqueUserFrequencyCount(value *AlertActionFiltersConditionsEventUniqueUserFrequencyCount)
	PutIssueCategory(value *AlertActionFiltersConditionsIssueCategory)
	PutIssueOccurrences(value *AlertActionFiltersConditionsIssueOccurrences)
	PutIssuePriorityDeescalating(value *AlertActionFiltersConditionsIssuePriorityDeescalating)
	PutIssuePriorityGreaterOrEqual(value *AlertActionFiltersConditionsIssuePriorityGreaterOrEqual)
	PutIssueType(value *AlertActionFiltersConditionsIssueType)
	PutLatestAdoptedRelease(value *AlertActionFiltersConditionsLatestAdoptedRelease)
	PutLatestRelease(value *AlertActionFiltersConditionsLatestRelease)
	PutLevel(value *AlertActionFiltersConditionsLevel)
	PutPercentSessionsCount(value *AlertActionFiltersConditionsPercentSessionsCount)
	PutPercentSessionsPercent(value *AlertActionFiltersConditionsPercentSessionsPercent)
	PutTaggedEvent(value *AlertActionFiltersConditionsTaggedEvent)
	ResetAgeComparison()
	ResetAssignedTo()
	ResetEventAttribute()
	ResetEventFrequencyCount()
	ResetEventFrequencyPercent()
	ResetEventUniqueUserFrequencyCount()
	ResetIssueCategory()
	ResetIssueOccurrences()
	ResetIssuePriorityDeescalating()
	ResetIssuePriorityGreaterOrEqual()
	ResetIssueType()
	ResetLatestAdoptedRelease()
	ResetLatestRelease()
	ResetLevel()
	ResetPercentSessionsCount()
	ResetPercentSessionsPercent()
	ResetTaggedEvent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertActionFiltersConditionsOutputReference
type jsiiProxy_AlertActionFiltersConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) AgeComparison() AlertActionFiltersConditionsAgeComparisonOutputReference {
	var returns AlertActionFiltersConditionsAgeComparisonOutputReference
	_jsii_.Get(
		j,
		"ageComparison",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) AgeComparisonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ageComparisonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) AssignedTo() AlertActionFiltersConditionsAssignedToOutputReference {
	var returns AlertActionFiltersConditionsAssignedToOutputReference
	_jsii_.Get(
		j,
		"assignedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) AssignedToInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignedToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) EventAttribute() AlertActionFiltersConditionsEventAttributeOutputReference {
	var returns AlertActionFiltersConditionsEventAttributeOutputReference
	_jsii_.Get(
		j,
		"eventAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) EventAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) EventFrequencyCount() AlertActionFiltersConditionsEventFrequencyCountOutputReference {
	var returns AlertActionFiltersConditionsEventFrequencyCountOutputReference
	_jsii_.Get(
		j,
		"eventFrequencyCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) EventFrequencyCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventFrequencyCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) EventFrequencyPercent() AlertActionFiltersConditionsEventFrequencyPercentOutputReference {
	var returns AlertActionFiltersConditionsEventFrequencyPercentOutputReference
	_jsii_.Get(
		j,
		"eventFrequencyPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) EventFrequencyPercentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventFrequencyPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) EventUniqueUserFrequencyCount() AlertActionFiltersConditionsEventUniqueUserFrequencyCountOutputReference {
	var returns AlertActionFiltersConditionsEventUniqueUserFrequencyCountOutputReference
	_jsii_.Get(
		j,
		"eventUniqueUserFrequencyCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) EventUniqueUserFrequencyCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventUniqueUserFrequencyCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) IssueCategory() AlertActionFiltersConditionsIssueCategoryOutputReference {
	var returns AlertActionFiltersConditionsIssueCategoryOutputReference
	_jsii_.Get(
		j,
		"issueCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) IssueCategoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) IssueOccurrences() AlertActionFiltersConditionsIssueOccurrencesOutputReference {
	var returns AlertActionFiltersConditionsIssueOccurrencesOutputReference
	_jsii_.Get(
		j,
		"issueOccurrences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) IssueOccurrencesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueOccurrencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) IssuePriorityDeescalating() AlertActionFiltersConditionsIssuePriorityDeescalatingOutputReference {
	var returns AlertActionFiltersConditionsIssuePriorityDeescalatingOutputReference
	_jsii_.Get(
		j,
		"issuePriorityDeescalating",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) IssuePriorityDeescalatingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuePriorityDeescalatingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) IssuePriorityGreaterOrEqual() AlertActionFiltersConditionsIssuePriorityGreaterOrEqualOutputReference {
	var returns AlertActionFiltersConditionsIssuePriorityGreaterOrEqualOutputReference
	_jsii_.Get(
		j,
		"issuePriorityGreaterOrEqual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) IssuePriorityGreaterOrEqualInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuePriorityGreaterOrEqualInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) IssueType() AlertActionFiltersConditionsIssueTypeOutputReference {
	var returns AlertActionFiltersConditionsIssueTypeOutputReference
	_jsii_.Get(
		j,
		"issueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) IssueTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) LatestAdoptedRelease() AlertActionFiltersConditionsLatestAdoptedReleaseOutputReference {
	var returns AlertActionFiltersConditionsLatestAdoptedReleaseOutputReference
	_jsii_.Get(
		j,
		"latestAdoptedRelease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) LatestAdoptedReleaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latestAdoptedReleaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) LatestRelease() AlertActionFiltersConditionsLatestReleaseOutputReference {
	var returns AlertActionFiltersConditionsLatestReleaseOutputReference
	_jsii_.Get(
		j,
		"latestRelease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) LatestReleaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latestReleaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) Level() AlertActionFiltersConditionsLevelOutputReference {
	var returns AlertActionFiltersConditionsLevelOutputReference
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) LevelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) PercentSessionsCount() AlertActionFiltersConditionsPercentSessionsCountOutputReference {
	var returns AlertActionFiltersConditionsPercentSessionsCountOutputReference
	_jsii_.Get(
		j,
		"percentSessionsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) PercentSessionsCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"percentSessionsCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) PercentSessionsPercent() AlertActionFiltersConditionsPercentSessionsPercentOutputReference {
	var returns AlertActionFiltersConditionsPercentSessionsPercentOutputReference
	_jsii_.Get(
		j,
		"percentSessionsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) PercentSessionsPercentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"percentSessionsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) TaggedEvent() AlertActionFiltersConditionsTaggedEventOutputReference {
	var returns AlertActionFiltersConditionsTaggedEventOutputReference
	_jsii_.Get(
		j,
		"taggedEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) TaggedEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taggedEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlertActionFiltersConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlertActionFiltersConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewAlertActionFiltersConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertActionFiltersConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.alert.AlertActionFiltersConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlertActionFiltersConditionsOutputReference_Override(a AlertActionFiltersConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.alert.AlertActionFiltersConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutAgeComparison(value *AlertActionFiltersConditionsAgeComparison) {
	if err := a.validatePutAgeComparisonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAgeComparison",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutAssignedTo(value *AlertActionFiltersConditionsAssignedTo) {
	if err := a.validatePutAssignedToParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAssignedTo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutEventAttribute(value *AlertActionFiltersConditionsEventAttribute) {
	if err := a.validatePutEventAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventAttribute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutEventFrequencyCount(value *AlertActionFiltersConditionsEventFrequencyCount) {
	if err := a.validatePutEventFrequencyCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventFrequencyCount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutEventFrequencyPercent(value *AlertActionFiltersConditionsEventFrequencyPercent) {
	if err := a.validatePutEventFrequencyPercentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventFrequencyPercent",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutEventUniqueUserFrequencyCount(value *AlertActionFiltersConditionsEventUniqueUserFrequencyCount) {
	if err := a.validatePutEventUniqueUserFrequencyCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventUniqueUserFrequencyCount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutIssueCategory(value *AlertActionFiltersConditionsIssueCategory) {
	if err := a.validatePutIssueCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIssueCategory",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutIssueOccurrences(value *AlertActionFiltersConditionsIssueOccurrences) {
	if err := a.validatePutIssueOccurrencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIssueOccurrences",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutIssuePriorityDeescalating(value *AlertActionFiltersConditionsIssuePriorityDeescalating) {
	if err := a.validatePutIssuePriorityDeescalatingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIssuePriorityDeescalating",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutIssuePriorityGreaterOrEqual(value *AlertActionFiltersConditionsIssuePriorityGreaterOrEqual) {
	if err := a.validatePutIssuePriorityGreaterOrEqualParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIssuePriorityGreaterOrEqual",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutIssueType(value *AlertActionFiltersConditionsIssueType) {
	if err := a.validatePutIssueTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIssueType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutLatestAdoptedRelease(value *AlertActionFiltersConditionsLatestAdoptedRelease) {
	if err := a.validatePutLatestAdoptedReleaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLatestAdoptedRelease",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutLatestRelease(value *AlertActionFiltersConditionsLatestRelease) {
	if err := a.validatePutLatestReleaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLatestRelease",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutLevel(value *AlertActionFiltersConditionsLevel) {
	if err := a.validatePutLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLevel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutPercentSessionsCount(value *AlertActionFiltersConditionsPercentSessionsCount) {
	if err := a.validatePutPercentSessionsCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPercentSessionsCount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutPercentSessionsPercent(value *AlertActionFiltersConditionsPercentSessionsPercent) {
	if err := a.validatePutPercentSessionsPercentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPercentSessionsPercent",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) PutTaggedEvent(value *AlertActionFiltersConditionsTaggedEvent) {
	if err := a.validatePutTaggedEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTaggedEvent",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetAgeComparison() {
	_jsii_.InvokeVoid(
		a,
		"resetAgeComparison",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetAssignedTo() {
	_jsii_.InvokeVoid(
		a,
		"resetAssignedTo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetEventAttribute() {
	_jsii_.InvokeVoid(
		a,
		"resetEventAttribute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetEventFrequencyCount() {
	_jsii_.InvokeVoid(
		a,
		"resetEventFrequencyCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetEventFrequencyPercent() {
	_jsii_.InvokeVoid(
		a,
		"resetEventFrequencyPercent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetEventUniqueUserFrequencyCount() {
	_jsii_.InvokeVoid(
		a,
		"resetEventUniqueUserFrequencyCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetIssueCategory() {
	_jsii_.InvokeVoid(
		a,
		"resetIssueCategory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetIssueOccurrences() {
	_jsii_.InvokeVoid(
		a,
		"resetIssueOccurrences",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetIssuePriorityDeescalating() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuePriorityDeescalating",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetIssuePriorityGreaterOrEqual() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuePriorityGreaterOrEqual",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetIssueType() {
	_jsii_.InvokeVoid(
		a,
		"resetIssueType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetLatestAdoptedRelease() {
	_jsii_.InvokeVoid(
		a,
		"resetLatestAdoptedRelease",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetLatestRelease() {
	_jsii_.InvokeVoid(
		a,
		"resetLatestRelease",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetPercentSessionsCount() {
	_jsii_.InvokeVoid(
		a,
		"resetPercentSessionsCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetPercentSessionsPercent() {
	_jsii_.InvokeVoid(
		a,
		"resetPercentSessionsPercent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ResetTaggedEvent() {
	_jsii_.InvokeVoid(
		a,
		"resetTaggedEvent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlertActionFiltersConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

