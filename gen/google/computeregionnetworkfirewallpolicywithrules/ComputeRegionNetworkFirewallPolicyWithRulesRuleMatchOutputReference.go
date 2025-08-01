package computeregionnetworkfirewallpolicywithrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/computeregionnetworkfirewallpolicywithrules/internal"
)

type ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference interface {
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
	DestAddressGroups() *[]*string
	SetDestAddressGroups(val *[]*string)
	DestAddressGroupsInput() *[]*string
	DestFqdns() *[]*string
	SetDestFqdns(val *[]*string)
	DestFqdnsInput() *[]*string
	DestIpRanges() *[]*string
	SetDestIpRanges(val *[]*string)
	DestIpRangesInput() *[]*string
	DestRegionCodes() *[]*string
	SetDestRegionCodes(val *[]*string)
	DestRegionCodesInput() *[]*string
	DestThreatIntelligences() *[]*string
	SetDestThreatIntelligences(val *[]*string)
	DestThreatIntelligencesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeRegionNetworkFirewallPolicyWithRulesRuleMatch
	SetInternalValue(val *ComputeRegionNetworkFirewallPolicyWithRulesRuleMatch)
	Layer4Config() ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchLayer4ConfigList
	Layer4ConfigInput() interface{}
	SrcAddressGroups() *[]*string
	SetSrcAddressGroups(val *[]*string)
	SrcAddressGroupsInput() *[]*string
	SrcFqdns() *[]*string
	SetSrcFqdns(val *[]*string)
	SrcFqdnsInput() *[]*string
	SrcIpRanges() *[]*string
	SetSrcIpRanges(val *[]*string)
	SrcIpRangesInput() *[]*string
	SrcRegionCodes() *[]*string
	SetSrcRegionCodes(val *[]*string)
	SrcRegionCodesInput() *[]*string
	SrcSecureTag() ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList
	SrcSecureTagInput() interface{}
	SrcThreatIntelligences() *[]*string
	SetSrcThreatIntelligences(val *[]*string)
	SrcThreatIntelligencesInput() *[]*string
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
	PutLayer4Config(value interface{})
	PutSrcSecureTag(value interface{})
	ResetDestAddressGroups()
	ResetDestFqdns()
	ResetDestIpRanges()
	ResetDestRegionCodes()
	ResetDestThreatIntelligences()
	ResetSrcAddressGroups()
	ResetSrcFqdns()
	ResetSrcIpRanges()
	ResetSrcRegionCodes()
	ResetSrcSecureTag()
	ResetSrcThreatIntelligences()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference
type jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestAddressGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destAddressGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestRegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destRegionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestThreatIntelligencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destThreatIntelligencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) InternalValue() *ComputeRegionNetworkFirewallPolicyWithRulesRuleMatch {
	var returns *ComputeRegionNetworkFirewallPolicyWithRulesRuleMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) Layer4Config() ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchLayer4ConfigList {
	var returns ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchLayer4ConfigList
	_jsii_.Get(
		j,
		"layer4Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) Layer4ConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"layer4ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcAddressGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcAddressGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcRegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcSecureTag() ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList {
	var returns ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList
	_jsii_.Get(
		j,
		"srcSecureTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcSecureTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srcSecureTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcThreatIntelligencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcThreatIntelligencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionNetworkFirewallPolicyWithRules.ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference_Override(c ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionNetworkFirewallPolicyWithRules.ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetDestAddressGroups(val *[]*string) {
	if err := j.validateSetDestAddressGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destAddressGroups",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetDestFqdns(val *[]*string) {
	if err := j.validateSetDestFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destFqdns",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetDestIpRanges(val *[]*string) {
	if err := j.validateSetDestIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destIpRanges",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetDestRegionCodes(val *[]*string) {
	if err := j.validateSetDestRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destRegionCodes",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetDestThreatIntelligences(val *[]*string) {
	if err := j.validateSetDestThreatIntelligencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destThreatIntelligences",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetInternalValue(val *ComputeRegionNetworkFirewallPolicyWithRulesRuleMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcAddressGroups(val *[]*string) {
	if err := j.validateSetSrcAddressGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcAddressGroups",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcFqdns(val *[]*string) {
	if err := j.validateSetSrcFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcFqdns",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcIpRanges(val *[]*string) {
	if err := j.validateSetSrcIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcIpRanges",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcRegionCodes(val *[]*string) {
	if err := j.validateSetSrcRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcRegionCodes",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcThreatIntelligences(val *[]*string) {
	if err := j.validateSetSrcThreatIntelligencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcThreatIntelligences",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) PutLayer4Config(value interface{}) {
	if err := c.validatePutLayer4ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLayer4Config",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) PutSrcSecureTag(value interface{}) {
	if err := c.validatePutSrcSecureTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSrcSecureTag",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetDestAddressGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetDestAddressGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetDestFqdns() {
	_jsii_.InvokeVoid(
		c,
		"resetDestFqdns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetDestIpRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetDestIpRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetDestRegionCodes() {
	_jsii_.InvokeVoid(
		c,
		"resetDestRegionCodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetDestThreatIntelligences() {
	_jsii_.InvokeVoid(
		c,
		"resetDestThreatIntelligences",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcAddressGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcAddressGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcFqdns() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcFqdns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcIpRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcIpRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcRegionCodes() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcRegionCodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcSecureTag() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcSecureTag",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcThreatIntelligences() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcThreatIntelligences",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

