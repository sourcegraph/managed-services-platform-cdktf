package organization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/organization/internal"
)

type OrganizationProfileOutputReference interface {
	cdktf.ComplexObject
	BusinessAddress() *string
	SetBusinessAddress(val *string)
	BusinessAddressInput() *string
	BusinessEmail() *string
	SetBusinessEmail(val *string)
	BusinessEmailInput() *string
	BusinessName() *string
	SetBusinessName(val *string)
	BusinessNameInput() *string
	BusinessPhone() *string
	SetBusinessPhone(val *string)
	BusinessPhoneInput() *string
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
	ExternalMetadata() *string
	SetExternalMetadata(val *string)
	ExternalMetadataInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrganizationProfileOutputReference
type jsiiProxy_OrganizationProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrganizationProfileOutputReference) BusinessAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) BusinessAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) BusinessEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) BusinessEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) BusinessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) BusinessNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) BusinessPhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessPhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) BusinessPhoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessPhoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) ExternalMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) ExternalMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrganizationProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrganizationProfileOutputReference {
	_init_.Initialize()

	if err := validateNewOrganizationProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrganizationProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.organization.OrganizationProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrganizationProfileOutputReference_Override(o OrganizationProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.organization.OrganizationProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrganizationProfileOutputReference)SetBusinessAddress(val *string) {
	if err := j.validateSetBusinessAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessAddress",
		val,
	)
}

func (j *jsiiProxy_OrganizationProfileOutputReference)SetBusinessEmail(val *string) {
	if err := j.validateSetBusinessEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessEmail",
		val,
	)
}

func (j *jsiiProxy_OrganizationProfileOutputReference)SetBusinessName(val *string) {
	if err := j.validateSetBusinessNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessName",
		val,
	)
}

func (j *jsiiProxy_OrganizationProfileOutputReference)SetBusinessPhone(val *string) {
	if err := j.validateSetBusinessPhoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessPhone",
		val,
	)
}

func (j *jsiiProxy_OrganizationProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrganizationProfileOutputReference)SetExternalMetadata(val *string) {
	if err := j.validateSetExternalMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalMetadata",
		val,
	)
}

func (j *jsiiProxy_OrganizationProfileOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrganizationProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrganizationProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

