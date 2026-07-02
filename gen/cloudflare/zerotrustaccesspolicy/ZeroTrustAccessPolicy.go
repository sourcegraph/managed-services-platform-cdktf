package zerotrustaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/zerotrustaccesspolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy cloudflare_zero_trust_access_policy}.
type ZeroTrustAccessPolicy interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	ApprovalGroups() ZeroTrustAccessPolicyApprovalGroupsList
	ApprovalGroupsInput() any
	ApprovalRequired() any
	SetApprovalRequired(val any)
	ApprovalRequiredInput() any
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	Decision() *string
	SetDecision(val *string)
	DecisionInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Exclude() ZeroTrustAccessPolicyExcludeList
	ExcludeInput() any
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	Include() ZeroTrustAccessPolicyIncludeList
	IncludeInput() any
	IsolationRequired() any
	SetIsolationRequired(val any)
	IsolationRequiredInput() any
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	PurposeJustificationPrompt() *string
	SetPurposeJustificationPrompt(val *string)
	PurposeJustificationPromptInput() *string
	PurposeJustificationRequired() any
	SetPurposeJustificationRequired(val any)
	PurposeJustificationRequiredInput() any
	// Experimental.
	RawOverrides() any
	Require() ZeroTrustAccessPolicyRequireList
	RequireInput() any
	SessionDuration() *string
	SetSessionDuration(val *string)
	SessionDurationInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
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
	HasResourceMove() any
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutApprovalGroups(value any)
	PutExclude(value any)
	PutInclude(value any)
	PutRequire(value any)
	ResetApprovalGroups()
	ResetApprovalRequired()
	ResetExclude()
	ResetIsolationRequired()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPurposeJustificationPrompt()
	ResetPurposeJustificationRequired()
	ResetRequire()
	ResetSessionDuration()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for ZeroTrustAccessPolicy
type jsiiProxy_ZeroTrustAccessPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ApprovalGroups() ZeroTrustAccessPolicyApprovalGroupsList {
	var returns ZeroTrustAccessPolicyApprovalGroupsList
	_jsii_.Get(
		j,
		"approvalGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ApprovalGroupsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"approvalGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ApprovalRequired() any {
	var returns any
	_jsii_.Get(
		j,
		"approvalRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ApprovalRequiredInput() any {
	var returns any
	_jsii_.Get(
		j,
		"approvalRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Decision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) DecisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Exclude() ZeroTrustAccessPolicyExcludeList {
	var returns ZeroTrustAccessPolicyExcludeList
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ExcludeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Include() ZeroTrustAccessPolicyIncludeList {
	var returns ZeroTrustAccessPolicyIncludeList
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) IncludeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) IsolationRequired() any {
	var returns any
	_jsii_.Get(
		j,
		"isolationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) IsolationRequiredInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isolationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) PurposeJustificationPrompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purposeJustificationPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) PurposeJustificationPromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purposeJustificationPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) PurposeJustificationRequired() any {
	var returns any
	_jsii_.Get(
		j,
		"purposeJustificationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) PurposeJustificationRequiredInput() any {
	var returns any
	_jsii_.Get(
		j,
		"purposeJustificationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) Require() ZeroTrustAccessPolicyRequireList {
	var returns ZeroTrustAccessPolicyRequireList
	_jsii_.Get(
		j,
		"require",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) RequireInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requireInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SessionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SessionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy cloudflare_zero_trust_access_policy} Resource.
func NewZeroTrustAccessPolicy(scope constructs.Construct, id *string, config *ZeroTrustAccessPolicyConfig) ZeroTrustAccessPolicy {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessPolicy{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_policy cloudflare_zero_trust_access_policy} Resource.
func NewZeroTrustAccessPolicy_Override(z ZeroTrustAccessPolicy, scope constructs.Construct, id *string, config *ZeroTrustAccessPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		[]any{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetApprovalRequired(val any) {
	if err := j.validateSetApprovalRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalRequired",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetDecision(val *string) {
	if err := j.validateSetDecisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decision",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetIsolationRequired(val any) {
	if err := j.validateSetIsolationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isolationRequired",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetPurposeJustificationPrompt(val *string) {
	if err := j.validateSetPurposeJustificationPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purposeJustificationPrompt",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetPurposeJustificationRequired(val any) {
	if err := j.validateSetPurposeJustificationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purposeJustificationRequired",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicy) SetSessionDuration(val *string) {
	if err := j.validateSetSessionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionDuration",
		val,
	)
}

// Generates CDKTF code for importing a ZeroTrustAccessPolicy resource upon running "cdktf plan <stack-name>".
func ZeroTrustAccessPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustAccessPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ZeroTrustAccessPolicy_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustAccessPolicy_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustAccessPolicy_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func ZeroTrustAccessPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) AddOverride(path *string, value any) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]any{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]any{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]any{id},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) MoveTo(moveTarget *string, index any) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]any{id},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) PutApprovalGroups(value any) {
	if err := z.validatePutApprovalGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putApprovalGroups",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) PutExclude(value any) {
	if err := z.validatePutExcludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExclude",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) PutInclude(value any) {
	if err := z.validatePutIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putInclude",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) PutRequire(value any) {
	if err := z.validatePutRequireParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putRequire",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetApprovalGroups() {
	_jsii_.InvokeVoid(
		z,
		"resetApprovalGroups",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetApprovalRequired() {
	_jsii_.InvokeVoid(
		z,
		"resetApprovalRequired",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetExclude() {
	_jsii_.InvokeVoid(
		z,
		"resetExclude",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetIsolationRequired() {
	_jsii_.InvokeVoid(
		z,
		"resetIsolationRequired",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetPurposeJustificationPrompt() {
	_jsii_.InvokeVoid(
		z,
		"resetPurposeJustificationPrompt",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetPurposeJustificationRequired() {
	_jsii_.InvokeVoid(
		z,
		"resetPurposeJustificationRequired",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetRequire() {
	_jsii_.InvokeVoid(
		z,
		"resetRequire",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ResetSessionDuration() {
	_jsii_.InvokeVoid(
		z,
		"resetSessionDuration",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicy) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
