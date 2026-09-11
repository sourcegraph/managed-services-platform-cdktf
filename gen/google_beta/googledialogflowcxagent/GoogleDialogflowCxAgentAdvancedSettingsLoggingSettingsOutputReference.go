package googledialogflowcxagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledialogflowcxagent/internal"
)

type GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference interface {
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
	EnableConsentBasedRedaction() any
	SetEnableConsentBasedRedaction(val any)
	EnableConsentBasedRedactionInput() any
	EnableInteractionLogging() any
	SetEnableInteractionLogging(val any)
	EnableInteractionLoggingInput() any
	EnableStackdriverLogging() any
	SetEnableStackdriverLogging(val any)
	EnableStackdriverLoggingInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowCxAgentAdvancedSettingsLoggingSettings
	SetInternalValue(val *GoogleDialogflowCxAgentAdvancedSettingsLoggingSettings)
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
	ResetEnableConsentBasedRedaction()
	ResetEnableInteractionLogging()
	ResetEnableStackdriverLogging()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference
type jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableConsentBasedRedaction() any {
	var returns any
	_jsii_.Get(
		j,
		"enableConsentBasedRedaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableConsentBasedRedactionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableConsentBasedRedactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableInteractionLogging() any {
	var returns any
	_jsii_.Get(
		j,
		"enableInteractionLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableInteractionLoggingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableInteractionLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableStackdriverLogging() any {
	var returns any
	_jsii_.Get(
		j,
		"enableStackdriverLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) EnableStackdriverLoggingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableStackdriverLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) InternalValue() *GoogleDialogflowCxAgentAdvancedSettingsLoggingSettings {
	var returns *GoogleDialogflowCxAgentAdvancedSettingsLoggingSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxAgent.GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference_Override(g GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxAgent.GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetEnableConsentBasedRedaction(val any) {
	if err := j.validateSetEnableConsentBasedRedactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConsentBasedRedaction",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetEnableInteractionLogging(val any) {
	if err := j.validateSetEnableInteractionLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInteractionLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetEnableStackdriverLogging(val any) {
	if err := j.validateSetEnableStackdriverLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStackdriverLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetInternalValue(val *GoogleDialogflowCxAgentAdvancedSettingsLoggingSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ResetEnableConsentBasedRedaction() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableConsentBasedRedaction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ResetEnableInteractionLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableInteractionLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ResetEnableStackdriverLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableStackdriverLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleDialogflowCxAgentAdvancedSettingsLoggingSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
