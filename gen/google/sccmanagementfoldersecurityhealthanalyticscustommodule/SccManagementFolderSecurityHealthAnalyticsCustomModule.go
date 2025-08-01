package sccmanagementfoldersecurityhealthanalyticscustommodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/sccmanagementfoldersecurityhealthanalyticscustommodule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/scc_management_folder_security_health_analytics_custom_module google_scc_management_folder_security_health_analytics_custom_module}.
type SccManagementFolderSecurityHealthAnalyticsCustomModule interface {
	cdktf.TerraformResource
	AncestorModule() *string
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
	CustomConfig() SccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference
	CustomConfigInput() *SccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfig
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EnablementState() *string
	SetEnablementState(val *string)
	EnablementStateInput() *string
	Folder() *string
	SetFolder(val *string)
	FolderInput() *string
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
	LastEditor() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SccManagementFolderSecurityHealthAnalyticsCustomModuleTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
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
	PutCustomConfig(value *SccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfig)
	PutTimeouts(value *SccManagementFolderSecurityHealthAnalyticsCustomModuleTimeouts)
	ResetCustomConfig()
	ResetDisplayName()
	ResetEnablementState()
	ResetId()
	ResetLocation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
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

// The jsii proxy struct for SccManagementFolderSecurityHealthAnalyticsCustomModule
type jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) AncestorModule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ancestorModule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) CustomConfig() SccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference {
	var returns SccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference
	_jsii_.Get(
		j,
		"customConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) CustomConfigInput() *SccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfig {
	var returns *SccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfig
	_jsii_.Get(
		j,
		"customConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) EnablementState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablementState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) EnablementStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablementStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) LastEditor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastEditor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) Timeouts() SccManagementFolderSecurityHealthAnalyticsCustomModuleTimeoutsOutputReference {
	var returns SccManagementFolderSecurityHealthAnalyticsCustomModuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/scc_management_folder_security_health_analytics_custom_module google_scc_management_folder_security_health_analytics_custom_module} Resource.
func NewSccManagementFolderSecurityHealthAnalyticsCustomModule(scope constructs.Construct, id *string, config *SccManagementFolderSecurityHealthAnalyticsCustomModuleConfig) SccManagementFolderSecurityHealthAnalyticsCustomModule {
	_init_.Initialize()

	if err := validateNewSccManagementFolderSecurityHealthAnalyticsCustomModuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule{}

	_jsii_.Create(
		"@cdktf/provider-google.sccManagementFolderSecurityHealthAnalyticsCustomModule.SccManagementFolderSecurityHealthAnalyticsCustomModule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/scc_management_folder_security_health_analytics_custom_module google_scc_management_folder_security_health_analytics_custom_module} Resource.
func NewSccManagementFolderSecurityHealthAnalyticsCustomModule_Override(s SccManagementFolderSecurityHealthAnalyticsCustomModule, scope constructs.Construct, id *string, config *SccManagementFolderSecurityHealthAnalyticsCustomModuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.sccManagementFolderSecurityHealthAnalyticsCustomModule.SccManagementFolderSecurityHealthAnalyticsCustomModule",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetEnablementState(val *string) {
	if err := j.validateSetEnablementStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablementState",
		val,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a SccManagementFolderSecurityHealthAnalyticsCustomModule resource upon running "cdktf plan <stack-name>".
func SccManagementFolderSecurityHealthAnalyticsCustomModule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSccManagementFolderSecurityHealthAnalyticsCustomModule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.sccManagementFolderSecurityHealthAnalyticsCustomModule.SccManagementFolderSecurityHealthAnalyticsCustomModule",
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
func SccManagementFolderSecurityHealthAnalyticsCustomModule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSccManagementFolderSecurityHealthAnalyticsCustomModule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.sccManagementFolderSecurityHealthAnalyticsCustomModule.SccManagementFolderSecurityHealthAnalyticsCustomModule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SccManagementFolderSecurityHealthAnalyticsCustomModule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSccManagementFolderSecurityHealthAnalyticsCustomModule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.sccManagementFolderSecurityHealthAnalyticsCustomModule.SccManagementFolderSecurityHealthAnalyticsCustomModule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SccManagementFolderSecurityHealthAnalyticsCustomModule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSccManagementFolderSecurityHealthAnalyticsCustomModule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.sccManagementFolderSecurityHealthAnalyticsCustomModule.SccManagementFolderSecurityHealthAnalyticsCustomModule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SccManagementFolderSecurityHealthAnalyticsCustomModule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.sccManagementFolderSecurityHealthAnalyticsCustomModule.SccManagementFolderSecurityHealthAnalyticsCustomModule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) PutCustomConfig(value *SccManagementFolderSecurityHealthAnalyticsCustomModuleCustomConfig) {
	if err := s.validatePutCustomConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) PutTimeouts(value *SccManagementFolderSecurityHealthAnalyticsCustomModuleTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ResetCustomConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ResetDisplayName() {
	_jsii_.InvokeVoid(
		s,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ResetEnablementState() {
	_jsii_.InvokeVoid(
		s,
		"resetEnablementState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ResetLocation() {
	_jsii_.InvokeVoid(
		s,
		"resetLocation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementFolderSecurityHealthAnalyticsCustomModule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

