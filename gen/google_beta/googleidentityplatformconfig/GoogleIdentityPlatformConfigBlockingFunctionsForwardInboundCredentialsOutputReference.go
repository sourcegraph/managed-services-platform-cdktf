package googleidentityplatformconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleidentityplatformconfig/internal"
)

type GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference interface {
	cdktf.ComplexObject
	AccessToken() any
	SetAccessToken(val any)
	AccessTokenInput() any
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
	IdToken() any
	SetIdToken(val any)
	IdTokenInput() any
	InternalValue() *GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials
	SetInternalValue(val *GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials)
	RefreshToken() any
	SetRefreshToken(val any)
	RefreshTokenInput() any
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
	ResetAccessToken()
	ResetIdToken()
	ResetRefreshToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference
type jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) AccessToken() any {
	var returns any
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) AccessTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) IdToken() any {
	var returns any
	_jsii_.Get(
		j,
		"idToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) IdTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"idTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) InternalValue() *GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials {
	var returns *GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) RefreshToken() any {
	var returns any
	_jsii_.Get(
		j,
		"refreshToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) RefreshTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"refreshTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference_Override(g GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) SetAccessToken(val any) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) SetIdToken(val any) {
	if err := j.validateSetIdTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idToken",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) SetInternalValue(val *GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) SetRefreshToken(val any) {
	if err := j.validateSetRefreshTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshToken",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ResetAccessToken() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ResetIdToken() {
	_jsii_.InvokeVoid(
		g,
		"resetIdToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ResetRefreshToken() {
	_jsii_.InvokeVoid(
		g,
		"resetRefreshToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
