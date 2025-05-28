package datatferegistrymodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/tfe/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/tfe/datatferegistrymodule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module tfe_registry_module}.
type DataTfeRegistryModule interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModuleProvider() *string
	SetModuleProvider(val *string)
	ModuleProviderInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	NoCode() cdktf.IResolvable
	NoCodeModuleId() *string
	NoCodeModuleSource() *string
	// The tree node.
	Node() constructs.Node
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	Permissions() DataTfeRegistryModulePermissionsList
	PermissionsInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	PublishingMechanism() *string
	// Experimental.
	RawOverrides() interface{}
	RegistryName() *string
	SetRegistryName(val *string)
	RegistryNameInput() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TestConfig() DataTfeRegistryModuleTestConfigList
	TestConfigInput() interface{}
	UpdatedAt() *string
	VcsRepo() DataTfeRegistryModuleVcsRepoList
	VcsRepoInput() interface{}
	VersionStatuses() DataTfeRegistryModuleVersionStatusesList
	VersionStatusesInput() interface{}
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
	PutPermissions(value interface{})
	PutTestConfig(value interface{})
	PutVcsRepo(value interface{})
	PutVersionStatuses(value interface{})
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermissions()
	ResetRegistryName()
	ResetTestConfig()
	ResetVcsRepo()
	ResetVersionStatuses()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataTfeRegistryModule
type jsiiProxy_DataTfeRegistryModule struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataTfeRegistryModule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) ModuleProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"moduleProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) ModuleProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"moduleProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) NoCode() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"noCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) NoCodeModuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noCodeModuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) NoCodeModuleSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noCodeModuleSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) Permissions() DataTfeRegistryModulePermissionsList {
	var returns DataTfeRegistryModulePermissionsList
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) PermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) PublishingMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishingMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) RegistryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) RegistryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) TestConfig() DataTfeRegistryModuleTestConfigList {
	var returns DataTfeRegistryModuleTestConfigList
	_jsii_.Get(
		j,
		"testConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) TestConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"testConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) VcsRepo() DataTfeRegistryModuleVcsRepoList {
	var returns DataTfeRegistryModuleVcsRepoList
	_jsii_.Get(
		j,
		"vcsRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) VcsRepoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vcsRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) VersionStatuses() DataTfeRegistryModuleVersionStatusesList {
	var returns DataTfeRegistryModuleVersionStatusesList
	_jsii_.Get(
		j,
		"versionStatuses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeRegistryModule) VersionStatusesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionStatusesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module tfe_registry_module} Data Source.
func NewDataTfeRegistryModule(scope constructs.Construct, id *string, config *DataTfeRegistryModuleConfig) DataTfeRegistryModule {
	_init_.Initialize()

	if err := validateNewDataTfeRegistryModuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfeRegistryModule{}

	_jsii_.Create(
		"@cdktf/provider-tfe.dataTfeRegistryModule.DataTfeRegistryModule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/data-sources/registry_module tfe_registry_module} Data Source.
func NewDataTfeRegistryModule_Override(d DataTfeRegistryModule, scope constructs.Construct, id *string, config *DataTfeRegistryModuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tfe.dataTfeRegistryModule.DataTfeRegistryModule",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataTfeRegistryModule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataTfeRegistryModule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataTfeRegistryModule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataTfeRegistryModule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataTfeRegistryModule)SetModuleProvider(val *string) {
	if err := j.validateSetModuleProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"moduleProvider",
		val,
	)
}

func (j *jsiiProxy_DataTfeRegistryModule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataTfeRegistryModule)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataTfeRegistryModule)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_DataTfeRegistryModule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataTfeRegistryModule)SetRegistryName(val *string) {
	if err := j.validateSetRegistryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registryName",
		val,
	)
}

// Generates CDKTF code for importing a DataTfeRegistryModule resource upon running "cdktf plan <stack-name>".
func DataTfeRegistryModule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataTfeRegistryModule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.dataTfeRegistryModule.DataTfeRegistryModule",
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
func DataTfeRegistryModule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfeRegistryModule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.dataTfeRegistryModule.DataTfeRegistryModule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfeRegistryModule_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfeRegistryModule_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.dataTfeRegistryModule.DataTfeRegistryModule",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfeRegistryModule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfeRegistryModule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-tfe.dataTfeRegistryModule.DataTfeRegistryModule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTfeRegistryModule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-tfe.dataTfeRegistryModule.DataTfeRegistryModule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTfeRegistryModule) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfeRegistryModule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataTfeRegistryModule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfeRegistryModule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfeRegistryModule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfeRegistryModule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfeRegistryModule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfeRegistryModule) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfeRegistryModule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfeRegistryModule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataTfeRegistryModule) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) PutPermissions(value interface{}) {
	if err := d.validatePutPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPermissions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) PutTestConfig(value interface{}) {
	if err := d.validatePutTestConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTestConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) PutVcsRepo(value interface{}) {
	if err := d.validatePutVcsRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVcsRepo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) PutVersionStatuses(value interface{}) {
	if err := d.validatePutVersionStatusesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVersionStatuses",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) ResetPermissions() {
	_jsii_.InvokeVoid(
		d,
		"resetPermissions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) ResetRegistryName() {
	_jsii_.InvokeVoid(
		d,
		"resetRegistryName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) ResetTestConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetTestConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) ResetVcsRepo() {
	_jsii_.InvokeVoid(
		d,
		"resetVcsRepo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) ResetVersionStatuses() {
	_jsii_.InvokeVoid(
		d,
		"resetVersionStatuses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfeRegistryModule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeRegistryModule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeRegistryModule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeRegistryModule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeRegistryModule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeRegistryModule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

