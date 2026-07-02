package alertpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/alertpolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/alert_policy opsgenie_alert_policy}.
type AlertPolicy interface {
	cdktf.TerraformResource
	Actions() *[]*string
	SetActions(val *[]*string)
	ActionsInput() *[]*string
	AlertDescription() *string
	SetAlertDescription(val *string)
	AlertDescriptionInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	ContinuePolicy() any
	SetContinuePolicy(val any)
	ContinuePolicyInput() any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
	Entity() *string
	SetEntity(val *string)
	EntityInput() *string
	Filter() AlertPolicyFilterOutputReference
	FilterInput() *AlertPolicyFilter
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
	IgnoreOriginalActions() any
	SetIgnoreOriginalActions(val any)
	IgnoreOriginalActionsInput() any
	IgnoreOriginalDetails() any
	SetIgnoreOriginalDetails(val any)
	IgnoreOriginalDetailsInput() any
	IgnoreOriginalResponders() any
	SetIgnoreOriginalResponders(val any)
	IgnoreOriginalRespondersInput() any
	IgnoreOriginalTags() any
	SetIgnoreOriginalTags(val any)
	IgnoreOriginalTagsInput() any
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Message() *string
	SetMessage(val *string)
	MessageInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PolicyDescription() *string
	SetPolicyDescription(val *string)
	PolicyDescriptionInput() *string
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
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
	Responders() AlertPolicyRespondersList
	RespondersInput() any
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	TeamId() *string
	SetTeamId(val *string)
	TeamIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	TimeRestriction() AlertPolicyTimeRestrictionOutputReference
	TimeRestrictionInput() *AlertPolicyTimeRestriction
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
	PutFilter(value *AlertPolicyFilter)
	PutResponders(value any)
	PutTimeRestriction(value *AlertPolicyTimeRestriction)
	ResetActions()
	ResetAlertDescription()
	ResetAlias()
	ResetContinuePolicy()
	ResetEnabled()
	ResetEntity()
	ResetFilter()
	ResetId()
	ResetIgnoreOriginalActions()
	ResetIgnoreOriginalDetails()
	ResetIgnoreOriginalResponders()
	ResetIgnoreOriginalTags()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyDescription()
	ResetPriority()
	ResetResponders()
	ResetSource()
	ResetTags()
	ResetTeamId()
	ResetTimeRestriction()
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

// The jsii proxy struct for AlertPolicy
type jsiiProxy_AlertPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AlertPolicy) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) ActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) AlertDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) AlertDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) ContinuePolicy() any {
	var returns any
	_jsii_.Get(
		j,
		"continuePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) ContinuePolicyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"continuePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Entity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) EntityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Filter() AlertPolicyFilterOutputReference {
	var returns AlertPolicyFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) FilterInput() *AlertPolicyFilter {
	var returns *AlertPolicyFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalActions() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreOriginalActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalActionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreOriginalActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalDetails() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreOriginalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalDetailsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreOriginalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalResponders() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreOriginalResponders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalRespondersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreOriginalRespondersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalTags() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreOriginalTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalTagsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreOriginalTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) MessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) PolicyDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) PolicyDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Responders() AlertPolicyRespondersList {
	var returns AlertPolicyRespondersList
	_jsii_.Get(
		j,
		"responders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) RespondersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"respondersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) TimeRestriction() AlertPolicyTimeRestrictionOutputReference {
	var returns AlertPolicyTimeRestrictionOutputReference
	_jsii_.Get(
		j,
		"timeRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) TimeRestrictionInput() *AlertPolicyTimeRestriction {
	var returns *AlertPolicyTimeRestriction
	_jsii_.Get(
		j,
		"timeRestrictionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/alert_policy opsgenie_alert_policy} Resource.
func NewAlertPolicy(scope constructs.Construct, id *string, config *AlertPolicyConfig) AlertPolicy {
	_init_.Initialize()

	if err := validateNewAlertPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertPolicy{}

	_jsii_.Create(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/alert_policy opsgenie_alert_policy} Resource.
func NewAlertPolicy_Override(a AlertPolicy, scope constructs.Construct, id *string, config *AlertPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
		[]any{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AlertPolicy) SetActions(val *[]*string) {
	if err := j.validateSetActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetAlertDescription(val *string) {
	if err := j.validateSetAlertDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertDescription",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetAlias(val *string) {
	if err := j.validateSetAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetContinuePolicy(val any) {
	if err := j.validateSetContinuePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continuePolicy",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetEntity(val *string) {
	if err := j.validateSetEntityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entity",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetIgnoreOriginalActions(val any) {
	if err := j.validateSetIgnoreOriginalActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreOriginalActions",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetIgnoreOriginalDetails(val any) {
	if err := j.validateSetIgnoreOriginalDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreOriginalDetails",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetIgnoreOriginalResponders(val any) {
	if err := j.validateSetIgnoreOriginalRespondersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreOriginalResponders",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetIgnoreOriginalTags(val any) {
	if err := j.validateSetIgnoreOriginalTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreOriginalTags",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetMessage(val *string) {
	if err := j.validateSetMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetPolicyDescription(val *string) {
	if err := j.validateSetPolicyDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDescription",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy) SetTeamId(val *string) {
	if err := j.validateSetTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

// Generates CDKTF code for importing a AlertPolicy resource upon running "cdktf plan <stack-name>".
func AlertPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAlertPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
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
func AlertPolicy_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateAlertPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlertPolicy_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateAlertPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlertPolicy_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateAlertPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func AlertPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AlertPolicy) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (a *jsiiProxy_AlertPolicy) AddOverride(path *string, value any) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]any{path, value},
	)
}

func (a *jsiiProxy_AlertPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]any{id, provider},
	)
}

func (a *jsiiProxy_AlertPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]any{id},
	)
}

func (a *jsiiProxy_AlertPolicy) MoveTo(moveTarget *string, index any) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (a *jsiiProxy_AlertPolicy) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]any{id},
	)
}

func (a *jsiiProxy_AlertPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (a *jsiiProxy_AlertPolicy) PutFilter(value *AlertPolicyFilter) {
	if err := a.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilter",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertPolicy) PutResponders(value any) {
	if err := a.validatePutRespondersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResponders",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertPolicy) PutTimeRestriction(value *AlertPolicyTimeRestriction) {
	if err := a.validatePutTimeRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeRestriction",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertPolicy) ResetActions() {
	_jsii_.InvokeVoid(
		a,
		"resetActions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetAlertDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetAlertDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetContinuePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetContinuePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetEntity() {
	_jsii_.InvokeVoid(
		a,
		"resetEntity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetIgnoreOriginalActions() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreOriginalActions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetIgnoreOriginalDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreOriginalDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetIgnoreOriginalResponders() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreOriginalResponders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetIgnoreOriginalTags() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreOriginalTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetPolicyDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetResponders() {
	_jsii_.InvokeVoid(
		a,
		"resetResponders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetSource() {
	_jsii_.InvokeVoid(
		a,
		"resetSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetTeamId() {
	_jsii_.InvokeVoid(
		a,
		"resetTeamId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) ResetTimeRestriction() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeRestriction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertPolicy) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
