package datacloudflarezerotrusttunnelcloudflaredvirtualnetwork

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/datacloudflarezerotrusttunnelcloudflaredvirtualnetwork/internal"
)

type DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference interface {
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsDefault() interface{}
	SetIsDefault(val interface{})
	IsDefaultInput() interface{}
	IsDeleted() interface{}
	SetIsDeleted(val interface{})
	IsDeletedInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ResetId()
	ResetIsDefault()
	ResetIsDeleted()
	ResetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference
type jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) IsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) IsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) IsDeleted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDeleted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) IsDeletedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDeletedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredVirtualNetwork.DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference_Override(d DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustTunnelCloudflaredVirtualNetwork.DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference)SetIsDefault(val interface{}) {
	if err := j.validateSetIsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDefault",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference)SetIsDeleted(val interface{}) {
	if err := j.validateSetIsDeletedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDeleted",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) ResetIsDefault() {
	_jsii_.InvokeVoid(
		d,
		"resetIsDefault",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) ResetIsDeleted() {
	_jsii_.InvokeVoid(
		d,
		"resetIsDeleted",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

