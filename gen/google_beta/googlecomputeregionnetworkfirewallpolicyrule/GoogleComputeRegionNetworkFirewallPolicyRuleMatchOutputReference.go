package googlecomputeregionnetworkfirewallpolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputeregionnetworkfirewallpolicyrule/internal"
)

type GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference interface {
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
	InternalValue() *GoogleComputeRegionNetworkFirewallPolicyRuleMatch
	SetInternalValue(val *GoogleComputeRegionNetworkFirewallPolicyRuleMatch)
	Layer4Configs() GoogleComputeRegionNetworkFirewallPolicyRuleMatchLayer4ConfigsList
	Layer4ConfigsInput() interface{}
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
	SrcSecureTags() GoogleComputeRegionNetworkFirewallPolicyRuleMatchSrcSecureTagsList
	SrcSecureTagsInput() interface{}
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
	PutLayer4Configs(value interface{})
	PutSrcSecureTags(value interface{})
	ResetDestAddressGroups()
	ResetDestFqdns()
	ResetDestIpRanges()
	ResetDestRegionCodes()
	ResetDestThreatIntelligences()
	ResetSrcAddressGroups()
	ResetSrcFqdns()
	ResetSrcIpRanges()
	ResetSrcRegionCodes()
	ResetSrcSecureTags()
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

// The jsii proxy struct for GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference
type jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) DestAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) DestAddressGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destAddressGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) DestFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) DestFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) DestIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) DestIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) DestRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) DestRegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destRegionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) DestThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) DestThreatIntelligencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destThreatIntelligencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) InternalValue() *GoogleComputeRegionNetworkFirewallPolicyRuleMatch {
	var returns *GoogleComputeRegionNetworkFirewallPolicyRuleMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) Layer4Configs() GoogleComputeRegionNetworkFirewallPolicyRuleMatchLayer4ConfigsList {
	var returns GoogleComputeRegionNetworkFirewallPolicyRuleMatchLayer4ConfigsList
	_jsii_.Get(
		j,
		"layer4Configs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) Layer4ConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"layer4ConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcAddressGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcAddressGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcRegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcSecureTags() GoogleComputeRegionNetworkFirewallPolicyRuleMatchSrcSecureTagsList {
	var returns GoogleComputeRegionNetworkFirewallPolicyRuleMatchSrcSecureTagsList
	_jsii_.Get(
		j,
		"srcSecureTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcSecureTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srcSecureTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) SrcThreatIntelligencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcThreatIntelligencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeRegionNetworkFirewallPolicyRule.GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference_Override(g GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeRegionNetworkFirewallPolicyRule.GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetDestAddressGroups(val *[]*string) {
	if err := j.validateSetDestAddressGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destAddressGroups",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetDestFqdns(val *[]*string) {
	if err := j.validateSetDestFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destFqdns",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetDestIpRanges(val *[]*string) {
	if err := j.validateSetDestIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destIpRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetDestRegionCodes(val *[]*string) {
	if err := j.validateSetDestRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destRegionCodes",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetDestThreatIntelligences(val *[]*string) {
	if err := j.validateSetDestThreatIntelligencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destThreatIntelligences",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetInternalValue(val *GoogleComputeRegionNetworkFirewallPolicyRuleMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetSrcAddressGroups(val *[]*string) {
	if err := j.validateSetSrcAddressGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcAddressGroups",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetSrcFqdns(val *[]*string) {
	if err := j.validateSetSrcFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcFqdns",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetSrcIpRanges(val *[]*string) {
	if err := j.validateSetSrcIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcIpRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetSrcRegionCodes(val *[]*string) {
	if err := j.validateSetSrcRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcRegionCodes",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetSrcThreatIntelligences(val *[]*string) {
	if err := j.validateSetSrcThreatIntelligencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcThreatIntelligences",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) PutLayer4Configs(value interface{}) {
	if err := g.validatePutLayer4ConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLayer4Configs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) PutSrcSecureTags(value interface{}) {
	if err := g.validatePutSrcSecureTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSrcSecureTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ResetDestAddressGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetDestAddressGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ResetDestFqdns() {
	_jsii_.InvokeVoid(
		g,
		"resetDestFqdns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ResetDestIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetDestIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ResetDestRegionCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetDestRegionCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ResetDestThreatIntelligences() {
	_jsii_.InvokeVoid(
		g,
		"resetDestThreatIntelligences",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcAddressGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcAddressGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcFqdns() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcFqdns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcRegionCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcRegionCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcSecureTags() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcSecureTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcThreatIntelligences() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcThreatIntelligences",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyRuleMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

