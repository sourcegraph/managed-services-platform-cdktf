package googlednsrecordset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlednsrecordset/internal"
)

type GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference interface {
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
	ExternalEndpoints() *[]*string
	SetExternalEndpoints(val *[]*string)
	ExternalEndpointsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalLoadBalancers() GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsInternalLoadBalancersList
	InternalLoadBalancersInput() interface{}
	InternalValue() *GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets
	SetInternalValue(val *GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets)
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
	PutInternalLoadBalancers(value interface{})
	ResetExternalEndpoints()
	ResetInternalLoadBalancers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference
type jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) ExternalEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) ExternalEndpointsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) InternalLoadBalancers() GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsInternalLoadBalancersList {
	var returns GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsInternalLoadBalancersList
	_jsii_.Get(
		j,
		"internalLoadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) InternalLoadBalancersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalLoadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) InternalValue() *GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets {
	var returns *GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDnsRecordSet.GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference_Override(g GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDnsRecordSet.GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference)SetExternalEndpoints(val *[]*string) {
	if err := j.validateSetExternalEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalEndpoints",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference)SetInternalValue(val *GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) PutInternalLoadBalancers(value interface{}) {
	if err := g.validatePutInternalLoadBalancersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInternalLoadBalancers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) ResetExternalEndpoints() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalEndpoints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) ResetInternalLoadBalancers() {
	_jsii_.InvokeVoid(
		g,
		"resetInternalLoadBalancers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

