package zerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/zerotrustaccessapplication/internal"
)

type ZeroTrustAccessApplicationScimConfigOutputReference interface {
	cdktf.ComplexObject
	Authentication() ZeroTrustAccessApplicationScimConfigAuthenticationOutputReference
	AuthenticationInput() any
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
	DeactivateOnDelete() any
	SetDeactivateOnDelete(val any)
	DeactivateOnDeleteInput() any
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
	// Experimental.
	Fqn() *string
	IdpUid() *string
	SetIdpUid(val *string)
	IdpUidInput() *string
	InternalValue() any
	SetInternalValue(val any)
	Mappings() ZeroTrustAccessApplicationScimConfigMappingsList
	MappingsInput() any
	RemoteUri() *string
	SetRemoteUri(val *string)
	RemoteUriInput() *string
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
	PutAuthentication(value *ZeroTrustAccessApplicationScimConfigAuthentication)
	PutMappings(value any)
	ResetAuthentication()
	ResetDeactivateOnDelete()
	ResetEnabled()
	ResetMappings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustAccessApplicationScimConfigOutputReference
type jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) Authentication() ZeroTrustAccessApplicationScimConfigAuthenticationOutputReference {
	var returns ZeroTrustAccessApplicationScimConfigAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) AuthenticationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) DeactivateOnDelete() any {
	var returns any
	_jsii_.Get(
		j,
		"deactivateOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) DeactivateOnDeleteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"deactivateOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) IdpUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) IdpUidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) Mappings() ZeroTrustAccessApplicationScimConfigMappingsList {
	var returns ZeroTrustAccessApplicationScimConfigMappingsList
	_jsii_.Get(
		j,
		"mappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) MappingsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"mappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) RemoteUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) RemoteUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewZeroTrustAccessApplicationScimConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZeroTrustAccessApplicationScimConfigOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessApplicationScimConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationScimConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZeroTrustAccessApplicationScimConfigOutputReference_Override(z ZeroTrustAccessApplicationScimConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationScimConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) SetDeactivateOnDelete(val any) {
	if err := j.validateSetDeactivateOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deactivateOnDelete",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) SetIdpUid(val *string) {
	if err := j.validateSetIdpUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpUid",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) SetRemoteUri(val *string) {
	if err := j.validateSetRemoteUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteUri",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) PutAuthentication(value *ZeroTrustAccessApplicationScimConfigAuthentication) {
	if err := z.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthentication",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) PutMappings(value any) {
	if err := z.validatePutMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putMappings",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) ResetAuthentication() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) ResetDeactivateOnDelete() {
	_jsii_.InvokeVoid(
		z,
		"resetDeactivateOnDelete",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) ResetMappings() {
	_jsii_.InvokeVoid(
		z,
		"resetMappings",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := z.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		z,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
