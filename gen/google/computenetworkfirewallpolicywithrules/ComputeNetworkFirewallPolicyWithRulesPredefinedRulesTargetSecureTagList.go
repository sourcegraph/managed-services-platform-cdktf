package computenetworkfirewallpolicywithrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/computenetworkfirewallpolicywithrules/internal"
)

type ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList
type jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList {
	_init_.Initialize()

	if err := validateNewComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList{}

	_jsii_.Create(
		"@cdktf/provider-google.computeNetworkFirewallPolicyWithRules.ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList_Override(c ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeNetworkFirewallPolicyWithRules.ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := c.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		c,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList) Get(index *float64) ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesPredefinedRulesTargetSecureTagList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

