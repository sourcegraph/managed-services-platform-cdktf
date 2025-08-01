package colabruntimetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/colabruntimetemplate/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template google_colab_runtime_template}.
type ColabRuntimeTemplate interface {
	cdktf.TerraformResource
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
	DataPersistentDiskSpec() ColabRuntimeTemplateDataPersistentDiskSpecOutputReference
	DataPersistentDiskSpecInput() *ColabRuntimeTemplateDataPersistentDiskSpec
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktf.StringMap
	EncryptionSpec() ColabRuntimeTemplateEncryptionSpecOutputReference
	EncryptionSpecInput() *ColabRuntimeTemplateEncryptionSpec
	EucConfig() ColabRuntimeTemplateEucConfigOutputReference
	EucConfigInput() *ColabRuntimeTemplateEucConfig
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
	IdleShutdownConfig() ColabRuntimeTemplateIdleShutdownConfigOutputReference
	IdleShutdownConfigInput() *ColabRuntimeTemplateIdleShutdownConfig
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MachineSpec() ColabRuntimeTemplateMachineSpecOutputReference
	MachineSpecInput() *ColabRuntimeTemplateMachineSpec
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkSpec() ColabRuntimeTemplateNetworkSpecOutputReference
	NetworkSpecInput() *ColabRuntimeTemplateNetworkSpec
	NetworkTags() *[]*string
	SetNetworkTags(val *[]*string)
	NetworkTagsInput() *[]*string
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
	ShieldedVmConfig() ColabRuntimeTemplateShieldedVmConfigOutputReference
	ShieldedVmConfigInput() *ColabRuntimeTemplateShieldedVmConfig
	SoftwareConfig() ColabRuntimeTemplateSoftwareConfigOutputReference
	SoftwareConfigInput() *ColabRuntimeTemplateSoftwareConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ColabRuntimeTemplateTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutDataPersistentDiskSpec(value *ColabRuntimeTemplateDataPersistentDiskSpec)
	PutEncryptionSpec(value *ColabRuntimeTemplateEncryptionSpec)
	PutEucConfig(value *ColabRuntimeTemplateEucConfig)
	PutIdleShutdownConfig(value *ColabRuntimeTemplateIdleShutdownConfig)
	PutMachineSpec(value *ColabRuntimeTemplateMachineSpec)
	PutNetworkSpec(value *ColabRuntimeTemplateNetworkSpec)
	PutShieldedVmConfig(value *ColabRuntimeTemplateShieldedVmConfig)
	PutSoftwareConfig(value *ColabRuntimeTemplateSoftwareConfig)
	PutTimeouts(value *ColabRuntimeTemplateTimeouts)
	ResetDataPersistentDiskSpec()
	ResetDescription()
	ResetEncryptionSpec()
	ResetEucConfig()
	ResetId()
	ResetIdleShutdownConfig()
	ResetLabels()
	ResetMachineSpec()
	ResetName()
	ResetNetworkSpec()
	ResetNetworkTags()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetShieldedVmConfig()
	ResetSoftwareConfig()
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

// The jsii proxy struct for ColabRuntimeTemplate
type jsiiProxy_ColabRuntimeTemplate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ColabRuntimeTemplate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) DataPersistentDiskSpec() ColabRuntimeTemplateDataPersistentDiskSpecOutputReference {
	var returns ColabRuntimeTemplateDataPersistentDiskSpecOutputReference
	_jsii_.Get(
		j,
		"dataPersistentDiskSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) DataPersistentDiskSpecInput() *ColabRuntimeTemplateDataPersistentDiskSpec {
	var returns *ColabRuntimeTemplateDataPersistentDiskSpec
	_jsii_.Get(
		j,
		"dataPersistentDiskSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) EncryptionSpec() ColabRuntimeTemplateEncryptionSpecOutputReference {
	var returns ColabRuntimeTemplateEncryptionSpecOutputReference
	_jsii_.Get(
		j,
		"encryptionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) EncryptionSpecInput() *ColabRuntimeTemplateEncryptionSpec {
	var returns *ColabRuntimeTemplateEncryptionSpec
	_jsii_.Get(
		j,
		"encryptionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) EucConfig() ColabRuntimeTemplateEucConfigOutputReference {
	var returns ColabRuntimeTemplateEucConfigOutputReference
	_jsii_.Get(
		j,
		"eucConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) EucConfigInput() *ColabRuntimeTemplateEucConfig {
	var returns *ColabRuntimeTemplateEucConfig
	_jsii_.Get(
		j,
		"eucConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) IdleShutdownConfig() ColabRuntimeTemplateIdleShutdownConfigOutputReference {
	var returns ColabRuntimeTemplateIdleShutdownConfigOutputReference
	_jsii_.Get(
		j,
		"idleShutdownConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) IdleShutdownConfigInput() *ColabRuntimeTemplateIdleShutdownConfig {
	var returns *ColabRuntimeTemplateIdleShutdownConfig
	_jsii_.Get(
		j,
		"idleShutdownConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) MachineSpec() ColabRuntimeTemplateMachineSpecOutputReference {
	var returns ColabRuntimeTemplateMachineSpecOutputReference
	_jsii_.Get(
		j,
		"machineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) MachineSpecInput() *ColabRuntimeTemplateMachineSpec {
	var returns *ColabRuntimeTemplateMachineSpec
	_jsii_.Get(
		j,
		"machineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) NetworkSpec() ColabRuntimeTemplateNetworkSpecOutputReference {
	var returns ColabRuntimeTemplateNetworkSpecOutputReference
	_jsii_.Get(
		j,
		"networkSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) NetworkSpecInput() *ColabRuntimeTemplateNetworkSpec {
	var returns *ColabRuntimeTemplateNetworkSpec
	_jsii_.Get(
		j,
		"networkSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) NetworkTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) NetworkTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) ShieldedVmConfig() ColabRuntimeTemplateShieldedVmConfigOutputReference {
	var returns ColabRuntimeTemplateShieldedVmConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedVmConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) ShieldedVmConfigInput() *ColabRuntimeTemplateShieldedVmConfig {
	var returns *ColabRuntimeTemplateShieldedVmConfig
	_jsii_.Get(
		j,
		"shieldedVmConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) SoftwareConfig() ColabRuntimeTemplateSoftwareConfigOutputReference {
	var returns ColabRuntimeTemplateSoftwareConfigOutputReference
	_jsii_.Get(
		j,
		"softwareConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) SoftwareConfigInput() *ColabRuntimeTemplateSoftwareConfig {
	var returns *ColabRuntimeTemplateSoftwareConfig
	_jsii_.Get(
		j,
		"softwareConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) Timeouts() ColabRuntimeTemplateTimeoutsOutputReference {
	var returns ColabRuntimeTemplateTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabRuntimeTemplate) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template google_colab_runtime_template} Resource.
func NewColabRuntimeTemplate(scope constructs.Construct, id *string, config *ColabRuntimeTemplateConfig) ColabRuntimeTemplate {
	_init_.Initialize()

	if err := validateNewColabRuntimeTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ColabRuntimeTemplate{}

	_jsii_.Create(
		"@cdktf/provider-google.colabRuntimeTemplate.ColabRuntimeTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template google_colab_runtime_template} Resource.
func NewColabRuntimeTemplate_Override(c ColabRuntimeTemplate, scope constructs.Construct, id *string, config *ColabRuntimeTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.colabRuntimeTemplate.ColabRuntimeTemplate",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetNetworkTags(val *[]*string) {
	if err := j.validateSetNetworkTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkTags",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ColabRuntimeTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ColabRuntimeTemplate resource upon running "cdktf plan <stack-name>".
func ColabRuntimeTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateColabRuntimeTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.colabRuntimeTemplate.ColabRuntimeTemplate",
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
func ColabRuntimeTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateColabRuntimeTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.colabRuntimeTemplate.ColabRuntimeTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ColabRuntimeTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateColabRuntimeTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.colabRuntimeTemplate.ColabRuntimeTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ColabRuntimeTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateColabRuntimeTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.colabRuntimeTemplate.ColabRuntimeTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ColabRuntimeTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.colabRuntimeTemplate.ColabRuntimeTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) PutDataPersistentDiskSpec(value *ColabRuntimeTemplateDataPersistentDiskSpec) {
	if err := c.validatePutDataPersistentDiskSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataPersistentDiskSpec",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) PutEncryptionSpec(value *ColabRuntimeTemplateEncryptionSpec) {
	if err := c.validatePutEncryptionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEncryptionSpec",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) PutEucConfig(value *ColabRuntimeTemplateEucConfig) {
	if err := c.validatePutEucConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEucConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) PutIdleShutdownConfig(value *ColabRuntimeTemplateIdleShutdownConfig) {
	if err := c.validatePutIdleShutdownConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIdleShutdownConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) PutMachineSpec(value *ColabRuntimeTemplateMachineSpec) {
	if err := c.validatePutMachineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMachineSpec",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) PutNetworkSpec(value *ColabRuntimeTemplateNetworkSpec) {
	if err := c.validatePutNetworkSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkSpec",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) PutShieldedVmConfig(value *ColabRuntimeTemplateShieldedVmConfig) {
	if err := c.validatePutShieldedVmConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putShieldedVmConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) PutSoftwareConfig(value *ColabRuntimeTemplateSoftwareConfig) {
	if err := c.validatePutSoftwareConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSoftwareConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) PutTimeouts(value *ColabRuntimeTemplateTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetDataPersistentDiskSpec() {
	_jsii_.InvokeVoid(
		c,
		"resetDataPersistentDiskSpec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetEncryptionSpec() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionSpec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetEucConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetEucConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetIdleShutdownConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetIdleShutdownConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetMachineSpec() {
	_jsii_.InvokeVoid(
		c,
		"resetMachineSpec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetNetworkSpec() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkSpec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetNetworkTags() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetShieldedVmConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetShieldedVmConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetSoftwareConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSoftwareConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabRuntimeTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabRuntimeTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

