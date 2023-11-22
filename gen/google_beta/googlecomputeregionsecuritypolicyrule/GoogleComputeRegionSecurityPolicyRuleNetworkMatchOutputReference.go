package googlecomputeregionsecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputeregionsecuritypolicyrule/internal"
)

type GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference interface {
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
	DestIpRanges() *[]*string
	SetDestIpRanges(val *[]*string)
	DestIpRangesInput() *[]*string
	DestPorts() *[]*string
	SetDestPorts(val *[]*string)
	DestPortsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeRegionSecurityPolicyRuleNetworkMatch
	SetInternalValue(val *GoogleComputeRegionSecurityPolicyRuleNetworkMatch)
	IpProtocols() *[]*string
	SetIpProtocols(val *[]*string)
	IpProtocolsInput() *[]*string
	SrcAsns() *[]*float64
	SetSrcAsns(val *[]*float64)
	SrcAsnsInput() *[]*float64
	SrcIpRanges() *[]*string
	SetSrcIpRanges(val *[]*string)
	SrcIpRangesInput() *[]*string
	SrcPorts() *[]*string
	SetSrcPorts(val *[]*string)
	SrcPortsInput() *[]*string
	SrcRegionCodes() *[]*string
	SetSrcRegionCodes(val *[]*string)
	SrcRegionCodesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserDefinedFields() GoogleComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFieldsList
	UserDefinedFieldsInput() interface{}
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
	PutUserDefinedFields(value interface{})
	ResetDestIpRanges()
	ResetDestPorts()
	ResetIpProtocols()
	ResetSrcAsns()
	ResetSrcIpRanges()
	ResetSrcPorts()
	ResetSrcRegionCodes()
	ResetUserDefinedFields()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference
type jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) DestIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) DestIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) DestPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) DestPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) InternalValue() *GoogleComputeRegionSecurityPolicyRuleNetworkMatch {
	var returns *GoogleComputeRegionSecurityPolicyRuleNetworkMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) IpProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) IpProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) SrcAsns() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"srcAsns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) SrcAsnsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"srcAsnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) SrcIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) SrcIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) SrcPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) SrcPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) SrcRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) SrcRegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) UserDefinedFields() GoogleComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFieldsList {
	var returns GoogleComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFieldsList
	_jsii_.Get(
		j,
		"userDefinedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) UserDefinedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedFieldsInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeRegionSecurityPolicyRule.GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference_Override(g GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeRegionSecurityPolicyRule.GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetDestIpRanges(val *[]*string) {
	if err := j.validateSetDestIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destIpRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetDestPorts(val *[]*string) {
	if err := j.validateSetDestPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destPorts",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetInternalValue(val *GoogleComputeRegionSecurityPolicyRuleNetworkMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetIpProtocols(val *[]*string) {
	if err := j.validateSetIpProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipProtocols",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetSrcAsns(val *[]*float64) {
	if err := j.validateSetSrcAsnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcAsns",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetSrcIpRanges(val *[]*string) {
	if err := j.validateSetSrcIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcIpRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetSrcPorts(val *[]*string) {
	if err := j.validateSetSrcPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcPorts",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetSrcRegionCodes(val *[]*string) {
	if err := j.validateSetSrcRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcRegionCodes",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) PutUserDefinedFields(value interface{}) {
	if err := g.validatePutUserDefinedFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserDefinedFields",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ResetDestIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetDestIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ResetDestPorts() {
	_jsii_.InvokeVoid(
		g,
		"resetDestPorts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ResetIpProtocols() {
	_jsii_.InvokeVoid(
		g,
		"resetIpProtocols",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ResetSrcAsns() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcAsns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ResetSrcIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ResetSrcPorts() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcPorts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ResetSrcRegionCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcRegionCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ResetUserDefinedFields() {
	_jsii_.InvokeVoid(
		g,
		"resetUserDefinedFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRuleNetworkMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

