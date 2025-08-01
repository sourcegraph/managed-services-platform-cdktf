package modelarmorfloorsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/modelarmorfloorsetting/internal"
)

type ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference interface {
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
	DeidentifyTemplate() *string
	SetDeidentifyTemplate(val *string)
	DeidentifyTemplateInput() *string
	// Experimental.
	Fqn() *string
	InspectTemplate() *string
	SetInspectTemplate(val *string)
	InspectTemplateInput() *string
	InternalValue() *ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig
	SetInternalValue(val *ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig)
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
	ResetDeidentifyTemplate()
	ResetInspectTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference
type jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) DeidentifyTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deidentifyTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) DeidentifyTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deidentifyTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) InspectTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) InspectTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) InternalValue() *ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig {
	var returns *ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference {
	_init_.Initialize()

	if err := validateNewModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference_Override(m ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference)SetDeidentifyTemplate(val *string) {
	if err := j.validateSetDeidentifyTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deidentifyTemplate",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference)SetInspectTemplate(val *string) {
	if err := j.validateSetInspectTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectTemplate",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference)SetInternalValue(val *ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) ResetDeidentifyTemplate() {
	_jsii_.InvokeVoid(
		m,
		"resetDeidentifyTemplate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) ResetInspectTemplate() {
	_jsii_.InvokeVoid(
		m,
		"resetInspectTemplate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigSdpSettingsAdvancedConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

