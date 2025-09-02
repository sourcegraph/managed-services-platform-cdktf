package googlemodelarmorfloorsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlemodelarmorfloorsetting/internal"
)

type GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference interface {
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
	EnableCloudLogging() interface{}
	SetEnableCloudLogging(val interface{})
	EnableCloudLoggingInput() interface{}
	// Experimental.
	Fqn() *string
	InspectAndBlock() interface{}
	SetInspectAndBlock(val interface{})
	InspectAndBlockInput() interface{}
	InspectOnly() interface{}
	SetInspectOnly(val interface{})
	InspectOnlyInput() interface{}
	InternalValue() *GoogleModelArmorFloorsettingAiPlatformFloorSetting
	SetInternalValue(val *GoogleModelArmorFloorsettingAiPlatformFloorSetting)
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
	ResetEnableCloudLogging()
	ResetInspectAndBlock()
	ResetInspectOnly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference
type jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) EnableCloudLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCloudLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) EnableCloudLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCloudLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InspectAndBlock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectAndBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InspectAndBlockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectAndBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InspectOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InspectOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InternalValue() *GoogleModelArmorFloorsettingAiPlatformFloorSetting {
	var returns *GoogleModelArmorFloorsettingAiPlatformFloorSetting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleModelArmorFloorsetting.GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference_Override(g GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleModelArmorFloorsetting.GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetEnableCloudLogging(val interface{}) {
	if err := j.validateSetEnableCloudLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCloudLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetInspectAndBlock(val interface{}) {
	if err := j.validateSetInspectAndBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectAndBlock",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetInspectOnly(val interface{}) {
	if err := j.validateSetInspectOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectOnly",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetInternalValue(val *GoogleModelArmorFloorsettingAiPlatformFloorSetting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ResetEnableCloudLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableCloudLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ResetInspectAndBlock() {
	_jsii_.InvokeVoid(
		g,
		"resetInspectAndBlock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ResetInspectOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetInspectOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

