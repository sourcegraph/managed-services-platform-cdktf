package googledialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledialogflowcxpage/internal"
)

type GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EndpointingTimeoutDuration() *string
	SetEndpointingTimeoutDuration(val *string)
	EndpointingTimeoutDurationInput() *string
	FinishDigit() *string
	SetFinishDigit(val *string)
	FinishDigitInput() *string
	// Experimental.
	Fqn() *string
	InterdigitTimeoutDuration() *string
	SetInterdigitTimeoutDuration(val *string)
	InterdigitTimeoutDurationInput() *string
	InternalValue() *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettings
	SetInternalValue(val *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettings)
	MaxDigits() *float64
	SetMaxDigits(val *float64)
	MaxDigitsInput() *float64
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
	ResetEnabled()
	ResetEndpointingTimeoutDuration()
	ResetFinishDigit()
	ResetInterdigitTimeoutDuration()
	ResetMaxDigits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference
type jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) EndpointingTimeoutDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointingTimeoutDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) EndpointingTimeoutDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointingTimeoutDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) FinishDigit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finishDigit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) FinishDigitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finishDigitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) InterdigitTimeoutDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interdigitTimeoutDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) InterdigitTimeoutDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interdigitTimeoutDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) InternalValue() *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettings {
	var returns *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) MaxDigits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDigits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) MaxDigitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDigitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxPage.GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference_Override(g GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxPage.GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference)SetEndpointingTimeoutDuration(val *string) {
	if err := j.validateSetEndpointingTimeoutDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointingTimeoutDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference)SetFinishDigit(val *string) {
	if err := j.validateSetFinishDigitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finishDigit",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference)SetInterdigitTimeoutDuration(val *string) {
	if err := j.validateSetInterdigitTimeoutDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interdigitTimeoutDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference)SetInternalValue(val *GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference)SetMaxDigits(val *float64) {
	if err := j.validateSetMaxDigitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDigits",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) ResetEndpointingTimeoutDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointingTimeoutDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) ResetFinishDigit() {
	_jsii_.InvokeVoid(
		g,
		"resetFinishDigit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) ResetInterdigitTimeoutDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetInterdigitTimeoutDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) ResetMaxDigits() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxDigits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsDtmfSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

