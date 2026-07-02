package zerotrustaccessgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/zerotrustaccessgroup/internal"
)

type ZeroTrustAccessGroupIncludeOutputReference interface {
	cdktf.ComplexObject
	AnyValidServiceToken() ZeroTrustAccessGroupIncludeAnyValidServiceTokenOutputReference
	AnyValidServiceTokenInput() any
	AuthContext() ZeroTrustAccessGroupIncludeAuthContextOutputReference
	AuthContextInput() any
	AuthMethod() ZeroTrustAccessGroupIncludeAuthMethodOutputReference
	AuthMethodInput() any
	AzureAd() ZeroTrustAccessGroupIncludeAzureAdOutputReference
	AzureAdInput() any
	Certificate() ZeroTrustAccessGroupIncludeCertificateOutputReference
	CertificateInput() any
	CommonName() ZeroTrustAccessGroupIncludeCommonNameOutputReference
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
	DevicePosture() ZeroTrustAccessGroupIncludeDevicePostureOutputReference
	DevicePostureInput() any
	Email() ZeroTrustAccessGroupIncludeEmailOutputReference
	EmailDomain() ZeroTrustAccessGroupIncludeEmailDomainOutputReference
	EmailDomainInput() any
	EmailInput() any
	EmailList() ZeroTrustAccessGroupIncludeEmailListStructOutputReference
	EmailListInput() any
	Everyone() ZeroTrustAccessGroupIncludeEveryoneOutputReference
	EveryoneInput() any
	ExternalEvaluation() ZeroTrustAccessGroupIncludeExternalEvaluationOutputReference
	ExternalEvaluationInput() any
	// Experimental.
	Fqn() *string
	Geo() ZeroTrustAccessGroupIncludeGeoOutputReference
	GeoInput() any
	GithubOrganization() ZeroTrustAccessGroupIncludeGithubOrganizationOutputReference
	GithubOrganizationInput() any
	Group() ZeroTrustAccessGroupIncludeGroupOutputReference
	GroupInput() any
	Gsuite() ZeroTrustAccessGroupIncludeGsuiteOutputReference
	GsuiteInput() any
	InternalValue() any
	SetInternalValue(val any)
	Ip() ZeroTrustAccessGroupIncludeIpOutputReference
	IpInput() any
	IpList() ZeroTrustAccessGroupIncludeIpListStructOutputReference
	IpListInput() any
	LoginMethod() ZeroTrustAccessGroupIncludeLoginMethodOutputReference
	LoginMethodInput() any
	Oidc() ZeroTrustAccessGroupIncludeOidcOutputReference
	OidcInput() any
	Okta() ZeroTrustAccessGroupIncludeOktaOutputReference
	OktaInput() any
	Saml() ZeroTrustAccessGroupIncludeSamlOutputReference
	SamlInput() any
	ServiceToken() ZeroTrustAccessGroupIncludeServiceTokenOutputReference
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
	PutAnyValidServiceToken(value *ZeroTrustAccessGroupIncludeAnyValidServiceToken)
	PutAuthContext(value *ZeroTrustAccessGroupIncludeAuthContext)
	PutAuthMethod(value *ZeroTrustAccessGroupIncludeAuthMethod)
	PutAzureAd(value *ZeroTrustAccessGroupIncludeAzureAd)
	PutCertificate(value *ZeroTrustAccessGroupIncludeCertificate)
	PutCommonName(value *ZeroTrustAccessGroupIncludeCommonName)
	PutDevicePosture(value *ZeroTrustAccessGroupIncludeDevicePosture)
	PutEmail(value *ZeroTrustAccessGroupIncludeEmail)
	PutEmailDomain(value *ZeroTrustAccessGroupIncludeEmailDomain)
	PutEmailList(value *ZeroTrustAccessGroupIncludeEmailListStruct)
	PutEveryone(value *ZeroTrustAccessGroupIncludeEveryone)
	PutExternalEvaluation(value *ZeroTrustAccessGroupIncludeExternalEvaluation)
	PutGeo(value *ZeroTrustAccessGroupIncludeGeo)
	PutGithubOrganization(value *ZeroTrustAccessGroupIncludeGithubOrganization)
	PutGroup(value *ZeroTrustAccessGroupIncludeGroup)
	PutGsuite(value *ZeroTrustAccessGroupIncludeGsuite)
	PutIp(value *ZeroTrustAccessGroupIncludeIp)
	PutIpList(value *ZeroTrustAccessGroupIncludeIpListStruct)
	PutLoginMethod(value *ZeroTrustAccessGroupIncludeLoginMethod)
	PutOidc(value *ZeroTrustAccessGroupIncludeOidc)
	PutOkta(value *ZeroTrustAccessGroupIncludeOkta)
	PutSaml(value *ZeroTrustAccessGroupIncludeSaml)
	PutServiceToken(value *ZeroTrustAccessGroupIncludeServiceToken)
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

// The jsii proxy struct for ZeroTrustAccessGroupIncludeOutputReference
type jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) AnyValidServiceToken() ZeroTrustAccessGroupIncludeAnyValidServiceTokenOutputReference {
	var returns ZeroTrustAccessGroupIncludeAnyValidServiceTokenOutputReference
	_jsii_.Get(
		j,
		"anyValidServiceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) AnyValidServiceTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"anyValidServiceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) AuthContext() ZeroTrustAccessGroupIncludeAuthContextOutputReference {
	var returns ZeroTrustAccessGroupIncludeAuthContextOutputReference
	_jsii_.Get(
		j,
		"authContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) AuthContextInput() any {
	var returns any
	_jsii_.Get(
		j,
		"authContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) AuthMethod() ZeroTrustAccessGroupIncludeAuthMethodOutputReference {
	var returns ZeroTrustAccessGroupIncludeAuthMethodOutputReference
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) AuthMethodInput() any {
	var returns any
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) AzureAd() ZeroTrustAccessGroupIncludeAzureAdOutputReference {
	var returns ZeroTrustAccessGroupIncludeAzureAdOutputReference
	_jsii_.Get(
		j,
		"azureAd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) AzureAdInput() any {
	var returns any
	_jsii_.Get(
		j,
		"azureAdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Certificate() ZeroTrustAccessGroupIncludeCertificateOutputReference {
	var returns ZeroTrustAccessGroupIncludeCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) CertificateInput() any {
	var returns any
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) CommonName() ZeroTrustAccessGroupIncludeCommonNameOutputReference {
	var returns ZeroTrustAccessGroupIncludeCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) CommonNameInput() any {
	var returns any
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) DevicePosture() ZeroTrustAccessGroupIncludeDevicePostureOutputReference {
	var returns ZeroTrustAccessGroupIncludeDevicePostureOutputReference
	_jsii_.Get(
		j,
		"devicePosture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) DevicePostureInput() any {
	var returns any
	_jsii_.Get(
		j,
		"devicePostureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Email() ZeroTrustAccessGroupIncludeEmailOutputReference {
	var returns ZeroTrustAccessGroupIncludeEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) EmailDomain() ZeroTrustAccessGroupIncludeEmailDomainOutputReference {
	var returns ZeroTrustAccessGroupIncludeEmailDomainOutputReference
	_jsii_.Get(
		j,
		"emailDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) EmailDomainInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) EmailInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) EmailList() ZeroTrustAccessGroupIncludeEmailListStructOutputReference {
	var returns ZeroTrustAccessGroupIncludeEmailListStructOutputReference
	_jsii_.Get(
		j,
		"emailList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) EmailListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"emailListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Everyone() ZeroTrustAccessGroupIncludeEveryoneOutputReference {
	var returns ZeroTrustAccessGroupIncludeEveryoneOutputReference
	_jsii_.Get(
		j,
		"everyone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) EveryoneInput() any {
	var returns any
	_jsii_.Get(
		j,
		"everyoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ExternalEvaluation() ZeroTrustAccessGroupIncludeExternalEvaluationOutputReference {
	var returns ZeroTrustAccessGroupIncludeExternalEvaluationOutputReference
	_jsii_.Get(
		j,
		"externalEvaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ExternalEvaluationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"externalEvaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Geo() ZeroTrustAccessGroupIncludeGeoOutputReference {
	var returns ZeroTrustAccessGroupIncludeGeoOutputReference
	_jsii_.Get(
		j,
		"geo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GeoInput() any {
	var returns any
	_jsii_.Get(
		j,
		"geoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GithubOrganization() ZeroTrustAccessGroupIncludeGithubOrganizationOutputReference {
	var returns ZeroTrustAccessGroupIncludeGithubOrganizationOutputReference
	_jsii_.Get(
		j,
		"githubOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GithubOrganizationInput() any {
	var returns any
	_jsii_.Get(
		j,
		"githubOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Group() ZeroTrustAccessGroupIncludeGroupOutputReference {
	var returns ZeroTrustAccessGroupIncludeGroupOutputReference
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GroupInput() any {
	var returns any
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Gsuite() ZeroTrustAccessGroupIncludeGsuiteOutputReference {
	var returns ZeroTrustAccessGroupIncludeGsuiteOutputReference
	_jsii_.Get(
		j,
		"gsuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GsuiteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"gsuiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Ip() ZeroTrustAccessGroupIncludeIpOutputReference {
	var returns ZeroTrustAccessGroupIncludeIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) IpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) IpList() ZeroTrustAccessGroupIncludeIpListStructOutputReference {
	var returns ZeroTrustAccessGroupIncludeIpListStructOutputReference
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) IpListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) LoginMethod() ZeroTrustAccessGroupIncludeLoginMethodOutputReference {
	var returns ZeroTrustAccessGroupIncludeLoginMethodOutputReference
	_jsii_.Get(
		j,
		"loginMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) LoginMethodInput() any {
	var returns any
	_jsii_.Get(
		j,
		"loginMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Oidc() ZeroTrustAccessGroupIncludeOidcOutputReference {
	var returns ZeroTrustAccessGroupIncludeOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) OidcInput() any {
	var returns any
	_jsii_.Get(
		j,
		"oidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Okta() ZeroTrustAccessGroupIncludeOktaOutputReference {
	var returns ZeroTrustAccessGroupIncludeOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) OktaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Saml() ZeroTrustAccessGroupIncludeSamlOutputReference {
	var returns ZeroTrustAccessGroupIncludeSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) SamlInput() any {
	var returns any
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ServiceToken() ZeroTrustAccessGroupIncludeServiceTokenOutputReference {
	var returns ZeroTrustAccessGroupIncludeServiceTokenOutputReference
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ServiceTokenInput() any {
	var returns any
	_jsii_.Get(
		j,
		"serviceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewZeroTrustAccessGroupIncludeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustAccessGroupIncludeOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessGroupIncludeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessGroup.ZeroTrustAccessGroupIncludeOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustAccessGroupIncludeOutputReference_Override(z ZeroTrustAccessGroupIncludeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessGroup.ZeroTrustAccessGroupIncludeOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutAnyValidServiceToken(value *ZeroTrustAccessGroupIncludeAnyValidServiceToken) {
	if err := z.validatePutAnyValidServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAnyValidServiceToken",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutAuthContext(value *ZeroTrustAccessGroupIncludeAuthContext) {
	if err := z.validatePutAuthContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthContext",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutAuthMethod(value *ZeroTrustAccessGroupIncludeAuthMethod) {
	if err := z.validatePutAuthMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAuthMethod",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutAzureAd(value *ZeroTrustAccessGroupIncludeAzureAd) {
	if err := z.validatePutAzureAdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putAzureAd",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutCertificate(value *ZeroTrustAccessGroupIncludeCertificate) {
	if err := z.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCertificate",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutCommonName(value *ZeroTrustAccessGroupIncludeCommonName) {
	if err := z.validatePutCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCommonName",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutDevicePosture(value *ZeroTrustAccessGroupIncludeDevicePosture) {
	if err := z.validatePutDevicePostureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDevicePosture",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutEmail(value *ZeroTrustAccessGroupIncludeEmail) {
	if err := z.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmail",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutEmailDomain(value *ZeroTrustAccessGroupIncludeEmailDomain) {
	if err := z.validatePutEmailDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailDomain",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutEmailList(value *ZeroTrustAccessGroupIncludeEmailListStruct) {
	if err := z.validatePutEmailListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEmailList",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutEveryone(value *ZeroTrustAccessGroupIncludeEveryone) {
	if err := z.validatePutEveryoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEveryone",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutExternalEvaluation(value *ZeroTrustAccessGroupIncludeExternalEvaluation) {
	if err := z.validatePutExternalEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExternalEvaluation",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutGeo(value *ZeroTrustAccessGroupIncludeGeo) {
	if err := z.validatePutGeoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGeo",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutGithubOrganization(value *ZeroTrustAccessGroupIncludeGithubOrganization) {
	if err := z.validatePutGithubOrganizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGithubOrganization",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutGroup(value *ZeroTrustAccessGroupIncludeGroup) {
	if err := z.validatePutGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGroup",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutGsuite(value *ZeroTrustAccessGroupIncludeGsuite) {
	if err := z.validatePutGsuiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putGsuite",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutIp(value *ZeroTrustAccessGroupIncludeIp) {
	if err := z.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIp",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutIpList(value *ZeroTrustAccessGroupIncludeIpListStruct) {
	if err := z.validatePutIpListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putIpList",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutLoginMethod(value *ZeroTrustAccessGroupIncludeLoginMethod) {
	if err := z.validatePutLoginMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLoginMethod",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutOidc(value *ZeroTrustAccessGroupIncludeOidc) {
	if err := z.validatePutOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOidc",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutOkta(value *ZeroTrustAccessGroupIncludeOkta) {
	if err := z.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putOkta",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutSaml(value *ZeroTrustAccessGroupIncludeSaml) {
	if err := z.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaml",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) PutServiceToken(value *ZeroTrustAccessGroupIncludeServiceToken) {
	if err := z.validatePutServiceTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putServiceToken",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetAnyValidServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetAnyValidServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetAuthContext() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthContext",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetAzureAd() {
	_jsii_.InvokeVoid(
		z,
		"resetAzureAd",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificate",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		z,
		"resetCommonName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetDevicePosture() {
	_jsii_.InvokeVoid(
		z,
		"resetDevicePosture",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		z,
		"resetEmail",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetEmailDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetEmailList() {
	_jsii_.InvokeVoid(
		z,
		"resetEmailList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetEveryone() {
	_jsii_.InvokeVoid(
		z,
		"resetEveryone",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetExternalEvaluation() {
	_jsii_.InvokeVoid(
		z,
		"resetExternalEvaluation",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetGeo() {
	_jsii_.InvokeVoid(
		z,
		"resetGeo",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetGithubOrganization() {
	_jsii_.InvokeVoid(
		z,
		"resetGithubOrganization",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		z,
		"resetGroup",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetGsuite() {
	_jsii_.InvokeVoid(
		z,
		"resetGsuite",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		z,
		"resetIp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetIpList() {
	_jsii_.InvokeVoid(
		z,
		"resetIpList",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetLoginMethod() {
	_jsii_.InvokeVoid(
		z,
		"resetLoginMethod",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetOidc() {
	_jsii_.InvokeVoid(
		z,
		"resetOidc",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetOkta() {
	_jsii_.InvokeVoid(
		z,
		"resetOkta",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		z,
		"resetSaml",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ResetServiceToken() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceToken",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (z *jsiiProxy_ZeroTrustAccessGroupIncludeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
