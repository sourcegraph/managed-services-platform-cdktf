package alert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/alert/internal"
)

type AlertActionFiltersActionsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Discord() AlertActionFiltersActionsDiscordOutputReference
	DiscordInput() interface{}
	Email() AlertActionFiltersActionsEmailOutputReference
	EmailInput() interface{}
	// Experimental.
	Fqn() *string
	Github() AlertActionFiltersActionsGithubOutputReference
	GithubInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Jira() AlertActionFiltersActionsJiraOutputReference
	JiraInput() interface{}
	JiraServer() AlertActionFiltersActionsJiraServerOutputReference
	JiraServerInput() interface{}
	Msteams() AlertActionFiltersActionsMsteamsOutputReference
	MsteamsInput() interface{}
	Opsgenie() AlertActionFiltersActionsOpsgenieOutputReference
	OpsgenieInput() interface{}
	Pagerduty() AlertActionFiltersActionsPagerdutyOutputReference
	PagerdutyInput() interface{}
	Plugin() AlertActionFiltersActionsPluginOutputReference
	PluginInput() interface{}
	SentryApp() AlertActionFiltersActionsSentryAppOutputReference
	SentryAppInput() interface{}
	Slack() AlertActionFiltersActionsSlackOutputReference
	SlackInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Vsts() AlertActionFiltersActionsVstsOutputReference
	VstsInput() interface{}
	Webhook() AlertActionFiltersActionsWebhookOutputReference
	WebhookInput() interface{}
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
	PutDiscord(value *AlertActionFiltersActionsDiscord)
	PutEmail(value *AlertActionFiltersActionsEmail)
	PutGithub(value *AlertActionFiltersActionsGithub)
	PutJira(value *AlertActionFiltersActionsJira)
	PutJiraServer(value *AlertActionFiltersActionsJiraServer)
	PutMsteams(value *AlertActionFiltersActionsMsteams)
	PutOpsgenie(value *AlertActionFiltersActionsOpsgenie)
	PutPagerduty(value *AlertActionFiltersActionsPagerduty)
	PutPlugin(value *AlertActionFiltersActionsPlugin)
	PutSentryApp(value *AlertActionFiltersActionsSentryApp)
	PutSlack(value *AlertActionFiltersActionsSlack)
	PutVsts(value *AlertActionFiltersActionsVsts)
	PutWebhook(value *AlertActionFiltersActionsWebhook)
	ResetDiscord()
	ResetEmail()
	ResetGithub()
	ResetJira()
	ResetJiraServer()
	ResetMsteams()
	ResetOpsgenie()
	ResetPagerduty()
	ResetPlugin()
	ResetSentryApp()
	ResetSlack()
	ResetVsts()
	ResetWebhook()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertActionFiltersActionsOutputReference
type jsiiProxy_AlertActionFiltersActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Discord() AlertActionFiltersActionsDiscordOutputReference {
	var returns AlertActionFiltersActionsDiscordOutputReference
	_jsii_.Get(
		j,
		"discord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) DiscordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Email() AlertActionFiltersActionsEmailOutputReference {
	var returns AlertActionFiltersActionsEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) EmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Github() AlertActionFiltersActionsGithubOutputReference {
	var returns AlertActionFiltersActionsGithubOutputReference
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) GithubInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Jira() AlertActionFiltersActionsJiraOutputReference {
	var returns AlertActionFiltersActionsJiraOutputReference
	_jsii_.Get(
		j,
		"jira",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) JiraInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) JiraServer() AlertActionFiltersActionsJiraServerOutputReference {
	var returns AlertActionFiltersActionsJiraServerOutputReference
	_jsii_.Get(
		j,
		"jiraServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) JiraServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Msteams() AlertActionFiltersActionsMsteamsOutputReference {
	var returns AlertActionFiltersActionsMsteamsOutputReference
	_jsii_.Get(
		j,
		"msteams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) MsteamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"msteamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Opsgenie() AlertActionFiltersActionsOpsgenieOutputReference {
	var returns AlertActionFiltersActionsOpsgenieOutputReference
	_jsii_.Get(
		j,
		"opsgenie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) OpsgenieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opsgenieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Pagerduty() AlertActionFiltersActionsPagerdutyOutputReference {
	var returns AlertActionFiltersActionsPagerdutyOutputReference
	_jsii_.Get(
		j,
		"pagerduty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) PagerdutyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pagerdutyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Plugin() AlertActionFiltersActionsPluginOutputReference {
	var returns AlertActionFiltersActionsPluginOutputReference
	_jsii_.Get(
		j,
		"plugin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) PluginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pluginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) SentryApp() AlertActionFiltersActionsSentryAppOutputReference {
	var returns AlertActionFiltersActionsSentryAppOutputReference
	_jsii_.Get(
		j,
		"sentryApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) SentryAppInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sentryAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Slack() AlertActionFiltersActionsSlackOutputReference {
	var returns AlertActionFiltersActionsSlackOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) SlackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Vsts() AlertActionFiltersActionsVstsOutputReference {
	var returns AlertActionFiltersActionsVstsOutputReference
	_jsii_.Get(
		j,
		"vsts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) VstsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vstsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) Webhook() AlertActionFiltersActionsWebhookOutputReference {
	var returns AlertActionFiltersActionsWebhookOutputReference
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference) WebhookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}


func NewAlertActionFiltersActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlertActionFiltersActionsOutputReference {
	_init_.Initialize()

	if err := validateNewAlertActionFiltersActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertActionFiltersActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.alert.AlertActionFiltersActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlertActionFiltersActionsOutputReference_Override(a AlertActionFiltersActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.alert.AlertActionFiltersActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertActionFiltersActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutDiscord(value *AlertActionFiltersActionsDiscord) {
	if err := a.validatePutDiscordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDiscord",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutEmail(value *AlertActionFiltersActionsEmail) {
	if err := a.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmail",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutGithub(value *AlertActionFiltersActionsGithub) {
	if err := a.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGithub",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutJira(value *AlertActionFiltersActionsJira) {
	if err := a.validatePutJiraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJira",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutJiraServer(value *AlertActionFiltersActionsJiraServer) {
	if err := a.validatePutJiraServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJiraServer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutMsteams(value *AlertActionFiltersActionsMsteams) {
	if err := a.validatePutMsteamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMsteams",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutOpsgenie(value *AlertActionFiltersActionsOpsgenie) {
	if err := a.validatePutOpsgenieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpsgenie",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutPagerduty(value *AlertActionFiltersActionsPagerduty) {
	if err := a.validatePutPagerdutyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPagerduty",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutPlugin(value *AlertActionFiltersActionsPlugin) {
	if err := a.validatePutPluginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPlugin",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutSentryApp(value *AlertActionFiltersActionsSentryApp) {
	if err := a.validatePutSentryAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSentryApp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutSlack(value *AlertActionFiltersActionsSlack) {
	if err := a.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlack",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutVsts(value *AlertActionFiltersActionsVsts) {
	if err := a.validatePutVstsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVsts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) PutWebhook(value *AlertActionFiltersActionsWebhook) {
	if err := a.validatePutWebhookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWebhook",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetDiscord() {
	_jsii_.InvokeVoid(
		a,
		"resetDiscord",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		a,
		"resetEmail",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetGithub() {
	_jsii_.InvokeVoid(
		a,
		"resetGithub",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetJira() {
	_jsii_.InvokeVoid(
		a,
		"resetJira",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetJiraServer() {
	_jsii_.InvokeVoid(
		a,
		"resetJiraServer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetMsteams() {
	_jsii_.InvokeVoid(
		a,
		"resetMsteams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetOpsgenie() {
	_jsii_.InvokeVoid(
		a,
		"resetOpsgenie",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetPagerduty() {
	_jsii_.InvokeVoid(
		a,
		"resetPagerduty",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetPlugin() {
	_jsii_.InvokeVoid(
		a,
		"resetPlugin",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetSentryApp() {
	_jsii_.InvokeVoid(
		a,
		"resetSentryApp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetVsts() {
	_jsii_.InvokeVoid(
		a,
		"resetVsts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ResetWebhook() {
	_jsii_.InvokeVoid(
		a,
		"resetWebhook",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlertActionFiltersActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

