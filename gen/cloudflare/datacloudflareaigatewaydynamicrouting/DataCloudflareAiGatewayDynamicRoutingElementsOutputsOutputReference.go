package datacloudflareaigatewaydynamicrouting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/datacloudflareaigatewaydynamicrouting/internal"
)

type DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference interface {
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
	ElementId() *string
	Fallback() DataCloudflareAiGatewayDynamicRoutingElementsOutputsFallbackOutputReference
	False() DataCloudflareAiGatewayDynamicRoutingElementsOutputsFalseOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataCloudflareAiGatewayDynamicRoutingElementsOutputs
	SetInternalValue(val *DataCloudflareAiGatewayDynamicRoutingElementsOutputs)
	Next() DataCloudflareAiGatewayDynamicRoutingElementsOutputsNextOutputReference
	Success() DataCloudflareAiGatewayDynamicRoutingElementsOutputsSuccessOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	True() DataCloudflareAiGatewayDynamicRoutingElementsOutputsTrueOutputReference
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

// The jsii proxy struct for DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference
type jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) ElementId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) Fallback() DataCloudflareAiGatewayDynamicRoutingElementsOutputsFallbackOutputReference {
	var returns DataCloudflareAiGatewayDynamicRoutingElementsOutputsFallbackOutputReference
	_jsii_.Get(
		j,
		"fallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) False() DataCloudflareAiGatewayDynamicRoutingElementsOutputsFalseOutputReference {
	var returns DataCloudflareAiGatewayDynamicRoutingElementsOutputsFalseOutputReference
	_jsii_.Get(
		j,
		"false",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) InternalValue() *DataCloudflareAiGatewayDynamicRoutingElementsOutputs {
	var returns *DataCloudflareAiGatewayDynamicRoutingElementsOutputs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) Next() DataCloudflareAiGatewayDynamicRoutingElementsOutputsNextOutputReference {
	var returns DataCloudflareAiGatewayDynamicRoutingElementsOutputsNextOutputReference
	_jsii_.Get(
		j,
		"next",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) Success() DataCloudflareAiGatewayDynamicRoutingElementsOutputsSuccessOutputReference {
	var returns DataCloudflareAiGatewayDynamicRoutingElementsOutputsSuccessOutputReference
	_jsii_.Get(
		j,
		"success",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) True() DataCloudflareAiGatewayDynamicRoutingElementsOutputsTrueOutputReference {
	var returns DataCloudflareAiGatewayDynamicRoutingElementsOutputsTrueOutputReference
	_jsii_.Get(
		j,
		"true",
		&returns,
	)
	return returns
}


func NewDataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareAiGatewayDynamicRouting.DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference_Override(d DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareAiGatewayDynamicRouting.DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference)SetInternalValue(val *DataCloudflareAiGatewayDynamicRoutingElementsOutputs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareAiGatewayDynamicRoutingElementsOutputsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

