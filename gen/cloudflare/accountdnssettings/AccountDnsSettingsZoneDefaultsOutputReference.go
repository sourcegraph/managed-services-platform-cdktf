package accountdnssettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/accountdnssettings/internal"
)

type AccountDnsSettingsZoneDefaultsOutputReference interface {
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
	FlattenAllCnames() any
	SetFlattenAllCnames(val any)
	FlattenAllCnamesInput() any
	FoundationDns() any
	SetFoundationDns(val any)
	FoundationDnsInput() any
	// Experimental.
	Fqn() *string
	InternalDns() AccountDnsSettingsZoneDefaultsInternalDnsOutputReference
	InternalDnsInput() any
	InternalValue() any
	SetInternalValue(val any)
	MultiProvider() any
	SetMultiProvider(val any)
	MultiProviderInput() any
	Nameservers() AccountDnsSettingsZoneDefaultsNameserversOutputReference
	NameserversInput() any
	NsTtl() *float64
	SetNsTtl(val *float64)
	NsTtlInput() *float64
	SecondaryOverrides() any
	SetSecondaryOverrides(val any)
	SecondaryOverridesInput() any
	Soa() AccountDnsSettingsZoneDefaultsSoaOutputReference
	SoaInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ZoneMode() *string
	SetZoneMode(val *string)
	ZoneModeInput() *string
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
	PutInternalDns(value *AccountDnsSettingsZoneDefaultsInternalDns)
	PutNameservers(value *AccountDnsSettingsZoneDefaultsNameservers)
	PutSoa(value *AccountDnsSettingsZoneDefaultsSoa)
	ResetFlattenAllCnames()
	ResetFoundationDns()
	ResetInternalDns()
	ResetMultiProvider()
	ResetNameservers()
	ResetNsTtl()
	ResetSecondaryOverrides()
	ResetSoa()
	ResetZoneMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccountDnsSettingsZoneDefaultsOutputReference
type jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) FlattenAllCnames() any {
	var returns any
	_jsii_.Get(
		j,
		"flattenAllCnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) FlattenAllCnamesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"flattenAllCnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) FoundationDns() any {
	var returns any
	_jsii_.Get(
		j,
		"foundationDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) FoundationDnsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"foundationDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) InternalDns() AccountDnsSettingsZoneDefaultsInternalDnsOutputReference {
	var returns AccountDnsSettingsZoneDefaultsInternalDnsOutputReference
	_jsii_.Get(
		j,
		"internalDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) InternalDnsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"internalDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) MultiProvider() any {
	var returns any
	_jsii_.Get(
		j,
		"multiProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) MultiProviderInput() any {
	var returns any
	_jsii_.Get(
		j,
		"multiProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) Nameservers() AccountDnsSettingsZoneDefaultsNameserversOutputReference {
	var returns AccountDnsSettingsZoneDefaultsNameserversOutputReference
	_jsii_.Get(
		j,
		"nameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) NameserversInput() any {
	var returns any
	_jsii_.Get(
		j,
		"nameserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) NsTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nsTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) NsTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nsTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SecondaryOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"secondaryOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SecondaryOverridesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"secondaryOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) Soa() AccountDnsSettingsZoneDefaultsSoaOutputReference {
	var returns AccountDnsSettingsZoneDefaultsSoaOutputReference
	_jsii_.Get(
		j,
		"soa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SoaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"soaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ZoneMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ZoneModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneModeInput",
		&returns,
	)
	return returns
}

func NewAccountDnsSettingsZoneDefaultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccountDnsSettingsZoneDefaultsOutputReference {
	_init_.Initialize()

	if err := validateNewAccountDnsSettingsZoneDefaultsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accountDnsSettings.AccountDnsSettingsZoneDefaultsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccountDnsSettingsZoneDefaultsOutputReference_Override(a AccountDnsSettingsZoneDefaultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.accountDnsSettings.AccountDnsSettingsZoneDefaultsOutputReference",
		[]any{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SetFlattenAllCnames(val any) {
	if err := j.validateSetFlattenAllCnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flattenAllCnames",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SetFoundationDns(val any) {
	if err := j.validateSetFoundationDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"foundationDns",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SetMultiProvider(val any) {
	if err := j.validateSetMultiProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiProvider",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SetNsTtl(val *float64) {
	if err := j.validateSetNsTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsTtl",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SetSecondaryOverrides(val any) {
	if err := j.validateSetSecondaryOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryOverrides",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) SetZoneMode(val *string) {
	if err := j.validateSetZoneModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneMode",
		val,
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) PutInternalDns(value *AccountDnsSettingsZoneDefaultsInternalDns) {
	if err := a.validatePutInternalDnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInternalDns",
		[]any{value},
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) PutNameservers(value *AccountDnsSettingsZoneDefaultsNameservers) {
	if err := a.validatePutNameserversParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNameservers",
		[]any{value},
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) PutSoa(value *AccountDnsSettingsZoneDefaultsSoa) {
	if err := a.validatePutSoaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSoa",
		[]any{value},
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetFlattenAllCnames() {
	_jsii_.InvokeVoid(
		a,
		"resetFlattenAllCnames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetFoundationDns() {
	_jsii_.InvokeVoid(
		a,
		"resetFoundationDns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetInternalDns() {
	_jsii_.InvokeVoid(
		a,
		"resetInternalDns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetMultiProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetNameservers() {
	_jsii_.InvokeVoid(
		a,
		"resetNameservers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetNsTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetNsTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetSecondaryOverrides() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondaryOverrides",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetSoa() {
	_jsii_.InvokeVoid(
		a,
		"resetSoa",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ResetZoneMode() {
	_jsii_.InvokeVoid(
		a,
		"resetZoneMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		a,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountDnsSettingsZoneDefaultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
