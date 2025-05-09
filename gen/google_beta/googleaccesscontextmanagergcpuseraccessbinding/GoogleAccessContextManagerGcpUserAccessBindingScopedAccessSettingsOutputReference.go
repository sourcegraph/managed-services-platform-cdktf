package googleaccesscontextmanagergcpuseraccessbinding

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleaccesscontextmanagergcpuseraccessbinding/internal"
)

type GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference interface {
	cdktf.ComplexObject
	ActiveSettings() GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettingsOutputReference
	ActiveSettingsInput() *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings
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
	DryRunSettings() GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettingsOutputReference
	DryRunSettingsInput() *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Scope() GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsScopeOutputReference
	ScopeInput() *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope
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
	PutActiveSettings(value *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings)
	PutDryRunSettings(value *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings)
	PutScope(value *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope)
	ResetActiveSettings()
	ResetDryRunSettings()
	ResetScope()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference
type jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ActiveSettings() GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettingsOutputReference {
	var returns GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettingsOutputReference
	_jsii_.Get(
		j,
		"activeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ActiveSettingsInput() *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings {
	var returns *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings
	_jsii_.Get(
		j,
		"activeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) DryRunSettings() GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettingsOutputReference {
	var returns GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettingsOutputReference
	_jsii_.Get(
		j,
		"dryRunSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) DryRunSettingsInput() *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings {
	var returns *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings
	_jsii_.Get(
		j,
		"dryRunSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) Scope() GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsScopeOutputReference {
	var returns GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsScopeOutputReference
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ScopeInput() *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope {
	var returns *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAccessContextManagerGcpUserAccessBinding.GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference_Override(g GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAccessContextManagerGcpUserAccessBinding.GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) PutActiveSettings(value *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings) {
	if err := g.validatePutActiveSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putActiveSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) PutDryRunSettings(value *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings) {
	if err := g.validatePutDryRunSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDryRunSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) PutScope(value *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope) {
	if err := g.validatePutScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScope",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ResetActiveSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetActiveSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ResetDryRunSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetDryRunSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		g,
		"resetScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

