package zerotrustgatewaypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/zerotrustgatewaypolicy/internal"
)

type ZeroTrustGatewayPolicyScheduleOutputReference interface {
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
	Fri() *string
	SetFri(val *string)
	FriInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Mon() *string
	SetMon(val *string)
	MonInput() *string
	Sat() *string
	SetSat(val *string)
	SatInput() *string
	Sun() *string
	SetSun(val *string)
	SunInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thu() *string
	SetThu(val *string)
	ThuInput() *string
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	Tue() *string
	SetTue(val *string)
	TueInput() *string
	Wed() *string
	SetWed(val *string)
	WedInput() *string
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
	ResetFri()
	ResetMon()
	ResetSat()
	ResetSun()
	ResetThu()
	ResetTimeZone()
	ResetTue()
	ResetWed()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustGatewayPolicyScheduleOutputReference
type jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) Fri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) FriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) Mon() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) MonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) Sat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) SatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"satInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) Sun() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) SunInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) Thu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ThuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) Tue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) TueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) Wed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) WedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wedInput",
		&returns,
	)
	return returns
}


func NewZeroTrustGatewayPolicyScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustGatewayPolicyScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustGatewayPolicyScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewayPolicy.ZeroTrustGatewayPolicyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustGatewayPolicyScheduleOutputReference_Override(z ZeroTrustGatewayPolicyScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustGatewayPolicy.ZeroTrustGatewayPolicyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetFri(val *string) {
	if err := j.validateSetFriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fri",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetMon(val *string) {
	if err := j.validateSetMonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mon",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetSat(val *string) {
	if err := j.validateSetSatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sat",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetSun(val *string) {
	if err := j.validateSetSunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sun",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetThu(val *string) {
	if err := j.validateSetThuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thu",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetTue(val *string) {
	if err := j.validateSetTueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference)SetWed(val *string) {
	if err := j.validateSetWedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wed",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ResetFri() {
	_jsii_.InvokeVoid(
		z,
		"resetFri",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ResetMon() {
	_jsii_.InvokeVoid(
		z,
		"resetMon",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ResetSat() {
	_jsii_.InvokeVoid(
		z,
		"resetSat",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ResetSun() {
	_jsii_.InvokeVoid(
		z,
		"resetSun",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ResetThu() {
	_jsii_.InvokeVoid(
		z,
		"resetThu",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		z,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ResetTue() {
	_jsii_.InvokeVoid(
		z,
		"resetTue",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ResetWed() {
	_jsii_.InvokeVoid(
		z,
		"resetWed",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := z.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustGatewayPolicyScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

