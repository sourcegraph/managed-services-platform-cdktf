package googlenetappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlenetappvolume/internal"
)

type GoogleNetappVolumeTieringPolicyOutputReference interface {
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
	CoolingThresholdDays() *float64
	SetCoolingThresholdDays(val *float64)
	CoolingThresholdDaysInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HotTierBypassModeEnabled() interface{}
	SetHotTierBypassModeEnabled(val interface{})
	HotTierBypassModeEnabledInput() interface{}
	InternalValue() *GoogleNetappVolumeTieringPolicy
	SetInternalValue(val *GoogleNetappVolumeTieringPolicy)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TierAction() *string
	SetTierAction(val *string)
	TierActionInput() *string
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
	ResetCoolingThresholdDays()
	ResetHotTierBypassModeEnabled()
	ResetTierAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetappVolumeTieringPolicyOutputReference
type jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) CoolingThresholdDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coolingThresholdDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) CoolingThresholdDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coolingThresholdDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) HotTierBypassModeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hotTierBypassModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) HotTierBypassModeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hotTierBypassModeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) InternalValue() *GoogleNetappVolumeTieringPolicy {
	var returns *GoogleNetappVolumeTieringPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) TierAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) TierActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierActionInput",
		&returns,
	)
	return returns
}


func NewGoogleNetappVolumeTieringPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetappVolumeTieringPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetappVolumeTieringPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleNetappVolume.GoogleNetappVolumeTieringPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetappVolumeTieringPolicyOutputReference_Override(g GoogleNetappVolumeTieringPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleNetappVolume.GoogleNetappVolumeTieringPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference)SetCoolingThresholdDays(val *float64) {
	if err := j.validateSetCoolingThresholdDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coolingThresholdDays",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference)SetHotTierBypassModeEnabled(val interface{}) {
	if err := j.validateSetHotTierBypassModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hotTierBypassModeEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference)SetInternalValue(val *GoogleNetappVolumeTieringPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference)SetTierAction(val *string) {
	if err := j.validateSetTierActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tierAction",
		val,
	)
}

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) ResetCoolingThresholdDays() {
	_jsii_.InvokeVoid(
		g,
		"resetCoolingThresholdDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) ResetHotTierBypassModeEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetHotTierBypassModeEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) ResetTierAction() {
	_jsii_.InvokeVoid(
		g,
		"resetTierAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetappVolumeTieringPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

