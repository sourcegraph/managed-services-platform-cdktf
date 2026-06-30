package zerotrustaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/zerotrustaccesspolicy/internal"
)

type ZeroTrustAccessPolicyRequireOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() ZeroTrustAccessPolicyRequireAnyValidServiceTokenOutputReference
	AnyValidServiceTokenInput() any
	AuthContext() ZeroTrustAccessPolicyRequireAuthContextOutputReference
	AuthContextInput() any
	AuthMethod() ZeroTrustAccessPolicyRequireAuthMethodOutputReference
	AuthMethodInput() any
	AzureAd() ZeroTrustAccessPolicyRequireAzureAdOutputReference
	AzureAdInput() any
	Certificate() ZeroTrustAccessPolicyRequireCertificateOutputReference
	CertificateInput() any
	CommonName() ZeroTrustAccessPolicyRequireCommonNameOutputReference
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
	DevicePosture() ZeroTrustAccessPolicyRequireDevicePostureOutputReference
	DevicePostureInput() any
	Email() ZeroTrustAccessPolicyRequireEmailOutputReference
	EmailDomain() ZeroTrustAccessPolicyRequireEmailDomainOutputReference
	EmailDomainInput() any
	EmailInput() any
	EmailList() ZeroTrustAccessPolicyRequireEmailListStructOutputReference
	EmailListInput() any
	Everyone() ZeroTrustAccessPolicyRequireEveryoneOutputReference
	EveryoneInput() any
	ExternalEvaluation() ZeroTrustAccessPolicyRequireExternalEvaluationOutputReference
	ExternalEvaluationInput() any
	// Experimental.
	Fqn() *string
	Geo() ZeroTrustAccessPolicyRequireGeoOutputReference
	GeoInput() any
	GithubOrganization() ZeroTrustAccessPolicyRequireGithubOrganizationOutputReference
	GithubOrganizationInput() any
	Group() ZeroTrustAccessPolicyRequireGroupOutputReference
	GroupInput() any
	Gsuite() ZeroTrustAccessPolicyRequireGsuiteOutputReference
	GsuiteInput() any
	InternalValue() any
	SetInternalValue(val any)
	Ip() ZeroTrustAccessPolicyRequireIpOutputReference
	IpInput() any
	IpList() ZeroTrustAccessPolicyRequireIpListStructOutputReference
	IpListInput() any
	LoginMethod() ZeroTrustAccessPolicyRequireLoginMethodOutputReference
	LoginMethodInput() any
	Oidc() ZeroTrustAccessPolicyRequireOidcOutputReference
	OidcInput() any
	Okta() ZeroTrustAccessPolicyRequireOktaOutputReference
	OktaInput() any
	Saml() ZeroTrustAccessPolicyRequireSamlOutputReference
	SamlInput() any
	ServiceToken() ZeroTrustAccessPolicyRequireServiceTokenOutputReference
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
	PutAnyValidServiceToken(value *ZeroTrustAccessPolicyRequireAnyValidServiceToken)
	PutAuthContext(value *ZeroTrustAccessPolicyRequireAuthContext)
	PutAuthMethod(value *ZeroTrustAccessPolicyRequireAuthMethod)
	PutAzureAd(value *ZeroTrustAccessPolicyRequireAzureAd)
	PutCertificate(value *ZeroTrustAccessPolicyRequireCertificate)
	PutCommonName(value *ZeroTrustAccessPolicyRequireCommonName)
	PutDevicePosture(value *ZeroTrustAccessPolicyRequireDevicePosture)
	PutEmail(value *ZeroTrustAccessPolicyRequireEmail)
	PutEmailDomain(value *ZeroTrustAccessPolicyRequireEmailDomain)
	PutEmailList(value *ZeroTrustAccessPolicyRequireEmailListStruct)
	PutEveryone(value *ZeroTrustAccessPolicyRequireEveryone)
	PutExternalEvaluation(value *ZeroTrustAccessPolicyRequireExternalEvaluation)
	PutGeo(value *ZeroTrustAccessPolicyRequireGeo)
	PutGithubOrganization(value *ZeroTrustAccessPolicyRequireGithubOrganization)
	PutGroup(value *ZeroTrustAccessPolicyRequireGroup)
	PutGsuite(value *ZeroTrustAccessPolicyRequireGsuite)
	PutIp(value *ZeroTrustAccessPolicyRequireIp)
	PutIpList(value *ZeroTrustAccessPolicyRequireIpListStruct)
	PutLoginMethod(value *ZeroTrustAccessPolicyRequireLoginMethod)
	PutOidc(value *ZeroTrustAccessPolicyRequireOidc)
	PutOkta(value *ZeroTrustAccessPolicyRequireOkta)
	PutSaml(value *ZeroTrustAccessPolicyRequireSaml)
	PutServiceToken(value *ZeroTrustAccessPolicyRequireServiceToken)
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

// The jsii proxy struct for ZeroTrustAccessPolicyRequireOutputReference
type jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) AnyValidServiceToken() ZeroTrustAccessPolicyRequireAnyValidServiceTokenOutputReference {
	var returns ZeroTrustAccessPolicyRequireAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) AnyValidServiceTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) AuthContext() ZeroTrustAccessPolicyRequireAuthContextOutputReference {
	var returns ZeroTrustAccessPolicyRequireAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) AuthContextInput() any {
	var returns any
	_jsii_.Get(
		j,
		"authContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) AuthMethod() ZeroTrustAccessPolicyRequireAuthMethodOutputReference {
	var returns ZeroTrustAccessPolicyRequireAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) AuthMethodInput() any {
	var returns any
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) AzureAd() ZeroTrustAccessPolicyRequireAzureAdOutputReference {
	var returns ZeroTrustAccessPolicyRequireAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) AzureAdInput() any {
	var returns any
	_jsii_.Get(
		j,
		"azureAdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Certificate() ZeroTrustAccessPolicyRequireCertificateOutputReference {
	var returns ZeroTrustAccessPolicyRequireCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) CertificateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) CommonName() ZeroTrustAccessPolicyRequireCommonNameOutputReference {
	var returns ZeroTrustAccessPolicyRequireCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) CommonNameInput() any {
	var returns any
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) DevicePosture() ZeroTrustAccessPolicyRequireDevicePostureOutputReference {
	var returns ZeroTrustAccessPolicyRequireDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) DevicePostureInput() any {
	var returns any
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Email() ZeroTrustAccessPolicyRequireEmailOutputReference {
	var returns ZeroTrustAccessPolicyRequireEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) EmailDomain() ZeroTrustAccessPolicyRequireEmailDomainOutputReference {
	var returns ZeroTrustAccessPolicyRequireEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) EmailDomainInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) EmailInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) EmailList() ZeroTrustAccessPolicyRequireEmailListStructOutputReference {
	var returns ZeroTrustAccessPolicyRequireEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) EmailListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Everyone() ZeroTrustAccessPolicyRequireEveryoneOutputReference {
	var returns ZeroTrustAccessPolicyRequireEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) EveryoneInput() any {
	var returns any
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ExternalEvaluation() ZeroTrustAccessPolicyRequireExternalEvaluationOutputReference {
	var returns ZeroTrustAccessPolicyRequireExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ExternalEvaluationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Geo() ZeroTrustAccessPolicyRequireGeoOutputReference {
	var returns ZeroTrustAccessPolicyRequireGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GeoInput() any {
	var returns any
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GithubOrganization() ZeroTrustAccessPolicyRequireGithubOrganizationOutputReference {
	var returns ZeroTrustAccessPolicyRequireGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GithubOrganizationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"githubOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Group() ZeroTrustAccessPolicyRequireGroupOutputReference {
	var returns ZeroTrustAccessPolicyRequireGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GroupInput() any {
	var returns any
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Gsuite() ZeroTrustAccessPolicyRequireGsuiteOutputReference {
	var returns ZeroTrustAccessPolicyRequireGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GsuiteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Ip() ZeroTrustAccessPolicyRequireIpOutputReference {
	var returns ZeroTrustAccessPolicyRequireIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) IpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) IpList() ZeroTrustAccessPolicyRequireIpListStructOutputReference {
	var returns ZeroTrustAccessPolicyRequireIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) IpListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) LoginMethod() ZeroTrustAccessPolicyRequireLoginMethodOutputReference {
	var returns ZeroTrustAccessPolicyRequireLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) LoginMethodInput() any {
	var returns any
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Oidc() ZeroTrustAccessPolicyRequireOidcOutputReference {
	var returns ZeroTrustAccessPolicyRequireOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) OidcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"oidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Okta() ZeroTrustAccessPolicyRequireOktaOutputReference {
	var returns ZeroTrustAccessPolicyRequireOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) OktaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Saml() ZeroTrustAccessPolicyRequireSamlOutputReference {
	var returns ZeroTrustAccessPolicyRequireSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) SamlInput() any {
	var returns any
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ServiceToken() ZeroTrustAccessPolicyRequireServiceTokenOutputReference {
	var returns ZeroTrustAccessPolicyRequireServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ServiceTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewZeroTrustAccessPolicyRequireOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustAccessPolicyRequireOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessPolicyRequireOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicyRequireOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustAccessPolicyRequireOutputReference_Override(z ZeroTrustAccessPolicyRequireOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessPolicy.ZeroTrustAccessPolicyRequireOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutAnyValidServiceToken(value *ZeroTrustAccessPolicyRequireAnyValidServiceToken) {
	if err := z.validatePutAnyValidServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnyValidServiceToken",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutAuthContext(value *ZeroTrustAccessPolicyRequireAuthContext) {
	if err := z.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthContext",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutAuthMethod(value *ZeroTrustAccessPolicyRequireAuthMethod) {
	if err := z.validatePutAuthMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthMethod",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutAzureAd(value *ZeroTrustAccessPolicyRequireAzureAd) {
	if err := z.validatePutAzureAdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAzureAd",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutCertificate(value *ZeroTrustAccessPolicyRequireCertificate) {
	if err := z.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCertificate",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutCommonName(value *ZeroTrustAccessPolicyRequireCommonName) {
	if err := z.validatePutCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCommonName",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutDevicePosture(value *ZeroTrustAccessPolicyRequireDevicePosture) {
	if err := z.validatePutDevicePostureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDevicePosture",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutEmail(value *ZeroTrustAccessPolicyRequireEmail) {
	if err := z.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmail",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutEmailDomain(value *ZeroTrustAccessPolicyRequireEmailDomain) {
	if err := z.validatePutEmailDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailDomain",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutEmailList(value *ZeroTrustAccessPolicyRequireEmailListStruct) {
	if err := z.validatePutEmailListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailList",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutEveryone(value *ZeroTrustAccessPolicyRequireEveryone) {
	if err := z.validatePutEveryoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEveryone",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutExternalEvaluation(value *ZeroTrustAccessPolicyRequireExternalEvaluation) {
	if err := z.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExternalEvaluation",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutGeo(value *ZeroTrustAccessPolicyRequireGeo) {
	if err := z.validatePutGeoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGeo",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutGithubOrganization(value *ZeroTrustAccessPolicyRequireGithubOrganization) {
	if err := z.validatePutGithubOrganizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGithubOrganization",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutGroup(value *ZeroTrustAccessPolicyRequireGroup) {
	if err := z.validatePutGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGroup",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutGsuite(value *ZeroTrustAccessPolicyRequireGsuite) {
	if err := z.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGsuite",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutIp(value *ZeroTrustAccessPolicyRequireIp) {
	if err := z.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIp",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutIpList(value *ZeroTrustAccessPolicyRequireIpListStruct) {
	if err := z.validatePutIpListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpList",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutLoginMethod(value *ZeroTrustAccessPolicyRequireLoginMethod) {
	if err := z.validatePutLoginMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLoginMethod",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutOidc(value *ZeroTrustAccessPolicyRequireOidc) {
	if err := z.validatePutOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOidc",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutOkta(value *ZeroTrustAccessPolicyRequireOkta) {
	if err := z.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOkta",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutSaml(value *ZeroTrustAccessPolicyRequireSaml) {
	if err := z.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaml",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) PutServiceToken(value *ZeroTrustAccessPolicyRequireServiceToken) {
	if err := z.validatePutServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putServiceToken",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetAuthContext() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthContext",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetAzureAd() {
	_jsii_.InvokeVoid(
		z,
		"resetAzureAd",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		z,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		z,
		"resetEmail",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetEmailList() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		z,
		"resetEveryone",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		z,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		z,
		"resetGeo",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetGithubOrganization() {
	_jsii_.InvokeVoid(
		z,
		"resetGithubOrganization",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		z,
		"resetGroup",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		z,
		"resetGsuite",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		z,
		"resetIp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		z,
		"resetIpList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetLoginMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetOidc() {
	_jsii_.InvokeVoid(
		z,
		"resetOidc",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		z,
		"resetOkta",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		z,
		"resetSaml",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (z *jsiiProxy_ZeroTrustAccessPolicyRequireOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
