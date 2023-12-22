package datatfesamlsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/tfe/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/tfe/datatfesamlsettings/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/data-sources/saml_settings tfe_saml_settings}.
type DataTfeSamlSettings interface {
	cdktf.TerraformDataSource
	AcsConsumerUrl() *string
	AttrGroups() *string
	AttrSiteAdmin() *string
	AttrUsername() *string
	AuthnRequestsSigned() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Debug() cdktf.IResolvable
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MetadataUrl() *string
	// The tree node.
	Node() constructs.Node
	OldIdpCert() *string
	PrivateKey() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SignatureDigestMethod() *string
	SignatureSigningMethod() *string
	SiteAdminRole() *string
	SloEndpointUrl() *string
	SsoApiTokenSessionTimeout() *float64
	SsoEndpointUrl() *string
	TeamManagementEnabled() cdktf.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WantAssertionsSigned() cdktf.IResolvable
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTfeSamlSettings
type jsiiProxy_DataTfeSamlSettings struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataTfeSamlSettings) AcsConsumerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acsConsumerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) AttrGroups() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) AttrSiteAdmin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSiteAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) AttrUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) AuthnRequestsSigned() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"authnRequestsSigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) Debug() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"debug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) IdpCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) MetadataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) OldIdpCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oldIdpCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) SignatureDigestMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureDigestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) SignatureSigningMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureSigningMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) SiteAdminRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteAdminRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) SloEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) SsoApiTokenSessionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ssoApiTokenSessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) SsoEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) TeamManagementEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"teamManagementEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeSamlSettings) WantAssertionsSigned() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"wantAssertionsSigned",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/data-sources/saml_settings tfe_saml_settings} Data Source.
func NewDataTfeSamlSettings(scope constructs.Construct, id *string, config *DataTfeSamlSettingsConfig) DataTfeSamlSettings {
	_init_.Initialize()

	if err := validateNewDataTfeSamlSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfeSamlSettings{}

	_jsii_.Create(
		"@cdktf/provider-tfe.dataTfeSamlSettings.DataTfeSamlSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/data-sources/saml_settings tfe_saml_settings} Data Source.
func NewDataTfeSamlSettings_Override(d DataTfeSamlSettings, scope constructs.Construct, id *string, config *DataTfeSamlSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tfe.dataTfeSamlSettings.DataTfeSamlSettings",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataTfeSamlSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataTfeSamlSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataTfeSamlSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataTfeSamlSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataTfeSamlSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
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
func DataTfeSamlSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfeSamlSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.dataTfeSamlSettings.DataTfeSamlSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfeSamlSettings_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfeSamlSettings_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.dataTfeSamlSettings.DataTfeSamlSettings",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfeSamlSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfeSamlSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.dataTfeSamlSettings.DataTfeSamlSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTfeSamlSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-tfe.dataTfeSamlSettings.DataTfeSamlSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTfeSamlSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTfeSamlSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfeSamlSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeSamlSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

