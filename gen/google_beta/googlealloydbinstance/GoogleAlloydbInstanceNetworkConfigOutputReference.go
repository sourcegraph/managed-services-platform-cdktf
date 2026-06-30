package googlealloydbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlealloydbinstance/internal"
)

type GoogleAlloydbInstanceNetworkConfigOutputReference interface {
	cdktf.ComplexObject
	AllocatedIpRangeOverride() *string
	SetAllocatedIpRangeOverride(val *string)
	AllocatedIpRangeOverrideInput() *string
	AuthorizedExternalNetworks() GoogleAlloydbInstanceNetworkConfigAuthorizedExternalNetworksList
	AuthorizedExternalNetworksInput() any
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
	EnableOutboundPublicIp() any
	SetEnableOutboundPublicIp(val any)
	EnableOutboundPublicIpInput() any
	EnablePublicIp() any
	SetEnablePublicIp(val any)
	EnablePublicIpInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleAlloydbInstanceNetworkConfig
	SetInternalValue(val *GoogleAlloydbInstanceNetworkConfig)
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
	PutAuthorizedExternalNetworks(value any)
	ResetAllocatedIpRangeOverride()
	ResetAuthorizedExternalNetworks()
	ResetEnableOutboundPublicIp()
	ResetEnablePublicIp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAlloydbInstanceNetworkConfigOutputReference
type jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) AllocatedIpRangeOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedIpRangeOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) AllocatedIpRangeOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedIpRangeOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) AuthorizedExternalNetworks() GoogleAlloydbInstanceNetworkConfigAuthorizedExternalNetworksList {
	var returns GoogleAlloydbInstanceNetworkConfigAuthorizedExternalNetworksList
	_jsii_.Get(
		j,
		"authorizedExternalNetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) AuthorizedExternalNetworksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"authorizedExternalNetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) EnableOutboundPublicIp() any {
	var returns any
	_jsii_.Get(
		j,
		"enableOutboundPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) EnableOutboundPublicIpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableOutboundPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) EnablePublicIp() any {
	var returns any
	_jsii_.Get(
		j,
		"enablePublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) EnablePublicIpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enablePublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) InternalValue() *GoogleAlloydbInstanceNetworkConfig {
	var returns *GoogleAlloydbInstanceNetworkConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleAlloydbInstanceNetworkConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAlloydbInstanceNetworkConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAlloydbInstanceNetworkConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAlloydbInstance.GoogleAlloydbInstanceNetworkConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAlloydbInstanceNetworkConfigOutputReference_Override(g GoogleAlloydbInstanceNetworkConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAlloydbInstance.GoogleAlloydbInstanceNetworkConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) SetAllocatedIpRangeOverride(val *string) {
	if err := j.validateSetAllocatedIpRangeOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedIpRangeOverride",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) SetEnableOutboundPublicIp(val any) {
	if err := j.validateSetEnableOutboundPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOutboundPublicIp",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) SetEnablePublicIp(val any) {
	if err := j.validateSetEnablePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePublicIp",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) SetInternalValue(val *GoogleAlloydbInstanceNetworkConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) PutAuthorizedExternalNetworks(value any) {
	if err := g.validatePutAuthorizedExternalNetworksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorizedExternalNetworks",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) ResetAllocatedIpRangeOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetAllocatedIpRangeOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) ResetAuthorizedExternalNetworks() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizedExternalNetworks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) ResetEnableOutboundPublicIp() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableOutboundPublicIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) ResetEnablePublicIp() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePublicIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleAlloydbInstanceNetworkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
