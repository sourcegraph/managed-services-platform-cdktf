package notificationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/notificationpolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy opsgenie_notification_policy}.
type NotificationPolicy interface {
	cdktf.TerraformResource
	AutoCloseAction() NotificationPolicyAutoCloseActionOutputReference
	AutoCloseActionInput() *NotificationPolicyAutoCloseAction
	AutoRestartAction() NotificationPolicyAutoRestartActionOutputReference
	AutoRestartActionInput() *NotificationPolicyAutoRestartAction
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeDuplicationAction() NotificationPolicyDeDuplicationActionOutputReference
	DeDuplicationActionInput() *NotificationPolicyDeDuplicationAction
	DelayAction() NotificationPolicyDelayActionOutputReference
	DelayActionInput() *NotificationPolicyDelayAction
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Filter() NotificationPolicyFilterOutputReference
	FilterInput() *NotificationPolicyFilter
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PolicyDescription() *string
	SetPolicyDescription(val *string)
	PolicyDescriptionInput() *string
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
	Suppress() interface{}
	SetSuppress(val interface{})
	SuppressInput() interface{}
	TeamId() *string
	SetTeamId(val *string)
	TeamIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeRestriction() NotificationPolicyTimeRestrictionOutputReference
	TimeRestrictionInput() *NotificationPolicyTimeRestriction
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
	PutAutoCloseAction(value *NotificationPolicyAutoCloseAction)
	PutAutoRestartAction(value *NotificationPolicyAutoRestartAction)
	PutDeDuplicationAction(value *NotificationPolicyDeDuplicationAction)
	PutDelayAction(value *NotificationPolicyDelayAction)
	PutFilter(value *NotificationPolicyFilter)
	PutTimeRestriction(value *NotificationPolicyTimeRestriction)
	ResetAutoCloseAction()
	ResetAutoRestartAction()
	ResetDeDuplicationAction()
	ResetDelayAction()
	ResetEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyDescription()
	ResetSuppress()
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

// The jsii proxy struct for NotificationPolicy
type jsiiProxy_NotificationPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NotificationPolicy) AutoCloseAction() NotificationPolicyAutoCloseActionOutputReference {
	var returns NotificationPolicyAutoCloseActionOutputReference
	_jsii_.Get(
		j,
		"autoCloseAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) AutoCloseActionInput() *NotificationPolicyAutoCloseAction {
	var returns *NotificationPolicyAutoCloseAction
	_jsii_.Get(
		j,
		"autoCloseActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) AutoRestartAction() NotificationPolicyAutoRestartActionOutputReference {
	var returns NotificationPolicyAutoRestartActionOutputReference
	_jsii_.Get(
		j,
		"autoRestartAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) AutoRestartActionInput() *NotificationPolicyAutoRestartAction {
	var returns *NotificationPolicyAutoRestartAction
	_jsii_.Get(
		j,
		"autoRestartActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) DeDuplicationAction() NotificationPolicyDeDuplicationActionOutputReference {
	var returns NotificationPolicyDeDuplicationActionOutputReference
	_jsii_.Get(
		j,
		"deDuplicationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) DeDuplicationActionInput() *NotificationPolicyDeDuplicationAction {
	var returns *NotificationPolicyDeDuplicationAction
	_jsii_.Get(
		j,
		"deDuplicationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) DelayAction() NotificationPolicyDelayActionOutputReference {
	var returns NotificationPolicyDelayActionOutputReference
	_jsii_.Get(
		j,
		"delayAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) DelayActionInput() *NotificationPolicyDelayAction {
	var returns *NotificationPolicyDelayAction
	_jsii_.Get(
		j,
		"delayActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Filter() NotificationPolicyFilterOutputReference {
	var returns NotificationPolicyFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) FilterInput() *NotificationPolicyFilter {
	var returns *NotificationPolicyFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) PolicyDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) PolicyDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) Suppress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) SuppressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) TimeRestriction() NotificationPolicyTimeRestrictionOutputReference {
	var returns NotificationPolicyTimeRestrictionOutputReference
	_jsii_.Get(
		j,
		"timeRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationPolicy) TimeRestrictionInput() *NotificationPolicyTimeRestriction {
	var returns *NotificationPolicyTimeRestriction
	_jsii_.Get(
		j,
		"timeRestrictionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy opsgenie_notification_policy} Resource.
func NewNotificationPolicy(scope constructs.Construct, id *string, config *NotificationPolicyConfig) NotificationPolicy {
	_init_.Initialize()

	if err := validateNewNotificationPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotificationPolicy{}

	_jsii_.Create(
		"@cdktf/provider-opsgenie.notificationPolicy.NotificationPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy opsgenie_notification_policy} Resource.
func NewNotificationPolicy_Override(n NotificationPolicy, scope constructs.Construct, id *string, config *NotificationPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opsgenie.notificationPolicy.NotificationPolicy",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetPolicyDescription(val *string) {
	if err := j.validateSetPolicyDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDescription",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetSuppress(val interface{}) {
	if err := j.validateSetSuppressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppress",
		val,
	)
}

func (j *jsiiProxy_NotificationPolicy)SetTeamId(val *string) {
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
func NotificationPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNotificationPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.notificationPolicy.NotificationPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NotificationPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNotificationPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.notificationPolicy.NotificationPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NotificationPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNotificationPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.notificationPolicy.NotificationPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NotificationPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opsgenie.notificationPolicy.NotificationPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NotificationPolicy) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NotificationPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NotificationPolicy) PutAutoCloseAction(value *NotificationPolicyAutoCloseAction) {
	if err := n.validatePutAutoCloseActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAutoCloseAction",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationPolicy) PutAutoRestartAction(value *NotificationPolicyAutoRestartAction) {
	if err := n.validatePutAutoRestartActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAutoRestartAction",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationPolicy) PutDeDuplicationAction(value *NotificationPolicyDeDuplicationAction) {
	if err := n.validatePutDeDuplicationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDeDuplicationAction",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationPolicy) PutDelayAction(value *NotificationPolicyDelayAction) {
	if err := n.validatePutDelayActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDelayAction",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationPolicy) PutFilter(value *NotificationPolicyFilter) {
	if err := n.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putFilter",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationPolicy) PutTimeRestriction(value *NotificationPolicyTimeRestriction) {
	if err := n.validatePutTimeRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeRestriction",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetAutoCloseAction() {
	_jsii_.InvokeVoid(
		n,
		"resetAutoCloseAction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetAutoRestartAction() {
	_jsii_.InvokeVoid(
		n,
		"resetAutoRestartAction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetDeDuplicationAction() {
	_jsii_.InvokeVoid(
		n,
		"resetDeDuplicationAction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetDelayAction() {
	_jsii_.InvokeVoid(
		n,
		"resetDelayAction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetPolicyDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetPolicyDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetSuppress() {
	_jsii_.InvokeVoid(
		n,
		"resetSuppress",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) ResetTimeRestriction() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeRestriction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

