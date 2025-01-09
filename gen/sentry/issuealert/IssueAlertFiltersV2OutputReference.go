package issuealert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/issuealert/internal"
)

type IssueAlertFiltersV2OutputReference interface {
	cdktf.ComplexObject
	AgeComparison() IssueAlertFiltersV2AgeComparisonOutputReference
	AgeComparisonInput() interface{}
	AssignedTo() IssueAlertFiltersV2AssignedToOutputReference
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
	EventAttribute() IssueAlertFiltersV2EventAttributeOutputReference
	EventAttributeInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IssueCategory() IssueAlertFiltersV2IssueCategoryOutputReference
	IssueCategoryInput() interface{}
	IssueOccurrences() IssueAlertFiltersV2IssueOccurrencesOutputReference
	IssueOccurrencesInput() interface{}
	LatestAdoptedRelease() IssueAlertFiltersV2LatestAdoptedReleaseOutputReference
	LatestAdoptedReleaseInput() interface{}
	LatestRelease() IssueAlertFiltersV2LatestReleaseOutputReference
	LatestReleaseInput() interface{}
	Level() IssueAlertFiltersV2LevelOutputReference
	LevelInput() interface{}
	TaggedEvent() IssueAlertFiltersV2TaggedEventOutputReference
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
	PutAgeComparison(value *IssueAlertFiltersV2AgeComparison)
	PutAssignedTo(value *IssueAlertFiltersV2AssignedTo)
	PutEventAttribute(value *IssueAlertFiltersV2EventAttribute)
	PutIssueCategory(value *IssueAlertFiltersV2IssueCategory)
	PutIssueOccurrences(value *IssueAlertFiltersV2IssueOccurrences)
	PutLatestAdoptedRelease(value *IssueAlertFiltersV2LatestAdoptedRelease)
	PutLatestRelease(value *IssueAlertFiltersV2LatestRelease)
	PutLevel(value *IssueAlertFiltersV2Level)
	PutTaggedEvent(value *IssueAlertFiltersV2TaggedEvent)
	ResetAgeComparison()
	ResetAssignedTo()
	ResetEventAttribute()
	ResetIssueCategory()
	ResetIssueOccurrences()
	ResetLatestAdoptedRelease()
	ResetLatestRelease()
	ResetLevel()
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

// The jsii proxy struct for IssueAlertFiltersV2OutputReference
type jsiiProxy_IssueAlertFiltersV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) AgeComparison() IssueAlertFiltersV2AgeComparisonOutputReference {
	var returns IssueAlertFiltersV2AgeComparisonOutputReference
	_jsii_.Get(
		j,
		"ageComparison",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) AgeComparisonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ageComparisonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) AssignedTo() IssueAlertFiltersV2AssignedToOutputReference {
	var returns IssueAlertFiltersV2AssignedToOutputReference
	_jsii_.Get(
		j,
		"assignedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) AssignedToInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignedToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) EventAttribute() IssueAlertFiltersV2EventAttributeOutputReference {
	var returns IssueAlertFiltersV2EventAttributeOutputReference
	_jsii_.Get(
		j,
		"eventAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) EventAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) IssueCategory() IssueAlertFiltersV2IssueCategoryOutputReference {
	var returns IssueAlertFiltersV2IssueCategoryOutputReference
	_jsii_.Get(
		j,
		"issueCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) IssueCategoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) IssueOccurrences() IssueAlertFiltersV2IssueOccurrencesOutputReference {
	var returns IssueAlertFiltersV2IssueOccurrencesOutputReference
	_jsii_.Get(
		j,
		"issueOccurrences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) IssueOccurrencesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issueOccurrencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) LatestAdoptedRelease() IssueAlertFiltersV2LatestAdoptedReleaseOutputReference {
	var returns IssueAlertFiltersV2LatestAdoptedReleaseOutputReference
	_jsii_.Get(
		j,
		"latestAdoptedRelease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) LatestAdoptedReleaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latestAdoptedReleaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) LatestRelease() IssueAlertFiltersV2LatestReleaseOutputReference {
	var returns IssueAlertFiltersV2LatestReleaseOutputReference
	_jsii_.Get(
		j,
		"latestRelease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) LatestReleaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latestReleaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) Level() IssueAlertFiltersV2LevelOutputReference {
	var returns IssueAlertFiltersV2LevelOutputReference
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) LevelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) TaggedEvent() IssueAlertFiltersV2TaggedEventOutputReference {
	var returns IssueAlertFiltersV2TaggedEventOutputReference
	_jsii_.Get(
		j,
		"taggedEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) TaggedEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taggedEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIssueAlertFiltersV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IssueAlertFiltersV2OutputReference {
	_init_.Initialize()

	if err := validateNewIssueAlertFiltersV2OutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IssueAlertFiltersV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.issueAlert.IssueAlertFiltersV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIssueAlertFiltersV2OutputReference_Override(i IssueAlertFiltersV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.issueAlert.IssueAlertFiltersV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IssueAlertFiltersV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) PutAgeComparison(value *IssueAlertFiltersV2AgeComparison) {
	if err := i.validatePutAgeComparisonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAgeComparison",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) PutAssignedTo(value *IssueAlertFiltersV2AssignedTo) {
	if err := i.validatePutAssignedToParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAssignedTo",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) PutEventAttribute(value *IssueAlertFiltersV2EventAttribute) {
	if err := i.validatePutEventAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEventAttribute",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) PutIssueCategory(value *IssueAlertFiltersV2IssueCategory) {
	if err := i.validatePutIssueCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIssueCategory",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) PutIssueOccurrences(value *IssueAlertFiltersV2IssueOccurrences) {
	if err := i.validatePutIssueOccurrencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIssueOccurrences",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) PutLatestAdoptedRelease(value *IssueAlertFiltersV2LatestAdoptedRelease) {
	if err := i.validatePutLatestAdoptedReleaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLatestAdoptedRelease",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) PutLatestRelease(value *IssueAlertFiltersV2LatestRelease) {
	if err := i.validatePutLatestReleaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLatestRelease",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) PutLevel(value *IssueAlertFiltersV2Level) {
	if err := i.validatePutLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLevel",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) PutTaggedEvent(value *IssueAlertFiltersV2TaggedEvent) {
	if err := i.validatePutTaggedEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTaggedEvent",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) ResetAgeComparison() {
	_jsii_.InvokeVoid(
		i,
		"resetAgeComparison",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) ResetAssignedTo() {
	_jsii_.InvokeVoid(
		i,
		"resetAssignedTo",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) ResetEventAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetEventAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) ResetIssueCategory() {
	_jsii_.InvokeVoid(
		i,
		"resetIssueCategory",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) ResetIssueOccurrences() {
	_jsii_.InvokeVoid(
		i,
		"resetIssueOccurrences",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) ResetLatestAdoptedRelease() {
	_jsii_.InvokeVoid(
		i,
		"resetLatestAdoptedRelease",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) ResetLatestRelease() {
	_jsii_.InvokeVoid(
		i,
		"resetLatestRelease",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		i,
		"resetLevel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) ResetTaggedEvent() {
	_jsii_.InvokeVoid(
		i,
		"resetTaggedEvent",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IssueAlertFiltersV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

