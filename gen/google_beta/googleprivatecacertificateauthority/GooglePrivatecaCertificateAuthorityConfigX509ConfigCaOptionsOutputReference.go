package googleprivatecacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleprivatecacertificateauthority/internal"
)

type GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference interface {
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
	InternalValue() *GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptions
	SetInternalValue(val *GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptions)
	IsCa() any
	SetIsCa(val any)
	IsCaInput() any
	MaxIssuerPathLength() *float64
	SetMaxIssuerPathLength(val *float64)
	MaxIssuerPathLengthInput() *float64
	NonCa() any
	SetNonCa(val any)
	NonCaInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ZeroMaxIssuerPathLength() any
	SetZeroMaxIssuerPathLength(val any)
	ZeroMaxIssuerPathLengthInput() any
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
	ResetMaxIssuerPathLength()
	ResetNonCa()
	ResetZeroMaxIssuerPathLength()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference
type jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) InternalValue() *GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptions {
	var returns *GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) IsCa() any {
	var returns any
	_jsii_.Get(
		j,
		"isCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) IsCaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"isCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) MaxIssuerPathLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIssuerPathLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) MaxIssuerPathLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIssuerPathLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) NonCa() any {
	var returns any
	_jsii_.Get(
		j,
		"nonCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) NonCaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"nonCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) ZeroMaxIssuerPathLength() any {
	var returns any
	_jsii_.Get(
		j,
		"zeroMaxIssuerPathLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) ZeroMaxIssuerPathLengthInput() any {
	var returns any
	_jsii_.Get(
		j,
		"zeroMaxIssuerPathLengthInput",
		&returns,
	)
	return returns
}

func NewGooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googlePrivatecaCertificateAuthority.GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference_Override(g GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googlePrivatecaCertificateAuthority.GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) SetInternalValue(val *GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) SetIsCa(val any) {
	if err := j.validateSetIsCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCa",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) SetMaxIssuerPathLength(val *float64) {
	if err := j.validateSetMaxIssuerPathLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIssuerPathLength",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) SetNonCa(val any) {
	if err := j.validateSetNonCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonCa",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) SetZeroMaxIssuerPathLength(val any) {
	if err := j.validateSetZeroMaxIssuerPathLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroMaxIssuerPathLength",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) ResetMaxIssuerPathLength() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxIssuerPathLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) ResetNonCa() {
	_jsii_.InvokeVoid(
		g,
		"resetNonCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) ResetZeroMaxIssuerPathLength() {
	_jsii_.InvokeVoid(
		g,
		"resetZeroMaxIssuerPathLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
