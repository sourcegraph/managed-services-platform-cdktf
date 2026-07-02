package zerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/zerotrustaccessapplication/internal"
)

type ZeroTrustAccessApplicationPoliciesRequireOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenOutputReference
	AnyValidServiceTokenInput() any
	AuthContext() ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference
	AuthContextInput() any
	AuthMethod() ZeroTrustAccessApplicationPoliciesRequireAuthMethodOutputReference
	AuthMethodInput() any
	AzureAd() ZeroTrustAccessApplicationPoliciesRequireAzureAdOutputReference
	AzureAdInput() any
	Certificate() ZeroTrustAccessApplicationPoliciesRequireCertificateOutputReference
	CertificateInput() any
	CommonName() ZeroTrustAccessApplicationPoliciesRequireCommonNameOutputReference
	CommonNameInput() any
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
	DevicePosture() ZeroTrustAccessApplicationPoliciesRequireDevicePostureOutputReference
	DevicePostureInput() any
	Email() ZeroTrustAccessApplicationPoliciesRequireEmailOutputReference
	EmailDomain() ZeroTrustAccessApplicationPoliciesRequireEmailDomainOutputReference
	EmailDomainInput() any
	EmailInput() any
	EmailList() ZeroTrustAccessApplicationPoliciesRequireEmailListStructOutputReference
	EmailListInput() any
	Everyone() ZeroTrustAccessApplicationPoliciesRequireEveryoneOutputReference
	EveryoneInput() any
	ExternalEvaluation() ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationOutputReference
	ExternalEvaluationInput() any
	// Experimental.
	Fqn() *string
	Geo() ZeroTrustAccessApplicationPoliciesRequireGeoOutputReference
	GeoInput() any
	GithubOrganization() ZeroTrustAccessApplicationPoliciesRequireGithubOrganizationOutputReference
	GithubOrganizationInput() any
	Group() ZeroTrustAccessApplicationPoliciesRequireGroupOutputReference
	GroupInput() any
	Gsuite() ZeroTrustAccessApplicationPoliciesRequireGsuiteOutputReference
	GsuiteInput() any
	InternalValue() any
	SetInternalValue(val any)
	Ip() ZeroTrustAccessApplicationPoliciesRequireIpOutputReference
	IpInput() any
	IpList() ZeroTrustAccessApplicationPoliciesRequireIpListStructOutputReference
	IpListInput() any
	LoginMethod() ZeroTrustAccessApplicationPoliciesRequireLoginMethodOutputReference
	LoginMethodInput() any
	Oidc() ZeroTrustAccessApplicationPoliciesRequireOidcOutputReference
	OidcInput() any
	Okta() ZeroTrustAccessApplicationPoliciesRequireOktaOutputReference
	OktaInput() any
	Saml() ZeroTrustAccessApplicationPoliciesRequireSamlOutputReference
	SamlInput() any
	ServiceToken() ZeroTrustAccessApplicationPoliciesRequireServiceTokenOutputReference
	ServiceTokenInput() any
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
	PutAnyValidServiceToken(value *ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceToken)
	PutAuthContext(value *ZeroTrustAccessApplicationPoliciesRequireAuthContext)
	PutAuthMethod(value *ZeroTrustAccessApplicationPoliciesRequireAuthMethod)
	PutAzureAd(value *ZeroTrustAccessApplicationPoliciesRequireAzureAd)
	PutCertificate(value *ZeroTrustAccessApplicationPoliciesRequireCertificate)
	PutCommonName(value *ZeroTrustAccessApplicationPoliciesRequireCommonName)
	PutDevicePosture(value *ZeroTrustAccessApplicationPoliciesRequireDevicePosture)
	PutEmail(value *ZeroTrustAccessApplicationPoliciesRequireEmail)
	PutEmailDomain(value *ZeroTrustAccessApplicationPoliciesRequireEmailDomain)
	PutEmailList(value *ZeroTrustAccessApplicationPoliciesRequireEmailListStruct)
	PutEveryone(value *ZeroTrustAccessApplicationPoliciesRequireEveryone)
	PutExternalEvaluation(value *ZeroTrustAccessApplicationPoliciesRequireExternalEvaluation)
	PutGeo(value *ZeroTrustAccessApplicationPoliciesRequireGeo)
	PutGithubOrganization(value *ZeroTrustAccessApplicationPoliciesRequireGithubOrganization)
	PutGroup(value *ZeroTrustAccessApplicationPoliciesRequireGroup)
	PutGsuite(value *ZeroTrustAccessApplicationPoliciesRequireGsuite)
	PutIp(value *ZeroTrustAccessApplicationPoliciesRequireIp)
	PutIpList(value *ZeroTrustAccessApplicationPoliciesRequireIpListStruct)
	PutLoginMethod(value *ZeroTrustAccessApplicationPoliciesRequireLoginMethod)
	PutOidc(value *ZeroTrustAccessApplicationPoliciesRequireOidc)
	PutOkta(value *ZeroTrustAccessApplicationPoliciesRequireOkta)
	PutSaml(value *ZeroTrustAccessApplicationPoliciesRequireSaml)
	PutServiceToken(value *ZeroTrustAccessApplicationPoliciesRequireServiceToken)
	ResetAnyValidServiceToken()
	ResetAuthContext()
	ResetAuthMethod()
	ResetAzureAd()
	ResetCertificate()
	ResetCommonName()
	ResetDevicePosture()
	ResetEmail()
	ResetEmailDomain()
	ResetEmailList()
	ResetEveryone()
	ResetExternalEvaluation()
	ResetGeo()
	ResetGithubOrganization()
	ResetGroup()
	ResetGsuite()
	ResetIp()
	ResetIpList()
	ResetLoginMethod()
	ResetOidc()
	ResetOkta()
	ResetSaml()
	ResetServiceToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustAccessApplicationPoliciesRequireOutputReference
type jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) AnyValidServiceToken() ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) AnyValidServiceTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) AuthContext() ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) AuthContextInput() any {
	var returns any
	_jsii_.Get(
		j,
		"authContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) AuthMethod() ZeroTrustAccessApplicationPoliciesRequireAuthMethodOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) AuthMethodInput() any {
	var returns any
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) AzureAd() ZeroTrustAccessApplicationPoliciesRequireAzureAdOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) AzureAdInput() any {
	var returns any
	_jsii_.Get(
		j,
		"azureAdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Certificate() ZeroTrustAccessApplicationPoliciesRequireCertificateOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) CertificateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) CommonName() ZeroTrustAccessApplicationPoliciesRequireCommonNameOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) CommonNameInput() any {
	var returns any
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) DevicePosture() ZeroTrustAccessApplicationPoliciesRequireDevicePostureOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) DevicePostureInput() any {
	var returns any
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Email() ZeroTrustAccessApplicationPoliciesRequireEmailOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) EmailDomain() ZeroTrustAccessApplicationPoliciesRequireEmailDomainOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) EmailDomainInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) EmailInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) EmailList() ZeroTrustAccessApplicationPoliciesRequireEmailListStructOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) EmailListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Everyone() ZeroTrustAccessApplicationPoliciesRequireEveryoneOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) EveryoneInput() any {
	var returns any
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ExternalEvaluation() ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ExternalEvaluationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Geo() ZeroTrustAccessApplicationPoliciesRequireGeoOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GeoInput() any {
	var returns any
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GithubOrganization() ZeroTrustAccessApplicationPoliciesRequireGithubOrganizationOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GithubOrganizationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"githubOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Group() ZeroTrustAccessApplicationPoliciesRequireGroupOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GroupInput() any {
	var returns any
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Gsuite() ZeroTrustAccessApplicationPoliciesRequireGsuiteOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GsuiteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Ip() ZeroTrustAccessApplicationPoliciesRequireIpOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) IpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) IpList() ZeroTrustAccessApplicationPoliciesRequireIpListStructOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) IpListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) LoginMethod() ZeroTrustAccessApplicationPoliciesRequireLoginMethodOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) LoginMethodInput() any {
	var returns any
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Oidc() ZeroTrustAccessApplicationPoliciesRequireOidcOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) OidcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"oidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Okta() ZeroTrustAccessApplicationPoliciesRequireOktaOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) OktaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Saml() ZeroTrustAccessApplicationPoliciesRequireSamlOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) SamlInput() any {
	var returns any
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ServiceToken() ZeroTrustAccessApplicationPoliciesRequireServiceTokenOutputReference {
	var returns ZeroTrustAccessApplicationPoliciesRequireServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ServiceTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewZeroTrustAccessApplicationPoliciesRequireOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustAccessApplicationPoliciesRequireOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessApplicationPoliciesRequireOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustAccessApplicationPoliciesRequireOutputReference_Override(z ZeroTrustAccessApplicationPoliciesRequireOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutAnyValidServiceToken(value *ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceToken) {
	if err := z.validatePutAnyValidServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnyValidServiceToken",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutAuthContext(value *ZeroTrustAccessApplicationPoliciesRequireAuthContext) {
	if err := z.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthContext",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutAuthMethod(value *ZeroTrustAccessApplicationPoliciesRequireAuthMethod) {
	if err := z.validatePutAuthMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthMethod",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutAzureAd(value *ZeroTrustAccessApplicationPoliciesRequireAzureAd) {
	if err := z.validatePutAzureAdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAzureAd",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutCertificate(value *ZeroTrustAccessApplicationPoliciesRequireCertificate) {
	if err := z.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCertificate",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutCommonName(value *ZeroTrustAccessApplicationPoliciesRequireCommonName) {
	if err := z.validatePutCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCommonName",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutDevicePosture(value *ZeroTrustAccessApplicationPoliciesRequireDevicePosture) {
	if err := z.validatePutDevicePostureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDevicePosture",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutEmail(value *ZeroTrustAccessApplicationPoliciesRequireEmail) {
	if err := z.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmail",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutEmailDomain(value *ZeroTrustAccessApplicationPoliciesRequireEmailDomain) {
	if err := z.validatePutEmailDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailDomain",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutEmailList(value *ZeroTrustAccessApplicationPoliciesRequireEmailListStruct) {
	if err := z.validatePutEmailListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailList",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutEveryone(value *ZeroTrustAccessApplicationPoliciesRequireEveryone) {
	if err := z.validatePutEveryoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEveryone",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutExternalEvaluation(value *ZeroTrustAccessApplicationPoliciesRequireExternalEvaluation) {
	if err := z.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExternalEvaluation",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutGeo(value *ZeroTrustAccessApplicationPoliciesRequireGeo) {
	if err := z.validatePutGeoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGeo",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutGithubOrganization(value *ZeroTrustAccessApplicationPoliciesRequireGithubOrganization) {
	if err := z.validatePutGithubOrganizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGithubOrganization",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutGroup(value *ZeroTrustAccessApplicationPoliciesRequireGroup) {
	if err := z.validatePutGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGroup",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutGsuite(value *ZeroTrustAccessApplicationPoliciesRequireGsuite) {
	if err := z.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGsuite",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutIp(value *ZeroTrustAccessApplicationPoliciesRequireIp) {
	if err := z.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIp",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutIpList(value *ZeroTrustAccessApplicationPoliciesRequireIpListStruct) {
	if err := z.validatePutIpListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpList",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutLoginMethod(value *ZeroTrustAccessApplicationPoliciesRequireLoginMethod) {
	if err := z.validatePutLoginMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLoginMethod",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutOidc(value *ZeroTrustAccessApplicationPoliciesRequireOidc) {
	if err := z.validatePutOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOidc",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutOkta(value *ZeroTrustAccessApplicationPoliciesRequireOkta) {
	if err := z.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOkta",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutSaml(value *ZeroTrustAccessApplicationPoliciesRequireSaml) {
	if err := z.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaml",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) PutServiceToken(value *ZeroTrustAccessApplicationPoliciesRequireServiceToken) {
	if err := z.validatePutServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putServiceToken",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetAuthContext() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthContext",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetAzureAd() {
	_jsii_.InvokeVoid(
		z,
		"resetAzureAd",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		z,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		z,
		"resetEmail",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetEmailList() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		z,
		"resetEveryone",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		z,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		z,
		"resetGeo",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetGithubOrganization() {
	_jsii_.InvokeVoid(
		z,
		"resetGithubOrganization",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		z,
		"resetGroup",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		z,
		"resetGsuite",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		z,
		"resetIp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		z,
		"resetIpList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetLoginMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetOidc() {
	_jsii_.InvokeVoid(
		z,
		"resetOidc",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		z,
		"resetOkta",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		z,
		"resetSaml",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (z *jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
