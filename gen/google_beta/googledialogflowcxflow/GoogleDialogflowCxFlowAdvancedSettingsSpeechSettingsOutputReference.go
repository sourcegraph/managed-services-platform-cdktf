package googledialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledialogflowcxflow/internal"
)

type GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EndpointerSensitivity() *float64
	SetEndpointerSensitivity(val *float64)
	EndpointerSensitivityInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowCxFlowAdvancedSettingsSpeechSettings
	SetInternalValue(val *GoogleDialogflowCxFlowAdvancedSettingsSpeechSettings)
	Models() *map[string]*string
	SetModels(val *map[string]*string)
	ModelsInput() *map[string]*string
	NoSpeechTimeout() *string
	SetNoSpeechTimeout(val *string)
	NoSpeechTimeoutInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseTimeoutBasedEndpointing() any
	SetUseTimeoutBasedEndpointing(val any)
	UseTimeoutBasedEndpointingInput() any
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
	ResetEndpointerSensitivity()
	ResetModels()
	ResetNoSpeechTimeout()
	ResetUseTimeoutBasedEndpointing()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference
type jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) EndpointerSensitivity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endpointerSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) EndpointerSensitivityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endpointerSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) InternalValue() *GoogleDialogflowCxFlowAdvancedSettingsSpeechSettings {
	var returns *GoogleDialogflowCxFlowAdvancedSettingsSpeechSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) Models() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"models",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ModelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"modelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) NoSpeechTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noSpeechTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) NoSpeechTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noSpeechTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) UseTimeoutBasedEndpointing() any {
	var returns any
	_jsii_.Get(
		j,
		"useTimeoutBasedEndpointing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) UseTimeoutBasedEndpointingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useTimeoutBasedEndpointingInput",
		&returns,
	)
	return returns
}

func NewGoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference_Override(g GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetEndpointerSensitivity(val *float64) {
	if err := j.validateSetEndpointerSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointerSensitivity",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetInternalValue(val *GoogleDialogflowCxFlowAdvancedSettingsSpeechSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetModels(val *map[string]*string) {
	if err := j.validateSetModelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"models",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetNoSpeechTimeout(val *string) {
	if err := j.validateSetNoSpeechTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSpeechTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) SetUseTimeoutBasedEndpointing(val any) {
	if err := j.validateSetUseTimeoutBasedEndpointingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTimeoutBasedEndpointing",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ResetEndpointerSensitivity() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointerSensitivity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ResetModels() {
	_jsii_.InvokeVoid(
		g,
		"resetModels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ResetNoSpeechTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetNoSpeechTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ResetUseTimeoutBasedEndpointing() {
	_jsii_.InvokeVoid(
		g,
		"resetUseTimeoutBasedEndpointing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
