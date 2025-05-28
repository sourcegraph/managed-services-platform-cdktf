package googleprivatecacertificatetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleprivatecacertificatetemplate/internal"
)

type GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference interface {
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
	Critical() interface{}
	SetCritical(val interface{})
	CriticalInput() interface{}
	ExcludedDnsNames() *[]*string
	SetExcludedDnsNames(val *[]*string)
	ExcludedDnsNamesInput() *[]*string
	ExcludedEmailAddresses() *[]*string
	SetExcludedEmailAddresses(val *[]*string)
	ExcludedEmailAddressesInput() *[]*string
	ExcludedIpRanges() *[]*string
	SetExcludedIpRanges(val *[]*string)
	ExcludedIpRangesInput() *[]*string
	ExcludedUris() *[]*string
	SetExcludedUris(val *[]*string)
	ExcludedUrisInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraints
	SetInternalValue(val *GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraints)
	PermittedDnsNames() *[]*string
	SetPermittedDnsNames(val *[]*string)
	PermittedDnsNamesInput() *[]*string
	PermittedEmailAddresses() *[]*string
	SetPermittedEmailAddresses(val *[]*string)
	PermittedEmailAddressesInput() *[]*string
	PermittedIpRanges() *[]*string
	SetPermittedIpRanges(val *[]*string)
	PermittedIpRangesInput() *[]*string
	PermittedUris() *[]*string
	SetPermittedUris(val *[]*string)
	PermittedUrisInput() *[]*string
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
	ResetExcludedDnsNames()
	ResetExcludedEmailAddresses()
	ResetExcludedIpRanges()
	ResetExcludedUris()
	ResetPermittedDnsNames()
	ResetPermittedEmailAddresses()
	ResetPermittedIpRanges()
	ResetPermittedUris()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference
type jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) Critical() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"critical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) CriticalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ExcludedDnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedDnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ExcludedDnsNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedDnsNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ExcludedEmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedEmailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ExcludedEmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedEmailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ExcludedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ExcludedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ExcludedUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ExcludedUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) InternalValue() *GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraints {
	var returns *GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraints
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) PermittedDnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) PermittedDnsNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) PermittedEmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedEmailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) PermittedEmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedEmailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) PermittedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) PermittedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) PermittedUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) PermittedUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference_Override(g GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetCritical(val interface{}) {
	if err := j.validateSetCriticalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"critical",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetExcludedDnsNames(val *[]*string) {
	if err := j.validateSetExcludedDnsNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedDnsNames",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetExcludedEmailAddresses(val *[]*string) {
	if err := j.validateSetExcludedEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedEmailAddresses",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetExcludedIpRanges(val *[]*string) {
	if err := j.validateSetExcludedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedIpRanges",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetExcludedUris(val *[]*string) {
	if err := j.validateSetExcludedUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedUris",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetInternalValue(val *GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraints) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetPermittedDnsNames(val *[]*string) {
	if err := j.validateSetPermittedDnsNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedDnsNames",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetPermittedEmailAddresses(val *[]*string) {
	if err := j.validateSetPermittedEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedEmailAddresses",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetPermittedIpRanges(val *[]*string) {
	if err := j.validateSetPermittedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedIpRanges",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetPermittedUris(val *[]*string) {
	if err := j.validateSetPermittedUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedUris",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ResetExcludedDnsNames() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedDnsNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ResetExcludedEmailAddresses() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedEmailAddresses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ResetExcludedIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ResetExcludedUris() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ResetPermittedDnsNames() {
	_jsii_.InvokeVoid(
		g,
		"resetPermittedDnsNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ResetPermittedEmailAddresses() {
	_jsii_.InvokeVoid(
		g,
		"resetPermittedEmailAddresses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ResetPermittedIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetPermittedIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ResetPermittedUris() {
	_jsii_.InvokeVoid(
		g,
		"resetPermittedUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesNameConstraintsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

