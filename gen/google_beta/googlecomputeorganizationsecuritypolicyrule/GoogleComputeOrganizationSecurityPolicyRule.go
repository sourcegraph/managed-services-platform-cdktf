package googlecomputeorganizationsecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputeorganizationsecuritypolicyrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_organization_security_policy_rule google_compute_organization_security_policy_rule}.
type GoogleComputeOrganizationSecurityPolicyRule interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	EnableLogging() any
	SetEnableLogging(val any)
	EnableLoggingInput() any
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Match() GoogleComputeOrganizationSecurityPolicyRuleMatchOutputReference
	MatchInput() *GoogleComputeOrganizationSecurityPolicyRuleMatch
	// The tree node.
	Node() constructs.Node
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	Preview() any
	SetPreview(val any)
	PreviewInput() any
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	TargetResources() *[]*string
	SetTargetResources(val *[]*string)
	TargetResourcesInput() *[]*string
	TargetServiceAccounts() *[]*string
	SetTargetServiceAccounts(val *[]*string)
	TargetServiceAccountsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeOrganizationSecurityPolicyRuleTimeoutsOutputReference
	TimeoutsInput() any
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
	PutMatch(value *GoogleComputeOrganizationSecurityPolicyRuleMatch)
	PutTimeouts(value *GoogleComputeOrganizationSecurityPolicyRuleTimeouts)
	ResetDescription()
	ResetDirection()
	ResetEnableLogging()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreview()
	ResetTargetResources()
	ResetTargetServiceAccounts()
	ResetTimeouts()
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

// The jsii proxy struct for GoogleComputeOrganizationSecurityPolicyRule
type jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) EnableLogging() any {
	var returns any
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) EnableLoggingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Match() GoogleComputeOrganizationSecurityPolicyRuleMatchOutputReference {
	var returns GoogleComputeOrganizationSecurityPolicyRuleMatchOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) MatchInput() *GoogleComputeOrganizationSecurityPolicyRuleMatch {
	var returns *GoogleComputeOrganizationSecurityPolicyRuleMatch
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Preview() any {
	var returns any
	_jsii_.Get(
		j,
		"preview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) PreviewInput() any {
	var returns any
	_jsii_.Get(
		j,
		"previewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) TargetResources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) TargetResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) TargetServiceAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) TargetServiceAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) Timeouts() GoogleComputeOrganizationSecurityPolicyRuleTimeoutsOutputReference {
	var returns GoogleComputeOrganizationSecurityPolicyRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_organization_security_policy_rule google_compute_organization_security_policy_rule} Resource.
func NewGoogleComputeOrganizationSecurityPolicyRule(scope constructs.Construct, id *string, config *GoogleComputeOrganizationSecurityPolicyRuleConfig) GoogleComputeOrganizationSecurityPolicyRule {
	_init_.Initialize()

	if err := validateNewGoogleComputeOrganizationSecurityPolicyRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeOrganizationSecurityPolicyRule.GoogleComputeOrganizationSecurityPolicyRule",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_organization_security_policy_rule google_compute_organization_security_policy_rule} Resource.
func NewGoogleComputeOrganizationSecurityPolicyRule_Override(g GoogleComputeOrganizationSecurityPolicyRule, scope constructs.Construct, id *string, config *GoogleComputeOrganizationSecurityPolicyRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeOrganizationSecurityPolicyRule.GoogleComputeOrganizationSecurityPolicyRule",
		[]any{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetEnableLogging(val any) {
	if err := j.validateSetEnableLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetPreview(val any) {
	if err := j.validateSetPreviewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preview",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetTargetResources(val *[]*string) {
	if err := j.validateSetTargetResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResources",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SetTargetServiceAccounts(val *[]*string) {
	if err := j.validateSetTargetServiceAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetServiceAccounts",
		val,
	)
}

// Generates CDKTF code for importing a GoogleComputeOrganizationSecurityPolicyRule resource upon running "cdktf plan <stack-name>".
func GoogleComputeOrganizationSecurityPolicyRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeOrganizationSecurityPolicyRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleComputeOrganizationSecurityPolicyRule.GoogleComputeOrganizationSecurityPolicyRule",
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
func GoogleComputeOrganizationSecurityPolicyRule_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeOrganizationSecurityPolicyRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleComputeOrganizationSecurityPolicyRule.GoogleComputeOrganizationSecurityPolicyRule",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeOrganizationSecurityPolicyRule_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeOrganizationSecurityPolicyRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleComputeOrganizationSecurityPolicyRule.GoogleComputeOrganizationSecurityPolicyRule",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeOrganizationSecurityPolicyRule_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeOrganizationSecurityPolicyRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleComputeOrganizationSecurityPolicyRule.GoogleComputeOrganizationSecurityPolicyRule",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func GoogleComputeOrganizationSecurityPolicyRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleComputeOrganizationSecurityPolicyRule.GoogleComputeOrganizationSecurityPolicyRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) AddOverride(path *string, value any) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]any{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]any{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]any{id},
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) MoveTo(moveTarget *string, index any) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]any{id},
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) PutMatch(value *GoogleComputeOrganizationSecurityPolicyRuleMatch) {
	if err := g.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMatch",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) PutTimeouts(value *GoogleComputeOrganizationSecurityPolicyRuleTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ResetDirection() {
	_jsii_.InvokeVoid(
		g,
		"resetDirection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ResetEnableLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ResetPreview() {
	_jsii_.InvokeVoid(
		g,
		"resetPreview",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ResetTargetResources() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ResetTargetServiceAccounts() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetServiceAccounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeOrganizationSecurityPolicyRule) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
