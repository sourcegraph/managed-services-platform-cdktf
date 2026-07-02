package magictransitsitelan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/magictransitsitelan/internal"
)

type MagicTransitSiteLanRoutedSubnetsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	Nat() MagicTransitSiteLanRoutedSubnetsNatOutputReference
	NatInput() any
	NextHop() *string
	SetNextHop(val *string)
	NextHopInput() *string
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
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
	PutNat(value *MagicTransitSiteLanRoutedSubnetsNat)
	ResetNat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MagicTransitSiteLanRoutedSubnetsOutputReference
type jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) Nat() MagicTransitSiteLanRoutedSubnetsNatOutputReference {
	var returns MagicTransitSiteLanRoutedSubnetsNatOutputReference
	_jsii_.Get(
		j,
		"nat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) NatInput() any {
	var returns any
	_jsii_.Get(
		j,
		"natInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) NextHop() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextHop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) NextHopInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextHopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMagicTransitSiteLanRoutedSubnetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MagicTransitSiteLanRoutedSubnetsOutputReference {
	_init_.Initialize()

	if err := validateNewMagicTransitSiteLanRoutedSubnetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicTransitSiteLan.MagicTransitSiteLanRoutedSubnetsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMagicTransitSiteLanRoutedSubnetsOutputReference_Override(m MagicTransitSiteLanRoutedSubnetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.magicTransitSiteLan.MagicTransitSiteLanRoutedSubnetsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) SetNextHop(val *string) {
	if err := j.validateSetNextHopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextHop",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) PutNat(value *MagicTransitSiteLanRoutedSubnetsNat) {
	if err := m.validatePutNatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNat",
		[]any{value},
	)
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) ResetNat() {
	_jsii_.InvokeVoid(
		m,
		"resetNat",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		m,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MagicTransitSiteLanRoutedSubnetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
