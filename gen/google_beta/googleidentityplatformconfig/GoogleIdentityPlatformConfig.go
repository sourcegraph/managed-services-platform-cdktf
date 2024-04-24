package googleidentityplatformconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleidentityplatformconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_identity_platform_config google_identity_platform_config}.
type GoogleIdentityPlatformConfig interface {
	cdktf.TerraformResource
	AuthorizedDomains() *[]*string
	SetAuthorizedDomains(val *[]*string)
	AuthorizedDomainsInput() *[]*string
	AutodeleteAnonymousUsers() interface{}
	SetAutodeleteAnonymousUsers(val interface{})
	AutodeleteAnonymousUsersInput() interface{}
	BlockingFunctions() GoogleIdentityPlatformConfigBlockingFunctionsOutputReference
	BlockingFunctionsInput() *GoogleIdentityPlatformConfigBlockingFunctions
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Client() GoogleIdentityPlatformConfigClientOutputReference
	ClientInput() *GoogleIdentityPlatformConfigClient
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Mfa() GoogleIdentityPlatformConfigMfaOutputReference
	MfaInput() *GoogleIdentityPlatformConfigMfa
	Monitoring() GoogleIdentityPlatformConfigMonitoringOutputReference
	MonitoringInput() *GoogleIdentityPlatformConfigMonitoring
	MultiTenant() GoogleIdentityPlatformConfigMultiTenantOutputReference
	MultiTenantInput() *GoogleIdentityPlatformConfigMultiTenant
	Name() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Quota() GoogleIdentityPlatformConfigQuotaOutputReference
	QuotaInput() *GoogleIdentityPlatformConfigQuota
	// Experimental.
	RawOverrides() interface{}
	SignIn() GoogleIdentityPlatformConfigSignInOutputReference
	SignInInput() *GoogleIdentityPlatformConfigSignIn
	SmsRegionConfig() GoogleIdentityPlatformConfigSmsRegionConfigOutputReference
	SmsRegionConfigInput() *GoogleIdentityPlatformConfigSmsRegionConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleIdentityPlatformConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutBlockingFunctions(value *GoogleIdentityPlatformConfigBlockingFunctions)
	PutClient(value *GoogleIdentityPlatformConfigClient)
	PutMfa(value *GoogleIdentityPlatformConfigMfa)
	PutMonitoring(value *GoogleIdentityPlatformConfigMonitoring)
	PutMultiTenant(value *GoogleIdentityPlatformConfigMultiTenant)
	PutQuota(value *GoogleIdentityPlatformConfigQuota)
	PutSignIn(value *GoogleIdentityPlatformConfigSignIn)
	PutSmsRegionConfig(value *GoogleIdentityPlatformConfigSmsRegionConfig)
	PutTimeouts(value *GoogleIdentityPlatformConfigTimeouts)
	ResetAuthorizedDomains()
	ResetAutodeleteAnonymousUsers()
	ResetBlockingFunctions()
	ResetClient()
	ResetId()
	ResetMfa()
	ResetMonitoring()
	ResetMultiTenant()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetQuota()
	ResetSignIn()
	ResetSmsRegionConfig()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleIdentityPlatformConfig
type jsiiProxy_GoogleIdentityPlatformConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) AuthorizedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) AuthorizedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) AutodeleteAnonymousUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autodeleteAnonymousUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) AutodeleteAnonymousUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autodeleteAnonymousUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) BlockingFunctions() GoogleIdentityPlatformConfigBlockingFunctionsOutputReference {
	var returns GoogleIdentityPlatformConfigBlockingFunctionsOutputReference
	_jsii_.Get(
		j,
		"blockingFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) BlockingFunctionsInput() *GoogleIdentityPlatformConfigBlockingFunctions {
	var returns *GoogleIdentityPlatformConfigBlockingFunctions
	_jsii_.Get(
		j,
		"blockingFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Client() GoogleIdentityPlatformConfigClientOutputReference {
	var returns GoogleIdentityPlatformConfigClientOutputReference
	_jsii_.Get(
		j,
		"client",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) ClientInput() *GoogleIdentityPlatformConfigClient {
	var returns *GoogleIdentityPlatformConfigClient
	_jsii_.Get(
		j,
		"clientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Mfa() GoogleIdentityPlatformConfigMfaOutputReference {
	var returns GoogleIdentityPlatformConfigMfaOutputReference
	_jsii_.Get(
		j,
		"mfa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) MfaInput() *GoogleIdentityPlatformConfigMfa {
	var returns *GoogleIdentityPlatformConfigMfa
	_jsii_.Get(
		j,
		"mfaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Monitoring() GoogleIdentityPlatformConfigMonitoringOutputReference {
	var returns GoogleIdentityPlatformConfigMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) MonitoringInput() *GoogleIdentityPlatformConfigMonitoring {
	var returns *GoogleIdentityPlatformConfigMonitoring
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) MultiTenant() GoogleIdentityPlatformConfigMultiTenantOutputReference {
	var returns GoogleIdentityPlatformConfigMultiTenantOutputReference
	_jsii_.Get(
		j,
		"multiTenant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) MultiTenantInput() *GoogleIdentityPlatformConfigMultiTenant {
	var returns *GoogleIdentityPlatformConfigMultiTenant
	_jsii_.Get(
		j,
		"multiTenantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Quota() GoogleIdentityPlatformConfigQuotaOutputReference {
	var returns GoogleIdentityPlatformConfigQuotaOutputReference
	_jsii_.Get(
		j,
		"quota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) QuotaInput() *GoogleIdentityPlatformConfigQuota {
	var returns *GoogleIdentityPlatformConfigQuota
	_jsii_.Get(
		j,
		"quotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) SignIn() GoogleIdentityPlatformConfigSignInOutputReference {
	var returns GoogleIdentityPlatformConfigSignInOutputReference
	_jsii_.Get(
		j,
		"signIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) SignInInput() *GoogleIdentityPlatformConfigSignIn {
	var returns *GoogleIdentityPlatformConfigSignIn
	_jsii_.Get(
		j,
		"signInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) SmsRegionConfig() GoogleIdentityPlatformConfigSmsRegionConfigOutputReference {
	var returns GoogleIdentityPlatformConfigSmsRegionConfigOutputReference
	_jsii_.Get(
		j,
		"smsRegionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) SmsRegionConfigInput() *GoogleIdentityPlatformConfigSmsRegionConfig {
	var returns *GoogleIdentityPlatformConfigSmsRegionConfig
	_jsii_.Get(
		j,
		"smsRegionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) Timeouts() GoogleIdentityPlatformConfigTimeoutsOutputReference {
	var returns GoogleIdentityPlatformConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_identity_platform_config google_identity_platform_config} Resource.
func NewGoogleIdentityPlatformConfig(scope constructs.Construct, id *string, config *GoogleIdentityPlatformConfigConfig) GoogleIdentityPlatformConfig {
	_init_.Initialize()

	if err := validateNewGoogleIdentityPlatformConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIdentityPlatformConfig{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_identity_platform_config google_identity_platform_config} Resource.
func NewGoogleIdentityPlatformConfig_Override(g GoogleIdentityPlatformConfig, scope constructs.Construct, id *string, config *GoogleIdentityPlatformConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfig",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig)SetAuthorizedDomains(val *[]*string) {
	if err := j.validateSetAuthorizedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedDomains",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig)SetAutodeleteAnonymousUsers(val interface{}) {
	if err := j.validateSetAutodeleteAnonymousUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodeleteAnonymousUsers",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleIdentityPlatformConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func GoogleIdentityPlatformConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIdentityPlatformConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleIdentityPlatformConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIdentityPlatformConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleIdentityPlatformConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIdentityPlatformConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleIdentityPlatformConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleIdentityPlatformConfig.GoogleIdentityPlatformConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIdentityPlatformConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) PutBlockingFunctions(value *GoogleIdentityPlatformConfigBlockingFunctions) {
	if err := g.validatePutBlockingFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBlockingFunctions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) PutClient(value *GoogleIdentityPlatformConfigClient) {
	if err := g.validatePutClientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClient",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) PutMfa(value *GoogleIdentityPlatformConfigMfa) {
	if err := g.validatePutMfaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMfa",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) PutMonitoring(value *GoogleIdentityPlatformConfigMonitoring) {
	if err := g.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMonitoring",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) PutMultiTenant(value *GoogleIdentityPlatformConfigMultiTenant) {
	if err := g.validatePutMultiTenantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMultiTenant",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) PutQuota(value *GoogleIdentityPlatformConfigQuota) {
	if err := g.validatePutQuotaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQuota",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) PutSignIn(value *GoogleIdentityPlatformConfigSignIn) {
	if err := g.validatePutSignInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSignIn",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) PutSmsRegionConfig(value *GoogleIdentityPlatformConfigSmsRegionConfig) {
	if err := g.validatePutSmsRegionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSmsRegionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) PutTimeouts(value *GoogleIdentityPlatformConfigTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetAuthorizedDomains() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizedDomains",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetAutodeleteAnonymousUsers() {
	_jsii_.InvokeVoid(
		g,
		"resetAutodeleteAnonymousUsers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetBlockingFunctions() {
	_jsii_.InvokeVoid(
		g,
		"resetBlockingFunctions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetClient() {
	_jsii_.InvokeVoid(
		g,
		"resetClient",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetMfa() {
	_jsii_.InvokeVoid(
		g,
		"resetMfa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetMonitoring() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetMultiTenant() {
	_jsii_.InvokeVoid(
		g,
		"resetMultiTenant",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetQuota() {
	_jsii_.InvokeVoid(
		g,
		"resetQuota",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetSignIn() {
	_jsii_.InvokeVoid(
		g,
		"resetSignIn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetSmsRegionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSmsRegionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIdentityPlatformConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

