package googlemodelarmortemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlemodelarmortemplate/internal"
)

type GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference interface {
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
	ConfidenceLevel() *string
	SetConfidenceLevel(val *string)
	ConfidenceLevelInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FilterEnforcement() *string
	SetFilterEnforcement(val *string)
	FilterEnforcementInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings
	SetInternalValue(val *GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings)
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
	ResetConfidenceLevel()
	ResetFilterEnforcement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference
type jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) ConfidenceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidenceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) ConfidenceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidenceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) FilterEnforcement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterEnforcement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) FilterEnforcementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterEnforcementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) InternalValue() *GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings {
	var returns *GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleModelArmorTemplate.GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference_Override(g GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleModelArmorTemplate.GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference)SetConfidenceLevel(val *string) {
	if err := j.validateSetConfidenceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidenceLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference)SetFilterEnforcement(val *string) {
	if err := j.validateSetFilterEnforcementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterEnforcement",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference)SetInternalValue(val *GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) ResetConfidenceLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidenceLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) ResetFilterEnforcement() {
	_jsii_.InvokeVoid(
		g,
		"resetFilterEnforcement",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

