package googleprivatecacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleprivatecacertificate/internal"
)

type GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference interface {
	cdktf.ComplexObject
	CertSign() any
	SetCertSign(val any)
	CertSignInput() any
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
	ContentCommitment() any
	SetContentCommitment(val any)
	ContentCommitmentInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrlSign() any
	SetCrlSign(val any)
	CrlSignInput() any
	DataEncipherment() any
	SetDataEncipherment(val any)
	DataEnciphermentInput() any
	DecipherOnly() any
	SetDecipherOnly(val any)
	DecipherOnlyInput() any
	DigitalSignature() any
	SetDigitalSignature(val any)
	DigitalSignatureInput() any
	EncipherOnly() any
	SetEncipherOnly(val any)
	EncipherOnlyInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage
	SetInternalValue(val *GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage)
	KeyAgreement() any
	SetKeyAgreement(val any)
	KeyAgreementInput() any
	KeyEncipherment() any
	SetKeyEncipherment(val any)
	KeyEnciphermentInput() any
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
	ResetCertSign()
	ResetContentCommitment()
	ResetCrlSign()
	ResetDataEncipherment()
	ResetDecipherOnly()
	ResetDigitalSignature()
	ResetEncipherOnly()
	ResetKeyAgreement()
	ResetKeyEncipherment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference
type jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) CertSign() any {
	var returns any
	_jsii_.Get(
		j,
		"certSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) CertSignInput() any {
	var returns any
	_jsii_.Get(
		j,
		"certSignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ContentCommitment() any {
	var returns any
	_jsii_.Get(
		j,
		"contentCommitment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ContentCommitmentInput() any {
	var returns any
	_jsii_.Get(
		j,
		"contentCommitmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) CrlSign() any {
	var returns any
	_jsii_.Get(
		j,
		"crlSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) CrlSignInput() any {
	var returns any
	_jsii_.Get(
		j,
		"crlSignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DataEncipherment() any {
	var returns any
	_jsii_.Get(
		j,
		"dataEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DataEnciphermentInput() any {
	var returns any
	_jsii_.Get(
		j,
		"dataEnciphermentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DecipherOnly() any {
	var returns any
	_jsii_.Get(
		j,
		"decipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DecipherOnlyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"decipherOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DigitalSignature() any {
	var returns any
	_jsii_.Get(
		j,
		"digitalSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DigitalSignatureInput() any {
	var returns any
	_jsii_.Get(
		j,
		"digitalSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) EncipherOnly() any {
	var returns any
	_jsii_.Get(
		j,
		"encipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) EncipherOnlyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"encipherOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) InternalValue() *GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage {
	var returns *GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) KeyAgreement() any {
	var returns any
	_jsii_.Get(
		j,
		"keyAgreement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) KeyAgreementInput() any {
	var returns any
	_jsii_.Get(
		j,
		"keyAgreementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) KeyEncipherment() any {
	var returns any
	_jsii_.Get(
		j,
		"keyEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) KeyEnciphermentInput() any {
	var returns any
	_jsii_.Get(
		j,
		"keyEnciphermentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googlePrivatecaCertificate.GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference_Override(g GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googlePrivatecaCertificate.GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetCertSign(val any) {
	if err := j.validateSetCertSignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certSign",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetContentCommitment(val any) {
	if err := j.validateSetContentCommitmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentCommitment",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetCrlSign(val any) {
	if err := j.validateSetCrlSignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crlSign",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetDataEncipherment(val any) {
	if err := j.validateSetDataEnciphermentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataEncipherment",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetDecipherOnly(val any) {
	if err := j.validateSetDecipherOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decipherOnly",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetDigitalSignature(val any) {
	if err := j.validateSetDigitalSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digitalSignature",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetEncipherOnly(val any) {
	if err := j.validateSetEncipherOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encipherOnly",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetInternalValue(val *GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetKeyAgreement(val any) {
	if err := j.validateSetKeyAgreementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyAgreement",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetKeyEncipherment(val any) {
	if err := j.validateSetKeyEnciphermentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyEncipherment",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetCertSign() {
	_jsii_.InvokeVoid(
		g,
		"resetCertSign",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetContentCommitment() {
	_jsii_.InvokeVoid(
		g,
		"resetContentCommitment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetCrlSign() {
	_jsii_.InvokeVoid(
		g,
		"resetCrlSign",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetDataEncipherment() {
	_jsii_.InvokeVoid(
		g,
		"resetDataEncipherment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetDecipherOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetDecipherOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetDigitalSignature() {
	_jsii_.InvokeVoid(
		g,
		"resetDigitalSignature",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetEncipherOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetEncipherOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetKeyAgreement() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyAgreement",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetKeyEncipherment() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyEncipherment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
