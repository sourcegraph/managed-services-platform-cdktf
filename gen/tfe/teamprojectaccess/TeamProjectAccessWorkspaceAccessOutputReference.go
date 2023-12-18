package teamprojectaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/tfe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/tfe/teamprojectaccess/internal"
)

type TeamProjectAccessWorkspaceAccessOutputReference interface {
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
	Create() interface{}
	SetCreate(val interface{})
	CreateInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() interface{}
	SetDelete(val interface{})
	DeleteInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Locking() interface{}
	SetLocking(val interface{})
	LockingInput() interface{}
	Move() interface{}
	SetMove(val interface{})
	MoveInput() interface{}
	Runs() *string
	SetRuns(val *string)
	RunsInput() *string
	RunTasks() interface{}
	SetRunTasks(val interface{})
	RunTasksInput() interface{}
	SentinelMocks() *string
	SetSentinelMocks(val *string)
	SentinelMocksInput() *string
	StateVersions() *string
	SetStateVersions(val *string)
	StateVersionsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Variables() *string
	SetVariables(val *string)
	VariablesInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetLocking()
	ResetMove()
	ResetRuns()
	ResetRunTasks()
	ResetSentinelMocks()
	ResetStateVersions()
	ResetVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TeamProjectAccessWorkspaceAccessOutputReference
type jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) Create() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) CreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) Delete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) DeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) Locking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) LockingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lockingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) Move() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"move",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) MoveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"moveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) Runs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) RunsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) RunTasks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) RunTasksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) SentinelMocks() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sentinelMocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) SentinelMocksInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sentinelMocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) StateVersions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) StateVersionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) Variables() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) VariablesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}


func NewTeamProjectAccessWorkspaceAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TeamProjectAccessWorkspaceAccessOutputReference {
	_init_.Initialize()

	if err := validateNewTeamProjectAccessWorkspaceAccessOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-tfe.teamProjectAccess.TeamProjectAccessWorkspaceAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewTeamProjectAccessWorkspaceAccessOutputReference_Override(t TeamProjectAccessWorkspaceAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tfe.teamProjectAccess.TeamProjectAccessWorkspaceAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetCreate(val interface{}) {
	if err := j.validateSetCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetDelete(val interface{}) {
	if err := j.validateSetDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetLocking(val interface{}) {
	if err := j.validateSetLockingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locking",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetMove(val interface{}) {
	if err := j.validateSetMoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"move",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetRuns(val *string) {
	if err := j.validateSetRunsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runs",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetRunTasks(val interface{}) {
	if err := j.validateSetRunTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runTasks",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetSentinelMocks(val *string) {
	if err := j.validateSetSentinelMocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sentinelMocks",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetStateVersions(val *string) {
	if err := j.validateSetStateVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stateVersions",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference)SetVariables(val *string) {
	if err := j.validateSetVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		t,
		"resetCreate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		t,
		"resetDelete",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ResetLocking() {
	_jsii_.InvokeVoid(
		t,
		"resetLocking",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ResetMove() {
	_jsii_.InvokeVoid(
		t,
		"resetMove",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ResetRuns() {
	_jsii_.InvokeVoid(
		t,
		"resetRuns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ResetRunTasks() {
	_jsii_.InvokeVoid(
		t,
		"resetRunTasks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ResetSentinelMocks() {
	_jsii_.InvokeVoid(
		t,
		"resetSentinelMocks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ResetStateVersions() {
	_jsii_.InvokeVoid(
		t,
		"resetStateVersions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ResetVariables() {
	_jsii_.InvokeVoid(
		t,
		"resetVariables",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamProjectAccessWorkspaceAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

