package workersscript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/workersscript/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/workers_script cloudflare_workers_script}.
type WorkersScript interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AnalyticsEngineBinding() WorkersScriptAnalyticsEngineBindingList
	AnalyticsEngineBindingInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CompatibilityDate() *string
	SetCompatibilityDate(val *string)
	CompatibilityDateInput() *string
	CompatibilityFlags() *[]*string
	SetCompatibilityFlags(val *[]*string)
	CompatibilityFlagsInput() *[]*string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	D1DatabaseBinding() WorkersScriptD1DatabaseBindingList
	D1DatabaseBindingInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DispatchNamespace() *string
	SetDispatchNamespace(val *string)
	DispatchNamespaceInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HyperdriveConfigBinding() WorkersScriptHyperdriveConfigBindingList
	HyperdriveConfigBindingInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	KvNamespaceBinding() WorkersScriptKvNamespaceBindingList
	KvNamespaceBindingInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logpush() interface{}
	SetLogpush(val interface{})
	LogpushInput() interface{}
	Module() interface{}
	SetModule(val interface{})
	ModuleInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Placement() WorkersScriptPlacementList
	PlacementInput() interface{}
	PlainTextBinding() WorkersScriptPlainTextBindingList
	PlainTextBindingInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueueBinding() WorkersScriptQueueBindingList
	QueueBindingInput() interface{}
	R2BucketBinding() WorkersScriptR2BucketBindingList
	R2BucketBindingInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	SecretTextBinding() WorkersScriptSecretTextBindingList
	SecretTextBindingInput() interface{}
	ServiceBinding() WorkersScriptServiceBindingList
	ServiceBindingInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WebassemblyBinding() WorkersScriptWebassemblyBindingList
	WebassemblyBindingInput() interface{}
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
	PutAnalyticsEngineBinding(value interface{})
	PutD1DatabaseBinding(value interface{})
	PutHyperdriveConfigBinding(value interface{})
	PutKvNamespaceBinding(value interface{})
	PutPlacement(value interface{})
	PutPlainTextBinding(value interface{})
	PutQueueBinding(value interface{})
	PutR2BucketBinding(value interface{})
	PutSecretTextBinding(value interface{})
	PutServiceBinding(value interface{})
	PutWebassemblyBinding(value interface{})
	ResetAnalyticsEngineBinding()
	ResetCompatibilityDate()
	ResetCompatibilityFlags()
	ResetD1DatabaseBinding()
	ResetDispatchNamespace()
	ResetHyperdriveConfigBinding()
	ResetId()
	ResetKvNamespaceBinding()
	ResetLogpush()
	ResetModule()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacement()
	ResetPlainTextBinding()
	ResetQueueBinding()
	ResetR2BucketBinding()
	ResetSecretTextBinding()
	ResetServiceBinding()
	ResetTags()
	ResetWebassemblyBinding()
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

// The jsii proxy struct for WorkersScript
type jsiiProxy_WorkersScript struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WorkersScript) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) AnalyticsEngineBinding() WorkersScriptAnalyticsEngineBindingList {
	var returns WorkersScriptAnalyticsEngineBindingList
	_jsii_.Get(
		j,
		"analyticsEngineBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) AnalyticsEngineBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analyticsEngineBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) CompatibilityDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) CompatibilityDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) CompatibilityFlags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) CompatibilityFlagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibilityFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) D1DatabaseBinding() WorkersScriptD1DatabaseBindingList {
	var returns WorkersScriptD1DatabaseBindingList
	_jsii_.Get(
		j,
		"d1DatabaseBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) D1DatabaseBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"d1DatabaseBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) DispatchNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dispatchNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) DispatchNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dispatchNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) HyperdriveConfigBinding() WorkersScriptHyperdriveConfigBindingList {
	var returns WorkersScriptHyperdriveConfigBindingList
	_jsii_.Get(
		j,
		"hyperdriveConfigBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) HyperdriveConfigBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hyperdriveConfigBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) KvNamespaceBinding() WorkersScriptKvNamespaceBindingList {
	var returns WorkersScriptKvNamespaceBindingList
	_jsii_.Get(
		j,
		"kvNamespaceBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) KvNamespaceBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kvNamespaceBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Logpush() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logpush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) LogpushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logpushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Module() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"module",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ModuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"moduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Placement() WorkersScriptPlacementList {
	var returns WorkersScriptPlacementList
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) PlacementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) PlainTextBinding() WorkersScriptPlainTextBindingList {
	var returns WorkersScriptPlainTextBindingList
	_jsii_.Get(
		j,
		"plainTextBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) PlainTextBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plainTextBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) QueueBinding() WorkersScriptQueueBindingList {
	var returns WorkersScriptQueueBindingList
	_jsii_.Get(
		j,
		"queueBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) QueueBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queueBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) R2BucketBinding() WorkersScriptR2BucketBindingList {
	var returns WorkersScriptR2BucketBindingList
	_jsii_.Get(
		j,
		"r2BucketBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) R2BucketBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"r2BucketBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) SecretTextBinding() WorkersScriptSecretTextBindingList {
	var returns WorkersScriptSecretTextBindingList
	_jsii_.Get(
		j,
		"secretTextBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) SecretTextBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretTextBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ServiceBinding() WorkersScriptServiceBindingList {
	var returns WorkersScriptServiceBindingList
	_jsii_.Get(
		j,
		"serviceBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) ServiceBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) WebassemblyBinding() WorkersScriptWebassemblyBindingList {
	var returns WorkersScriptWebassemblyBindingList
	_jsii_.Get(
		j,
		"webassemblyBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkersScript) WebassemblyBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webassemblyBindingInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/workers_script cloudflare_workers_script} Resource.
func NewWorkersScript(scope constructs.Construct, id *string, config *WorkersScriptConfig) WorkersScript {
	_init_.Initialize()

	if err := validateNewWorkersScriptParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkersScript{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/workers_script cloudflare_workers_script} Resource.
func NewWorkersScript_Override(w WorkersScript, scope constructs.Construct, id *string, config *WorkersScriptConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkersScript)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetCompatibilityDate(val *string) {
	if err := j.validateSetCompatibilityDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityDate",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetCompatibilityFlags(val *[]*string) {
	if err := j.validateSetCompatibilityFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityFlags",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetDispatchNamespace(val *string) {
	if err := j.validateSetDispatchNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dispatchNamespace",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetLogpush(val interface{}) {
	if err := j.validateSetLogpushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logpush",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetModule(val interface{}) {
	if err := j.validateSetModuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"module",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WorkersScript)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a WorkersScript resource upon running "cdktf plan <stack-name>".
func WorkersScript_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWorkersScript_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
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
func WorkersScript_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkersScript_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkersScript_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkersScript_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkersScript_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkersScript_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkersScript_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.workersScript.WorkersScript",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WorkersScript) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WorkersScript) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WorkersScript) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WorkersScript) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkersScript) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WorkersScript) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkersScript) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkersScript) PutAnalyticsEngineBinding(value interface{}) {
	if err := w.validatePutAnalyticsEngineBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAnalyticsEngineBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutD1DatabaseBinding(value interface{}) {
	if err := w.validatePutD1DatabaseBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putD1DatabaseBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutHyperdriveConfigBinding(value interface{}) {
	if err := w.validatePutHyperdriveConfigBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHyperdriveConfigBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutKvNamespaceBinding(value interface{}) {
	if err := w.validatePutKvNamespaceBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putKvNamespaceBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutPlacement(value interface{}) {
	if err := w.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putPlacement",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutPlainTextBinding(value interface{}) {
	if err := w.validatePutPlainTextBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putPlainTextBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutQueueBinding(value interface{}) {
	if err := w.validatePutQueueBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueueBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutR2BucketBinding(value interface{}) {
	if err := w.validatePutR2BucketBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putR2BucketBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutSecretTextBinding(value interface{}) {
	if err := w.validatePutSecretTextBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSecretTextBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutServiceBinding(value interface{}) {
	if err := w.validatePutServiceBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putServiceBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) PutWebassemblyBinding(value interface{}) {
	if err := w.validatePutWebassemblyBindingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putWebassemblyBinding",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkersScript) ResetAnalyticsEngineBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetAnalyticsEngineBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetCompatibilityDate() {
	_jsii_.InvokeVoid(
		w,
		"resetCompatibilityDate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetCompatibilityFlags() {
	_jsii_.InvokeVoid(
		w,
		"resetCompatibilityFlags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetD1DatabaseBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetD1DatabaseBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetDispatchNamespace() {
	_jsii_.InvokeVoid(
		w,
		"resetDispatchNamespace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetHyperdriveConfigBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetHyperdriveConfigBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetKvNamespaceBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetKvNamespaceBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetLogpush() {
	_jsii_.InvokeVoid(
		w,
		"resetLogpush",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetModule() {
	_jsii_.InvokeVoid(
		w,
		"resetModule",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetPlacement() {
	_jsii_.InvokeVoid(
		w,
		"resetPlacement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetPlainTextBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetPlainTextBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetQueueBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetQueueBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetR2BucketBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetR2BucketBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetSecretTextBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetSecretTextBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetServiceBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetServiceBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) ResetWebassemblyBinding() {
	_jsii_.InvokeVoid(
		w,
		"resetWebassemblyBinding",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkersScript) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkersScript) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

