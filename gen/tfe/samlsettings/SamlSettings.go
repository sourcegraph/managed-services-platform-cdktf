package samlsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/tfe/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/tfe/samlsettings/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/saml_settings tfe_saml_settings}.
type SamlSettings interface {
	cdktf.TerraformResource
	AcsConsumerUrl() *string
	AttrGroups() *string
	SetAttrGroups(val *string)
	AttrGroupsInput() *string
	AttrSiteAdmin() *string
	SetAttrSiteAdmin(val *string)
	AttrSiteAdminInput() *string
	AttrUsername() *string
	SetAttrUsername(val *string)
	AttrUsernameInput() *string
	AuthnRequestsSigned() interface{}
	SetAuthnRequestsSigned(val interface{})
	AuthnRequestsSignedInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Debug() interface{}
	SetDebug(val interface{})
	DebugInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IdpCert() *string
	SetIdpCert(val *string)
	IdpCertInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MetadataUrl() *string
	// The tree node.
	Node() constructs.Node
	OldIdpCert() *string
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SignatureDigestMethod() *string
	SetSignatureDigestMethod(val *string)
	SignatureDigestMethodInput() *string
	SignatureSigningMethod() *string
	SetSignatureSigningMethod(val *string)
	SignatureSigningMethodInput() *string
	SiteAdminRole() *string
	SetSiteAdminRole(val *string)
	SiteAdminRoleInput() *string
	SloEndpointUrl() *string
	SetSloEndpointUrl(val *string)
	SloEndpointUrlInput() *string
	SsoApiTokenSessionTimeout() *float64
	SetSsoApiTokenSessionTimeout(val *float64)
	SsoApiTokenSessionTimeoutInput() *float64
	SsoEndpointUrl() *string
	SetSsoEndpointUrl(val *string)
	SsoEndpointUrlInput() *string
	TeamManagementEnabled() interface{}
	SetTeamManagementEnabled(val interface{})
	TeamManagementEnabledInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WantAssertionsSigned() interface{}
	SetWantAssertionsSigned(val interface{})
	WantAssertionsSignedInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAttrGroups()
	ResetAttrSiteAdmin()
	ResetAttrUsername()
	ResetAuthnRequestsSigned()
	ResetCertificate()
	ResetDebug()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateKey()
	ResetSignatureDigestMethod()
	ResetSignatureSigningMethod()
	ResetSiteAdminRole()
	ResetSsoApiTokenSessionTimeout()
	ResetTeamManagementEnabled()
	ResetWantAssertionsSigned()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SamlSettings
type jsiiProxy_SamlSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SamlSettings) AcsConsumerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acsConsumerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) AttrGroups() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) AttrGroupsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) AttrSiteAdmin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSiteAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) AttrSiteAdminInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSiteAdminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) AttrUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) AttrUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) AuthnRequestsSigned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authnRequestsSigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) AuthnRequestsSignedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authnRequestsSignedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) Debug() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) DebugInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) IdpCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) IdpCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) MetadataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) OldIdpCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldIdpCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SignatureDigestMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureDigestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SignatureDigestMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureDigestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SignatureSigningMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureSigningMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SignatureSigningMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureSigningMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SiteAdminRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteAdminRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SiteAdminRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteAdminRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SloEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SloEndpointUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloEndpointUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SsoApiTokenSessionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ssoApiTokenSessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SsoApiTokenSessionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ssoApiTokenSessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SsoEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) SsoEndpointUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoEndpointUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) TeamManagementEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"teamManagementEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) TeamManagementEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"teamManagementEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) WantAssertionsSigned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wantAssertionsSigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlSettings) WantAssertionsSignedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wantAssertionsSignedInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/saml_settings tfe_saml_settings} Resource.
func NewSamlSettings(scope constructs.Construct, id *string, config *SamlSettingsConfig) SamlSettings {
	_init_.Initialize()

	if err := validateNewSamlSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SamlSettings{}

	_jsii_.Create(
		"@cdktf/provider-tfe.samlSettings.SamlSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/saml_settings tfe_saml_settings} Resource.
func NewSamlSettings_Override(s SamlSettings, scope constructs.Construct, id *string, config *SamlSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tfe.samlSettings.SamlSettings",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SamlSettings)SetAttrGroups(val *string) {
	if err := j.validateSetAttrGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attrGroups",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetAttrSiteAdmin(val *string) {
	if err := j.validateSetAttrSiteAdminParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attrSiteAdmin",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetAttrUsername(val *string) {
	if err := j.validateSetAttrUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attrUsername",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetAuthnRequestsSigned(val interface{}) {
	if err := j.validateSetAuthnRequestsSignedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authnRequestsSigned",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetDebug(val interface{}) {
	if err := j.validateSetDebugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"debug",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetIdpCert(val *string) {
	if err := j.validateSetIdpCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpCert",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetPrivateKey(val *string) {
	if err := j.validateSetPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetSignatureDigestMethod(val *string) {
	if err := j.validateSetSignatureDigestMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureDigestMethod",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetSignatureSigningMethod(val *string) {
	if err := j.validateSetSignatureSigningMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureSigningMethod",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetSiteAdminRole(val *string) {
	if err := j.validateSetSiteAdminRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"siteAdminRole",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetSloEndpointUrl(val *string) {
	if err := j.validateSetSloEndpointUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sloEndpointUrl",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetSsoApiTokenSessionTimeout(val *float64) {
	if err := j.validateSetSsoApiTokenSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoApiTokenSessionTimeout",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetSsoEndpointUrl(val *string) {
	if err := j.validateSetSsoEndpointUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoEndpointUrl",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetTeamManagementEnabled(val interface{}) {
	if err := j.validateSetTeamManagementEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamManagementEnabled",
		val,
	)
}

func (j *jsiiProxy_SamlSettings)SetWantAssertionsSigned(val interface{}) {
	if err := j.validateSetWantAssertionsSignedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wantAssertionsSigned",
		val,
	)
}

// Generates CDKTF code for importing a SamlSettings resource upon running "cdktf plan <stack-name>".
func SamlSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSamlSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.samlSettings.SamlSettings",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SamlSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.samlSettings.SamlSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SamlSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.samlSettings.SamlSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SamlSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.samlSettings.SamlSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SamlSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-tfe.samlSettings.SamlSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SamlSettings) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SamlSettings) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SamlSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SamlSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SamlSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SamlSettings) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SamlSettings) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SamlSettings) ResetAttrGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetAttrGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetAttrSiteAdmin() {
	_jsii_.InvokeVoid(
		s,
		"resetAttrSiteAdmin",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetAttrUsername() {
	_jsii_.InvokeVoid(
		s,
		"resetAttrUsername",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetAuthnRequestsSigned() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthnRequestsSigned",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetCertificate() {
	_jsii_.InvokeVoid(
		s,
		"resetCertificate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetDebug() {
	_jsii_.InvokeVoid(
		s,
		"resetDebug",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetSignatureDigestMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetSignatureDigestMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetSignatureSigningMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetSignatureSigningMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetSiteAdminRole() {
	_jsii_.InvokeVoid(
		s,
		"resetSiteAdminRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetSsoApiTokenSessionTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetSsoApiTokenSessionTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetTeamManagementEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetTeamManagementEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) ResetWantAssertionsSigned() {
	_jsii_.InvokeVoid(
		s,
		"resetWantAssertionsSigned",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

