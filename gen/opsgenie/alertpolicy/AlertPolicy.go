package alertpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/alertpolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy opsgenie_alert_policy}.
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
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContinuePolicy() interface{}
	SetContinuePolicy(val interface{})
	ContinuePolicyInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	IgnoreOriginalActions() interface{}
	SetIgnoreOriginalActions(val interface{})
	IgnoreOriginalActionsInput() interface{}
	IgnoreOriginalDetails() interface{}
	SetIgnoreOriginalDetails(val interface{})
	IgnoreOriginalDetailsInput() interface{}
	IgnoreOriginalResponders() interface{}
	SetIgnoreOriginalResponders(val interface{})
	IgnoreOriginalRespondersInput() interface{}
	IgnoreOriginalTags() interface{}
	SetIgnoreOriginalTags(val interface{})
	IgnoreOriginalTagsInput() interface{}
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Responders() AlertPolicyRespondersList
	RespondersInput() interface{}
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
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeRestriction() AlertPolicyTimeRestrictionOutputReference
	TimeRestrictionInput() *AlertPolicyTimeRestriction
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutFilter(value *AlertPolicyFilter)
	PutResponders(value interface{})
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
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
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

func (j *jsiiProxy_AlertPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) ContinuePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) ContinuePolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertPolicy) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) EnabledInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalActions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreOriginalActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreOriginalActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreOriginalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreOriginalDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalResponders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreOriginalResponders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalRespondersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreOriginalRespondersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreOriginalTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) IgnoreOriginalTagsInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertPolicy) RawOverrides() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertPolicy) RespondersInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
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


// Create a new {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy opsgenie_alert_policy} Resource.
func NewAlertPolicy(scope constructs.Construct, id *string, config *AlertPolicyConfig) AlertPolicy {
	_init_.Initialize()

	if err := validateNewAlertPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertPolicy{}

	_jsii_.Create(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy opsgenie_alert_policy} Resource.
func NewAlertPolicy_Override(a AlertPolicy, scope constructs.Construct, id *string, config *AlertPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AlertPolicy)SetActions(val *[]*string) {
	if err := j.validateSetActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetAlertDescription(val *string) {
	if err := j.validateSetAlertDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertDescription",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetAlias(val *string) {
	if err := j.validateSetAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetContinuePolicy(val interface{}) {
	if err := j.validateSetContinuePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continuePolicy",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetEntity(val *string) {
	if err := j.validateSetEntityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entity",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetIgnoreOriginalActions(val interface{}) {
	if err := j.validateSetIgnoreOriginalActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreOriginalActions",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetIgnoreOriginalDetails(val interface{}) {
	if err := j.validateSetIgnoreOriginalDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreOriginalDetails",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetIgnoreOriginalResponders(val interface{}) {
	if err := j.validateSetIgnoreOriginalRespondersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreOriginalResponders",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetIgnoreOriginalTags(val interface{}) {
	if err := j.validateSetIgnoreOriginalTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreOriginalTags",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetMessage(val *string) {
	if err := j.validateSetMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetPolicyDescription(val *string) {
	if err := j.validateSetPolicyDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDescription",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AlertPolicy)SetTeamId(val *string) {
	if err := j.validateSetTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
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
func AlertPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlertPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlertPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlertPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlertPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlertPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.alertPolicy.AlertPolicy",
		"isTerraformResource",
		[]interface{}{x},
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

func (a *jsiiProxy_AlertPolicy) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AlertPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlertPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlertPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlertPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlertPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlertPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlertPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlertPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlertPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AlertPolicy) PutFilter(value *AlertPolicyFilter) {
	if err := a.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertPolicy) PutResponders(value interface{}) {
	if err := a.validatePutRespondersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResponders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertPolicy) PutTimeRestriction(value *AlertPolicyTimeRestriction) {
	if err := a.validatePutTimeRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeRestriction",
		[]interface{}{value},
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

func (a *jsiiProxy_AlertPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertPolicy) ToMetadata() interface{} {
	var returns interface{}

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

func (a *jsiiProxy_AlertPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

