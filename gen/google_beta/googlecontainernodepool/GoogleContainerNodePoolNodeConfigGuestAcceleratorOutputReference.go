package googlecontainernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecontainernodepool/internal"
)

type GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference interface {
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GpuDriverInstallationConfig() GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfigOutputReference
	GpuDriverInstallationConfigInput() *GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfig
	GpuPartitionSize() *string
	SetGpuPartitionSize(val *string)
	GpuPartitionSizeInput() *string
	GpuSharingConfig() GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference
	GpuSharingConfigInput() *GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuSharingConfig
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutGpuDriverInstallationConfig(value *GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfig)
	PutGpuSharingConfig(value *GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuSharingConfig)
	ResetGpuDriverInstallationConfig()
	ResetGpuPartitionSize()
	ResetGpuSharingConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference
type jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GpuDriverInstallationConfig() GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfigOutputReference
	_jsii_.Get(
		j,
		"gpuDriverInstallationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GpuDriverInstallationConfigInput() *GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfig {
	var returns *GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfig
	_jsii_.Get(
		j,
		"gpuDriverInstallationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GpuPartitionSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuPartitionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GpuPartitionSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuPartitionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GpuSharingConfig() GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference
	_jsii_.Get(
		j,
		"gpuSharingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GpuSharingConfigInput() *GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuSharingConfig {
	var returns *GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuSharingConfig
	_jsii_.Get(
		j,
		"gpuSharingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference_Override(g GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference)SetCount(val *float64) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference)SetGpuPartitionSize(val *string) {
	if err := j.validateSetGpuPartitionSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuPartitionSize",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) PutGpuDriverInstallationConfig(value *GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfig) {
	if err := g.validatePutGpuDriverInstallationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGpuDriverInstallationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) PutGpuSharingConfig(value *GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuSharingConfig) {
	if err := g.validatePutGpuSharingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGpuSharingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) ResetGpuDriverInstallationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGpuDriverInstallationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) ResetGpuPartitionSize() {
	_jsii_.InvokeVoid(
		g,
		"resetGpuPartitionSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) ResetGpuSharingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGpuSharingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigGuestAcceleratorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

