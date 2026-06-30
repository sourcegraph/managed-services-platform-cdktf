package zerotrustaccessgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/zerotrustaccessgroup/internal"
)

type ZeroTrustAccessGroupRequireOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() ZeroTrustAccessGroupRequireAnyValidServiceTokenOutputReference
	AnyValidServiceTokenInput() any
	AuthContext() ZeroTrustAccessGroupRequireAuthContextOutputReference
	AuthContextInput() any
	AuthMethod() ZeroTrustAccessGroupRequireAuthMethodOutputReference
	AuthMethodInput() any
	AzureAd() ZeroTrustAccessGroupRequireAzureAdOutputReference
	AzureAdInput() any
	Certificate() ZeroTrustAccessGroupRequireCertificateOutputReference
	CertificateInput() any
	CommonName() ZeroTrustAccessGroupRequireCommonNameOutputReference
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
	DevicePosture() ZeroTrustAccessGroupRequireDevicePostureOutputReference
	DevicePostureInput() any
	Email() ZeroTrustAccessGroupRequireEmailOutputReference
	EmailDomain() ZeroTrustAccessGroupRequireEmailDomainOutputReference
	EmailDomainInput() any
	EmailInput() any
	EmailList() ZeroTrustAccessGroupRequireEmailListStructOutputReference
	EmailListInput() any
	Everyone() ZeroTrustAccessGroupRequireEveryoneOutputReference
	EveryoneInput() any
	ExternalEvaluation() ZeroTrustAccessGroupRequireExternalEvaluationOutputReference
	ExternalEvaluationInput() any
	// Experimental.
	Fqn() *string
	Geo() ZeroTrustAccessGroupRequireGeoOutputReference
	GeoInput() any
	GithubOrganization() ZeroTrustAccessGroupRequireGithubOrganizationOutputReference
	GithubOrganizationInput() any
	Group() ZeroTrustAccessGroupRequireGroupOutputReference
	GroupInput() any
	Gsuite() ZeroTrustAccessGroupRequireGsuiteOutputReference
	GsuiteInput() any
	InternalValue() any
	SetInternalValue(val any)
	Ip() ZeroTrustAccessGroupRequireIpOutputReference
	IpInput() any
	IpList() ZeroTrustAccessGroupRequireIpListStructOutputReference
	IpListInput() any
	LoginMethod() ZeroTrustAccessGroupRequireLoginMethodOutputReference
	LoginMethodInput() any
	Oidc() ZeroTrustAccessGroupRequireOidcOutputReference
	OidcInput() any
	Okta() ZeroTrustAccessGroupRequireOktaOutputReference
	OktaInput() any
	Saml() ZeroTrustAccessGroupRequireSamlOutputReference
	SamlInput() any
	ServiceToken() ZeroTrustAccessGroupRequireServiceTokenOutputReference
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
	PutAnyValidServiceToken(value *ZeroTrustAccessGroupRequireAnyValidServiceToken)
	PutAuthContext(value *ZeroTrustAccessGroupRequireAuthContext)
	PutAuthMethod(value *ZeroTrustAccessGroupRequireAuthMethod)
	PutAzureAd(value *ZeroTrustAccessGroupRequireAzureAd)
	PutCertificate(value *ZeroTrustAccessGroupRequireCertificate)
	PutCommonName(value *ZeroTrustAccessGroupRequireCommonName)
	PutDevicePosture(value *ZeroTrustAccessGroupRequireDevicePosture)
	PutEmail(value *ZeroTrustAccessGroupRequireEmail)
	PutEmailDomain(value *ZeroTrustAccessGroupRequireEmailDomain)
	PutEmailList(value *ZeroTrustAccessGroupRequireEmailListStruct)
	PutEveryone(value *ZeroTrustAccessGroupRequireEveryone)
	PutExternalEvaluation(value *ZeroTrustAccessGroupRequireExternalEvaluation)
	PutGeo(value *ZeroTrustAccessGroupRequireGeo)
	PutGithubOrganization(value *ZeroTrustAccessGroupRequireGithubOrganization)
	PutGroup(value *ZeroTrustAccessGroupRequireGroup)
	PutGsuite(value *ZeroTrustAccessGroupRequireGsuite)
	PutIp(value *ZeroTrustAccessGroupRequireIp)
	PutIpList(value *ZeroTrustAccessGroupRequireIpListStruct)
	PutLoginMethod(value *ZeroTrustAccessGroupRequireLoginMethod)
	PutOidc(value *ZeroTrustAccessGroupRequireOidc)
	PutOkta(value *ZeroTrustAccessGroupRequireOkta)
	PutSaml(value *ZeroTrustAccessGroupRequireSaml)
	PutServiceToken(value *ZeroTrustAccessGroupRequireServiceToken)
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

// The jsii proxy struct for ZeroTrustAccessGroupRequireOutputReference
type jsiiProxy_ZeroTrustAccessGroupRequireOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AnyValidServiceToken() ZeroTrustAccessGroupRequireAnyValidServiceTokenOutputReference {
	var returns ZeroTrustAccessGroupRequireAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AnyValidServiceTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AuthContext() ZeroTrustAccessGroupRequireAuthContextOutputReference {
	var returns ZeroTrustAccessGroupRequireAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AuthContextInput() any {
	var returns any
	_jsii_.Get(
		j,
		"authContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AuthMethod() ZeroTrustAccessGroupRequireAuthMethodOutputReference {
	var returns ZeroTrustAccessGroupRequireAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AuthMethodInput() any {
	var returns any
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AzureAd() ZeroTrustAccessGroupRequireAzureAdOutputReference {
	var returns ZeroTrustAccessGroupRequireAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) AzureAdInput() any {
	var returns any
	_jsii_.Get(
		j,
		"azureAdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Certificate() ZeroTrustAccessGroupRequireCertificateOutputReference {
	var returns ZeroTrustAccessGroupRequireCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) CertificateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) CommonName() ZeroTrustAccessGroupRequireCommonNameOutputReference {
	var returns ZeroTrustAccessGroupRequireCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) CommonNameInput() any {
	var returns any
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) DevicePosture() ZeroTrustAccessGroupRequireDevicePostureOutputReference {
	var returns ZeroTrustAccessGroupRequireDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) DevicePostureInput() any {
	var returns any
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Email() ZeroTrustAccessGroupRequireEmailOutputReference {
	var returns ZeroTrustAccessGroupRequireEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EmailDomain() ZeroTrustAccessGroupRequireEmailDomainOutputReference {
	var returns ZeroTrustAccessGroupRequireEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EmailDomainInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EmailInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EmailList() ZeroTrustAccessGroupRequireEmailListStructOutputReference {
	var returns ZeroTrustAccessGroupRequireEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EmailListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Everyone() ZeroTrustAccessGroupRequireEveryoneOutputReference {
	var returns ZeroTrustAccessGroupRequireEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) EveryoneInput() any {
	var returns any
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ExternalEvaluation() ZeroTrustAccessGroupRequireExternalEvaluationOutputReference {
	var returns ZeroTrustAccessGroupRequireExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ExternalEvaluationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Geo() ZeroTrustAccessGroupRequireGeoOutputReference {
	var returns ZeroTrustAccessGroupRequireGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GeoInput() any {
	var returns any
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GithubOrganization() ZeroTrustAccessGroupRequireGithubOrganizationOutputReference {
	var returns ZeroTrustAccessGroupRequireGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GithubOrganizationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"githubOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Group() ZeroTrustAccessGroupRequireGroupOutputReference {
	var returns ZeroTrustAccessGroupRequireGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GroupInput() any {
	var returns any
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Gsuite() ZeroTrustAccessGroupRequireGsuiteOutputReference {
	var returns ZeroTrustAccessGroupRequireGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GsuiteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Ip() ZeroTrustAccessGroupRequireIpOutputReference {
	var returns ZeroTrustAccessGroupRequireIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) IpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) IpList() ZeroTrustAccessGroupRequireIpListStructOutputReference {
	var returns ZeroTrustAccessGroupRequireIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) IpListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) LoginMethod() ZeroTrustAccessGroupRequireLoginMethodOutputReference {
	var returns ZeroTrustAccessGroupRequireLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) LoginMethodInput() any {
	var returns any
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Oidc() ZeroTrustAccessGroupRequireOidcOutputReference {
	var returns ZeroTrustAccessGroupRequireOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) OidcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"oidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Okta() ZeroTrustAccessGroupRequireOktaOutputReference {
	var returns ZeroTrustAccessGroupRequireOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) OktaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Saml() ZeroTrustAccessGroupRequireSamlOutputReference {
	var returns ZeroTrustAccessGroupRequireSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) SamlInput() any {
	var returns any
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ServiceToken() ZeroTrustAccessGroupRequireServiceTokenOutputReference {
	var returns ZeroTrustAccessGroupRequireServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ServiceTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewZeroTrustAccessGroupRequireOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustAccessGroupRequireOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessGroupRequireOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessGroupRequireOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessGroup.ZeroTrustAccessGroupRequireOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustAccessGroupRequireOutputReference_Override(z ZeroTrustAccessGroupRequireOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessGroup.ZeroTrustAccessGroupRequireOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutAnyValidServiceToken(value *ZeroTrustAccessGroupRequireAnyValidServiceToken) {
	if err := z.validatePutAnyValidServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnyValidServiceToken",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutAuthContext(value *ZeroTrustAccessGroupRequireAuthContext) {
	if err := z.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthContext",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutAuthMethod(value *ZeroTrustAccessGroupRequireAuthMethod) {
	if err := z.validatePutAuthMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthMethod",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutAzureAd(value *ZeroTrustAccessGroupRequireAzureAd) {
	if err := z.validatePutAzureAdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAzureAd",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutCertificate(value *ZeroTrustAccessGroupRequireCertificate) {
	if err := z.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCertificate",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutCommonName(value *ZeroTrustAccessGroupRequireCommonName) {
	if err := z.validatePutCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCommonName",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutDevicePosture(value *ZeroTrustAccessGroupRequireDevicePosture) {
	if err := z.validatePutDevicePostureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDevicePosture",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutEmail(value *ZeroTrustAccessGroupRequireEmail) {
	if err := z.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmail",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutEmailDomain(value *ZeroTrustAccessGroupRequireEmailDomain) {
	if err := z.validatePutEmailDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailDomain",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutEmailList(value *ZeroTrustAccessGroupRequireEmailListStruct) {
	if err := z.validatePutEmailListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailList",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutEveryone(value *ZeroTrustAccessGroupRequireEveryone) {
	if err := z.validatePutEveryoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEveryone",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutExternalEvaluation(value *ZeroTrustAccessGroupRequireExternalEvaluation) {
	if err := z.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExternalEvaluation",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutGeo(value *ZeroTrustAccessGroupRequireGeo) {
	if err := z.validatePutGeoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGeo",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutGithubOrganization(value *ZeroTrustAccessGroupRequireGithubOrganization) {
	if err := z.validatePutGithubOrganizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGithubOrganization",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutGroup(value *ZeroTrustAccessGroupRequireGroup) {
	if err := z.validatePutGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGroup",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutGsuite(value *ZeroTrustAccessGroupRequireGsuite) {
	if err := z.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGsuite",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutIp(value *ZeroTrustAccessGroupRequireIp) {
	if err := z.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIp",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutIpList(value *ZeroTrustAccessGroupRequireIpListStruct) {
	if err := z.validatePutIpListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpList",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutLoginMethod(value *ZeroTrustAccessGroupRequireLoginMethod) {
	if err := z.validatePutLoginMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLoginMethod",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutOidc(value *ZeroTrustAccessGroupRequireOidc) {
	if err := z.validatePutOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOidc",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutOkta(value *ZeroTrustAccessGroupRequireOkta) {
	if err := z.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOkta",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutSaml(value *ZeroTrustAccessGroupRequireSaml) {
	if err := z.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaml",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) PutServiceToken(value *ZeroTrustAccessGroupRequireServiceToken) {
	if err := z.validatePutServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putServiceToken",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetAuthContext() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthContext",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetAzureAd() {
	_jsii_.InvokeVoid(
		z,
		"resetAzureAd",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		z,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		z,
		"resetEmail",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetEmailList() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		z,
		"resetEveryone",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		z,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		z,
		"resetGeo",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetGithubOrganization() {
	_jsii_.InvokeVoid(
		z,
		"resetGithubOrganization",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		z,
		"resetGroup",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		z,
		"resetGsuite",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		z,
		"resetIp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		z,
		"resetIpList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetLoginMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetOidc() {
	_jsii_.InvokeVoid(
		z,
		"resetOidc",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		z,
		"resetOkta",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		z,
		"resetSaml",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (z *jsiiProxy_ZeroTrustAccessGroupRequireOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
