package googlesecurityscannerscanconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlesecurityscannerscanconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_security_scanner_scan_config google_security_scanner_scan_config}.
type GoogleSecurityScannerScanConfig interface {
	cdktf.TerraformResource
	Authentication() GoogleSecurityScannerScanConfigAuthenticationOutputReference
	AuthenticationInput() *GoogleSecurityScannerScanConfigAuthentication
	BlacklistPatterns() *[]*string
	SetBlacklistPatterns(val *[]*string)
	BlacklistPatternsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ExportToSecurityCommandCenter() *string
	SetExportToSecurityCommandCenter(val *string)
	ExportToSecurityCommandCenterInput() *string
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
	MaxQps() *float64
	SetMaxQps(val *float64)
	MaxQpsInput() *float64
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
	// Experimental.
	RawOverrides() interface{}
	Schedule() GoogleSecurityScannerScanConfigScheduleOutputReference
	ScheduleInput() *GoogleSecurityScannerScanConfigSchedule
	StartingUrls() *[]*string
	SetStartingUrls(val *[]*string)
	StartingUrlsInput() *[]*string
	TargetPlatforms() *[]*string
	SetTargetPlatforms(val *[]*string)
	TargetPlatformsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleSecurityScannerScanConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserAgent() *string
	SetUserAgent(val *string)
	UserAgentInput() *string
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
	PutAuthentication(value *GoogleSecurityScannerScanConfigAuthentication)
	PutSchedule(value *GoogleSecurityScannerScanConfigSchedule)
	PutTimeouts(value *GoogleSecurityScannerScanConfigTimeouts)
	ResetAuthentication()
	ResetBlacklistPatterns()
	ResetExportToSecurityCommandCenter()
	ResetId()
	ResetMaxQps()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSchedule()
	ResetTargetPlatforms()
	ResetTimeouts()
	ResetUserAgent()
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

// The jsii proxy struct for GoogleSecurityScannerScanConfig
type jsiiProxy_GoogleSecurityScannerScanConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Authentication() GoogleSecurityScannerScanConfigAuthenticationOutputReference {
	var returns GoogleSecurityScannerScanConfigAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) AuthenticationInput() *GoogleSecurityScannerScanConfigAuthentication {
	var returns *GoogleSecurityScannerScanConfigAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) BlacklistPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blacklistPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) BlacklistPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blacklistPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) ExportToSecurityCommandCenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToSecurityCommandCenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) ExportToSecurityCommandCenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToSecurityCommandCenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) MaxQps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) MaxQpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Schedule() GoogleSecurityScannerScanConfigScheduleOutputReference {
	var returns GoogleSecurityScannerScanConfigScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) ScheduleInput() *GoogleSecurityScannerScanConfigSchedule {
	var returns *GoogleSecurityScannerScanConfigSchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) StartingUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startingUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) StartingUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startingUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) TargetPlatforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetPlatforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) TargetPlatformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetPlatformsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) Timeouts() GoogleSecurityScannerScanConfigTimeoutsOutputReference {
	var returns GoogleSecurityScannerScanConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) UserAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig) UserAgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_security_scanner_scan_config google_security_scanner_scan_config} Resource.
func NewGoogleSecurityScannerScanConfig(scope constructs.Construct, id *string, config *GoogleSecurityScannerScanConfigConfig) GoogleSecurityScannerScanConfig {
	_init_.Initialize()

	if err := validateNewGoogleSecurityScannerScanConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSecurityScannerScanConfig{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleSecurityScannerScanConfig.GoogleSecurityScannerScanConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_security_scanner_scan_config google_security_scanner_scan_config} Resource.
func NewGoogleSecurityScannerScanConfig_Override(g GoogleSecurityScannerScanConfig, scope constructs.Construct, id *string, config *GoogleSecurityScannerScanConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleSecurityScannerScanConfig.GoogleSecurityScannerScanConfig",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetBlacklistPatterns(val *[]*string) {
	if err := j.validateSetBlacklistPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blacklistPatterns",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetExportToSecurityCommandCenter(val *string) {
	if err := j.validateSetExportToSecurityCommandCenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToSecurityCommandCenter",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetMaxQps(val *float64) {
	if err := j.validateSetMaxQpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxQps",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetStartingUrls(val *[]*string) {
	if err := j.validateSetStartingUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingUrls",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetTargetPlatforms(val *[]*string) {
	if err := j.validateSetTargetPlatformsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPlatforms",
		val,
	)
}

func (j *jsiiProxy_GoogleSecurityScannerScanConfig)SetUserAgent(val *string) {
	if err := j.validateSetUserAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userAgent",
		val,
	)
}

// Generates CDKTF code for importing a GoogleSecurityScannerScanConfig resource upon running "cdktf plan <stack-name>".
func GoogleSecurityScannerScanConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleSecurityScannerScanConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleSecurityScannerScanConfig.GoogleSecurityScannerScanConfig",
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
func GoogleSecurityScannerScanConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSecurityScannerScanConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleSecurityScannerScanConfig.GoogleSecurityScannerScanConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleSecurityScannerScanConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSecurityScannerScanConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleSecurityScannerScanConfig.GoogleSecurityScannerScanConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleSecurityScannerScanConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSecurityScannerScanConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleSecurityScannerScanConfig.GoogleSecurityScannerScanConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleSecurityScannerScanConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleSecurityScannerScanConfig.GoogleSecurityScannerScanConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) PutAuthentication(value *GoogleSecurityScannerScanConfigAuthentication) {
	if err := g.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) PutSchedule(value *GoogleSecurityScannerScanConfigSchedule) {
	if err := g.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchedule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) PutTimeouts(value *GoogleSecurityScannerScanConfigTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ResetAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ResetBlacklistPatterns() {
	_jsii_.InvokeVoid(
		g,
		"resetBlacklistPatterns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ResetExportToSecurityCommandCenter() {
	_jsii_.InvokeVoid(
		g,
		"resetExportToSecurityCommandCenter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ResetMaxQps() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxQps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ResetSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ResetTargetPlatforms() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetPlatforms",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ResetUserAgent() {
	_jsii_.InvokeVoid(
		g,
		"resetUserAgent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecurityScannerScanConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

