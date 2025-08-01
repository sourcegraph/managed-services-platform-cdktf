package computeregionsecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/computeregionsecuritypolicy/internal"
)

type ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference interface {
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
	InternalValue() *ComputeRegionSecurityPolicyRulesNetworkMatch
	SetInternalValue(val *ComputeRegionSecurityPolicyRulesNetworkMatch)
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
	UserDefinedFields() ComputeRegionSecurityPolicyRulesNetworkMatchUserDefinedFieldsList
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

// The jsii proxy struct for ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference
type jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) DestIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) DestIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) DestPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) DestPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) InternalValue() *ComputeRegionSecurityPolicyRulesNetworkMatch {
	var returns *ComputeRegionSecurityPolicyRulesNetworkMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) IpProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) IpProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) SrcAsns() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"srcAsns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) SrcAsnsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"srcAsnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) SrcIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) SrcIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) SrcPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) SrcPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) SrcRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) SrcRegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) UserDefinedFields() ComputeRegionSecurityPolicyRulesNetworkMatchUserDefinedFieldsList {
	var returns ComputeRegionSecurityPolicyRulesNetworkMatchUserDefinedFieldsList
	_jsii_.Get(
		j,
		"userDefinedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) UserDefinedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedFieldsInput",
		&returns,
	)
	return returns
}


func NewComputeRegionSecurityPolicyRulesNetworkMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionSecurityPolicyRulesNetworkMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionSecurityPolicy.ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionSecurityPolicyRulesNetworkMatchOutputReference_Override(c ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionSecurityPolicy.ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetDestIpRanges(val *[]*string) {
	if err := j.validateSetDestIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destIpRanges",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetDestPorts(val *[]*string) {
	if err := j.validateSetDestPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destPorts",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetInternalValue(val *ComputeRegionSecurityPolicyRulesNetworkMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetIpProtocols(val *[]*string) {
	if err := j.validateSetIpProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipProtocols",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetSrcAsns(val *[]*float64) {
	if err := j.validateSetSrcAsnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcAsns",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetSrcIpRanges(val *[]*string) {
	if err := j.validateSetSrcIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcIpRanges",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetSrcPorts(val *[]*string) {
	if err := j.validateSetSrcPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcPorts",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetSrcRegionCodes(val *[]*string) {
	if err := j.validateSetSrcRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcRegionCodes",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) PutUserDefinedFields(value interface{}) {
	if err := c.validatePutUserDefinedFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUserDefinedFields",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ResetDestIpRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetDestIpRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ResetDestPorts() {
	_jsii_.InvokeVoid(
		c,
		"resetDestPorts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ResetIpProtocols() {
	_jsii_.InvokeVoid(
		c,
		"resetIpProtocols",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ResetSrcAsns() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcAsns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ResetSrcIpRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcIpRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ResetSrcPorts() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcPorts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ResetSrcRegionCodes() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcRegionCodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ResetUserDefinedFields() {
	_jsii_.InvokeVoid(
		c,
		"resetUserDefinedFields",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

