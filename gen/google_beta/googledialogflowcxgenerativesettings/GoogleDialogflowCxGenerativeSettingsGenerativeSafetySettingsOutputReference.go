package googledialogflowcxgenerativesettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledialogflowcxgenerativesettings/internal"
)

type GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference interface {
	cdktf.ComplexObject
	BannedPhrases() GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsBannedPhrasesList
	BannedPhrasesInput() interface{}
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
	DefaultBannedPhraseMatchStrategy() *string
	SetDefaultBannedPhraseMatchStrategy(val *string)
	DefaultBannedPhraseMatchStrategyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettings
	SetInternalValue(val *GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettings)
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
	PutBannedPhrases(value interface{})
	ResetBannedPhrases()
	ResetDefaultBannedPhraseMatchStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference
type jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) BannedPhrases() GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsBannedPhrasesList {
	var returns GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsBannedPhrasesList
	_jsii_.Get(
		j,
		"bannedPhrases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) BannedPhrasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bannedPhrasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) DefaultBannedPhraseMatchStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBannedPhraseMatchStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) DefaultBannedPhraseMatchStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBannedPhraseMatchStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) InternalValue() *GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettings {
	var returns *GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxGenerativeSettings.GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference_Override(g GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxGenerativeSettings.GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetDefaultBannedPhraseMatchStrategy(val *string) {
	if err := j.validateSetDefaultBannedPhraseMatchStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBannedPhraseMatchStrategy",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetInternalValue(val *GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) PutBannedPhrases(value interface{}) {
	if err := g.validatePutBannedPhrasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBannedPhrases",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ResetBannedPhrases() {
	_jsii_.InvokeVoid(
		g,
		"resetBannedPhrases",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ResetDefaultBannedPhraseMatchStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultBannedPhraseMatchStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

