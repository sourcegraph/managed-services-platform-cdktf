package alert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/alert/internal"
)

type AlertActionFiltersActionsJiraOutputReference interface {
	cdktf.ComplexObject
	AdditionalFields() *map[string]*string
	SetAdditionalFields(val *map[string]*string)
	AdditionalFieldsInput() *map[string]*string
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
	Components() *[]*string
	SetComponents(val *[]*string)
	ComponentsInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IntegrationId() *string
	SetIntegrationId(val *string)
	IntegrationIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IssueType() *string
	SetIssueType(val *string)
	IssueTypeInput() *string
	Labels() *string
	SetLabels(val *string)
	LabelsInput() *string
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Reporter() *string
	SetReporter(val *string)
	ReporterInput() *string
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
	ResetAdditionalFields()
	ResetComponents()
	ResetLabels()
	ResetPriority()
	ResetReporter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertActionFiltersActionsJiraOutputReference
type jsiiProxy_AlertActionFiltersActionsJiraOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) AdditionalFields() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) AdditionalFieldsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) Components() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ComponentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"componentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) IntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) IntegrationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) IssueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) IssueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) Labels() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) LabelsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) Reporter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reporter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ReporterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reporterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlertActionFiltersActionsJiraOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlertActionFiltersActionsJiraOutputReference {
	_init_.Initialize()

	if err := validateNewAlertActionFiltersActionsJiraOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertActionFiltersActionsJiraOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.alert.AlertActionFiltersActionsJiraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlertActionFiltersActionsJiraOutputReference_Override(a AlertActionFiltersActionsJiraOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.alert.AlertActionFiltersActionsJiraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetAdditionalFields(val *map[string]*string) {
	if err := j.validateSetAdditionalFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalFields",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetComponents(val *[]*string) {
	if err := j.validateSetComponentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"components",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetIntegrationId(val *string) {
	if err := j.validateSetIntegrationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationId",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetIssueType(val *string) {
	if err := j.validateSetIssueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueType",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetLabels(val *string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetReporter(val *string) {
	if err := j.validateSetReporterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reporter",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsJiraOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ResetAdditionalFields() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalFields",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ResetComponents() {
	_jsii_.InvokeVoid(
		a,
		"resetComponents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		a,
		"resetLabels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ResetReporter() {
	_jsii_.InvokeVoid(
		a,
		"resetReporter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlertActionFiltersActionsJiraOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

