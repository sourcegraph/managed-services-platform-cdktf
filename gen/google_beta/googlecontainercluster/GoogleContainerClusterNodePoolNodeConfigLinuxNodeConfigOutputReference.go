package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference interface {
	cdktf.ComplexObject
	CgroupMode() *string
	SetCgroupMode(val *string)
	CgroupModeInput() *string
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
	// Experimental.
	Fqn() *string
	HugepagesConfig() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference
	HugepagesConfigInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfig
	InternalValue() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig
	SetInternalValue(val *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig)
	Sysctls() *map[string]*string
	SetSysctls(val *map[string]*string)
	SysctlsInput() *map[string]*string
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
	PutHugepagesConfig(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfig)
	ResetCgroupMode()
	ResetHugepagesConfig()
	ResetSysctls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference
type jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) CgroupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) CgroupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) HugepagesConfig() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference
	_jsii_.Get(
		j,
		"hugepagesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) HugepagesConfigInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfig
	_jsii_.Get(
		j,
		"hugepagesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) InternalValue() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) Sysctls() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) SysctlsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleContainerCluster.GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference_Override(g GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleContainerCluster.GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference)SetCgroupMode(val *string) {
	if err := j.validateSetCgroupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cgroupMode",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference)SetInternalValue(val *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference)SetSysctls(val *map[string]*string) {
	if err := j.validateSetSysctlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sysctls",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) PutHugepagesConfig(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigHugepagesConfig) {
	if err := g.validatePutHugepagesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHugepagesConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) ResetCgroupMode() {
	_jsii_.InvokeVoid(
		g,
		"resetCgroupMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) ResetHugepagesConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHugepagesConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) ResetSysctls() {
	_jsii_.InvokeVoid(
		g,
		"resetSysctls",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

