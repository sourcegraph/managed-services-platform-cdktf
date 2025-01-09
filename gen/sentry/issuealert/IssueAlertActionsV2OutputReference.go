package issuealert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/issuealert/internal"
)

type IssueAlertActionsV2OutputReference interface {
	cdktf.ComplexObject
	AzureDevopsCreateTicket() IssueAlertActionsV2AzureDevopsCreateTicketOutputReference
	AzureDevopsCreateTicketInput() interface{}
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
	DiscordNotifyService() IssueAlertActionsV2DiscordNotifyServiceOutputReference
	DiscordNotifyServiceInput() interface{}
	// Experimental.
	Fqn() *string
	GithubCreateTicket() IssueAlertActionsV2GithubCreateTicketOutputReference
	GithubCreateTicketInput() interface{}
	GithubEnterpriseCreateTicket() IssueAlertActionsV2GithubEnterpriseCreateTicketOutputReference
	GithubEnterpriseCreateTicketInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JiraCreateTicket() IssueAlertActionsV2JiraCreateTicketOutputReference
	JiraCreateTicketInput() interface{}
	JiraServerCreateTicket() IssueAlertActionsV2JiraServerCreateTicketOutputReference
	JiraServerCreateTicketInput() interface{}
	MsteamsNotifyService() IssueAlertActionsV2MsteamsNotifyServiceOutputReference
	MsteamsNotifyServiceInput() interface{}
	NotifyEmail() IssueAlertActionsV2NotifyEmailOutputReference
	NotifyEmailInput() interface{}
	NotifyEvent() IssueAlertActionsV2NotifyEventOutputReference
	NotifyEventInput() interface{}
	NotifyEventSentryApp() IssueAlertActionsV2NotifyEventSentryAppOutputReference
	NotifyEventSentryAppInput() interface{}
	NotifyEventService() IssueAlertActionsV2NotifyEventServiceOutputReference
	NotifyEventServiceInput() interface{}
	OpsgenieNotifyTeam() IssueAlertActionsV2OpsgenieNotifyTeamOutputReference
	OpsgenieNotifyTeamInput() interface{}
	PagerdutyNotifyService() IssueAlertActionsV2PagerdutyNotifyServiceOutputReference
	PagerdutyNotifyServiceInput() interface{}
	SlackNotifyService() IssueAlertActionsV2SlackNotifyServiceOutputReference
	SlackNotifyServiceInput() interface{}
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
	PutAzureDevopsCreateTicket(value *IssueAlertActionsV2AzureDevopsCreateTicket)
	PutDiscordNotifyService(value *IssueAlertActionsV2DiscordNotifyService)
	PutGithubCreateTicket(value *IssueAlertActionsV2GithubCreateTicket)
	PutGithubEnterpriseCreateTicket(value *IssueAlertActionsV2GithubEnterpriseCreateTicket)
	PutJiraCreateTicket(value *IssueAlertActionsV2JiraCreateTicket)
	PutJiraServerCreateTicket(value *IssueAlertActionsV2JiraServerCreateTicket)
	PutMsteamsNotifyService(value *IssueAlertActionsV2MsteamsNotifyService)
	PutNotifyEmail(value *IssueAlertActionsV2NotifyEmail)
	PutNotifyEvent(value *IssueAlertActionsV2NotifyEvent)
	PutNotifyEventSentryApp(value *IssueAlertActionsV2NotifyEventSentryApp)
	PutNotifyEventService(value *IssueAlertActionsV2NotifyEventService)
	PutOpsgenieNotifyTeam(value *IssueAlertActionsV2OpsgenieNotifyTeam)
	PutPagerdutyNotifyService(value *IssueAlertActionsV2PagerdutyNotifyService)
	PutSlackNotifyService(value *IssueAlertActionsV2SlackNotifyService)
	ResetAzureDevopsCreateTicket()
	ResetDiscordNotifyService()
	ResetGithubCreateTicket()
	ResetGithubEnterpriseCreateTicket()
	ResetJiraCreateTicket()
	ResetJiraServerCreateTicket()
	ResetMsteamsNotifyService()
	ResetNotifyEmail()
	ResetNotifyEvent()
	ResetNotifyEventSentryApp()
	ResetNotifyEventService()
	ResetOpsgenieNotifyTeam()
	ResetPagerdutyNotifyService()
	ResetSlackNotifyService()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IssueAlertActionsV2OutputReference
type jsiiProxy_IssueAlertActionsV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) AzureDevopsCreateTicket() IssueAlertActionsV2AzureDevopsCreateTicketOutputReference {
	var returns IssueAlertActionsV2AzureDevopsCreateTicketOutputReference
	_jsii_.Get(
		j,
		"azureDevopsCreateTicket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) AzureDevopsCreateTicketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureDevopsCreateTicketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) DiscordNotifyService() IssueAlertActionsV2DiscordNotifyServiceOutputReference {
	var returns IssueAlertActionsV2DiscordNotifyServiceOutputReference
	_jsii_.Get(
		j,
		"discordNotifyService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) DiscordNotifyServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discordNotifyServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) GithubCreateTicket() IssueAlertActionsV2GithubCreateTicketOutputReference {
	var returns IssueAlertActionsV2GithubCreateTicketOutputReference
	_jsii_.Get(
		j,
		"githubCreateTicket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) GithubCreateTicketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubCreateTicketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) GithubEnterpriseCreateTicket() IssueAlertActionsV2GithubEnterpriseCreateTicketOutputReference {
	var returns IssueAlertActionsV2GithubEnterpriseCreateTicketOutputReference
	_jsii_.Get(
		j,
		"githubEnterpriseCreateTicket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) GithubEnterpriseCreateTicketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubEnterpriseCreateTicketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) JiraCreateTicket() IssueAlertActionsV2JiraCreateTicketOutputReference {
	var returns IssueAlertActionsV2JiraCreateTicketOutputReference
	_jsii_.Get(
		j,
		"jiraCreateTicket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) JiraCreateTicketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraCreateTicketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) JiraServerCreateTicket() IssueAlertActionsV2JiraServerCreateTicketOutputReference {
	var returns IssueAlertActionsV2JiraServerCreateTicketOutputReference
	_jsii_.Get(
		j,
		"jiraServerCreateTicket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) JiraServerCreateTicketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jiraServerCreateTicketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) MsteamsNotifyService() IssueAlertActionsV2MsteamsNotifyServiceOutputReference {
	var returns IssueAlertActionsV2MsteamsNotifyServiceOutputReference
	_jsii_.Get(
		j,
		"msteamsNotifyService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) MsteamsNotifyServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"msteamsNotifyServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) NotifyEmail() IssueAlertActionsV2NotifyEmailOutputReference {
	var returns IssueAlertActionsV2NotifyEmailOutputReference
	_jsii_.Get(
		j,
		"notifyEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) NotifyEmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) NotifyEvent() IssueAlertActionsV2NotifyEventOutputReference {
	var returns IssueAlertActionsV2NotifyEventOutputReference
	_jsii_.Get(
		j,
		"notifyEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) NotifyEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyEventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) NotifyEventSentryApp() IssueAlertActionsV2NotifyEventSentryAppOutputReference {
	var returns IssueAlertActionsV2NotifyEventSentryAppOutputReference
	_jsii_.Get(
		j,
		"notifyEventSentryApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) NotifyEventSentryAppInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyEventSentryAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) NotifyEventService() IssueAlertActionsV2NotifyEventServiceOutputReference {
	var returns IssueAlertActionsV2NotifyEventServiceOutputReference
	_jsii_.Get(
		j,
		"notifyEventService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) NotifyEventServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyEventServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) OpsgenieNotifyTeam() IssueAlertActionsV2OpsgenieNotifyTeamOutputReference {
	var returns IssueAlertActionsV2OpsgenieNotifyTeamOutputReference
	_jsii_.Get(
		j,
		"opsgenieNotifyTeam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) OpsgenieNotifyTeamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opsgenieNotifyTeamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) PagerdutyNotifyService() IssueAlertActionsV2PagerdutyNotifyServiceOutputReference {
	var returns IssueAlertActionsV2PagerdutyNotifyServiceOutputReference
	_jsii_.Get(
		j,
		"pagerdutyNotifyService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) PagerdutyNotifyServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pagerdutyNotifyServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) SlackNotifyService() IssueAlertActionsV2SlackNotifyServiceOutputReference {
	var returns IssueAlertActionsV2SlackNotifyServiceOutputReference
	_jsii_.Get(
		j,
		"slackNotifyService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) SlackNotifyServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slackNotifyServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIssueAlertActionsV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IssueAlertActionsV2OutputReference {
	_init_.Initialize()

	if err := validateNewIssueAlertActionsV2OutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IssueAlertActionsV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.issueAlert.IssueAlertActionsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIssueAlertActionsV2OutputReference_Override(i IssueAlertActionsV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.issueAlert.IssueAlertActionsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IssueAlertActionsV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutAzureDevopsCreateTicket(value *IssueAlertActionsV2AzureDevopsCreateTicket) {
	if err := i.validatePutAzureDevopsCreateTicketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAzureDevopsCreateTicket",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutDiscordNotifyService(value *IssueAlertActionsV2DiscordNotifyService) {
	if err := i.validatePutDiscordNotifyServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDiscordNotifyService",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutGithubCreateTicket(value *IssueAlertActionsV2GithubCreateTicket) {
	if err := i.validatePutGithubCreateTicketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putGithubCreateTicket",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutGithubEnterpriseCreateTicket(value *IssueAlertActionsV2GithubEnterpriseCreateTicket) {
	if err := i.validatePutGithubEnterpriseCreateTicketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putGithubEnterpriseCreateTicket",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutJiraCreateTicket(value *IssueAlertActionsV2JiraCreateTicket) {
	if err := i.validatePutJiraCreateTicketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putJiraCreateTicket",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutJiraServerCreateTicket(value *IssueAlertActionsV2JiraServerCreateTicket) {
	if err := i.validatePutJiraServerCreateTicketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putJiraServerCreateTicket",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutMsteamsNotifyService(value *IssueAlertActionsV2MsteamsNotifyService) {
	if err := i.validatePutMsteamsNotifyServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMsteamsNotifyService",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutNotifyEmail(value *IssueAlertActionsV2NotifyEmail) {
	if err := i.validatePutNotifyEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNotifyEmail",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutNotifyEvent(value *IssueAlertActionsV2NotifyEvent) {
	if err := i.validatePutNotifyEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNotifyEvent",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutNotifyEventSentryApp(value *IssueAlertActionsV2NotifyEventSentryApp) {
	if err := i.validatePutNotifyEventSentryAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNotifyEventSentryApp",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutNotifyEventService(value *IssueAlertActionsV2NotifyEventService) {
	if err := i.validatePutNotifyEventServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNotifyEventService",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutOpsgenieNotifyTeam(value *IssueAlertActionsV2OpsgenieNotifyTeam) {
	if err := i.validatePutOpsgenieNotifyTeamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOpsgenieNotifyTeam",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutPagerdutyNotifyService(value *IssueAlertActionsV2PagerdutyNotifyService) {
	if err := i.validatePutPagerdutyNotifyServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPagerdutyNotifyService",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) PutSlackNotifyService(value *IssueAlertActionsV2SlackNotifyService) {
	if err := i.validatePutSlackNotifyServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSlackNotifyService",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetAzureDevopsCreateTicket() {
	_jsii_.InvokeVoid(
		i,
		"resetAzureDevopsCreateTicket",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetDiscordNotifyService() {
	_jsii_.InvokeVoid(
		i,
		"resetDiscordNotifyService",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetGithubCreateTicket() {
	_jsii_.InvokeVoid(
		i,
		"resetGithubCreateTicket",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetGithubEnterpriseCreateTicket() {
	_jsii_.InvokeVoid(
		i,
		"resetGithubEnterpriseCreateTicket",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetJiraCreateTicket() {
	_jsii_.InvokeVoid(
		i,
		"resetJiraCreateTicket",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetJiraServerCreateTicket() {
	_jsii_.InvokeVoid(
		i,
		"resetJiraServerCreateTicket",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetMsteamsNotifyService() {
	_jsii_.InvokeVoid(
		i,
		"resetMsteamsNotifyService",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetNotifyEmail() {
	_jsii_.InvokeVoid(
		i,
		"resetNotifyEmail",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetNotifyEvent() {
	_jsii_.InvokeVoid(
		i,
		"resetNotifyEvent",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetNotifyEventSentryApp() {
	_jsii_.InvokeVoid(
		i,
		"resetNotifyEventSentryApp",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetNotifyEventService() {
	_jsii_.InvokeVoid(
		i,
		"resetNotifyEventService",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetOpsgenieNotifyTeam() {
	_jsii_.InvokeVoid(
		i,
		"resetOpsgenieNotifyTeam",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetPagerdutyNotifyService() {
	_jsii_.InvokeVoid(
		i,
		"resetPagerdutyNotifyService",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ResetSlackNotifyService() {
	_jsii_.InvokeVoid(
		i,
		"resetSlackNotifyService",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IssueAlertActionsV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

