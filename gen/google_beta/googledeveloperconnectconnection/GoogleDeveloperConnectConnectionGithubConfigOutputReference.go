package googledeveloperconnectconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledeveloperconnectconnection/internal"
)

type GoogleDeveloperConnectConnectionGithubConfigOutputReference interface {
	cdktf.ComplexObject
	AppInstallationId() *string
	SetAppInstallationId(val *string)
	AppInstallationIdInput() *string
	AuthorizerCredential() GoogleDeveloperConnectConnectionGithubConfigAuthorizerCredentialOutputReference
	AuthorizerCredentialInput() *GoogleDeveloperConnectConnectionGithubConfigAuthorizerCredential
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
	GithubApp() *string
	SetGithubApp(val *string)
	GithubAppInput() *string
	InstallationUri() *string
	InternalValue() *GoogleDeveloperConnectConnectionGithubConfig
	SetInternalValue(val *GoogleDeveloperConnectConnectionGithubConfig)
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
	PutAuthorizerCredential(value *GoogleDeveloperConnectConnectionGithubConfigAuthorizerCredential)
	ResetAppInstallationId()
	ResetAuthorizerCredential()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDeveloperConnectConnectionGithubConfigOutputReference
type jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) AppInstallationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInstallationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) AppInstallationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInstallationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) AuthorizerCredential() GoogleDeveloperConnectConnectionGithubConfigAuthorizerCredentialOutputReference {
	var returns GoogleDeveloperConnectConnectionGithubConfigAuthorizerCredentialOutputReference
	_jsii_.Get(
		j,
		"authorizerCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) AuthorizerCredentialInput() *GoogleDeveloperConnectConnectionGithubConfigAuthorizerCredential {
	var returns *GoogleDeveloperConnectConnectionGithubConfigAuthorizerCredential
	_jsii_.Get(
		j,
		"authorizerCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) GithubApp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"githubApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) GithubAppInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"githubAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) InstallationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) InternalValue() *GoogleDeveloperConnectConnectionGithubConfig {
	var returns *GoogleDeveloperConnectConnectionGithubConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDeveloperConnectConnectionGithubConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDeveloperConnectConnectionGithubConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDeveloperConnectConnectionGithubConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnectionGithubConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDeveloperConnectConnectionGithubConfigOutputReference_Override(g GoogleDeveloperConnectConnectionGithubConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnectionGithubConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference)SetAppInstallationId(val *string) {
	if err := j.validateSetAppInstallationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appInstallationId",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference)SetGithubApp(val *string) {
	if err := j.validateSetGithubAppParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"githubApp",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference)SetInternalValue(val *GoogleDeveloperConnectConnectionGithubConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) PutAuthorizerCredential(value *GoogleDeveloperConnectConnectionGithubConfigAuthorizerCredential) {
	if err := g.validatePutAuthorizerCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorizerCredential",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) ResetAppInstallationId() {
	_jsii_.InvokeVoid(
		g,
		"resetAppInstallationId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) ResetAuthorizerCredential() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizerCredential",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDeveloperConnectConnectionGithubConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

