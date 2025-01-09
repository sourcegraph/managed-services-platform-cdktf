package datasentryissuealert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/datasentryissuealert/internal"
)

type DataSentryIssueAlertActionsV2OutputReference interface {
	cdktf.ComplexObject
	AzureDevopsCreateTicket() DataSentryIssueAlertActionsV2AzureDevopsCreateTicketOutputReference
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
	DiscordNotifyService() DataSentryIssueAlertActionsV2DiscordNotifyServiceOutputReference
	// Experimental.
	Fqn() *string
	GithubCreateTicket() DataSentryIssueAlertActionsV2GithubCreateTicketOutputReference
	GithubEnterpriseCreateTicket() DataSentryIssueAlertActionsV2GithubEnterpriseCreateTicketOutputReference
	InternalValue() *DataSentryIssueAlertActionsV2
	SetInternalValue(val *DataSentryIssueAlertActionsV2)
	JiraCreateTicket() DataSentryIssueAlertActionsV2JiraCreateTicketOutputReference
	JiraServerCreateTicket() DataSentryIssueAlertActionsV2JiraServerCreateTicketOutputReference
	MsteamsNotifyService() DataSentryIssueAlertActionsV2MsteamsNotifyServiceOutputReference
	NotifyEmail() DataSentryIssueAlertActionsV2NotifyEmailOutputReference
	NotifyEvent() DataSentryIssueAlertActionsV2NotifyEventOutputReference
	NotifyEventSentryApp() DataSentryIssueAlertActionsV2NotifyEventSentryAppOutputReference
	NotifyEventService() DataSentryIssueAlertActionsV2NotifyEventServiceOutputReference
	OpsgenieNotifyTeam() DataSentryIssueAlertActionsV2OpsgenieNotifyTeamOutputReference
	PagerdutyNotifyService() DataSentryIssueAlertActionsV2PagerdutyNotifyServiceOutputReference
	SlackNotifyService() DataSentryIssueAlertActionsV2SlackNotifyServiceOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataSentryIssueAlertActionsV2OutputReference
type jsiiProxy_DataSentryIssueAlertActionsV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) AzureDevopsCreateTicket() DataSentryIssueAlertActionsV2AzureDevopsCreateTicketOutputReference {
	var returns DataSentryIssueAlertActionsV2AzureDevopsCreateTicketOutputReference
	_jsii_.Get(
		j,
		"azureDevopsCreateTicket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) DiscordNotifyService() DataSentryIssueAlertActionsV2DiscordNotifyServiceOutputReference {
	var returns DataSentryIssueAlertActionsV2DiscordNotifyServiceOutputReference
	_jsii_.Get(
		j,
		"discordNotifyService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) GithubCreateTicket() DataSentryIssueAlertActionsV2GithubCreateTicketOutputReference {
	var returns DataSentryIssueAlertActionsV2GithubCreateTicketOutputReference
	_jsii_.Get(
		j,
		"githubCreateTicket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) GithubEnterpriseCreateTicket() DataSentryIssueAlertActionsV2GithubEnterpriseCreateTicketOutputReference {
	var returns DataSentryIssueAlertActionsV2GithubEnterpriseCreateTicketOutputReference
	_jsii_.Get(
		j,
		"githubEnterpriseCreateTicket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) InternalValue() *DataSentryIssueAlertActionsV2 {
	var returns *DataSentryIssueAlertActionsV2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) JiraCreateTicket() DataSentryIssueAlertActionsV2JiraCreateTicketOutputReference {
	var returns DataSentryIssueAlertActionsV2JiraCreateTicketOutputReference
	_jsii_.Get(
		j,
		"jiraCreateTicket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) JiraServerCreateTicket() DataSentryIssueAlertActionsV2JiraServerCreateTicketOutputReference {
	var returns DataSentryIssueAlertActionsV2JiraServerCreateTicketOutputReference
	_jsii_.Get(
		j,
		"jiraServerCreateTicket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) MsteamsNotifyService() DataSentryIssueAlertActionsV2MsteamsNotifyServiceOutputReference {
	var returns DataSentryIssueAlertActionsV2MsteamsNotifyServiceOutputReference
	_jsii_.Get(
		j,
		"msteamsNotifyService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) NotifyEmail() DataSentryIssueAlertActionsV2NotifyEmailOutputReference {
	var returns DataSentryIssueAlertActionsV2NotifyEmailOutputReference
	_jsii_.Get(
		j,
		"notifyEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) NotifyEvent() DataSentryIssueAlertActionsV2NotifyEventOutputReference {
	var returns DataSentryIssueAlertActionsV2NotifyEventOutputReference
	_jsii_.Get(
		j,
		"notifyEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) NotifyEventSentryApp() DataSentryIssueAlertActionsV2NotifyEventSentryAppOutputReference {
	var returns DataSentryIssueAlertActionsV2NotifyEventSentryAppOutputReference
	_jsii_.Get(
		j,
		"notifyEventSentryApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) NotifyEventService() DataSentryIssueAlertActionsV2NotifyEventServiceOutputReference {
	var returns DataSentryIssueAlertActionsV2NotifyEventServiceOutputReference
	_jsii_.Get(
		j,
		"notifyEventService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) OpsgenieNotifyTeam() DataSentryIssueAlertActionsV2OpsgenieNotifyTeamOutputReference {
	var returns DataSentryIssueAlertActionsV2OpsgenieNotifyTeamOutputReference
	_jsii_.Get(
		j,
		"opsgenieNotifyTeam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) PagerdutyNotifyService() DataSentryIssueAlertActionsV2PagerdutyNotifyServiceOutputReference {
	var returns DataSentryIssueAlertActionsV2PagerdutyNotifyServiceOutputReference
	_jsii_.Get(
		j,
		"pagerdutyNotifyService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) SlackNotifyService() DataSentryIssueAlertActionsV2SlackNotifyServiceOutputReference {
	var returns DataSentryIssueAlertActionsV2SlackNotifyServiceOutputReference
	_jsii_.Get(
		j,
		"slackNotifyService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataSentryIssueAlertActionsV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSentryIssueAlertActionsV2OutputReference {
	_init_.Initialize()

	if err := validateNewDataSentryIssueAlertActionsV2OutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSentryIssueAlertActionsV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.dataSentryIssueAlert.DataSentryIssueAlertActionsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSentryIssueAlertActionsV2OutputReference_Override(d DataSentryIssueAlertActionsV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.dataSentryIssueAlert.DataSentryIssueAlertActionsV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference)SetInternalValue(val *DataSentryIssueAlertActionsV2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSentryIssueAlertActionsV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

