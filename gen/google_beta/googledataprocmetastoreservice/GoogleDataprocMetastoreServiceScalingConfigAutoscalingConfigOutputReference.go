package googledataprocmetastoreservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledataprocmetastoreservice/internal"
)

type GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference interface {
	cdktf.ComplexObject
	AutoscalingEnabled() any
	SetAutoscalingEnabled(val any)
	AutoscalingEnabledInput() any
	AutoscalingFactor() *float64
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
	InternalValue() *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfig
	SetInternalValue(val *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfig)
	LimitConfig() GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference
	LimitConfigInput() *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig
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
	PutLimitConfig(value *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig)
	ResetAutoscalingEnabled()
	ResetLimitConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference
type jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) AutoscalingEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"autoscalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) AutoscalingEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"autoscalingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) AutoscalingFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) InternalValue() *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfig {
	var returns *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) LimitConfig() GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference {
	var returns GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference
	_jsii_.Get(
		j,
		"limitConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) LimitConfigInput() *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig {
	var returns *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig
	_jsii_.Get(
		j,
		"limitConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDataprocMetastoreService.GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference_Override(g GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDataprocMetastoreService.GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) SetAutoscalingEnabled(val any) {
	if err := j.validateSetAutoscalingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) SetInternalValue(val *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) PutLimitConfig(value *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig) {
	if err := g.validatePutLimitConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLimitConfig",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ResetAutoscalingEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoscalingEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ResetLimitConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLimitConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
