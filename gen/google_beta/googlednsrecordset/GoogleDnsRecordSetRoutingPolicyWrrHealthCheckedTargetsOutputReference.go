package googlednsrecordset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlednsrecordset/internal"
)

type GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExternalEndpoints() *[]*string
	SetExternalEndpoints(val *[]*string)
	ExternalEndpointsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalLoadBalancers() GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsInternalLoadBalancersList
	InternalLoadBalancersInput() any
	InternalValue() *GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargets
	SetInternalValue(val *GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargets)
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
	PutInternalLoadBalancers(value any)
	ResetExternalEndpoints()
	ResetInternalLoadBalancers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference
type jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) ExternalEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) ExternalEndpointsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) InternalLoadBalancers() GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsInternalLoadBalancersList {
	var returns GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsInternalLoadBalancersList
	_jsii_.Get(
		j,
		"internalLoadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) InternalLoadBalancersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"internalLoadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) InternalValue() *GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargets {
	var returns *GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDnsRecordSet.GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference_Override(g GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDnsRecordSet.GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) SetExternalEndpoints(val *[]*string) {
	if err := j.validateSetExternalEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalEndpoints",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) SetInternalValue(val *GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) PutInternalLoadBalancers(value any) {
	if err := g.validatePutInternalLoadBalancersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInternalLoadBalancers",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) ResetExternalEndpoints() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalEndpoints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) ResetInternalLoadBalancers() {
	_jsii_.InvokeVoid(
		g,
		"resetInternalLoadBalancers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
