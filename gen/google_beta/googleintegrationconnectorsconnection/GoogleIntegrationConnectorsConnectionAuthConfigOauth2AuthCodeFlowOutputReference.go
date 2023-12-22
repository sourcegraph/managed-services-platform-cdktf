package googleintegrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleintegrationconnectorsconnection/internal"
)

type GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference interface {
	cdktf.ComplexObject
	AuthUri() *string
	SetAuthUri(val *string)
	AuthUriInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecretOutputReference
	ClientSecretInput() *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret
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
	EnablePkce() interface{}
	SetEnablePkce(val interface{})
	EnablePkceInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow
	SetInternalValue(val *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow)
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
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
	PutClientSecret(value *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret)
	ResetAuthUri()
	ResetClientId()
	ResetClientSecret()
	ResetEnablePkce()
	ResetScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference
type jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) AuthUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) AuthUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ClientSecret() GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecretOutputReference {
	var returns GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecretOutputReference
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ClientSecretInput() *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret {
	var returns *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) EnablePkce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePkce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) EnablePkceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePkceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) InternalValue() *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow {
	var returns *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference_Override(g GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetAuthUri(val *string) {
	if err := j.validateSetAuthUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUri",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetEnablePkce(val interface{}) {
	if err := j.validateSetEnablePkceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePkce",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetInternalValue(val *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) PutClientSecret(value *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret) {
	if err := g.validatePutClientSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientSecret",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ResetAuthUri() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		g,
		"resetClientId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ResetEnablePkce() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePkce",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ResetScopes() {
	_jsii_.InvokeVoid(
		g,
		"resetScopes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

