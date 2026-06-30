package googlesqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlesqldatabaseinstance/internal"
)

type GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference interface {
	cdktf.ComplexObject
	AllowedConsumerProjects() *[]*string
	SetAllowedConsumerProjects(val *[]*string)
	AllowedConsumerProjectsInput() *[]*string
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
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	PscAutoConnections() GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsList
	PscAutoConnectionsInput() any
	PscEnabled() any
	SetPscEnabled(val any)
	PscEnabledInput() any
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
	PutPscAutoConnections(value any)
	ResetAllowedConsumerProjects()
	ResetPscAutoConnections()
	ResetPscEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference
type jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) AllowedConsumerProjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedConsumerProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) AllowedConsumerProjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedConsumerProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) PscAutoConnections() GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsList {
	var returns GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsList
	_jsii_.Get(
		j,
		"pscAutoConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) PscAutoConnectionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"pscAutoConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) PscEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"pscEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) PscEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"pscEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference_Override(g GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) SetAllowedConsumerProjects(val *[]*string) {
	if err := j.validateSetAllowedConsumerProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedConsumerProjects",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) SetPscEnabled(val any) {
	if err := j.validateSetPscEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pscEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) PutPscAutoConnections(value any) {
	if err := g.validatePutPscAutoConnectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPscAutoConnections",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) ResetAllowedConsumerProjects() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedConsumerProjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) ResetPscAutoConnections() {
	_jsii_.InvokeVoid(
		g,
		"resetPscAutoConnections",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) ResetPscEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetPscEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
