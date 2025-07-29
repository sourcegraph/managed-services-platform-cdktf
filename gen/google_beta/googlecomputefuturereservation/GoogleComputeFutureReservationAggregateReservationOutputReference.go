package googlecomputefuturereservation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputefuturereservation/internal"
)

type GoogleComputeFutureReservationAggregateReservationOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeFutureReservationAggregateReservation
	SetInternalValue(val *GoogleComputeFutureReservationAggregateReservation)
	ReservedResources() GoogleComputeFutureReservationAggregateReservationReservedResourcesList
	ReservedResourcesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmFamily() *string
	SetVmFamily(val *string)
	VmFamilyInput() *string
	WorkloadType() *string
	SetWorkloadType(val *string)
	WorkloadTypeInput() *string
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
	PutReservedResources(value interface{})
	ResetVmFamily()
	ResetWorkloadType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeFutureReservationAggregateReservationOutputReference
type jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) InternalValue() *GoogleComputeFutureReservationAggregateReservation {
	var returns *GoogleComputeFutureReservationAggregateReservation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) ReservedResources() GoogleComputeFutureReservationAggregateReservationReservedResourcesList {
	var returns GoogleComputeFutureReservationAggregateReservationReservedResourcesList
	_jsii_.Get(
		j,
		"reservedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) ReservedResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reservedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) VmFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) VmFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) WorkloadType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) WorkloadTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadTypeInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeFutureReservationAggregateReservationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeFutureReservationAggregateReservationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeFutureReservationAggregateReservationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeFutureReservation.GoogleComputeFutureReservationAggregateReservationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeFutureReservationAggregateReservationOutputReference_Override(g GoogleComputeFutureReservationAggregateReservationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeFutureReservation.GoogleComputeFutureReservationAggregateReservationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference)SetInternalValue(val *GoogleComputeFutureReservationAggregateReservation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference)SetVmFamily(val *string) {
	if err := j.validateSetVmFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmFamily",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference)SetWorkloadType(val *string) {
	if err := j.validateSetWorkloadTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadType",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) PutReservedResources(value interface{}) {
	if err := g.validatePutReservedResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReservedResources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) ResetVmFamily() {
	_jsii_.InvokeVoid(
		g,
		"resetVmFamily",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) ResetWorkloadType() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkloadType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeFutureReservationAggregateReservationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

