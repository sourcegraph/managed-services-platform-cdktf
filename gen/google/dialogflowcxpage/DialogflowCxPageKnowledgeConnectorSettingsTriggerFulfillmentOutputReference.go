package dialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/dialogflowcxpage/internal"
)

type DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference interface {
	cdktf.ComplexObject
	AdvancedSettings() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsOutputReference
	AdvancedSettingsInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettings
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
	ConditionalCases() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentConditionalCasesList
	ConditionalCasesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableGenerativeFallback() interface{}
	SetEnableGenerativeFallback(val interface{})
	EnableGenerativeFallbackInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillment
	SetInternalValue(val *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillment)
	Messages() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesList
	MessagesInput() interface{}
	ReturnPartialResponses() interface{}
	SetReturnPartialResponses(val interface{})
	ReturnPartialResponsesInput() interface{}
	SetParameterActions() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentSetParameterActionsList
	SetParameterActionsInput() interface{}
	Tag() *string
	SetTag(val *string)
	TagInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Webhook() *string
	SetWebhook(val *string)
	WebhookInput() *string
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
	PutAdvancedSettings(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettings)
	PutConditionalCases(value interface{})
	PutMessages(value interface{})
	PutSetParameterActions(value interface{})
	ResetAdvancedSettings()
	ResetConditionalCases()
	ResetEnableGenerativeFallback()
	ResetMessages()
	ResetReturnPartialResponses()
	ResetSetParameterActions()
	ResetTag()
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

// The jsii proxy struct for DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference
type jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) AdvancedSettings() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsOutputReference {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsOutputReference
	_jsii_.Get(
		j,
		"advancedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) AdvancedSettingsInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettings {
	var returns *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettings
	_jsii_.Get(
		j,
		"advancedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ConditionalCases() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentConditionalCasesList {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentConditionalCasesList
	_jsii_.Get(
		j,
		"conditionalCases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ConditionalCasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionalCasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) EnableGenerativeFallback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGenerativeFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) EnableGenerativeFallbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGenerativeFallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) InternalValue() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillment {
	var returns *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) Messages() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesList {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesList
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) MessagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ReturnPartialResponses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnPartialResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ReturnPartialResponsesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnPartialResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) SetParameterActions() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentSetParameterActionsList {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentSetParameterActionsList
	_jsii_.Get(
		j,
		"setParameterActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) SetParameterActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setParameterActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) Webhook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) WebhookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference_Override(d DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference)SetEnableGenerativeFallback(val interface{}) {
	if err := j.validateSetEnableGenerativeFallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableGenerativeFallback",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference)SetInternalValue(val *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference)SetReturnPartialResponses(val interface{}) {
	if err := j.validateSetReturnPartialResponsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnPartialResponses",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference)SetTag(val *string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference)SetWebhook(val *string) {
	if err := j.validateSetWebhookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhook",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) PutAdvancedSettings(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettings) {
	if err := d.validatePutAdvancedSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAdvancedSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) PutConditionalCases(value interface{}) {
	if err := d.validatePutConditionalCasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConditionalCases",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) PutMessages(value interface{}) {
	if err := d.validatePutMessagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMessages",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) PutSetParameterActions(value interface{}) {
	if err := d.validatePutSetParameterActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSetParameterActions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ResetAdvancedSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetAdvancedSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ResetConditionalCases() {
	_jsii_.InvokeVoid(
		d,
		"resetConditionalCases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ResetEnableGenerativeFallback() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableGenerativeFallback",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ResetMessages() {
	_jsii_.InvokeVoid(
		d,
		"resetMessages",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ResetReturnPartialResponses() {
	_jsii_.InvokeVoid(
		d,
		"resetReturnPartialResponses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ResetSetParameterActions() {
	_jsii_.InvokeVoid(
		d,
		"resetSetParameterActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		d,
		"resetTag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ResetWebhook() {
	_jsii_.InvokeVoid(
		d,
		"resetWebhook",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

