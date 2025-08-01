package customhostname

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/customhostname/internal"
)

type CustomHostnameSslOutputReference interface {
	cdktf.ComplexObject
	BundleMethod() *string
	SetBundleMethod(val *string)
	BundleMethodInput() *string
	CertificateAuthority() *string
	SetCertificateAuthority(val *string)
	CertificateAuthorityInput() *string
	CloudflareBranding() interface{}
	SetCloudflareBranding(val interface{})
	CloudflareBrandingInput() interface{}
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
	CustomCertBundle() CustomHostnameSslCustomCertBundleList
	CustomCertBundleInput() interface{}
	CustomCertificate() *string
	SetCustomCertificate(val *string)
	CustomCertificateInput() *string
	CustomKey() *string
	SetCustomKey(val *string)
	CustomKeyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Method() *string
	SetMethod(val *string)
	MethodInput() *string
	Settings() CustomHostnameSslSettingsOutputReference
	SettingsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Wildcard() interface{}
	SetWildcard(val interface{})
	WildcardInput() interface{}
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
	PutCustomCertBundle(value interface{})
	PutSettings(value *CustomHostnameSslSettings)
	ResetBundleMethod()
	ResetCertificateAuthority()
	ResetCloudflareBranding()
	ResetCustomCertBundle()
	ResetCustomCertificate()
	ResetCustomKey()
	ResetMethod()
	ResetSettings()
	ResetType()
	ResetWildcard()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomHostnameSslOutputReference
type jsiiProxy_CustomHostnameSslOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) BundleMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) BundleMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) CertificateAuthority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) CertificateAuthorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) CloudflareBranding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudflareBranding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) CloudflareBrandingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudflareBrandingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) CustomCertBundle() CustomHostnameSslCustomCertBundleList {
	var returns CustomHostnameSslCustomCertBundleList
	_jsii_.Get(
		j,
		"customCertBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) CustomCertBundleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customCertBundleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) CustomCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) CustomCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) CustomKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) CustomKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) Settings() CustomHostnameSslSettingsOutputReference {
	var returns CustomHostnameSslSettingsOutputReference
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) SettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) Wildcard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomHostnameSslOutputReference) WildcardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcardInput",
		&returns,
	)
	return returns
}


func NewCustomHostnameSslOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomHostnameSslOutputReference {
	_init_.Initialize()

	if err := validateNewCustomHostnameSslOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomHostnameSslOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.customHostname.CustomHostnameSslOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomHostnameSslOutputReference_Override(c CustomHostnameSslOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.customHostname.CustomHostnameSslOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetBundleMethod(val *string) {
	if err := j.validateSetBundleMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bundleMethod",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetCertificateAuthority(val *string) {
	if err := j.validateSetCertificateAuthorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateAuthority",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetCloudflareBranding(val interface{}) {
	if err := j.validateSetCloudflareBrandingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudflareBranding",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetCustomCertificate(val *string) {
	if err := j.validateSetCustomCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customCertificate",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetCustomKey(val *string) {
	if err := j.validateSetCustomKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customKey",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CustomHostnameSslOutputReference)SetWildcard(val interface{}) {
	if err := j.validateSetWildcardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildcard",
		val,
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) PutCustomCertBundle(value interface{}) {
	if err := c.validatePutCustomCertBundleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomCertBundle",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) PutSettings(value *CustomHostnameSslSettings) {
	if err := c.validatePutSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ResetBundleMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetBundleMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ResetCertificateAuthority() {
	_jsii_.InvokeVoid(
		c,
		"resetCertificateAuthority",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ResetCloudflareBranding() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudflareBranding",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ResetCustomCertBundle() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomCertBundle",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ResetCustomCertificate() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomCertificate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ResetCustomKey() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ResetSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ResetWildcard() {
	_jsii_.InvokeVoid(
		c,
		"resetWildcard",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomHostnameSslOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

