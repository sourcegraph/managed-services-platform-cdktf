package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference interface {
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
	ContainerdConfig() GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigOutputReference
	ContainerdConfigInput() *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GcfsConfig() GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfigOutputReference
	GcfsConfigInput() *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig
	InsecureKubeletReadonlyPortEnabled() *string
	SetInsecureKubeletReadonlyPortEnabled(val *string)
	InsecureKubeletReadonlyPortEnabledInput() *string
	InternalValue() *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaults
	SetInternalValue(val *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaults)
	LoggingVariant() *string
	SetLoggingVariant(val *string)
	LoggingVariantInput() *string
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
	PutContainerdConfig(value *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig)
	PutGcfsConfig(value *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig)
	ResetContainerdConfig()
	ResetGcfsConfig()
	ResetInsecureKubeletReadonlyPortEnabled()
	ResetLoggingVariant()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference
type jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ContainerdConfig() GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigOutputReference {
	var returns GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigOutputReference
	_jsii_.Get(
		j,
		"containerdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ContainerdConfigInput() *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig {
	var returns *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig
	_jsii_.Get(
		j,
		"containerdConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GcfsConfig() GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfigOutputReference {
	var returns GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfigOutputReference
	_jsii_.Get(
		j,
		"gcfsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GcfsConfigInput() *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig {
	var returns *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig
	_jsii_.Get(
		j,
		"gcfsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) InsecureKubeletReadonlyPortEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insecureKubeletReadonlyPortEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) InsecureKubeletReadonlyPortEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insecureKubeletReadonlyPortEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) InternalValue() *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaults {
	var returns *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaults
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) LoggingVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) LoggingVariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleContainerCluster.GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference_Override(g GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleContainerCluster.GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) SetInsecureKubeletReadonlyPortEnabled(val *string) {
	if err := j.validateSetInsecureKubeletReadonlyPortEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureKubeletReadonlyPortEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) SetInternalValue(val *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaults) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) SetLoggingVariant(val *string) {
	if err := j.validateSetLoggingVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingVariant",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) PutContainerdConfig(value *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig) {
	if err := g.validatePutContainerdConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainerdConfig",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) PutGcfsConfig(value *GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig) {
	if err := g.validatePutGcfsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcfsConfig",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ResetContainerdConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerdConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ResetGcfsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGcfsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ResetInsecureKubeletReadonlyPortEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetInsecureKubeletReadonlyPortEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ResetLoggingVariant() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingVariant",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
