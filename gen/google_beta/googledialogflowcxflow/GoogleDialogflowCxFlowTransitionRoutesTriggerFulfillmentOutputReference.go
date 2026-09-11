package googledialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledialogflowcxflow/internal"
)

type GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConditionalCases() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentConditionalCasesList
	ConditionalCasesInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillment
	SetInternalValue(val *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillment)
	Messages() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesList
	MessagesInput() any
	ReturnPartialResponses() any
	SetReturnPartialResponses(val any)
	ReturnPartialResponsesInput() any
	SetParameterActions() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentSetParameterActionsList
	SetParameterActionsInput() any
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutConditionalCases(value any)
	PutMessages(value any)
	PutSetParameterActions(value any)
	ResetConditionalCases()
	ResetMessages()
	ResetReturnPartialResponses()
	ResetSetParameterActions()
	ResetTag()
	ResetWebhook()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference
type jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ConditionalCases() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentConditionalCasesList {
	var returns GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentConditionalCasesList
	_jsii_.Get(
		j,
		"conditionalCases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ConditionalCasesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"conditionalCasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) InternalValue() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillment {
	var returns *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) Messages() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesList {
	var returns GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesList
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) MessagesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"messagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ReturnPartialResponses() any {
	var returns any
	_jsii_.Get(
		j,
		"returnPartialResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ReturnPartialResponsesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"returnPartialResponsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) SetParameterActions() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentSetParameterActionsList {
	var returns GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentSetParameterActionsList
	_jsii_.Get(
		j,
		"setParameterActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) SetParameterActionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"setParameterActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) Webhook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) WebhookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookInput",
		&returns,
	)
	return returns
}

func NewGoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference_Override(g GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) SetInternalValue(val *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) SetReturnPartialResponses(val any) {
	if err := j.validateSetReturnPartialResponsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"returnPartialResponses",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) SetTag(val *string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) SetWebhook(val *string) {
	if err := j.validateSetWebhookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhook",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) PutConditionalCases(value any) {
	if err := g.validatePutConditionalCasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConditionalCases",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) PutMessages(value any) {
	if err := g.validatePutMessagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMessages",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) PutSetParameterActions(value any) {
	if err := g.validatePutSetParameterActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSetParameterActions",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ResetConditionalCases() {
	_jsii_.InvokeVoid(
		g,
		"resetConditionalCases",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ResetMessages() {
	_jsii_.InvokeVoid(
		g,
		"resetMessages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ResetReturnPartialResponses() {
	_jsii_.InvokeVoid(
		g,
		"resetReturnPartialResponses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ResetSetParameterActions() {
	_jsii_.InvokeVoid(
		g,
		"resetSetParameterActions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		g,
		"resetTag",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ResetWebhook() {
	_jsii_.InvokeVoid(
		g,
		"resetWebhook",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
