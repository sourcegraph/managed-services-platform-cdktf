package computenetworkfirewallpolicywithrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/computenetworkfirewallpolicywithrules/internal"
)

type ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference interface {
	cdktf.ComplexObject
	Action() *string
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
	Description() *string
	Direction() *string
	Disabled() cdktf.IResolvable
	EnableLogging() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeNetworkFirewallPolicyWithRulesPredefinedRules
	SetInternalValue(val *ComputeNetworkFirewallPolicyWithRulesPredefinedRules)
	Match() ComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchList
	Priority() *float64
	RuleName() *string
	SecurityProfileGroup() *string
	TargetSecureTag() ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList
	TargetServiceAccounts() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsInspect() cdktf.IResolvable
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference
type jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) Disabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) EnableLogging() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) InternalValue() *ComputeNetworkFirewallPolicyWithRulesPredefinedRules {
	var returns *ComputeNetworkFirewallPolicyWithRulesPredefinedRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) Match() ComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchList {
	var returns ComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchList
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) SecurityProfileGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) TargetSecureTag() ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList {
	var returns ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList
	_jsii_.Get(
		j,
		"targetSecureTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) TargetServiceAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) TlsInspect() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"tlsInspect",
		&returns,
	)
	return returns
}


func NewComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference {
	_init_.Initialize()

	if err := validateNewComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeNetworkFirewallPolicyWithRules.ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference_Override(c ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeNetworkFirewallPolicyWithRules.ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference)SetInternalValue(val *ComputeNetworkFirewallPolicyWithRulesPredefinedRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

